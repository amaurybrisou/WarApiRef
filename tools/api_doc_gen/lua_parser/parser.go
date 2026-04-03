package lua_parser

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"roraddons/tools/api_doc_gen/graph"
)

type functionSpan struct {
	RawName           string
	ShortName         string
	Local             bool
	DeclarationKind   string
	ReceiverName      string
	ReceiverSeparator string
	Start             int
	BodyStart         int
	End               int
	Params            []string
	Line              int
	EndLine           int
}

type constValue struct {
	Kind string
	Text string
	Rows [][]string
}

type scanContext struct {
	addon           string
	file            string
	functionName    string
	aliasMap        map[string]string
	localNames      map[string]bool
	consts          map[string]constValue
	savedRoots      map[string]bool
	functionSymbols map[string]string
}

type loopInfo struct {
	ValueVar  string
	TableName string
	BodyStart int
	End       int
}

func ParseFile(addonName string, filePath string, manifest graph.Manifest) (graph.LuaFileResult, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return graph.LuaFileResult{}, fmt.Errorf("read lua file %s: %w", filePath, err)
	}

	tokens, err := Tokenize(string(bytes))
	if err != nil {
		return graph.LuaFileResult{}, fmt.Errorf("tokenize lua file %s: %w", filePath, err)
	}

	spans := findFunctionSpans(tokens)
	inFunction := buildFunctionMask(len(tokens), spans)
	aliases, locals, consts := collectScope(tokens, 0, len(tokens)-1, inFunction, nil, nil, nil)
	exports := findExportedLocalFunctions(tokens, spans, inFunction, aliases)
	savedRoots := make(map[string]bool, len(manifest.SavedVariables))
	for _, variable := range manifest.SavedVariables {
		savedRoots[graph.RootSegment(variable)] = true
	}

	modules := []graph.Module{{
		Addon: addonName,
		Name:  addonName,
		File:  graph.NormalizePath(filePath),
		Line:  1,
		Kind:  "addon-root",
	}}
	functions := make([]graph.Function, 0, len(spans))
	events := []graph.EventRegistration{}
	state := []graph.StateVariable{}

	functionSymbols := map[string]string{}
	for _, span := range spans {
		if span.Local {
			qualifiedName := qualifyLocalFunction(addonName, span.ShortName)
			if exported, ok := exports[span.ShortName]; ok {
				functionSymbols[span.ShortName] = exported
				functionSymbols[qualifiedName] = exported
			} else {
				functionSymbols[span.ShortName] = qualifiedName
				functionSymbols[qualifiedName] = qualifiedName
			}
		}
	}

	for _, span := range spans {
		functionName, aliasesForFunction, localFlag := resolveFunctionName(addonName, span, aliases, exports)
		functionAliasMap := copyStringMap(aliases)
		for shortName, resolved := range functionSymbols {
			functionAliasMap[shortName] = resolved
		}
		localNames := copyBoolMap(locals)
		for _, parameter := range span.Params {
			localNames[parameter] = true
		}
		functionConsts := copyConstMap(consts)
		functionAliasMap, localNames, functionConsts = collectScope(tokens, span.BodyStart, span.End, nil, functionAliasMap, localNames, functionConsts)
		calls, registrations, writes, discoveredModules := scanRange(tokens, span.BodyStart, span.End, nil, scanContext{
			addon:           addonName,
			file:            graph.NormalizePath(filePath),
			functionName:    functionName,
			aliasMap:        functionAliasMap,
			localNames:      localNames,
			consts:          functionConsts,
			savedRoots:      savedRoots,
			functionSymbols: functionSymbols,
		})

		function := graph.Function{
			Addon:             addonName,
			Name:              functionName,
			DeclaredName:      declaredFunctionName(span),
			ShortName:         span.ShortName,
			ScopeKind:         scopeKind(localFlag),
			DeclarationKind:   declarationKind(span),
			ReceiverName:      span.ReceiverName,
			ReceiverSeparator: span.ReceiverSeparator,
			Aliases:           aliasesForFunction,
			Module:            graph.OwnerName(functionName),
			File:              graph.NormalizePath(filePath),
			Line:              span.Line,
			EndLine:           span.EndLine,
			DeclarationOrder:  span.Start,
			Params:            span.Params,
			Calls:             calls,
			Events:            registrations,
			StateWrites:       stateNames(writes),
			Local:             localFlag,
			Kind:              functionKind(functionName),
		}
		functions = append(functions, function)
		events = append(events, registrations...)
		state = append(state, writes...)
		modules = append(modules, discoveredModules...)
		if function.Module != "" {
			modules = append(modules, graph.Module{
				Addon: addonName,
				Name:  function.Module,
				File:  graph.NormalizePath(filePath),
				Line:  span.Line,
				Kind:  "function-owner",
			})
		}
	}

	_, topLevelRegistrations, topLevelState, topLevelModules := scanRange(tokens, 0, len(tokens)-1, inFunction, scanContext{
		addon:           addonName,
		file:            graph.NormalizePath(filePath),
		aliasMap:        aliases,
		localNames:      locals,
		consts:          consts,
		savedRoots:      savedRoots,
		functionSymbols: functionSymbols,
	})
	events = append(events, topLevelRegistrations...)
	state = append(state, topLevelState...)
	modules = append(modules, topLevelModules...)
	for _, variable := range state {
		if variable.Owner != "" {
			modules = append(modules, graph.Module{
				Addon: variable.Addon,
				Name:  variable.Owner,
				File:  variable.File,
				Line:  variable.Line,
				Kind:  "state-owner",
			})
		}
	}

	return graph.LuaFileResult{
		Functions: graph.UniqueSortedFunctions(functions),
		Modules:   graph.UniqueSortedModules(modules),
		Events:    graph.UniqueSortedEvents(events),
		State:     graph.UniqueSortedState(state),
	}, nil
}

func findFunctionSpans(tokens []Token) []functionSpan {
	spans := []functionSpan{}
	for index := 0; index < len(tokens); index++ {
		functionTokenIndex := -1
		span := functionSpan{}
		switch {
		case tokens[index].IsKeyword("function"):
			functionTokenIndex = index
			span.DeclarationKind = "function_decl"
			if rawName, next, ok := readNameChain(tokens, index+1); ok {
				span.RawName = rawName
				span.ShortName = graph.LastSegment(rawName)
				span.ReceiverName, span.ReceiverSeparator = splitReceiver(rawName)
				if next < len(tokens) && tokens[next].IsSymbol("(") {
					closeIndex := matchPair(tokens, next, "(", ")")
					span.Params = parseParameters(tokens[next+1 : closeIndex])
					span.BodyStart = closeIndex + 1
				}
			}
		case tokens[index].IsKeyword("local") && index+1 < len(tokens) && tokens[index+1].IsKeyword("function"):
			functionTokenIndex = index + 1
			span.Local = true
			span.DeclarationKind = "local_function_decl"
			if index+2 < len(tokens) && tokens[index+2].Kind == tokenIdentifier {
				span.ShortName = tokens[index+2].Text
			}
			if index+3 < len(tokens) && tokens[index+3].IsSymbol("(") {
				closeIndex := matchPair(tokens, index+3, "(", ")")
				span.Params = parseParameters(tokens[index+4 : closeIndex])
				span.BodyStart = closeIndex + 1
			}
		case tokens[index].IsKeyword("local") && index+3 < len(tokens) && tokens[index+1].Kind == tokenIdentifier && tokens[index+2].IsSymbol("=") && tokens[index+3].IsKeyword("function"):
			functionTokenIndex = index + 3
			span.Local = true
			span.DeclarationKind = "assignment_function"
			span.ShortName = tokens[index+1].Text
			if index+4 < len(tokens) && tokens[index+4].IsSymbol("(") {
				closeIndex := matchPair(tokens, index+4, "(", ")")
				span.Params = parseParameters(tokens[index+5 : closeIndex])
				span.BodyStart = closeIndex + 1
			}
		default:
			if index > 0 && tokens[index-1].IsKeyword("local") {
				continue
			}
			rawName, next, ok := readNameChain(tokens, index)
			if ok && next+1 < len(tokens) && tokens[next].IsSymbol("=") && tokens[next+1].IsKeyword("function") {
				functionTokenIndex = next + 1
				span.DeclarationKind = "assignment_function"
				span.RawName = rawName
				span.ShortName = graph.LastSegment(rawName)
				span.ReceiverName, span.ReceiverSeparator = splitReceiver(rawName)
				if next+2 < len(tokens) && tokens[next+2].IsSymbol("(") {
					closeIndex := matchPair(tokens, next+2, "(", ")")
					span.Params = parseParameters(tokens[next+3 : closeIndex])
					span.BodyStart = closeIndex + 1
				}
			}
		}

		if functionTokenIndex < 0 || span.BodyStart <= 0 {
			continue
		}
		span.Start = functionTokenIndex
		span.Line = tokens[functionTokenIndex].Line
		span.End = findFunctionEnd(tokens, functionTokenIndex, span.BodyStart)
		span.EndLine = tokens[span.End].Line
		spans = append(spans, span)
	}

	sort.Slice(spans, func(i, j int) bool {
		if spans[i].Start == spans[j].Start {
			return spans[i].End < spans[j].End
		}
		return spans[i].Start < spans[j].Start
	})
	return spans
}

func buildFunctionMask(length int, spans []functionSpan) []bool {
	mask := make([]bool, length)
	for _, span := range spans {
		for index := span.Start; index <= span.End && index < length; index++ {
			mask[index] = true
		}
	}
	return mask
}

func collectScope(tokens []Token, start, end int, skip []bool, inheritedAliases map[string]string, inheritedLocals map[string]bool, inheritedConsts map[string]constValue) (map[string]string, map[string]bool, map[string]constValue) {
	aliases := copyStringMap(inheritedAliases)
	locals := copyBoolMap(inheritedLocals)
	consts := copyConstMap(inheritedConsts)

	for index := maxInt(start, 0); index <= end && index < len(tokens); index++ {
		if skip != nil && skip[index] {
			continue
		}
		if !tokens[index].IsKeyword("local") {
			continue
		}

		if index+1 < len(tokens) && tokens[index+1].IsKeyword("function") {
			if index+2 < len(tokens) && tokens[index+2].Kind == tokenIdentifier {
				locals[tokens[index+2].Text] = true
			}
			continue
		}

		names := []string{}
		cursor := index + 1
		for cursor <= end && cursor < len(tokens) {
			if tokens[cursor].Kind == tokenIdentifier {
				names = append(names, tokens[cursor].Text)
				locals[tokens[cursor].Text] = true
				cursor++
				continue
			}
			if tokens[cursor].IsSymbol(",") {
				cursor++
				continue
			}
			break
		}
		if len(names) == 0 || cursor > end || cursor >= len(tokens) || !tokens[cursor].IsSymbol("=") {
			continue
		}

		rhsStart := cursor + 1
		if rhsStart > end || rhsStart >= len(tokens) {
			continue
		}
		if len(names) == 1 {
			if name, _, ok := readNameChain(tokens, rhsStart); ok {
				aliases[names[0]] = resolveName(name, aliases)
			}
			if tokens[rhsStart].Kind == tokenString {
				consts[names[0]] = constValue{Kind: "string", Text: strings.Trim(tokens[rhsStart].Text, "\"'")}
				continue
			}
			if tokens[rhsStart].Kind == tokenNumber {
				consts[names[0]] = constValue{Kind: "number", Text: tokens[rhsStart].Text}
				continue
			}
			if tokens[rhsStart].IsKeyword("true") || tokens[rhsStart].IsKeyword("false") {
				consts[names[0]] = constValue{Kind: "boolean", Text: tokens[rhsStart].Text}
				continue
			}
			if tokens[rhsStart].IsKeyword("nil") {
				consts[names[0]] = constValue{Kind: "nil", Text: "nil"}
				continue
			}
			if tokens[rhsStart].IsSymbol("{") {
				closeIndex := matchPair(tokens, rhsStart, "{", "}")
				consts[names[0]] = parseTableLiteral(tokens[rhsStart : closeIndex+1])
				continue
			}
		}
	}

	return aliases, locals, consts
}

func findExportedLocalFunctions(tokens []Token, spans []functionSpan, skip []bool, aliases map[string]string) map[string]string {
	localFunctions := map[string]bool{}
	for _, span := range spans {
		if span.Local && span.ShortName != "" {
			localFunctions[span.ShortName] = true
		}
	}

	exports := map[string]string{}
	for index := 0; index < len(tokens)-2; index++ {
		if skip != nil && skip[index] {
			continue
		}
		name, next, ok := readNameChain(tokens, index)
		if !ok || next+1 >= len(tokens) || !tokens[next].IsSymbol("=") {
			continue
		}
		rhs := tokens[next+1]
		if rhs.Kind != tokenIdentifier || !localFunctions[rhs.Text] {
			continue
		}
		exports[rhs.Text] = resolveName(name, aliases)
	}
	return exports
}

func resolveFunctionName(addonName string, span functionSpan, aliases map[string]string, exports map[string]string) (string, []string, bool) {
	if span.Local {
		qualifiedLocal := qualifyLocalFunction(addonName, span.ShortName)
		if exported, ok := exports[span.ShortName]; ok {
			return exported, []string{qualifiedLocal}, false
		}
		return qualifiedLocal, nil, true
	}
	resolved := resolveName(span.RawName, aliases)
	if resolved == "" {
		resolved = span.RawName
	}
	return resolved, nil, false
}

func scanRange(tokens []Token, start, end int, skip []bool, context scanContext) ([]graph.Call, []graph.EventRegistration, []graph.StateVariable, []graph.Module) {
	calls := []graph.Call{}
	registrations := []graph.EventRegistration{}
	state := []graph.StateVariable{}
	modules := []graph.Module{}

	for index := maxInt(start, 0); index <= end && index < len(tokens); index++ {
		if skip != nil && skip[index] {
			continue
		}

		if tokens[index].IsKeyword("for") {
			if loop, ok := parseGenericForLoop(tokens, index); ok {
				registrations = append(registrations, expandLoopRegistrations(tokens, loop, context)...)
				index = loop.End
				continue
			}
		}

		name, next, ok := readNameChain(tokens, index)
		if !ok {
			continue
		}
		resolvedName := resolveName(name, context.aliasMap)
		root := graph.RootSegment(name)
		if context.localNames[root] && context.aliasMap[root] == "" && context.functionSymbols[root] == "" {
			continue
		}

		if next < len(tokens) && tokens[next].IsSymbol("(") {
			closeIndex := matchPair(tokens, next, "(", ")")
			argumentTokens := splitTopLevel(tokens[next+1:closeIndex], ",")
			arguments := make([]string, 0, len(argumentTokens))
			for _, group := range argumentTokens {
				arguments = append(arguments, strings.TrimSpace(expressionString(group)))
			}
			calls = append(calls, graph.Call{
				Name:           resolvedName,
				CalleeRaw:      name,
				CalleeResolved: resolvedName,
				Line:           tokens[index].Line,
				Arguments:      arguments,
			})
			registrations = append(registrations, parseEventRegistration(name, resolvedName, arguments, tokens[index].Line, context)...)
		}

		if next < len(tokens) && tokens[next].IsSymbol("=") {
			if assignmentExportsFunction(tokens, next+1, context) {
				continue
			}
			valueType := inferValueType(tokens, next+1)
			if capturesModule(resolvedName, valueType, context) {
				modules = append(modules, graph.Module{
					Addon: context.addon,
					Name:  resolvedName,
					File:  context.file,
					Line:  tokens[index].Line,
					Kind:  "table",
				})
			}
			if capturesState(resolvedName, context) {
				owner := graph.OwnerName(resolvedName)
				state = append(state, graph.StateVariable{
					Addon:     context.addon,
					Name:      resolvedName,
					Owner:     owner,
					File:      context.file,
					Line:      tokens[index].Line,
					ValueType: valueType,
					Saved:     context.savedRoots[graph.RootSegment(resolvedName)],
					Scope:     context.functionName,
				})
			}
		}

		if next-1 > index {
			index = next - 1
		}
	}

	return dedupeCalls(calls), graph.UniqueSortedEvents(registrations), graph.UniqueSortedState(state), graph.UniqueSortedModules(modules)
}

func parseEventRegistration(rawCallee string, callee string, arguments []string, line int, context scanContext) []graph.EventRegistration {
	resolve := func(value string) string {
		return resolveExpression(value, context)
	}
	resolveHandler := func(value string) string {
		resolved := resolveExpression(value, context)
		if target, ok := context.functionSymbols[resolved]; ok {
			return target
		}
		return resolved
	}
	result := []graph.EventRegistration{}
	appendRegistration := func(scope string, eventArgIndex int, handlerArgIndex int, windowArgIndex int) {
		if len(arguments) <= handlerArgIndex || len(arguments) <= eventArgIndex {
			return
		}
		eventRaw := strings.TrimSpace(arguments[eventArgIndex])
		handlerRaw := strings.TrimSpace(arguments[handlerArgIndex])
		eventName := resolve(arguments[eventArgIndex])
		handlerName := resolveHandler(arguments[handlerArgIndex])
		if strings.Contains(eventName, "[") || strings.Contains(handlerName, "[") || eventName == "" || handlerName == "" {
			return
		}
		windowName := ""
		windowRaw := ""
		if windowArgIndex >= 0 && len(arguments) > windowArgIndex {
			windowRaw = strings.TrimSpace(arguments[windowArgIndex])
			windowName = resolve(arguments[windowArgIndex])
		}
		result = append(result, graph.EventRegistration{
			Addon:           context.addon,
			Registrar:       callee,
			Event:           eventName,
			EventRaw:        eventRaw,
			EventResolved:   eventName,
			Handler:         handlerName,
			HandlerRaw:      handlerRaw,
			HandlerResolved: handlerName,
			Window:          windowName,
			WindowRaw:       windowRaw,
			WindowResolved:  windowName,
			Scope:           scope,
			SourceFunction:  context.functionName,
			File:            context.file,
			Line:            line,
		})
	}

	switch callee {
	case "RegisterEventHandler":
		appendRegistration("global", 0, 1, -1)
	case "WindowRegisterEventHandler":
		appendRegistration("window", 1, 2, 0)
	case "WindowRegisterCoreEventHandler":
		appendRegistration("core", 1, 2, 0)
	case "LibGroup.Register":
		appendRegistration("library", 0, 1, -1)
	default:
		_ = rawCallee
		if strings.HasSuffix(callee, ".Register") {
			appendRegistration("library", 0, len(arguments)-1, -1)
		}
		if strings.HasSuffix(callee, ".AddEventHandler") && len(arguments) >= 3 {
			appendRegistration("addon", 1, 2, -1)
		}
	}
	return result
}

func expandLoopRegistrations(tokens []Token, loop loopInfo, context scanContext) []graph.EventRegistration {
	constTable, ok := context.consts[loop.TableName]
	if !ok || constTable.Kind != "table" {
		return nil
	}
	registrations := []graph.EventRegistration{}
	for index := loop.BodyStart; index <= loop.End && index < len(tokens); index++ {
		name, next, ok := readNameChain(tokens, index)
		if !ok {
			continue
		}
		resolvedName := resolveName(name, context.aliasMap)
		if next >= len(tokens) || !tokens[next].IsSymbol("(") {
			continue
		}
		closeIndex := matchPair(tokens, next, "(", ")")
		argumentTokens := splitTopLevel(tokens[next+1:closeIndex], ",")
		arguments := make([]string, 0, len(argumentTokens))
		for _, group := range argumentTokens {
			arguments = append(arguments, strings.TrimSpace(expressionString(group)))
		}

		makeRegistration := func(eventValue string, handlerValue string) {
			registration := parseEventRegistration(resolvedName, resolvedName, []string{eventValue, handlerValue}, tokens[index].Line, context)
			if strings.EqualFold(resolvedName, "WindowRegisterEventHandler") || strings.EqualFold(resolvedName, "WindowRegisterCoreEventHandler") {
				return
			}
			registrations = append(registrations, registration...)
		}

		switch resolvedName {
		case "RegisterEventHandler":
			if len(arguments) < 2 || !matchesIndexedVar(arguments[0], loop.ValueVar, 1) || !matchesIndexedVar(arguments[1], loop.ValueVar, 2) {
				continue
			}
			for _, row := range constTable.Rows {
				if len(row) >= 2 {
					makeRegistration(row[0], row[1])
				}
			}
		case "LibGroup.Register":
			if len(arguments) < 2 || !matchesIndexedVar(arguments[0], loop.ValueVar, 1) || !matchesIndexedVar(arguments[1], loop.ValueVar, 2) {
				continue
			}
			for _, row := range constTable.Rows {
				if len(row) >= 2 {
					registrations = append(registrations, parseEventRegistration(resolvedName, resolvedName, []string{row[0], row[1]}, tokens[index].Line, context)...)
				}
			}
		}
	}
	return graph.UniqueSortedEvents(registrations)
}

func parseGenericForLoop(tokens []Token, start int) (loopInfo, bool) {
	if start >= len(tokens) || !tokens[start].IsKeyword("for") {
		return loopInfo{}, false
	}
	variables := []string{}
	index := start + 1
	for index < len(tokens) {
		if tokens[index].IsKeyword("in") {
			break
		}
		if tokens[index].Kind == tokenIdentifier {
			variables = append(variables, tokens[index].Text)
		}
		index++
	}
	if index >= len(tokens) || !tokens[index].IsKeyword("in") || len(variables) == 0 {
		return loopInfo{}, false
	}

	iteratorName, next, ok := readNameChain(tokens, index+1)
	if !ok || (iteratorName != "pairs" && iteratorName != "ipairs") || next >= len(tokens) || !tokens[next].IsSymbol("(") {
		return loopInfo{}, false
	}
	closeIndex := matchPair(tokens, next, "(", ")")
	arguments := splitTopLevel(tokens[next+1:closeIndex], ",")
	if len(arguments) != 1 {
		return loopInfo{}, false
	}
	tableName := strings.TrimSpace(expressionString(arguments[0]))
	if closeIndex+1 >= len(tokens) || !tokens[closeIndex+1].IsKeyword("do") {
		return loopInfo{}, false
	}
	endIndex := findBlockEnd(tokens, closeIndex+2)
	if endIndex <= closeIndex+1 {
		return loopInfo{}, false
	}
	return loopInfo{
		ValueVar:  variables[len(variables)-1],
		TableName: tableName,
		BodyStart: closeIndex + 2,
		End:       endIndex,
	}, true
}

func findFunctionEnd(tokens []Token, functionIndex int, bodyStart int) int {
	depth := 1
	for index := bodyStart; index < len(tokens); index++ {
		if tokens[index].Kind != tokenKeyword {
			continue
		}
		switch tokens[index].Text {
		case "function":
			if index != functionIndex {
				depth++
			}
		case "if", "for", "while", "repeat":
			depth++
		case "end", "until":
			depth--
			if depth == 0 {
				return index
			}
		}
	}
	return len(tokens) - 1
}

func findBlockEnd(tokens []Token, start int) int {
	depth := 1
	for index := start; index < len(tokens); index++ {
		if tokens[index].Kind != tokenKeyword {
			continue
		}
		switch tokens[index].Text {
		case "function", "if", "for", "while", "repeat":
			depth++
		case "end", "until":
			depth--
			if depth == 0 {
				return index
			}
		}
	}
	return len(tokens) - 1
}

func parseParameters(tokens []Token) []string {
	parameters := []string{}
	for _, token := range tokens {
		if token.Kind == tokenIdentifier || token.Text == "..." {
			parameters = append(parameters, token.Text)
		}
	}
	return parameters
}

func readNameChain(tokens []Token, start int) (string, int, bool) {
	if start >= len(tokens) || tokens[start].Kind != tokenIdentifier {
		return "", start, false
	}
	parts := []string{tokens[start].Text}
	index := start + 1
	for index+1 < len(tokens) {
		switch {
		case tokens[index].IsSymbol(".") || tokens[index].IsSymbol(":"):
			if tokens[index+1].Kind != tokenIdentifier {
				return strings.Join(parts, ""), index, true
			}
			parts = append(parts, tokens[index].Text, tokens[index+1].Text)
			index += 2
		case tokens[index].IsSymbol("[") && index+2 < len(tokens) && tokens[index+1].Kind == tokenString && tokens[index+2].IsSymbol("]"):
			key := strings.Trim(tokens[index+1].Text, "\"'")
			parts = append(parts, ".", key)
			index += 3
		default:
			return strings.Join(parts, ""), index, true
		}
	}
	return strings.Join(parts, ""), index, true
}

func matchPair(tokens []Token, start int, open string, close string) int {
	depth := 1
	for index := start + 1; index < len(tokens); index++ {
		if tokens[index].IsSymbol(open) {
			depth++
			continue
		}
		if tokens[index].IsSymbol(close) {
			depth--
			if depth == 0 {
				return index
			}
		}
	}
	return len(tokens) - 1
}

func splitTopLevel(tokens []Token, delimiter string) [][]Token {
	if len(tokens) == 0 {
		return nil
	}
	groups := [][]Token{}
	start := 0
	parenDepth := 0
	braceDepth := 0
	bracketDepth := 0
	for index, token := range tokens {
		switch token.Text {
		case "(":
			parenDepth++
		case ")":
			parenDepth--
		case "{":
			braceDepth++
		case "}":
			braceDepth--
		case "[":
			bracketDepth++
		case "]":
			bracketDepth--
		}
		if token.Text == delimiter && parenDepth == 0 && braceDepth == 0 && bracketDepth == 0 {
			groups = append(groups, tokens[start:index])
			start = index + 1
		}
	}
	if start <= len(tokens) {
		groups = append(groups, tokens[start:])
	}
	return groups
}

func expressionString(tokens []Token) string {
	if len(tokens) == 0 {
		return ""
	}
	var builder strings.Builder
	for index, token := range tokens {
		if index > 0 && needsSpace(tokens[index-1], token) {
			builder.WriteRune(' ')
		}
		builder.WriteString(token.Text)
	}
	return builder.String()
}

func needsSpace(left Token, right Token) bool {
	if left.Kind == tokenIdentifier || left.Kind == tokenKeyword || left.Kind == tokenNumber || left.Kind == tokenString {
		return right.Kind == tokenIdentifier || right.Kind == tokenKeyword || right.Kind == tokenNumber || right.Kind == tokenString
	}
	if left.Text == "]" && (right.Kind == tokenIdentifier || right.Kind == tokenNumber) {
		return true
	}
	return false
}

func parseTableLiteral(tokens []Token) constValue {
	value := constValue{
		Kind: "table",
		Text: expressionString(tokens),
		Rows: [][]string{},
	}
	if len(tokens) < 2 {
		return value
	}
	items := splitTopLevel(tokens[1:len(tokens)-1], ",")
	for _, item := range items {
		trimmed := trimTokenSlice(item)
		if len(trimmed) == 0 {
			continue
		}
		if trimmed[0].IsSymbol("{") && trimmed[len(trimmed)-1].IsSymbol("}") {
			rowTokens := splitTopLevel(trimmed[1:len(trimmed)-1], ",")
			row := []string{}
			for _, group := range rowTokens {
				row = append(row, strings.TrimSpace(expressionString(group)))
			}
			value.Rows = append(value.Rows, row)
			continue
		}
		value.Rows = append(value.Rows, []string{strings.TrimSpace(expressionString(trimmed))})
	}
	return value
}

func trimTokenSlice(tokens []Token) []Token {
	start := 0
	for start < len(tokens) && tokens[start].Text == "," {
		start++
	}
	end := len(tokens)
	for end > start && tokens[end-1].Text == "," {
		end--
	}
	return tokens[start:end]
}

func resolveName(name string, aliases map[string]string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return ""
	}
	if aliases == nil {
		return trimmed
	}
	separatorIndex := strings.IndexAny(trimmed, ".:")
	if separatorIndex < 0 {
		if resolved, ok := aliases[trimmed]; ok {
			return resolved
		}
		return trimmed
	}
	root := trimmed[:separatorIndex]
	if resolvedRoot, ok := aliases[root]; ok {
		return resolvedRoot + trimmed[separatorIndex:]
	}
	return trimmed
}

func resolveExpression(value string, context scanContext) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return ""
	}
	if constant, ok := context.consts[trimmed]; ok && constant.Kind != "table" {
		return constant.Text
	}
	if target, ok := context.functionSymbols[trimmed]; ok {
		return target
	}
	if strings.HasPrefix(trimmed, "\"") || strings.HasPrefix(trimmed, "'") {
		return strings.Trim(trimmed, "\"'")
	}
	return resolveName(trimmed, context.aliasMap)
}

func inferValueType(tokens []Token, start int) string {
	if start >= len(tokens) {
		return "unknown"
	}
	token := tokens[start]
	switch {
	case token.IsKeyword("true") || token.IsKeyword("false"):
		return "boolean"
	case token.IsKeyword("nil"):
		return "nil"
	case token.Kind == tokenString:
		return "string"
	case token.Kind == tokenNumber:
		return "number"
	case token.IsSymbol("{"):
		return "table"
	case token.IsKeyword("function"):
		return "function"
	case token.Kind == tokenIdentifier:
		if start+1 < len(tokens) && tokens[start+1].IsSymbol("(") {
			return "call-result"
		}
		if start+2 < len(tokens) && tokens[start+1].IsKeyword("or") && tokens[start+2].IsSymbol("{") {
			return "table"
		}
		return "reference"
	default:
		return "expression"
	}
}

func capturesModule(name string, valueType string, context scanContext) bool {
	if valueType != "table" {
		return false
	}
	root := graph.RootSegment(name)
	if root == "" {
		return false
	}
	if !isPublicRoot(root, context) {
		return false
	}
	if context.localNames[root] && context.aliasMap[root] == "" {
		return false
	}
	return strings.Contains(name, ".") || root == context.addon || isExportedIdentifier(root)
}

func capturesState(name string, context scanContext) bool {
	root := graph.RootSegment(name)
	if root == "" {
		return false
	}
	if !isPublicRoot(root, context) {
		return false
	}
	if context.localNames[root] && context.aliasMap[root] == "" && context.functionSymbols[root] == "" {
		return false
	}
	if context.savedRoots[root] {
		return true
	}
	if strings.Contains(name, ".") {
		return true
	}
	return root == context.addon
}

func assignmentExportsFunction(tokens []Token, start int, context scanContext) bool {
	if start >= len(tokens) {
		return false
	}
	if tokens[start].IsKeyword("function") {
		return true
	}
	if tokens[start].Kind == tokenIdentifier {
		if _, ok := context.functionSymbols[tokens[start].Text]; ok {
			return true
		}
	}
	return false
}

func matchesIndexedVar(expression string, variable string, index int) bool {
	normalized := strings.ReplaceAll(strings.TrimSpace(expression), " ", "")
	return normalized == fmt.Sprintf("%s[%d]", variable, index)
}

func qualifyLocalFunction(addonName string, shortName string) string {
	return addonName + ".local." + shortName
}

func stateNames(state []graph.StateVariable) []string {
	names := make([]string, 0, len(state))
	for _, variable := range state {
		names = append(names, variable.Name)
	}
	return graph.UniqueStrings(names)
}

func functionKind(name string) string {
	segment := strings.ToLower(graph.LastSegment(name))
	switch {
	case strings.HasPrefix(segment, "on"):
		return "handler"
	case strings.Contains(segment, "initialize") || strings.Contains(segment, "init"):
		return "lifecycle"
	case strings.Contains(segment, "shutdown") || strings.Contains(segment, "cleanup"):
		return "lifecycle"
	case strings.Contains(segment, "update") || strings.Contains(segment, "refresh"):
		return "update"
	case strings.Contains(segment, "load") || strings.Contains(segment, "save"):
		return "settings"
	default:
		return "function"
	}
}

func dedupeCalls(calls []graph.Call) []graph.Call {
	seen := map[string]bool{}
	result := make([]graph.Call, 0, len(calls))
	for _, call := range calls {
		key := call.CalleeRaw + "|" + call.CalleeResolved + "|" + fmt.Sprintf("%d", call.Line) + "|" + strings.Join(call.Arguments, ",")
		if seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, call)
	}
	return result
}

func copyStringMap(input map[string]string) map[string]string {
	output := map[string]string{}
	for key, value := range input {
		output[key] = value
	}
	return output
}

func copyBoolMap(input map[string]bool) map[string]bool {
	output := map[string]bool{}
	for key, value := range input {
		output[key] = value
	}
	return output
}

func copyConstMap(input map[string]constValue) map[string]constValue {
	output := map[string]constValue{}
	for key, value := range input {
		rows := make([][]string, 0, len(value.Rows))
		for _, row := range value.Rows {
			rows = append(rows, append([]string{}, row...))
		}
		output[key] = constValue{Kind: value.Kind, Text: value.Text, Rows: rows}
	}
	return output
}

func isExportedIdentifier(value string) bool {
	if value == "" {
		return false
	}
	first := value[0]
	return first >= 'A' && first <= 'Z'
}

func isPublicRoot(root string, context scanContext) bool {
	if root == context.addon || context.savedRoots[root] {
		return true
	}
	return isExportedIdentifier(root)
}

func declaredFunctionName(span functionSpan) string {
	if strings.TrimSpace(span.RawName) != "" {
		return span.RawName
	}
	return span.ShortName
}

func scopeKind(local bool) string {
	if local {
		return "local"
	}
	return "global"
}

func declarationKind(span functionSpan) string {
	if strings.TrimSpace(span.DeclarationKind) != "" {
		return span.DeclarationKind
	}
	if span.Local {
		return "local_function_decl"
	}
	if strings.TrimSpace(span.RawName) != "" {
		return "function_decl"
	}
	return "assignment_function"
}

func splitReceiver(rawName string) (string, string) {
	trimmed := strings.TrimSpace(rawName)
	if trimmed == "" {
		return "", ""
	}
	lastDot := strings.LastIndex(trimmed, ".")
	lastColon := strings.LastIndex(trimmed, ":")
	if lastDot < 0 && lastColon < 0 {
		return "", ""
	}
	if lastColon > lastDot {
		return trimmed[:lastColon], ":"
	}
	return trimmed[:lastDot], "."
}

func maxInt(left int, right int) int {
	if left > right {
		return left
	}
	return right
}

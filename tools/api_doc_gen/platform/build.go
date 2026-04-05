package platform

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"roraddons/tools/api_doc_gen/confidence"
	"roraddons/tools/api_doc_gen/graph"
)

var namespaceTokenPattern = regexp.MustCompile(`\b(?:SystemData|GameData)(?:\.[A-Za-z_][A-Za-z0-9_]*)+\b`)

var standardLuaCalls = map[string]bool{
	"assert":         true,
	"collectgarbage": true,
	"error":          true,
	"getmetatable":   true,
	"ipairs":         true,
	"next":           true,
	"pairs":          true,
	"pcall":          true,
	"print":          true,
	"rawequal":       true,
	"rawget":         true,
	"rawset":         true,
	"select":         true,
	"setmetatable":   true,
	"tonumber":       true,
	"tostring":       true,
	"type":           true,
	"unpack":         true,
	"xpcall":         true,
	"coroutine":      true,
	"debug":          true,
	"io":             true,
	"math":           true,
	"os":             true,
	"string":         true,
	"table":          true,
}

var windowFunctionPrefixes = []string{
	"Window",
	"Label",
	"Button",
	"DynamicImage",
	"TextEditBox",
	"ComboBox",
	"ScrollWindow",
	"StatusBar",
}

var windowFunctionRoots = map[string]bool{
	"LayoutEditor": true,
}

var sharedTableRoots = map[string]bool{
	"EA_ChatWindow": true,
	"LibSlash":      true,
	"wstring":       true,
}

var explicitPlatformTables = map[string]bool{
	"ActiveWindow":      true,
	"AlertText":         true,
	"AlertTextWindow":   true,
	"AuctionWindow":     true,
	"BankWindow":        true,
	"CharacterWindow":   true,
	"CraftingSystem":    true,
	"CultivationWindow": true,
	"Cursor":            true,
	"DefaultColor":      true,
	"DialogManager":     true,
	"GroupWindow":       true,
	"HelpTips":          true,
	"Icons":             true,
	"InterfaceCore":     true,
	"MailWindow":        true,
	"MailWindowTabSend": true,
	"MoneyFrame":        true,
	"PartyUtils":        true,
	"PetWindow":         true,
	"PlayerMenuWindow":  true,
	"Window":            true,
}

var signalGlobalFunctions = map[string]bool{
	"BroadcastEvent":                 true,
	"CreateWindow":                   true,
	"CreateWindowFromTemplate":       true,
	"DestroyWindow":                  true,
	"DoesWindowExist":                true,
	"GetBuffs":                       true,
	"GetIconData":                    true,
	"GetInventoryItemData":           true,
	"RegisterEventHandler":           true,
	"SendUseItem":                    true,
	"UnregisterEventHandler":         true,
	"WindowRegisterCoreEventHandler": true,
	"WindowRegisterEventHandler":     true,
	"towstring":                      true,
}

type addonScope struct {
	roots        map[string]bool
	lastSegments map[string]bool
}

type callObservation struct {
	Addon     string
	Caller    string
	Arguments []string
	Raw       string
	Source    string
	Line      int
}

type functionAccumulator struct {
	Name         string
	Addons       map[string]bool
	Aliases      map[string]bool
	Observations []callObservation
}

type eventAccumulator struct {
	Name        string
	Addons      map[string]bool
	Registrars  map[string]bool
	Handlers    map[string]bool
	TriggeredBy map[string]bool
	Examples    []UsageExample
}

type xmlHandlerAccumulator struct {
	Name         string
	Addons       map[string]bool
	ElementTypes map[string]bool
	Examples     []UsageExample
}

type fieldAccumulator struct {
	Name     string
	Addons   map[string]bool
	Contexts map[string]bool
	Files    map[string]bool
}

type tableAccumulator struct {
	Name      string
	Addons    map[string]bool
	Functions map[string]bool
	Members   map[string]bool
	Examples  []UsageExample
	Notes     []string
}

type constantAccumulator struct {
	Name   string
	Addons map[string]bool
	UsedBy map[string]bool
	Notes  []string
	Files  map[string]bool
}

type elementAccumulator struct {
	Name              string
	Addons            map[string]bool
	Attributes        map[string]int
	Handlers          map[string]int
	HandlerFunctions  map[string]int            // Lua function names bound via XML handlers
	HandlerBindings   map[string]map[string]int // event → {lua_function → count}
	Inherits          map[string]int
	ChildTypes        map[string]int // structural (unnamed) child element types
	ChildElementTypes map[string]int // named child frame element types
	ParentTypes       map[string]int // element types that contain this one
	Snippets          []string       // CompositionSnippet candidates from real frames
	Examples          []UsageExample
}

type lifecycleAccumulator struct {
	Phase    string
	Addons   map[string]bool
	Evidence []string
	Details  []string
}

type scoringContext struct {
	input                 contractSemanticInput
	addonNames            []string
	definitionCounts      map[string]int
	localDefinitionCounts map[string]int
}

type candidateCollector struct {
	items []CandidateSummary
}

func (collector *candidateCollector) record(name string, category string, assessment confidence.Assessment) {
	status := CandidateStatusCandidate
	if confidence.IsRejectedAddonLocal(assessment) {
		status = CandidateStatusRejected
	} else if confidence.ShouldPromote(assessment) {
		status = CandidateStatusCanonical
	}
	collector.items = append(collector.items, CandidateSummary{
		Name:       name,
		Category:   category,
		Status:     status,
		RawScore:   assessment.RawScore,
		Score:      assessment.FinalScore,
		Confidence: toPlatformConfidence(assessment.Level),
		Signals:    assessment.Signals,
		Rationale:  assessment.Rationale,
		Evidence:   assessment.Evidence,
	})
}

func newScoringContext(input contractSemanticInput) scoringContext {
	addonSet := map[string]bool{}
	definitionCounts := map[string]int{}
	localDefinitionCounts := map[string]int{}
	for _, doc := range input.Functions {
		addonSet[doc.Addon] = true
		definitionCounts[doc.Name]++
		if doc.Local {
			localDefinitionCounts[doc.Name]++
		}
		for _, alias := range doc.Aliases {
			definitionCounts[alias]++
			if doc.Local {
				localDefinitionCounts[alias]++
			}
		}
	}
	for _, frame := range input.Frames {
		addonSet[frame.Addon] = true
	}
	for _, handler := range input.Handlers {
		addonSet[handler.Addon] = true
	}

	return scoringContext{
		input:                 input,
		addonNames:            mapKeys(addonSet),
		definitionCounts:      definitionCounts,
		localDefinitionCounts: localDefinitionCounts,
	}
}

func Build(contracts ContractModel) Corpus {
	return BuildWithOptions(contracts, BuildOptions{})
}

// BuildOptions configures optional behaviour for [Build].
type BuildOptions struct {
	// SourceRoot, when non-empty, enables the source-first pipeline path.
	// The directory must contain addon sub-directories with manifest files
	// (*.mod or *.toc). When set, real XML and Lua source files are parsed
	// directly from the addons directory.
	//
	// In contract-only mode SourceRoot is required; the pipeline panics
	// if it is empty.
	SourceRoot string
}

// BuildWithOptions builds the platform corpus with explicit options.
// Use this instead of [Build] to enable the source-first pipeline path.
func BuildWithOptions(contracts ContractModel, opts BuildOptions) Corpus {
	if err := validateContractModelConsistency(contracts); err != nil {
		panic("contract-only pipeline guard: " + err.Error())
	}

	input := semanticInputFromContracts(contracts)
	corpus := Corpus{
		SourceRoot:  contracts.Root,
		Contracts:   contracts,
		GeneratedAt: time.Now().UTC(),
	}
	ctx := newScoringContext(input)
	collector := &candidateCollector{}

	scopes := buildAddonScopes(input)
	frameTypes := map[string]string{}
	globalFunctions := map[string]*functionAccumulator{}
	windowFunctions := map[string]*functionAccumulator{}
	gameEvents := map[string]*eventAccumulator{}
	windowEvents := map[string]*eventAccumulator{}
	xmlHandlers := map[string]*xmlHandlerAccumulator{}
	systemFields := map[string]*fieldAccumulator{}
	gameFields := map[string]*fieldAccumulator{}
	tables := map[string]*tableAccumulator{}
	constants := map[string]*constantAccumulator{}
	elements := map[string]*elementAccumulator{}
	lifecycle := map[string]*lifecycleAccumulator{}

	for _, frame := range input.Frames {
		frameName := canonicalSymbolName(frame.Name)
		frameType := canonicalElementTypeName(frame.Type)
		frameTypes[frame.Addon+"|"+frameName] = frameType
		acc := ensureElementAccumulator(elements, frameType)
		acc.Addons[frame.Addon] = true
		for key := range frame.Attributes {
			if canonicalKey := canonicalAttributeName(key); canonicalKey != "" {
				acc.Attributes[canonicalKey]++
			}
		}
		for _, handler := range frame.Handlers {
			eventName := canonicalEventName(handler.Event)
			acc.Handlers[eventName]++
			if functionName := canonicalSymbolName(handler.Function); functionName != "" {
				acc.HandlerFunctions[functionName]++
				// Track the (event, lua_function) binding pair so element pages can
				// surface which Lua functions are commonly bound per XML event.
				if acc.HandlerBindings[eventName] == nil {
					acc.HandlerBindings[eventName] = map[string]int{}
				}
				acc.HandlerBindings[eventName][functionName]++
			}
		}
		if inherits := canonicalSymbolName(frame.Inherits); inherits != "" {
			acc.Inherits[inherits]++
		}
		// Track the element type of the parent that contains this frame.
		if parentType := canonicalElementTypeName(frame.ParentType); parentType != "" {
			acc.ParentTypes[parentType]++
		}
		// Track named child element types.
		for _, childElemType := range frame.ChildElementTypes {
			if childType := canonicalElementTypeName(childElemType); childType != "" {
				acc.ChildElementTypes[childType]++
				// Record the reverse: this child type has the current frame type as a parent.
				childAcc := ensureElementAccumulator(elements, childType)
				childAcc.ParentTypes[frameType]++
			}
		}
		for _, childType := range frame.StructuralChildTypes {
			if structuralType := canonicalElementTypeName(childType); structuralType != "" {
				acc.ChildTypes[structuralType]++
			}
		}
		// Feed structural children through their own elementAccumulator entries so that
		// they get first-class pages alongside named element types (ListBox, Button, etc.).
		for _, childType := range frame.StructuralChildTypes {
			structuralType := canonicalElementTypeName(childType)
			if structuralType == "" {
				continue
			}
			childAcc := ensureElementAccumulator(elements, structuralType)
			childAcc.Addons[frame.Addon] = true
			// Carry over attribute keys captured from the XML source.
			for _, attrKey := range frame.StructuralChildAttrKeys[childType] {
				if canonicalKey := canonicalAttributeName(attrKey); canonicalKey != "" {
					childAcc.Attributes[canonicalKey]++
				}
			}
			childAcc.Examples = appendUniqueExample(childAcc.Examples, UsageExample{
				Addon:   frame.Addon,
				Caller:  frameName,
				Snippet: structuralType + " in " + frameType + " " + frameName,
				Source:  frame.Source,
			})
		}
		acc.Examples = appendUniqueExample(acc.Examples, UsageExample{
			Addon:   frame.Addon,
			Caller:  frameName,
			Snippet: frameType + " " + frameName,
			Source:  frame.Source,
		})
		// Collect etree-derived composition snippets for later selection.
		if frame.CompositionSnippet != "" {
			acc.Snippets = append(acc.Snippets, frame.CompositionSnippet)
		}

		if shouldKeepConstant(frame.Inherits, frame.Addon, nil) {
			constantAcc := ensureConstantAccumulator(constants, frame.Inherits)
			constantAcc.Addons[frame.Addon] = true
			constantAcc.UsedBy[frameName] = true
			constantAcc.Files[frame.Source] = true
		}
		if shouldKeepConstant(frame.Parent, frame.Addon, nil) {
			constantAcc := ensureConstantAccumulator(constants, frame.Parent)
			constantAcc.Addons[frame.Addon] = true
			constantAcc.UsedBy[frameName] = true
			constantAcc.Files[frame.Source] = true
		}
	}

	for _, doc := range input.Functions {
		for _, item := range doc.EventRegistrations {
			eventName := canonicalEventName(item.Event)
			acc := ensureEventAccumulator(selectEventAccumulator(eventName, gameEvents, windowEvents), eventName)
			acc.Addons[doc.Addon] = true
			if handlerName := canonicalSymbolName(item.Handler); handlerName != "" {
				acc.Handlers[handlerName] = true
			}
			if scopeName := canonicalSymbolName(item.Scope); scopeName != "" {
				acc.Registrars[scopeName] = true
			}
			acc.Examples = appendUniqueExample(acc.Examples, UsageExample{
				Addon:   doc.Addon,
				Caller:  doc.Name,
				Snippet: eventName + " -> " + canonicalSymbolName(item.Handler),
				Source:  doc.Source,
			})
			extractNamespaceTokens(eventName, doc.Addon, canonicalSymbolName(item.Handler), doc.Source, "event_registration", systemFields, gameFields)
		}

		for _, call := range dedupeCallAliases(doc.Calls) {
			extractNamespaceTokens(strings.Join(call.Arguments, ", "), doc.Addon, doc.Name, doc.Source, "lua_call", systemFields, gameFields)
			if call.Name == "" || isStandardLuaCall(call.Name) || isLocalCall(doc.Addon, call.Name, scopes) {
				continue
			}
			if !shouldKeepFunction(call.Name, doc.Addon, scopes, call, classifyFunctionCategory(call.Name)) {
				continue
			}
			accMap := globalFunctions
			if classifyFunctionCategory(call.Name) == "Window Function" {
				accMap = windowFunctions
			}
			acc := ensureFunctionAccumulator(accMap, call.Name)
			acc.Addons[doc.Addon] = true
			acc.Observations = append(acc.Observations, callObservation{
				Addon:     doc.Addon,
				Caller:    doc.Name,
				Arguments: call.Arguments,
				Raw:       call.Raw,
				Source:    doc.Source,
				Line:      call.Line,
			})

			if root := graph.RootSegment(call.Name); root != "" && root != call.Name && root != "SystemData" && root != "GameData" {
				tableAcc := ensureTableAccumulator(tables, root)
				tableAcc.Addons[doc.Addon] = true
				tableAcc.Functions[call.Name] = true
				tableAcc.Examples = appendUniqueExample(tableAcc.Examples, UsageExample{
					Addon:   doc.Addon,
					Caller:  doc.Name,
					Snippet: formatSnippet(call.Name, call.Arguments),
					Source:  doc.Source,
				})
			}
		}
	}

	for _, handler := range input.Handlers {
		eventName := canonicalEventName(handler.Event)
		frameName := canonicalSymbolName(handler.Frame)
		functionName := canonicalSymbolName(handler.Function)
		xmlAcc := ensureXMLHandlerAccumulator(xmlHandlers, eventName)
		xmlAcc.Addons[handler.Addon] = true
		if elementType := frameTypes[handler.Addon+"|"+frameName]; elementType != "" {
			xmlAcc.ElementTypes[elementType] = true
		}
		xmlAcc.Examples = appendUniqueExample(xmlAcc.Examples, UsageExample{
			Addon:   handler.Addon,
			Caller:  frameName,
			Snippet: frameName + "." + eventName + " -> " + functionName,
			Source:  handler.Source,
		})

		eventAcc := ensureEventAccumulator(selectEventAccumulator(eventName, gameEvents, windowEvents), eventName)
		eventAcc.Addons[handler.Addon] = true
		if functionName != "" {
			eventAcc.Handlers[functionName] = true
		}
		eventAcc.Examples = appendUniqueExample(eventAcc.Examples, UsageExample{
			Addon:   handler.Addon,
			Caller:  frameName,
			Snippet: frameName + "." + eventName + " -> " + functionName,
			Source:  handler.Source,
		})
	}

	for _, event := range input.Events {
		eventName := canonicalEventName(event.Name)
		acc := ensureEventAccumulator(selectEventAccumulator(eventName, gameEvents, windowEvents), eventName)
		for _, registration := range event.LuaRegistrations {
			acc.Addons[registration.Addon] = true
			if registrar := canonicalSymbolName(registration.Registrar); registrar != "" {
				acc.Registrars[registrar] = true
			}
			if handlerName := canonicalSymbolName(registration.Handler); handlerName != "" {
				acc.Handlers[handlerName] = true
			}
			acc.Examples = appendUniqueExample(acc.Examples, UsageExample{
				Addon:   registration.Addon,
				Caller:  canonicalSymbolName(registration.Handler),
				Snippet: canonicalSymbolName(registration.Registrar) + "(" + eventName + ", " + canonicalSymbolName(registration.Handler) + ")",
				Source:  eventName,
			})
		}
		for _, handler := range event.XMLHandlers {
			acc.Addons[handler.Addon] = true
			if functionName := canonicalSymbolName(handler.Function); functionName != "" {
				acc.Handlers[functionName] = true
			}
			acc.Examples = appendUniqueExample(acc.Examples, UsageExample{
				Addon:   handler.Addon,
				Caller:  canonicalSymbolName(handler.Frame),
				Snippet: canonicalSymbolName(handler.Frame) + "." + eventName + " -> " + canonicalSymbolName(handler.Function),
				Source:  eventName,
			})
		}
		for _, trigger := range event.TriggeredBy {
			if trigger != "" {
				acc.TriggeredBy[trigger] = true
			}
		}
		extractNamespaceTokens(eventName, strings.Join(mapKeys(acc.Addons), ", "), eventName, "", "event_page", systemFields, gameFields)
	}

	for _, binding := range input.Bindings {
		eventName := canonicalEventName(binding.Event)
		xmlAcc := ensureXMLHandlerAccumulator(xmlHandlers, eventName)
		xmlAcc.Addons[binding.Addon] = true
		xmlAcc.Examples = appendUniqueExample(xmlAcc.Examples, UsageExample{
			Addon:   binding.Addon,
			Caller:  canonicalSymbolName(binding.Frame),
			Snippet: canonicalSymbolName(binding.Frame) + "." + eventName + " -> " + canonicalSymbolName(binding.LuaFunction),
			Source:  binding.XMLFunction,
		})
	}

	// Contract artifacts do not currently carry saved-variable, flow, or curated
	// example docs. Those pathways are intentionally disabled in contract-only mode.

	corpus.GlobalFunctions = finalizeFunctionSymbols(globalFunctions, "Global Function", ctx, collector)
	corpus.WindowFunctions = finalizeFunctionSymbols(windowFunctions, "Window Function", ctx, collector)
	corpus.GlobalTables = finalizeTableSymbols(tables, ctx, collector)
	corpus.Constants = finalizeConstantSymbols(constants, ctx, collector)
	corpus.ElementTypes = finalizeElementSymbols(elements, ctx, collector)
	corpus.XMLHandlers = finalizeXMLHandlerSymbols(xmlHandlers, ctx, collector)

	// Run the new phased XML↔Lua semantic pipeline (Phases 1-4).
	// This enriches ElementTypes with structured attribute profiles,
	// structural child profiles, Lua API call aggregations, handler
	// argument patterns, and .mod lifecycle facts.
	pipelineResult := runPhasedPipeline(corpus.ElementTypes, opts.SourceRoot)
	corpus.ElementTypes = pipelineResult.ElementTypes
	corpus.AddonLifecycleSemantics = pipelineResult.AddonLifecycleSemantics
	corpus.FunctionLifecycleRoles = pipelineResult.FunctionLifecycleRoles
	corpus.LifecycleDiagnostics = pipelineResult.LifecycleDiagnostics
	corpus.GameEvents = finalizeEventSymbols(gameEvents, "Game Event", ctx, collector)
	corpus.WindowEvents = finalizeEventSymbols(windowEvents, "Window Event", ctx, collector)
	corpus.SystemDataFields = finalizeFieldSymbols(systemFields, "SystemData", ctx, collector)
	corpus.GameDataFields = finalizeFieldSymbols(gameFields, "GameData", ctx, collector)
	corpus.SlashCommandPatterns = buildSlashPatterns(globalFunctions)
	corpus.WindowPatterns = buildWindowPatterns(globalFunctions, windowFunctions, input.Bindings)
	corpus.EventPatterns = buildEventPatterns(globalFunctions, corpus.GameEvents, corpus.WindowEvents)
	corpus.Lifecycle = finalizeLifecycle(lifecycle)
	corpus.Conventions = buildConventions(corpus, input)
	corpus.InferenceRules = defaultInferenceRules()
	corpus.Coverage = buildCoverage(corpus, collector.items, input)
	enrichSemanticArtifacts(&corpus, input)
	return corpus
}

func buildAddonScopes(input contractSemanticInput) map[string]addonScope {
	scopes := map[string]addonScope{}
	ensure := func(addon string) addonScope {
		scope, ok := scopes[addon]
		if !ok {
			scope = addonScope{roots: map[string]bool{}, lastSegments: map[string]bool{}}
		}
		return scope
	}
	for _, doc := range input.Functions {
		scope := ensure(doc.Addon)
		scope.roots[graph.RootSegment(doc.Addon)] = true
		if root := graph.RootSegment(doc.Module); root != "" {
			scope.roots[root] = true
		}
		if root := graph.RootSegment(doc.Name); root != "" {
			scope.roots[root] = true
		}
		scope.lastSegments[graph.LastSegment(doc.Name)] = true
		for _, alias := range doc.Aliases {
			scope.lastSegments[graph.LastSegment(alias)] = true
		}
		scopes[doc.Addon] = scope
	}
	return scopes
}

func dedupeCallAliases(calls []CallDoc) []CallDoc {
	grouped := map[string][]CallDoc{}
	for _, call := range calls {
		key := fmt.Sprintf("%d|%s|%s", call.Line, call.Raw, graph.LastSegment(call.Name))
		grouped[key] = append(grouped[key], call)
	}
	result := []CallDoc{}
	for _, group := range grouped {
		preferred := []CallDoc{}
		for _, call := range group {
			if strings.Contains(call.Name, ".") {
				preferred = append(preferred, call)
			}
		}
		if len(preferred) == 0 {
			preferred = group
		}
		seen := map[string]bool{}
		for _, call := range preferred {
			if seen[call.Name] {
				continue
			}
			seen[call.Name] = true
			result = append(result, call)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Line == result[j].Line {
			return result[i].Name < result[j].Name
		}
		return result[i].Line < result[j].Line
	})
	return result
}

func isStandardLuaCall(name string) bool {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return true
	}
	root := graph.RootSegment(trimmed)
	if standardLuaCalls[trimmed] {
		return true
	}
	return standardLuaCalls[root] && root != "wstring"
}

func isLocalCall(addon string, name string, scopes map[string]addonScope) bool {
	scope, ok := scopes[addon]
	if !ok {
		return false
	}
	root := graph.RootSegment(name)
	if root != "" && root != name {
		if root == "SystemData" || root == "GameData" || sharedTableRoots[root] || windowFunctionRoots[root] {
			return false
		}
		for _, prefix := range windowFunctionPrefixes {
			if root == prefix {
				return false
			}
		}
		return scope.roots[root]
	}
	return scope.lastSegments[name]
}

func shouldKeepFunction(name string, addon string, scopes map[string]addonScope, call CallDoc, category string) bool {
	if strings.Contains(name, ":") {
		return false
	}
	if category == "Window Function" {
		return true
	}
	if signalGlobalFunctions[name] {
		return true
	}
	root := graph.RootSegment(name)
	if sharedTableRoots[root] {
		return true
	}
	if strings.Contains(name, ".") && !isLocalCall(addon, name, scopes) && isEngineishGlobalRoot(root) {
		return true
	}
	return false
}

func classifyFunctionCategory(name string) string {
	if isWindowFunctionName(name) {
		return "Window Function"
	}
	return "Global Function"
}

func isWindowFunctionName(name string) bool {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" || strings.Contains(trimmed, ":") {
		return false
	}
	root := graph.RootSegment(trimmed)
	if strings.Contains(trimmed, ".") {
		return windowFunctionRoots[root]
	}
	for _, prefix := range windowFunctionPrefixes {
		if strings.HasPrefix(trimmed, prefix) && len(trimmed) > len(prefix) {
			next := trimmed[len(prefix)]
			if next >= 'A' && next <= 'Z' {
				return true
			}
		}
	}
	return false
}

func isEngineishGlobalRoot(root string) bool {
	trimmed := strings.TrimSpace(root)
	if trimmed == "" {
		return false
	}
	if sharedTableRoots[trimmed] {
		return true
	}
	if strings.HasPrefix(trimmed, "Lib") || strings.HasPrefix(trimmed, "LIB") {
		return false
	}
	if trimmed == "GameData" || trimmed == "SystemData" || trimmed == "InterfaceCore" || trimmed == "DialogManager" || trimmed == "Icons" || trimmed == "PartyUtils" || trimmed == "Cursor" || trimmed == "HelpTips" || trimmed == "DefaultColor" {
		return true
	}
	if strings.HasPrefix(trimmed, "EA_") || strings.HasSuffix(trimmed, "Window") {
		return true
	}
	if strings.HasPrefix(trimmed, "Auction") || strings.HasPrefix(trimmed, "Bank") || strings.HasPrefix(trimmed, "Crafting") || strings.HasPrefix(trimmed, "Mail") || strings.HasPrefix(trimmed, "Money") || strings.HasPrefix(trimmed, "Alert") || strings.HasPrefix(trimmed, "TextLog") {
		return true
	}
	return false
}

func selectEventAccumulator(name string, game map[string]*eventAccumulator, window map[string]*eventAccumulator) map[string]*eventAccumulator {
	if isWindowEventName(name) {
		return window
	}
	return game
}

func isWindowEventName(name string) bool {
	trimmed := strings.TrimSpace(name)
	return strings.HasPrefix(trimmed, "On") && !strings.Contains(trimmed, "SystemData.")
}

func extractNamespaceTokens(text string, addon string, context string, source string, kind string, system map[string]*fieldAccumulator, game map[string]*fieldAccumulator) {
	for _, token := range namespaceTokenPattern.FindAllString(text, -1) {
		target := system
		if strings.HasPrefix(token, "GameData.") {
			target = game
		}
		acc := ensureFieldAccumulator(target, token)
		acc.Addons[addon] = true
		if context != "" {
			acc.Contexts[context] = true
		}
		if source != "" {
			acc.Files[source] = true
		}
		if kind != "" {
			acc.Contexts[kind] = true
		}
	}
}

func shouldKeepConstant(name string, addon string, counts map[string]map[string]bool) bool {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" || strings.EqualFold(trimmed, "none") {
		return false
	}
	if trimmed == "Root" || strings.HasPrefix(trimmed, "EA_") {
		return true
	}
	if counts == nil {
		return false
	}
	return len(counts[trimmed]) >= 2
}

func finalizeFunctionSymbols(values map[string]*functionAccumulator, category string, ctx scoringContext, collector *candidateCollector) []FunctionSymbol {
	result := make([]FunctionSymbol, 0, len(values))
	for _, acc := range values {
		signature, parameters, arity := inferSignature(acc.Name, acc.Observations)
		assessment := confidence.Score(buildFunctionEvidence(acc, category, ctx, arity))
		collector.record(acc.Name, category, assessment)
		if !confidence.ShouldPromote(assessment) {
			continue
		}
		level := toPlatformConfidence(assessment.Level)
		result = append(result, FunctionSymbol{
			Name:          acc.Name,
			Category:      category,
			Confidence:    level,
			RawScore:      assessment.RawScore,
			Score:         assessment.FinalScore,
			Signals:       assessment.Signals,
			Rationale:     assessment.Rationale,
			Evidence:      assessment.Evidence,
			Description:   describeFunction(acc.Name, category, len(acc.Addons)),
			Signature:     signature,
			Parameters:    parameters,
			Returns:       inferReturns(acc.Name),
			SideEffects:   inferSideEffects(acc.Name),
			SeenIn:        mapKeys(acc.Addons),
			Examples:      firstFunctionExamples(acc.Name, acc.Observations, 6),
			Aliases:       mapKeys(acc.Aliases),
			ObservedArity: arity,
			Notes:         functionNotes(category, level, len(acc.Addons)),
		})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func finalizeEventSymbols(values map[string]*eventAccumulator, category string, ctx scoringContext, collector *candidateCollector) []EventSymbol {
	result := make([]EventSymbol, 0, len(values))
	for _, acc := range values {
		assessment := confidence.Score(buildEventEvidence(acc, category, ctx))
		collector.record(acc.Name, category, assessment)
		if !confidence.ShouldPromote(assessment) {
			continue
		}
		level := toPlatformConfidence(assessment.Level)
		result = append(result, EventSymbol{
			Name:           acc.Name,
			Category:       category,
			Confidence:     level,
			RawScore:       assessment.RawScore,
			Score:          assessment.FinalScore,
			Signals:        assessment.Signals,
			Rationale:      assessment.Rationale,
			Evidence:       assessment.Evidence,
			Description:    describeEvent(acc.Name, category, len(acc.Addons), len(acc.Registrars), len(acc.Handlers)),
			HandlerPattern: inferEventHandlerPattern(acc.Name),
			Payload:        inferEventPayload(acc.Name),
			SeenIn:         mapKeys(acc.Addons),
			Examples:       firstUsageExamples(acc.Examples, 6),
			Registrars:     graph.UniqueStrings(append(mapKeys(acc.Registrars), mapKeys(acc.Handlers)...)),
			Notes:          eventNotes(acc),
		})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func finalizeXMLHandlerSymbols(values map[string]*xmlHandlerAccumulator, ctx scoringContext, collector *candidateCollector) []XMLHandlerSymbol {
	result := make([]XMLHandlerSymbol, 0, len(values))
	for _, acc := range values {
		assessment := confidence.Score(buildXMLHandlerEvidence(acc, ctx))
		collector.record(acc.Name, "XML Handler", assessment)
		if !confidence.ShouldPromote(assessment) {
			continue
		}
		level := toPlatformConfidence(assessment.Level)
		result = append(result, XMLHandlerSymbol{
			Name:            acc.Name,
			Confidence:      level,
			RawScore:        assessment.RawScore,
			Score:           assessment.FinalScore,
			Signals:         assessment.Signals,
			Rationale:       assessment.Rationale,
			Evidence:        assessment.Evidence,
			Description:     describeXMLHandler(acc.Name, len(acc.Addons), mapKeys(acc.ElementTypes)),
			ExpectedBinding: inferXMLBinding(acc.Name),
			ElementTypes:    mapKeys(acc.ElementTypes),
			SeenIn:          mapKeys(acc.Addons),
			Examples:        firstUsageExamples(acc.Examples, 6),
			Notes:           xmlHandlerNotes(acc.Name),
		})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func finalizeFieldSymbols(values map[string]*fieldAccumulator, namespace string, ctx scoringContext, collector *candidateCollector) []FieldSymbol {
	result := make([]FieldSymbol, 0, len(values))
	for _, acc := range values {
		assessment := confidence.Score(buildFieldEvidence(acc, namespace, ctx))
		collector.record(acc.Name, namespace+" Field", assessment)
		if !confidence.ShouldPromote(assessment) {
			continue
		}
		level := toPlatformConfidence(assessment.Level)
		result = append(result, FieldSymbol{
			Name:        acc.Name,
			Confidence:  level,
			RawScore:    assessment.RawScore,
			Score:       assessment.FinalScore,
			Signals:     assessment.Signals,
			Rationale:   assessment.Rationale,
			Evidence:    assessment.Evidence,
			Description: describeField(acc.Name, namespace, len(acc.Addons), mapKeys(acc.Contexts)),
			SeenIn:      mapKeys(acc.Addons),
			Notes:       fieldNotes(acc),
		})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func finalizeTableSymbols(values map[string]*tableAccumulator, ctx scoringContext, collector *candidateCollector) []TableSymbol {
	result := make([]TableSymbol, 0, len(values))
	for _, acc := range values {
		if acc.Name == "SystemData" || acc.Name == "GameData" {
			continue
		}
		if !strings.HasPrefix(acc.Name, "EA_") && !sharedTableRoots[acc.Name] && !explicitPlatformTables[acc.Name] {
			continue
		}
		if len(acc.Addons) < 2 && !sharedTableRoots[acc.Name] && !strings.HasPrefix(acc.Name, "EA_") && acc.Name != "Window" && acc.Name != "wstring" && acc.Name != "InterfaceCore" && acc.Name != "DialogManager" && acc.Name != "Icons" && acc.Name != "PartyUtils" && acc.Name != "Cursor" && acc.Name != "HelpTips" && acc.Name != "DefaultColor" {
			continue
		}
		if len(acc.Addons) < 2 && len(acc.Functions) == 0 {
			continue
		}
		assessment := confidence.Score(buildTableEvidence(acc, ctx))
		collector.record(acc.Name, "Global Table", assessment)
		if !confidence.ShouldPromote(assessment) {
			continue
		}
		level := toPlatformConfidence(assessment.Level)
		result = append(result, TableSymbol{
			Name:                 acc.Name,
			Confidence:           level,
			RawScore:             assessment.RawScore,
			Score:                assessment.FinalScore,
			Signals:              assessment.Signals,
			Rationale:            assessment.Rationale,
			Evidence:             assessment.Evidence,
			Description:          describeTable(acc.Name, len(acc.Addons), len(acc.Members), len(acc.Functions)),
			SeenIn:               mapKeys(acc.Addons),
			Functions:            mapKeys(acc.Functions),
			ObservedMembers:      mapKeys(acc.Members),
			RepresentativeUsages: firstUsageExamples(acc.Examples, 6),
			Notes:                graph.UniqueStrings(acc.Notes),
		})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func finalizeConstantSymbols(values map[string]*constantAccumulator, ctx scoringContext, collector *candidateCollector) []ConstantSymbol {
	result := make([]ConstantSymbol, 0, len(values))
	for _, acc := range values {
		assessment := confidence.Score(buildConstantEvidence(acc, ctx))
		collector.record(acc.Name, "Constant", assessment)
		if !confidence.ShouldPromote(assessment) {
			continue
		}
		level := toPlatformConfidence(assessment.Level)
		result = append(result, ConstantSymbol{
			Name:        acc.Name,
			Confidence:  level,
			RawScore:    assessment.RawScore,
			Score:       assessment.FinalScore,
			Signals:     assessment.Signals,
			Rationale:   assessment.Rationale,
			Evidence:    assessment.Evidence,
			Description: describeConstant(acc.Name, len(acc.Addons), len(acc.UsedBy)),
			SeenIn:      mapKeys(acc.Addons),
			UsedBy:      mapKeys(acc.UsedBy),
			Notes:       graph.UniqueStrings(acc.Notes),
		})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func finalizeElementSymbols(values map[string]*elementAccumulator, ctx scoringContext, collector *candidateCollector) []ElementTypeSymbol {
	result := make([]ElementTypeSymbol, 0, len(values))
	for _, acc := range values {
		assessment := confidence.Score(buildElementEvidence(acc, ctx))
		collector.record(acc.Name, "XML Element Type", assessment)
		if !confidence.ShouldPromote(assessment) {
			continue
		}
		level := toPlatformConfidence(assessment.Level)
		symbol := ElementTypeSymbol{
			Name:                    acc.Name,
			Confidence:              level,
			RawScore:                assessment.RawScore,
			Score:                   assessment.FinalScore,
			Signals:                 assessment.Signals,
			Rationale:               assessment.Rationale,
			Evidence:                assessment.Evidence,
			Description:             "",
			SeenIn:                  mapKeys(acc.Addons),
			CommonAttributes:        topKeysByCount(acc.Attributes, 12),
			CommonHandlers:          topKeysByCount(acc.Handlers, 12),
			CommonHandlerFunctions:  topKeysByCount(acc.HandlerFunctions, 12),
			CommonInherits:          topKeysByCount(acc.Inherits, 12),
			CommonChildTypes:        topKeysByCount(acc.ChildTypes, 8),
			CommonChildElementTypes: topKeysByCount(acc.ChildElementTypes, 8),
			CommonParentTypes:       topKeysByCount(acc.ParentTypes, 6),
			CompositionSnippet:      bestCompositionSnippet(acc.Snippets),
			XMLEventBindings:        buildXMLEventBindings(acc.HandlerBindings),
			Examples:                firstUsageExamples(acc.Examples, 6),
			Notes:                   nil,
		}
		symbol.Description = inferElementDescription(symbol)
		result = append(result, symbol)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func finalizeLifecycle(values map[string]*lifecycleAccumulator) []LifecyclePhase {
	order := []string{"manifest", "initialize", "saved-variables", "xml", "runtime", "update", "shutdown"}
	seen := map[string]bool{}
	result := []LifecyclePhase{}
	appendPhase := func(name string) {
		acc, ok := values[name]
		if !ok || seen[name] {
			return
		}
		seen[name] = true
		result = append(result, LifecyclePhase{
			Phase:       name,
			Description: describeLifecycle(name, acc),
			SeenInCount: len(acc.Addons),
			Evidence:    firstStrings(graph.UniqueStrings(append(acc.Evidence, acc.Details...)), 12),
		})
	}
	for _, phase := range order {
		appendPhase(phase)
	}
	for _, phase := range mapKeys(values) {
		appendPhase(phase)
	}
	return result
}

func buildSlashPatterns(functions map[string]*functionAccumulator) []PatternDoc {
	patterns := []PatternDoc{}
	if acc, ok := functions["LibSlash.RegisterSlashCmd"]; ok {
		patterns = append(patterns, PatternDoc{
			Name:        "LibSlash registration",
			Category:    "slash-commands",
			Confidence:  confidenceForFunction("Global Function", len(acc.Addons), acc.Name),
			Description: "Observed slash commands being registered through the shared LibSlash table.",
			Evidence:    examplesToEvidence(firstFunctionExamples(acc.Name, acc.Observations, 6)),
		})
	}
	if acc, ok := functions["LibSlash.UnregisterSlashCmd"]; ok {
		patterns = append(patterns, PatternDoc{
			Name:        "LibSlash cleanup",
			Category:    "slash-commands",
			Confidence:  confidenceForFunction("Global Function", len(acc.Addons), acc.Name),
			Description: "Observed slash commands being explicitly unregistered during addon shutdown paths.",
			Evidence:    examplesToEvidence(firstFunctionExamples(acc.Name, acc.Observations, 6)),
		})
	}
	return patterns
}

func buildWindowPatterns(globalFunctions map[string]*functionAccumulator, windowFunctions map[string]*functionAccumulator, bindings []BindingDoc) []PatternDoc {
	patterns := []PatternDoc{}
	if acc, ok := globalFunctions["CreateWindow"]; ok {
		patterns = append(patterns, PatternDoc{
			Name:        "Window creation",
			Category:    "window-api",
			Confidence:  confidenceForFunction("Global Function", len(acc.Addons), acc.Name),
			Description: "Observed top-level UI windows being created from XML definitions at initialization time.",
			Evidence:    examplesToEvidence(firstFunctionExamples(acc.Name, acc.Observations, 6)),
		})
	}
	if acc, ok := globalFunctions["CreateWindowFromTemplate"]; ok {
		patterns = append(patterns, PatternDoc{
			Name:        "Template instantiation",
			Category:    "window-api",
			Confidence:  confidenceForFunction("Global Function", len(acc.Addons), acc.Name),
			Description: "Observed repeated UI elements being instantiated from XML templates at runtime.",
			Evidence:    examplesToEvidence(firstFunctionExamples(acc.Name, acc.Observations, 6)),
		})
	}
	if acc, ok := windowFunctions["WindowAddAnchor"]; ok {
		patterns = append(patterns, PatternDoc{
			Name:        "Runtime anchoring",
			Category:    "window-api",
			Confidence:  confidenceForFunction("Window Function", len(acc.Addons), acc.Name),
			Description: "Observed layouts being finalized in Lua by clearing and re-adding anchors after window creation.",
			Evidence:    examplesToEvidence(firstFunctionExamples(acc.Name, acc.Observations, 6)),
		})
	}
	if len(bindings) > 0 {
		evidence := []string{}
		for _, binding := range bindings {
			evidence = append(evidence, binding.Addon+": "+binding.Frame+"."+binding.Event+" -> "+binding.LuaFunction)
			if len(evidence) == 6 {
				break
			}
		}
		patterns = append(patterns, PatternDoc{
			Name:        "XML to Lua binding",
			Category:    "window-api",
			Confidence:  ConfidenceHigh,
			Description: "Observed XML handler events routing directly into addon Lua functions through shared engine bindings.",
			Evidence:    evidence,
		})
	}
	return patterns
}

func buildEventPatterns(globalFunctions map[string]*functionAccumulator, gameEvents []EventSymbol, windowEvents []EventSymbol) []PatternDoc {
	patterns := []PatternDoc{}
	for _, name := range []string{"RegisterEventHandler", "UnregisterEventHandler", "WindowRegisterEventHandler", "WindowRegisterCoreEventHandler", "BroadcastEvent"} {
		acc, ok := globalFunctions[name]
		if !ok {
			continue
		}
		patterns = append(patterns, PatternDoc{
			Name:        name,
			Category:    "events",
			Confidence:  confidenceForFunction("Global Function", len(acc.Addons), acc.Name),
			Description: describeFunction(name, "Global Function", len(acc.Addons)),
			Evidence:    examplesToEvidence(firstFunctionExamples(acc.Name, acc.Observations, 6)),
		})
	}
	if len(gameEvents) > 0 {
		patterns = append(patterns, PatternDoc{
			Name:        "Game event fan-in",
			Category:    "events",
			Confidence:  ConfidenceHigh,
			Description: "Observed multiple addons converging on a small set of runtime events such as LOADING_END, RELOAD_INTERFACE, and player update events.",
			Evidence:    eventSymbolNames(gameEvents, 8),
		})
	}
	if len(windowEvents) > 0 {
		patterns = append(patterns, PatternDoc{
			Name:        "Window event hooks",
			Category:    "events",
			Confidence:  ConfidenceHigh,
			Description: "Observed XML and window-scoped handlers using On* hooks to bridge engine UI events into Lua.",
			Evidence:    eventSymbolNames(windowEvents, 8),
		})
	}
	return patterns
}

func buildConventions(corpus Corpus, input contractSemanticInput) []PatternDoc {
	patterns := []PatternDoc{}
	seedRuntimeEvidence, seedListEvidence := loadXMLConventionSeedEvidence()

	patterns = append(patterns, PatternDoc{
		Name:        "Initialization pattern",
		Category:    "conventions",
		Confidence:  ConfidenceHigh,
		Description: "Addons consistently move from manifest loading into initialize hooks, then into XML creation and runtime event registration.",
		Evidence:    lifecycleEvidence(corpus.Lifecycle, "initialize", "runtime", "xml"),
	})
	patterns = append(patterns, PatternDoc{
		Name:        "Event registration pattern",
		Category:    "conventions",
		Confidence:  ConfidenceHigh,
		Description: "Runtime events are typically wired through RegisterEventHandler or window-scoped variants, with handler names staying module-qualified.",
		Evidence:    combinePatternEvidence(corpus.EventPatterns, 8),
	})
	patterns = append(patterns, PatternDoc{
		Name:        "UI creation pattern",
		Category:    "conventions",
		Confidence:  ConfidenceHigh,
		Description: "UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.",
		Evidence:    combinePatternEvidence(corpus.WindowPatterns, 8),
	})
	patterns = append(patterns, PatternDoc{
		Name:        "XML to Lua binding pattern",
		Category:    "conventions",
		Confidence:  ConfidenceHigh,
		Description: "XML handler names map directly to Lua functions and can be cross-checked through the bindings page.",
		Evidence:    bindingEvidence(input.Bindings, 8),
	})
	patterns = append(patterns, PatternDoc{
		Name:        "XML runtime caveats",
		Category:    "conventions",
		Confidence:  ConfidenceMedium,
		Description: "Implementation-validated findings show that XML input, anchoring, and scroll layout behavior can depend on ancestor state, stable parent containers, and outer-window sizing even when child nodes appear correctly configured.",
		Evidence: uniqueNonEmptyStrings(append([]string{
			"WhoHealedMe: a child `OnLButtonUp` target remained inert until the parent or template input chain was made input-enabled.",
			"guidance: validate ancestor `handleinput` state across the clickable chain, not only on the child node.",
			"caveat: treat this as a reusable runtime warning, not a guaranteed engine contract.",
			"WhoHealedMe: nested scroll content dimensions initially under-reported usable space during early layout.",
			"guidance: compute size from the outer parent first, then resize child content and call `ScrollWindowUpdateScrollRect`.",
		}, seedRuntimeEvidence...)),
	})
	patterns = append(patterns, PatternDoc{
		Name:        "XML list binding pattern",
		Category:    "conventions",
		Confidence:  ConfidenceMedium,
		Description: "ListBox rows are commonly bound through ListData-backed Lua tables, with ListColumns supplying text fields and Lua population callbacks handling extra row setup such as icons or reordered display.",
		Evidence: uniqueNonEmptyStrings(append([]string{
			"QuickTacticSwitch: `ListData table=\"QTS.listDisplayData\" populationfunction=\"QTS.PopulateDisplay\"` binds a ListBox to Lua-backed row data.",
			"QuickTacticSwitch: `ListColumns` binds `Name` and `Enemy`, while `QTS.PopulateDisplay` uses `QuickTacticSwitchWindowList.PopulatorIndices` to populate row icons.",
			"QuickTacticSwitch: `ListBoxSetDisplayOrder` and `ListBoxGetDataIndex` are used to manage visible ordering and row-to-data mapping.",
			"AggroMeter: `ListData table=\"AggroMeter.Listdata\" populationfunction=\"\"` suggests column-only text binding works without a custom population callback.",
		}, seedListEvidence...)),
	})
	return patterns
}

func loadXMLConventionSeedEvidence() (runtimeEvidence []string, listEvidence []string) {
	candidates := []string{
		"docs/platform/seeds/xml_conventions.md",
		"/workspace/docs/platform/seeds/xml_conventions.md",
	}

	var content string
	for _, candidate := range candidates {
		data, err := os.ReadFile(candidate)
		if err == nil {
			content = string(data)
			break
		}
	}
	if content == "" {
		return nil, nil
	}

	section := ""
	inPromotedObservation := false
	for _, rawLine := range strings.Split(content, "\n") {
		line := strings.TrimSpace(rawLine)
		if strings.HasPrefix(line, "## ") {
			section = strings.TrimSpace(strings.TrimPrefix(line, "## "))
			inPromotedObservation = false
			continue
		}
		if strings.HasPrefix(line, "<!-- OBSERVATION:") {
			inPromotedObservation = true
			continue
		}

		switch {
		case inPromotedObservation:
			if evidence, ok := seedEvidenceLine(line); ok {
				runtimeEvidence = append(runtimeEvidence, evidence)
			}
		case section == "Observed Layout Caveats":
			if evidence, ok := seedEvidenceLine(line); ok {
				runtimeEvidence = append(runtimeEvidence, evidence)
			}
		case section == "Observed List Binding Patterns":
			if evidence, ok := seedEvidenceLine(line); ok {
				listEvidence = append(listEvidence, evidence)
			}
		}
	}

	return uniqueNonEmptyStrings(runtimeEvidence), uniqueNonEmptyStrings(listEvidence)
}

func seedEvidenceLine(line string) (string, bool) {
	if line == "" || strings.HasPrefix(line, ">") {
		return "", false
	}
	if strings.HasPrefix(line, "- ") {
		return strings.TrimSpace(strings.TrimPrefix(line, "- ")), true
	}
	return "", false
}

func uniqueNonEmptyStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := make(map[string]bool, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	return result
}

func toPlatformConfidence(level confidence.Level) Confidence {
	switch level {
	case confidence.LevelHigh:
		return ConfidenceHigh
	case confidence.LevelMedium:
		return ConfidenceMedium
	default:
		return ConfidenceLow
	}
}

func buildFunctionEvidence(acc *functionAccumulator, category string, ctx scoringContext, observedArity int) confidence.Evidence {
	files := map[string]bool{}
	exampleLocations := []string{}
	sourceKinds := map[string]bool{}
	for _, observation := range acc.Observations {
		if observation.Source != "" {
			files[observation.Source] = true
		}
		exampleLocations = append(exampleLocations, observation.Addon+": "+observation.Caller)
		sourceKinds["lua_calls"] = true
	}
	xmlUsageCount := 0
	xmlAttributeUsageCount := 0
	for _, handler := range ctx.input.Handlers {
		if handler.Function != acc.Name {
			continue
		}
		xmlUsageCount++
		xmlAttributeUsageCount++
		if handler.Source != "" {
			files[handler.Source] = true
		}
		exampleLocations = append(exampleLocations, handler.Addon+": "+handler.Frame+"."+handler.Event)
		sourceKinds["xml_handlers"] = true
	}
	for _, binding := range ctx.input.Bindings {
		if binding.LuaFunction != acc.Name && binding.XMLFunction != acc.Name {
			continue
		}
		xmlUsageCount++
		xmlAttributeUsageCount++
		exampleLocations = append(exampleLocations, binding.Addon+": "+binding.Frame+"."+binding.Event)
		sourceKinds["bindings"] = true
	}
	knownEngineNamespace := isKnownEngineNamespace(acc.Name, category)
	return confidence.Evidence{
		SymbolName:                  acc.Name,
		SymbolKind:                  category,
		AddonsSeenIn:                mapKeys(acc.Addons),
		FilesSeenIn:                 mapKeys(files),
		XMLUsageCount:               xmlUsageCount,
		LuaUsageCount:               len(acc.Observations),
		GlobalUsageCount:            len(acc.Observations),
		LocalDefinitionCount:        ctx.localDefinitionCounts[acc.Name],
		NamespacesDetected:          detectNamespaces(acc.Name),
		KnownEngineNamespace:        knownEngineNamespace,
		DefaultUIPresence:           hasDefaultUIPresence(acc.Name, category),
		EventBindingPresence:        isEventBindingSymbol(acc.Name),
		ExampleLocations:            firstStrings(graph.UniqueStrings(exampleLocations), 16),
		XMLAttributeUsageCount:      xmlAttributeUsageCount,
		DocumentationReferenceCount: documentationReferenceCount(sourceKinds, "lua_calls"),
		TOCInitReferenceCount:       0,
		UsedByXMLAndLua:             xmlUsageCount > 0 && len(acc.Observations) > 0,
		ConsistentRole:              len(acc.Addons) >= 2,
		ConsistentArguments:         hasConsistentArgumentPattern(acc.Observations),
		ConsistentReturns:           hasConsistentReturnPattern(acc.Name, observedArity, len(acc.Observations)),
		SlashCommandPresence:        isSlashSymbol(acc.Name),
		WeakUsageOnly:               len(acc.Observations) <= 1 && len(acc.Addons) <= 1 && xmlUsageCount == 0 && !knownEngineNamespace,
		ProjectSpecificName:         hasAddonSpecificName(acc.Name, ctx.addonNames),
		PlaceholderLikeName:         isPlaceholderLikeName(acc.Name, category),
		ConflictingSignatures:       hasConflictingSignatures(acc.Observations),
		ConflictingRoles:            false,
		WrapperLikely:               likelyWrapper(ctx.localDefinitionCounts[acc.Name], len(acc.Observations), xmlUsageCount, knownEngineNamespace),
		NeverOutsideLocalGraph:      len(acc.Addons) <= 1 && len(sourceKinds) == 1 && sourceKinds["lua_calls"] && !knownEngineNamespace,
		LocalHelperOnly:             ctx.localDefinitionCounts[acc.Name] > 0 && len(sourceKinds) == 1 && sourceKinds["lua_calls"],
		SourceKinds:                 mapKeys(sourceKinds),
	}
}

func buildEventEvidence(acc *eventAccumulator, category string, ctx scoringContext) confidence.Evidence {
	files := map[string]bool{}
	exampleLocations := []string{}
	sourceKinds := map[string]bool{"event_page": true}
	luaUsageCount := 0
	xmlUsageCount := 0
	xmlAttributeUsageCount := 0
	for _, doc := range ctx.input.Functions {
		for _, event := range doc.EventRegistrations {
			if event.Event != acc.Name {
				continue
			}
			luaUsageCount++
			if doc.Source != "" {
				files[doc.Source] = true
			}
			exampleLocations = append(exampleLocations, doc.Addon+": "+doc.Name)
			sourceKinds["lua_event_registration"] = true
		}
	}
	for _, handler := range ctx.input.Handlers {
		if handler.Event != acc.Name {
			continue
		}
		xmlUsageCount++
		xmlAttributeUsageCount++
		if handler.Source != "" {
			files[handler.Source] = true
		}
		exampleLocations = append(exampleLocations, handler.Addon+": "+handler.Frame+"."+handler.Event)
		sourceKinds["xml_handlers"] = true
	}
	knownEngineNamespace := isKnownEngineNamespace(acc.Name, category)
	return confidence.Evidence{
		SymbolName:                  acc.Name,
		SymbolKind:                  category,
		AddonsSeenIn:                mapKeys(acc.Addons),
		FilesSeenIn:                 mapKeys(files),
		XMLUsageCount:               xmlUsageCount,
		LuaUsageCount:               luaUsageCount,
		GlobalUsageCount:            luaUsageCount,
		LocalDefinitionCount:        0,
		NamespacesDetected:          detectNamespaces(acc.Name),
		KnownEngineNamespace:        knownEngineNamespace,
		DefaultUIPresence:           hasDefaultUIPresence(acc.Name, category),
		EventBindingPresence:        true,
		ExampleLocations:            firstStrings(graph.UniqueStrings(exampleLocations), 16),
		XMLAttributeUsageCount:      xmlAttributeUsageCount,
		DocumentationReferenceCount: documentationReferenceCount(sourceKinds, "event_page"),
		TOCInitReferenceCount:       0,
		UsedByXMLAndLua:             xmlUsageCount > 0 && luaUsageCount > 0,
		ConsistentRole:              len(acc.Addons) >= 2,
		ConsistentArguments:         false,
		ConsistentReturns:           false,
		SlashCommandPresence:        false,
		WeakUsageOnly:               (luaUsageCount+xmlUsageCount) <= 1 && len(acc.Addons) <= 1 && !knownEngineNamespace,
		ProjectSpecificName:         hasAddonSpecificName(acc.Name, ctx.addonNames),
		PlaceholderLikeName:         isPlaceholderLikeName(acc.Name, category),
		ConflictingSignatures:       false,
		ConflictingRoles:            category == "Game Event" && isWindowEventName(acc.Name),
		WrapperLikely:               false,
		NeverOutsideLocalGraph:      len(acc.Addons) <= 1 && len(sourceKinds) == 1 && sourceKinds["event_page"] && !knownEngineNamespace,
		LocalHelperOnly:             false,
		SourceKinds:                 mapKeys(sourceKinds),
	}
}

func buildXMLHandlerEvidence(acc *xmlHandlerAccumulator, ctx scoringContext) confidence.Evidence {
	files := map[string]bool{}
	exampleLocations := []string{}
	sourceKinds := map[string]bool{"xml_handlers": true}
	xmlUsageCount := 0
	luaUsageCount := 0
	for _, handler := range ctx.input.Handlers {
		if handler.Event != acc.Name {
			continue
		}
		xmlUsageCount++
		if handler.Source != "" {
			files[handler.Source] = true
		}
		exampleLocations = append(exampleLocations, handler.Addon+": "+handler.Frame+"."+handler.Event)
	}
	for _, binding := range ctx.input.Bindings {
		if binding.Event != acc.Name {
			continue
		}
		luaUsageCount++
		exampleLocations = append(exampleLocations, binding.Addon+": "+binding.Frame+"."+binding.Event)
		sourceKinds["bindings"] = true
	}
	return confidence.Evidence{
		SymbolName:                  acc.Name,
		SymbolKind:                  "XML Handler",
		AddonsSeenIn:                mapKeys(acc.Addons),
		FilesSeenIn:                 mapKeys(files),
		XMLUsageCount:               xmlUsageCount,
		LuaUsageCount:               luaUsageCount,
		GlobalUsageCount:            0,
		LocalDefinitionCount:        0,
		NamespacesDetected:          detectNamespaces(acc.Name),
		KnownEngineNamespace:        false,
		DefaultUIPresence:           false,
		EventBindingPresence:        true,
		ExampleLocations:            firstStrings(graph.UniqueStrings(exampleLocations), 16),
		XMLAttributeUsageCount:      xmlUsageCount,
		DocumentationReferenceCount: documentationReferenceCount(sourceKinds, "xml_handlers"),
		TOCInitReferenceCount:       0,
		UsedByXMLAndLua:             xmlUsageCount > 0 && luaUsageCount > 0,
		ConsistentRole:              len(acc.Addons) >= 2,
		ConsistentArguments:         false,
		ConsistentReturns:           false,
		SlashCommandPresence:        false,
		WeakUsageOnly:               xmlUsageCount <= 1 && len(acc.Addons) <= 1,
		ProjectSpecificName:         !strings.HasPrefix(acc.Name, "On") && hasAddonSpecificName(acc.Name, ctx.addonNames),
		PlaceholderLikeName:         isPlaceholderLikeName(acc.Name, "XML Handler"),
		ConflictingSignatures:       false,
		ConflictingRoles:            false,
		WrapperLikely:               false,
		NeverOutsideLocalGraph:      false,
		LocalHelperOnly:             false,
		SourceKinds:                 mapKeys(sourceKinds),
	}
}

func buildFieldEvidence(acc *fieldAccumulator, namespace string, ctx scoringContext) confidence.Evidence {
	sourceKinds := map[string]bool{}
	for context := range acc.Contexts {
		if context == "lua_call" || context == "event_registration" || context == "flow" || context == "event_page" {
			sourceKinds[context] = true
		}
	}
	if len(sourceKinds) == 0 {
		sourceKinds["observed"] = true
	}
	return confidence.Evidence{
		SymbolName:                  acc.Name,
		SymbolKind:                  namespace + " Field",
		AddonsSeenIn:                mapKeys(acc.Addons),
		FilesSeenIn:                 mapKeys(acc.Files),
		XMLUsageCount:               0,
		LuaUsageCount:               len(acc.Contexts),
		GlobalUsageCount:            len(acc.Contexts),
		LocalDefinitionCount:        0,
		NamespacesDetected:          detectNamespaces(acc.Name),
		KnownEngineNamespace:        true,
		DefaultUIPresence:           true,
		EventBindingPresence:        strings.HasPrefix(acc.Name, "SystemData.Events."),
		ExampleLocations:            firstStrings(mapKeys(acc.Contexts), 16),
		XMLAttributeUsageCount:      0,
		DocumentationReferenceCount: documentationReferenceCount(sourceKinds, "observed"),
		TOCInitReferenceCount:       0,
		UsedByXMLAndLua:             false,
		ConsistentRole:              len(acc.Addons) >= 2,
		ConsistentArguments:         false,
		ConsistentReturns:           false,
		SlashCommandPresence:        false,
		WeakUsageOnly:               len(acc.Addons) <= 1 && len(acc.Contexts) <= 1,
		ProjectSpecificName:         false,
		PlaceholderLikeName:         isPlaceholderLikeName(acc.Name, namespace+" Field"),
		ConflictingSignatures:       false,
		ConflictingRoles:            false,
		WrapperLikely:               false,
		NeverOutsideLocalGraph:      false,
		LocalHelperOnly:             false,
		SourceKinds:                 mapKeys(sourceKinds),
	}
}

func buildTableEvidence(acc *tableAccumulator, ctx scoringContext) confidence.Evidence {
	files := usageFiles(acc.Examples)
	sourceKinds := map[string]bool{"lua_calls": true}
	knownEngineNamespace := isKnownEngineNamespace(acc.Name, "Global Table")
	localDefinitionCount := 0
	return confidence.Evidence{
		SymbolName:                  acc.Name,
		SymbolKind:                  "Global Table",
		AddonsSeenIn:                mapKeys(acc.Addons),
		FilesSeenIn:                 mapKeys(files),
		XMLUsageCount:               0,
		LuaUsageCount:               len(acc.Examples),
		GlobalUsageCount:            len(acc.Functions),
		LocalDefinitionCount:        localDefinitionCount,
		NamespacesDetected:          detectNamespaces(acc.Name),
		KnownEngineNamespace:        knownEngineNamespace,
		DefaultUIPresence:           hasDefaultUIPresence(acc.Name, "Global Table"),
		EventBindingPresence:        false,
		ExampleLocations:            firstStrings(exampleLocationsFromExamples(acc.Examples), 16),
		XMLAttributeUsageCount:      0,
		DocumentationReferenceCount: documentationReferenceCount(sourceKinds, "lua_calls"),
		TOCInitReferenceCount:       0,
		UsedByXMLAndLua:             false,
		ConsistentRole:              len(acc.Addons) >= 2 || knownEngineNamespace,
		ConsistentArguments:         false,
		ConsistentReturns:           false,
		SlashCommandPresence:        strings.HasPrefix(acc.Name, "LibSlash"),
		WeakUsageOnly:               len(acc.Examples) <= 1 && len(acc.Addons) <= 1 && !knownEngineNamespace,
		ProjectSpecificName:         hasAddonSpecificName(acc.Name, ctx.addonNames),
		PlaceholderLikeName:         isPlaceholderLikeName(acc.Name, "Global Table"),
		ConflictingSignatures:       false,
		ConflictingRoles:            false,
		WrapperLikely:               false,
		NeverOutsideLocalGraph:      len(acc.Addons) <= 1 && len(sourceKinds) == 1 && sourceKinds["lua_calls"] && !knownEngineNamespace,
		LocalHelperOnly:             localDefinitionCount > 0 && len(acc.Addons) <= 1 && !knownEngineNamespace,
		SourceKinds:                 mapKeys(sourceKinds),
	}
}

func buildConstantEvidence(acc *constantAccumulator, ctx scoringContext) confidence.Evidence {
	knownEngineNamespace := isKnownEngineNamespace(acc.Name, "Constant")
	sourceKinds := map[string]bool{"xml_attributes": true}
	return confidence.Evidence{
		SymbolName:                  acc.Name,
		SymbolKind:                  "Constant",
		AddonsSeenIn:                mapKeys(acc.Addons),
		FilesSeenIn:                 mapKeys(acc.Files),
		XMLUsageCount:               len(acc.UsedBy),
		LuaUsageCount:               0,
		GlobalUsageCount:            0,
		LocalDefinitionCount:        0,
		NamespacesDetected:          detectNamespaces(acc.Name),
		KnownEngineNamespace:        knownEngineNamespace,
		DefaultUIPresence:           hasDefaultUIPresence(acc.Name, "Constant"),
		EventBindingPresence:        false,
		ExampleLocations:            firstStrings(mapKeys(acc.UsedBy), 16),
		XMLAttributeUsageCount:      len(acc.UsedBy),
		DocumentationReferenceCount: documentationReferenceCount(sourceKinds, "xml_attributes"),
		TOCInitReferenceCount:       0,
		UsedByXMLAndLua:             false,
		ConsistentRole:              len(acc.Addons) >= 2 || knownEngineNamespace,
		ConsistentArguments:         false,
		ConsistentReturns:           false,
		SlashCommandPresence:        false,
		WeakUsageOnly:               len(acc.Addons) <= 1 && !knownEngineNamespace,
		ProjectSpecificName:         hasAddonSpecificName(acc.Name, ctx.addonNames),
		PlaceholderLikeName:         isPlaceholderLikeName(acc.Name, "Constant"),
		ConflictingSignatures:       false,
		ConflictingRoles:            false,
		WrapperLikely:               false,
		NeverOutsideLocalGraph:      false,
		LocalHelperOnly:             false,
		SourceKinds:                 mapKeys(sourceKinds),
	}
}

func buildElementEvidence(acc *elementAccumulator, ctx scoringContext) confidence.Evidence {
	files := usageFiles(acc.Examples)
	sourceKinds := map[string]bool{"xml_frames": true}
	if len(acc.Handlers) > 0 {
		sourceKinds["xml_handlers"] = true
	}
	knownEngineNamespace := isKnownEngineNamespace(acc.Name, "XML Element Type")
	return confidence.Evidence{
		SymbolName:                  acc.Name,
		SymbolKind:                  "XML Element Type",
		AddonsSeenIn:                mapKeys(acc.Addons),
		FilesSeenIn:                 mapKeys(files),
		XMLUsageCount:               len(acc.Examples),
		LuaUsageCount:               len(acc.Handlers),
		GlobalUsageCount:            0,
		LocalDefinitionCount:        0,
		NamespacesDetected:          detectNamespaces(acc.Name),
		KnownEngineNamespace:        knownEngineNamespace,
		DefaultUIPresence:           hasDefaultUIPresence(acc.Name, "XML Element Type"),
		EventBindingPresence:        len(acc.Handlers) > 0,
		ExampleLocations:            firstStrings(exampleLocationsFromExamples(acc.Examples), 16),
		XMLAttributeUsageCount:      len(acc.Examples),
		DocumentationReferenceCount: documentationReferenceCount(sourceKinds, "xml_frames"),
		TOCInitReferenceCount:       0,
		UsedByXMLAndLua:             len(acc.Handlers) > 0,
		ConsistentRole:              len(acc.Addons) >= 2 || knownEngineNamespace,
		ConsistentArguments:         false,
		ConsistentReturns:           false,
		SlashCommandPresence:        false,
		WeakUsageOnly:               len(acc.Addons) <= 1 && !knownEngineNamespace,
		ProjectSpecificName:         false,
		PlaceholderLikeName:         isPlaceholderLikeName(acc.Name, "XML Element Type"),
		ConflictingSignatures:       false,
		ConflictingRoles:            false,
		WrapperLikely:               false,
		NeverOutsideLocalGraph:      false,
		LocalHelperOnly:             false,
		SourceKinds:                 mapKeys(sourceKinds),
	}
}

func detectNamespaces(name string) []string {
	parts := []string{}
	trimmed := strings.TrimSpace(name)
	if root := graph.RootSegment(trimmed); root != "" {
		parts = append(parts, root)
	}
	if strings.HasPrefix(trimmed, "SystemData.") {
		parts = append(parts, "SystemData")
	}
	if strings.HasPrefix(trimmed, "GameData.") {
		parts = append(parts, "GameData")
	}
	return graph.UniqueStrings(parts)
}

func isKnownEngineNamespace(name string, category string) bool {
	trimmed := strings.TrimSpace(name)
	root := graph.RootSegment(trimmed)
	if strings.HasPrefix(trimmed, "SystemData.") || strings.HasPrefix(trimmed, "GameData.") || strings.HasPrefix(trimmed, "EA_Window_") || strings.HasPrefix(trimmed, "EA_") {
		return true
	}
	if category == "Window Function" {
		return true
	}
	if category == "XML Element Type" {
		switch trimmed {
		case "Window", "Button", "Label", "DynamicImage", "ComboBox", "ScrollWindow", "StatusBar", "EditBox", "MapDisplay", "SliderBar", "ListBox", "FullResizeImage", "HorizontalResizeImage", "VerticalResizeImage", "ListData", "ListColumns", "ListColumn":
			return true
		}
	}
	if explicitPlatformTables[trimmed] || sharedTableRoots[trimmed] {
		return true
	}
	if windowFunctionRoots[root] {
		return true
	}
	for _, prefix := range windowFunctionPrefixes {
		if strings.HasPrefix(trimmed, prefix) || root == prefix {
			return true
		}
	}
	switch root {
	case "TextLog", "InterfaceCore", "DialogManager", "Icons", "PartyUtils", "Cursor", "HelpTips", "DefaultColor", "LibSlash", "wstring":
		return true
	}
	return false
}

func hasDefaultUIPresence(name string, category string) bool {
	trimmed := strings.TrimSpace(name)
	root := graph.RootSegment(trimmed)
	if strings.HasPrefix(trimmed, "EA_") || strings.HasPrefix(trimmed, "EA_Window_") || trimmed == "Root" {
		return true
	}
	if explicitPlatformTables[trimmed] {
		return true
	}
	if category == "Window Function" || category == "XML Element Type" {
		return true
	}
	if root == "SystemData" || root == "GameData" || root == "InterfaceCore" || root == "DialogManager" {
		return true
	}
	return false
}

func hasAddonSpecificName(name string, addonNames []string) bool {
	trimmed := strings.ToLower(strings.TrimSpace(name))
	if trimmed == "" || isKnownEngineNamespace(name, "") {
		return false
	}
	for _, addonName := range addonNames {
		candidate := strings.ToLower(strings.TrimSpace(addonName))
		if len(candidate) < 4 {
			continue
		}
		if strings.Contains(trimmed, candidate) {
			return true
		}
	}
	return false
}

func isPlaceholderLikeName(name string, category string) bool {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" || trimmed == "..." {
		return true
	}
	if strings.ContainsAny(trimmed, "()\"'+") {
		return true
	}
	placeholderTokens := map[string]bool{
		"e":         true,
		"ev":        true,
		"evt":       true,
		"event":     true,
		"events":    true,
		"fx":        true,
		"f":         true,
		"fn":        true,
		"func":      true,
		"handler":   true,
		"callback":  true,
		"coreevent": true,
	}
	if placeholderTokens[strings.ToLower(trimmed)] {
		return true
	}
	if !strings.Contains(trimmed, ".") && (strings.Contains(category, "Event") || category == "XML Handler") {
		lower := strings.ToLower(trimmed)
		if lower == trimmed && len(trimmed) <= 8 {
			return true
		}
	}
	return false
}

func isEventBindingSymbol(name string) bool {
	return name == "RegisterEventHandler" || name == "UnregisterEventHandler" || name == "WindowRegisterEventHandler" || name == "WindowRegisterCoreEventHandler" || name == "BroadcastEvent"
}

func isSlashSymbol(name string) bool {
	return strings.HasPrefix(name, "LibSlash.") || strings.Contains(name, "SlashCmd")
}

func hasConsistentArgumentPattern(observations []callObservation) bool {
	if len(observations) < 2 {
		return false
	}
	seen := map[int]bool{}
	for _, observation := range observations {
		seen[len(observation.Arguments)] = true
	}
	return len(seen) == 1
}

func hasConflictingSignatures(observations []callObservation) bool {
	if len(observations) < 2 {
		return false
	}
	seen := map[int]bool{}
	for _, observation := range observations {
		seen[len(observation.Arguments)] = true
	}
	return len(seen) > 1
}

func hasConsistentReturnPattern(name string, observedArity int, usageCount int) bool {
	if usageCount < 2 {
		return false
	}
	last := graph.LastSegment(name)
	return strings.HasPrefix(last, "Get") || strings.HasPrefix(last, "Does") || strings.HasPrefix(last, "Is") || (strings.HasPrefix(last, "Has") && observedArity <= 2)
}

func likelyWrapper(localDefinitionCount int, usageCount int, xmlUsageCount int, knownEngineNamespace bool) bool {
	return localDefinitionCount > 0 && usageCount <= 2 && xmlUsageCount == 0 && !knownEngineNamespace
}

func documentationReferenceCount(sourceKinds map[string]bool, primary string) int {
	count := 0
	for kind := range sourceKinds {
		if kind == primary {
			continue
		}
		count++
	}
	return count
}

func usageFiles(examples []UsageExample) map[string]bool {
	files := map[string]bool{}
	for _, example := range examples {
		if example.Source != "" && strings.Contains(example.Source, "/") {
			files[example.Source] = true
		}
	}
	return files
}

func exampleLocationsFromExamples(examples []UsageExample) []string {
	result := []string{}
	for _, example := range examples {
		label := example.Addon
		if example.Caller != "" {
			label += ": " + example.Caller
		}
		result = append(result, label)
	}
	return graph.UniqueStrings(result)
}

func sortCandidateSummaries(values []CandidateSummary) []CandidateSummary {
	sort.Slice(values, func(i, j int) bool {
		if values[i].Score == values[j].Score {
			if values[i].Category == values[j].Category {
				return values[i].Name < values[j].Name
			}
			return values[i].Category < values[j].Category
		}
		return values[i].Score > values[j].Score
	})
	return values
}

func matrixRow(summary CandidateSummary) ConfidenceMatrixRow {
	row := ConfidenceMatrixRow{Name: summary.Name, Category: summary.Category, Score: summary.Score, Confidence: summary.Confidence}
	for _, signal := range summary.Signals {
		switch signal.Key {
		case "addon_spread_4_plus", "addon_spread_2_3":
			row.CrossAddon += signal.Weight
		case "known_engine_namespace":
			row.Namespace += signal.Weight
		case "xml_handler_attribute", "xml_and_lua":
			row.XMLExposure += signal.Weight
		case "event_binding":
			row.EventLinkage += signal.Weight
		case "default_ui_presence", "known_namespace_override":
			row.DefaultUI += signal.Weight
		case "single_addon_local_definition", "single_addon_project_specific", "local_helper_only", "addon_prefix_penalty", "single_internal_usage", "weak_single_usage", "wrapper_penalty", "never_outside_local_graph", "placeholder_or_computed_name":
			row.LocalPenalty += signal.Weight
		case "consistent_arguments", "consistent_returns", "conflicting_signatures":
			row.Signature += signal.Weight
		}
	}
	return row
}

func buildCoverage(corpus Corpus, candidates []CandidateSummary, input contractSemanticInput) Coverage {
	coverage := Coverage{
		SourceCounts: map[string]int{
			"function_docs": len(input.Functions),
			"frame_docs":    len(input.Frames),
			"handler_docs":  len(input.Handlers),
			"event_docs":    len(input.Events),
			"bindings":      len(input.Bindings),
			"flows":         0,
			"examples":      0,
		},
		SymbolCounts: map[string]int{
			"global_functions": len(corpus.GlobalFunctions),
			"window_functions": len(corpus.WindowFunctions),
			"xml_handlers":     len(corpus.XMLHandlers),
			"game_events":      len(corpus.GameEvents),
			"window_events":    len(corpus.WindowEvents),
			"systemdata":       len(corpus.SystemDataFields),
			"gamedata":         len(corpus.GameDataFields),
			"tables":           len(corpus.GlobalTables),
			"constants":        len(corpus.Constants),
			"element_types":    len(corpus.ElementTypes),
		},
		ConfidenceCounts: map[Confidence]int{},
	}
	for _, summary := range candidates {
		coverage.ConfidenceCounts[summary.Confidence]++
		seenIn := len(summary.Evidence.AddonsSeenIn)
		if seenIn <= 1 {
			coverage.SeenOnceCount++
		} else {
			coverage.SeenManyCount++
		}
		if summary.Confidence == ConfidenceLow {
			coverage.UncertainSymbols = append(coverage.UncertainSymbols, summary.Name)
		}
		coverage.ExplanationMatrix = append(coverage.ExplanationMatrix, matrixRow(summary))
		switch summary.Status {
		case CandidateStatusCanonical:
			if summary.Confidence == ConfidenceHigh {
				coverage.HighConfidencePlatform = append(coverage.HighConfidencePlatform, summary)
			} else {
				coverage.MediumConfidenceCandidates = append(coverage.MediumConfidenceCandidates, summary)
			}
		case CandidateStatusRejected:
			coverage.RejectedAddonLocal = append(coverage.RejectedAddonLocal, summary)
		default:
			if summary.Confidence == ConfidenceLow {
				coverage.LowConfidenceSymbols = append(coverage.LowConfidenceSymbols, summary)
			} else {
				coverage.MediumConfidenceCandidates = append(coverage.MediumConfidenceCandidates, summary)
			}
		}
	}
	sort.Strings(coverage.UncertainSymbols)
	if len(coverage.UncertainSymbols) > 24 {
		coverage.UncertainSymbols = coverage.UncertainSymbols[:24]
	}
	coverage.HighConfidencePlatform = sortCandidateSummaries(coverage.HighConfidencePlatform)
	coverage.MediumConfidenceCandidates = sortCandidateSummaries(coverage.MediumConfidenceCandidates)
	coverage.LowConfidenceSymbols = sortCandidateSummaries(coverage.LowConfidenceSymbols)
	coverage.RejectedAddonLocal = sortCandidateSummaries(coverage.RejectedAddonLocal)
	sort.Slice(coverage.ExplanationMatrix, func(i, j int) bool {
		if coverage.ExplanationMatrix[i].Score == coverage.ExplanationMatrix[j].Score {
			return coverage.ExplanationMatrix[i].Name < coverage.ExplanationMatrix[j].Name
		}
		return coverage.ExplanationMatrix[i].Score > coverage.ExplanationMatrix[j].Score
	})
	if len(coverage.ExplanationMatrix) > 40 {
		coverage.ExplanationMatrix = coverage.ExplanationMatrix[:40]
	}
	return coverage
}

func inferSignature(name string, observations []callObservation) (string, []ParameterDoc, int) {
	if len(observations) == 0 {
		return name + "()", nil, 0
	}
	counts := map[int]int{}
	byIndex := map[int][]string{}
	for _, observation := range observations {
		arity := len(observation.Arguments)
		counts[arity]++
		for index, argument := range observation.Arguments {
			byIndex[index] = append(byIndex[index], argument)
		}
	}
	bestArity := 0
	bestCount := -1
	for arity, count := range counts {
		if count > bestCount || (count == bestCount && arity < bestArity) {
			bestArity = arity
			bestCount = count
		}
	}
	parameters := make([]ParameterDoc, 0, bestArity)
	names := make([]string, 0, bestArity)
	for index := 0; index < bestArity; index++ {
		nameHint, role, evidence := inferParameter(name, index, byIndex[index])
		names = append(names, nameHint)
		parameters = append(parameters, ParameterDoc{Name: nameHint, Role: role, Evidence: evidence})
	}
	return name + "(" + strings.Join(names, ", ") + ")", parameters, bestArity
}

func inferParameter(functionName string, index int, samples []string) (string, string, string) {
	trimmedName := functionName
	if strings.HasPrefix(trimmedName, "RegisterEventHandler") || strings.HasPrefix(trimmedName, "UnregisterEventHandler") {
		if index == 0 {
			return "eventId", "Observed as a SystemData or runtime event identifier.", summarizeSamples(samples)
		}
		if index == 1 {
			return "handlerName", "Observed as a Lua handler function reference.", summarizeSamples(samples)
		}
	}
	if trimmedName == "WindowRegisterEventHandler" {
		switch index {
		case 0:
			return "windowName", "Observed as a target window name.", summarizeSamples(samples)
		case 1:
			return "eventId", "Observed as a SystemData event identifier.", summarizeSamples(samples)
		case 2:
			return "handlerName", "Observed as a Lua handler reference.", summarizeSamples(samples)
		}
	}
	if trimmedName == "WindowRegisterCoreEventHandler" {
		switch index {
		case 0:
			return "windowName", "Observed as a target window name.", summarizeSamples(samples)
		case 1:
			return "windowEvent", "Observed as an On* window event string.", summarizeSamples(samples)
		case 2:
			return "handlerName", "Observed as a Lua handler reference.", summarizeSamples(samples)
		}
	}
	if trimmedName == "BroadcastEvent" {
		return "eventId", "Observed as a runtime event identifier dispatched to listeners.", summarizeSamples(samples)
	}
	if trimmedName == "CreateWindowFromTemplate" {
		switch index {
		case 0:
			return "windowName", "Observed as a runtime window instance name.", summarizeSamples(samples)
		case 1:
			return "templateName", "Observed as an XML template name.", summarizeSamples(samples)
		case 2:
			return "parentWindow", "Observed as a parent window name.", summarizeSamples(samples)
		}
	}
	if trimmedName == "CreateWindow" {
		if index == 0 {
			return "windowName", "Observed as a top-level window name.", summarizeSamples(samples)
		}
		if index == 1 {
			return "showOnCreate", "Observed as a boolean visibility flag.", summarizeSamples(samples)
		}
	}
	if strings.HasPrefix(trimmedName, "WindowGet") || strings.HasPrefix(trimmedName, "WindowSet") || trimmedName == "DestroyWindow" || trimmedName == "DoesWindowExist" || trimmedName == "WindowClearAnchors" {
		if index == 0 {
			return "windowName", "Observed as a target window name.", summarizeSamples(samples)
		}
	}
	if trimmedName == "WindowAddAnchor" {
		switch index {
		case 0:
			return "windowName", "Observed as the window being positioned.", summarizeSamples(samples)
		case 1:
			return "point", "Observed as an anchor point string.", summarizeSamples(samples)
		case 2:
			return "relativeTo", "Observed as a reference window name.", summarizeSamples(samples)
		case 3:
			return "relativePoint", "Observed as a reference anchor point string.", summarizeSamples(samples)
		case 4:
			return "offsetX", "Observed as a numeric horizontal offset.", summarizeSamples(samples)
		case 5:
			return "offsetY", "Observed as a numeric vertical offset.", summarizeSamples(samples)
		}
	}
	if trimmedName == "LabelSetText" || trimmedName == "ButtonSetText" || trimmedName == "TextEditBoxSetText" {
		if index == 0 {
			return "windowName", "Observed as a target control name.", summarizeSamples(samples)
		}
		if index == 1 {
			return "text", "Observed as a text or wstring payload.", summarizeSamples(samples)
		}
	}
	if trimmedName == "DynamicImageSetTexture" {
		switch index {
		case 0:
			return "windowName", "Observed as a target image control name.", summarizeSamples(samples)
		case 1:
			return "texture", "Observed as a texture resource name.", summarizeSamples(samples)
		case 2:
			return "textureX", "Observed as a numeric texture coordinate or atlas offset.", summarizeSamples(samples)
		case 3:
			return "textureY", "Observed as a numeric texture coordinate or atlas offset.", summarizeSamples(samples)
		}
	}
	if trimmedName == "LibSlash.RegisterSlashCmd" {
		if index == 0 {
			return "slashName", "Observed as a slash command token.", summarizeSamples(samples)
		}
		if index == 1 {
			return "handler", "Observed as a command handler callback.", summarizeSamples(samples)
		}
	}
	if trimmedName == "LibSlash.UnregisterSlashCmd" {
		return "slashName", "Observed as a slash command token.", summarizeSamples(samples)
	}
	role := inferGenericRole(samples)
	return fmt.Sprintf("arg%d", index+1), role, summarizeSamples(samples)
}

func inferReturns(name string) []string {
	last := graph.LastSegment(name)
	if strings.HasPrefix(last, "Get") || strings.HasPrefix(last, "Does") {
		return []string{"Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone."}
	}
	return []string{"Not confidently inferable from contract artifacts alone."}
}

func inferSideEffects(name string) []string {
	switch {
	case strings.HasPrefix(name, "Register") || strings.Contains(name, "Register"):
		return []string{"Registers engine or library callbacks for later dispatch."}
	case strings.HasPrefix(name, "Unregister") || strings.Contains(name, "Unregister"):
		return []string{"Removes a previously registered callback or command binding."}
	case strings.HasPrefix(name, "Broadcast"):
		return []string{"Dispatches an event into the runtime event system."}
	case strings.HasPrefix(name, "Create"):
		return []string{"Creates or instantiates UI objects from XML-backed definitions."}
	case strings.HasPrefix(name, "Destroy"):
		return []string{"Destroys a runtime-created window instance."}
	case strings.Contains(name, "Set"):
		return []string{"Mutates engine-visible UI state or presentation."}
	case strings.Contains(name, "AddAnchor") || strings.Contains(name, "ClearAnchors"):
		return []string{"Mutates runtime window layout state."}
	case strings.Contains(name, "Print"):
		return []string{"Writes output to the chat or UI surface."}
	case strings.HasPrefix(name, "Send"):
		return []string{"Sends a client action request back into the engine."}
	default:
		return []string{"No side effect is confidently inferable from contract artifacts alone."}
	}
}

func confidenceForFunction(category string, addonCount int, name string) Confidence {
	if addonCount >= 4 {
		return ConfidenceHigh
	}
	if addonCount >= 2 {
		return ConfidenceMedium
	}
	if category == "Window Function" || signalGlobalFunctions[name] || sharedTableRoots[graph.RootSegment(name)] {
		return ConfidenceMedium
	}
	return ConfidenceLow
}

func confidenceForEvent(category string, addonCount int, name string) Confidence {
	if addonCount >= 4 || strings.HasPrefix(name, "SystemData.Events.") {
		return ConfidenceHigh
	}
	if addonCount >= 2 || category == "Window Event" {
		return ConfidenceMedium
	}
	return ConfidenceLow
}

func confidenceForXMLHandler(addonCount int, name string) Confidence {
	if addonCount >= 4 {
		return ConfidenceHigh
	}
	if addonCount >= 2 || strings.HasPrefix(name, "On") {
		return ConfidenceMedium
	}
	return ConfidenceLow
}

func confidenceForField(addonCount int, name string) Confidence {
	if addonCount >= 4 {
		return ConfidenceHigh
	}
	if addonCount >= 2 || strings.HasPrefix(name, "SystemData.Events.") {
		return ConfidenceMedium
	}
	return ConfidenceLow
}

func confidenceForTable(addonCount int, name string) Confidence {
	if addonCount >= 3 || sharedTableRoots[name] {
		return ConfidenceHigh
	}
	if addonCount >= 2 {
		return ConfidenceMedium
	}
	return ConfidenceLow
}

func confidenceForConstant(addonCount int, name string) Confidence {
	if strings.HasPrefix(name, "EA_") || addonCount >= 3 {
		return ConfidenceHigh
	}
	if addonCount >= 2 || name == "Root" {
		return ConfidenceMedium
	}
	return ConfidenceLow
}

func confidenceForElement(addonCount int, name string) Confidence {
	if addonCount >= 4 {
		return ConfidenceHigh
	}
	if addonCount >= 2 || name == "Window" || name == "Button" {
		return ConfidenceMedium
	}
	return ConfidenceLow
}

func describeFunction(name string, category string, addonCount int) string {
	switch {
	case name == "RegisterEventHandler":
		return "Observed registering global runtime handlers against shared event identifiers."
	case name == "UnregisterEventHandler":
		return "Observed removing previously registered global runtime handlers."
	case name == "WindowRegisterEventHandler":
		return "Observed binding SystemData events directly to a named window."
	case name == "WindowRegisterCoreEventHandler":
		return "Observed binding On* window events directly to a named window."
	case name == "BroadcastEvent":
		return "Observed triggering a runtime event so existing handlers are notified."
	case name == "CreateWindow":
		return "Observed creating a top-level XML window from a loaded definition."
	case name == "CreateWindowFromTemplate":
		return "Observed instantiating repeated UI elements from an XML template."
	case name == "DestroyWindow":
		return "Observed tearing down runtime-created windows."
	case name == "DoesWindowExist":
		return "Observed guarding runtime window creation and cleanup by checking whether a named window exists."
	case strings.HasPrefix(name, "WindowAddAnchor"):
		return "Observed positioning windows relative to other runtime UI elements."
	case strings.HasPrefix(name, "WindowClearAnchors"):
		return "Observed resetting a window layout before applying new runtime anchors."
	case strings.HasPrefix(name, "WindowSet"):
		return "Observed mutating runtime window state or presentation."
	case strings.HasPrefix(name, "WindowGet"):
		return "Observed querying runtime window state or metadata."
	case strings.HasPrefix(name, "LabelSet"):
		return "Observed updating label text or label styling on existing controls."
	case strings.HasPrefix(name, "ButtonSet"):
		return "Observed mutating button text or pressed state on existing controls."
	case strings.HasPrefix(name, "DynamicImageSet"):
		return "Observed mutating runtime image resources on existing controls."
	case strings.HasPrefix(name, "TextEditBox"):
		return "Observed reading from or writing to edit-box controls."
	case strings.Contains(name, "SlashCmd"):
		return "Observed wiring slash commands through a shared command-registration table."
	case strings.HasPrefix(name, "Get"):
		return fmt.Sprintf("Observed as a shared query API across %d addons.", addonCount)
	case strings.HasPrefix(name, "Send"):
		return fmt.Sprintf("Observed as a shared action API across %d addons.", addonCount)
	default:
		return fmt.Sprintf("Observed as a %s across %d addons.", strings.ToLower(category), addonCount)
	}
}

func describeEvent(name string, category string, addonCount int, registrarCount int, handlerCount int) string {
	switch category {
	case "Window Event":
		if handlerCount > 1 {
			return fmt.Sprintf("Engine-supplied UI event hook bound by %d addons through window handlers.", addonCount)
		}
		return "Engine-supplied UI event hook for window-scoped event handling."
	case "System Event":
		if registrarCount > handlerCount && handlerCount > 0 {
			return fmt.Sprintf("Shared SystemData runtime event registered by %d addons and handled by %d handler functions.", registrarCount, handlerCount)
		}
		return fmt.Sprintf("Shared SystemData runtime event observed across %d addons.", addonCount)
	default:
		if handlerCount > 0 {
			return fmt.Sprintf("Runtime event with %d handler registrations observed across %d addons.", handlerCount, addonCount)
		}
		return fmt.Sprintf("Observed as a runtime event across %d addons.", addonCount)
	}
}

func describeXMLHandler(name string, addonCount int, elementTypes []string) string {
	if len(elementTypes) > 0 {
		typesStr := strings.Join(firstStrings(elementTypes, 2), ", ")
		if len(elementTypes) > 2 {
			typesStr += " and others"
		}
		return fmt.Sprintf("XML handler event commonly bound to %s elements in %d addons.", typesStr, addonCount)
	}
	return fmt.Sprintf("XML handler event observed across %d addons.", addonCount)
}

func describeField(name string, namespace string, addonCount int, contextTypes []string) string {
	if len(contextTypes) > 0 {
		contextStr := contextTypes[0]
		if len(contextTypes) > 1 {
			contextStr += " and " + strings.Join(contextTypes[1:], ", ")
		}
		return fmt.Sprintf("%s.%s field accessed by %d addons; commonly found in %s contexts.", namespace, name, addonCount, contextStr)
	}
	return fmt.Sprintf("%s field from %s namespace accessed by %d addons.", name, namespace, addonCount)
}

func describeTable(name string, addonCount int, memberCount int, functionCount int) string {
	if memberCount > 0 && functionCount > 0 {
		return fmt.Sprintf("Shared table with %d member functions and %d data members accessed by %d addons.", functionCount, memberCount, addonCount)
	}
	if functionCount > 0 {
		return fmt.Sprintf("Shared function table with %d member functions; the primary API surface for %d addons.", functionCount, addonCount)
	}
	if memberCount > 0 {
		return fmt.Sprintf("Shared data table with %d observed members accessed by %d addons.", memberCount, addonCount)
	}
	return fmt.Sprintf("Shared global table or namespace accessed by %d addons.", addonCount)
}

func describeConstant(name string, addonCount int, usageCount int) string {
	switch {
	case strings.HasPrefix(name, "EA_"):
		if strings.Contains(name, "Color") || strings.Contains(name, "COLOR") {
			return fmt.Sprintf("Engine XML color constant for standard UI styling; inherited by %d addons.", addonCount)
		}
		if strings.HasPrefix(name, "EA_BUTTON") || strings.HasPrefix(name, "EA_WINDOW") {
			return fmt.Sprintf("Engine XML template or element constant referenced by %d addons.", addonCount)
		}
		return fmt.Sprintf("Engine-supplied XML constant or template class referenced by %d addons.", addonCount)
	case usageCount > 1:
		return fmt.Sprintf("Shared constant used by %d function contexts across %d addons.", usageCount, addonCount)
	default:
		return fmt.Sprintf("Shared constant referenced by %d addons.", addonCount)
	}
}

func describeElementFallback(name string) string {
	return name + " is an XML element type observed in contract artifacts, but structured evidence is currently sparse."
}

func inferElementDescription(symbol ElementTypeSymbol) string {
	if text, ok := coreElementDescriptions[symbol.Name]; ok {
		return text
	}

	parents := firstStrings(graph.UniqueStrings(symbol.CommonParentTypes), 2)
	structuralChildren := firstStrings(graph.UniqueStrings(symbol.CommonChildTypes), 3)
	namedChildren := firstStrings(graph.UniqueStrings(symbol.CommonChildElementTypes), 3)
	inherits := firstStrings(graph.UniqueStrings(append(symbol.CommonInherits, symbol.InheritsBases...)), 2)
	handlers := firstStrings(graph.UniqueStrings(symbol.CommonHandlers), 3)
	handlerFns := firstStrings(graph.UniqueStrings(symbol.CommonHandlerFunctions), 2)
	luaManipulators := firstStrings(graph.UniqueStrings(symbol.LuaManipulators), 2)

	bindingEvents := make([]string, 0, len(symbol.XMLEventBindings))
	for _, binding := range symbol.XMLEventBindings {
		event := strings.TrimSpace(binding.Event)
		if event != "" {
			bindingEvents = append(bindingEvents, event)
		}
	}
	bindingEvents = firstStrings(graph.UniqueStrings(bindingEvents), 3)

	if len(parents) == 0 && len(structuralChildren) == 0 && len(namedChildren) == 0 && len(inherits) == 0 && len(handlers) == 0 && len(bindingEvents) == 0 && len(handlerFns) == 0 && len(luaManipulators) == 0 {
		return describeElementFallback(symbol.Name)
	}

	role := "XML UI element"
	if isLikelyStructuralElement(symbol.Name, handlers, bindingEvents, handlerFns, namedChildren) {
		role = "structural XML sub-element"
	} else if hasInputHandlers(append(handlers, bindingEvents...)) {
		role = "interactive XML control"
	} else if len(structuralChildren) > 0 || len(namedChildren) > 0 {
		role = "container-style XML element"
	} else if len(handlers) > 0 || len(bindingEvents) > 0 || len(handlerFns) > 0 {
		role = "event-capable XML element"
	}

	where := ""
	if len(parents) > 0 {
		where = " It commonly appears under " + strings.Join(parents, " and ") + "."
	} else if len(inherits) > 0 {
		where = " It is commonly instantiated from " + strings.Join(inherits, " and ") + " templates."
	}

	doesParts := []string{}
	if len(structuralChildren) > 0 {
		doesParts = append(doesParts, "organize structural children such as "+strings.Join(structuralChildren, ", "))
	}
	if len(namedChildren) > 0 {
		doesParts = append(doesParts, "host named child elements such as "+strings.Join(namedChildren, ", "))
	}
	if len(bindingEvents) > 0 {
		doesParts = append(doesParts, "bind XML events like "+strings.Join(bindingEvents, ", ")+" to Lua")
	} else if len(handlers) > 0 {
		doesParts = append(doesParts, "declare handlers such as "+strings.Join(handlers, ", "))
	}
	if len(handlerFns) > 0 {
		doesParts = append(doesParts, "route callbacks to Lua functions like "+strings.Join(handlerFns, ", "))
	}
	if len(luaManipulators) > 0 {
		doesParts = append(doesParts, "be manipulated from Lua by functions such as "+strings.Join(luaManipulators, ", "))
	}

	what := symbol.Name + " is " + indefiniteRolePhrase(role) + "."

	if len(doesParts) == 0 {
		if len(where) > 0 {
			return what + where
		}
		return what + " " + describeElementFallback(symbol.Name)
	}

	does := " It is typically used to " + strings.Join(firstStrings(doesParts, 2), " and ") + "."
	return what + where + does
}

var coreElementDescriptions = map[string]string{
	"Window":       "Window is the primary XML container element for addon UI frames. It commonly hosts named child elements and binds lifecycle or input events to Lua handlers.",
	"Button":       "Button is an interactive XML control used for click and mouse-driven callbacks. It usually binds input events to Lua handler functions.",
	"Label":        "Label is a display-focused XML element used to present text content in UI layouts. It commonly appears inside container elements such as Window.",
	"Text":         "Text is a structural XML sub-element used to define display text inside parent controls. It commonly appears as a child element in label or button-style definitions.",
	"ListBox":      "ListBox is a container-style XML element for row-based UI lists. It commonly works with structural children such as ListData and ListColumns and binds events to Lua callbacks.",
	"ListData":     "ListData is a structural XML sub-element that binds list controls to Lua-backed row data. It commonly appears under ListBox as list data wiring.",
	"ListColumns":  "ListColumns is a structural XML container that declares list column mappings. It commonly appears under ListBox and contains ListColumn entries.",
	"ListColumn":   "ListColumn is a structural XML sub-element that maps one list field to a target row child element. It commonly appears under ListColumns.",
	"ScrollWindow": "ScrollWindow is a container-style XML element for scrollable UI content. It commonly hosts child elements and binds scroll-related or input handlers to Lua.",
}

func isLikelyStructuralElement(name string, handlers []string, bindingEvents []string, handlerFns []string, namedChildren []string) bool {
	if len(handlers) > 0 || len(bindingEvents) > 0 || len(handlerFns) > 0 {
		return false
	}
	if len(namedChildren) > 0 {
		return false
	}
	knownStructural := map[string]bool{
		"Anchor": true, "Anchors": true, "Dimensions": true, "Size": true,
		"Point": true, "AbsPoint": true, "Text": true, "TextColors": true,
		"Textures": true, "TexCoords": true, "States": true, "Layer": true,
		"Layers": true, "ListData": true, "ListColumns": true, "ListColumn": true,
	}
	if knownStructural[name] {
		return true
	}
	return strings.HasSuffix(name, "s") && !strings.HasPrefix(name, "On")
}

func hasInputHandlers(events []string) bool {
	for _, event := range events {
		lower := strings.ToLower(event)
		if strings.Contains(lower, "click") || strings.Contains(lower, "mouse") || strings.Contains(lower, "button") || strings.Contains(lower, "drag") || strings.Contains(lower, "key") {
			return true
		}
	}
	return false
}

func indefiniteRolePhrase(role string) string {
	if role == "" {
		return "an XML element"
	}
	first := strings.ToLower(string(role[0]))
	if first == "a" || first == "e" || first == "i" || first == "o" || first == "u" {
		return "an " + role
	}
	return "a " + role
}

// bestCompositionSnippet picks the most representative composition snippet from
// all observed frames of an element type. "Most representative" means the one
// with the most XML structure lines (more sub-elements visible), which tends to
// come from the most fully-populated real usage.
func bestCompositionSnippet(snippets []string) string {
	if len(snippets) == 0 {
		return ""
	}
	best := ""
	bestLines := 0
	for _, s := range snippets {
		n := strings.Count(s, "\n")
		if n > bestLines {
			bestLines = n
			best = s
		}
	}
	if best == "" {
		return snippets[0]
	}
	return best
}

func describeLifecycle(phase string, acc *lifecycleAccumulator) string {
	switch phase {
	case "manifest":
		return "Addon manifests declare file load order and bootstrap hooks before Lua runtime logic begins."
	case "saved-variables":
		return "Saved-variable roots appear before normal runtime hooks and provide persistent addon state."
	case "initialize":
		return "Initialization hooks create windows, hydrate settings, and bind runtime callbacks."
	case "xml":
		return "XML windows, templates, and handlers become available as the UI layer is created."
	case "runtime":
		return "Runtime logic pivots into event-driven updates wired through shared event registration APIs."
	case "update":
		return "Optional update hooks provide repeated processing after initialization has completed."
	case "shutdown":
		return "Shutdown hooks unregister commands or handlers and persist addon-owned state."
	default:
		return fmt.Sprintf("Observed lifecycle phase used by %d addons.", len(acc.Addons))
	}
}

func inferEventHandlerPattern(name string) string {
	if isWindowEventName(name) {
		return "Observed as an On* callback routed into a module-qualified Lua function."
	}
	return "Observed as a runtime event ID routed through RegisterEventHandler-style APIs."
}

func inferEventPayload(name string) []string {
	if strings.HasPrefix(name, "SystemData.Events.") {
		return []string{"Payload shape is not inferable from contract artifacts alone; treat this as an engine event identifier."}
	}
	if isWindowEventName(name) {
		return []string{"Window callback arguments are not fully inferable from contract artifacts alone."}
	}
	return []string{"Payload not inferable from addon-level documentation alone."}
}

// knownXMLHandlerArgs maps well-known WAR XML handler event names to their
// expected Lua callback signature string.  Confidence is HIGH for engine-standard
// events whose argument conventions are widely observed across many addons.
var knownXMLHandlerArgs = map[string]string{
	// Lifecycle hooks – no engine-supplied arguments.
	"OnInitialize": "function()",
	"OnShutdown":   "function()",
	"OnShow":       "function()",
	"OnHide":       "function()",
	"OnShown":      "function()",
	"OnHidden":     "function()",
	// Mouse / pointer events.
	"OnMouseOver":   "function()",
	"OnMouseOut":    "function()",
	"OnMouseDown":   "function(button)",
	"OnMouseUp":     "function(button)",
	"OnClick":       "function()",
	"OnDoubleClick": "function()",
	"OnMouseWheel":  "function(delta)",
	// Keyboard events.
	"OnEnterPressed":  "function()",
	"OnEscapePressed": "function()",
	"OnKeyDown":       "function(key)",
	"OnKeyUp":         "function(key)",
	// Value / text change events.
	"OnTextChanged":  "function()",
	"OnValueChanged": "function(value)",
	// Drag events.
	"OnDragStart":   "function()",
	"OnDragStop":    "function()",
	"OnReceiveDrag": "function()",
	// Resize / move events.
	"OnSizeChanged": "function(width, height)",
	"OnMove":        "function()",
	// Per-frame update hook – elapsed time in seconds.
	"OnUpdate": "function(elapsed)",
	// Edit box events.
	"OnCursorChanged":        "function()",
	"OnInputLanguageChanged": "function()",
	// Selection / list events.
	"OnSelectionChanged": "function(selectedRow)",
	"OnScrollChanged":    "function()",
}

func inferXMLBinding(name string) string {
	if sig, ok := knownXMLHandlerArgs[name]; ok {
		return sig
	}
	return "function(...)"
}

// xmlHandlerArgsConfidence returns "HIGH", "MEDIUM", or "LOW" for the inferred
// argument signature of an XML handler event name.
func xmlHandlerArgsConfidence(name string) string {
	if _, ok := knownXMLHandlerArgs[name]; ok {
		return "MEDIUM" // known pattern, but runtime engine args are not verified from source
	}
	return "LOW"
}

// buildXMLEventBindings converts the per-event Lua-function binding counts
// from an elementAccumulator into a sorted slice of XMLEventBinding records.
// Events with no associated Lua functions are skipped.
func buildXMLEventBindings(bindings map[string]map[string]int) []XMLEventBinding {
	if len(bindings) == 0 {
		return nil
	}
	// Order by event name for stable output.
	events := make([]string, 0, len(bindings))
	for event := range bindings {
		if strings.TrimSpace(event) != "" {
			events = append(events, event)
		}
	}
	sort.Strings(events)
	result := make([]XMLEventBinding, 0, len(events))
	for _, event := range events {
		funcCounts := bindings[event]
		topFuncs := topKeysByCount(funcCounts, 6)
		if len(topFuncs) == 0 {
			continue
		}
		result = append(result, XMLEventBinding{
			Event:          event,
			LuaFunctions:   topFuncs,
			InferredArgs:   inferXMLBinding(event),
			ArgsConfidence: xmlHandlerArgsConfidence(event),
		})
	}
	return result
}

func inferGenericRole(samples []string) string {
	counts := map[string]int{"event": 0, "handler": 0, "window": 0, "boolean": 0, "number": 0, "text": 0}
	for _, sample := range samples {
		trimmed := strings.TrimSpace(sample)
		switch {
		case strings.HasPrefix(trimmed, "SystemData.Events."):
			counts["event"]++
		case strings.Contains(trimmed, ".") && !strings.HasPrefix(trimmed, "\""):
			counts["handler"]++
		case trimmed == "true" || trimmed == "false":
			counts["boolean"]++
		case strings.HasPrefix(trimmed, "\"") || strings.HasPrefix(trimmed, "L ") || strings.HasPrefix(trimmed, "L\""):
			counts["text"]++
		case isNumericLike(trimmed):
			counts["number"]++
		default:
			counts["window"]++
		}
	}
	bestKey := "value"
	bestCount := -1
	for key, count := range counts {
		if count > bestCount {
			bestKey = key
			bestCount = count
		}
	}
	switch bestKey {
	case "event":
		return "Observed as an event identifier."
	case "handler":
		return "Observed as a function or method reference."
	case "boolean":
		return "Observed as a boolean toggle."
	case "number":
		return "Observed as a numeric value."
	case "text":
		return "Observed as a text or wstring payload."
	case "window":
		return "Observed as a runtime window or control identifier."
	default:
		return "Observed argument role is uncertain."
	}
}

func summarizeSamples(samples []string) string {
	unique := graph.UniqueStrings(samples)
	if len(unique) == 0 {
		return "No direct examples available."
	}
	return "Observed values: " + strings.Join(firstStrings(unique, 3), ", ")
}

func functionNotes(category string, confidence Confidence, addonCount int) []string {
	notes := []string{}
	if confidence == ConfidenceLow {
		notes = append(notes, "Only weakly supported by the generated addon corpus; verify against in-game behavior.")
	}
	if addonCount < 2 {
		notes = append(notes, "Only one addon surfaced this symbol in the current corpus.")
	}
	if category == "Global Function" {
		notes = append(notes, "Canonical entry built from observed call sites, not from engine source or decompiled definitions.")
	}
	return notes
}

func eventNotes(acc *eventAccumulator) []string {
	notes := []string{}
	if len(acc.TriggeredBy) > 0 {
		notes = append(notes, "Triggered-by evidence: "+strings.Join(mapKeys(acc.TriggeredBy), ", "))
	}
	if len(acc.Addons) < 2 {
		notes = append(notes, "Only one addon surfaced this event in the current addon-api corpus.")
	}
	return notes
}

func xmlHandlerNotes(name string) []string {
	var notes []string
	if _, known := knownXMLHandlerArgs[name]; known {
		conf := xmlHandlerArgsConfidence(name)
		notes = append(notes, "Expected callback signature inferred from common WAR XML handler conventions ("+conf+" confidence).")
	} else {
		notes = append(notes, "Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.")
	}
	if !strings.HasPrefix(name, "On") {
		notes = append(notes, "Non-On* handler names should be reviewed manually; most XML hooks are On* events.")
	}
	return notes
}

func fieldNotes(acc *fieldAccumulator) []string {
	contexts := mapKeys(acc.Contexts)
	if len(contexts) == 0 {
		return nil
	}
	return []string{"Observed in contexts: " + strings.Join(firstStrings(contexts, 6), ", ")}
}

func defaultInferenceRules() []string {
	return []string{
		"Parsed only generated markdown under docs/addon-api; raw addon Lua and XML were not rescanned.",
		"Dropped call-site aliases when a qualified and unqualified symbol shared caller, line, arguments, and final segment.",
		"Filtered addon-local calls when the symbol root or final segment matched addon-owned namespaces or function names.",
		"Promoted symbols to platform API candidates when they appeared in multiple addons, in XML/window handlers, or under known WAR/UI namespaces.",
		"Marked high, medium, or low confidence strictly from observed addon spread and namespace signal, not from external API databases.",
		"Inferred signatures only from repeated argument positions and clear naming signal; uncertain return values remain explicitly marked unknown.",
		"Extracted SystemData and GameData members only from observed tokens in calls, event pages, flow evidence, and generated globals tables.",
	}
}

func ensureFunctionAccumulator(target map[string]*functionAccumulator, name string) *functionAccumulator {
	name = canonicalSymbolName(name)
	acc, ok := target[name]
	if !ok {
		acc = &functionAccumulator{Name: name, Addons: map[string]bool{}, Aliases: map[string]bool{}}
		target[name] = acc
	}
	return acc
}

func ensureEventAccumulator(target map[string]*eventAccumulator, name string) *eventAccumulator {
	name = canonicalEventName(name)
	acc, ok := target[name]
	if !ok {
		acc = &eventAccumulator{Name: name, Addons: map[string]bool{}, Registrars: map[string]bool{}, Handlers: map[string]bool{}, TriggeredBy: map[string]bool{}}
		target[name] = acc
	}
	return acc
}

func ensureXMLHandlerAccumulator(target map[string]*xmlHandlerAccumulator, name string) *xmlHandlerAccumulator {
	name = canonicalEventName(name)
	acc, ok := target[name]
	if !ok {
		acc = &xmlHandlerAccumulator{Name: name, Addons: map[string]bool{}, ElementTypes: map[string]bool{}}
		target[name] = acc
	}
	return acc
}

func ensureFieldAccumulator(target map[string]*fieldAccumulator, name string) *fieldAccumulator {
	acc, ok := target[name]
	if !ok {
		acc = &fieldAccumulator{Name: name, Addons: map[string]bool{}, Contexts: map[string]bool{}, Files: map[string]bool{}}
		target[name] = acc
	}
	return acc
}

func ensureTableAccumulator(target map[string]*tableAccumulator, name string) *tableAccumulator {
	name = canonicalSymbolName(name)
	acc, ok := target[name]
	if !ok {
		acc = &tableAccumulator{Name: name, Addons: map[string]bool{}, Functions: map[string]bool{}, Members: map[string]bool{}}
		target[name] = acc
	}
	return acc
}

func ensureConstantAccumulator(target map[string]*constantAccumulator, name string) *constantAccumulator {
	name = canonicalSymbolName(name)
	acc, ok := target[name]
	if !ok {
		acc = &constantAccumulator{Name: name, Addons: map[string]bool{}, UsedBy: map[string]bool{}, Files: map[string]bool{}}
		target[name] = acc
	}
	return acc
}

func ensureElementAccumulator(target map[string]*elementAccumulator, name string) *elementAccumulator {
	name = canonicalElementTypeName(name)
	acc, ok := target[name]
	if !ok {
		acc = &elementAccumulator{
			Name:              name,
			Addons:            map[string]bool{},
			Attributes:        map[string]int{},
			Handlers:          map[string]int{},
			HandlerFunctions:  map[string]int{},
			HandlerBindings:   map[string]map[string]int{},
			Inherits:          map[string]int{},
			ChildTypes:        map[string]int{},
			ChildElementTypes: map[string]int{},
			ParentTypes:       map[string]int{},
		}
		target[name] = acc
	}
	return acc
}

func canonicalSymbolName(name string) string {
	return strings.TrimSpace(name)
}

func canonicalElementTypeName(name string) string {
	return canonicalSymbolName(name)
}

func canonicalAttributeName(name string) string {
	return canonicalSymbolName(name)
}

func canonicalEventName(name string) string {
	return graph.NormalizeEventName(name)
}

func ensureLifecycleAccumulator(target map[string]*lifecycleAccumulator, phase string) *lifecycleAccumulator {
	acc, ok := target[phase]
	if !ok {
		acc = &lifecycleAccumulator{Phase: phase, Addons: map[string]bool{}}
		target[phase] = acc
	}
	return acc
}

func appendUniqueExample(existing []UsageExample, example UsageExample) []UsageExample {
	key := example.Addon + "|" + example.Caller + "|" + example.Snippet
	for _, item := range existing {
		if item.Addon+"|"+item.Caller+"|"+item.Snippet == key {
			return existing
		}
	}
	return append(existing, example)
}

func appendUniqueString(existing []string, value string) []string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return existing
	}
	for _, item := range existing {
		if item == trimmed {
			return existing
		}
	}
	return append(existing, trimmed)
}

func addToSetMap(target map[string]map[string]bool, key string, value string) {
	set, ok := target[key]
	if !ok {
		set = map[string]bool{}
		target[key] = set
	}
	set[value] = true
}

func mapKeys[T any](source map[string]T) []string {
	keys := make([]string, 0, len(source))
	for key := range source {
		trimmed := strings.TrimSpace(key)
		if trimmed != "" {
			keys = append(keys, trimmed)
		}
	}
	sort.Strings(keys)
	return keys
}

func topKeysByCount(source map[string]int, limit int) []string {
	type pair struct {
		Name  string
		Count int
	}
	pairs := make([]pair, 0, len(source))
	for key, count := range source {
		if strings.TrimSpace(key) == "" {
			continue
		}
		pairs = append(pairs, pair{Name: key, Count: count})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Count == pairs[j].Count {
			return pairs[i].Name < pairs[j].Name
		}
		return pairs[i].Count > pairs[j].Count
	})
	result := make([]string, 0, limit)
	for _, item := range pairs {
		result = append(result, item.Name)
		if len(result) == limit {
			break
		}
	}
	return result
}

func firstFunctionExamples(name string, observations []callObservation, limit int) []UsageExample {
	examples := make([]UsageExample, 0, limit)
	for _, observation := range observations {
		examples = appendUniqueExample(examples, UsageExample{
			Addon:   observation.Addon,
			Caller:  observation.Caller,
			Snippet: formatSnippet(name, observation.Arguments),
			Source:  observation.Source,
		})
		if len(examples) == limit {
			break
		}
	}
	return examples
}

func firstUsageExamples(examples []UsageExample, limit int) []UsageExample {
	if len(examples) <= limit {
		return append([]UsageExample{}, examples...)
	}
	return append([]UsageExample{}, examples[:limit]...)
}

func firstStrings(values []string, limit int) []string {
	if len(values) <= limit {
		return append([]string{}, values...)
	}
	return append([]string{}, values[:limit]...)
}

func formatSnippet(name string, arguments []string) string {
	if name == "" {
		return "(" + strings.Join(arguments, ", ") + ")"
	}
	return name + "(" + strings.Join(arguments, ", ") + ")"
}

func examplesToEvidence(examples []UsageExample) []string {
	evidence := []string{}
	for _, example := range examples {
		evidence = append(evidence, example.Addon+": "+example.Snippet)
	}
	return evidence
}

func eventSymbolNames(values []EventSymbol, limit int) []string {
	result := []string{}
	for _, value := range values {
		result = append(result, value.Name+" ("+string(value.Confidence)+")")
		if len(result) == limit {
			break
		}
	}
	return result
}

func bindingEvidence(bindings []BindingDoc, limit int) []string {
	result := []string{}
	for _, binding := range bindings {
		result = append(result, binding.Addon+": "+binding.Frame+"."+binding.Event+" -> "+binding.LuaFunction)
		if len(result) == limit {
			break
		}
	}
	return result
}

func lifecycleEvidence(phases []LifecyclePhase, names ...string) []string {
	lookup := map[string]LifecyclePhase{}
	for _, phase := range phases {
		lookup[phase.Phase] = phase
	}
	result := []string{}
	for _, name := range names {
		phase, ok := lookup[name]
		if !ok {
			continue
		}
		result = append(result, name+": "+strings.Join(firstStrings(phase.Evidence, 2), ", "))
	}
	return result
}

func combinePatternEvidence(patterns []PatternDoc, limit int) []string {
	result := []string{}
	for _, pattern := range patterns {
		for _, evidence := range pattern.Evidence {
			result = append(result, pattern.Name+": "+evidence)
			if len(result) == limit {
				return result
			}
		}
	}
	return result
}

func isNumericLike(value string) bool {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return false
	}
	for _, runeValue := range trimmed {
		if (runeValue < '0' || runeValue > '9') && runeValue != '.' && runeValue != '-' {
			return false
		}
	}
	return true
}

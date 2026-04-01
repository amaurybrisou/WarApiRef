package deep_analysis

import (
	"regexp"
	"sort"
	"strings"
)

type downstreamSignal struct {
	Position   int
	Type       string
	Role       string
	Pattern    string
	Confidence int
	Evidence   string
}

type callSiteObservation struct {
	CallerPath string
	Line       int
	CalleeName string
	Assigned   []string
	Raw        string
	Signals    []downstreamSignal
}

// AdvancedReturnAnalyzer reconstructs return semantics using multiple evidence sources.
type AdvancedReturnAnalyzer struct {
	functions     map[string]FunctionSource
	byName        map[string][]string
	callsByCallee map[string][]callSiteObservation
	callsByName   map[string][]callSiteObservation
}

func NewAdvancedReturnAnalyzer() *AdvancedReturnAnalyzer {
	return &AdvancedReturnAnalyzer{
		functions:     map[string]FunctionSource{},
		byName:        map[string][]string{},
		callsByCallee: map[string][]callSiteObservation{},
		callsByName:   map[string][]callSiteObservation{},
	}
}

func (a *AdvancedReturnAnalyzer) BuildIndex(functions []FunctionSource) {
	for _, fn := range functions {
		a.functions[fn.Path] = fn
		nameKey := normalizeName(fn.Name)
		a.byName[nameKey] = append(a.byName[nameKey], fn.Path)
	}

	for _, fn := range functions {
		obs := extractCallAssignments(fn.Path, fn.Source)
		for _, o := range obs {
			nameKey := normalizeName(o.CalleeName)
			a.callsByName[nameKey] = append(a.callsByName[nameKey], o)

			resolved := a.resolveCalleePaths(o.CalleeName)
			for _, calleePath := range resolved {
				a.callsByCallee[calleePath] = append(a.callsByCallee[calleePath], o)
			}
		}
	}
}

// AnalyzeFunctionByName performs return inference from call-site flows even when
// the callee has no local source definition in the corpus.
func (a *AdvancedReturnAnalyzer) AnalyzeFunctionByName(functionName string) AdvancedReturnReport {
	nameKey := normalizeName(functionName)
	observations := a.callsByName[nameKey]
	if len(observations) == 0 {
		return AdvancedReturnReport{FunctionPath: functionName}
	}

	report := AdvancedReturnReport{
		FunctionPath:     functionName,
		BranchSensitive:  false,
		WrapperAffected:  false,
		RuntimeObserved:  false,
		Evidence:         []ProvenanceRecord{},
		UncertaintyNotes: []string{},
	}

	observedArities := map[int]bool{}
	for _, obs := range observations {
		if len(obs.Assigned) > 0 {
			observedArities[len(obs.Assigned)] = true
		}
		report.Evidence = append(report.Evidence, ProvenanceRecord{
			FunctionPath:   functionName,
			SourceFunction: obs.CallerPath,
			Line:           obs.Line,
			Pattern:        "call_site_assignment",
			Snippet:        obs.Raw,
			Observation:    "Return value consumed at call site",
		})
	}
	report.ObservedReturnArity = mapKeysInt(observedArities)
	if len(report.ObservedReturnArity) > 0 {
		report.InferredReturnArity = report.ObservedReturnArity[len(report.ObservedReturnArity)-1]
	}

	report.ReturnVariants = buildReturnVariants(nil, observations)
	report.ReturnPositions = buildPositionInferences(report, nil, observations)
	report.Rationale = buildReturnRationale(report, observations)

	if len(report.ReturnVariants) > 1 {
		report.UncertaintyNotes = append(report.UncertaintyNotes, "Contradictory return behavior observed across call sites.")
	}

	return report
}

func (a *AdvancedReturnAnalyzer) AnalyzeAll() map[string]AdvancedReturnReport {
	reports := map[string]AdvancedReturnReport{}
	for path := range a.functions {
		reports[path] = a.AnalyzeFunction(path)
	}
	return reports
}

func (a *AdvancedReturnAnalyzer) AnalyzeFunction(functionPath string) AdvancedReturnReport {
	fn, ok := a.functions[functionPath]
	if !ok {
		return AdvancedReturnReport{FunctionPath: functionPath}
	}

	returnShapes, wrapperAffected, branchSensitive, shapeEvidence := analyzeReturnStatements(fn)
	callObs := a.callsByCallee[functionPath]

	report := AdvancedReturnReport{
		FunctionPath:     functionPath,
		BranchSensitive:  branchSensitive,
		WrapperAffected:  wrapperAffected,
		RuntimeObserved:  false,
		Evidence:         shapeEvidence,
		UncertaintyNotes: []string{},
	}

	observedArities := map[int]bool{}
	for _, shape := range returnShapes {
		observedArities[shape.Arity] = true
	}
	for _, obs := range callObs {
		if len(obs.Assigned) > 0 {
			observedArities[len(obs.Assigned)] = true
		}
	}
	report.ObservedReturnArity = mapKeysInt(observedArities)
	if len(report.ObservedReturnArity) > 0 {
		report.InferredReturnArity = report.ObservedReturnArity[len(report.ObservedReturnArity)-1]
	}

	report.ReturnVariants = buildReturnVariants(returnShapes, callObs)
	report.ReturnPositions = buildPositionInferences(report, returnShapes, callObs)
	report.Rationale = buildReturnRationale(report, callObs)

	if len(report.ReturnVariants) > 1 {
		report.UncertaintyNotes = append(report.UncertaintyNotes, "Multiple return variants observed; semantics may depend on branch or call context.")
	}
	if report.WrapperAffected {
		report.UncertaintyNotes = append(report.UncertaintyNotes, "Wrapper transformations detected; raw callee semantics may be altered.")
	}
	if report.BranchSensitive {
		report.UncertaintyNotes = append(report.UncertaintyNotes, "Branch-sensitive return behavior detected.")
	}

	return report
}

type returnShape struct {
	Arity      int
	Values     []string
	Types      []string
	Roles      []string
	BranchNote string
	Evidence   []ProvenanceRecord
}

func analyzeReturnStatements(fn FunctionSource) ([]returnShape, bool, bool, []ProvenanceRecord) {
	lines := strings.Split(fn.Source, "\n")
	shapes := []returnShape{}
	allEvidence := []ProvenanceRecord{}
	wrapperAffected := false
	branchSensitive := false

	returnPattern := regexp.MustCompile(`^\s*return\s*(.*)$`)
	assignmentFromCall := regexp.MustCompile(`^\s*(?:local\s+)?([A-Za-z_][A-Za-z0-9_]*(?:\s*,\s*[A-Za-z_][A-Za-z0-9_]*)*)\s*=\s*([A-Za-z_][A-Za-z0-9_\.]*?)\s*\(`)
	ifPattern := regexp.MustCompile(`\bif\b|\belseif\b|\belse\b`)

	for i, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}
		if ifPattern.MatchString(line) {
			branchSensitive = true
		}

		if m := assignmentFromCall.FindStringSubmatch(line); len(m) > 2 {
			_ = m
		}

		m := returnPattern.FindStringSubmatch(line)
		if len(m) < 2 {
			continue
		}

		expr := strings.TrimSpace(m[1])
		if expr == "" {
			shapes = append(shapes, returnShape{Arity: 0, Values: []string{}, Types: []string{}, Roles: []string{}})
			continue
		}

		if strings.Contains(expr, "not not ") || strings.Contains(expr, "~= nil") {
			wrapperAffected = true
		}
		if regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_\.]*\s*\(`).MatchString(expr) {
			wrapperAffected = true
		}

		parts := splitCSV(expr)
		types := make([]string, 0, len(parts))
		roles := make([]string, 0, len(parts))
		for _, part := range parts {
			types = append(types, classifyExpressionType(part))
			roles = append(roles, inferRoleFromExpression(part))
		}

		ev := ProvenanceRecord{
			FunctionPath:   fn.Path,
			SourceFunction: fn.Path,
			Line:           i + 1,
			Pattern:        "return_statement",
			Snippet:        strings.TrimSpace(rawLine),
			Observation:    "Observed direct return statement",
		}
		allEvidence = append(allEvidence, ev)
		shapes = append(shapes, returnShape{
			Arity:      len(parts),
			Values:     parts,
			Types:      types,
			Roles:      roles,
			BranchNote: "direct_return",
			Evidence:   []ProvenanceRecord{ev},
		})
	}

	if len(shapes) > 1 {
		seen := map[string]bool{}
		for _, s := range shapes {
			key := strings.Join(s.Types, "|")
			seen[key] = true
		}
		if len(seen) > 1 {
			branchSensitive = true
		}
	}

	return shapes, wrapperAffected, branchSensitive, allEvidence
}

func extractCallAssignments(callerPath, source string) []callSiteObservation {
	lines := strings.Split(source, "\n")
	obs := []callSiteObservation{}

	assignPattern := regexp.MustCompile(`^\s*(?:local\s+)?([A-Za-z_][A-Za-z0-9_]*(?:\s*,\s*[A-Za-z_][A-Za-z0-9_]*)*)\s*=\s*([A-Za-z_][A-Za-z0-9_\.]*)\s*\(`)
	for i, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		m := assignPattern.FindStringSubmatch(line)
		if len(m) < 3 {
			continue
		}
		lhs := splitCSV(m[1])
		callee := strings.TrimSpace(m[2])
		signals := []downstreamSignal{}
		for pos, variable := range lhs {
			signals = append(signals, inferDownstreamSignals(variable, pos+1, lines, i+1)...)
		}
		obs = append(obs, callSiteObservation{
			CallerPath: callerPath,
			Line:       i + 1,
			CalleeName: callee,
			Assigned:   lhs,
			Raw:        line,
			Signals:    signals,
		})
	}

	return obs
}

func inferDownstreamSignals(variable string, position int, lines []string, startLine int) []downstreamSignal {
	signals := []downstreamSignal{}
	name := regexp.QuoteMeta(variable)
	window := 40
	end := startLine + window
	if end > len(lines) {
		end = len(lines)
	}

	patterns := []struct {
		re         *regexp.Regexp
		typeName   string
		role       string
		confidence int
		pattern    string
	}{
		{regexp.MustCompile(`\bif\s+` + name + `\b|\b` + name + `\s*~=\s*nil|\b` + name + `\s*==\s*(true|false)`), "boolean", "success_flag", 78, "comparison"},
		{regexp.MustCompile(name + `\.[A-Za-z_][A-Za-z0-9_]*`), "table-like", "object_table", 68, "field_access"},
		{regexp.MustCompile(name + `\[[^\]]+\]`), "table-like", "list_table", 66, "index_access"},
		{regexp.MustCompile(`pairs\(` + name + `\)|ipairs\(` + name + `\)`), "table-like", "list_table", 72, "iteration"},
		{regexp.MustCompile(name + `\s*\.\.|\.\.\s*` + name), "string-like", "text", 65, "concatenation"},
		{regexp.MustCompile(name + `\s*[\+\-\*/]`), "number-like", "count", 70, "arithmetic"},
	}

	passPattern := regexp.MustCompile(`([A-Za-z_][A-Za-z0-9_\.]*)\s*\(` + name + `(,|\))`)

	for i := startLine; i < end; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		for _, p := range patterns {
			if p.re.MatchString(line) {
				signals = append(signals, downstreamSignal{
					Position:   position,
					Type:       p.typeName,
					Role:       p.role,
					Pattern:    p.pattern,
					Confidence: p.confidence,
					Evidence:   line,
				})
			}
		}

		if m := passPattern.FindStringSubmatch(line); len(m) > 1 {
			role := inferRoleFromCalleeName(m[1])
			signals = append(signals, downstreamSignal{
				Position:   position,
				Type:       inferTypeFromRole(role),
				Role:       role,
				Pattern:    "forwarded_call",
				Confidence: 74,
				Evidence:   line,
			})
		}
	}

	return signals
}

func (a *AdvancedReturnAnalyzer) resolveCalleePaths(callee string) []string {
	trimmed := strings.TrimSpace(callee)
	if trimmed == "" {
		return nil
	}
	parts := strings.Split(trimmed, ".")
	name := normalizeName(parts[len(parts)-1])
	return append([]string{}, a.byName[name]...)
}

func normalizeName(name string) string {
	name = strings.TrimSpace(name)
	if idx := strings.LastIndex(name, "."); idx >= 0 {
		name = name[idx+1:]
	}
	return strings.ToLower(name)
}

func splitCSV(expr string) []string {
	raw := strings.Split(expr, ",")
	parts := make([]string, 0, len(raw))
	for _, p := range raw {
		t := strings.TrimSpace(p)
		if t != "" {
			parts = append(parts, t)
		}
	}
	return parts
}

func mapKeysInt(values map[int]bool) []int {
	out := make([]int, 0, len(values))
	for key := range values {
		out = append(out, key)
	}
	sort.Ints(out)
	return out
}

func classifyExpressionType(expr string) string {
	e := strings.TrimSpace(expr)
	switch {
	case e == "nil":
		return "nil"
	case e == "true" || e == "false" || strings.Contains(e, "not not") || strings.Contains(e, "~= nil"):
		return "boolean"
	case regexp.MustCompile(`^-?\d+(\.\d+)?$`).MatchString(e):
		return "number-like"
	case regexp.MustCompile(`^L?".*"$|^'.*'$`).MatchString(e):
		return "string-like"
	case strings.HasPrefix(e, "{") && strings.HasSuffix(e, "}"):
		return "table-like"
	case strings.Contains(e, ".") || strings.Contains(e, "["):
		return "table-like"
	case regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_\.]*\s*\(`).MatchString(e):
		return "unknown-call-result"
	default:
		return "unknown"
	}
}

func inferRoleFromExpression(expr string) string {
	e := strings.TrimSpace(expr)
	switch {
	case e == "true" || e == "false" || strings.Contains(e, "not not") || strings.Contains(e, "~= nil"):
		return "success_flag"
	case strings.Contains(strings.ToLower(e), "err"):
		return "status_code"
	case regexp.MustCompile(`^L?".*"$|^'.*'$`).MatchString(e):
		return "text"
	case regexp.MustCompile(`^-?\d+(\.\d+)?$`).MatchString(e):
		return "count"
	case strings.Contains(e, "Window"):
		return "window_name"
	default:
		return "unknown"
	}
}

func inferRoleFromCalleeName(name string) string {
	lower := strings.ToLower(name)
	switch {
	case strings.Contains(lower, "window"):
		return "window_name"
	case strings.Contains(lower, "frame"):
		return "frame_name"
	case strings.Contains(lower, "event"):
		return "event_payload"
	case strings.Contains(lower, "name") || strings.Contains(lower, "text") || strings.Contains(lower, "string"):
		return "text"
	case strings.Contains(lower, "id"):
		return "id"
	case strings.Contains(lower, "count") || strings.Contains(lower, "size"):
		return "count"
	default:
		return "unknown"
	}
}

func inferTypeFromRole(role string) string {
	switch role {
	case "success_flag":
		return "boolean"
	case "window_name", "frame_name", "text", "status_code":
		return "string-like"
	case "id", "count", "enum-like value":
		return "number-like"
	case "object_table", "list_table", "event_payload":
		return "table-like"
	default:
		return "unknown"
	}
}

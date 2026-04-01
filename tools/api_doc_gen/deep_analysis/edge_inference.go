package deep_analysis

import (
	"regexp"
	"strings"
)

// EdgeInference contains logic for discovering missing relationships between symbols
type EdgeInference struct {
	AllFunctions map[string]*FunctionAnalysis // map of function ID to analysis
	CallGraph    map[string][]string          // function -> functions it calls
	GlobalAccess map[string][]string          // function -> global tables/data it accesses
	EventChains  map[string][]string          // event name -> functions that handle it
}

// FunctionAnalysis holds extracted information about a function
type FunctionAnalysis struct {
	Name              string
	FullPath          string
	Source            string // raw function source code
	IsInitFunction    bool
	IsUpdateFunction  bool
	IsEventHandler    bool
	EventName         string // if IsEventHandler, what event
	AccessedGlobals   []string
	CallsSites        []CallSite
	ReturnStmts       []string
	AssignmentTargets []string
	UIOperations      []UIOperation
}

// CallSite records a function call
type CallSite struct {
	CalledFunction string
	Line           int
	Args           []string
}

// UIOperation records UI manipulation
type UIOperation struct {
	Operation string // "SetShowing", "Show", "Hide", "SetAnchor", etc
	Target    string // what frame/window
	Line      int
	Args      []string
}

// NewEdgeInference creates inference engine
func NewEdgeInference() *EdgeInference {
	return &EdgeInference{
		AllFunctions: make(map[string]*FunctionAnalysis),
		CallGraph:    make(map[string][]string),
		GlobalAccess: make(map[string][]string),
		EventChains:  make(map[string][]string),
	}
}

// AnalyzeFunctionSource parses function source and extracts calls, data access, UI operations
func (ei *EdgeInference) AnalyzeFunctionSource(functionPath, source string, isInit, isUpdate, isHandler bool, eventName string) *FunctionAnalysis {
	fa := &FunctionAnalysis{
		FullPath:          functionPath,
		Source:            source,
		IsInitFunction:    isInit,
		IsUpdateFunction:  isUpdate,
		IsEventHandler:    isHandler,
		EventName:         eventName,
		AccessedGlobals:   []string{},
		CallsSites:        []CallSite{},
		ReturnStmts:       []string{},
		AssignmentTargets: []string{},
		UIOperations:      []UIOperation{},
	}

	parts := strings.Split(functionPath, ".")
	if len(parts) > 0 {
		fa.Name = parts[len(parts)-1]
	}

	// Extract function calls: func() or module.func()
	callRegex := regexp.MustCompile(`(\w+(?:\.\w+)*)\s*\(`)
	callMatches := callRegex.FindAllStringSubmatchIndex(source, -1)
	for _, match := range callMatches {
		callName := source[match[2]:match[3]]
		fa.CallsSites = append(fa.CallsSites, CallSite{CalledFunction: callName})
	}

	// Extract global access patterns
	globalAccessPatterns := []string{
		`SystemData\.(\w+)`,
		`GameData\.(\w+)`,
		`d\.module\.(\w+)`,
		`\b(global)\b`,
	}
	for _, pattern := range globalAccessPatterns {
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllString(source, -1)
		fa.AccessedGlobals = append(fa.AccessedGlobals, matches...)
	}

	// Extract return statements
	returnRegex := regexp.MustCompile(`return\s+([^;\n]*)`)
	returnMatches := returnRegex.FindAllString(source, -1)
	fa.ReturnStmts = append(fa.ReturnStmts, returnMatches...)

	// Extract UI operations: SetShowing, Show, Hide, SetAnchor, SetSize, SetPosition
	uiOpsPatterns := []string{
		`SetShowing\s*\(\s*(\w+)`,
		`(\w+):Show\(\)`,
		`(\w+):Hide\(\)`,
		`(\w+):SetAnchor`,
		`(\w+):SetSize`,
		`(\w+):SetPosition`,
	}
	for _, pattern := range uiOpsPatterns {
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllStringSubmatch(source, -1)
		for _, m := range matches {
			if len(m) > 1 {
				op := performanceRuleMatch(pattern)
				fa.UIOperations = append(fa.UIOperations, UIOperation{
					Operation: op,
					Target:    m[1],
				})
			}
		}
	}

	// Extract assignment targets
	assignmentRegex := regexp.MustCompile(`(\w+)\s*=`)
	assignMatches := assignmentRegex.FindAllString(source, -1)
	fa.AssignmentTargets = append(fa.AssignmentTargets, assignMatches...)

	ei.AllFunctions[functionPath] = fa
	return fa
}

// performanceRuleMatch extracts operation name from pattern
func performanceRuleMatch(pattern string) string {
	if strings.Contains(pattern, "SetShowing") {
		return "SetShowing"
	}
	if strings.Contains(pattern, ":Show") {
		return "Show"
	}
	if strings.Contains(pattern, ":Hide") {
		return "Hide"
	}
	if strings.Contains(pattern, "SetAnchor") {
		return "SetAnchor"
	}
	if strings.Contains(pattern, "SetSize") {
		return "SetSize"
	}
	if strings.Contains(pattern, "SetPosition") {
		return "SetPosition"
	}
	return "unknown"
}

// InferReadsSystemData finds functions accessing SystemData
func (ei *EdgeInference) InferReadsSystemData(functionPath string) bool {
	fa, ok := ei.AllFunctions[functionPath]
	if !ok {
		return false
	}
	for _, global := range fa.AccessedGlobals {
		if strings.Contains(global, "SystemData") {
			return true
		}
	}
	return false
}

// InferReadsGameData finds functions accessing GameData
func (ei *EdgeInference) InferReadsGameData(functionPath string) bool {
	fa, ok := ei.AllFunctions[functionPath]
	if !ok {
		return false
	}
	for _, global := range fa.AccessedGlobals {
		if strings.Contains(global, "GameData") {
			return true
		}
	}
	return false
}

// InferUpdatesUI detects UI update patterns
func (ei *EdgeInference) InferUpdatesUI(functionPath string) bool {
	fa, ok := ei.AllFunctions[functionPath]
	if !ok {
		return false
	}
	for _, call := range fa.CallsSites {
		if strings.Contains(call.CalledFunction, "Refresh") ||
			strings.Contains(call.CalledFunction, "Update") ||
			strings.Contains(call.CalledFunction, "Invalidate") {
			return true
		}
	}
	// Also check for direct UI operations in this function
	return len(fa.UIOperations) > 0
}

// InferTogglesVisibility detects Show/Hide patterns
func (ei *EdgeInference) InferTogglesVisibility(functionPath string) bool {
	fa, ok := ei.AllFunctions[functionPath]
	if !ok {
		return false
	}
	for _, op := range fa.UIOperations {
		if op.Operation == "Show" || op.Operation == "Hide" || op.Operation == "SetShowing" {
			return true
		}
	}
	return false
}

// InferUpdatesLayout detects layout modification patterns
func (ei *EdgeInference) InferUpdatesLayout(functionPath string) bool {
	fa, ok := ei.AllFunctions[functionPath]
	if !ok {
		return false
	}
	for _, op := range fa.UIOperations {
		if op.Operation == "SetAnchor" || op.Operation == "SetSize" || op.Operation == "SetPosition" {
			return true
		}
	}
	return false
}

// InferInitializes detects init-time calls
func (ei *EdgeInference) InferInitializes(functionPath string) bool {
	fa, ok := ei.AllFunctions[functionPath]
	if !ok {
		return false
	}
	return fa.IsInitFunction
}

// InferRefreshes detects periodic update patterns
func (ei *EdgeInference) InferRefreshes(functionPath string) bool {
	fa, ok := ei.AllFunctions[functionPath]
	if !ok {
		return false
	}
	return fa.IsUpdateFunction
}

// InferEventTriggered connects event handlers to source events
func (ei *EdgeInference) InferEventTriggered(functionPath string) (string, bool) {
	fa, ok := ei.AllFunctions[functionPath]
	if !ok {
		return "", false
	}
	if fa.IsEventHandler {
		return fa.EventName, true
	}
	return "", false
}

// InferCommonlyUsedWith finds co-occurrence patterns
func (ei *EdgeInference) InferCommonlyUsedWith(from, to string) bool {
	f1, ok1 := ei.AllFunctions[from]
	if !ok1 {
		return false
	}

	// Check if target is called from source
	for _, call := range f1.CallsSites {
		if strings.HasSuffix(to, call.CalledFunction) {
			return true
		}
	}
	return false
}

// InferAppearsInSameFlow finds functions in same event/call chain
func (ei *EdgeInference) InferAppearsInSameFlow(from, to string) bool {
	f1, ok1 := ei.AllFunctions[from]
	if !ok1 {
		return false
	}
	f2, ok2 := ei.AllFunctions[to]
	if !ok2 {
		return false
	}

	// Same event handler or same init/update phase
	if f1.IsEventHandler && f2.IsEventHandler && f1.EventName == f2.EventName {
		return true
	}
	if f1.IsInitFunction && f2.IsInitFunction {
		return true
	}
	if f1.IsUpdateFunction && f2.IsUpdateFunction {
		return true
	}

	return false
}

// CountRelatedEdges returns potential new edges and their confidence
func (ei *EdgeInference) CountRelatedEdges() map[string]int {
	counts := make(map[string]int)
	for path, fa := range ei.AllFunctions {
		_ = path // use all functions

		if ei.InferReadsSystemData(path) {
			counts["reads_systemdata"]++
		}
		if ei.InferReadsGameData(path) {
			counts["reads_gamedata"]++
		}
		if ei.InferUpdatesUI(path) {
			counts["updates_ui"]++
		}
		if ei.InferTogglesVisibility(path) {
			counts["toggles_visibility"]++
		}
		if ei.InferUpdatesLayout(path) {
			counts["updates_layout"]++
		}
		if fa.IsInitFunction {
			counts["initializes"]++
		}
		if fa.IsUpdateFunction {
			counts["refreshes"]++
		}
	}
	return counts
}

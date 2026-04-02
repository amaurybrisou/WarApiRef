// Package lua_semantic implements Phase 3 of the XML↔Lua analysis pipeline.
//
// Phase 3 — LUA FUNCTION ANALYSIS: Deeply analyze correlated Lua functions
// to extract parameters, inferred argument names/types, callback shapes,
// downstream calls, state reads/writes, table variable internals, and
// return shapes. All inferences are evidence-based and conservative.
//
// This phase consumes the binding set from Phase 2 and the source model
// to produce [LuaSemanticCorpus].
package lua_semantic

// LuaSemanticCorpus is the output of Phase 3, containing deep analysis
// results for Lua functions that are correlated with XML elements.
type LuaSemanticCorpus struct {
	// Functions maps qualified function name → analysis result.
	Functions map[string]*FunctionAnalysis

	// HandlerAnalyses maps "addon|frame|event" → handler-specific analysis.
	HandlerAnalyses map[string]*HandlerAnalysis
}

// FunctionAnalysis holds the deep semantic analysis of a single Lua function.
type FunctionAnalysis struct {
	// Identity
	Name          string   // Function name as declared
	QualifiedName string   // Fully qualified name (e.g. "MyAddon.OnInitialize")
	Addon         string   // Source addon
	File          string   // Source file
	Line          int      // Declaration line
	EndLine       int      // End line

	// Parameters
	Parameters []ParameterAnalysis // Ordered parameter analysis

	// Return analysis
	Returns         ReturnAnalysis
	ObservedArity   []int  // Observed return value counts
	InferredArity   int    // Best-guess return arity

	// Downstream calls
	Calls []CallAnalysis // Functions called within this function

	// State access
	StateReads  []StateAccess // Global/module state reads
	StateWrites []StateAccess // Global/module state writes

	// Callback shape
	CallbackShape *CallbackShape // Non-nil if function appears to be a callback

	// Evidence
	AnalysisConfidence string // Overall confidence: HIGH, MEDIUM, LOW
	EvidenceCount      int    // Number of independent observations
	Notes              []string
}

// ParameterAnalysis describes what we know about a function parameter.
type ParameterAnalysis struct {
	Name       string // Parameter name from source
	Position   int    // 0-based position
	Role       string // Inferred role: "self", "callback", "flag", "identifier", "data", "unknown"
	Type       string // Inferred type: "number", "string", "boolean", "table", "function", "wstring", "unknown"
	IsOptional bool   // Evidence suggests parameter may be omitted

	// Evidence supporting the inference
	UsagePatterns []string // How the parameter is used in the function body
	CallSiteHints []string // What callers pass for this parameter
	Confidence    string   // HIGH, MEDIUM, LOW
	Notes         []string
}

// ReturnAnalysis describes what a function returns.
type ReturnAnalysis struct {
	HasReturn bool   // Function contains return statements
	Variants  []ReturnVariant // Different return shapes observed

	// Best-guess summary
	PrimaryType string // Most likely return type
	PrimaryRole string // What the return value represents
	Confidence  string // HIGH, MEDIUM, LOW
	Notes       []string
}

// ReturnVariant describes one observed return shape.
type ReturnVariant struct {
	Arity      int      // Number of values returned
	Types      []string // Inferred type per position
	Roles      []string // Inferred role per position
	Frequency  int      // How often this variant is observed
	Confidence string   // HIGH, MEDIUM, LOW
}

// CallAnalysis records a function call made from within the analyzed function.
type CallAnalysis struct {
	Callee       string   // Called function name
	QualifiedName string  // Fully qualified callee name
	Arguments    []string // Argument expressions (raw)
	Line         int      // Source line
	IsAPICall    bool     // Appears to be a platform/engine API call
	IsWindowCall bool     // Appears to be a Window/UI API call
	Category     string   // Classification: "ui", "data", "event", "utility", "unknown"
}

// StateAccess records a read or write to module/global state.
type StateAccess struct {
	Variable      string // Variable name (e.g. "MyAddon.settings")
	Owner         string // Module/table owner (e.g. "MyAddon")
	Field         string // Specific field if applicable
	AccessType    string // "read" or "write"
	Line          int    // Source line
	ValueHint     string // Inferred type/value of what's stored (for writes)
}

// CallbackShape describes the expected signature when a function is used as a callback.
type CallbackShape struct {
	ExpectedParams []CallbackParam // Parameters the callback is expected to receive
	CallerContext  string          // What calls this callback (e.g. "XML EventHandler OnClick")
	Pattern        string          // Classification: "event_handler", "data_callback", "update_loop", "unknown"
	Confidence     string          // HIGH, MEDIUM, LOW
}

// CallbackParam describes an expected parameter for a callback function.
type CallbackParam struct {
	Position int    // 0-based position
	Name     string // Expected parameter name (from convention or documentation)
	Type     string // Expected type
	Role     string // What this parameter carries
	Source   string // Where this inference comes from: "xml_handler", "api_docs", "usage_pattern"
}

// HandlerAnalysis specialises FunctionAnalysis for XML event handler functions,
// carrying additional context about the XML element and event that bind it.
type HandlerAnalysis struct {
	*FunctionAnalysis

	// XML context
	FrameName string // XML frame name
	FrameType string // XML element type tag
	Event     string // XML event name (e.g. "OnInitialize")
	Addon     string // Addon

	// Handler-specific inference
	InferredHandlerArgs []CallbackParam // Expected arguments based on the event type
	HandlerCategory     string          // "lifecycle", "input", "data", "custom"
	HandlerConfidence   string          // HIGH, MEDIUM, LOW
}

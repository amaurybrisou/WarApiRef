// Package semantic_merge implements Phase 4 of the XML↔Lua analysis pipeline.
//
// Phase 4 — UPDATE XML SEMANTICS WITH LUA FINDINGS: Merge Lua analysis
// results back into the XML semantic model so that XML element documentation
// can expose parents, children, structural sub-elements, XML handlers/events,
// Lua-bound functions, tag properties, callback/argument semantics, and
// common XML↔Lua composition patterns.
//
// The output is an [EnrichedElementCatalog] that serves as the single source
// of truth for downstream documentation, graph, and MCP artifact generation.
package semantic_merge

// EnrichedElementCatalog is the output of Phase 4 — a fully merged model
// where each element type carries both XML structural data and Lua semantic
// findings, ready for Phase 5 (generation).
type EnrichedElementCatalog struct {
	Elements map[string]*EnrichedElement // Tag name → enriched element
}

// EnrichedElement is the merged model for a single XML element type,
// combining structural XML data with Lua analysis results.
type EnrichedElement struct {
	// Identity
	Tag         string // Element type tag (e.g. "Button", "ListBox")
	Description string // Generated description

	// Confidence
	Confidence string // Overall confidence: HIGH, MEDIUM, LOW
	Score      int    // Numeric score 0-100

	// Occurrence
	SeenIn     []string // Addon names where this type appears
	TotalCount int      // Total instances across all addons

	// === XML STRUCTURE (from Phase 1) ===

	// Relationships
	Parents            []ElementRef // Element types that contain this one
	Children           []ElementRef // Named child element types
	StructuralChildren []StructuralChildRef // Unnamed structural sub-elements

	// Attributes
	AttributeProfiles []AttributeRef // Attribute reference with required/optional

	// Inheritance
	InheritsBases []string // Template bases commonly used
	IsTemplate    bool     // Whether this type is itself a template

	// Composition
	CompositionSnippet string // Representative XML hierarchy snippet

	// === XML↔LUA BINDINGS (from Phase 2 + Phase 3) ===

	// Event bindings with Lua semantic data
	EventBindings []EnrichedEventBinding

	// Lua functions that manipulate frames of this type
	LuaManipulators []LuaManipulatorRef

	// === LUA SEMANTIC DATA (from Phase 3) ===

	// Aggregated Lua API calls made from handlers on this element type
	LuaAPICalls []LuaAPICallRef

	// Common callback/argument patterns for handlers on this type
	HandlerArgPatterns []HandlerArgPattern

	// === NOTES ===
	Notes []string
}

// ElementRef is a reference to another element type with frequency data.
type ElementRef struct {
	Tag        string // Element type tag
	Count      int    // How many times this relationship was observed
	Confidence string // HIGH, MEDIUM, LOW
}

// StructuralChildRef describes an unnamed structural child element type.
type StructuralChildRef struct {
	Tag            string         // Child element tag (e.g. "ListData")
	Count          int            // Observation count
	Attributes     []AttributeRef // Attributes observed on instances of this child
	Confidence     string
}

// AttributeRef describes an observed attribute with usage statistics.
type AttributeRef struct {
	Name         string   // Attribute key
	IsRequired   bool     // Present in >80% of instances
	SampleValues []string // Up to N distinct observed values (not sensitive data)
	Count        int      // How many elements have this attribute
	TotalCount   int      // Total elements of this type (for computing required %)
}

// EnrichedEventBinding merges XML handler declaration data with Lua semantic
// analysis of the handler function.
type EnrichedEventBinding struct {
	Event string // XML event name (e.g. "OnInitialize", "OnLButtonUp")

	// Lua functions bound to this event (with frequency)
	LuaFunctions []BoundFunction

	// Inferred handler argument signature
	InferredArgs   string // Human-readable signature (e.g. "self, mouseButton")
	ArgsConfidence string // HIGH, MEDIUM, LOW

	// Per-event Lua API calls (aggregated from handler function analysis)
	LuaAPICalls []string // Downstream API calls made by handler functions

	// Handler category
	Category string // "lifecycle", "input", "data", "custom"

	// Observation count
	TotalCount int
}

// BoundFunction describes a Lua function bound to an XML event with resolution status.
type BoundFunction struct {
	Name       string // Lua function name
	Count      int    // How many times this binding was observed
	Resolved   bool   // Whether the function was found in Lua source
	HasAnalysis bool  // Whether Phase 3 produced deep analysis for this function
}

// LuaManipulatorRef describes a Lua function that manipulates frames of a specific type.
type LuaManipulatorRef struct {
	Function    string // Lua function name
	CallPattern string // API pattern used (e.g. "CreateWindowFromTemplate")
	Count       int    // Observation count
	Confidence  string
}

// LuaAPICallRef describes an API call commonly made from handlers on an element type.
type LuaAPICallRef struct {
	Function   string // Called API function name
	Count      int    // How many handlers call this
	Category   string // Classification of the call
	FromEvents []string // Which events' handlers make this call
}

// HandlerArgPattern describes a common callback argument pattern for handlers.
type HandlerArgPattern struct {
	Event          string          // XML event name
	ExpectedParams []ExpectedParam // Expected parameters for handlers of this event
	Confidence     string          // HIGH, MEDIUM, LOW
	Source         string          // Where this pattern comes from
}

// ExpectedParam describes an expected parameter for a handler callback.
type ExpectedParam struct {
	Position int    // 0-based
	Name     string // Expected name
	Type     string // Expected type
	Role     string // What it represents
}

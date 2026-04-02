// Package xml_lua_binding implements Phase 2 of the XML↔Lua analysis pipeline.
//
// Phase 2 — LUA CORRELATION: Using the XML structural model from Phase 1,
// correlate Lua code with XML elements. This phase builds real XML↔Lua
// bindings by matching:
//   - XML frame/tag names referenced in Lua code
//   - XML EventHandler function="..." declarations to Lua function definitions
//   - Lua functions that manipulate identified XML frames/windows
//   - Event/handler ownership by XML element type
//
// The output is an [XMLLuaBindingSet] that maps XML elements to their
// correlated Lua functions and vice versa.
package xml_lua_binding

// XMLLuaBindingSet is the output of Phase 2, containing all discovered
// correlations between XML elements and Lua code.
type XMLLuaBindingSet struct {
	// HandlerBindings maps XML handler declarations to their Lua functions.
	// Key: "addon|frame_name|event" → binding details
	HandlerBindings []*HandlerBinding

	// FrameManipulations records Lua functions that reference specific XML frames.
	FrameManipulations []*FrameManipulation

	// ElementTypeBindings aggregates bindings per element type tag.
	ElementTypeBindings map[string]*ElementTypeBinding

	// UnresolvedHandlers lists handler declarations where the Lua function was not found.
	UnresolvedHandlers []*UnresolvedHandler

	// Statistics provides summary statistics about the binding process.
	Statistics BindingStatistics
}

// BindingStatistics tracks resolution success rates and binding quality.
type BindingStatistics struct {
	TotalHandlers     int // Total XML handler declarations processed
	ResolvedHandlers  int // Successfully resolved to Lua functions
	UnresolvedCount   int // Handler declarations with no Lua match
	HighConfidence    int // HIGH confidence resolutions
	MediumConfidence  int // MEDIUM confidence resolutions
	LowConfidence     int // LOW confidence resolutions (including unresolved)
	UniqueEvents      int // Distinct event names seen
	UniqueElementTypes int // Distinct element types with handlers
	ManipulationCount int // Frame manipulation patterns found
}

// HandlerBinding is a resolved link between an XML EventHandler declaration
// and the Lua function it references.
type HandlerBinding struct {
	// XML side
	Addon     string // Addon name
	FrameName string // Named frame containing the handler
	FrameType string // Element type tag of the frame
	Event     string // Handler event (e.g. "OnInitialize")
	XMLFile   string // XML source file
	XMLLine   int    // Line of the EventHandler declaration

	// Lua side
	LuaFunction     string // Resolved Lua function name
	LuaFile         string // Lua source file (empty if not resolved to a file)
	LuaLine         int    // Line in Lua file (0 if not resolved)
	LuaParams       []string // Parameter names from function definition
	LuaIsLocal      bool   // Whether the Lua function is local
	LuaQualifiedName string // Fully qualified Lua name (e.g. "Module.Function")

	// Resolution
	Resolved   bool   // Whether the Lua function was found in the Lua source
	Confidence string // HIGH, MEDIUM, LOW
}

// FrameManipulation records a Lua function that references an XML frame/window
// by name (e.g. via CreateWindowFromTemplate, GetWindowByName, or direct
// variable reference to a frame name).
type FrameManipulation struct {
	LuaFunction string // Lua function performing the manipulation
	LuaFile     string // Source file
	LuaLine     int    // Source line
	FrameName   string // Referenced frame name
	FrameType   string // Element type of the referenced frame (if known)
	Addon       string // Addon context
	CallPattern string // What API call references the frame (e.g. "CreateWindowFromTemplate")
	Confidence  string // HIGH, MEDIUM, LOW
}

// ElementTypeBinding aggregates all bindings for a specific element type tag.
type ElementTypeBinding struct {
	Tag string // Element type tag (e.g. "Button", "ListBox")

	// Handler bindings per event
	EventBindings map[string]*EventBinding // Event name → aggregated binding

	// Functions that manipulate frames of this type
	ManipulatingFunctions []string // Lua function names
	ManipulationPatterns  map[string]int // API call pattern → count
}

// EventBinding aggregates handler observations for one event on one element type.
type EventBinding struct {
	Event        string            // Event name (e.g. "OnLButtonUp")
	LuaFunctions map[string]int    // Lua function → count across addons
	ByAddon      map[string]int    // Addon → count
	TotalCount   int               // Total observations
	Resolved     int               // How many were resolved to actual Lua definitions
	Unresolved   int               // How many were unresolved
}

// UnresolvedHandler is an XML handler declaration where the referenced Lua
// function could not be located in any parsed Lua source.
type UnresolvedHandler struct {
	Addon       string
	FrameName   string
	FrameType   string
	Event       string
	LuaFunction string
	XMLFile     string
	XMLLine     int
}

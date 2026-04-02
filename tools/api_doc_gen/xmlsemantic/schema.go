// Package xmlsemantic provides a structured semantic model for XML element
// types used in WAR addon UI definitions. It bridges the gap between raw parsed
// XML frame/handler data and the richer type-level documentation that describes:
//
//   - Attribute profiles (which attributes are required vs optional, typical values)
//   - Structural child composition (named and unnamed children per element type)
//   - XML event handler bindings (which Lua functions handle which events)
//   - Lua API calls derived from handler functions (what each handler actually does)
//
// All data is derived from corpus evidence — observed frames, handlers, and Lua
// function call graphs — not from static dictionaries or graph proximity.
package xmlsemantic

// FrameInput carries the frame data needed for XML semantic analysis.
// The caller converts platform.FrameDoc records into this form.
type FrameInput struct {
	Addon                string
	Name                 string
	Type                 string
	ParentType           string
	Attributes           map[string]string   // attribute name → observed value
	StructuralChildTypes []string            // unnamed structural child element types
	NamedChildTypes      []string            // element type tags of named child frames
}

// HandlerInput carries one XML event handler binding record.
// The caller converts platform.HandlerDoc (or FrameHandlerDoc) records into this form.
type HandlerInput struct {
	Addon    string
	Frame    string // frame name (used to resolve element type via FrameInput lookup)
	Event    string
	Function string // bound Lua function name
}

// FunctionCallInput carries the minimal Lua function data needed to derive
// Lua API call chains for handler analysis.
// The caller converts platform.FunctionDoc records into this form.
type FunctionCallInput struct {
	Name  string
	Addon string
	Calls []string // names of Lua API functions called by this function
}

// AttributeProfile captures usage statistics for one attribute observed on
// an XML element type, including a required/optional heuristic.
type AttributeProfile struct {
	Name       string   // attribute name (e.g. "rowdef")
	Count      int      // number of observed element instances that carry this attribute
	Total      int      // total observed instances of the element type
	TopValues  []string // up to 5 most commonly observed values
	IsRequired bool     // true when Count/Total >= 0.75
}

// LuaAPICall describes one Lua API function call observed from handler functions
// bound to an XML element type. It is derived from actual call graph analysis.
type LuaAPICall struct {
	FunctionName string // e.g. "WindowGetId" or "EA_Window_ContextMenu.CreateContextMenu"
	ViaEvent     string // the XML handler event through which it is invoked (e.g. "OnLButtonUp")
	Count        int    // total observed invocations across all addons
}

// HandlerEventBinding describes the observed binding of one XML event to the
// Lua functions that handle it on a specific element type, plus the Lua API
// calls those handler functions make.
type HandlerEventBinding struct {
	Event          string       // XML event name (e.g. "OnLButtonUp")
	DirectFuncs    []string     // Lua function names directly bound (by frequency, up to 6)
	LuaAPICalls    []LuaAPICall // Lua API calls made by the handler functions (up to 8)
	InferredArgs   string       // best-effort Lua callback signature
	ArgsConfidence string       // "HIGH", "MEDIUM", or "LOW"
}

// ElementTypeSchema is the structured semantic model for one XML element type.
// It is built from corpus evidence rather than flat frequency counts alone.
type ElementTypeSchema struct {
	Name               string
	TotalInstances     int                   // total observed frames of this type
	AttributeProfiles  []AttributeProfile    // per-attribute profile, sorted by frequency
	StructuralChildren map[string]int        // unnamed structural child types → observed count
	NamedChildTypes    map[string]int        // named child element types → observed count
	ParentTypes        map[string]int        // container element types → observed count
	HandlerBindings    []HandlerEventBinding // per-event handler bindings, sorted by event name
	Addons             []string              // addons where this element type is observed
}

// Package xml_structure implements Phase 1 of the XML↔Lua analysis pipeline.
//
// Phase 1 — XML STRUCTURE: Parse XML as a real tree and produce a strong
// structural model capturing parent/child relationships, named elements,
// derived names from $parent, unnamed structural children, attributes,
// event handler declarations, and template/inheritance information.
//
// The output of this phase is an [XMLTree] per file and an aggregated
// [XMLCorpus] across all addons. Downstream phases consume these models
// rather than the flat Frame/Handler lists from the graph package.
package xml_structure

// XMLTree represents the full parsed structure of a single XML file.
type XMLTree struct {
	Addon    string        // Addon that owns this file
	File     string        // Normalised file path
	Root     []*XMLElement // Top-level elements (usually <Interface> or direct frames)
	AllNodes []*XMLElement // All elements in document order (flat index)
}

// XMLElement is a single node in the XML tree.
type XMLElement struct {
	// Identity
	Tag      string            // Element tag name (e.g. "Window", "Button", "ListData")
	Name     string            // Resolved name attribute (with $parent expanded), empty for unnamed
	RawName  string            // Original name attribute value before $parent expansion
	Addon    string            // Owning addon
	File     string            // Source file
	Line     int               // Source line

	// Structure
	Parent   *XMLElement       // Direct parent element (nil for roots)
	Children []*XMLElement     // All direct child elements

	// Attributes
	Attributes map[string]string // All XML attributes including name, inherits, etc.

	// Classification
	IsNamed       bool   // Has a non-empty name attribute
	IsStructural  bool   // Unnamed child of a named frame (e.g. ListData inside ListBox)
	IsTemplate    bool   // Name or inherits contains "Template"
	IsIgnored     bool   // Layout/anchoring element not tracked as semantic

	// Inheritance / Templates
	Inherits string // Resolved inherits attribute

	// Parent reference
	ParentFrameName string // Name of nearest named ancestor frame
	ParentFrameType string // Tag of nearest named ancestor frame

	// Event handlers declared on this element
	Handlers []XMLHandlerDecl
}

// XMLHandlerDecl represents an <EventHandler> declaration found inside an element's
// <EventHandlers> block.
type XMLHandlerDecl struct {
	Event    string // Handler event name (e.g. "OnInitialize", "OnClick")
	Function string // Lua function name to call
	Line     int    // Source line of the EventHandler element
}

// -----------------------------------------------------------------------
// Aggregated corpus model
// -----------------------------------------------------------------------

// XMLCorpus is the aggregated structural model across all parsed XML files.
// It is the output of Phase 1 and the primary input to Phase 2.
type XMLCorpus struct {
	Trees []*XMLTree // One per parsed XML file

	// Indexes built after parsing for efficient lookup
	ByName       map[string][]*XMLElement  // name → elements with that name
	ByTag        map[string][]*XMLElement  // tag → all elements with that tag
	ByAddon      map[string][]*XMLTree     // addon name → trees
	ByFile       map[string]*XMLTree       // file path → tree

	// Aggregated element type profiles
	ElementTypes map[string]*ElementTypeProfile // tag name → profile
}

// ElementTypeProfile aggregates observations about a specific XML element type
// (e.g. "Button", "ListBox") across all addons and files.
type ElementTypeProfile struct {
	Tag string // Element type tag

	// Occurrence counts
	TotalCount    int            // Total number of elements with this tag
	NamedCount    int            // How many have a name attribute
	UnnamedCount  int            // How many are structural (unnamed)
	ByAddon       map[string]int // Count per addon

	// Attributes observed across all instances
	Attributes map[string]*AttributeProfile // attr name → profile

	// Structural relationships
	ParentTags    map[string]int // Tags of direct parent elements → count
	ChildTags     map[string]int // Tags of direct child elements → count
	NamedChildren map[string]int // Tags of named child elements → count
	StructuralChildren map[string]int // Tags of unnamed structural children → count

	// Inheritance
	InheritsBases map[string]int // Template bases → count
	IsTemplate    bool           // At least one instance is a template

	// Event handlers
	HandlerEvents map[string]int          // Event name → occurrence count
	HandlerBindings map[string]map[string]int // Event → {lua function → count}

	// Composition snippets (best examples of XML structure)
	CompositionSnippets []string // Representative XML hierarchy snippets
}

// AttributeProfile tracks observations about a specific attribute across all
// instances of an element type.
type AttributeProfile struct {
	Name         string         // Attribute key name
	TotalCount   int            // How many elements have this attribute
	SampleValues []string       // Up to N distinct observed values
	ByAddon      map[string]int // Count per addon
	IsRequired   bool           // Appears in >80% of instances (computed after aggregation)
}

package platform

import (
	"time"

	"roraddons/tools/api_doc_gen/confidence"
)

type FunctionDoc struct {
	Name               string
	Addon              string
	Kind               string
	Module             string
	Local              bool
	Source             string
	Parameters         []string
	Aliases            []string
	Calls              []CallDoc
	EventRegistrations []EventRegistrationDoc
	StateWrites        []string
}

type CallDoc struct {
	Name      string
	Line      int
	Arguments []string
	Raw       string
}

type EventRegistrationDoc struct {
	Event   string
	Scope   string
	Handler string
}

type FrameDoc struct {
	Name                    string
	RawName                 string
	Addon                   string
	Type                    string
	Parent                  string
	ParentType              string // element type tag of the direct parent named frame
	Inherits                string
	Template                bool
	Source                  string
	Children                []string
	ChildElementTypes       []string            // element type tags of named child frames
	StructuralChildTypes    []string            // unnamed structural element type names inside this frame
	StructuralChildAttrKeys map[string][]string // attribute keys per structural child type
	CompositionSnippet      string              // etree-derived XML snippet showing the structural hierarchy
	Attributes              map[string]string
	Handlers                []FrameHandlerDoc
}

type FrameHandlerDoc struct {
	Event    string
	Function string
}

type HandlerDoc struct {
	Addon    string
	Frame    string
	Event    string
	Function string
	Source   string
}

type EventDoc struct {
	Name             string
	LuaRegistrations []EventUsageDoc
	XMLHandlers      []EventHandlerUsageDoc
	TriggeredBy      []string
}

type EventUsageDoc struct {
	Addon     string
	Registrar string
	Handler   string
	Scope     string
}

type EventHandlerUsageDoc struct {
	Addon    string
	Frame    string
	Function string
}

type BindingDoc struct {
	Addon       string
	Frame       string
	Event       string
	XMLFunction string
	LuaFunction string
	Resolved    bool
}

type Confidence string

const (
	ConfidenceHigh   Confidence = "HIGH"
	ConfidenceMedium Confidence = "MEDIUM"
	ConfidenceLow    Confidence = "LOW"
)

type CandidateStatus string

const (
	CandidateStatusCanonical CandidateStatus = "canonical"
	CandidateStatusCandidate CandidateStatus = "candidate"
	CandidateStatusRejected  CandidateStatus = "rejected"
)

type UsageExample struct {
	Addon   string
	Caller  string
	Snippet string
	Source  string
}

type ParameterDoc struct {
	Name     string
	Role     string
	Evidence string
}

type ReturnVariantDoc struct {
	Label      string
	Arity      int
	Shape      []string
	Confidence int
	Notes      []string
}

type ReturnPositionDoc struct {
	Position        int
	InferredType    string
	InferredRole    string
	ConfidenceScore int
	ConfidenceLevel string
	Evidence        []string
	Optional        bool
	Stable          bool
}

type FunctionSymbol struct {
	Name          string
	Category      string
	Confidence    Confidence
	RawScore      int
	Score         int
	Signals       []confidence.Signal
	Rationale     string
	Evidence      confidence.Evidence
	Description   string
	Signature     string
	Parameters    []ParameterDoc
	Returns       []string
	SideEffects   []string
	SeenIn        []string
	Examples      []UsageExample
	Aliases       []string
	ObservedArity int
	Notes         []string

	ObservedReturnArity []int
	InferredReturnArity int
	ReturnVariants      []ReturnVariantDoc
	ReturnPositions     []ReturnPositionDoc
	BranchSensitive     bool
	WrapperAffected     bool
	RuntimeObserved     bool
	ReturnRationale     string
	ReturnUncertainty   []string
}

type EventSymbol struct {
	Name           string
	Category       string
	Confidence     Confidence
	RawScore       int
	Score          int
	Signals        []confidence.Signal
	Rationale      string
	Evidence       confidence.Evidence
	Description    string
	HandlerPattern string
	Payload        []string
	SeenIn         []string
	Examples       []UsageExample
	Registrars     []string
	Notes          []string
}

type XMLHandlerSymbol struct {
	Name            string
	Confidence      Confidence
	RawScore        int
	Score           int
	Signals         []confidence.Signal
	Rationale       string
	Evidence        confidence.Evidence
	Description     string
	ExpectedBinding string
	ElementTypes    []string
	SeenIn          []string
	Examples        []UsageExample
	Notes           []string
}

type FieldSymbol struct {
	Name        string
	Confidence  Confidence
	RawScore    int
	Score       int
	Signals     []confidence.Signal
	Rationale   string
	Evidence    confidence.Evidence
	Description string
	SeenIn      []string
	Notes       []string
}

type TableSymbol struct {
	Name                 string
	Confidence           Confidence
	RawScore             int
	Score                int
	Signals              []confidence.Signal
	Rationale            string
	Evidence             confidence.Evidence
	Description          string
	SeenIn               []string
	Functions            []string
	ObservedMembers      []string
	RepresentativeUsages []UsageExample
	Notes                []string
}

type ConstantSymbol struct {
	Name        string
	Confidence  Confidence
	RawScore    int
	Score       int
	Signals     []confidence.Signal
	Rationale   string
	Evidence    confidence.Evidence
	Description string
	SeenIn      []string
	UsedBy      []string
	Notes       []string
}

// XMLEventBinding records the observed association between one XML handler event
// name and the Lua functions commonly bound to it on a specific element type.
// InferredArgs is a best-effort Lua function signature string; ArgsConfidence
// is "HIGH", "MEDIUM", or "LOW" to indicate how reliable the inference is.
type XMLEventBinding struct {
	Event          string
	LuaFunctions   []string
	InferredArgs   string
	ArgsConfidence string
	Category       string   // Handler category: "lifecycle", "input", "data", "layout", "focus", "tooltip", "custom"
	LuaAPICalls    []string // Per-event downstream API calls made by handlers of this event
}

type ElementTypeSymbol struct {
	Name                    string
	Confidence              Confidence
	RawScore                int
	Score                   int
	Signals                 []confidence.Signal
	Rationale               string
	Evidence                confidence.Evidence
	Description             string
	SeenIn                  []string
	CommonAttributes        []string
	CommonHandlers          []string
	CommonHandlerFunctions  []string // Lua function names commonly bound via XML handlers on this element type
	CommonInherits          []string
	CommonChildTypes        []string          // structural (unnamed) child element types (e.g. ListData in ListBox)
	CommonChildElementTypes []string          // named child frame element types (e.g. Button inside Window)
	CommonParentTypes       []string          // element types that commonly contain this one (e.g. Window for Button)
	CompositionSnippet      string            // representative XML snippet showing the hierarchy of structural children
	XMLEventBindings        []XMLEventBinding // per-event bindings with Lua functions and inferred args
	Examples                []UsageExample
	Notes                   []string

	// Phase 4 enrichment: structured data from the semantic merge pipeline
	AttributeProfiles       []AttributeProfileEntry  // Attribute reference table with required/optional and sample values
	StructuralChildProfiles []StructuralChildProfile // Detailed structural sub-element profiles
	LuaAPICallsFromHandlers []LuaAPICallEntry        // Aggregated Lua API calls made from XML event handlers
	HandlerArgPatterns      []HandlerArgPatternEntry // Expected callback argument patterns per event
	LuaManipulators         []string                 // Lua functions that manipulate frames of this type

	// Phase 4 enrichment: relationship data with counts and confidence
	ParentRefs          []ElementRelRef // Parent element types with observation counts
	ChildRefs           []ElementRelRef // Named child element types with observation counts
	StructuralChildRefs []ElementRelRef // Structural children with observation counts
	InheritsBases       []string        // Template bases commonly inherited
	IsTemplate          bool            // Whether this type commonly acts as a template

	// Phase 4 enrichment: binding resolution statistics
	BindingResolvedPct   int // Percentage of handler bindings resolved to Lua source
	BindingTotalHandlers int // Total handler declarations observed
	BindingResolvedCount int // Number resolved to Lua functions

	// .mod lifecycle enrichment: structured startup-window facts.
	// Notes about startup-created windows are derived from these records.
	StartupWindowFacts []WindowLifecycleSemantic
}

// ElementRelRef describes a relationship to another element type with frequency data.
type ElementRelRef struct {
	Tag        string // Element type tag
	Count      int    // Observation count
	Confidence string // HIGH, MEDIUM, LOW
}

// AttributeProfileEntry describes an observed XML attribute with usage statistics.
type AttributeProfileEntry struct {
	Name         string   // Attribute key
	IsRequired   bool     // Present in >80% of instances
	SampleValues []string // Distinct observed values (up to 8)
	Count        int      // How many elements have this attribute
	TotalCount   int      // Total elements of this type
}

// StructuralChildProfile describes an unnamed structural child element type
// with its own attribute data.
type StructuralChildProfile struct {
	Tag        string                  // Child element tag (e.g. "ListData")
	Count      int                     // Observation count
	Attributes []AttributeProfileEntry // Attributes observed on this child type
}

// LuaAPICallEntry describes an API call commonly made from event handlers.
type LuaAPICallEntry struct {
	Function   string   // Called function name
	Count      int      // How many handlers make this call
	Category   string   // Call classification: "ui", "data", "event", "utility", "unknown"
	FromEvents []string // Which XML events' handlers make this call
}

// HandlerArgPatternEntry describes expected callback argument patterns for a handler event.
type HandlerArgPatternEntry struct {
	Event      string                 // XML event name
	Params     []HandlerExpectedParam // Expected parameters
	Confidence string                 // HIGH, MEDIUM, LOW
}

// HandlerExpectedParam describes one expected parameter in a handler callback.
type HandlerExpectedParam struct {
	Position int    // 0-based position
	Name     string // Expected parameter name
	Type     string // Expected type
	Role     string // Semantic role
}

// ModProvenance records that an enriched fact originated from a .mod manifest.
type ModProvenance struct {
	AddonName   string // Addon whose .mod manifest produced this fact
	HookKind    string // Lifecycle hook kind: "OnInitialize", "OnUpdate", "OnShutdown", or ""
	HookTag     string // Raw element tag of the lifecycle hook
	ActionTag   string // Raw element tag of the action (e.g. "CreateWindow", "CallFunction")
	ActionIndex int    // Zero-based position of the action within its hook
	Resolution  string // "exact", "ambiguous", or "unresolved"
	Confidence  string // "HIGH", "MEDIUM", or "" (empty when unresolved)
}

// FunctionLifecycleRole describes the lifecycle role of a Lua function as
// derived from .mod manifest semantic analysis.
type FunctionLifecycleRole struct {
	FuncName   string // Qualified Lua function name (or RefName when unresolved)
	RefName    string // Original .mod name attribute value
	Role       string // "startup_entrypoint", "shutdown_entrypoint", "update_callback", "unresolved_lifecycle_ref", "unknown_lifecycle_ref"
	Provenance ModProvenance
}

// WindowLifecycleSemantic holds structured facts about a window that is
// created during an addon lifecycle hook, derived from .mod semantic analysis.
type WindowLifecycleSemantic struct {
	FrameName   string // XML frame name from the .mod CreateWindow action
	ElementType string // Element type tag resolved for the frame (empty if unresolved)
	HookKind    string // Lifecycle hook kind: "OnInitialize", "OnUpdate", "OnShutdown", or ""
	Resolution  string // "exact", "ambiguous", or "unresolved"
	Confidence  string // "HIGH", "MEDIUM", or ""
	Provenance  ModProvenance
}

// AddonLifecycleSemantic holds structured addon-level lifecycle facts derived
// from .mod manifest semantic analysis.  One entry per .mod addon.
type AddonLifecycleSemantic struct {
	AddonName string
	HookKinds []string // Distinct lifecycle hook kinds found

	// Actions grouped by lifecycle phase
	StartupActions  []LifecycleActionRecord // Actions from OnInitialize
	ShutdownActions []LifecycleActionRecord // Actions from OnShutdown
	UpdateActions   []LifecycleActionRecord // Actions from OnUpdate
	UnknownActions  []LifecycleActionRecord // Actions from unrecognised lifecycle sections

	// All unresolved action references (across all lifecycle sections)
	UnresolvedRefs []LifecycleActionRecord

	// Addon manifest declarations
	SavedVariables []string
}

// LifecycleExclusion records one addon exclusion decision from lifecycle
// discovery diagnostics.
type LifecycleExclusion struct {
	AddonName  string // Addon display name (manifest name when available)
	Directory  string // Source directory name under SourceRoot
	ReasonCode string // Stable code: no_manifest, no_resolved_source_files, toc_only
	Reason     string // Human-readable reason
}

// LifecycleDiscoveryDiagnostics captures source discovery accounting so that
// lifecycle reports can clearly explain inclusion/exclusion behavior.
type LifecycleDiscoveryDiagnostics struct {
	SourceRoot                 string
	SourceDirectoryCount       int
	ManifestDiscoveredCount    int // Addons discovered by manifest scan (.mod/.toc)
	SourceScannedAddonCount    int // Addons with at least one manifest-listed XML/Lua file on disk
	ModSemanticAddonCount      int // Source-scanned addons with .mod manifest semantics
	LifecycleCatalogAddonCount int // Final addons emitted in AddonLifecycleSemantics

	NoManifestDirectories []string
	Exclusions            []LifecycleExclusion
}

// LifecycleActionRecord is a structured record of a single lifecycle action
// from a .mod manifest.
type LifecycleActionRecord struct {
	ActionKind   string   // "CallFunction", "CreateWindow", or "" (unknown)
	ActionTag    string   // Raw element tag from the .mod file
	Name         string   // "name" attribute value from the action
	HookKind     string   // Parent hook kind
	HookTag      string   // Parent hook raw tag
	Resolution   string   // "exact", "ambiguous", or "unresolved"
	Confidence   string   // "HIGH", "MEDIUM", or ""
	MatchedNames []string // Matched Lua function names or XML frame names
}

type PatternDoc struct {
	Name        string
	Category    string
	Confidence  Confidence
	Description string
	Evidence    []string
}

type LifecyclePhase struct {
	Phase       string
	Description string
	SeenInCount int
	Evidence    []string
}

type Coverage struct {
	SourceCounts               map[string]int
	SymbolCounts               map[string]int
	ConfidenceCounts           map[Confidence]int
	SeenOnceCount              int
	SeenManyCount              int
	UncertainSymbols           []string
	HighConfidencePlatform     []CandidateSummary
	MediumConfidenceCandidates []CandidateSummary
	LowConfidenceSymbols       []CandidateSummary
	RejectedAddonLocal         []CandidateSummary
	ExplanationMatrix          []ConfidenceMatrixRow
}

type CandidateSummary struct {
	Name       string
	Category   string
	Status     CandidateStatus
	RawScore   int
	Score      int
	Confidence Confidence
	Signals    []confidence.Signal
	Rationale  string
	Evidence   confidence.Evidence
}

type ConfidenceMatrixRow struct {
	Name         string
	Category     string
	Score        int
	Confidence   Confidence
	CrossAddon   int
	Namespace    int
	XMLExposure  int
	EventLinkage int
	DefaultUI    int
	LocalPenalty int
	Signature    int
}

type DocLink struct {
	ID         string     `json:"id,omitempty"`
	Label      string     `json:"label"`
	Type       string     `json:"type"`
	Category   string     `json:"category,omitempty"`
	Path       string     `json:"path,omitempty"`
	Confidence Confidence `json:"confidence,omitempty"`
	Score      int        `json:"score,omitempty"`
	Summary    string     `json:"summary,omitempty"`
}

type SemanticLinks struct {
	RelatedAPIs []DocLink `json:"related_apis,omitempty"`
	UsedWith    []DocLink `json:"used_with,omitempty"`
	TriggeredBy []DocLink `json:"triggered_by,omitempty"`
	Affects     []DocLink `json:"affects,omitempty"`
}

type NavigationNode struct {
	ID         string           `json:"id,omitempty"`
	Label      string           `json:"label"`
	Type       string           `json:"type"`
	Path       string           `json:"path,omitempty"`
	Importance int              `json:"importance,omitempty"`
	Children   []NavigationNode `json:"children,omitempty"`
}

type SidebarItem struct {
	Label string        `json:"label"`
	Type  string        `json:"type,omitempty"`
	Path  string        `json:"path,omitempty"`
	Order int           `json:"order,omitempty"`
	Items []SidebarItem `json:"items,omitempty"`
}

type GraphNode struct {
	ID         string     `json:"id"`
	Label      string     `json:"label"`
	Type       string     `json:"type"`
	Category   string     `json:"category,omitempty"`
	Path       string     `json:"path,omitempty"`
	Confidence Confidence `json:"confidence,omitempty"`
	Score      int        `json:"score,omitempty"`
	Summary    string     `json:"summary,omitempty"`
}

type GraphEdge struct {
	From            string   `json:"from"`
	To              string   `json:"to"`
	Type            string   `json:"type"`
	ConfidenceScore int      `json:"confidence_score,omitempty"` // 0-100
	EvidenceCount   int      `json:"evidence_count,omitempty"`   // number of observations
	EvidenceSources []string `json:"evidence_sources,omitempty"` // file:line references
	Rationale       string   `json:"rationale,omitempty"`        // explanation
	Weight          int      `json:"weight,omitempty"`           // backward compatibility
	Evidence        []string `json:"evidence,omitempty"`         // backward compatibility
}

type APIGraph struct {
	GeneratedAt time.Time   `json:"generated_at"`
	Nodes       []GraphNode `json:"nodes"`
	Edges       []GraphEdge `json:"edges"`
}

type RelationChain struct {
	ID           string    `json:"id"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	Description  string    `json:"description,omitempty"`
	Participants []DocLink `json:"participants,omitempty"`
	Flow         []string  `json:"flow,omitempty"`
	Evidence     []string  `json:"evidence,omitempty"`
	Weight       int       `json:"weight,omitempty"`
}

type RelationReport struct {
	GeneratedAt         time.Time       `json:"generated_at"`
	FrequentCallChains  []RelationChain `json:"frequent_call_chains,omitempty"`
	CommonCombinations  []RelationChain `json:"common_api_combinations,omitempty"`
	EventHandlerUIFlows []RelationChain `json:"event_handler_ui_flows,omitempty"`
}

type PatternPage struct {
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Category    string     `json:"category"`
	Confidence  Confidence `json:"confidence"`
	Description string     `json:"description"`
	Involved    []DocLink  `json:"involved_apis,omitempty"`
	Flow        []string   `json:"flow,omitempty"`
	ExampleCode string     `json:"example_code,omitempty"`
	Evidence    []string   `json:"evidence,omitempty"`
	Path        string     `json:"path"`
}

type Corpus struct {
	SourceRoot           string
	Contracts            ContractModel
	GlobalFunctions      []FunctionSymbol
	WindowFunctions      []FunctionSymbol
	GlobalTables         []TableSymbol
	Constants            []ConstantSymbol
	ElementTypes         []ElementTypeSymbol
	XMLHandlers          []XMLHandlerSymbol
	GameEvents           []EventSymbol
	WindowEvents         []EventSymbol
	SystemDataFields     []FieldSymbol
	GameDataFields       []FieldSymbol
	SlashCommandPatterns []PatternDoc
	WindowPatterns       []PatternDoc
	EventPatterns        []PatternDoc
	Lifecycle            []LifecyclePhase
	Conventions          []PatternDoc
	InferenceRules       []string
	Coverage             Coverage
	SymbolLinks          map[string]SemanticLinks
	NavigationTree       NavigationNode
	Sidebar              []SidebarItem
	APIGraph             APIGraph
	Relations            RelationReport
	PatternPages         []PatternPage
	GeneratedAt          time.Time

	// .mod lifecycle enrichment: addon-level and function-level lifecycle facts.
	AddonLifecycleSemantics []AddonLifecycleSemantic // One per .mod addon
	FunctionLifecycleRoles  []FunctionLifecycleRole  // Function-level lifecycle roles
	LifecycleDiagnostics    LifecycleDiscoveryDiagnostics
}

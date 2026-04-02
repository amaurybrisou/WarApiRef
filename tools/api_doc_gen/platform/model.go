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
	Name                     string
	Addon                    string
	Type                     string
	Parent                   string
	ParentType               string              // element type tag of the direct parent named frame
	Inherits                 string
	Template                 bool
	Source                   string
	Children                 []string
	ChildElementTypes        []string            // element type tags of named child frames
	StructuralChildTypes     []string            // unnamed structural element type names inside this frame
	StructuralChildAttrKeys  map[string][]string // attribute keys per structural child type
	// StructuralChildAttrValues samples up to 8 unique observed values per structural
	// child attribute key, enabling attribute-role inference in the platform build step.
	StructuralChildAttrValues map[string]map[string][]string
	CompositionSnippet        string              // etree-derived XML snippet showing the structural hierarchy
	Attributes                map[string]string
	Handlers                  []FrameHandlerDoc
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

type FlowDoc struct {
	Addon string
	Steps []FlowStepDoc
}

type FlowStepDoc struct {
	Phase    string
	Detail   string
	Evidence []string
}

type ExampleDoc struct {
	Addon       string
	Frame       string
	Event       string
	LuaFunction string
}

type GlobalNamespaceDoc struct {
	Addon string
	Name  string
}

type SavedVariableDoc struct {
	Addon string
	Name  string
}

type GlobalsDoc struct {
	Namespaces     []GlobalNamespaceDoc
	SavedVariables []SavedVariableDoc
}

type SourceModel struct {
	Root      string
	Functions []FunctionDoc
	Frames    []FrameDoc
	Handlers  []HandlerDoc
	Events    []EventDoc
	Bindings  []BindingDoc
	Flows     []FlowDoc
	Examples  []ExampleDoc
	Globals   GlobalsDoc
	LoadedAt  time.Time
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
	Event         string
	LuaFunctions  []string
	InferredArgs  string
	ArgsConfidence string
}

// StructuralChildProfile captures the observed structural (unnamed) XML child
// element for an element type, with its own attribute keys and sampled values.
// This enables element pages to show not just "ListData" but also what attributes
// it carries (e.g. populationfunction, table) and what those values look like.
type StructuralChildProfile struct {
	Tag              string
	ObservedCount    int                    // how many frames of the parent type carry this child
	CommonAttrKeys   []string               // attribute names ordered by frequency
	AttrValueSamples map[string][]string    // key → up to 8 representative unique values
}

// AttributeProfile captures observed attribute semantics for an XML element type,
// including sampled values and a best-effort inference of the attribute's role.
// InferredRole is one of: "boolean", "number", "frame-ref", "lua-function", "string".
type AttributeProfile struct {
	Name          string
	ObservedCount int
	InferredRole  string   // best-effort classification of the attribute's value type
	SampleValues  []string // up to 8 representative unique observed values
	Confidence    string   // HIGH / MEDIUM / LOW — reflects how many samples informed the inference
}

type ElementTypeSymbol struct {
	Name                     string
	Confidence               Confidence
	RawScore                 int
	Score                    int
	Signals                  []confidence.Signal
	Rationale                string
	Evidence                 confidence.Evidence
	Description              string
	SeenIn                   []string
	CommonAttributes         []string
	CommonHandlers           []string
	CommonHandlerFunctions   []string // Lua function names commonly bound via XML handlers on this element type
	CommonInherits           []string
	CommonChildTypes         []string // structural (unnamed) child element types (e.g. ListData in ListBox)
	CommonChildElementTypes  []string // named child frame element types (e.g. Button inside Window)
	CommonParentTypes        []string // element types that commonly contain this one (e.g. Window for Button)
	CompositionSnippet       string   // representative XML snippet showing the hierarchy of structural children
	XMLEventBindings         []XMLEventBinding // per-event bindings with Lua functions and inferred args
	// StructuralChildProfiles replaces/extends CommonChildTypes with per-child attribute profiles.
	// Each entry describes one observed structural (unnamed) XML child type and its attribute conventions.
	StructuralChildProfiles  []StructuralChildProfile
	// AttributeProfiles extends CommonAttributes with value samples and inferred attribute roles.
	AttributeProfiles        []AttributeProfile
	Examples                 []UsageExample
	Notes                    []string
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
	Source               SourceModel
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
}

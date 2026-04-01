package deep_analysis

// FunctionSource is the minimal source input for advanced return analysis.
type FunctionSource struct {
	Path   string
	Name   string
	Addon  string
	Source string
}

// ReturnPositionInference stores per-slot return inference.
type ReturnPositionInference struct {
	Position        int
	InferredType    string
	InferredRole    string
	ConfidenceScore int
	ConfidenceLevel string
	Evidence        []string
	Optional        bool
	Stable          bool
	RuntimeObserved bool
}

// ReturnVariant captures one observed/plausible return shape.
type ReturnVariant struct {
	Label      string
	Arity      int
	Shape      []string
	Contexts   []string
	Confidence int
	Notes      []string
}

// AdvancedReturnReport contains full function return reconstruction details.
type AdvancedReturnReport struct {
	FunctionPath        string
	ObservedReturnArity []int
	InferredReturnArity int
	ReturnVariants      []ReturnVariant
	ReturnPositions     []ReturnPositionInference
	BranchSensitive     bool
	WrapperAffected     bool
	RuntimeObserved     bool
	Rationale           string
	UncertaintyNotes    []string
	Evidence            []ProvenanceRecord
}

func confidenceLevelFromScore(score int) string {
	switch {
	case score >= 75:
		return "HIGH"
	case score >= 50:
		return "MEDIUM"
	default:
		return "LOW"
	}
}

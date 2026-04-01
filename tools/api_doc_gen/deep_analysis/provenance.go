package deep_analysis

import "fmt"

// ProvenanceRecord captures where an inference came from.
type ProvenanceRecord struct {
	FunctionPath   string
	SourceFunction string
	Line           int
	Pattern        string
	Snippet        string
	Observation    string
}

func (p ProvenanceRecord) Summary() string {
	if p.Line > 0 {
		return fmt.Sprintf("%s:%d %s", p.SourceFunction, p.Line, p.Observation)
	}
	return fmt.Sprintf("%s %s", p.SourceFunction, p.Observation)
}

package deep_analysis

import (
	"fmt"
	"sort"
	"strings"
)

// EvidenceSource represents where evidence came from
type EvidenceSource struct {
	File        string // path to addon/generated docs file
	Line        int    // line number if applicable
	Type        string // "call_site", "return_location", "xml_binding", "event_handler", etc
	Description string // human-readable evidence details
}

// Evidence tracks observations about a symbol with provenance
type Evidence struct {
	Sources   []EvidenceSource // all places this was observed
	Count     int              // how many independent observations
	Patterns  []string         // identified patterns (e.g., "always_assigned_to_boolean_var")
	Weight    float64          // calculated importance
	Rationale string           // explanation of what this evidence means
}

// ConfidenceLevel describes how confident we are in an inference
type ConfidenceLevel struct {
	Score      int    // 0-100
	Level      string // "HIGH", "MEDIUM", "LOW", "UNCERTAIN"
	Basis      string // brief explanation
	Caveats    string // known limitations or alternative explanations
	RuntimeObs bool   // true if any runtime evidence supports this
}

// Inference represents a single inferred fact about a symbol
type Inference struct {
	Subject    string          // what we're inferring about (function name, parameter, etc)
	Property   string          // what property (return_type, argument_role, edge_type)
	Value      interface{}     // the inferred value
	Evidence   Evidence        // supporting evidence
	Confidence ConfidenceLevel // how confident we are
	Observed   bool            // true if directly observed (not inferred)
	Uncertain  bool            // true if we're not sure enough to include
}

// EdgeRelation represents a relationship between two symbols
type EdgeRelation struct {
	From                string // source symbol ID
	To                  string // target symbol ID
	RelationType        string // "calls", "reads_data", "triggered_by", etc
	ConfidenceScore     int    // 0-100
	EvidenceCount       int    // how many independent observations
	EvidenceSources     []EvidenceSource
	Rationale           string
	SupportsCommonUsage bool // used together in practice?
	IsTransitive        bool // only holds through indirect relationships?
}

// NewEvidence creates evidence with source tracking
func NewEvidence(sourceFile string, sourceLine int, sourceType, description string) Evidence {
	return Evidence{
		Sources: []EvidenceSource{
			{File: sourceFile, Line: sourceLine, Type: sourceType, Description: description},
		},
		Count: 1,
	}
}

// AddSource adds another observation point to the evidence
func (e *Evidence) AddSource(sourceFile string, sourceLine int, sourceType, description string) {
	e.Sources = append(e.Sources, EvidenceSource{
		File:        sourceFile,
		Line:        sourceLine,
		Type:        sourceType,
		Description: description,
	})
	e.Count++
	e.UpdateWeight()
}

// AddPattern marks an observed pattern
func (e *Evidence) AddPattern(pattern string) {
	e.Patterns = append(e.Patterns, pattern)
	e.UpdateWeight()
}

// UpdateWeight recalculates evidence importance
func (e *Evidence) UpdateWeight() {
	// Weight based on count and pattern multiplicity
	e.Weight = float64(e.Count) + float64(len(e.Patterns))*0.5
}

// Sources returns unique evidence locations
func (e Evidence) UniqueSourceLocations() []string {
	locations := []string{}
	seen := map[string]bool{}
	for _, src := range e.Sources {
		loc := fmt.Sprintf("%s:%d", src.File, src.Line)
		if src.Line == 0 {
			loc = src.File
		}
		if !seen[loc] {
			locations = append(locations, loc)
			seen[loc] = true
		}
	}
	sort.Strings(locations)
	return locations
}

// SourcesByType groups evidence by type
func (e Evidence) SourcesByType() map[string]int {
	count := map[string]int{}
	for _, src := range e.Sources {
		count[src.Type]++
	}
	return count
}

// Summary provides a brief evidence summary
func (e Evidence) Summary() string {
	if e.Count == 0 {
		return "no evidence"
	}
	byType := e.SourcesByType()
	parts := []string{}
	for t := range byType {
		parts = append(parts, fmt.Sprintf("%d %s", byType[t], t))
	}
	sort.Strings(parts)
	return strings.Join(parts, ", ")
}

// ScoreConfidence calculates confidence based on evidence
func ScoreConfidence(evidenceCount int, patterns []string, hasRuntimeObs bool, contradictions int) ConfidenceLevel {
	score := 50                // base score
	score += evidenceCount * 8 // +8 per independent observation (capped at 100)
	score += len(patterns) * 5 // +5 per pattern
	if hasRuntimeObs {
		score += 15 // significant boost for runtime observation
	}
	score -= contradictions * 10 // penalties for conflicting evidence

	if score > 100 {
		score = 100
	}
	if score < 0 {
		score = 0
	}

	var level, basis string
	if score >= 80 {
		level = "HIGH"
		basis = "strong evidence from multiple sources"
	} else if score >= 60 {
		level = "MEDIUM"
		basis = "moderate evidence, fairly confident"
	} else if score >= 40 {
		level = "LOW"
		basis = "limited evidence, tentative"
	} else {
		level = "UNCERTAIN"
		basis = "insufficient evidence for confident inference"
	}

	return ConfidenceLevel{
		Score: score,
		Level: level,
		Basis: basis,
	}
}

package deep_analysis

import (
	"fmt"
	"strings"
)

// ConfidenceScorer calculates evidence-based confidence scores for inferences
type ConfidenceScorer struct {
	DefaultWeights map[string]float64
}

// NewConfidenceScorer creates scorer with default evidence weights
func NewConfidenceScorer() *ConfidenceScorer {
	return &ConfidenceScorer{
		DefaultWeights: map[string]float64{
			// Return type evidence weights
			"explicit_annotation": 1.0,  // 100% confidence if explicitly stated
			"literal":             0.75, // 75% for literals (usually correct)
			"assignment_target":   0.60, // 60% for how it's assigned
			"comparison":          0.70, // 70% for comparison contexts
			"table_expected":      0.65, // 65% if used as table
			"call_result":         0.50, // 50% if returned from unknown

			// Argument evidence weights
			"call_site":         0.75, // 75% for observed call sites
			"usage_pattern":     0.65, // 65% for usage patterns
			"field_access":      0.70, // 70% for field access
			"indexing":          0.75, // 75% for bracket indexing
			"method_call":       0.80, // 80% for method calls
			"boolean_condition": 0.75, // 75% for if conditions
			"arithmetic":        0.85, // 85% for arithmetic operations
			"array_like":        0.85, // 85% for array iteration
			"table_like":        0.85, // 85% for table iteration
			"callable":          0.90, // 90% if called as function

			// Edge evidence weights
			"direct_observation":  0.95, // 95% for direct observation
			"static_analysis":     0.70, // 70% for static analysis
			"pattern_match":       0.60, // 60% for pattern matching
			"inferred_from_usage": 0.55, // 55% for usage inference
			"heuristic":           0.50, // 50% for heuristics
			"runtime_observed":    0.99, // 99% for runtime observations
		},
	}
}

// ScoreEvidence calculates confidence for a set of evidence items
func (cs *ConfidenceScorer) ScoreEvidence(evidences []Evidence) int {
	if len(evidences) == 0 {
		return 0
	}

	totalWeight := 0.0
	observationCount := 0

	for _, evidence := range evidences {
		// Each evidence piece contributes
		baseWeight := cs.DefaultWeights[evidence.Sources[0].Type]
		if baseWeight == 0 {
			baseWeight = 0.50 // default for unknown types
		}

		// Multiple observations increase confidence
		for range evidence.Sources {
			observationCount++
		}

		// Patterns increase confidence
		patternBonus := float64(len(evidence.Patterns)) * 0.05

		totalWeight += baseWeight + patternBonus
	}

	// Average confidence, but boost with multiple observations
	avgConfidence := (totalWeight / float64(len(evidences))) * 100.0

	// Observation bonus: multiple independent observations boost confidence
	if observationCount > 1 {
		avgConfidence += float64((observationCount - 1) * 5)
	}

	// Cap at 100
	if avgConfidence > 100 {
		avgConfidence = 100
	}

	return int(avgConfidence)
}

// ScoreReturnTypeConfidence scores a return type inference
func (cs *ConfidenceScorer) ScoreReturnTypeConfidence(analysis *ReturnAnalysis) int {
	if len(analysis.ExplicitReturns) > 0 {
		return 95 // Explicit annotations are very reliable
	}

	// Score based on evidence count and type
	score := 30 // base score

	// Literal returns are good evidence
	score += len(analysis.LiteralReturns) * 15

	// Multiple independent inferred types suggest confidence
	if len(analysis.InferredTypes) > 0 {
		score += 20
	}
	if len(analysis.InferredTypes) > 1 {
		score += 15 // Multiple different inferred types reduce confidence slightly
	}

	// Assignment and comparison contexts add evidence
	score += len(analysis.AssignmentContexts) * 10
	score += len(analysis.ComparisonContexts) * 10

	// Multiple evidences boost
	if len(analysis.Evidence) > 3 {
		score += 15
	}

	// Cap at reasonable max
	if score > 85 {
		score = 85 // Never show 100% confidence for inferred types
	}

	return score
}

// ScoreArgumentTypeConfidence scores an argument type inference
func (cs *ConfidenceScorer) ScoreArgumentTypeConfidence(analysis *ArgumentAnalysis) int {
	score := 40 // base score

	// Count type signals
	signalCount := len(analysis.TypeSignals)
	score += signalCount * 12 // +12 per distinct signal

	// Usage patterns add evidence
	score += len(analysis.Usage) * 8

	// If same type appears multiple times in signals, boost confidence
	typeFrequency := make(map[string]int)
	for _, sig := range analysis.TypeSignals {
		typeFrequency[sig.Signal]++
	}
	for _, freq := range typeFrequency {
		if freq > 1 {
			score += 10 // +10 for reinforced signals
		}
	}

	// Call patterns evidence
	score += len(analysis.Evidence) * 5

	// Cap at 80 for inferred arguments
	if score > 80 {
		score = 80
	}

	return score
}

// ScoreEdgeConfidence scores a relationship inference
func (cs *ConfidenceScorer) ScoreEdgeConfidence(edge *EdgeRelation) int {
	score := 50 // base score

	// Evidence count is a primary factor
	score += edge.EvidenceCount * 15

	// Multiple evidence sources
	if len(edge.EvidenceSources) > 1 {
		score += 15
	}

	// Direct observation is highly confident
	for _, src := range edge.EvidenceSources {
		if src.Type == "direct_observation" || src.Type == "runtime_observed" {
			score += 25
		}
		if src.Type == "static_analysis" {
			score += 10
		}
	}

	// Cap based on relation type certainty
	score = cs.capScoreByRelationType(score, edge.RelationType)

	return score
}

// capScoreByRelationType limits confidence based on how certain this relation type typically is
func (cs *ConfidenceScorer) capScoreByRelationType(score int, relType string) int {
	switch relType {
	case "calls":
		// Very high confidence possible for function calls
		if score > 95 {
			score = 95
		}
	case "handled_by":
		// Event handlers are generally well-defined
		if score > 90 {
			score = 90
		}
	case "triggered_by":
		// Event triggers are fairly definite
		if score > 90 {
			score = 90
		}
	case "defined_in":
		// Definitions are usually 100% certain
		if score > 99 {
			score = 99
		}
	case "reads_systemdata", "reads_gamedata":
		// Data access can be inferred but has some uncertainty
		if score > 85 {
			score = 85
		}
	case "commonly_used_with":
		// Co-occurrence patterns are weaker signals
		if score > 75 {
			score = 75
		}
	case "appears_in_same_flow":
		// Sequence inference has moderate confidence
		if score > 70 {
			score = 70
		}
	case "updates_ui", "toggles_visibility", "updates_layout":
		// UI operations can be inferred from patterns
		if score > 80 {
			score = 80
		}
	case "initializes", "refreshes":
		// Init/refresh patterns are fairly deterministic
		if score > 85 {
			score = 85
		}
	default:
		// Unknown relation types get lower caps
		if score > 65 {
			score = 65
		}
	}

	return score
}

// ConfidenceRating provides interpretation of confidence score
func ConfidenceRating(score int) string {
	switch {
	case score >= 95:
		return "DEFINITE"
	case score >= 85:
		return "VERY_HIGH"
	case score >= 75:
		return "HIGH"
	case score >= 60:
		return "MEDIUM"
	case score >= 45:
		return "LOW"
	case score >= 30:
		return "VERY_LOW"
	default:
		return "UNCERTAIN"
	}
}

// ShouldIncludeInference determines if an inference should be included in final results
func (cs *ConfidenceScorer) ShouldIncludeInference(confidence int, inferenceType string) bool {
	switch inferenceType {
	case "return_type":
		// Include return types with medium+ confidence
		return confidence >= 50
	case "argument_type":
		// Include argument types with medium+ confidence
		return confidence >= 55
	case "edge":
		// Include edges with medium+ confidence
		return confidence >= 50
	case "edge_runtime":
		// Runtime-observed edges included with lower threshold
		return confidence >= 40
	default:
		return confidence >= 60
	}
}

// ExplanationForScore provides human-readable explanation of confidence
func ExplanationForScore(score int, evidence string) string {
	rating := ConfidenceRating(score)
	explanation := fmt.Sprintf("Confidence: %s (%d%%)", rating, score)

	if evidence != "" {
		explanation += fmt.Sprintf("\nBased on: %s", evidence)
	}

	// Add interpretation
	switch rating {
	case "DEFINITE":
		explanation += "\nThis inference is definite and should be trusted."
	case "VERY_HIGH":
		explanation += "\nThis inference is very likely to be correct."
	case "HIGH":
		explanation += "\nThis inference is probably correct."
	case "MEDIUM":
		explanation += "\nThis inference is reasonably likely, but has some uncertainty."
	case "LOW":
		explanation += "\nThis inference should be treated tentatively."
	case "VERY_LOW":
		explanation += "\nThis inference is speculative and uncertain."
	case "UNCERTAIN":
		explanation += "\nThis inference is not sufficiently supported by evidence."
	}

	return explanation
}

// CombineConfidences merges multiple confidence scores
func CombineConfidences(scores ...int) int {
	if len(scores) == 0 {
		return 0
	}

	// Use geometric mean for combining confidences
	// This prevents 100% * 50% = 50% (which is wrong)
	// Instead: sqrt(100 * 50) = 70%, which better reflects "mostly confident"
	product := 1.0
	for _, score := range scores {
		product *= float64(score) / 100.0
	}

	// Geometric mean
	geometricMean := 1.0
	for range scores {
		geometricMean *= product / float64(len(scores))
	}

	combinedScore := geometricMean * 100.0
	if combinedScore > 100 {
		combinedScore = 100
	}

	return int(combinedScore)
}

// AggregateEvidence combines multiple evidence items into a single assessment
func (cs *ConfidenceScorer) AggregateEvidence(evidences map[string]Evidence) (int, string) {
	if len(evidences) == 0 {
		return 0, "no evidence"
	}

	totalScore := 0.0
	totalWeight := 0.0

	evidenceDetails := []string{}

	for evidenceType, evidence := range evidences {
		weight := cs.DefaultWeights[evidenceType]
		if weight == 0 {
			weight = 0.5 // default
		}

		// Weight by number of observations
		adjustedWeight := weight * float64(evidence.Count)
		totalScore += adjustedWeight * 100.0
		totalWeight += adjustedWeight

		detail := fmt.Sprintf("%s (n=%d)", evidenceType, evidence.Count)
		evidenceDetails = append(evidenceDetails, detail)
	}

	if totalWeight == 0 {
		return 0, "no weighted evidence"
	}

	finalScore := totalScore / totalWeight
	if finalScore > 100 {
		finalScore = 100
	}

	explanation := strings.Join(evidenceDetails, "; ")
	return int(finalScore), explanation
}

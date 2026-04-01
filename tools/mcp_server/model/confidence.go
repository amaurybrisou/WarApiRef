package model

// Confidence bundles both source confidence and semantic confidence.
type Confidence struct {
	ConfidenceScore             int    `json:"confidence_score"`
	ConfidenceLevel             string `json:"confidence_level"`
	ConfidenceRationale         string `json:"confidence_rationale"`
	SemanticConfidenceScore     int    `json:"semantic_confidence_score"`
	SemanticConfidenceLevel     string `json:"semantic_confidence_level"`
	SemanticConfidenceRationale string `json:"semantic_confidence_rationale"`
}

func LevelFromScore(score int) string {
	switch {
	case score >= 80:
		return "HIGH"
	case score >= 55:
		return "MEDIUM"
	default:
		return "LOW"
	}
}

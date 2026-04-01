package model

// APIError is the shared structured error model for tool/resource failures.
type APIError struct {
	ErrorCode        string           `json:"error_code"`
	ErrorMessage     string           `json:"error_message"`
	Details          map[string]any   `json:"details,omitempty"`
	Suggestion       string           `json:"suggestion,omitempty"`
	CandidateMatches []CandidateMatch `json:"candidate_matches,omitempty"`
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	return e.ErrorCode + ": " + e.ErrorMessage
}

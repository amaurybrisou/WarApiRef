package model

// Warning is a structured uncertainty signal returned by tools/resources.
type Warning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SymbolSummary is the shared canonical object shape returned by tools.
type SymbolSummary struct {
	ID             string                 `json:"id"`
	Kind           string                 `json:"kind"`
	CanonicalName  string                 `json:"canonical_name"`
	Title          string                 `json:"title"`
	Summary        string                 `json:"summary"`
	Type           string                 `json:"type"`
	Namespace      string                 `json:"namespace"`
	DocPath        string                 `json:"doc_path"`
	NavigationPath string                 `json:"navigation_path"`
	RelatedIDs     []string               `json:"related_ids"`
	Examples       []string               `json:"examples"`
	Warnings       []Warning              `json:"warnings"`
	Metadata       map[string]interface{} `json:"metadata"`
	Confidence
}

// CandidateMatch is used in suggestions and alternatives.
type CandidateMatch struct {
	ID            string `json:"id"`
	CanonicalName string `json:"canonical_name"`
	Type          string `json:"type"`
	DocPath       string `json:"doc_path"`
}

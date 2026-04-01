package schema

import "roraddons/tools/mcp_server/model"

type LookupSymbolResponse struct {
	ExactMatch   bool                   `json:"exact_match"`
	Symbol       model.SymbolSummary    `json:"symbol"`
	Alternatives []model.CandidateMatch `json:"alternatives,omitempty"`
	Warnings     []model.Warning        `json:"warnings,omitempty"`
}

type SearchMatch struct {
	ID                      string `json:"id"`
	CanonicalName           string `json:"canonical_name"`
	Type                    string `json:"type"`
	Namespace               string `json:"namespace"`
	Summary                 string `json:"summary"`
	ConfidenceScore         int    `json:"confidence_score"`
	SemanticConfidenceScore int    `json:"semantic_confidence_score"`
	DocPath                 string `json:"doc_path"`
}

type SearchSymbolsResponse struct {
	TotalMatches int             `json:"total_matches"`
	Matches      []SearchMatch   `json:"matches"`
	Warnings     []model.Warning `json:"warnings,omitempty"`
}

type RelatedItem struct {
	ID              string `json:"id"`
	CanonicalName   string `json:"canonical_name"`
	Type            string `json:"type"`
	RelationType    string `json:"relation_type"`
	Summary         string `json:"summary"`
	ConfidenceScore int    `json:"confidence_score"`
	DocPath         string `json:"doc_path"`
}

type GetRelatedSymbolsResponse struct {
	SourceSymbol model.SymbolSummary      `json:"source_symbol"`
	Relations    map[string][]RelatedItem `json:"relations"`
	GraphSummary map[string]any           `json:"graph_summary"`
	Warnings     []model.Warning          `json:"warnings,omitempty"`
}

type FindUsageExamplesResponse struct {
	QueryType    string               `json:"query_type"`
	TotalMatches int                  `json:"total_matches"`
	Examples     []model.UsageExample `json:"examples"`
}

type GeneratedArtifact struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

type ScaffoldAddonSnippetResponse struct {
	GeneratedArtifacts    []GeneratedArtifact   `json:"generated_artifacts"`
	UsedDocumentedSymbols []model.SymbolSummary `json:"used_documented_symbols"`
	AvoidedSymbols        []string              `json:"avoided_symbols"`
	UncertaintyNotes      []string              `json:"uncertainty_notes"`
	Rationale             string                `json:"rationale"`
}

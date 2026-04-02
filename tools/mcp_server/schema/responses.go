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

type IngestObservationResponse struct {
	Accepted      bool            `json:"accepted"`
	EntryID       string          `json:"entry_id,omitempty"`
	SchemaVersion string          `json:"schema_version,omitempty"`
	SourcePath    string          `json:"source_path,omitempty"`
	Errors        []string        `json:"errors,omitempty"`
	Warnings      []model.Warning `json:"warnings,omitempty"`
}

type IngestObservationBatchItem struct {
	SourcePath string   `json:"source_path"`
	Accepted   bool     `json:"accepted"`
	EntryID    string   `json:"entry_id,omitempty"`
	Errors     []string `json:"errors,omitempty"`
}

type IngestObservationBatchResponse struct {
	TotalFiles     int                          `json:"total_files"`
	ProcessedFiles int                          `json:"processed_files"`
	AcceptedCount  int                          `json:"accepted_count"`
	RejectedCount  int                          `json:"rejected_count"`
	Results        []IngestObservationBatchItem `json:"results"`
	Warnings       []model.Warning              `json:"warnings,omitempty"`
}

// ObservationSummary is a compact representation of a queued observation for listing.
type ObservationSummary struct {
	ObservationID string   `json:"observation_id"`
	Status        string   `json:"status"`
	TargetSeeds   []string `json:"target_seeds"`
	ClaimSummary  string   `json:"claim_summary"`
	Confidence    string   `json:"confidence"`
	EvidenceCount int      `json:"evidence_count"`
	CreatedAt     string   `json:"created_at"`
	SourceAddon   string   `json:"source_addon,omitempty"`
}

type ListPendingObservationsResponse struct {
	TotalCount   int                  `json:"total_count"`
	Observations []ObservationSummary `json:"observations"`
	Warnings     []model.Warning      `json:"warnings,omitempty"`
}

type ReviewObservationResponse struct {
	ObservationID string          `json:"observation_id"`
	Status        string          `json:"status"`
	Verdict       string          `json:"verdict"`
	Warnings      []model.Warning `json:"warnings,omitempty"`
	Errors        []string        `json:"errors,omitempty"`
}

// SeedUpdate describes one seed-file change performed during promotion.
type SeedUpdate struct {
	SeedPath  string `json:"seed_path"`
	Appended  bool   `json:"appended"`
	Duplicate bool   `json:"duplicate,omitempty"`
}

type PromoteObservationResponse struct {
	ObservationID string          `json:"observation_id"`
	Promoted      bool            `json:"promoted"`
	TargetSeeds   []string        `json:"target_seeds"`
	SeedUpdates   []SeedUpdate    `json:"seed_updates"`
	DryRun        bool            `json:"dry_run,omitempty"`
	Warnings      []model.Warning `json:"warnings,omitempty"`
	Errors        []string        `json:"errors,omitempty"`
}

// RejectedObservationSummary is a compact view of a rejected record.
type RejectedObservationSummary struct {
	ObservationID   string   `json:"observation_id"`
	Status          string   `json:"status"`
	TargetSeeds     []string `json:"target_seeds"`
	ClaimSummary    string   `json:"claim_summary"`
	Confidence      string   `json:"confidence"`
	RejectionReason string   `json:"rejection_reason,omitempty"`
	RejectedAt      string   `json:"rejected_at,omitempty"`
	Reviewer        string   `json:"reviewer,omitempty"`
	CreatedAt       string   `json:"created_at"`
}

type ListRejectedObservationsResponse struct {
	TotalCount   int                          `json:"total_count"`
	Observations []RejectedObservationSummary `json:"observations"`
	Warnings     []model.Warning              `json:"warnings,omitempty"`
}

// RegenerationStep describes one command in a regeneration run.
type RegenerationStep struct {
	Label   string `json:"label"`
	Command string `json:"command"`
	Success bool   `json:"success,omitempty"`
	Output  string `json:"output,omitempty"`
	Error   string `json:"error,omitempty"`
}

type RegenerateResponse struct {
	Scope    string             `json:"scope"`
	DryRun   bool               `json:"dry_run"`
	Steps    []RegenerationStep `json:"steps"`
	Success  bool               `json:"success"`
	Errors   []string           `json:"errors,omitempty"`
	Warnings []model.Warning    `json:"warnings,omitempty"`
}

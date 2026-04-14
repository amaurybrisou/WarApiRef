package schema

type LookupSymbolRequest struct {
	SymbolName      string `json:"symbol_name"`
	SymbolType      string `json:"symbol_type,omitempty"`
	IncludeExamples bool   `json:"include_examples,omitempty"`
	IncludeRelated  bool   `json:"include_related,omitempty"`
}

type SearchSymbolsRequest struct {
	Query           string `json:"query"`
	TypeFilter      string `json:"type_filter,omitempty"`
	NamespaceFilter string `json:"namespace_filter,omitempty"`
	ConfidenceMin   int    `json:"confidence_min,omitempty"`
	Limit           int    `json:"limit,omitempty"`
	IncludeExamples bool   `json:"include_examples,omitempty"`
}

type GetRelatedSymbolsRequest struct {
	SymbolName       string   `json:"symbol_name"`
	RelationTypes    []string `json:"relation_types,omitempty"`
	Depth            int      `json:"depth,omitempty"`
	LimitPerRelation int      `json:"limit_per_relation,omitempty"`
}

type GetEventFlowRequest struct {
	EventName         string `json:"event_name"`
	IncludeExamples   bool   `json:"include_examples,omitempty"`
	IncludeDownstream bool   `json:"include_downstream,omitempty"`
}

type GetXMLHandlerInfoRequest struct {
	HandlerName     string `json:"handler_name"`
	IncludeExamples bool   `json:"include_examples,omitempty"`
	IncludeRelated  bool   `json:"include_related,omitempty"`
}

type FindUsageExamplesRequest struct {
	SymbolName  string `json:"symbol_name,omitempty"`
	PatternName string `json:"pattern_name,omitempty"`
	AddonFilter string `json:"addon_filter,omitempty"`
	Limit       int    `json:"limit,omitempty"`
}

type ExplainConfidenceRequest struct {
	SymbolName string `json:"symbol_name"`
}

type ExplainSymbolUsageRequest struct {
	SymbolName      string `json:"symbol_name"`
	IncludePatterns bool   `json:"include_patterns,omitempty"`
	IncludePitfalls bool   `json:"include_pitfalls,omitempty"`
}

type ScaffoldAddonSnippetRequest struct {
	TaskDescription string   `json:"task_description"`
	DesiredAPIs     []string `json:"desired_apis,omitempty"`
	OutputStyle     string   `json:"output_style,omitempty"`
	IncludeXML      bool     `json:"include_xml,omitempty"`
	IncludeComments bool     `json:"include_comments,omitempty"`
}

type IngestObservationRequest struct {
	Observation map[string]any `json:"observation"`
	DryRun      bool           `json:"dry_run,omitempty"`
	Persist     bool           `json:"persist,omitempty"`
	QueueFile   string         `json:"queue_file,omitempty"`
	SourcePath  string         `json:"source_path,omitempty"`
}

type IngestObservationBatchRequest struct {
	DryRun          bool   `json:"dry_run,omitempty"`
	Persist         bool   `json:"persist,omitempty"`
	AutoAccept      bool   `json:"auto_accept,omitempty"`
	QueueFile       string `json:"queue_file,omitempty"`
	ContinueOnError bool   `json:"continue_on_error,omitempty"`
	Limit           int    `json:"limit,omitempty"`
}

type ListPendingObservationsRequest struct {
	StatusFilter string `json:"status_filter,omitempty"` // candidate, accepted, or empty for both
	TargetFilter string `json:"target_filter,omitempty"` // match against target_seeds entries
	Limit        int    `json:"limit,omitempty"`
}

type ReviewObservationRequest struct {
	ObservationID string `json:"observation_id"`
	Verdict       string `json:"verdict"` // "accept" or "reject"
	Reviewer      string `json:"reviewer,omitempty"`
	Notes         string `json:"notes,omitempty"`
}

type PromoteObservationRequest struct {
	ObservationID    string `json:"observation_id"`
	DryRun           bool   `json:"dry_run,omitempty"`
	SeedPathOverride string `json:"seed_path_override,omitempty"`
}

type ListRejectedObservationsRequest struct {
	Limit int `json:"limit,omitempty"`
}

type RegenerateRequest struct {
	Scope  string `json:"scope,omitempty"` // "platform", "site", or "full" (default)
	DryRun bool   `json:"dry_run,omitempty"`
}

type PromoteAllAcceptedRequest struct {
	DryRun bool `json:"dry_run,omitempty"`
}

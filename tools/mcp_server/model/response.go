package model

// ResourceEnvelope is the stable resource response shape.
type ResourceEnvelope struct {
	URI            string                 `json:"uri"`
	Kind           string                 `json:"kind"`
	Title          string                 `json:"title"`
	Summary        string                 `json:"summary"`
	ContentType    string                 `json:"content_type"`
	StructuredData interface{}            `json:"structured_data"`
	SourcePaths    []string               `json:"source_paths"`
	Warnings       []Warning              `json:"warnings"`
	Metadata       map[string]interface{} `json:"metadata"`
}

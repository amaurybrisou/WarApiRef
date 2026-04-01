package server

import "roraddons/tools/mcp_server/schema"

type ToolDescriptor struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	InputSchema    any      `json:"inputSchema"`
	RequestFields  []string `json:"request_fields"`
	ResponseFields []string `json:"response_fields"`
}

type ResourceDescriptor struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

func ToolRegistry() []ToolDescriptor {
	return []ToolDescriptor{
		{Name: "lookup_symbol", Description: "Exact/normalized/fuzzy symbol lookup with explicit match status.", InputSchema: map[string]any{"type": "object", "additionalProperties": false, "properties": map[string]any{"symbol_name": map[string]any{"type": "string", "minLength": 1}, "symbol_type": map[string]any{"type": "string"}, "include_examples": map[string]any{"type": "boolean"}, "include_related": map[string]any{"type": "boolean"}}, "required": []string{"symbol_name"}}, RequestFields: []string{"symbol_name", "symbol_type", "include_examples", "include_related"}, ResponseFields: []string{"exact_match", "symbol", "alternatives", "warnings"}},
		{Name: "search_symbols", Description: "Search canonical symbols with filters and confidence bounds.", InputSchema: map[string]any{"type": "object", "additionalProperties": false, "properties": map[string]any{"query": map[string]any{"type": "string", "minLength": 1}, "type_filter": map[string]any{"type": "string", "enum": schema.SupportedSymbolTypes}, "namespace_filter": map[string]any{"type": "string"}, "confidence_min": map[string]any{"type": "integer", "minimum": 0, "maximum": 100}, "limit": map[string]any{"type": "integer", "minimum": 1, "maximum": 100}, "include_examples": map[string]any{"type": "boolean"}}, "required": []string{"query"}}, RequestFields: []string{"query", "type_filter", "namespace_filter", "confidence_min", "limit", "include_examples"}, ResponseFields: []string{"total_matches", "matches", "warnings"}},
		{Name: "get_related_symbols", Description: "Traverse related symbols grouped by relation type.", InputSchema: map[string]any{"type": "object", "additionalProperties": false, "properties": map[string]any{"symbol_name": map[string]any{"type": "string", "minLength": 1}, "relation_types": map[string]any{"type": "array", "items": map[string]any{"type": "string", "enum": schema.SupportedRelationTypes}}, "depth": map[string]any{"type": "integer", "minimum": 0, "maximum": 3}, "limit_per_relation": map[string]any{"type": "integer", "minimum": 1, "maximum": 50}}, "required": []string{"symbol_name"}}, RequestFields: []string{"symbol_name", "relation_types", "depth", "limit_per_relation"}, ResponseFields: []string{"source_symbol", "relations", "graph_summary", "warnings"}},
		{Name: "get_event_flow", Description: "Return observed handler and call flows for an event.", InputSchema: map[string]any{"type": "object", "additionalProperties": false, "properties": map[string]any{"event_name": map[string]any{"type": "string", "minLength": 1}, "include_examples": map[string]any{"type": "boolean"}, "include_downstream": map[string]any{"type": "boolean"}}, "required": []string{"event_name"}}, RequestFields: []string{"event_name", "include_examples", "include_downstream"}, ResponseFields: []string{"event", "observed_handlers", "common_flows", "downstream_symbols", "examples", "warnings"}},
		{Name: "get_xml_handler_info", Description: "Return XML handler role, callback shape, and pairings.", InputSchema: map[string]any{"type": "object", "additionalProperties": false, "properties": map[string]any{"handler_name": map[string]any{"type": "string", "minLength": 1}, "include_examples": map[string]any{"type": "boolean"}, "include_related": map[string]any{"type": "boolean"}}, "required": []string{"handler_name"}}, RequestFields: []string{"handler_name", "include_examples", "include_related"}, ResponseFields: []string{"handler", "lifecycle_role", "expected_callback_shape", "common_element_types", "paired_symbols", "examples", "warnings"}},
		{Name: "find_usage_examples", Description: "Find symbol or pattern usage examples.", InputSchema: map[string]any{"type": "object", "additionalProperties": false, "properties": map[string]any{"symbol_name": map[string]any{"type": "string"}, "pattern_name": map[string]any{"type": "string"}, "addon_filter": map[string]any{"type": "string"}, "limit": map[string]any{"type": "integer", "minimum": 1, "maximum": 100}}, "anyOf": []map[string]any{{"required": []string{"symbol_name"}}, {"required": []string{"pattern_name"}}}}, RequestFields: []string{"symbol_name", "pattern_name", "addon_filter", "limit"}, ResponseFields: []string{"query_type", "total_matches", "examples"}},
		{Name: "explain_confidence", Description: "Explain confidence and semantic confidence evidence.", InputSchema: map[string]any{"type": "object", "additionalProperties": false, "properties": map[string]any{"symbol_name": map[string]any{"type": "string", "minLength": 1}}, "required": []string{"symbol_name"}}, RequestFields: []string{"symbol_name"}, ResponseFields: []string{"symbol", "confidence_breakdown", "semantic_confidence_breakdown", "evidence_signals", "penalties", "rationale", "warnings"}},
		{Name: "explain_symbol_usage", Description: "Explain purpose, parameters, returns, and uncertainty.", InputSchema: map[string]any{"type": "object", "additionalProperties": false, "properties": map[string]any{"symbol_name": map[string]any{"type": "string", "minLength": 1}, "include_patterns": map[string]any{"type": "boolean"}, "include_pitfalls": map[string]any{"type": "boolean"}}, "required": []string{"symbol_name"}}, RequestFields: []string{"symbol_name", "include_patterns", "include_pitfalls"}, ResponseFields: []string{"symbol", "purpose", "usage_context", "parameters", "returns", "side_effects", "preconditions", "postconditions", "common_patterns", "pitfalls_or_uncertainty", "examples", "warnings"}},
		{Name: "scaffold_addon_snippet", Description: "Generate snippets using documented APIs only.", InputSchema: map[string]any{"type": "object", "additionalProperties": false, "properties": map[string]any{"task_description": map[string]any{"type": "string", "minLength": 1}, "desired_apis": map[string]any{"type": "array", "items": map[string]any{"type": "string"}}, "output_style": map[string]any{"type": "string"}, "include_xml": map[string]any{"type": "boolean"}, "include_comments": map[string]any{"type": "boolean"}}, "required": []string{"task_description"}}, RequestFields: []string{"task_description", "desired_apis", "output_style", "include_xml", "include_comments"}, ResponseFields: []string{"generated_artifacts", "used_documented_symbols", "avoided_symbols", "uncertainty_notes", "rationale"}},
	}
}

func ResourceRegistry() []ResourceDescriptor {
	out := []ResourceDescriptor{}
	for _, uri := range schema.ResourceURIs {
		out = append(out, ResourceDescriptor{Name: uri, URI: uri})
	}
	return out
}

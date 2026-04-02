package schema

var ToolNames = []string{
	"lookup_symbol",
	"search_symbols",
	"get_related_symbols",
	"get_event_flow",
	"get_xml_handler_info",
	"find_usage_examples",
	"explain_confidence",
	"explain_symbol_usage",
	"scaffold_addon_snippet",
	"ingest_observation",
	"ingest_observation_batch",
	"list_pending_observations",
	"review_observation",
	"promote_observation",
	"list_rejected_observations",
	"regenerate_from_promoted_knowledge",
}

var SupportedSymbolTypes = []string{"function", "event", "xml_handler", "data_structure", "xml_element"}

var SupportedRelationTypes = []string{"related", "used_with", "triggered_by", "affects", "calls", "handled_by", "commonly_used_with"}

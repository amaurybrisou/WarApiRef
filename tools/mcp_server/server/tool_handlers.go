package server

import (
	"encoding/json"

	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
	toolops "roraddons/tools/mcp_server/tools"
)

func (a *App) HandleToolCall(name string, args json.RawMessage) (interface{}, *model.APIError) {
	decoded, verr := a.validator.decodeAndValidate(name, args)
	if verr != nil {
		return nil, verr
	}
	switch name {
	case "lookup_symbol":
		return toolops.LookupSymbol(a.store, decoded.(schema.LookupSymbolRequest))
	case "search_symbols":
		return toolops.SearchSymbols(a.store, decoded.(schema.SearchSymbolsRequest)), nil
	case "get_related_symbols":
		return toolops.GetRelatedSymbols(a.store, decoded.(schema.GetRelatedSymbolsRequest))
	case "get_event_flow":
		return toolops.GetEventFlow(a.store, decoded.(schema.GetEventFlowRequest))
	case "get_xml_handler_info":
		return toolops.GetXMLHandlerInfo(a.store, decoded.(schema.GetXMLHandlerInfoRequest))
	case "find_usage_examples":
		return toolops.FindUsageExamples(a.store, decoded.(schema.FindUsageExamplesRequest)), nil
	case "explain_confidence":
		return toolops.ExplainConfidence(a.store, decoded.(schema.ExplainConfidenceRequest)), nil
	case "explain_symbol_usage":
		return toolops.ExplainSymbolUsage(a.store, decoded.(schema.ExplainSymbolUsageRequest)), nil
	case "scaffold_addon_snippet":
		return toolops.ScaffoldAddonSnippet(a.store, decoded.(schema.ScaffoldAddonSnippetRequest)), nil
	default:
		return nil, &model.APIError{ErrorCode: "unknown_tool", ErrorMessage: "unknown tool"}
	}
}

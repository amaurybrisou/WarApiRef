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
	if name == "ingest_observation" {
		return a.ingestObservation(decoded.(schema.IngestObservationRequest)), nil
	}
	if name == "ingest_observation_batch" {
		return a.ingestObservationBatch(decoded.(schema.IngestObservationBatchRequest)), nil
	}

	store, storeErr := a.storeReady()
	if storeErr != nil {
		return nil, storeErr
	}
	switch name {
	case "lookup_symbol":
		return toolops.LookupSymbol(store, decoded.(schema.LookupSymbolRequest))
	case "search_symbols":
		return toolops.SearchSymbols(store, decoded.(schema.SearchSymbolsRequest)), nil
	case "get_related_symbols":
		return toolops.GetRelatedSymbols(store, decoded.(schema.GetRelatedSymbolsRequest))
	case "get_event_flow":
		return toolops.GetEventFlow(store, decoded.(schema.GetEventFlowRequest))
	case "get_xml_handler_info":
		return toolops.GetXMLHandlerInfo(store, decoded.(schema.GetXMLHandlerInfoRequest))
	case "find_usage_examples":
		return toolops.FindUsageExamples(store, decoded.(schema.FindUsageExamplesRequest)), nil
	case "explain_confidence":
		return toolops.ExplainConfidence(store, decoded.(schema.ExplainConfidenceRequest)), nil
	case "explain_symbol_usage":
		return toolops.ExplainSymbolUsage(store, decoded.(schema.ExplainSymbolUsageRequest)), nil
	case "scaffold_addon_snippet":
		return toolops.ScaffoldAddonSnippet(store, decoded.(schema.ScaffoldAddonSnippetRequest)), nil
	case "ingest_observation":
		return a.ingestObservation(decoded.(schema.IngestObservationRequest)), nil
	case "ingest_observation_batch":
		return a.ingestObservationBatch(decoded.(schema.IngestObservationBatchRequest)), nil
	default:
		return nil, &model.APIError{ErrorCode: "unknown_tool", ErrorMessage: "unknown tool"}
	}
}

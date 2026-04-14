package server

import (
	"encoding/json"
	"strings"

	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
)

type Validator struct{}

func NewValidator() *Validator { return &Validator{} }

func (v *Validator) decodeAndValidate(toolName string, raw json.RawMessage) (interface{}, *model.APIError) {
	switch toolName {
	case "lookup_symbol":
		var req schema.LookupSymbolRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("lookup_symbol", "invalid request payload")
		}
		if strings.TrimSpace(req.SymbolName) == "" {
			return nil, invalidParams("lookup_symbol", "symbol_name is required")
		}
		return req, nil
	case "search_symbols":
		var req schema.SearchSymbolsRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("search_symbols", "invalid request payload")
		}
		if strings.TrimSpace(req.Query) == "" {
			return nil, invalidParams("search_symbols", "query is required")
		}
		if req.TypeFilter != "" && !contains(schema.SupportedSymbolTypes, req.TypeFilter) {
			return nil, invalidParams("search_symbols", "unsupported type_filter")
		}
		return req, nil
	case "get_related_symbols":
		var req schema.GetRelatedSymbolsRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("get_related_symbols", "invalid request payload")
		}
		if strings.TrimSpace(req.SymbolName) == "" {
			return nil, invalidParams("get_related_symbols", "symbol_name is required")
		}
		for _, r := range req.RelationTypes {
			if !contains(schema.SupportedRelationTypes, r) {
				return nil, invalidParams("get_related_symbols", "invalid relation type: "+r)
			}
		}
		return req, nil
	case "get_event_flow":
		var req schema.GetEventFlowRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("get_event_flow", "invalid request payload")
		}
		if strings.TrimSpace(req.EventName) == "" {
			return nil, invalidParams("get_event_flow", "event_name is required")
		}
		return req, nil
	case "get_xml_handler_info":
		var req schema.GetXMLHandlerInfoRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("get_xml_handler_info", "invalid request payload")
		}
		if strings.TrimSpace(req.HandlerName) == "" {
			return nil, invalidParams("get_xml_handler_info", "handler_name is required")
		}
		return req, nil
	case "find_usage_examples":
		var req schema.FindUsageExamplesRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("find_usage_examples", "invalid request payload")
		}
		if strings.TrimSpace(req.SymbolName) == "" && strings.TrimSpace(req.PatternName) == "" {
			return nil, invalidParams("find_usage_examples", "at least one of symbol_name or pattern_name is required")
		}
		return req, nil
	case "explain_confidence":
		var req schema.ExplainConfidenceRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("explain_confidence", "invalid request payload")
		}
		if strings.TrimSpace(req.SymbolName) == "" {
			return nil, invalidParams("explain_confidence", "symbol_name is required")
		}
		return req, nil
	case "explain_symbol_usage":
		var req schema.ExplainSymbolUsageRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("explain_symbol_usage", "invalid request payload")
		}
		if strings.TrimSpace(req.SymbolName) == "" {
			return nil, invalidParams("explain_symbol_usage", "symbol_name is required")
		}
		return req, nil
	case "scaffold_addon_snippet":
		var req schema.ScaffoldAddonSnippetRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("scaffold_addon_snippet", "invalid request payload")
		}
		if strings.TrimSpace(req.TaskDescription) == "" {
			return nil, invalidParams("scaffold_addon_snippet", "task_description is required")
		}
		return req, nil
	case "ingest_observation":
		var req schema.IngestObservationRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("ingest_observation", "invalid request payload")
		}
		if len(req.Observation) == 0 {
			return nil, invalidParams("ingest_observation", "observation is required")
		}
		return req, nil
	case "ingest_observation_batch":
		var req schema.IngestObservationBatchRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("ingest_observation_batch", "invalid request payload")
		}
		if req.Limit < 0 {
			return nil, invalidParams("ingest_observation_batch", "limit must be >= 0")
		}
		return req, nil
	case "list_pending_observations":
		var req schema.ListPendingObservationsRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("list_pending_observations", "invalid request payload")
		}
		if req.Limit < 0 {
			return nil, invalidParams("list_pending_observations", "limit must be >= 0")
		}
		return req, nil
	case "review_observation":
		var req schema.ReviewObservationRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("review_observation", "invalid request payload")
		}
		if strings.TrimSpace(req.ObservationID) == "" {
			return nil, invalidParams("review_observation", "observation_id is required")
		}
		if req.Verdict != "accept" && req.Verdict != "reject" {
			return nil, invalidParams("review_observation", "verdict must be 'accept' or 'reject'")
		}
		return req, nil
	case "promote_observation":
		var req schema.PromoteObservationRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("promote_observation", "invalid request payload")
		}
		if strings.TrimSpace(req.ObservationID) == "" {
			return nil, invalidParams("promote_observation", "observation_id is required")
		}
		return req, nil
	case "promote_all_accepted":
		var req schema.PromoteAllAcceptedRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("promote_all_accepted", "invalid request payload")
		}
		return req, nil
	case "list_rejected_observations":
		var req schema.ListRejectedObservationsRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("list_rejected_observations", "invalid request payload")
		}
		if req.Limit < 0 {
			return nil, invalidParams("list_rejected_observations", "limit must be >= 0")
		}
		return req, nil
	case "regenerate_from_promoted_knowledge":
		var req schema.RegenerateRequest
		if err := json.Unmarshal(raw, &req); err != nil {
			return nil, invalidParams("regenerate_from_promoted_knowledge", "invalid request payload")
		}
		if req.Scope != "" && req.Scope != "platform" && req.Scope != "site" && req.Scope != "full" {
			return nil, invalidParams("regenerate_from_promoted_knowledge", "scope must be 'platform', 'site', or 'full'")
		}
		return req, nil
	default:
		return nil, &model.APIError{ErrorCode: "unknown_tool", ErrorMessage: "unknown tool: " + toolName}
	}
}

func invalidParams(tool, msg string) *model.APIError {
	return &model.APIError{ErrorCode: "invalid_params", ErrorMessage: msg, Details: map[string]any{"tool": tool}}
}

func contains(list []string, v string) bool {
	for _, e := range list {
		if strings.EqualFold(e, v) {
			return true
		}
	}
	return false
}

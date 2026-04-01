package tools

import (
	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
)

func GetEventFlow(b Backend, req schema.GetEventFlowRequest) (model.EventFlow, *model.APIError) {
	flow := b.EventFlow(req.EventName, req.IncludeExamples, req.IncludeDownstream)
	if flow.Event.ID == "" {
		return model.EventFlow{}, &model.APIError{ErrorCode: "event_not_found", ErrorMessage: "Event symbol was not found."}
	}
	return flow, nil
}

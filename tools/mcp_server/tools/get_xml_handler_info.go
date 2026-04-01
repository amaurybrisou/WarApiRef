package tools

import (
	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
)

func GetXMLHandlerInfo(b Backend, req schema.GetXMLHandlerInfoRequest) (model.XMLHandlerInfo, *model.APIError) {
	info := b.XMLHandlerInfo(req.HandlerName, req.IncludeExamples, req.IncludeRelated)
	if info.Handler.ID == "" {
		return model.XMLHandlerInfo{}, &model.APIError{ErrorCode: "xml_handler_not_found", ErrorMessage: "XML handler was not found."}
	}
	return info, nil
}

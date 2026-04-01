package server

import (
	"strings"

	"roraddons/tools/mcp_server/model"
	resourceops "roraddons/tools/mcp_server/resources"
)

func (a *App) HandleResourceRead(uri string) (model.ResourceEnvelope, *model.APIError) {
	switch {
	case strings.HasPrefix(uri, "war-api://symbols/"):
		return resourceops.ResolveSymbolResource(a.store, uri)
	case strings.HasPrefix(uri, "war-api://events/"):
		return resourceops.ResolveEventResource(a.store, uri)
	case strings.HasPrefix(uri, "war-api://xml-handlers/"):
		return resourceops.ResolveHandlerResource(a.store, uri)
	case strings.HasPrefix(uri, "war-api://patterns/"):
		return resourceops.ResolvePatternResource(a.store, uri)
	case strings.HasPrefix(uri, "war-api://graph/"):
		return resourceops.ResolveGraphResource(a.store, uri)
	case strings.HasPrefix(uri, "war-api://meta/") || strings.HasPrefix(uri, "war-api://navigation/"):
		return resourceops.ResolveMetaResource(a.store, uri)
	default:
		return model.ResourceEnvelope{}, &model.APIError{ErrorCode: "unsupported_uri", ErrorMessage: "unsupported resource URI"}
	}
}

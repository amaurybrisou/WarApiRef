package server

import (
	"strings"

	"roraddons/tools/mcp_server/model"
	resourceops "roraddons/tools/mcp_server/resources"
)

func (a *App) HandleResourceRead(uri string) (model.ResourceEnvelope, *model.APIError) {
	store, storeErr := a.storeReady()
	if storeErr != nil {
		return model.ResourceEnvelope{}, storeErr
	}
	switch {
	case strings.HasPrefix(uri, "war-api://symbols/"):
		return resourceops.ResolveSymbolResource(store, uri)
	case strings.HasPrefix(uri, "war-api://events/"):
		return resourceops.ResolveEventResource(store, uri)
	case strings.HasPrefix(uri, "war-api://xml-handlers/"):
		return resourceops.ResolveHandlerResource(store, uri)
	case strings.HasPrefix(uri, "war-api://patterns/"):
		return resourceops.ResolvePatternResource(store, uri)
	case strings.HasPrefix(uri, "war-api://graph/"):
		return resourceops.ResolveGraphResource(store, uri)
	case strings.HasPrefix(uri, "war-api://meta/") || strings.HasPrefix(uri, "war-api://navigation/"):
		return resourceops.ResolveMetaResource(store, uri)
	default:
		return model.ResourceEnvelope{}, &model.APIError{ErrorCode: "unsupported_uri", ErrorMessage: "unsupported resource URI"}
	}
}

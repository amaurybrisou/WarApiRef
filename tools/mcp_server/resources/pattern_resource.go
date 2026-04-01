package resources

import "roraddons/tools/mcp_server/model"

func ResolvePatternResource(b Backend, uri string) (model.ResourceEnvelope, *model.APIError) {
	return b.ResolveResource(uri)
}

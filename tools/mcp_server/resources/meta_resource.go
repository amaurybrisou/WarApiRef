package resources

import "roraddons/tools/mcp_server/model"

func ResolveMetaResource(b Backend, uri string) (model.ResourceEnvelope, *model.APIError) {
	return b.ResolveResource(uri)
}

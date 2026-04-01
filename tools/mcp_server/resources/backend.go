package resources

import "roraddons/tools/mcp_server/model"

type Backend interface {
	ResolveResource(uri string) (model.ResourceEnvelope, *model.APIError)
}

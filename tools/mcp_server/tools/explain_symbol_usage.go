package tools

import "roraddons/tools/mcp_server/schema"

func ExplainSymbolUsage(b Backend, req schema.ExplainSymbolUsageRequest) map[string]any {
	sym, body := b.ExplainUsage(req.SymbolName)
	body["symbol"] = sym
	body["warnings"] = sym.Warnings
	return body
}

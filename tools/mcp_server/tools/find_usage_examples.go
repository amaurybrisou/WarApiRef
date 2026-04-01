package tools

import "roraddons/tools/mcp_server/schema"

func FindUsageExamples(b Backend, req schema.FindUsageExamplesRequest) schema.FindUsageExamplesResponse {
	examples := b.UsageExamples(req.SymbolName, req.PatternName, req.AddonFilter, req.Limit)
	queryType := "symbol"
	if req.PatternName != "" {
		queryType = "pattern"
	}
	return schema.FindUsageExamplesResponse{QueryType: queryType, TotalMatches: len(examples), Examples: examples}
}

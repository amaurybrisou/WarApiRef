package tools

import "roraddons/tools/mcp_server/schema"

func SearchSymbols(b Backend, req schema.SearchSymbolsRequest) schema.SearchSymbolsResponse {
	limit := req.Limit
	if limit <= 0 {
		limit = 20
	}
	matches := b.Search(req.Query, req.TypeFilter, req.NamespaceFilter, req.ConfidenceMin, limit)
	out := make([]schema.SearchMatch, 0, len(matches))
	for _, m := range matches {
		sm := schema.SearchMatch{
			ID:                      m.ID,
			CanonicalName:           m.CanonicalName,
			Type:                    m.Type,
			Namespace:               m.Namespace,
			Summary:                 m.Summary,
			ConfidenceScore:         m.ConfidenceScore,
			SemanticConfidenceScore: m.SemanticConfidenceScore,
			DocPath:                 m.DocPath,
		}
		out = append(out, sm)
	}
	return schema.SearchSymbolsResponse{TotalMatches: len(out), Matches: out}
}

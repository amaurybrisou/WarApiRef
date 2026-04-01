package tools

import (
	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
)

func GetRelatedSymbols(b Backend, req schema.GetRelatedSymbolsRequest) (schema.GetRelatedSymbolsResponse, *model.APIError) {
	_, found, source, suggestions := b.LookupSymbol(req.SymbolName, "")
	if !found {
		return schema.GetRelatedSymbolsResponse{}, &model.APIError{ErrorCode: "symbol_not_found", ErrorMessage: "Source symbol not found.", CandidateMatches: suggestions}
	}
	rel := b.Related(source, req.RelationTypes, req.Depth, req.LimitPerRelation)
	grouped := map[string][]schema.RelatedItem{}
	total := 0
	for rt, items := range rel {
		for _, it := range items {
			grouped[rt] = append(grouped[rt], schema.RelatedItem{ID: it.ID, CanonicalName: it.CanonicalName, Type: it.Type, RelationType: rt, Summary: it.Summary, ConfidenceScore: it.ConfidenceScore, DocPath: it.DocPath})
			total++
		}
	}
	resp := schema.GetRelatedSymbolsResponse{SourceSymbol: source, Relations: grouped, GraphSummary: map[string]any{"relation_groups": len(grouped), "related_total": total}}
	return resp, nil
}

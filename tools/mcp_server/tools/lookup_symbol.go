package tools

import (
	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
)

func LookupSymbol(b Backend, req schema.LookupSymbolRequest) (schema.LookupSymbolResponse, *model.APIError) {
	exact, found, sym, suggestions := b.LookupSymbol(req.SymbolName, req.SymbolType)
	if !found {
		return schema.LookupSymbolResponse{}, &model.APIError{
			ErrorCode:        "symbol_not_found",
			ErrorMessage:     "No symbol matched the requested name.",
			Suggestion:       "Use search_symbols to discover canonical names.",
			CandidateMatches: suggestions,
		}
	}
	if !req.IncludeExamples {
		sym.Examples = nil
	}
	if !req.IncludeRelated {
		sym.RelatedIDs = nil
	}
	resp := schema.LookupSymbolResponse{ExactMatch: exact, Symbol: sym}
	if !exact {
		resp.Warnings = append(resp.Warnings, model.Warning{Code: "inferred_match", Message: "Resolved by normalized name match, not exact canonical name."})
	}
	resp.Alternatives = suggestions
	return resp, nil
}

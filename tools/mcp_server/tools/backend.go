package tools

import "roraddons/tools/mcp_server/model"

type Backend interface {
	LookupSymbol(name, symbolType string) (exact bool, found bool, symbol model.SymbolSummary, suggestions []model.CandidateMatch)
	Search(query, symbolType, namespace string, confidenceMin, limit int) []model.SymbolSummary
	Related(source model.SymbolSummary, relationTypes []string, depth, limit int) map[string][]model.SymbolSummary
	EventFlow(eventName string, includeExamples, includeDownstream bool) model.EventFlow
	XMLHandlerInfo(handlerName string, includeExamples, includeRelated bool) model.XMLHandlerInfo
	UsageExamples(symbolName, patternName, addonFilter string, limit int) []model.UsageExample
	ExplainConfidence(symbolName string) (model.SymbolSummary, map[string]any, map[string]any, []string, []string, string)
	ExplainUsage(symbolName string) (model.SymbolSummary, map[string]any)
}

package tools

import "roraddons/tools/mcp_server/schema"

func ExplainConfidence(b Backend, req schema.ExplainConfidenceRequest) map[string]any {
	sym, cb, scb, evidence, penalties, rationale := b.ExplainConfidence(req.SymbolName)
	return map[string]any{
		"symbol":                        sym,
		"confidence_breakdown":          cb,
		"semantic_confidence_breakdown": scb,
		"evidence_signals":              evidence,
		"penalties":                     penalties,
		"rationale":                     rationale,
		"warnings":                      sym.Warnings,
	}
}

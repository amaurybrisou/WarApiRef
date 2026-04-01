package deep_analysis

// GraphIntegration bridges the deep_analysis package with the existing platform/semantic.go
// This package helps apply deep analysis findings to enrich graph edges with confidence scores
// and richer evidence tracking.

import (
	"fmt"
	"strings"
)

// EnrichedEdge represents an edge enriched with deep analysis findings
type EnrichedEdge struct {
	From            string
	To              string
	Type            string
	ConfidenceScore int
	EvidenceCount   int
	EvidenceSources []string
	Rationale       string
	AnalysisMethod  string                 // how this edge was inferred
	RawEvidence     map[string]interface{} // detailed evidence for debugging
}

// EdgeEnricher applies deep analysis to enrich existing edges
type EdgeEnricher struct {
	EdgeInference     *EdgeInference
	ReturnInference   *ReturnInference
	ArgumentInference *ArgumentInference
	UsageClustering   *UsageClustering
	ConfidenceScorer  *ConfidenceScorer
}

// NewEdgeEnricher creates enricher with all analysis engines
func NewEdgeEnricher() *EdgeEnricher {
	return &EdgeEnricher{
		EdgeInference:     NewEdgeInference(),
		ReturnInference:   NewReturnInference(),
		ArgumentInference: NewArgumentInference(),
		UsageClustering:   NewUsageClustering(),
		ConfidenceScorer:  NewConfidenceScorer(),
	}
}

// EnrichEdge takes a basic edge and enriches it with confidence and evidence
func (ee *EdgeEnricher) EnrichEdge(from, to, edgeType string, baseEvidence []string) EnrichedEdge {
	enriched := EnrichedEdge{
		From:            from,
		To:              to,
		Type:            edgeType,
		RawEvidence:     make(map[string]interface{}),
		EvidenceSources: baseEvidence,
		EvidenceCount:   len(baseEvidence),
	}

	// Calculate confidence based on edge type and evidence
	confidence := ee.calculateEdgeConfidence(from, to, edgeType, baseEvidence)
	enriched.ConfidenceScore = confidence

	// Generate rationale
	enriched.Rationale = ee.generateRationale(from, to, edgeType, baseEvidence, confidence)
	enriched.AnalysisMethod = "evidence_based_inference"

	return enriched
}

// calculateEdgeConfidence determines confidence for an edge
func (ee *EdgeEnricher) calculateEdgeConfidence(from, to, edgeType string, sources []string) int {
	baseScore := 50 // default base

	// Type-specific scoring
	switch edgeType {
	case "calls":
		// Direct function calls are high confidence
		baseScore = 85 + len(sources)*3
	case "handled_by":
		// Event handlers are well-defined
		baseScore = 80 + len(sources)*3
	case "triggered_by":
		// Event triggers are fairly definite
		baseScore = 80 + len(sources)*3
	case "defined_in":
		// Definitions are nearly certain
		baseScore = 95
	case "reads_systemdata":
		// Data access from static analysis
		baseScore = 70 + len(sources)*4
	case "reads_gamedata":
		// Data access from static analysis
		baseScore = 70 + len(sources)*4
	case "writes_state":
		// State modification detection
		baseScore = 65 + len(sources)*4
	case "updates_ui":
		// UI operations from pattern matching
		baseScore = 68 + len(sources)*3
	case "toggles_visibility":
		// Show/Hide patterns
		baseScore = 72 + len(sources)*3
	case "updates_layout":
		// Layout modification patterns
		baseScore = 70 + len(sources)*3
	case "initializes":
		// Init-phase detection
		baseScore = 78 + len(sources)*2
	case "refreshes":
		// Update-phase detection
		baseScore = 75 + len(sources)*2
	case "commonly_used_with":
		// Co-occurrence patterns (weaker signal)
		baseScore = 55 + len(sources)*2
	case "appears_in_same_flow":
		// Sequence patterns
		baseScore = 60 + len(sources)*3
	default:
		baseScore = 50
	}

	// Evidence multiplier: more sources = higher confidence
	if len(sources) > 5 {
		baseScore += 10
	}
	if len(sources) > 10 {
		baseScore += 5
	}

	// Cap at reasonable maximum based on type
	switch edgeType {
	case "calls", "handled_by", "triggered_by", "defined_in":
		if baseScore > 95 {
			baseScore = 95
		}
	case "reads_systemdata", "reads_gamedata", "initializes", "refreshes":
		if baseScore > 85 {
			baseScore = 85
		}
	case "updates_ui", "toggles_visibility", "updates_layout", "writes_state":
		if baseScore > 80 {
			baseScore = 80
		}
	case "commonly_used_with", "appears_in_same_flow":
		if baseScore > 75 {
			baseScore = 75
		}
	}

	return baseScore
}

// generateRationale creates human-readable explanation for confidence
func (ee *EdgeEnricher) generateRationale(from, to, edgeType string, sources []string, confidence int) string {
	var rationale string

	switch edgeType {
	case "calls":
		rationale = fmt.Sprintf("%s calls %s (observed %d times)", from, to, len(sources))
	case "handled_by":
		rationale = fmt.Sprintf("%s is handled by %s (event handler binding)", from, to)
	case "triggered_by":
		rationale = fmt.Sprintf("%s triggers %s (event registration)", from, to)
	case "defined_in":
		rationale = fmt.Sprintf("%s is defined in module %s", from, to)
	case "reads_systemdata":
		rationale = fmt.Sprintf("%s reads from SystemData (%d access points)", from, len(sources))
	case "reads_gamedata":
		rationale = fmt.Sprintf("%s reads from GameData (%d access points)", from, len(sources))
	case "writes_state":
		rationale = fmt.Sprintf("%s modifies state (%d modification points)", from, len(sources))
	case "updates_ui":
		rationale = fmt.Sprintf("%s updates UI elements (%d update patterns)", from, len(sources))
	case "toggles_visibility":
		rationale = fmt.Sprintf("%s toggles visibility of UI elements (%d patterns)", from, len(sources))
	case "updates_layout":
		rationale = fmt.Sprintf("%s modifies layout (position/size) (%d patterns)", from, len(sources))
	case "initializes":
		rationale = fmt.Sprintf("%s is called during initialization phase", from)
	case "refreshes":
		rationale = fmt.Sprintf("%s is called during update/refresh phase", from)
	case "commonly_used_with":
		rationale = fmt.Sprintf("%s and %s are frequently used together (%d co-occurrences)", from, to, len(sources))
	case "appears_in_same_flow":
		rationale = fmt.Sprintf("%s and %s appear in related event/execution flows", from, to)
	default:
		rationale = fmt.Sprintf("%s relates to %s via %s", from, to, edgeType)
	}

	if confidence >= 80 {
		rationale += " [high confidence]"
	} else if confidence >= 60 {
		rationale += " [medium confidence]"
	} else if confidence >= 40 {
		rationale += " [low confidence]"
	}

	return rationale
}

// BuildMissingEdges uses deep analysis to discover relationships missed by initial graph generation
func (ee *EdgeEnricher) BuildMissingEdges() []*EnrichedEdge {
	edges := []*EnrichedEdge{}

	// Find data access relationships
	for functionPath, fa := range ee.EdgeInference.AllFunctions {
		if ee.EdgeInference.InferReadsSystemData(functionPath) {
			edge := ee.EnrichEdge(functionPath, "SystemData", "reads_systemdata",
				[]string{"static_analysis:data_access"})
			edges = append(edges, &edge)
		}
		if ee.EdgeInference.InferReadsGameData(functionPath) {
			edge := ee.EnrichEdge(functionPath, "GameData", "reads_gamedata",
				[]string{"static_analysis:data_access"})
			edges = append(edges, &edge)
		}
		if ee.EdgeInference.InferUpdatesUI(functionPath) {
			edge := ee.EnrichEdge(functionPath, "UI", "updates_ui",
				[]string{"static_analysis:ui_operations"})
			edges = append(edges, &edge)
		}
		if ee.EdgeInference.InferTogglesVisibility(functionPath) {
			edge := ee.EnrichEdge(functionPath, "Visibility", "toggles_visibility",
				[]string{"static_analysis:visibility_operations"})
			edges = append(edges, &edge)
		}
		if ee.EdgeInference.InferUpdatesLayout(functionPath) {
			edge := ee.EnrichEdge(functionPath, "Layout", "updates_layout",
				[]string{"static_analysis:layout_operations"})
			edges = append(edges, &edge)
		}

		// Init/refresh detection
		if fa.IsInitFunction {
			edge := ee.EnrichEdge(functionPath, "InitPhase", "initializes",
				[]string{"semantic:init_detection"})
			edges = append(edges, &edge)
		}
		if fa.IsUpdateFunction {
			edge := ee.EnrichEdge(functionPath, "UpdatePhase", "refreshes",
				[]string{"semantic:update_detection"})
			edges = append(edges, &edge)
		}
	}

	// Find commonly-used combinations
	for functionPath, fa := range ee.EdgeInference.AllFunctions {
		for _, callSite := range fa.CallsSites {
			calledFunc := callSite.CalledFunction
			otherPath, found := ee.findRelatedFunction(calledFunc)
			if found && ee.EdgeInference.InferCommonlyUsedWith(functionPath, otherPath) {
				edge := ee.EnrichEdge(functionPath, otherPath, "commonly_used_with",
					[]string{"static_analysis:call_pattern"})
				edges = append(edges, &edge)
			}
		}
	}

	return edges
}

// findRelatedFunction locates a function by name in the analysis
func (ee *EdgeEnricher) findRelatedFunction(funcName string) (string, bool) {
	for path := range ee.EdgeInference.AllFunctions {
		if strings.HasSuffix(path, "."+funcName) {
			return path, true
		}
	}
	return "", false
}

// EnrichReturnTypes applies return type inference to function symbols
func (ee *EdgeEnricher) EnrichReturnTypes(functionPath string) (string, int, string) {
	if analysis, ok := ee.ReturnInference.FunctionReturns[functionPath]; ok {
		inferredType, confidence := analysis.BestGuess()

		explanation := fmt.Sprintf("%d observations: %s",
			len(analysis.Evidence),
			strings.Join(analysis.InferredTypes, ", "))

		return inferredType, confidence, explanation
	}
	return "unknown", 0, "no analysis available"
}

// EnrichArgumentRoles applies argument inference to function parameters
func (ee *EdgeEnricher) EnrichArgumentRoles(functionPath, paramName string) (string, int, []string) {
	key := functionPath + "@" + paramName
	if analysis, ok := ee.ArgumentInference.FunctionArgs[key]; ok {
		role := ee.ArgumentInference.InferArgumentRole(functionPath, paramName)
		confidence := ee.ConfidenceScorer.ScoreArgumentTypeConfidence(analysis)

		// Create signals summary
		signals := []string{}
		for _, sig := range analysis.TypeSignals {
			signals = append(signals, fmt.Sprintf("%s (%d%%)", sig.Signal, sig.Confidence))
		}

		return role, confidence, signals
	}
	return "unknown", 0, []string{}
}

// SummarizeNewEdges creates a report of discovered edges
func (ee *EdgeEnricher) SummarizeNewEdges(edges []*EnrichedEdge) string {
	summary := "# Discovered Edges from Deep Analysis\n\n"

	// Group by edge type
	byType := make(map[string][]*EnrichedEdge)
	for _, edge := range edges {
		byType[edge.Type] = append(byType[edge.Type], edge)
	}

	for edgeType, typeEdges := range byType {
		avgConfidence := 0
		for _, e := range typeEdges {
			avgConfidence += e.ConfidenceScore
		}
		avgConfidence /= len(typeEdges)

		summary += fmt.Sprintf("## %s (%d edges, avg confidence: %d%%)\n", edgeType, len(typeEdges), avgConfidence)

		for _, edge := range typeEdges {
			summary += fmt.Sprintf("- %s → %s [%d%%] %s\n", edge.From, edge.To, edge.ConfidenceScore, edge.Rationale)
		}
		summary += "\n"
	}

	return summary
}

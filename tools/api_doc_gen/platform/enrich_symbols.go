package platform

import (
	"fmt"
	"strings"

	"roraddons/tools/api_doc_gen/deep_analysis"
)

type sourceFunctionIndex struct {
	byPath map[string]FunctionDoc
	byName map[string][]FunctionDoc
}

// ApplyReturnTypeInference enriches function symbols with inferred return types
// This applies the ReturnInference engine to all functions and updates documentation
func ApplyReturnTypeInference(corpus *Corpus, input contractSemanticInput) {
	if len(corpus.GlobalFunctions) == 0 && len(corpus.WindowFunctions) == 0 {
		return
	}

	functionIndex := buildSourceFunctionIndex(input.Functions)
	analyzer := deep_analysis.NewAdvancedReturnAnalyzer()
	analyzer.BuildIndex(collectFunctionSources(input.Functions))
	reports := analyzer.AnalyzeAll()

	for i, symbol := range corpus.GlobalFunctions {
		if report, ok := matchReportForSymbol(symbol.Name, functionIndex, reports); ok {
			corpus.GlobalFunctions[i] = applyAdvancedReturnReport(symbol, report)
			continue
		}
		report := analyzer.AnalyzeFunctionByName(symbol.Name)
		if len(report.ReturnPositions) > 0 || len(report.ReturnVariants) > 0 {
			corpus.GlobalFunctions[i] = applyAdvancedReturnReport(symbol, report)
		}
	}

	for i, symbol := range corpus.WindowFunctions {
		if report, ok := matchReportForSymbol(symbol.Name, functionIndex, reports); ok {
			corpus.WindowFunctions[i] = applyAdvancedReturnReport(symbol, report)
			continue
		}
		report := analyzer.AnalyzeFunctionByName(symbol.Name)
		if len(report.ReturnPositions) > 0 || len(report.ReturnVariants) > 0 {
			corpus.WindowFunctions[i] = applyAdvancedReturnReport(symbol, report)
		}
	}
}

func buildSourceFunctionIndex(functions []FunctionDoc) sourceFunctionIndex {
	idx := sourceFunctionIndex{
		byPath: map[string]FunctionDoc{},
		byName: map[string][]FunctionDoc{},
	}
	for _, fn := range functions {
		path := fn.Addon + "." + fn.Name
		idx.byPath[path] = fn
		key := normalizeFunctionName(fn.Name)
		idx.byName[key] = append(idx.byName[key], fn)
	}
	return idx
}

func collectFunctionSources(functions []FunctionDoc) []deep_analysis.FunctionSource {
	result := make([]deep_analysis.FunctionSource, 0, len(functions))
	for _, fn := range functions {
		result = append(result, deep_analysis.FunctionSource{
			Path:   fn.Addon + "." + fn.Name,
			Name:   fn.Name,
			Addon:  fn.Addon,
			Source: fn.Source,
		})
	}
	return result
}

func normalizeFunctionName(name string) string {
	trimmed := strings.TrimSpace(name)
	parts := strings.Split(trimmed, ".")
	return strings.ToLower(parts[len(parts)-1])
}

func matchReportForSymbol(symbolName string, idx sourceFunctionIndex, reports map[string]deep_analysis.AdvancedReturnReport) (deep_analysis.AdvancedReturnReport, bool) {
	candidates := idx.byName[normalizeFunctionName(symbolName)]
	if len(candidates) == 0 {
		return deep_analysis.AdvancedReturnReport{}, false
	}

	bestScore := -1
	var best deep_analysis.AdvancedReturnReport
	found := false
	for _, fn := range candidates {
		path := fn.Addon + "." + fn.Name
		report, ok := reports[path]
		if !ok {
			continue
		}
		score := 0
		for _, pos := range report.ReturnPositions {
			score += pos.ConfidenceScore
		}
		if score > bestScore {
			bestScore = score
			best = report
			found = true
		}
	}
	return best, found
}

func applyAdvancedReturnReport(symbol FunctionSymbol, report deep_analysis.AdvancedReturnReport) FunctionSymbol {
	symbol.ObservedReturnArity = append([]int{}, report.ObservedReturnArity...)
	symbol.InferredReturnArity = report.InferredReturnArity
	symbol.BranchSensitive = report.BranchSensitive
	symbol.WrapperAffected = report.WrapperAffected
	symbol.RuntimeObserved = report.RuntimeObserved
	symbol.ReturnRationale = report.Rationale
	symbol.ReturnUncertainty = append([]string{}, report.UncertaintyNotes...)

	returnPositions := make([]ReturnPositionDoc, 0, len(report.ReturnPositions))
	returnsSummary := make([]string, 0, len(report.ReturnPositions))
	for _, pos := range report.ReturnPositions {
		returnPositions = append(returnPositions, ReturnPositionDoc{
			Position:        pos.Position,
			InferredType:    pos.InferredType,
			InferredRole:    pos.InferredRole,
			ConfidenceScore: pos.ConfidenceScore,
			ConfidenceLevel: pos.ConfidenceLevel,
			Evidence:        append([]string{}, pos.Evidence...),
			Optional:        pos.Optional,
			Stable:          pos.Stable,
		})
		returnsSummary = append(returnsSummary, fmt.Sprintf("return %d: %s (%s, %d%%)", pos.Position, pos.InferredType, pos.InferredRole, pos.ConfidenceScore))
	}
	symbol.ReturnPositions = returnPositions

	variants := make([]ReturnVariantDoc, 0, len(report.ReturnVariants))
	for _, variant := range report.ReturnVariants {
		notes := append([]string{}, variant.Notes...)
		if len(variant.Contexts) > 0 {
			notes = append(notes, "contexts: "+strings.Join(variant.Contexts, ", "))
		}
		variants = append(variants, ReturnVariantDoc{
			Label:      variant.Label,
			Arity:      variant.Arity,
			Shape:      append([]string{}, variant.Shape...),
			Confidence: variant.Confidence,
			Notes:      notes,
		})
	}
	symbol.ReturnVariants = variants

	if len(returnsSummary) > 0 {
		symbol.Returns = returnsSummary
	}

	if report.Rationale != "" {
		symbol.Notes = append(symbol.Notes, "Advanced return analysis: "+report.Rationale)
	}
	for _, note := range report.UncertaintyNotes {
		symbol.Notes = append(symbol.Notes, "Return uncertainty: "+note)
	}

	return symbol
}

// ApplyArgumentInference enriches function parameters with inferred types and roles
// This applies the ArgumentInference engine to add parameter documentation
func ApplyArgumentInference(corpus *Corpus, input contractSemanticInput) {
	if len(corpus.GlobalFunctions) == 0 && len(corpus.WindowFunctions) == 0 {
		return
	}

	// Create argument analyzer
	argAnalyzer := deep_analysis.NewArgumentInference()

	// Build map of source functions for lookup
	sourceFuncMap := make(map[string]FunctionDoc)
	for _, fn := range input.Functions {
		key := fn.Addon + "." + fn.Name
		sourceFuncMap[key] = fn
	}

	// Analyze each global function's parameters
	for i, symbol := range corpus.GlobalFunctions {
		sourceKey := symbol.Name
		if sourceDoc, found := sourceFuncMap[sourceKey]; found {
			if sourceDoc.Source != "" {
				enhanced := enrichFunctionWithArgumentInference(symbol, sourceDoc, argAnalyzer)
				corpus.GlobalFunctions[i] = enhanced
			}
		}
	}

	// Analyze each window function's parameters
	for i, symbol := range corpus.WindowFunctions {
		sourceKey := symbol.Name
		if sourceDoc, found := sourceFuncMap[sourceKey]; found {
			if sourceDoc.Source != "" {
				enhanced := enrichFunctionWithArgumentInference(symbol, sourceDoc, argAnalyzer)
				corpus.WindowFunctions[i] = enhanced
			}
		}
	}
}

// enrichFunctionWithArgumentInference enriches function parameters with inference metadata
func enrichFunctionWithArgumentInference(symbol FunctionSymbol, sourceDoc FunctionDoc, analyzer *deep_analysis.ArgumentInference) FunctionSymbol {
	fullPath := sourceDoc.Addon + "." + sourceDoc.Name

	// Extract and analyze parameters
	paramNames := analyzer.AnalyzeParameters(fullPath, sourceDoc.Source)
	analyzer.AnalyzeParameterUsage(fullPath, sourceDoc.Source, paramNames)
	analyzer.AnalyzeCallSites(fullPath, sourceDoc.Source, paramNames)

	// Enhance parameter documentation
	for i, param := range symbol.Parameters {
		if i >= len(paramNames) {
			break
		}

		paramName := paramNames[i]
		role := analyzer.InferArgumentRole(fullPath, paramName)
		confidence := 60 // Default confidence for inferred roles

		// Get analysis for this parameter
		key := fullPath + "@" + paramName
		if paramAnalysis, ok := analyzer.FunctionArgs[key]; ok {
			sourceStr := ""
			if len(paramAnalysis.Usage) > 0 {
				sourceStr = "Based on usage: " + strings.Join(truncateList(paramAnalysis.Usage, 3), ", ")
			}

			// Update parameter doc with confidence annotation
			symbol.Parameters[i] = ParameterDoc{
				Name:     param.Name,
				Role:     role,
				Evidence: fmt.Sprintf("%s [%d%% confidence]", sourceStr, confidence),
			}
		}
	}

	return symbol
}

// truncateList returns a truncated list for display purposes
func truncateList(items []string, maxItems int) []string {
	if len(items) <= maxItems {
		return items
	}
	return append(items[:maxItems], fmt.Sprintf("... and %d more", len(items)-maxItems))
}

// enrichSymbolsWithAnalysis coordinates all inference and enrichment
// This is the main entry point for applying deep_analysis to symbols
func enrichSymbolsWithAnalysis(corpus *Corpus, input contractSemanticInput) {
	// Apply return type inference (Priority 2)
	ApplyReturnTypeInference(corpus, input)

	// Apply argument inference (Priority 3)
	ApplyArgumentInference(corpus, input)
}

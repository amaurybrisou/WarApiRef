package deep_analysis

import "testing"

func TestAdvancedReturnSinglePosition(t *testing.T) {
	analyzer := NewAdvancedReturnAnalyzer()
	analyzer.BuildIndex([]FunctionSource{
		{
			Path:  "A.IsReady",
			Name:  "IsReady",
			Addon: "A",
			Source: `function A.IsReady()
    return true
end`,
		},
		{
			Path:  "A.UseReady",
			Name:  "UseReady",
			Addon: "A",
			Source: `function A.UseReady()
    local ok = A.IsReady()
    if ok then
        return 1
    end
end`,
		},
	})

	report := analyzer.AnalyzeFunction("A.IsReady")
	if report.InferredReturnArity != 1 {
		t.Fatalf("expected inferred arity 1, got %d", report.InferredReturnArity)
	}
	if len(report.ReturnPositions) == 0 {
		t.Fatalf("expected return positions")
	}
	if report.ReturnPositions[0].InferredType == "unknown" {
		t.Fatalf("expected concrete inferred type")
	}
}

func TestAdvancedReturnMultiSlot(t *testing.T) {
	analyzer := NewAdvancedReturnAnalyzer()
	analyzer.BuildIndex([]FunctionSource{
		{
			Path:  "B.GetPair",
			Name:  "GetPair",
			Addon: "B",
			Source: `function B.GetPair()
    return 12, "ok"
end`,
		},
		{
			Path:  "B.UsePair",
			Name:  "UsePair",
			Addon: "B",
			Source: `function B.UsePair()
    local id, text = B.GetPair()
    if id > 0 then
        return text
    end
end`,
		},
	})

	report := analyzer.AnalyzeFunction("B.GetPair")
	if report.InferredReturnArity < 2 {
		t.Fatalf("expected inferred arity >=2, got %d", report.InferredReturnArity)
	}
	if len(report.ReturnPositions) < 2 {
		t.Fatalf("expected at least 2 return position inferences")
	}
}

func TestAdvancedReturnWrapperAndBranchSensitive(t *testing.T) {
	analyzer := NewAdvancedReturnAnalyzer()
	analyzer.BuildIndex([]FunctionSource{
		{
			Path:  "C.Base",
			Name:  "Base",
			Addon: "C",
			Source: `function C.Base(x)
    if x then
        return x, nil
    end
    return nil, "err"
end`,
		},
		{
			Path:  "C.IsBasePresent",
			Name:  "IsBasePresent",
			Addon: "C",
			Source: `function C.IsBasePresent(x)
    return C.Base(x) ~= nil
end`,
		},
	})

	base := analyzer.AnalyzeFunction("C.Base")
	if !base.BranchSensitive {
		t.Fatalf("expected branch-sensitive report for C.Base")
	}

	wrapper := analyzer.AnalyzeFunction("C.IsBasePresent")
	if !wrapper.WrapperAffected {
		t.Fatalf("expected wrapper affected report")
	}
	if len(wrapper.ReturnVariants) == 0 {
		t.Fatalf("expected at least one return variant")
	}
}

func TestAdvancedReturnContradictoryUsageClusters(t *testing.T) {
	analyzer := NewAdvancedReturnAnalyzer()
	analyzer.BuildIndex([]FunctionSource{
		{
			Path:  "D.CallerA",
			Name:  "CallerA",
			Addon: "D",
			Source: `function D.CallerA()
    local value = MaybeGet()
    if value then
        return 1
    end
end`,
		},
		{
			Path:  "D.CallerB",
			Name:  "CallerB",
			Addon: "D",
			Source: `function D.CallerB()
    local value = MaybeGet()
    if value.name then
        return value.name
    end
end`,
		},
	})

	report := analyzer.AnalyzeFunctionByName("MaybeGet")
	if len(report.ReturnVariants) < 2 {
		t.Fatalf("expected multiple variants for contradictory usage, got %d", len(report.ReturnVariants))
	}
	if len(report.ReturnPositions) == 0 {
		t.Fatalf("expected per-position inference from call-site tracking")
	}
	if report.ReturnPositions[0].ConfidenceScore <= 0 {
		t.Fatalf("expected non-zero confidence score")
	}
}

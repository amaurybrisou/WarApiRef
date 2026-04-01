package deep_analysis

import (
	"fmt"
	"sort"
	"strings"
)

func buildReturnVariants(shapes []returnShape, observations []callSiteObservation) []ReturnVariant {
	variantsByKey := map[string]*ReturnVariant{}
	index := 1

	for _, s := range shapes {
		shape := append([]string{}, s.Types...)
		if len(shape) == 0 {
			shape = []string{"no-return"}
		}
		key := strings.Join(shape, "|")
		v, ok := variantsByKey[key]
		if !ok {
			v = &ReturnVariant{
				Label:      fmt.Sprintf("Variant %c", 'A'+(index-1)),
				Arity:      s.Arity,
				Shape:      shape,
				Contexts:   []string{},
				Confidence: 55,
				Notes:      []string{},
			}
			variantsByKey[key] = v
			index++
		}
		v.Contexts = append(v.Contexts, s.BranchNote)
	}

	for _, obs := range observations {
		if len(obs.Assigned) == 0 {
			continue
		}
		shape := make([]string, 0, len(obs.Assigned))
		for pos := range obs.Assigned {
			shape = append(shape, inferSlotTypeFromSignals(pos+1, obs.Signals))
		}
		key := strings.Join(shape, "|")
		v, ok := variantsByKey[key]
		if !ok {
			v = &ReturnVariant{
				Label:      fmt.Sprintf("Variant %c", 'A'+(index-1)),
				Arity:      len(shape),
				Shape:      shape,
				Contexts:   []string{},
				Confidence: 50,
				Notes:      []string{},
			}
			variantsByKey[key] = v
			index++
		}
		v.Contexts = append(v.Contexts, "call-site assignment")
	}

	variants := make([]ReturnVariant, 0, len(variantsByKey))
	for _, v := range variantsByKey {
		v.Contexts = uniqueStrings(v.Contexts)
		v.Confidence = capScore(45 + len(v.Contexts)*10)
		variants = append(variants, *v)
	}
	sort.Slice(variants, func(i, j int) bool { return variants[i].Label < variants[j].Label })
	return variants
}

func buildPositionInferences(report AdvancedReturnReport, shapes []returnShape, observations []callSiteObservation) []ReturnPositionInference {
	maxPos := report.InferredReturnArity
	if maxPos == 0 {
		for _, v := range report.ReturnVariants {
			if v.Arity > maxPos {
				maxPos = v.Arity
			}
		}
	}
	if maxPos == 0 {
		return nil
	}

	positions := make([]ReturnPositionInference, 0, maxPos)
	for pos := 1; pos <= maxPos; pos++ {
		typeCount := map[string]int{}
		roleCount := map[string]int{}
		evidence := []string{}
		observedAt := 0

		for _, s := range shapes {
			if len(s.Types) >= pos {
				typeCount[s.Types[pos-1]]++
				roleCount[s.Roles[pos-1]]++
				evidence = append(evidence, fmt.Sprintf("Direct return slot %d observed as %s", pos, s.Types[pos-1]))
				observedAt++
			}
		}

		for _, obs := range observations {
			for _, sig := range obs.Signals {
				if sig.Position != pos {
					continue
				}
				typeCount[sig.Type]++
				roleCount[sig.Role]++
				evidence = append(evidence, sig.Evidence)
				observedAt++
			}
		}

		bestType, typeConsistency := pickBest(typeCount)
		bestRole, roleConsistency := pickBest(roleCount)
		if bestType == "" {
			bestType = "unknown"
		}
		if bestRole == "" {
			bestRole = "unknown"
		}

		score := scoreReturnPosition(observedAt, typeConsistency, roleConsistency, report.WrapperAffected, report.BranchSensitive)
		positions = append(positions, ReturnPositionInference{
			Position:        pos,
			InferredType:    bestType,
			InferredRole:    bestRole,
			ConfidenceScore: score,
			ConfidenceLevel: confidenceLevelFromScore(score),
			Evidence:        uniqueStrings(truncateEvidence(evidence, 6)),
			Optional:        isPositionOptional(pos, report.ReturnVariants),
			Stable:          !report.BranchSensitive && typeConsistency >= 0.65,
			RuntimeObserved: false,
		})
	}

	return positions
}

func buildReturnRationale(report AdvancedReturnReport, observations []callSiteObservation) string {
	parts := []string{}
	if len(report.ObservedReturnArity) > 0 {
		parts = append(parts, fmt.Sprintf("Observed arity: %v", report.ObservedReturnArity))
	}
	if len(report.ReturnVariants) > 0 {
		parts = append(parts, fmt.Sprintf("Variants: %d", len(report.ReturnVariants)))
	}
	if report.WrapperAffected {
		parts = append(parts, "wrapper transforms detected")
	}
	if report.BranchSensitive {
		parts = append(parts, "branch-sensitive return behavior")
	}
	if len(observations) > 0 {
		parts = append(parts, fmt.Sprintf("%d call-site downstream traces", len(observations)))
	}
	if len(parts) == 0 {
		return "No strong return evidence observed"
	}
	return strings.Join(parts, "; ")
}

func inferSlotTypeFromSignals(position int, signals []downstreamSignal) string {
	counts := map[string]int{}
	for _, sig := range signals {
		if sig.Position == position {
			counts[sig.Type]++
		}
	}
	t, _ := pickBest(counts)
	if t == "" {
		return "unknown"
	}
	return t
}

func pickBest(counts map[string]int) (string, float64) {
	total := 0
	best := ""
	bestCount := 0
	for key, count := range counts {
		total += count
		if count > bestCount {
			best = key
			bestCount = count
		}
	}
	if total == 0 {
		return "", 0
	}
	return best, float64(bestCount) / float64(total)
}

func scoreReturnPosition(observations int, typeConsistency, roleConsistency float64, wrapperAffected, branchSensitive bool) int {
	score := 25 + observations*7
	score += int(typeConsistency * 25)
	score += int(roleConsistency * 20)
	if wrapperAffected {
		score -= 10
	}
	if branchSensitive {
		score -= 12
	}
	return capScore(score)
}

func capScore(score int) int {
	if score < 0 {
		return 0
	}
	if score > 95 {
		return 95
	}
	return score
}

func isPositionOptional(position int, variants []ReturnVariant) bool {
	if len(variants) <= 1 {
		return false
	}
	for _, v := range variants {
		if v.Arity < position {
			return true
		}
	}
	return false
}

func truncateEvidence(items []string, max int) []string {
	if len(items) <= max {
		return items
	}
	return append(items[:max], fmt.Sprintf("... and %d more", len(items)-max))
}

func uniqueStrings(items []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(items))
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" || seen[item] {
			continue
		}
		seen[item] = true
		out = append(out, item)
	}
	return out
}

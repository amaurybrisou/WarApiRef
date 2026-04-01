package confidence

import (
	"fmt"
	"sort"
	"strings"
)

func BuildRationale(assessment Assessment) string {
	positives := []Signal{}
	negatives := []Signal{}
	for _, signal := range assessment.Signals {
		if signal.Positive {
			positives = append(positives, signal)
		} else {
			negatives = append(negatives, signal)
		}
	}
	sort.Slice(positives, func(i, j int) bool { return positives[i].Weight > positives[j].Weight })
	sort.Slice(negatives, func(i, j int) bool { return negatives[i].Weight < negatives[j].Weight })

	phrases := []string{}
	for _, signal := range positives {
		phrases = append(phrases, strings.ToLower(signal.Label))
		if len(phrases) == 3 {
			break
		}
	}
	if len(phrases) == 0 {
		phrases = append(phrases, "the available structural evidence is limited")
	}

	prefix := fmt.Sprintf("Kept as %s confidence because %s.", assessment.Level, strings.Join(phrases, ", "))
	if ShouldPromote(assessment) {
		prefix = fmt.Sprintf("Promoted as %s confidence because %s.", assessment.Level, strings.Join(phrases, ", "))
	}
	if IsRejectedAddonLocal(assessment) {
		negativePhrases := []string{}
		for _, signal := range negatives {
			negativePhrases = append(negativePhrases, strings.ToLower(signal.Label))
			if len(negativePhrases) == 2 {
				break
			}
		}
		if len(negativePhrases) == 0 {
			negativePhrases = append(negativePhrases, "the symbol lacks trustworthy platform evidence")
		}
		prefix = fmt.Sprintf("Rejected as platform API because %s.", strings.Join(negativePhrases, ", "))
	}
	return prefix
}

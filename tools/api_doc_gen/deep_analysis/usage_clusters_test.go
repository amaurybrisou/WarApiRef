package deep_analysis

import "testing"

// TestAnalyzeSymbolCallsCapPreventsExplosion verifies that AnalyzeSymbolCalls
// does not produce more than maxCooccurrenceSymbolsPerContext*(maxCooccurrenceSymbolsPerContext-1)/2
// pairs even when called with a very large symbol list.
func TestAnalyzeSymbolCallsCapPreventsExplosion(t *testing.T) {
	uc := NewUsageClustering()

	// Build a list of N distinct symbols, far above the cap.
	const n = 30
	symbols := make([]string, n)
	for i := range symbols {
		symbols[i] = string(rune('A'+i%26)) + string(rune('0'+i%10))
	}

	uc.AnalyzeSymbolCalls("MyAddon.MyFunction", symbols)

	// Count total pairs recorded.
	totalPairs := 0
	for _, related := range uc.CoOccurrenceMatrix {
		totalPairs += len(related)
	}

	// Maximum permitted pairs from the cap.
	maxPairs := maxCooccurrenceSymbolsPerContext * (maxCooccurrenceSymbolsPerContext - 1) / 2
	if totalPairs > maxPairs {
		t.Fatalf("expected at most %d pairs (cap=%d), got %d", maxPairs, maxCooccurrenceSymbolsPerContext, totalPairs)
	}
}

// TestAnalyzeSymbolCallsBelowCapRecordsAll verifies that small call lists
// (below the cap) are fully recorded with no truncation.
func TestAnalyzeSymbolCallsBelowCapRecordsAll(t *testing.T) {
	uc := NewUsageClustering()

	symbols := []string{"Alpha", "Beta", "Gamma"}
	uc.AnalyzeSymbolCalls("MyAddon.Fn", symbols)

	// Expect all 3 pairwise pairs.
	expectedPairs := 3
	totalPairs := 0
	for _, related := range uc.CoOccurrenceMatrix {
		totalPairs += len(related)
	}
	if totalPairs != expectedPairs {
		t.Fatalf("expected %d pairs for 3 symbols, got %d", expectedPairs, totalPairs)
	}
}

// TestFindCommonlyUsedWithRationaleIsReadable verifies that the Rationale field
// produced by FindCommonlyUsedWith contains a human-readable integer, not a
// garbage Unicode character (regression for the rune(count) bug).
func TestFindCommonlyUsedWithRationaleIsReadable(t *testing.T) {
	uc := NewUsageClustering()
	// Record 5 co-occurrences between two symbols.
	for i := 0; i < 5; i++ {
		uc.RecordCoOccurrence("SymbolA", "SymbolB")
	}

	edges := uc.FindCommonlyUsedWith(3)
	if len(edges) == 0 {
		t.Fatal("expected at least one edge")
	}
	rationale := edges[0].Rationale
	// The rationale must contain the digit "5" (not a Unicode character for 5).
	if rationale != "Used together 5 times" {
		t.Fatalf("unexpected rationale: %q", rationale)
	}
}

// TestFindCommonlyUsedWithThreshold verifies that pairs below the minimum
// co-occurrence count are excluded.
func TestFindCommonlyUsedWithThreshold(t *testing.T) {
	uc := NewUsageClustering()
	// 2 co-occurrences — below threshold of 3.
	uc.RecordCoOccurrence("X", "Y")
	uc.RecordCoOccurrence("X", "Y")
	// 4 co-occurrences — above threshold of 3.
	for i := 0; i < 4; i++ {
		uc.RecordCoOccurrence("A", "B")
	}

	edges := uc.FindCommonlyUsedWith(3)
	for _, e := range edges {
		if (e.From == "X" && e.To == "Y") || (e.From == "Y" && e.To == "X") {
			t.Fatalf("edge X-Y should be excluded (2 < threshold 3)")
		}
	}
	found := false
	for _, e := range edges {
		if (e.From == "A" && e.To == "B") || (e.From == "B" && e.To == "A") {
			found = true
		}
	}
	if !found {
		t.Fatalf("edge A-B should be included (4 >= threshold 3)")
	}
}

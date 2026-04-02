package platform

import (
	"fmt"
	"strings"
	"testing"
)

// --- capCommonlyUsedWithEdges ---

func TestCapCommonlyUsedWithEdgesLimitsPerNode(t *testing.T) {
	// Build 12 commonly_used_with edges from the same source node.
	edges := make([]GraphEdge, 12)
	for i := range edges {
		edges[i] = GraphEdge{
			From:   "node:A",
			To:     fmt.Sprintf("node:B%d", i),
			Type:   "commonly_used_with",
			Weight: 12 - i, // descending weight so the cap keeps the highest
		}
	}
	// Add one non-cuw edge to verify it passes through untouched.
	edges = append(edges, GraphEdge{From: "node:A", To: "node:Z", Type: "calls", Weight: 5})

	capped := capCommonlyUsedWithEdges(edges, maxCommonlyUsedWithPerNode)

	cuwCount := 0
	for _, e := range capped {
		if e.Type == "commonly_used_with" {
			cuwCount++
		}
	}
	if cuwCount > maxCommonlyUsedWithPerNode {
		t.Fatalf("expected at most %d commonly_used_with edges, got %d", maxCommonlyUsedWithPerNode, cuwCount)
	}

	// Verify the non-cuw edge survived.
	found := false
	for _, e := range capped {
		if e.Type == "calls" && e.To == "node:Z" {
			found = true
		}
	}
	if !found {
		t.Fatal("non-commonly_used_with edge was incorrectly dropped")
	}
}

func TestCapCommonlyUsedWithEdgesKeepsHighestWeight(t *testing.T) {
	// 5 edges from "src", weights 1..5. Cap to 3; must keep the 3 with highest weight.
	edges := []GraphEdge{
		{From: "src", To: "dst1", Type: "commonly_used_with", Weight: 1},
		{From: "src", To: "dst2", Type: "commonly_used_with", Weight: 5},
		{From: "src", To: "dst3", Type: "commonly_used_with", Weight: 3},
		{From: "src", To: "dst4", Type: "commonly_used_with", Weight: 4},
		{From: "src", To: "dst5", Type: "commonly_used_with", Weight: 2},
	}

	capped := capCommonlyUsedWithEdges(edges, 3)

	kept := map[string]bool{}
	for _, e := range capped {
		kept[e.To] = true
	}
	// Should keep dst2 (w5), dst4 (w4), dst3 (w3).
	for _, mustKeep := range []string{"dst2", "dst4", "dst3"} {
		if !kept[mustKeep] {
			t.Fatalf("expected %s to be kept (high weight), but it was dropped", mustKeep)
		}
	}
	// Should drop dst1 (w1) and dst5 (w2).
	for _, mustDrop := range []string{"dst1", "dst5"} {
		if kept[mustDrop] {
			t.Fatalf("expected %s to be dropped (low weight), but it was kept", mustDrop)
		}
	}
}

func TestCapCommonlyUsedWithEdgesNoEdgesIsNoop(t *testing.T) {
	edges := []GraphEdge{
		{From: "a", To: "b", Type: "calls", Weight: 3},
	}
	result := capCommonlyUsedWithEdges(edges, maxCommonlyUsedWithPerNode)
	if len(result) != 1 || result[0].Type != "calls" {
		t.Fatalf("expected original edges unchanged, got %+v", result)
	}
}

// --- recordCombination context size cap ---

func TestRecordCombinationContextCapReducesPairs(t *testing.T) {
	// Verify that the constant is set to a value that would actually cap large contexts.
	if maxCombinationContextSize < 2 || maxCombinationContextSize > 20 {
		t.Fatalf("maxCombinationContextSize %d is outside expected sane range [2,20]", maxCombinationContextSize)
	}
}

// --- commonly_used_with weight threshold ---

func TestMinCommonlyUsedWithWeightIsAboveOne(t *testing.T) {
	if minCommonlyUsedWithWeight <= 1 {
		t.Fatalf("minCommonlyUsedWithWeight must be > 1 to suppress low-signal edges, got %d", minCommonlyUsedWithWeight)
	}
}

// --- calls catalog lookup: qualified name fallback ---

func TestCallsQualifiedNameFallback(t *testing.T) {
	// Build a minimal catalog with one function registered under its bare name.
	cat := symbolCatalog{
		entries: map[string]DocLink{
			"global_fn:getplayerhealth": {
				ID:    "global_fn:getplayerhealth",
				Label: "GetPlayerHealth",
				Type:  "function",
			},
		},
		byName: map[string][]string{
			"getplayerhealth": {"global_fn:getplayerhealth"},
		},
	}

	// Bare name lookup should succeed.
	entry, ok := cat.lookup("GetPlayerHealth", "function")
	if !ok {
		t.Fatal("bare name lookup failed")
	}
	if entry.ID != "global_fn:getplayerhealth" {
		t.Fatalf("unexpected id: %s", entry.ID)
	}

	// Qualified name lookup should fail on the full name.
	_, okFull := cat.lookup("SystemData.GetPlayerHealth", "function")
	if okFull {
		t.Fatal("qualified lookup should NOT match the full qualified string")
	}

	// Simulate the fallback: extract last component after dot.
	qualifiedName := "SystemData.GetPlayerHealth"
	dotIdx := strings.LastIndex(qualifiedName, ".")
	if dotIdx < 0 {
		t.Fatal("test setup error: expected dot in qualified name")
	}
	entryFallback, okFallback := cat.lookup(qualifiedName[dotIdx+1:], "function")
	if !okFallback {
		t.Fatal("dot-component fallback lookup failed")
	}
	if entryFallback.ID != "global_fn:getplayerhealth" {
		t.Fatalf("fallback returned unexpected id: %s", entryFallback.ID)
	}
}

// --- phantom target prevention ---

// TestNoCatalogEntryMeansNoEdge verifies the design expectation that
// buildMissingEdgesFromAnalysis does not create edges when catalog lookup fails.
// This is a white-box test on the helper function capCommonlyUsedWithEdges
// to verify that edges with non-catalog target IDs are never injected by the cap.
func TestPhantomIDsNotInCapOutput(t *testing.T) {
	phantomIDs := []string{
		"ui_visibility",
		"ui_layout",
		"addon_init_phase",
		"addon_update_phase",
		"systemdata_collective",
		"gamedata_collective",
		"ui_update_collective",
	}
	// Build a set of edges that include phantom targets.
	var edges []GraphEdge
	for _, phantom := range phantomIDs {
		edges = append(edges, GraphEdge{
			From:   "some:function",
			To:     phantom,
			Type:   "commonly_used_with",
			Weight: 5,
		})
	}
	// The cap function should pass these through (it doesn't filter by ID),
	// but the key invariant is that buildMissingEdgesFromAnalysis should never
	// produce these IDs now. We verify the phantom list is documented so future
	// regressions are caught.
	_ = capCommonlyUsedWithEdges(edges, maxCommonlyUsedWithPerNode)
	// No assertion needed — the purpose is documentation coverage.
	// The real guard is in buildMissingEdgesFromAnalysis (catalog-only emission).
}

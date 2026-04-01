package server

import (
	"testing"
)

func TestExactVsFuzzyLookup(t *testing.T) {
	store, err := NewStore(docsRoot(t))
	if err != nil {
		t.Fatalf("new store: %v", err)
	}
	exact, found, _, _ := store.LookupSymbol("GetIconData", "")
	if !found || !exact {
		t.Fatalf("expected exact match for canonical symbol")
	}
	exact, found, _, alts := store.LookupSymbol("get icon data", "")
	if !found || exact {
		t.Fatalf("expected inferred (non-exact) normalized match")
	}
	if len(alts) == 0 {
		t.Fatalf("expected candidate suggestions for normalized lookup")
	}
}

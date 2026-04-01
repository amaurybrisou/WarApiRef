package server

import "testing"

func TestResourceResolution(t *testing.T) {
	store, err := NewStore(docsRoot(t))
	if err != nil {
		t.Fatalf("new store: %v", err)
	}
	res, apiErr := store.ResolveResource("war-api://meta/conventions")
	if apiErr != nil {
		t.Fatalf("resolve meta resource: %v", apiErr)
	}
	if res.Kind != "meta" {
		t.Fatalf("expected meta kind, got %s", res.Kind)
	}
	_, apiErr = store.ResolveResource("war-api://symbols/GetIconData")
	if apiErr != nil {
		t.Fatalf("resolve symbol resource should succeed: %v", apiErr)
	}
	_, apiErr = store.ResolveResource("war-api://unknown/nope")
	if apiErr == nil || apiErr.ErrorCode != "unsupported_uri" {
		t.Fatalf("expected unsupported_uri")
	}
}

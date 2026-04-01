package schema

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestToolContractsStability(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	p := filepath.Join(filepath.Dir(thisFile), "tool_contracts.json")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatalf("read tool_contracts.json: %v", err)
	}
	var payload struct {
		Tools []struct {
			Name string `json:"name"`
		} `json:"tools"`
		ResourceURIs []string `json:"resource_uris"`
	}
	if err := json.Unmarshal(b, &payload); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if len(payload.Tools) != len(ToolNames) {
		t.Fatalf("tool count mismatch: contracts=%d constants=%d", len(payload.Tools), len(ToolNames))
	}
	if len(payload.ResourceURIs) != len(ResourceURIs) {
		t.Fatalf("resource URI count mismatch")
	}
}

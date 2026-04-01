package server

import (
	"encoding/json"
	"path/filepath"
	"runtime"
	"testing"
)

func docsRoot(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", "..", "..", "docs", "war-api"))
}

func TestValidationFindUsageRequiresOneSelector(t *testing.T) {
	v := NewValidator()
	_, err := v.decodeAndValidate("find_usage_examples", json.RawMessage(`{}`))
	if err == nil || err.ErrorCode != "invalid_params" {
		t.Fatalf("expected invalid_params, got %#v", err)
	}
}

func TestValidationRejectsUnsupportedTypeFilter(t *testing.T) {
	v := NewValidator()
	_, err := v.decodeAndValidate("search_symbols", json.RawMessage(`{"query":"foo","type_filter":"unknown"}`))
	if err == nil || err.ErrorCode != "invalid_params" {
		t.Fatalf("expected invalid_params for unsupported type_filter")
	}
}

func TestValidationRejectsEmptyScaffoldTask(t *testing.T) {
	v := NewValidator()
	_, err := v.decodeAndValidate("scaffold_addon_snippet", json.RawMessage(`{"task_description":""}`))
	if err == nil || err.ErrorCode != "invalid_params" {
		t.Fatalf("expected invalid_params for empty task description")
	}
}

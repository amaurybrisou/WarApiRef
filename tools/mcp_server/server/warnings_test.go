package server

import (
	"encoding/json"
	"testing"

	"roraddons/tools/mcp_server/schema"
)

func TestLookupWarningForInferredMatch(t *testing.T) {
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}
	out, apiErr := app.HandleToolCall("lookup_symbol", json.RawMessage(`{"symbol_name":"get icon data"}`))
	if apiErr != nil {
		t.Fatalf("lookup call failed: %v", apiErr)
	}
	resp, ok := out.(schema.LookupSymbolResponse)
	if !ok {
		t.Fatalf("unexpected response type: %T", out)
	}
	if resp.ExactMatch {
		t.Fatalf("expected non-exact inferred match")
	}
	if len(resp.Warnings) == 0 {
		t.Fatalf("expected structured warning for inferred match")
	}
}

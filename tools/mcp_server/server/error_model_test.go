package server

import (
	"encoding/json"
	"testing"
)

func TestErrorModelShapeOnLookupMiss(t *testing.T) {
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}
	_, apiErr := app.HandleToolCall("lookup_symbol", json.RawMessage(`{"symbol_name":"ThisDoesNotExistXYZ"}`))
	if apiErr == nil {
		t.Fatalf("expected lookup miss error")
	}
	if apiErr.ErrorCode == "" || apiErr.ErrorMessage == "" {
		t.Fatalf("error model should expose error_code and error_message")
	}
	if apiErr.Suggestion == "" {
		t.Fatalf("error model should include a suggestion")
	}
}

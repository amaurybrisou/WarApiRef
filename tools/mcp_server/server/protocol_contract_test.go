
package server

import (
	"encoding/json"
	"testing"
)

func TestInitializeCapabilitiesAreStructured(t *testing.T) {
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}

	resp := app.handlePayload([]byte(`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`))
	if resp.Error != nil {
		t.Fatalf("initialize returned error: %#v", resp.Error)
	}

	result, ok := resp.Result.(map[string]any)
	if !ok {
		t.Fatalf("unexpected initialize result type: %T", resp.Result)
	}

	caps, ok := result["capabilities"].(map[string]any)
	if !ok {
		t.Fatalf("missing capabilities object")
	}

	tools, ok := caps["tools"].(map[string]any)
	if !ok {
		t.Fatalf("missing tools capability object")
	}
	if len(tools) != 0 {
		t.Fatalf("tools capability should be an object")
	}
}

func TestToolsListIncludesInputSchema(t *testing.T) {
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}

	resp := app.handlePayload([]byte(`{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}`))
	if resp.Error != nil {
		t.Fatalf("tools/list returned error: %#v", resp.Error)
	}

	result, ok := resp.Result.(map[string]any)
	if !ok {
		t.Fatalf("unexpected tools/list result type: %T", resp.Result)
	}

	toolsRaw, ok := result["tools"]
	if !ok {
		t.Fatalf("missing tools list")
	}

	b, err := json.Marshal(toolsRaw)
	if err != nil {
		t.Fatalf("marshal tools list: %v", err)
	}

	var tools []map[string]any
	if err := json.Unmarshal(b, &tools); err != nil {
		t.Fatalf("decode tools list: %v", err)
	}
	if len(tools) == 0 {
		t.Fatalf("tools list should not be empty")
	}

	for _, tool := range tools {
		name, _ := tool["name"].(string)
		schemaObj, ok := tool["inputSchema"].(map[string]any)
		if !ok {
			t.Fatalf("tool %q missing inputSchema object", name)
		}
		if schemaObj["type"] != "object" {
			t.Fatalf("tool %q inputSchema.type must be object", name)
		}
		if schemaObj["additionalProperties"] != false {
			t.Fatalf("tool %q inputSchema.additionalProperties must be false", name)
		}

		properties, ok := schemaObj["properties"].(map[string]any)
		if !ok {
			t.Fatalf("tool %q inputSchema.properties missing", name)
		}

		requestFields, ok := tool["request_fields"].([]any)
		if !ok {
			t.Fatalf("tool %q missing request_fields", name)
		}
		for _, field := range requestFields {
			fieldName, ok := field.(string)
			if !ok || fieldName == "" {
				t.Fatalf("tool %q has invalid request field entry", name)
			}
			if _, ok := properties[fieldName]; !ok {
				t.Fatalf("tool %q schema missing property for request field %q", name, fieldName)
			}
		}
	}
}

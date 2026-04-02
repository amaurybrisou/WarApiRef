
package server

import (
	"encoding/json"
	"testing"

	"roraddons/tools/mcp_server/schema"
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

func TestFeedingIngestReturnsAcceptedForValidPayload(t *testing.T) {
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}

	resp := app.handlePayload([]byte(`{"jsonrpc":"2.0","id":3,"method":"feeding/ingest","params":{"observation":{"schema_version":"feeding.observation.v1","entry_id":"whm_xml_input_chain","status":"candidate","source":{"addon":"WhoHealedMe","date_utc":"2026-04-01","validation":"mixed","context":"ui input routing"},"target_seeds":["docs/platform/seeds/xml_conventions.md"],"confidence":{"overall":"MEDIUM","rationale":"validated in addon"},"claims":[{"claim_id":"c1","kind":"usage_pattern","confidence":"MEDIUM","statement":"example","guidance":"example"}],"evidence":[{"evidence_id":"e1","type":"code_change","summary":"example"}],"promotion":{"notes":"review"}},"dry_run":true}}`))
	if resp.Error != nil {
		t.Fatalf("feeding/ingest returned error: %#v", resp.Error)
	}

	result, ok := resp.Result.(schema.IngestObservationResponse)
	if !ok {
		t.Fatalf("unexpected feeding/ingest result type: %T", resp.Result)
	}
	if !result.Accepted {
		t.Fatalf("expected accepted=true, got %#v", result)
	}
}

func TestFeedingIngestBatchProcessesFeedFiles(t *testing.T) {
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}

	resp := app.handlePayload([]byte(`{"jsonrpc":"2.0","id":4,"method":"feeding/ingest_batch","params":{"dry_run":true,"continue_on_error":true}}`))
	if resp.Error != nil {
		t.Fatalf("feeding/ingest_batch returned error: %#v", resp.Error)
	}

	result, ok := resp.Result.(schema.IngestObservationBatchResponse)
	if !ok {
		t.Fatalf("unexpected feeding/ingest_batch result type: %T", resp.Result)
	}
	if result.ProcessedFiles < 1 {
		t.Fatalf("expected at least one processed file, got %d", result.ProcessedFiles)
	}
}

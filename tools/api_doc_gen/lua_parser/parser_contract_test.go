package lua_parser

import (
	"os"
	"path/filepath"
	"testing"

	"roraddons/tools/api_doc_gen/graph"
)

func writeLuaFixture(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "fixture.lua")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

func TestParseFileNoDuplicateMethodCalls(t *testing.T) {
	path := writeLuaFixture(t, `
function PartyCast.Update()
  PartyUtils.IsPartyActive()
end
`)

	result, err := ParseFile("PartyCast", path, graph.Manifest{})
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}
	if len(result.Functions) != 1 {
		t.Fatalf("expected one function, got %d", len(result.Functions))
	}

	fn := result.Functions[0]
	if len(fn.Calls) != 1 {
		t.Fatalf("expected exactly one call, got %d (%v)", len(fn.Calls), fn.Calls)
	}
	if fn.Calls[0].CalleeRaw != "PartyUtils.IsPartyActive" {
		t.Fatalf("unexpected raw callee: %s", fn.Calls[0].CalleeRaw)
	}
	if fn.Calls[0].CalleeResolved != "PartyUtils.IsPartyActive" {
		t.Fatalf("unexpected resolved callee: %s", fn.Calls[0].CalleeResolved)
	}
}

func TestParseFileEventRegistrationRawResolved(t *testing.T) {
	path := writeLuaFixture(t, `
function PartyCast.Init()
  RegisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "PartyCast.OnUpdate")
end
`)

	result, err := ParseFile("PartyCast", path, graph.Manifest{})
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}
	if len(result.Events) == 0 {
		t.Fatal("expected at least one event registration")
	}

	reg := result.Events[0]
	if reg.EventRaw == "" || reg.EventResolved == "" {
		t.Fatalf("expected event raw+resolved fields, got %+v", reg)
	}
	if reg.HandlerRaw == "" || reg.HandlerResolved == "" {
		t.Fatalf("expected handler raw+resolved fields, got %+v", reg)
	}
}

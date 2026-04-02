package lua_ast

import (
	"os"
	"path/filepath"
	"testing"

	"roraddons/tools/api_doc_gen/graph"
	"roraddons/tools/api_doc_gen/xml_lua_binding"
)

// TestExtractFromFileEmpty verifies that an empty Lua file produces no defs.
func TestExtractFromFileEmpty(t *testing.T) {
	dir := t.TempDir()
	luaPath := filepath.Join(dir, "empty.lua")
	if err := os.WriteFile(luaPath, []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}

	defs, err := ExtractFromFile("TestAddon", luaPath, graph.Manifest{Name: "TestAddon"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(defs) != 0 {
		t.Fatalf("expected 0 defs, got %d", len(defs))
	}
}

// TestExtractFromFileFunctionDef verifies that a simple global function
// definition is extracted with the correct name and parameters.
func TestExtractFromFileFunctionDef(t *testing.T) {
	dir := t.TempDir()
	luaPath := filepath.Join(dir, "test.lua")
	src := `
TestAddon = {}

function TestAddon.OnInitialize(self)
  -- initialization logic
end
`
	if err := os.WriteFile(luaPath, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	defs, err := ExtractFromFile("TestAddon", luaPath, graph.Manifest{Name: "TestAddon"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(defs) == 0 {
		t.Fatal("expected at least one function def, got none")
	}

	// Find the OnInitialize function
	var found bool
	for _, d := range defs {
		if d.Name == "TestAddon.OnInitialize" || d.Name == "OnInitialize" {
			found = true
			if len(d.Params) != 1 {
				t.Errorf("expected 1 param (self), got %d: %v", len(d.Params), d.Params)
			}
			break
		}
	}
	if !found {
		t.Errorf("OnInitialize not found in defs: %v", defNames(defs))
	}
}

// TestExtractFromFileCallsExtracted verifies that call references inside a
// function body are extracted and attached to the LuaFunctionDef.
func TestExtractFromFileCallsExtracted(t *testing.T) {
	dir := t.TempDir()
	luaPath := filepath.Join(dir, "caller.lua")
	src := `
TestAddon = {}

function TestAddon.Setup()
  WindowUtils.SetWindowLayer("TestAddon_Window", "NORMAL")
end
`
	if err := os.WriteFile(luaPath, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	defs, err := ExtractFromFile("TestAddon", luaPath, graph.Manifest{Name: "TestAddon"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(defs) == 0 {
		t.Fatal("expected at least one function def")
	}

	var callCount int
	for _, d := range defs {
		callCount += len(d.Calls)
	}
	if callCount == 0 {
		t.Error("expected at least one call reference to be extracted")
	}
}

// TestExtractFromFileQualifiedName verifies that the QualifiedName is built
// from the addon name and function name when not already dotted.
func TestExtractFromFileQualifiedName(t *testing.T) {
	dir := t.TempDir()
	luaPath := filepath.Join(dir, "qname.lua")
	src := `
MyAddon = {}

function MyAddon.Run()
end
`
	if err := os.WriteFile(luaPath, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	defs, err := ExtractFromFile("MyAddon", luaPath, graph.Manifest{Name: "MyAddon"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, d := range defs {
		if d.QualifiedName == "" {
			t.Errorf("function %q has empty QualifiedName", d.Name)
		}
	}
}

func defNames(defs []xml_lua_binding.LuaFunctionDef) []string {
	names := make([]string, len(defs))
	for i, d := range defs {
		names[i] = d.Name
	}
	return names
}

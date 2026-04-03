package generator

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"roraddons/tools/api_doc_gen/graph"
	"roraddons/tools/api_doc_gen/lua_parser"
)

func luaRepoRoot(t *testing.T) string {
	t.Helper()
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to locate test file path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(currentFile), "..", "..", ".."))
}

func existingLuaPath(root string, candidates ...string) string {
	for _, rel := range candidates {
		full := filepath.Join(root, filepath.FromSlash(rel))
		if _, err := os.Stat(full); err == nil {
			return rel
		}
	}
	return ""
}

func TestBuildLuaAnalysisArtifactPartyCast(t *testing.T) {
	root := luaRepoRoot(t)
	rel := existingLuaPath(root, "addons/PartyCast/PartyCast.lua", "all_addons/PartyCast/PartyCast.lua")
	if rel == "" {
		t.Skip("PartyCast.lua not present in workspace")
	}

	parsed, err := lua_parser.ParseFile("PartyCast", filepath.Join(root, filepath.FromSlash(rel)), graph.Manifest{})
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}
	artifact := buildLuaAnalysisArtifact("PartyCast", rel, parsed)

	if artifact.SchemaVersion != luaAnalysisSchemaVersion {
		t.Fatalf("unexpected schema version: %s", artifact.SchemaVersion)
	}
	if len(artifact.Functions) == 0 {
		t.Fatal("expected functions")
	}
	if len(artifact.Namespaces) == 0 {
		t.Fatal("expected namespaces")
	}

	for i, fn := range artifact.Functions {
		if fn.FunctionID == "" {
			t.Fatalf("missing function_id at index %d", i)
		}
		if fn.QualifiedName == "" {
			t.Fatalf("missing qualified_name at index %d", i)
		}
		for _, call := range fn.Calls {
			if call.CallID == "" {
				t.Fatal("missing call_id")
			}
			if call.CalleeRaw == "" || call.CalleeResolved == "" {
				t.Fatalf("call raw/resolved missing: %+v", call)
			}
		}
	}
}

func TestBuildLuaAnalysisArtifactAdvancedRenown(t *testing.T) {
	root := luaRepoRoot(t)
	rel := existingLuaPath(root, "addons/advancedrenowntrainer/AdvancedRenownTraining.lua", "all_addons/advancedrenowntrainer/AdvancedRenownTraining.lua")
	if rel == "" {
		t.Skip("AdvancedRenownTraining.lua not present in workspace")
	}

	parsed, err := lua_parser.ParseFile("AdvancedRenownTrainer", filepath.Join(root, filepath.FromSlash(rel)), graph.Manifest{})
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}
	artifact := buildLuaAnalysisArtifact("AdvancedRenownTrainer", rel, parsed)

	if len(artifact.Functions) == 0 {
		t.Fatal("expected advanced renown functions")
	}
	for _, reg := range artifact.Registrations {
		if reg.RegistrationID == "" {
			t.Fatal("missing registration id")
		}
		if reg.EventRaw == "" || reg.EventResolved == "" {
			t.Fatalf("registration event identity missing: %+v", reg)
		}
		if reg.HandlerRaw == "" || reg.HandlerResolved == "" {
			t.Fatalf("registration handler identity missing: %+v", reg)
		}
	}
}

package generator

import "testing"

func TestBuildXMLLuaLinkArtifact_ContractOnlyLinks(t *testing.T) {
	xmlArtifacts := []phase1TreeArtifact{
		{
			SchemaVersion: phase1TreeSchemaVersion,
			Addon:         "TestAddon",
			SourceFile:    "ui/Test.xml",
			Roots:         []string{"n1"},
			Nodes: []phase1TreeNode{
				{
					NodeID:       "n1",
					ParentID:     nil,
					ChildIndex:   0,
					Tag:          "Window",
					RawName:      "Main",
					ResolvedName: "Main",
					Attributes:   sortedStringMap{"name": "Main"},
					Flags:        phase1TreeFlags{IsNamed: true, IsStructural: true},
				},
				{
					NodeID:       "n2",
					ParentID:     ptr("n1"),
					ChildIndex:   0,
					Tag:          "EventHandler",
					RawName:      "",
					ResolvedName: "",
					Attributes:   sortedStringMap{"event": "OnShown", "function": "TestAddon.OnShown"},
					Flags:        phase1TreeFlags{IsStructural: true},
				},
				{
					NodeID:       "n3",
					ParentID:     ptr("n1"),
					ChildIndex:   1,
					Tag:          "Window",
					RawName:      "$parentChild",
					ResolvedName: "MainChild",
					Attributes:   sortedStringMap{"name": "$parentChild"},
					Flags:        phase1TreeFlags{IsNamed: true, IsStructural: true},
				},
			},
		},
	}

	luaArtifacts := []luaAnalysisArtifact{
		{
			SchemaVersion: luaAnalysisSchemaVersion,
			Addon:         "TestAddon",
			SourceFile:    "Test.lua",
			Functions: []luaAnalysisFunctionEntry{
				{
					FunctionID:      "f1",
					DeclaredName:    "TestAddon.OnShown",
					QualifiedName:   "TestAddon.OnShown",
					ShortName:       "OnShown",
					ScopeKind:       "global",
					DeclarationKind: "function_decl",
					Source:          luaAnalysisSourceRange{LineStart: 10, LineEnd: 20},
					Calls:           []luaAnalysisCallEntry{},
				},
				{
					FunctionID:      "f2",
					DeclaredName:    "TestAddon.Show",
					QualifiedName:   "TestAddon.Show",
					ShortName:       "Show",
					ScopeKind:       "global",
					DeclarationKind: "function_decl",
					Source:          luaAnalysisSourceRange{LineStart: 30, LineEnd: 50},
					Calls: []luaAnalysisCallEntry{
						{
							CallID:         "c1",
							CalleeRaw:      "WindowSetShowing",
							CalleeResolved: "WindowSetShowing",
							Source:         luaAnalysisSource{Line: 35},
							Arguments:      []string{"\"MainChild\"", "true"},
							IsUIAPI:        true,
						},
					},
				},
			},
			Registrations: []luaAnalysisRegistration{
				{
					RegistrationID:   "r1",
					SourceFunctionID: "f2",
					Registrar:        "WindowRegisterEventHandler",
					EventRaw:         "SystemData.Events.SAMPLE",
					EventResolved:    "SystemData.Events.SAMPLE",
					HandlerRaw:       "\"TestAddon.OnShown\"",
					HandlerResolved:  "TestAddon.OnShown",
					Scope:            "window",
					WindowRaw:        "\"MainChild\"",
					WindowResolved:   "MainChild",
					Source:           luaAnalysisSource{Line: 36},
				},
			},
		},
	}

	artifact := buildXMLLuaLinkArtifact(
		"TestAddon",
		[]string{"ui/Test.xml.tree.json"},
		[]string{"Test.lua.analysis.json"},
		xmlArtifacts,
		luaArtifacts,
	)

	if artifact.SchemaVersion != xmlLuaLinkSchemaVersion {
		t.Fatalf("expected schema %q, got %q", xmlLuaLinkSchemaVersion, artifact.SchemaVersion)
	}
	if len(artifact.HandlerLinks) != 1 {
		t.Fatalf("expected 1 handler link, got %d", len(artifact.HandlerLinks))
	}
	if len(artifact.UIUsageLinks) != 1 {
		t.Fatalf("expected 1 ui usage link, got %d", len(artifact.UIUsageLinks))
	}
	if len(artifact.RegistrationLinks) != 1 {
		t.Fatalf("expected 1 registration link, got %d", len(artifact.RegistrationLinks))
	}

	if artifact.HandlerLinks[0].LinkID != "h1" {
		t.Fatalf("expected deterministic handler link id h1, got %q", artifact.HandlerLinks[0].LinkID)
	}
	if artifact.UIUsageLinks[0].LinkID != "u1" {
		t.Fatalf("expected deterministic ui usage link id u1, got %q", artifact.UIUsageLinks[0].LinkID)
	}
	if artifact.RegistrationLinks[0].LinkID != "r1" {
		t.Fatalf("expected deterministic registration link id r1, got %q", artifact.RegistrationLinks[0].LinkID)
	}

	if artifact.UIUsageLinks[0].XML.ResolvedName != "MainChild" {
		t.Fatalf("expected resolved xml name MainChild, got %q", artifact.UIUsageLinks[0].XML.ResolvedName)
	}
	if artifact.UIUsageLinks[0].XML.RawName != "$parentChild" {
		t.Fatalf("expected raw xml name $parentChild, got %q", artifact.UIUsageLinks[0].XML.RawName)
	}

	if len(artifact.Unresolved.HandlerBindings) != 0 {
		t.Fatalf("expected no unresolved handler bindings, got %d", len(artifact.Unresolved.HandlerBindings))
	}
	if len(artifact.Unresolved.Registrations) != 0 {
		t.Fatalf("expected no unresolved registrations, got %d", len(artifact.Unresolved.Registrations))
	}
}

func TestBuildXMLLuaLinkArtifact_EmitsUnresolved(t *testing.T) {
	xmlArtifacts := []phase1TreeArtifact{
		{
			SchemaVersion: phase1TreeSchemaVersion,
			Addon:         "TestAddon",
			SourceFile:    "ui/Test.xml",
			Roots:         []string{"n1"},
			Nodes: []phase1TreeNode{
				{
					NodeID:     "n1",
					ParentID:   nil,
					Tag:        "EventHandler",
					Attributes: sortedStringMap{"event": "OnShown", "function": "Missing.Handler"},
					Flags:      phase1TreeFlags{IsStructural: true},
				},
			},
		},
	}

	artifact := buildXMLLuaLinkArtifact("TestAddon", []string{"ui/Test.xml.tree.json"}, []string{}, xmlArtifacts, []luaAnalysisArtifact{})
	if len(artifact.HandlerLinks) != 0 {
		t.Fatalf("expected no handler links, got %d", len(artifact.HandlerLinks))
	}
	if len(artifact.Unresolved.HandlerBindings) != 1 {
		t.Fatalf("expected one unresolved handler, got %d", len(artifact.Unresolved.HandlerBindings))
	}
	if artifact.Unresolved.HandlerBindings[0].Reason != "no_lua_function_match" {
		t.Fatalf("expected no_lua_function_match, got %q", artifact.Unresolved.HandlerBindings[0].Reason)
	}
}

func ptr(value string) *string {
	return &value
}

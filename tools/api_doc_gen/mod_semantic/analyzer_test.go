package mod_semantic_test

import (
	"testing"

	"roraddons/tools/api_doc_gen/graph"
	"roraddons/tools/api_doc_gen/mod_semantic"
)

// --- helpers ---

func makeTree(root *graph.ModNode) *graph.ModuleTree {
	return &graph.ModuleTree{Root: root}
}

func makeNode(tag string, attrs map[string]string, children ...*graph.ModNode) *graph.ModNode {
	if attrs == nil {
		attrs = map[string]string{}
	}
	return &graph.ModNode{Tag: tag, Attrs: attrs, Children: children}
}

// --- semantic classification ---

// TestKnownHooksAreClassified verifies that OnInitialize, OnUpdate, and
// OnShutdown sections are classified with the correct HookKind.
func TestKnownHooksAreClassified(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil, makeNode("CallFunction", map[string]string{"name": "A.Init"})),
		makeNode("OnUpdate", nil, makeNode("CallFunction", map[string]string{"name": "A.Update"})),
		makeNode("OnShutdown", nil, makeNode("CallFunction", map[string]string{"name": "A.Shutdown"})),
	)
	tree := makeTree(root)

	sem := mod_semantic.AnalyzeTree("TestAddon", tree, nil, nil)

	if len(sem.LifecycleHooks) != 3 {
		t.Fatalf("expected 3 hooks, got %d", len(sem.LifecycleHooks))
	}
	kinds := map[mod_semantic.HookKind]bool{}
	for _, h := range sem.LifecycleHooks {
		kinds[h.Kind] = true
	}
	for _, k := range []mod_semantic.HookKind{
		mod_semantic.HookKindOnInitialize,
		mod_semantic.HookKindOnUpdate,
		mod_semantic.HookKindOnShutdown,
	} {
		if !kinds[k] {
			t.Errorf("expected hook kind %q not found", k)
		}
	}
}

// TestUnknownSectionsArePreserved verifies that non-metadata sections whose
// tags are not in the known set are preserved as HookKindUnknown, not dropped.
func TestUnknownSectionsArePreserved(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil, makeNode("CallFunction", map[string]string{"name": "A.Init"})),
		makeNode("OnLoad", nil, makeNode("CallFunction", map[string]string{"name": "A.OnLoad"})),
		makeNode("CustomLifecycle", nil, makeNode("DoThing", map[string]string{"name": "X"})),
	)
	tree := makeTree(root)

	sem := mod_semantic.AnalyzeTree("TestAddon", tree, nil, nil)

	if len(sem.LifecycleHooks) != 3 {
		t.Fatalf("expected 3 hooks (1 known + 2 unknown), got %d", len(sem.LifecycleHooks))
	}
	unknownCount := 0
	for _, h := range sem.LifecycleHooks {
		if h.Kind == mod_semantic.HookKindUnknown {
			unknownCount++
			if h.RawNode == nil {
				t.Error("unknown hook missing RawNode reference")
			}
		}
	}
	if unknownCount != 2 {
		t.Errorf("expected 2 unknown hooks, got %d", unknownCount)
	}
}

// TestMetadataSectionsAreNotTreatedAsHooks verifies that Files, SavedVariables,
// Author, and Description sections are not emitted as lifecycle hooks.
func TestMetadataSectionsAreNotTreatedAsHooks(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("Files", nil, makeNode("File", map[string]string{"name": "foo.lua"})),
		makeNode("SavedVariables", nil, makeNode("SavedVariable", map[string]string{"name": "Cfg"})),
		makeNode("Author", map[string]string{"name": "AuthorName"}),
		makeNode("Description", map[string]string{"text": "desc"}),
		makeNode("OnInitialize", nil, makeNode("CallFunction", map[string]string{"name": "A.Init"})),
	)
	tree := makeTree(root)

	sem := mod_semantic.AnalyzeTree("TestAddon", tree, nil, nil)

	if len(sem.LifecycleHooks) != 1 {
		t.Fatalf("expected only OnInitialize hook (metadata sections excluded), got %d hooks", len(sem.LifecycleHooks))
	}
	if sem.LifecycleHooks[0].Kind != mod_semantic.HookKindOnInitialize {
		t.Errorf("expected OnInitialize, got %q", sem.LifecycleHooks[0].Kind)
	}
}

// TestMixedKnownAndUnknownActionsPreserved verifies that a lifecycle hook with
// both recognised and unrecognised action tags retains all of them.
func TestMixedKnownAndUnknownActionsPreserved(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil,
			makeNode("CallFunction", map[string]string{"name": "A.Init"}),
			makeNode("CreateWindow", map[string]string{"name": "A_MainWindow"}),
			makeNode("FutureAction", map[string]string{"name": "X"}),
		),
	)
	tree := makeTree(root)

	sem := mod_semantic.AnalyzeTree("TestAddon", tree, nil, nil)
	if len(sem.LifecycleHooks) != 1 {
		t.Fatalf("expected 1 hook, got %d", len(sem.LifecycleHooks))
	}
	hook := sem.LifecycleHooks[0]
	if len(hook.Actions) != 3 {
		t.Fatalf("expected 3 actions (2 known + 1 unknown), got %d", len(hook.Actions))
	}
	kinds := map[mod_semantic.ActionKind]bool{}
	for _, a := range hook.Actions {
		kinds[a.Kind] = true
	}
	if !kinds[mod_semantic.ActionKindCallFunction] {
		t.Error("CallFunction action missing")
	}
	if !kinds[mod_semantic.ActionKindCreateWindow] {
		t.Error("CreateWindow action missing")
	}
	if !kinds[mod_semantic.ActionKindUnknown] {
		t.Error("unknown action missing")
	}
}

// TestActionOrderIsPreserved verifies that the Index field on each action
// reflects document order inside the hook.
func TestActionOrderIsPreserved(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil,
			makeNode("CallFunction", map[string]string{"name": "A.First"}),
			makeNode("CreateWindow", map[string]string{"name": "A_Win"}),
			makeNode("CallFunction", map[string]string{"name": "A.Third"}),
		),
	)
	tree := makeTree(root)

	sem := mod_semantic.AnalyzeTree("TestAddon", tree, nil, nil)
	hook := sem.LifecycleHooks[0]
	for i, action := range hook.Actions {
		if action.Index != i {
			t.Errorf("action[%d].Index = %d, want %d", i, action.Index, i)
		}
	}
}

// TestRawNodeReferencesAreSet verifies that every hook and action carries a
// non-nil RawNode pointing into the original tree.
func TestRawNodeReferencesAreSet(t *testing.T) {
	actionNode := makeNode("CallFunction", map[string]string{"name": "A.Init"})
	hookNode := makeNode("OnInitialize", nil, actionNode)
	root := makeNode("UiMod", nil, hookNode)
	tree := makeTree(root)

	sem := mod_semantic.AnalyzeTree("TestAddon", tree, nil, nil)
	if len(sem.LifecycleHooks) == 0 {
		t.Fatal("no hooks produced")
	}
	hook := sem.LifecycleHooks[0]
	if hook.RawNode != hookNode {
		t.Error("hook.RawNode does not point to the original hookNode")
	}
	if len(hook.Actions) == 0 {
		t.Fatal("no actions produced")
	}
	if hook.Actions[0].RawNode != actionNode {
		t.Error("action.RawNode does not point to the original actionNode")
	}
}

// TestRawTreeIsNeverMutated verifies that analysis does not modify the
// original ModuleTree.
func TestRawTreeIsNeverMutated(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil,
			makeNode("CallFunction", map[string]string{"name": "A.Init"}),
		),
	)
	// Record original pointers and attribute values.
	origHookNode := root.Children[0]
	origActionNode := root.Children[0].Children[0]
	origAttrValue := origActionNode.Attrs["name"]

	tree := makeTree(root)
	sem := mod_semantic.AnalyzeTree("TestAddon", tree, nil, nil)

	if sem.RawTree != tree {
		t.Error("sem.RawTree should be the same pointer as the input tree")
	}
	if root.Children[0] != origHookNode {
		t.Error("root children were mutated")
	}
	if root.Children[0].Children[0] != origActionNode {
		t.Error("hook children were mutated")
	}
	if origActionNode.Attrs["name"] != origAttrValue {
		t.Errorf("action attrs mutated: got %q, want %q", origActionNode.Attrs["name"], origAttrValue)
	}
}

// --- correlation ---

// TestCallFunctionResolvesToLuaExact verifies that a CallFunction whose name
// normalises to exactly one Lua function definition gets ResolutionExact.
func TestCallFunctionResolvesToLuaExact(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil,
			makeNode("CallFunction", map[string]string{"name": "MyAddon.Initialize"}),
		),
	)
	tree := makeTree(root)

	luaIndex := mod_semantic.BuildLuaNameIndex([]string{"MyAddon.Initialize"})

	sem := mod_semantic.AnalyzeTree("MyAddon", tree, luaIndex, nil)
	action := sem.LifecycleHooks[0].Actions[0]

	if action.Resolution != mod_semantic.ResolutionExact {
		t.Errorf("expected ResolutionExact, got %q", action.Resolution)
	}
	if len(action.MatchedLuaFunctions) != 1 || action.MatchedLuaFunctions[0] != "MyAddon.Initialize" {
		t.Errorf("expected MatchedLuaFunctions=[\"MyAddon.Initialize\"], got %v", action.MatchedLuaFunctions)
	}
	if action.Confidence != "HIGH" {
		t.Errorf("expected HIGH confidence, got %q", action.Confidence)
	}
	if action.EvidenceSource != ".mod" {
		t.Errorf("expected EvidenceSource=.mod, got %q", action.EvidenceSource)
	}
}

// TestCallFunctionUnresolved verifies that a CallFunction whose name is not in
// the Lua index retains ResolutionUnresolved and remains visible.
func TestCallFunctionUnresolved(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil,
			makeNode("CallFunction", map[string]string{"name": "MyAddon.Missing"}),
		),
	)
	tree := makeTree(root)

	luaIndex := mod_semantic.BuildLuaNameIndex([]string{"OtherAddon.Initialize"})

	sem := mod_semantic.AnalyzeTree("MyAddon", tree, luaIndex, nil)
	action := sem.LifecycleHooks[0].Actions[0]

	if action.Resolution != mod_semantic.ResolutionUnresolved {
		t.Errorf("expected ResolutionUnresolved, got %q", action.Resolution)
	}
	if len(action.MatchedLuaFunctions) != 0 {
		t.Errorf("expected no matches, got %v", action.MatchedLuaFunctions)
	}
	// The action must still be visible (not dropped).
	if action.Name != "MyAddon.Missing" {
		t.Errorf("unresolved action Name lost; got %q", action.Name)
	}
}

// TestCallFunctionAmbiguous verifies that when multiple Lua functions share the
// same normalised name, the resolution is ResolutionAmbiguous.
func TestCallFunctionAmbiguous(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil,
			makeNode("CallFunction", map[string]string{"name": "A.Initialize"}),
		),
	)
	tree := makeTree(root)

	// Two qualified names share the normalised suffix "initialize".
	luaIndex := mod_semantic.BuildLuaNameIndex([]string{"A.Initialize", "B.Initialize"})

	sem := mod_semantic.AnalyzeTree("A", tree, luaIndex, nil)
	action := sem.LifecycleHooks[0].Actions[0]

	if action.Resolution != mod_semantic.ResolutionAmbiguous {
		t.Errorf("expected ResolutionAmbiguous, got %q", action.Resolution)
	}
	if len(action.MatchedLuaFunctions) != 2 {
		t.Errorf("expected 2 ambiguous matches, got %d", len(action.MatchedLuaFunctions))
	}
	if action.Confidence != "MEDIUM" {
		t.Errorf("expected MEDIUM confidence for ambiguous, got %q", action.Confidence)
	}
}

// TestCreateWindowResolvesToXMLExact verifies that a CreateWindow whose name
// matches a known XML frame gets ResolutionExact.
func TestCreateWindowResolvesToXMLExact(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil,
			makeNode("CreateWindow", map[string]string{"name": "MyAddon_MainWindow"}),
		),
	)
	tree := makeTree(root)

	xmlIndex := mod_semantic.BuildXMLFrameNameIndex([]string{"MyAddon_MainWindow", "OtherAddon_Win"})

	sem := mod_semantic.AnalyzeTree("MyAddon", tree, nil, xmlIndex)
	action := sem.LifecycleHooks[0].Actions[0]

	if action.Resolution != mod_semantic.ResolutionExact {
		t.Errorf("expected ResolutionExact, got %q", action.Resolution)
	}
	if len(action.MatchedXMLNames) != 1 || action.MatchedXMLNames[0] != "MyAddon_MainWindow" {
		t.Errorf("expected MatchedXMLNames=[\"MyAddon_MainWindow\"], got %v", action.MatchedXMLNames)
	}
	if action.Confidence != "HIGH" {
		t.Errorf("expected HIGH confidence, got %q", action.Confidence)
	}
}

// TestCreateWindowUnresolved verifies that a CreateWindow whose name is not in
// the XML frame index retains ResolutionUnresolved and remains visible.
func TestCreateWindowUnresolved(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil,
			makeNode("CreateWindow", map[string]string{"name": "MyAddon_MissingWindow"}),
		),
	)
	tree := makeTree(root)
	xmlIndex := mod_semantic.BuildXMLFrameNameIndex([]string{"OtherAddon_Win"})

	sem := mod_semantic.AnalyzeTree("MyAddon", tree, nil, xmlIndex)
	action := sem.LifecycleHooks[0].Actions[0]

	if action.Resolution != mod_semantic.ResolutionUnresolved {
		t.Errorf("expected ResolutionUnresolved, got %q", action.Resolution)
	}
	if action.Name != "MyAddon_MissingWindow" {
		t.Errorf("unresolved CreateWindow Name lost; got %q", action.Name)
	}
}

// TestNilTreeProducesEmptyResult verifies that passing a nil tree does not panic
// and returns a valid empty ModuleSemantic.
func TestNilTreeProducesEmptyResult(t *testing.T) {
	sem := mod_semantic.AnalyzeTree("NoMod", nil, nil, nil)
	if sem == nil {
		t.Fatal("expected non-nil ModuleSemantic for nil tree")
	}
	if len(sem.LifecycleHooks) != 0 {
		t.Errorf("expected 0 hooks for nil tree, got %d", len(sem.LifecycleHooks))
	}
}

// TestEmptyHookBodyProducesHookWithNoActions verifies that a known lifecycle
// hook with no children is still emitted as a LifecycleHook with empty Actions.
func TestEmptyHookBodyProducesHookWithNoActions(t *testing.T) {
	root := makeNode("UiMod", nil,
		makeNode("OnInitialize", nil), // no children
	)
	tree := makeTree(root)

	sem := mod_semantic.AnalyzeTree("TestAddon", tree, nil, nil)
	if len(sem.LifecycleHooks) != 1 {
		t.Fatalf("expected 1 hook, got %d", len(sem.LifecycleHooks))
	}
	if len(sem.LifecycleHooks[0].Actions) != 0 {
		t.Errorf("expected 0 actions for empty hook, got %d", len(sem.LifecycleHooks[0].Actions))
	}
}

// TestBuildLuaNameIndex verifies index normalisation across dot-separated names.
func TestBuildLuaNameIndex(t *testing.T) {
	idx := mod_semantic.BuildLuaNameIndex([]string{
		"MyAddon.Initialize",
		"OtherAddon.Initialize",
		"SingleName",
	})
	if len(idx["initialize"]) != 2 {
		t.Errorf("expected 2 entries for 'initialize', got %d", len(idx["initialize"]))
	}
	if len(idx["singlename"]) != 1 || idx["singlename"][0] != "SingleName" {
		t.Errorf("expected [SingleName] for 'singlename', got %v", idx["singlename"])
	}
}

// TestBuildXMLFrameNameIndex verifies basic presence index construction.
func TestBuildXMLFrameNameIndex(t *testing.T) {
	idx := mod_semantic.BuildXMLFrameNameIndex([]string{"Win1", "Win2", ""})
	if !idx["Win1"] || !idx["Win2"] {
		t.Error("expected Win1 and Win2 in index")
	}
	if idx[""] {
		t.Error("empty string should not appear in index")
	}
}

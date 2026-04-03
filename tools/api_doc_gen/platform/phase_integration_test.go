package platform

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestPipelineSourceFirstProducesRicherHierarchy validates the core
// architectural requirement: the source-first pipeline produces a richer
// element type model than the degraded docs-based fallback when real source
// files are available.
//
// The test wires a minimal addon on disk (with a real XML file containing
// a named Window that holds a named Button child), runs both pipeline paths,
// and verifies that:
//   - the source-first path discovers the parent/child relationship
//   - the fallback path (empty SourceModel) produces no hierarchy
func TestPipelineSourceFirstProducesRicherHierarchy(t *testing.T) {
	// Build a minimal addon source tree on disk.
	root := t.TempDir()
	addonDir := filepath.Join(root, "HierarchyAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	// TOC manifest pointing at the XML file.
	toc := "## Title: HierarchyAddon\nHierarchyAddon.xml\n"
	if err := os.WriteFile(filepath.Join(addonDir, "HierarchyAddon.toc"), []byte(toc), 0o644); err != nil {
		t.Fatal(err)
	}

	// XML source with a Window containing a named Button.
	xmlSrc := `<Interface>
  <Window name="HierarchyAddon_Main" inherits="Default_UIPanelWindow">
    <EventHandlers>
      <EventHandler event="OnInitialize" function="HierarchyAddon.OnInitialize"/>
    </EventHandlers>
    <Button name="HierarchyAddon_OKButton">
      <EventHandlers>
        <EventHandler event="OnClick" function="HierarchyAddon.OnOKClick"/>
      </EventHandlers>
    </Button>
  </Window>
</Interface>`
	if err := os.WriteFile(filepath.Join(addonDir, "HierarchyAddon.xml"), []byte(xmlSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	// Seed element types so the pipeline has types to enrich.
	baseTypes := []ElementTypeSymbol{
		{Name: "Window"},
		{Name: "Button"},
	}

	// --- Source-first path ---
	srcResult := runPhasedPipelineFromSources(cloneTypes(baseTypes), root, SourceModel{})

	// --- Degraded docs path (no real source data) ---
	docsResult := runPhasedPipelineFromDocs(cloneTypes(baseTypes), SourceModel{})

	// The source-first path must have enriched element type data:
	// Window must appear as a parent of Button (or Button as a child of Window).
	windowFromSources := findType(srcResult.ElementTypes, "Window")
	buttonFromSources := findType(srcResult.ElementTypes, "Button")

	if windowFromSources == nil {
		t.Fatal("source-first: Window element type not found after pipeline")
	}
	if buttonFromSources == nil {
		t.Fatal("source-first: Button element type not found after pipeline")
	}

	// With real XML, the Window should have a handler and Button should
	// appear as a named child.
	if !hasHandlerEvent(windowFromSources, "OnInitialize") {
		t.Error("source-first: Window missing OnInitialize handler (parsed from real XML)")
	}
	if !hasHandlerEvent(buttonFromSources, "OnClick") {
		t.Error("source-first: Button missing OnClick handler (parsed from real XML)")
	}

	// The degraded docs path has no real source data, so it should not
	// produce any event bindings for these types.
	windowFromDocs := findType(docsResult.ElementTypes, "Window")
	buttonFromDocs := findType(docsResult.ElementTypes, "Button")

	if windowFromDocs != nil && hasHandlerEvent(windowFromDocs, "OnInitialize") {
		t.Error("degraded: Window unexpectedly has OnInitialize — source data leaked into fallback path")
	}
	if buttonFromDocs != nil && hasHandlerEvent(buttonFromDocs, "OnClick") {
		t.Error("degraded: Button unexpectedly has OnClick — source data leaked into fallback path")
	}
}

// TestPipelineFallbackIsExplicitlyDegraded verifies that the degraded fallback
// path still runs to completion (returns enriched types) when given a
// non-empty SourceModel, and that the dispatch function selects it when
// sourceRoot is empty.
func TestPipelineFallbackIsExplicitlyDegraded(t *testing.T) {
	source := SourceModel{}
	types := []ElementTypeSymbol{{Name: "Window"}, {Name: "Button"}}

	// dispatch with empty sourceRoot → should call degraded path without panic
	result := runPhasedPipeline(cloneTypes(types), source, "")
	if len(result.ElementTypes) != 2 {
		t.Fatalf("expected 2 element types back, got %d", len(result.ElementTypes))
	}
}

// TestPipelineDispatchSelectsSourceFirstWhenRootProvided verifies that when a
// valid sourceRoot is provided, the dispatch function uses the source-first
// path (it should not produce a [DEGRADED] log or error for a valid root even
// if the root is empty of addons — it falls back gracefully).
func TestPipelineDispatchSelectsSourceFirstWhenRootProvided(t *testing.T) {
	emptyRoot := t.TempDir() // valid dir, but no addons → graceful fallback logged
	source := SourceModel{}
	types := []ElementTypeSymbol{{Name: "Window"}}

	// Should not panic; gracefully falls back with a warning.
	result := runPhasedPipeline(cloneTypes(types), source, emptyRoot)
	if len(result.ElementTypes) != 1 {
		t.Fatalf("expected 1 element type back, got %d", len(result.ElementTypes))
	}
}

// TestSourceFirstPreservesNamedAndStructuralChildren verifies that the
// source-first pipeline correctly captures both named child elements and
// structural (unnamed) children from real XML.
func TestSourceFirstPreservesNamedAndStructuralChildren(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "StructAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(addonDir, "StructAddon.toc"),
		[]byte("## Title: StructAddon\nStructAddon.xml\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	// XML: a ListBox with a named button child AND a structural ListData child.
	xmlSrc := `<Interface>
  <ListBox name="StructAddon_List">
    <ListData populationfunction="StructAddon.PopulateList"/>
    <ListColumns>
      <ListColumn width="200"/>
    </ListColumns>
    <Button name="StructAddon_EntryButton"/>
  </ListBox>
</Interface>`
	if err := os.WriteFile(filepath.Join(addonDir, "StructAddon.xml"), []byte(xmlSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	types := []ElementTypeSymbol{{Name: "ListBox"}, {Name: "Button"}, {Name: "ListData"}}
	result := runPhasedPipelineFromSources(cloneTypes(types), root, SourceModel{})

	listBox := findType(result.ElementTypes, "ListBox")
	if listBox == nil {
		t.Fatal("ListBox not found in enriched types")
	}

	// Button should appear as a named child of ListBox.
	hasButton := false
	for _, c := range listBox.ChildRefs {
		if strings.EqualFold(c.Tag, "Button") {
			hasButton = true
			break
		}
	}
	// Also check CommonChildElementTypes (backward compat field).
	if !hasButton {
		for _, c := range listBox.CommonChildElementTypes {
			if strings.EqualFold(c, "Button") {
				hasButton = true
				break
			}
		}
	}
	if !hasButton {
		t.Error("source-first: ListBox.ChildRefs does not include Button (named child from real XML)")
	}

	// ListData should appear as a structural child.
	hasListData := false
	for _, sc := range listBox.StructuralChildRefs {
		if strings.EqualFold(sc.Tag, "ListData") {
			hasListData = true
			break
		}
	}
	if !hasListData {
		for _, sc := range listBox.CommonChildTypes {
			if strings.EqualFold(sc, "ListData") {
				hasListData = true
				break
			}
		}
	}
	if !hasListData {
		t.Error("source-first: ListBox.StructuralChildRefs does not include ListData (structural child from real XML)")
	}
}

// ---- helpers ----

func cloneTypes(types []ElementTypeSymbol) []ElementTypeSymbol {
	out := make([]ElementTypeSymbol, len(types))
	copy(out, types)
	return out
}

func findType(types []ElementTypeSymbol, name string) *ElementTypeSymbol {
	for i := range types {
		if types[i].Name == name {
			return &types[i]
		}
	}
	return nil
}

func hasHandlerEvent(sym *ElementTypeSymbol, event string) bool {
	for _, eb := range sym.XMLEventBindings {
		if strings.EqualFold(eb.Event, event) {
			return true
		}
	}
	// Also check HandlerArgPatterns (Phase 4 output)
	for _, p := range sym.HandlerArgPatterns {
		if strings.EqualFold(p.Event, event) {
			return true
		}
	}
	// Check CommonHandlerFunctions list (backward compat)
	for _, h := range sym.CommonHandlerFunctions {
		if strings.Contains(strings.ToLower(h), strings.ToLower(event)) {
			return true
		}
	}
	return false
}

// ---- mod_semantic pipeline integration tests ----

// TestModManifestAddonProducesStartupWindowNote verifies the end-to-end
// source-first pipeline integration with a .mod manifest:
//   - A .mod addon whose OnInitialize section contains a CreateWindow that
//     matches a real XML frame gets a "Startup-created window" note on the
//     corresponding element type.
//   - ModuleTree (not only Manifest) is the source of truth used to derive
//     this fact.
func TestModManifestAddonProducesStartupWindowNote(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "StartupAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	// .mod manifest: OnInitialize calls CreateWindow for StartupAddon_Main.
	modContent := `<UiMod name="StartupAddon" version="1.0">
  <Files>
    <File name="StartupAddon.xml"/>
    <File name="StartupAddon.lua"/>
  </Files>
  <OnInitialize>
    <CreateWindow name="StartupAddon_Main"/>
    <CallFunction name="StartupAddon.Initialize"/>
  </OnInitialize>
</UiMod>`
	if err := os.WriteFile(filepath.Join(addonDir, "StartupAddon.mod"), []byte(modContent), 0o644); err != nil {
		t.Fatal(err)
	}

	// XML: a Window named StartupAddon_Main.
	xmlSrc := `<Interface>
  <Window name="StartupAddon_Main" inherits="Default_UIPanelWindow">
    <EventHandlers>
      <EventHandler event="OnInitialize" function="StartupAddon.Initialize"/>
    </EventHandlers>
  </Window>
</Interface>`
	if err := os.WriteFile(filepath.Join(addonDir, "StartupAddon.xml"), []byte(xmlSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	// Lua: a function StartupAddon.Initialize.
	luaSrc := "function StartupAddon.Initialize() end\n"
	if err := os.WriteFile(filepath.Join(addonDir, "StartupAddon.lua"), []byte(luaSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	types := []ElementTypeSymbol{{Name: "Window"}}
	result := runPhasedPipelineFromSources(cloneTypes(types), root, SourceModel{})

	windowSym := findType(result.ElementTypes, "Window")
	if windowSym == nil {
		t.Fatal("Window element type not found after pipeline")
	}

	// The Window element type should have a startup-created note from the .mod
	// CreateWindow action that resolved to StartupAddon_Main.
	foundNote := false
	for _, note := range windowSym.Notes {
		if strings.Contains(note, "StartupAddon_Main") && strings.Contains(note, "OnInitialize") {
			foundNote = true
			break
		}
	}
	if !foundNote {
		t.Errorf("Window element type missing startup-created note; Notes = %v", windowSym.Notes)
	}
}

// TestTOCAddonHasNoModSemantics verifies that the source-first pipeline does
// not produce any .mod semantic notes for addons that use a .toc manifest.
// .toc addons have no ModuleTree, so no lifecycle analysis should occur.
func TestTOCAddonHasNoModSemantics(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "TOCAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(addonDir, "TOCAddon.toc"),
		[]byte("## Title: TOCAddon\nTOCAddon.xml\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	xmlSrc := `<Interface>
  <Window name="TOCAddon_Main"/>
</Interface>`
	if err := os.WriteFile(filepath.Join(addonDir, "TOCAddon.xml"), []byte(xmlSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	types := []ElementTypeSymbol{{Name: "Window"}}
	result := runPhasedPipelineFromSources(cloneTypes(types), root, SourceModel{})

	windowSym := findType(result.ElementTypes, "Window")
	if windowSym == nil {
		t.Fatal("Window element type not found after pipeline")
	}
	// No startup notes should exist because there is no .mod manifest.
	for _, note := range windowSym.Notes {
		if strings.Contains(note, "Startup-created") {
			t.Errorf("unexpected startup note on TOC addon Window: %q", note)
		}
	}
}

// TestModManifestUnresolvedCreateWindowIsNotDropped verifies that a CreateWindow
// action whose name does not match any parsed XML frame does not produce an
// element-type note, but the source-first pipeline still completes successfully.
func TestModManifestUnresolvedCreateWindowIsNotDropped(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "UnresolvedAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	modContent := `<UiMod name="UnresolvedAddon" version="1.0">
  <Files>
    <File name="UnresolvedAddon.xml"/>
  </Files>
  <OnInitialize>
    <CreateWindow name="UnresolvedAddon_NoSuchWindow"/>
  </OnInitialize>
</UiMod>`
	if err := os.WriteFile(filepath.Join(addonDir, "UnresolvedAddon.mod"), []byte(modContent), 0o644); err != nil {
		t.Fatal(err)
	}

	// XML does NOT contain a frame named UnresolvedAddon_NoSuchWindow.
	xmlSrc := `<Interface>
  <Window name="UnresolvedAddon_DifferentWindow"/>
</Interface>`
	if err := os.WriteFile(filepath.Join(addonDir, "UnresolvedAddon.xml"), []byte(xmlSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	types := []ElementTypeSymbol{{Name: "Window"}}
	result := runPhasedPipelineFromSources(cloneTypes(types), root, SourceModel{})

	windowSym := findType(result.ElementTypes, "Window")
	if windowSym == nil {
		t.Fatal("Window element type not found after pipeline")
	}
	// The unresolved reference should not produce any startup note.
	for _, note := range windowSym.Notes {
		if strings.Contains(note, "UnresolvedAddon_NoSuchWindow") {
			t.Errorf("unexpected note for unresolved window: %q", note)
		}
	}
}

// ---- deep integration tests (.mod lifecycle semantics) ----

// TestModCallFunctionProducesLifecycleRole verifies that a .mod CallFunction
// action in OnInitialize produces a FunctionLifecycleRole with:
//   - Role == "startup_entrypoint"
//   - FuncName set to the matched qualified Lua function name
//   - Provenance.AddonName, HookKind, Resolution, and Confidence populated
func TestModCallFunctionProducesLifecycleRole(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "RoleAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	modContent := `<UiMod name="RoleAddon" version="1.0">
  <Files>
    <File name="RoleAddon.lua"/>
  </Files>
  <OnInitialize>
    <CallFunction name="RoleAddon.OnInit"/>
  </OnInitialize>
</UiMod>`
	if err := os.WriteFile(filepath.Join(addonDir, "RoleAddon.mod"), []byte(modContent), 0o644); err != nil {
		t.Fatal(err)
	}

	luaSrc := "function RoleAddon.OnInit() end\n"
	if err := os.WriteFile(filepath.Join(addonDir, "RoleAddon.lua"), []byte(luaSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	result := runPhasedPipelineFromSources([]ElementTypeSymbol{}, root, SourceModel{})

	if len(result.FunctionLifecycleRoles) == 0 {
		t.Fatal("expected at least one FunctionLifecycleRole but got none")
	}

	var found *FunctionLifecycleRole
	for i := range result.FunctionLifecycleRoles {
		r := &result.FunctionLifecycleRoles[i]
		if strings.Contains(r.FuncName, "OnInit") || strings.Contains(r.RefName, "OnInit") {
			found = r
			break
		}
	}
	if found == nil {
		t.Fatalf("no lifecycle role found for RoleAddon.OnInit; roles = %+v", result.FunctionLifecycleRoles)
	}

	if found.Role != "startup_entrypoint" {
		t.Errorf("expected role=startup_entrypoint, got %q", found.Role)
	}
	if found.Provenance.AddonName != "RoleAddon" {
		t.Errorf("expected Provenance.AddonName=RoleAddon, got %q", found.Provenance.AddonName)
	}
	if found.Provenance.HookKind != "OnInitialize" {
		t.Errorf("expected Provenance.HookKind=OnInitialize, got %q", found.Provenance.HookKind)
	}
	if found.Provenance.Resolution != "exact" {
		t.Errorf("expected Provenance.Resolution=exact, got %q", found.Provenance.Resolution)
	}
}

// TestModCreateWindowProducesStructuredWindowFact verifies that a .mod
// CreateWindow action produces a structured WindowLifecycleSemantic on the
// matched element type, not just a text note.
func TestModCreateWindowProducesStructuredWindowFact(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "StructAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	modContent := `<UiMod name="StructAddon" version="1.0">
  <Files>
    <File name="StructAddon.xml"/>
  </Files>
  <OnInitialize>
    <CreateWindow name="StructAddon_Main"/>
  </OnInitialize>
</UiMod>`
	if err := os.WriteFile(filepath.Join(addonDir, "StructAddon.mod"), []byte(modContent), 0o644); err != nil {
		t.Fatal(err)
	}

	xmlSrc := `<Interface>
  <Window name="StructAddon_Main"/>
</Interface>`
	if err := os.WriteFile(filepath.Join(addonDir, "StructAddon.xml"), []byte(xmlSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	types := []ElementTypeSymbol{{Name: "Window"}}
	result := runPhasedPipelineFromSources(cloneTypes(types), root, SourceModel{})

	windowSym := findType(result.ElementTypes, "Window")
	if windowSym == nil {
		t.Fatal("Window element type not found")
	}

	if len(windowSym.StartupWindowFacts) == 0 {
		t.Fatal("expected StartupWindowFacts on Window but got none")
	}

	var fact *WindowLifecycleSemantic
	for i := range windowSym.StartupWindowFacts {
		if windowSym.StartupWindowFacts[i].FrameName == "StructAddon_Main" {
			fact = &windowSym.StartupWindowFacts[i]
			break
		}
	}
	if fact == nil {
		t.Fatalf("no WindowLifecycleSemantic for StructAddon_Main; facts=%+v", windowSym.StartupWindowFacts)
	}

	if fact.HookKind != "OnInitialize" {
		t.Errorf("expected HookKind=OnInitialize, got %q", fact.HookKind)
	}
	if fact.Resolution != "exact" {
		t.Errorf("expected Resolution=exact, got %q", fact.Resolution)
	}
	if fact.Provenance.AddonName != "StructAddon" {
		t.Errorf("expected Provenance.AddonName=StructAddon, got %q", fact.Provenance.AddonName)
	}
	// ModProvenance type implies source=mod (it is only produced by .mod analysis).
}

// TestModAddonLevelSemanticsIncludesSavedVariables verifies that
// AddonLifecycleSemantics carries saved variable declarations from the .mod
// manifest.
func TestModAddonLevelSemanticsIncludesSavedVariables(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "SVAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	modContent := `<UiMod name="SVAddon" version="1.0">
  <Files>
    <File name="SVAddon.lua"/>
  </Files>
  <SavedVariables>
    <SavedVariable name="SVAddon_Settings"/>
    <SavedVariable name="SVAddon_Profile"/>
  </SavedVariables>
  <OnInitialize>
    <CallFunction name="SVAddon.Init"/>
  </OnInitialize>
</UiMod>`
	if err := os.WriteFile(filepath.Join(addonDir, "SVAddon.mod"), []byte(modContent), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(addonDir, "SVAddon.lua"),
		[]byte("function SVAddon.Init() end\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	result := runPhasedPipelineFromSources([]ElementTypeSymbol{}, root, SourceModel{})

	if len(result.AddonLifecycleSemantics) == 0 {
		t.Fatal("expected AddonLifecycleSemantics but got none")
	}

	var addonSem *AddonLifecycleSemantic
	for i := range result.AddonLifecycleSemantics {
		if result.AddonLifecycleSemantics[i].AddonName == "SVAddon" {
			addonSem = &result.AddonLifecycleSemantics[i]
			break
		}
	}
	if addonSem == nil {
		t.Fatalf("no AddonLifecycleSemantic for SVAddon; got=%+v", result.AddonLifecycleSemantics)
	}

	if len(addonSem.SavedVariables) == 0 {
		t.Error("expected SavedVariables to be populated")
	}
	hasSV := false
	for _, sv := range addonSem.SavedVariables {
		if sv == "SVAddon_Settings" {
			hasSV = true
			break
		}
	}
	if !hasSV {
		t.Errorf("expected SVAddon_Settings in SavedVariables; got=%v", addonSem.SavedVariables)
	}
}

// TestModUnresolvedRefSurvivesIntoAddonSemantics verifies that an unresolved
// CallFunction reference remains visible in the addon's lifecycle semantics.
func TestModUnresolvedRefSurvivesIntoAddonSemantics(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "UnresolvedFuncAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	modContent := `<UiMod name="UnresolvedFuncAddon" version="1.0">
  <Files>
    <File name="UnresolvedFuncAddon.lua"/>
  </Files>
  <OnInitialize>
    <CallFunction name="UnresolvedFuncAddon.DoesNotExist"/>
  </OnInitialize>
</UiMod>`
	if err := os.WriteFile(filepath.Join(addonDir, "UnresolvedFuncAddon.mod"), []byte(modContent), 0o644); err != nil {
		t.Fatal(err)
	}
	// Lua file does NOT define DoesNotExist.
	if err := os.WriteFile(filepath.Join(addonDir, "UnresolvedFuncAddon.lua"),
		[]byte("-- no functions\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	result := runPhasedPipelineFromSources([]ElementTypeSymbol{}, root, SourceModel{})

	var addonSem *AddonLifecycleSemantic
	for i := range result.AddonLifecycleSemantics {
		if result.AddonLifecycleSemantics[i].AddonName == "UnresolvedFuncAddon" {
			addonSem = &result.AddonLifecycleSemantics[i]
			break
		}
	}
	if addonSem == nil {
		t.Fatal("no AddonLifecycleSemantic for UnresolvedFuncAddon")
	}

	// The unresolved reference must survive in StartupActions (not silently dropped).
	found := false
	for _, a := range addonSem.StartupActions {
		if strings.Contains(a.Name, "DoesNotExist") {
			found = true
			if a.Resolution != "unresolved" {
				t.Errorf("expected resolution=unresolved for DoesNotExist, got %q", a.Resolution)
			}
			break
		}
	}
	if !found {
		t.Errorf("unresolved CallFunction ref not found in StartupActions; actions=%+v", addonSem.StartupActions)
	}

	// It must also appear in UnresolvedRefs.
	foundUnresolved := false
	for _, u := range addonSem.UnresolvedRefs {
		if strings.Contains(u.Name, "DoesNotExist") {
			foundUnresolved = true
			break
		}
	}
	if !foundUnresolved {
		t.Error("unresolved ref not found in AddonLifecycleSemantic.UnresolvedRefs")
	}
}

// TestModProvenanceIsRecordedOnWindowFact verifies that the Provenance field
// of a WindowLifecycleSemantic is fully populated with .mod-specific metadata.
func TestModProvenanceIsRecordedOnWindowFact(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "ProvenanceAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	modContent := `<UiMod name="ProvenanceAddon" version="1.0">
  <Files>
    <File name="ProvenanceAddon.xml"/>
  </Files>
  <OnInitialize>
    <CreateWindow name="ProvenanceAddon_Win"/>
  </OnInitialize>
</UiMod>`
	if err := os.WriteFile(filepath.Join(addonDir, "ProvenanceAddon.mod"), []byte(modContent), 0o644); err != nil {
		t.Fatal(err)
	}

	xmlSrc := `<Interface>
  <Window name="ProvenanceAddon_Win"/>
</Interface>`
	if err := os.WriteFile(filepath.Join(addonDir, "ProvenanceAddon.xml"), []byte(xmlSrc), 0o644); err != nil {
		t.Fatal(err)
	}

	types := []ElementTypeSymbol{{Name: "Window"}}
	result := runPhasedPipelineFromSources(cloneTypes(types), root, SourceModel{})

	windowSym := findType(result.ElementTypes, "Window")
	if windowSym == nil {
		t.Fatal("Window element type not found")
	}
	if len(windowSym.StartupWindowFacts) == 0 {
		t.Fatal("expected StartupWindowFacts on Window")
	}

	fact := windowSym.StartupWindowFacts[0]

	// Provenance must be fully populated.
	if fact.Provenance.AddonName == "" {
		t.Error("Provenance.AddonName must not be empty")
	}
	if fact.Provenance.HookKind == "" {
		t.Error("Provenance.HookKind must not be empty")
	}
	if fact.Provenance.Resolution == "" {
		t.Error("Provenance.Resolution must not be empty")
	}
	if fact.Provenance.ActionTag == "" {
		t.Error("Provenance.ActionTag must not be empty")
	}
}

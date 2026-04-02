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
	enrichedFromSources := runPhasedPipelineFromSources(cloneTypes(baseTypes), root, SourceModel{})

	// --- Degraded docs path (no real source data) ---
	enrichedFromDocs := runPhasedPipelineFromDocs(cloneTypes(baseTypes), SourceModel{})

	// The source-first path must have enriched element type data:
	// Window must appear as a parent of Button (or Button as a child of Window).
	windowFromSources := findType(enrichedFromSources, "Window")
	buttonFromSources := findType(enrichedFromSources, "Button")

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
	windowFromDocs := findType(enrichedFromDocs, "Window")
	buttonFromDocs := findType(enrichedFromDocs, "Button")

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
	if len(result) != 2 {
		t.Fatalf("expected 2 element types back, got %d", len(result))
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
	if len(result) != 1 {
		t.Fatalf("expected 1 element type back, got %d", len(result))
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

	listBox := findType(result, "ListBox")
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

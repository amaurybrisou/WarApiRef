package source_scan

import (
	"os"
	"path/filepath"
	"testing"
)

// TestDiscoverAddonSourcesEmpty verifies that DiscoverAddonSources returns
// an empty slice (not an error) when the source root contains no addons.
func TestDiscoverAddonSourcesEmpty(t *testing.T) {
	dir := t.TempDir()
	sources, err := DiscoverAddonSources(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sources) != 0 {
		t.Fatalf("expected 0 sources, got %d", len(sources))
	}
}

// TestDiscoverAddonSourcesWithAddon verifies that a well-formed addon directory
// (with a .mod manifest and real XML/Lua files) is discovered and its files
// are correctly categorised.
func TestDiscoverAddonSourcesWithAddon(t *testing.T) {
	root := t.TempDir()

	// Create a minimal addon directory with a .mod manifest.
	addonDir := filepath.Join(root, "TestAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	// Write a minimal .toc manifest referencing two files.
	tocContent := "## Title: TestAddon\nTestAddon.xml\nTestAddon.lua\n"
	if err := os.WriteFile(filepath.Join(addonDir, "TestAddon.toc"), []byte(tocContent), 0o644); err != nil {
		t.Fatal(err)
	}

	// Create the source files referenced by the manifest.
	if err := os.WriteFile(filepath.Join(addonDir, "TestAddon.xml"), []byte(`<Interface/>`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(addonDir, "TestAddon.lua"), []byte(`function TestAddon.Init() end`), 0o644); err != nil {
		t.Fatal(err)
	}

	sources, err := DiscoverAddonSources(root)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sources) != 1 {
		t.Fatalf("expected 1 source addon, got %d", len(sources))
	}

	src := sources[0]
	if src.AddonName != "TestAddon" {
		t.Errorf("expected addon name TestAddon, got %q", src.AddonName)
	}
	if len(src.XMLFiles) != 1 {
		t.Errorf("expected 1 XML file, got %d", len(src.XMLFiles))
	}
	if len(src.LuaFiles) != 1 {
		t.Errorf("expected 1 Lua file, got %d", len(src.LuaFiles))
	}
}

// TestDiscoverAddonSourcesMissingFiles verifies that manifest entries for
// files that do not exist on disk are silently skipped, not treated as errors.
func TestDiscoverAddonSourcesMissingFiles(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "SparseAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	// Manifest lists two files; only one actually exists.
	tocContent := "## Title: SparseAddon\nExists.lua\nMissing.lua\n"
	if err := os.WriteFile(filepath.Join(addonDir, "SparseAddon.toc"), []byte(tocContent), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(addonDir, "Exists.lua"), []byte(`function SparseAddon.Run() end`), 0o644); err != nil {
		t.Fatal(err)
	}

	sources, err := DiscoverAddonSources(root)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sources) != 1 {
		t.Fatalf("expected 1 source, got %d", len(sources))
	}
	if len(sources[0].LuaFiles) != 1 {
		t.Errorf("expected 1 Lua file (the existing one), got %d", len(sources[0].LuaFiles))
	}
}

// TestDiscoverAddonSourcesModManifestTree verifies that when an addon uses a
// .mod manifest, the returned AddonSource carries a non-nil ModuleTree that
// faithfully represents the manifest's XML tree — including any unknown
// sections that should be preserved.
func TestDiscoverAddonSourcesModManifestTree(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "ModAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	modContent := `<?xml version="1.0" encoding="UTF-8"?>
<UiMod name="ModAddon" version="1.0">
  <Files>
    <File name="ModAddon.lua"/>
  </Files>
  <OnInitialize>
    <CallFunction name="ModAddon.Init"/>
  </OnInitialize>
  <CustomSection key="discovery_test"/>
</UiMod>`
	if err := os.WriteFile(filepath.Join(addonDir, "ModAddon.mod"), []byte(modContent), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(addonDir, "ModAddon.lua"), []byte(`function ModAddon.Init() end`), 0o644); err != nil {
		t.Fatal(err)
	}

	sources, err := DiscoverAddonSources(root)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sources) != 1 {
		t.Fatalf("expected 1 source addon, got %d", len(sources))
	}

	src := sources[0]
	if src.ModuleTree == nil {
		t.Fatal("expected non-nil ModuleTree for .mod manifest addon")
	}
	if src.ModuleTree.Root == nil {
		t.Fatal("expected non-nil ModuleTree.Root")
	}
	if src.ModuleTree.Root.Tag != "UiMod" {
		t.Errorf("root tag: got %q, want UiMod", src.ModuleTree.Root.Tag)
	}

	// The unknown <CustomSection> node must be preserved in the tree.
	found := false
	for _, child := range src.ModuleTree.Root.Children {
		if child.Tag == "CustomSection" {
			found = true
			if child.Attrs["key"] != "discovery_test" {
				t.Errorf("<CustomSection> attr key: got %q, want discovery_test", child.Attrs["key"])
			}
		}
	}
	if !found {
		t.Error("<CustomSection> was discarded from ModuleTree; expected preservation")
	}

	// Manifest is still correctly populated.
	if src.Manifest.Name != "ModAddon" {
		t.Errorf("manifest name: got %q", src.Manifest.Name)
	}
}

// TestDiscoverAddonSourcesTOCManifestNoTree verifies that when an addon uses a
// .toc manifest (not .mod), ModuleTree is nil — the field is only populated
// for .mod-based addons.
func TestDiscoverAddonSourcesTOCManifestNoTree(t *testing.T) {
	root := t.TempDir()
	addonDir := filepath.Join(root, "TOCAddon")
	if err := os.MkdirAll(addonDir, 0o755); err != nil {
		t.Fatal(err)
	}

	toc := "## Title: TOCAddon\nTOCAddon.lua\n"
	if err := os.WriteFile(filepath.Join(addonDir, "TOCAddon.toc"), []byte(toc), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(addonDir, "TOCAddon.lua"), []byte(`function TOCAddon.Init() end`), 0o644); err != nil {
		t.Fatal(err)
	}

	sources, err := DiscoverAddonSources(root)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sources) != 1 {
		t.Fatalf("expected 1 source addon, got %d", len(sources))
	}
	if sources[0].ModuleTree != nil {
		t.Error("expected nil ModuleTree for .toc manifest addon")
	}
}


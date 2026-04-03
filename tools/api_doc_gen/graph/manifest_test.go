package graph

import (
	"os"
	"path/filepath"
	"testing"
)

// writeModFile is a helper that writes a .mod file into a temp directory and
// returns the full path.
func writeModFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write mod file: %v", err)
	}
	return path
}

// TestParseModManifestTreeKnownSections verifies that all standard .mod
// sections are parsed into the Manifest and that the root ModNode carries
// every classified child.
func TestParseModManifestTreeKnownSections(t *testing.T) {
	dir := t.TempDir()
	modContent := `<?xml version="1.0" encoding="UTF-8"?>
<UiMod name="MyAddon" version="1.2" date="01/01/2024">
  <Author name="Alice"/>
  <Description text="A test addon"/>
  <Files>
    <File name="MyAddon.lua"/>
    <File name="MyAddon.xml"/>
  </Files>
  <SavedVariables>
    <SavedVariable name="MyAddon_Settings"/>
  </SavedVariables>
  <OnInitialize>
    <CreateWindow name="MyAddon_Main" show="true"/>
    <CallFunction name="MyAddon.OnInit"/>
  </OnInitialize>
  <OnUpdate>
    <CallFunction name="MyAddon.OnUpdate"/>
  </OnUpdate>
  <OnShutdown>
    <CallFunction name="MyAddon.OnShutdown"/>
  </OnShutdown>
</UiMod>`

	path := writeModFile(t, dir, "MyAddon.mod", modContent)
	tree, err := ParseModManifestTree(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Root node must be present and carry the UiMod tag.
	if tree.Root == nil {
		t.Fatal("expected non-nil Root")
	}
	if tree.Root.Tag != "UiMod" {
		t.Errorf("root tag: got %q, want %q", tree.Root.Tag, "UiMod")
	}

	// Root attributes should be captured generically.
	if tree.Root.Attrs["name"] != "MyAddon" {
		t.Errorf("root attr name: got %q, want %q", tree.Root.Attrs["name"], "MyAddon")
	}
	if tree.Root.Attrs["version"] != "1.2" {
		t.Errorf("root attr version: got %q", tree.Root.Attrs["version"])
	}

	// All seven known sections should appear among root's children.
	wantSections := map[ModSectionKind]bool{
		ModSectionAuthor:       false,
		ModSectionDescription:  false,
		ModSectionFiles:        false,
		ModSectionSavedVars:    false,
		ModSectionOnInitialize: false,
		ModSectionOnUpdate:     false,
		ModSectionOnShutdown:   false,
	}
	for _, child := range tree.Root.Children {
		if child.Section != ModSectionUnknown {
			wantSections[child.Section] = true
		}
	}
	for kind, found := range wantSections {
		if !found {
			t.Errorf("section %q not classified in tree children", kind)
		}
	}

	// Manifest extraction: Files.
	m := tree.Manifest
	if m.Name != "MyAddon" {
		t.Errorf("manifest name: got %q", m.Name)
	}
	if m.Version != "1.2" {
		t.Errorf("manifest version: got %q", m.Version)
	}
	if m.Metadata["author"] != "Alice" {
		t.Errorf("manifest author: got %q", m.Metadata["author"])
	}
	if m.Metadata["description"] != "A test addon" {
		t.Errorf("manifest description: got %q", m.Metadata["description"])
	}
	if len(m.Files) != 2 {
		t.Errorf("manifest files: got %d, want 2", len(m.Files))
	}
	if len(m.SavedVariables) != 1 || m.SavedVariables[0] != "MyAddon_Settings" {
		t.Errorf("manifest saved variables: got %v", m.SavedVariables)
	}
	if len(m.CreateWindows) != 1 || m.CreateWindows[0] != "MyAddon_Main" {
		t.Errorf("manifest create windows: got %v", m.CreateWindows)
	}
	if len(m.InitializeCalls) != 1 || m.InitializeCalls[0] != "MyAddon.OnInit" {
		t.Errorf("manifest init calls: got %v", m.InitializeCalls)
	}
	if len(m.UpdateCalls) != 1 || m.UpdateCalls[0] != "MyAddon.OnUpdate" {
		t.Errorf("manifest update calls: got %v", m.UpdateCalls)
	}
	if len(m.ShutdownCalls) != 1 || m.ShutdownCalls[0] != "MyAddon.OnShutdown" {
		t.Errorf("manifest shutdown calls: got %v", m.ShutdownCalls)
	}
}

func TestParseModManifestTreeModuleFileWrapper(t *testing.T) {
	dir := t.TempDir()
	path := writeModFile(t, dir, "WrappedAddon.mod", `<?xml version="1.0" encoding="UTF-8"?>
<ModuleFile>
  <UiMod name="WrappedAddon" version="1.0">
    <Files>
      <File name="WrappedAddon.lua"/>
      <File name="WrappedAddonWindow.xml"/>
    </Files>
    <SavedVariables>
      <SavedVariable name="WrappedAddon.Settings"/>
    </SavedVariables>
    <OnInitialize>
      <CreateWindow name="WrappedAddonWindow" show="false"/>
      <CallFunction name="WrappedAddon.Initialize"/>
    </OnInitialize>
  </UiMod>
</ModuleFile>`)

	tree, err := ParseModManifestTree(path)
	if err != nil {
		t.Fatalf("ParseModManifestTree failed: %v", err)
	}
	if tree.Root == nil {
		t.Fatal("expected non-nil root")
	}
	if tree.Root.Tag != "UiMod" {
		t.Fatalf("root tag: got %q, want UiMod", tree.Root.Tag)
	}
	if tree.Manifest.Name != "WrappedAddon" {
		t.Fatalf("manifest name: got %q, want WrappedAddon", tree.Manifest.Name)
	}
	if len(tree.Manifest.Files) != 2 {
		t.Fatalf("manifest files: got %d, want 2", len(tree.Manifest.Files))
	}
	if tree.Manifest.Files[1] != "WrappedAddonWindow.xml" {
		t.Fatalf("second manifest file: got %q", tree.Manifest.Files[1])
	}
	if len(tree.Manifest.CreateWindows) != 1 || tree.Manifest.CreateWindows[0] != "WrappedAddonWindow" {
		t.Fatalf("create windows: got %v", tree.Manifest.CreateWindows)
	}
	if len(tree.Manifest.InitializeCalls) != 1 || tree.Manifest.InitializeCalls[0] != "WrappedAddon.Initialize" {
		t.Fatalf("initialize calls: got %v", tree.Manifest.InitializeCalls)
	}
}

// TestParseModManifestTreeUnknownSections verifies that sections with tags
// not in the known set are preserved in the tree with Section ==
// ModSectionUnknown and are not silently discarded.
func TestParseModManifestTreeUnknownSections(t *testing.T) {
	dir := t.TempDir()
	modContent := `<?xml version="1.0" encoding="UTF-8"?>
<UiMod name="WeirdAddon" version="0.1">
  <Files>
    <File name="Weird.lua"/>
  </Files>
  <OnLoad>
    <CallFunction name="WeirdAddon.BootstrapPhase"/>
    <RegisterHook target="SomeSystem"/>
  </OnLoad>
  <CustomExtension key="experimental" value="true"/>
</UiMod>`

	path := writeModFile(t, dir, "WeirdAddon.mod", modContent)
	tree, err := ParseModManifestTree(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tree.Root == nil {
		t.Fatal("expected non-nil Root")
	}

	// Count unknown and known sections.
	var unknownTags []string
	for _, child := range tree.Root.Children {
		if child.Section == ModSectionUnknown {
			unknownTags = append(unknownTags, child.Tag)
		}
	}

	// "OnLoad" and "CustomExtension" are unknown; they must be preserved.
	foundOnLoad := false
	foundCustom := false
	for _, tag := range unknownTags {
		if tag == "OnLoad" {
			foundOnLoad = true
		}
		if tag == "CustomExtension" {
			foundCustom = true
		}
	}
	if !foundOnLoad {
		t.Error("unknown section <OnLoad> was discarded; expected preservation")
	}
	if !foundCustom {
		t.Error("unknown section <CustomExtension> was discarded; expected preservation")
	}

	// Children of <OnLoad> must also be present in the tree.
	for _, child := range tree.Root.Children {
		if child.Tag != "OnLoad" {
			continue
		}
		if len(child.Children) != 2 {
			t.Errorf("<OnLoad> children: got %d, want 2", len(child.Children))
		}
		break
	}

	// <CustomExtension> attrs must be captured generically.
	for _, child := range tree.Root.Children {
		if child.Tag != "CustomExtension" {
			continue
		}
		if child.Attrs["key"] != "experimental" {
			t.Errorf("<CustomExtension> attr key: got %q", child.Attrs["key"])
		}
		if child.Attrs["value"] != "true" {
			t.Errorf("<CustomExtension> attr value: got %q", child.Attrs["value"])
		}
	}
}

// TestParseModManifestTreeFullHierarchy verifies that deeply nested nodes are
// preserved without flattening.
func TestParseModManifestTreeFullHierarchy(t *testing.T) {
	dir := t.TempDir()
	modContent := `<?xml version="1.0" encoding="UTF-8"?>
<UiMod name="DeepAddon" version="1.0">
  <OnInitialize>
    <Group label="phase1">
      <CallFunction name="DeepAddon.PhaseOne"/>
      <CallFunction name="DeepAddon.PhaseTwo"/>
    </Group>
  </OnInitialize>
</UiMod>`

	path := writeModFile(t, dir, "DeepAddon.mod", modContent)
	tree, err := ParseModManifestTree(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Find OnInitialize section.
	var initSection *ModNode
	for _, child := range tree.Root.Children {
		if child.Section == ModSectionOnInitialize {
			initSection = child
			break
		}
	}
	if initSection == nil {
		t.Fatal("OnInitialize section not found")
	}

	// The <Group> child must be preserved with its own children.
	if len(initSection.Children) != 1 {
		t.Fatalf("OnInitialize children: got %d, want 1", len(initSection.Children))
	}
	group := initSection.Children[0]
	if group.Tag != "Group" {
		t.Errorf("child tag: got %q, want Group", group.Tag)
	}
	if group.Attrs["label"] != "phase1" {
		t.Errorf("Group attr label: got %q", group.Attrs["label"])
	}
	if len(group.Children) != 2 {
		t.Errorf("Group children: got %d, want 2", len(group.Children))
	}
}

// TestParseModManifestTreeBackwardsCompatManifest verifies that ParseModManifest
// still returns the same Manifest it always did — the tree-based implementation
// must not regress the flat manifest extraction.
func TestParseModManifestTreeBackwardsCompatManifest(t *testing.T) {
	dir := t.TempDir()
	modContent := `<?xml version="1.0" encoding="UTF-8"?>
<UiMod name="CompatAddon" version="2.0" date="06/15/2023">
  <Author name="Bob"/>
  <Files>
    <File name="CompatAddon.lua"/>
  </Files>
  <OnInitialize>
    <CallFunction name="CompatAddon.Init"/>
  </OnInitialize>
</UiMod>`

	path := writeModFile(t, dir, "CompatAddon.mod", modContent)

	// ParseModManifest delegates to ParseModManifestTree and returns Manifest.
	manifest, err := ParseModManifest(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if manifest.Name != "CompatAddon" {
		t.Errorf("name: got %q", manifest.Name)
	}
	if manifest.Version != "2.0" {
		t.Errorf("version: got %q", manifest.Version)
	}
	if manifest.Metadata["author"] != "Bob" {
		t.Errorf("author: got %q", manifest.Metadata["author"])
	}
	if len(manifest.Files) != 1 || manifest.Files[0] != "CompatAddon.lua" {
		t.Errorf("files: got %v", manifest.Files)
	}
	if len(manifest.InitializeCalls) != 1 || manifest.InitializeCalls[0] != "CompatAddon.Init" {
		t.Errorf("init calls: got %v", manifest.InitializeCalls)
	}
}

// TestParseModManifestTreeFallback verifies that a non-parseable XML file
// produces an error rather than silently returning an empty result.
func TestParseModManifestTreeFallback(t *testing.T) {
	dir := t.TempDir()
	// Completely malformed XML that even the fallback cannot recover.
	modContent := `NOT XML AT ALL {{{{`
	path := writeModFile(t, dir, "Broken.mod", modContent)

	_, err := ParseModManifestTree(path)
	if err == nil {
		t.Error("expected an error for completely malformed content, got nil")
	}
}

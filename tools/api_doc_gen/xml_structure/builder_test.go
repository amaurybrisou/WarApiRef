package xml_structure

import (
	"os"
	"strings"
	"testing"
)

const testXML = `<?xml version="1.0" encoding="UTF-8"?>
<Interface>
  <Windows>
    <Window name="MyWindow" inherits="BaseTemplate">
      <Button name="$parentOkButton" handleinput="true">
        <EventHandlers>
          <EventHandler event="OnLButtonUp" function="MyWindow.OnOkClicked"/>
        </EventHandlers>
      </Button>
      <ListBox name="$parentList">
        <ListData populationfunction="MyWindow.PopulateList" table="Items"/>
        <ListColumns>
          <ListColumn text="Name" width="200"/>
          <ListColumn text="Value" width="100"/>
        </ListColumns>
      </ListBox>
      <Label name="$parentTitle"/>
    </Window>
  </Windows>
</Interface>`

func writeTestXML(t *testing.T, content string) string {
	t.Helper()
	tmpFile := t.TempDir() + "/test.xml"
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}
	return tmpFile
}

func TestParseFileBasicTree(t *testing.T) {
	tmpFile := writeTestXML(t, testXML)

	tree, err := ParseFile("TestAddon", tmpFile)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	if tree.Addon != "TestAddon" {
		t.Errorf("expected addon TestAddon, got %s", tree.Addon)
	}

	if len(tree.AllNodes) == 0 {
		t.Fatal("expected non-empty AllNodes")
	}

	// Find named elements
	var named []*XMLElement
	for _, node := range tree.AllNodes {
		if node.IsNamed {
			named = append(named, node)
		}
	}

	if len(named) < 3 {
		t.Errorf("expected at least 3 named elements, got %d", len(named))
		for _, n := range named {
			t.Logf("  named: %s (%s)", n.Name, n.Tag)
		}
	}

	// Check $parent expansion
	foundOkButton := false
	for _, n := range named {
		if strings.Contains(n.Name, "OkButton") {
			foundOkButton = true
			if n.Tag != "Button" {
				t.Errorf("expected OkButton to be Button, got %s", n.Tag)
			}
		}
	}
	if !foundOkButton {
		t.Error("expected $parent expansion to produce an element containing 'OkButton'")
	}
}

func TestParseFileHandlers(t *testing.T) {
	tmpFile := writeTestXML(t, testXML)

	tree, err := ParseFile("TestAddon", tmpFile)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	var handlerCount int
	for _, node := range tree.AllNodes {
		handlerCount += len(node.Handlers)
	}

	if handlerCount == 0 {
		t.Error("expected at least one handler declaration")
	}

	for _, node := range tree.AllNodes {
		for _, h := range node.Handlers {
			if h.Event == "OnLButtonUp" {
				if h.Function != "MyWindow.OnOkClicked" {
					t.Errorf("expected function MyWindow.OnOkClicked, got %s", h.Function)
				}
				return
			}
		}
	}
	t.Error("OnLButtonUp handler not found")
}

func TestParseFileStructuralChildren(t *testing.T) {
	tmpFile := writeTestXML(t, testXML)

	tree, err := ParseFile("TestAddon", tmpFile)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	var structural []*XMLElement
	for _, node := range tree.AllNodes {
		if node.IsStructural {
			structural = append(structural, node)
		}
	}

	if len(structural) == 0 {
		t.Error("expected structural children (e.g. ListData, ListColumns, ListColumn)")
	}

	foundListData := false
	for _, s := range structural {
		if s.Tag == "ListData" {
			foundListData = true
			if s.Attributes["populationfunction"] == "" {
				t.Error("expected ListData to have populationfunction attribute")
			}
		}
	}
	if !foundListData {
		t.Error("expected ListData as a structural child")
	}
}

func TestBuildCorpus(t *testing.T) {
	tmpFile := writeTestXML(t, testXML)

	tree, err := ParseFile("TestAddon", tmpFile)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	corpus := BuildCorpus([]*XMLTree{tree})

	if len(corpus.ElementTypes) == 0 {
		t.Fatal("expected non-empty ElementTypes in corpus")
	}

	windowProfile, ok := corpus.ElementTypes["Window"]
	if !ok {
		t.Fatal("expected Window element type profile")
	}
	if windowProfile.TotalCount < 1 {
		t.Errorf("expected Window count >= 1, got %d", windowProfile.TotalCount)
	}

	buttonProfile, ok := corpus.ElementTypes["Button"]
	if !ok {
		t.Fatal("expected Button element type profile")
	}
	if buttonProfile.HandlerEvents["OnLButtonUp"] == 0 {
		t.Error("expected Button to have OnLButtonUp handler event")
	}

	listBoxProfile, ok := corpus.ElementTypes["ListBox"]
	if !ok {
		t.Fatal("expected ListBox element type profile")
	}
	if len(listBoxProfile.StructuralChildren) == 0 {
		t.Error("expected ListBox to have structural children")
	}
}

func TestBuildCompositionSnippet(t *testing.T) {
	tmpFile := writeTestXML(t, testXML)

	tree, err := ParseFile("TestAddon", tmpFile)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	for _, node := range tree.AllNodes {
		if node.Tag == "ListBox" && node.IsNamed {
			snippet := BuildCompositionSnippet(node, 4)
			if snippet == "" {
				t.Error("expected non-empty composition snippet for ListBox")
			}
			if !strings.Contains(snippet, "ListData") {
				t.Error("expected ListData in composition snippet")
			}
			return
		}
	}
	t.Error("ListBox element not found")
}

func TestParentTypeTracking(t *testing.T) {
	tmpFile := writeTestXML(t, testXML)

	tree, err := ParseFile("TestAddon", tmpFile)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	corpus := BuildCorpus([]*XMLTree{tree})

	buttonProfile, ok := corpus.ElementTypes["Button"]
	if !ok {
		t.Fatal("expected Button element type profile")
	}
	if buttonProfile.ParentTags["Window"] == 0 {
		t.Error("expected Button to have Window as a parent type")
	}
}

func TestAttributeProfiles(t *testing.T) {
	tmpFile := writeTestXML(t, testXML)

	tree, err := ParseFile("TestAddon", tmpFile)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	corpus := BuildCorpus([]*XMLTree{tree})

	lcProfile, ok := corpus.ElementTypes["ListColumn"]
	if !ok {
		t.Fatal("expected ListColumn element type profile")
	}

	if _, ok := lcProfile.Attributes["text"]; !ok {
		t.Error("expected ListColumn to track 'text' attribute")
	}
	if _, ok := lcProfile.Attributes["width"]; !ok {
		t.Error("expected ListColumn to track 'width' attribute")
	}
}

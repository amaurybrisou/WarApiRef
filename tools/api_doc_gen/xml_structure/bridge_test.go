package xml_structure

import (
	"os"
	"testing"
)

const horizontalResizeXML = `<?xml version="1.0" encoding="UTF-8"?>
<Interface>
  <Windows>
    <HorizontalResizeImage name="LOL_BAR" texture="EA_Training_Specialization" textureScale="0.89">
      <Size>
        <AbsPoint x="61" y="61" />
      </Size>
      <Sizes left="7" middle="61" right="0" />
      <TexCoords>
        <Left x="72" y="241" />
        <Middle x="79" y="241" />
        <Right x="119" y="241" />
      </TexCoords>
    </HorizontalResizeImage>
  </Windows>
</Interface>`

func writeXMLFixture(t *testing.T, content string) string {
	t.Helper()
	filePath := t.TempDir() + "/fixture.xml"
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write fixture XML: %v", err)
	}
	return filePath
}

func TestToFramesAndHandlersPreservesHorizontalResizeStructuralTags(t *testing.T) {
	filePath := writeXMLFixture(t, horizontalResizeXML)

	tree, err := ParseFile("PartyCast", filePath)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	frames, _ := tree.ToFramesAndHandlers()
	if len(frames) == 0 {
		t.Fatal("expected at least one frame")
	}

	var bar *Frame
	for i := range frames {
		if frames[i].Name == "LOL_BAR" {
			bar = &frames[i]
			break
		}
	}

	if bar == nil {
		t.Fatal("expected LOL_BAR frame")
	}

	want := map[string]bool{
		"Size":      false,
		"Sizes":     false,
		"TexCoords": false,
	}

	for _, tag := range bar.StructuralChildTypes {
		if _, ok := want[tag]; ok {
			want[tag] = true
		}
	}

	for tag, found := range want {
		if !found {
			t.Fatalf("expected structural child tag %s on LOL_BAR, got %v", tag, bar.StructuralChildTypes)
		}
	}
}

func TestParseFilePreservesWindowsHierarchyForHorizontalResizeImage(t *testing.T) {
	filePath := writeXMLFixture(t, horizontalResizeXML)

	tree, err := ParseFile("PartyCast", filePath)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	var windows *XMLElement
	var bar *XMLElement
	for _, node := range tree.AllNodes {
		if node.Tag == "Windows" && windows == nil {
			windows = node
		}
		if node.Tag == "HorizontalResizeImage" && node.Name == "LOL_BAR" {
			bar = node
		}
	}

	if windows == nil {
		t.Fatal("expected Windows node to exist in Phase 1 XML tree")
	}
	if bar == nil {
		t.Fatal("expected HorizontalResizeImage LOL_BAR node")
	}
	if bar.Parent == nil || bar.Parent.Tag != "Windows" {
		t.Fatalf("expected LOL_BAR direct parent tag Windows, got %v", parentTag(bar.Parent))
	}

	childTags := make(map[string]bool)
	for _, child := range bar.Children {
		childTags[child.Tag] = true
	}

	if !childTags["Size"] {
		t.Fatalf("expected Size child under LOL_BAR, got children %v", tagsOf(bar.Children))
	}
	if !childTags["TexCoords"] {
		t.Fatalf("expected TexCoords child under LOL_BAR, got children %v", tagsOf(bar.Children))
	}
}

func parentTag(elem *XMLElement) string {
	if elem == nil {
		return ""
	}
	return elem.Tag
}

func tagsOf(children []*XMLElement) []string {
	tags := make([]string, 0, len(children))
	for _, child := range children {
		tags = append(tags, child.Tag)
	}
	return tags
}

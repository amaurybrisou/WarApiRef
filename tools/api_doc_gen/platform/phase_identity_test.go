package platform

import "testing"

func TestSynthesizeXMLTreesPreservesDerivedRawName(t *testing.T) {
	source := SourceModel{
		Frames: []FrameDoc{
			{
				Addon:      "advancedrenowntrainer",
				Name:       "AdvancedRenownTrainingPresetsWindowBackground",
				RawName:    "$parentBackground",
				Type:       "Window",
				Parent:     "AdvancedRenownTrainingPresetsWindow",
				ParentType: "Window",
				Source:     "advancedrenowntrainer/AdvancedRenownTrainingPresets.xml:13",
				Attributes: map[string]string{"name": "$parentBackground", "inherits": "EA_Window_DefaultFrame"},
			},
		},
	}

	trees := synthesizeXMLTrees(source)
	if len(trees) != 1 || len(trees[0].AllNodes) != 1 {
		t.Fatalf("expected one synthesized XML node, got %d trees and %d nodes", len(trees), len(trees[0].AllNodes))
	}

	node := trees[0].AllNodes[0]
	if node.Name != "AdvancedRenownTrainingPresetsWindowBackground" {
		t.Fatalf("expected resolved frame name to survive, got %q", node.Name)
	}
	if node.RawName != "$parentBackground" {
		t.Fatalf("expected raw frame name to survive, got %q", node.RawName)
	}
	if node.Attributes["name"] != "$parentBackground" {
		t.Fatalf("expected raw name attribute to survive, got %q", node.Attributes["name"])
	}
	if node.ParentFrameName != "AdvancedRenownTrainingPresetsWindow" {
		t.Fatalf("expected parent frame name to survive, got %q", node.ParentFrameName)
	}
}

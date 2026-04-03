package platform

import "testing"

func TestSemanticInputFromContractsPreservesRawNameAndParent(t *testing.T) {
	parentID := "n_parent"
	childID := "n_child"

	model := ContractModel{
		Root: "addons-api",
		XMLTrees: []contractXMLTree{
			{
				SchemaVersion: "phase1-tree/v1",
				Addon:         "advancedrenowntrainer",
				SourceFile:    "advancedrenowntrainer/AdvancedRenownTrainingPresets.xml",
				Nodes: []contractXMLNode{
					{
						NodeID:       parentID,
						Tag:          "Window",
						RawName:      "AdvancedRenownTrainingPresetsWindow",
						ResolvedName: "AdvancedRenownTrainingPresetsWindow",
						Attributes:   map[string]string{"name": "AdvancedRenownTrainingPresetsWindow"},
						Flags:        contractNodeFlags{IsNamed: true},
					},
					{
						NodeID:       childID,
						ParentID:     &parentID,
						Tag:          "Window",
						RawName:      "$parentBackground",
						ResolvedName: "AdvancedRenownTrainingPresetsWindowBackground",
						Attributes: map[string]string{
							"name":     "$parentBackground",
							"inherits": "EA_Window_DefaultFrame",
						},
						Flags: contractNodeFlags{IsNamed: true},
					},
				},
			},
		},
	}

	input := semanticInputFromContracts(model)
	if len(input.Frames) == 0 {
		t.Fatal("expected at least one frame in semantic input")
	}

	var child *FrameDoc
	for i := range input.Frames {
		if input.Frames[i].Name == "AdvancedRenownTrainingPresetsWindowBackground" {
			child = &input.Frames[i]
			break
		}
	}
	if child == nil {
		t.Fatal("expected child frame in semantic input")
	}
	if child.RawName != "$parentBackground" {
		t.Fatalf("expected raw frame name to survive, got %q", child.RawName)
	}
	if child.Attributes["name"] != "$parentBackground" {
		t.Fatalf("expected raw name attribute to survive, got %q", child.Attributes["name"])
	}
	if child.Parent != "AdvancedRenownTrainingPresetsWindow" {
		t.Fatalf("expected parent frame name to survive, got %q", child.Parent)
	}
}

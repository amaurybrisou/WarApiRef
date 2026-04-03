package platform

import "testing"

func TestContractConsistencyAllowsEmptyLinksWithoutLuaAnalysis(t *testing.T) {
	model := ContractModel{
		XMLTrees: []contractXMLTree{{
			SchemaVersion: "phase1-tree/v1",
			Addon:         "AbilityNotifier",
			SourceFile:    "AbilityNotifier.xml",
			Roots:         []string{"n1"},
			Nodes: []contractXMLNode{{
				NodeID: "n1",
				Tag:    "Window",
			}},
		}},
		LuaAnalyses: []contractLuaAnalysis{{
			SchemaVersion: "phase2-lua/v1",
			Addon:         "SomeOtherAddon",
			SourceFile:    "other.lua",
			Functions:     []contractLuaFunction{},
			Registrations: []contractLuaRegistration{},
		}},
		Links: []contractXMLLuaLinks{{
			SchemaVersion: "phase3-xml-lua-link/v1",
			Addon:         "AbilityNotifier",
			HandlerLinks:  []contractHandlerLink{},
		}},
	}

	if err := validateContractModelConsistency(model); err != nil {
		t.Fatalf("expected empty-link addon without lua-analysis to be allowed, got: %v", err)
	}
}

func TestContractConsistencyRequiresLuaAnalysisWhenHandlerLinksExist(t *testing.T) {
	model := ContractModel{
		XMLTrees: []contractXMLTree{{
			SchemaVersion: "phase1-tree/v1",
			Addon:         "AbilityNotifier",
			SourceFile:    "AbilityNotifier.xml",
			Roots:         []string{"n1"},
			Nodes: []contractXMLNode{{
				NodeID: "n1",
				Tag:    "Window",
			}},
		}},
		LuaAnalyses: []contractLuaAnalysis{{
			SchemaVersion: "phase2-lua/v1",
			Addon:         "SomeOtherAddon",
			SourceFile:    "other.lua",
			Functions:     []contractLuaFunction{},
			Registrations: []contractLuaRegistration{},
		}},
		Links: []contractXMLLuaLinks{{
			SchemaVersion: "phase3-xml-lua-link/v1",
			Addon:         "AbilityNotifier",
			HandlerLinks: []contractHandlerLink{{
				XML: contractHandlerXMLRef{ParentNodeID: "n1", Event: "OnInitialize", FunctionRaw: "AbHelp.Initialize"},
				Lua: contractHandlerLuaRef{QualifiedName: "AbHelp.Initialize"},
			}},
		}},
	}

	if err := validateContractModelConsistency(model); err == nil {
		t.Fatalf("expected missing lua-analysis mismatch error when handler links exist")
	}
}

func TestContractConsistencyAllowsEmptyLinksWithoutXMLTreeArtifact(t *testing.T) {
	model := ContractModel{
		XMLTrees: []contractXMLTree{{
			SchemaVersion: "phase1-tree/v1",
			Addon:         "SomeXMLAddon",
			SourceFile:    "some.xml",
			Roots:         []string{"n1"},
			Nodes: []contractXMLNode{{
				NodeID: "n1",
				Tag:    "Window",
			}},
		}},
		LuaAnalyses: []contractLuaAnalysis{{
			SchemaVersion: "phase2-lua/v1",
			Addon:         "SomeLuaAddon",
			SourceFile:    "other.lua",
			Functions:     []contractLuaFunction{},
			Registrations: []contractLuaRegistration{},
		}},
		Links: []contractXMLLuaLinks{{
			SchemaVersion: "phase3-xml-lua-link/v1",
			Addon:         "ActionBarCD",
			HandlerLinks:  []contractHandlerLink{},
		}},
	}

	if err := validateContractModelConsistency(model); err != nil {
		t.Fatalf("expected empty-link addon without xml-tree artifact to be allowed, got: %v", err)
	}
}

func TestContractConsistencyRequiresXMLTreeWhenHandlerLinksExist(t *testing.T) {
	model := ContractModel{
		XMLTrees: []contractXMLTree{{
			SchemaVersion: "phase1-tree/v1",
			Addon:         "SomeXMLAddon",
			SourceFile:    "some.xml",
			Roots:         []string{"n1"},
			Nodes: []contractXMLNode{{
				NodeID: "n1",
				Tag:    "Window",
			}},
		}},
		LuaAnalyses: []contractLuaAnalysis{{
			SchemaVersion: "phase2-lua/v1",
			Addon:         "ActionBarCD",
			SourceFile:    "actionbarcd.lua",
			Functions: []contractLuaFunction{{
				FunctionID:    "f1",
				QualifiedName: "ActionBarCD.Init",
				Calls:         []contractLuaCall{},
			}},
			Registrations: []contractLuaRegistration{},
		}},
		Links: []contractXMLLuaLinks{{
			SchemaVersion: "phase3-xml-lua-link/v1",
			Addon:         "ActionBarCD",
			HandlerLinks: []contractHandlerLink{{
				XML: contractHandlerXMLRef{ParentNodeID: "n1", Event: "OnInitialize", FunctionRaw: "ActionBarCD.Init"},
				Lua: contractHandlerLuaRef{QualifiedName: "ActionBarCD.Init"},
			}},
		}},
	}

	if err := validateContractModelConsistency(model); err == nil {
		t.Fatalf("expected missing xml-tree mismatch error when handler links exist")
	}
}

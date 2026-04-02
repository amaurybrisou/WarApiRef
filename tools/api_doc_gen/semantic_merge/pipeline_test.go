package semantic_merge

import (
	"testing"

	"roraddons/tools/api_doc_gen/xml_lua_binding"
	"roraddons/tools/api_doc_gen/xml_structure"
)

func TestRunPipelineEmpty(t *testing.T) {
	output := RunPipeline(&PipelineInput{})

	if output.XMLCorpus == nil {
		t.Error("expected non-nil XMLCorpus")
	}
	if output.Bindings == nil {
		t.Error("expected non-nil Bindings")
	}
	if output.LuaSemantic == nil {
		t.Error("expected non-nil LuaSemantic")
	}
	if output.Catalog == nil {
		t.Error("expected non-nil Catalog")
	}
}

func TestRunPipelineWithSyntheticData(t *testing.T) {
	// Build a synthetic XMLTree
	buttonElem := &xml_structure.XMLElement{
		Tag:        "Button",
		Name:       "MyButton",
		Addon:      "TestAddon",
		File:       "test.xml",
		IsNamed:    true,
		Attributes: map[string]string{"handleinput": "true"},
		Handlers: []xml_structure.XMLHandlerDecl{
			{Event: "OnLButtonUp", Function: "MyAddon.OnClick"},
		},
	}

	windowElem := &xml_structure.XMLElement{
		Tag:        "Window",
		Name:       "MyWindow",
		Addon:      "TestAddon",
		File:       "test.xml",
		IsNamed:    true,
		Attributes: map[string]string{"inherits": "BaseTemplate"},
		Children:   []*xml_structure.XMLElement{buttonElem},
	}
	buttonElem.Parent = windowElem
	buttonElem.ParentFrameName = "MyWindow"
	buttonElem.ParentFrameType = "Window"

	tree := &xml_structure.XMLTree{
		Addon:    "TestAddon",
		File:     "test.xml",
		Root:     []*xml_structure.XMLElement{windowElem},
		AllNodes: []*xml_structure.XMLElement{windowElem, buttonElem},
	}

	// Build Lua function defs
	luaDefs := []xml_lua_binding.LuaFunctionDef{
		{
			Name:          "OnClick",
			QualifiedName: "TestAddon.MyAddon.OnClick",
			Addon:         "TestAddon",
			File:          "test.lua",
			Line:          10,
			Params:        []string{"self", "flags"},
			Calls: []xml_lua_binding.LuaCallRef{
				{Name: "WindowSetShowing", Arguments: []string{"MyButton", "false"}, Line: 12},
			},
		},
	}

	output := RunPipeline(&PipelineInput{
		XMLTrees:     []*xml_structure.XMLTree{tree},
		LuaFunctions: luaDefs,
	})

	// Phase 1: Check XML corpus
	if len(output.XMLCorpus.ElementTypes) == 0 {
		t.Error("expected element types in XML corpus")
	}
	if _, ok := output.XMLCorpus.ElementTypes["Button"]; !ok {
		t.Error("expected Button in element types")
	}
	if _, ok := output.XMLCorpus.ElementTypes["Window"]; !ok {
		t.Error("expected Window in element types")
	}

	// Phase 2: Check bindings
	if len(output.Bindings.HandlerBindings) == 0 {
		t.Error("expected handler bindings")
	}

	// Phase 4: Check catalog
	if len(output.Catalog.Elements) == 0 {
		t.Error("expected enriched elements")
	}

	buttonEnriched, ok := output.Catalog.Elements["Button"]
	if !ok {
		t.Fatal("expected Button in enriched catalog")
	}

	if len(buttonEnriched.EventBindings) == 0 {
		t.Error("expected event bindings on Button")
	}

	windowEnriched, ok := output.Catalog.Elements["Window"]
	if !ok {
		t.Fatal("expected Window in enriched catalog")
	}

	if len(windowEnriched.Children) == 0 {
		t.Error("expected Window to have children")
	}
}

func TestBuildEnrichedCatalogParentChildRelationships(t *testing.T) {
	child := &xml_structure.XMLElement{
		Tag:             "Button",
		Name:            "Btn1",
		Addon:           "A",
		IsNamed:         true,
		Attributes:      map[string]string{},
		ParentFrameType: "Window",
	}
	parent := &xml_structure.XMLElement{
		Tag:        "Window",
		Name:       "Win1",
		Addon:      "A",
		IsNamed:    true,
		Attributes: map[string]string{},
		Children:   []*xml_structure.XMLElement{child},
	}
	child.Parent = parent

	tree := &xml_structure.XMLTree{
		Addon:    "A",
		File:     "t.xml",
		Root:     []*xml_structure.XMLElement{parent},
		AllNodes: []*xml_structure.XMLElement{parent, child},
	}

	corpus := xml_structure.BuildCorpus([]*xml_structure.XMLTree{tree})
	catalog := BuildEnrichedCatalog(corpus, nil, nil)

	btn, ok := catalog.Elements["Button"]
	if !ok {
		t.Fatal("expected Button in catalog")
	}
	foundWindowParent := false
	for _, p := range btn.Parents {
		if p.Tag == "Window" {
			foundWindowParent = true
		}
	}
	if !foundWindowParent {
		t.Error("expected Button to have Window as parent")
	}

	win, ok := catalog.Elements["Window"]
	if !ok {
		t.Fatal("expected Window in catalog")
	}
	foundButtonChild := false
	for _, c := range win.Children {
		if c.Tag == "Button" {
			foundButtonChild = true
		}
	}
	if !foundButtonChild {
		t.Error("expected Window to have Button as named child")
	}
}

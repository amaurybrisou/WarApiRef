package xml_lua_binding

import (
	"testing"

	"roraddons/tools/api_doc_gen/xml_structure"
)

func TestBuildBindingsResolvesHandlers(t *testing.T) {
	// Create an XML corpus with a button that has an OnClick handler
	buttonElem := &xml_structure.XMLElement{
		Tag:        "Button",
		Name:       "MyButton",
		Addon:      "TestAddon",
		File:       "test.xml",
		IsNamed:    true,
		Attributes: map[string]string{},
		Handlers: []xml_structure.XMLHandlerDecl{
			{Event: "OnLButtonUp", Function: "TestAddon.OnClick"},
		},
	}

	tree := &xml_structure.XMLTree{
		Addon:    "TestAddon",
		File:     "test.xml",
		Root:     []*xml_structure.XMLElement{buttonElem},
		AllNodes: []*xml_structure.XMLElement{buttonElem},
	}

	corpus := xml_structure.BuildCorpus([]*xml_structure.XMLTree{tree})

	// Create a Lua function index
	luaIndex := &LuaFunctionIndex{
		ByName:          make(map[string][]LuaFunctionDef),
		ByQualifiedName: make(map[string]LuaFunctionDef),
	}
	luaIndex.ByQualifiedName["TestAddon.TestAddon.OnClick"] = LuaFunctionDef{
		Name:          "OnClick",
		QualifiedName: "TestAddon.TestAddon.OnClick",
		Addon:         "TestAddon",
		File:          "test.lua",
		Params:        []string{"flags", "x", "y"},
	}
	luaIndex.ByName["onclick"] = []LuaFunctionDef{
		{
			Name:          "OnClick",
			QualifiedName: "TestAddon.TestAddon.OnClick",
			Addon:         "TestAddon",
			File:          "test.lua",
			Params:        []string{"flags", "x", "y"},
		},
	}

	bindings := BuildBindings(corpus, luaIndex)

	if len(bindings.HandlerBindings) == 0 {
		t.Fatal("expected at least one handler binding")
	}

	// Check the binding was resolved
	found := false
	for _, b := range bindings.HandlerBindings {
		if b.Event == "OnLButtonUp" {
			found = true
			if !b.Resolved {
				t.Error("expected binding to be resolved")
			}
			if len(b.LuaParams) != 3 {
				t.Errorf("expected 3 Lua params, got %d", len(b.LuaParams))
			}
		}
	}
	if !found {
		t.Error("OnLButtonUp binding not found")
	}

	// Check statistics
	if bindings.Statistics.TotalHandlers == 0 {
		t.Error("expected TotalHandlers > 0")
	}
	if bindings.Statistics.ResolvedHandlers == 0 {
		t.Error("expected ResolvedHandlers > 0")
	}
}

func TestBuildBindingsTracksUnresolved(t *testing.T) {
	buttonElem := &xml_structure.XMLElement{
		Tag:        "Button",
		Name:       "MyButton",
		Addon:      "TestAddon",
		File:       "test.xml",
		IsNamed:    true,
		Attributes: map[string]string{},
		Handlers: []xml_structure.XMLHandlerDecl{
			{Event: "OnClick", Function: "NonExistent.Handler"},
		},
	}

	tree := &xml_structure.XMLTree{
		Addon:    "TestAddon",
		File:     "test.xml",
		Root:     []*xml_structure.XMLElement{buttonElem},
		AllNodes: []*xml_structure.XMLElement{buttonElem},
	}

	corpus := xml_structure.BuildCorpus([]*xml_structure.XMLTree{tree})

	// Empty Lua index — nothing to resolve against
	luaIndex := &LuaFunctionIndex{
		ByName:          make(map[string][]LuaFunctionDef),
		ByQualifiedName: make(map[string]LuaFunctionDef),
	}

	bindings := BuildBindings(corpus, luaIndex)

	if len(bindings.UnresolvedHandlers) == 0 {
		t.Error("expected at least one unresolved handler")
	}

	if bindings.Statistics.UnresolvedCount == 0 {
		t.Error("expected UnresolvedCount > 0")
	}
}

func TestBuildBindingsAggregatesElementTypes(t *testing.T) {
	// Two buttons with different handlers
	btn1 := &xml_structure.XMLElement{
		Tag:        "Button",
		Name:       "Btn1",
		Addon:      "Addon1",
		File:       "a.xml",
		IsNamed:    true,
		Attributes: map[string]string{},
		Handlers: []xml_structure.XMLHandlerDecl{
			{Event: "OnLButtonUp", Function: "Addon1.OnClick1"},
		},
	}
	btn2 := &xml_structure.XMLElement{
		Tag:        "Button",
		Name:       "Btn2",
		Addon:      "Addon2",
		File:       "b.xml",
		IsNamed:    true,
		Attributes: map[string]string{},
		Handlers: []xml_structure.XMLHandlerDecl{
			{Event: "OnLButtonUp", Function: "Addon2.OnClick2"},
		},
	}

	tree1 := &xml_structure.XMLTree{
		Addon: "Addon1", File: "a.xml",
		Root: []*xml_structure.XMLElement{btn1}, AllNodes: []*xml_structure.XMLElement{btn1},
	}
	tree2 := &xml_structure.XMLTree{
		Addon: "Addon2", File: "b.xml",
		Root: []*xml_structure.XMLElement{btn2}, AllNodes: []*xml_structure.XMLElement{btn2},
	}

	corpus := xml_structure.BuildCorpus([]*xml_structure.XMLTree{tree1, tree2})

	luaIndex := &LuaFunctionIndex{
		ByName:          make(map[string][]LuaFunctionDef),
		ByQualifiedName: make(map[string]LuaFunctionDef),
	}

	bindings := BuildBindings(corpus, luaIndex)

	// Check element type binding for Button
	etb, ok := bindings.ElementTypeBindings["Button"]
	if !ok {
		t.Fatal("expected Button element type binding")
	}

	eb, ok := etb.EventBindings["OnLButtonUp"]
	if !ok {
		t.Fatal("expected OnLButtonUp event binding")
	}

	if eb.TotalCount != 2 {
		t.Errorf("expected TotalCount 2, got %d", eb.TotalCount)
	}

	if len(eb.ByAddon) != 2 {
		t.Errorf("expected 2 addons, got %d", len(eb.ByAddon))
	}
}

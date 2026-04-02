package lua_semantic

import (
	"testing"

	"roraddons/tools/api_doc_gen/xml_lua_binding"
)

func TestBuildSemanticCorpusResolvedHandler(t *testing.T) {
	luaIndex := &xml_lua_binding.LuaFunctionIndex{
		ByName:          make(map[string][]xml_lua_binding.LuaFunctionDef),
		ByQualifiedName: make(map[string]xml_lua_binding.LuaFunctionDef),
	}
	luaIndex.ByQualifiedName["TestAddon.OnClick"] = xml_lua_binding.LuaFunctionDef{
		Name:          "OnClick",
		QualifiedName: "TestAddon.OnClick",
		Addon:         "TestAddon",
		File:          "test.lua",
		Line:          10,
		Params:        []string{"flags", "x", "y"},
		Calls: []xml_lua_binding.LuaCallRef{
			{Name: "WindowSetShowing", Arguments: []string{"myWnd", "false"}, Line: 12},
			{Name: "LabelSetText", Arguments: []string{"myLabel", "clicked"}, Line: 13},
			{Name: "SystemData.GetPlayerName", Arguments: nil, Line: 14},
		},
	}

	bindings := &xml_lua_binding.XMLLuaBindingSet{
		ElementTypeBindings: make(map[string]*xml_lua_binding.ElementTypeBinding),
		HandlerBindings: []*xml_lua_binding.HandlerBinding{
			{
				Addon:            "TestAddon",
				FrameName:        "MyButton",
				FrameType:        "Button",
				Event:            "OnLButtonUp",
				LuaFunction:      "OnClick",
				LuaQualifiedName: "TestAddon.OnClick",
				Resolved:         true,
				Confidence:       "HIGH",
				LuaParams:        []string{"flags", "x", "y"},
			},
		},
	}

	corpus := BuildSemanticCorpus(bindings, luaIndex)

	// Check function analysis
	fa, ok := corpus.Functions["TestAddon.OnClick"]
	if !ok {
		t.Fatal("expected function analysis for TestAddon.OnClick")
	}
	if len(fa.Parameters) != 3 {
		t.Errorf("expected 3 parameters, got %d", len(fa.Parameters))
	}

	// Check parameter inference
	if fa.Parameters[0].Role != "modifier_flags" {
		t.Errorf("expected role 'modifier_flags' for flags param, got '%s'", fa.Parameters[0].Role)
	}
	if fa.Parameters[1].Type != "number" {
		t.Errorf("expected type 'number' for x param, got '%s'", fa.Parameters[1].Type)
	}

	// Check API call classification
	uiCalls := 0
	dataCalls := 0
	for _, call := range fa.Calls {
		if call.Category == "ui" {
			uiCalls++
		}
		if call.Category == "data" {
			dataCalls++
		}
	}
	if uiCalls < 2 {
		t.Errorf("expected at least 2 UI calls, got %d", uiCalls)
	}
	if dataCalls < 1 {
		t.Errorf("expected at least 1 data call, got %d", dataCalls)
	}

	// Check handler analysis
	ha := corpus.HandlerAnalyses["TestAddon|MyButton|OnLButtonUp"]
	if ha == nil {
		t.Fatal("expected handler analysis for OnLButtonUp")
	}
	if ha.HandlerCategory != "input" {
		t.Errorf("expected category 'input', got '%s'", ha.HandlerCategory)
	}
	if len(ha.InferredHandlerArgs) != 3 {
		t.Errorf("expected 3 inferred args for OnLButtonUp, got %d", len(ha.InferredHandlerArgs))
	}
}

func TestBuildSemanticCorpusUnresolved(t *testing.T) {
	luaIndex := &xml_lua_binding.LuaFunctionIndex{
		ByName:          make(map[string][]xml_lua_binding.LuaFunctionDef),
		ByQualifiedName: make(map[string]xml_lua_binding.LuaFunctionDef),
	}

	bindings := &xml_lua_binding.XMLLuaBindingSet{
		ElementTypeBindings: make(map[string]*xml_lua_binding.ElementTypeBinding),
		HandlerBindings: []*xml_lua_binding.HandlerBinding{
			{
				Addon:       "TestAddon",
				FrameName:   "MyWnd",
				FrameType:   "Window",
				Event:       "OnInitialize",
				LuaFunction: "MyWnd.OnInit",
				Resolved:    false,
				Confidence:  "LOW",
			},
		},
		UnresolvedHandlers: []*xml_lua_binding.UnresolvedHandler{
			{
				Addon:       "TestAddon",
				FrameName:   "MyWnd",
				FrameType:   "Window",
				Event:       "OnInitialize",
				LuaFunction: "MyWnd.OnInit",
			},
		},
	}

	corpus := BuildSemanticCorpus(bindings, luaIndex)

	// Should still have a handler analysis for unresolved handlers
	ha := corpus.HandlerAnalyses["TestAddon|MyWnd|OnInitialize"]
	if ha == nil {
		t.Fatal("expected handler analysis for unresolved OnInitialize handler")
	}
	if ha.HandlerCategory != "lifecycle" {
		t.Errorf("expected category 'lifecycle', got '%s'", ha.HandlerCategory)
	}
}

func TestKnownHandlerArgsCompleteness(t *testing.T) {
	// Ensure we have handler arg info for all major events
	mustHave := []string{
		"OnInitialize", "OnShutdown", "OnUpdate",
		"OnLButtonUp", "OnLButtonDown", "OnRButtonUp", "OnRButtonDown",
		"OnMouseOver", "OnMouseOut", "OnMouseWheel",
		"OnTextChanged", "OnSelChanged", "OnScrollPosChanged",
		"OnShow", "OnHide", "OnShown", "OnHidden",
		"OnKeyDown", "OnKeyUp",
		"OnMButtonUp", "OnMButtonDown",
		"OnMouseDrag",
		"OnLButtonDblClk",
		"OnClick", "OnDoubleClick",
		"OnDragStart", "OnDragStop",
		"OnEnterPressed", "OnEscapePressed",
		"OnSizeChanged",
		"OnSliderChanged",
		"OnColorSelected",
	}

	for _, event := range mustHave {
		if _, ok := knownXMLHandlerArgs[event]; !ok {
			t.Errorf("missing handler args for event: %s", event)
		}
	}
}

func TestWindowAPICallsCompleteness(t *testing.T) {
	// Ensure core window APIs are tracked
	mustHave := []string{
		"WindowSetShowing", "WindowGetShowing",
		"WindowSetAlpha", "WindowGetAlpha",
		"WindowSetDimensions", "WindowGetDimensions",
		"WindowClearAnchors", "WindowAddAnchor",
		"LabelSetText", "LabelGetText",
		"ButtonSetText", "ButtonGetText",
		"DynamicImageSetTexture",
		"TextEditBoxGetText", "TextEditBoxSetText",
		"ComboBoxSetSelectedIndex", "ComboBoxGetSelectedIndex",
		"ScrollWindowSetOffset", "ScrollWindowGetOffset",
		"StatusBarSetCurrentValue", "StatusBarSetMaximumValue",
		"CreateWindowFromTemplate", "CreateWindow", "DestroyWindow",
		"DoesWindowExist",
		"WindowSetScale", "WindowSetLayer",
		"SliderBarSetCurrentPosition",
		"ListBoxSetDisplay", "ListBoxGetItemCount",
	}

	for _, api := range mustHave {
		if !windowAPICalls[api] {
			t.Errorf("missing window API: %s", api)
		}
	}
}

func TestClassifyCallCategories(t *testing.T) {
	tests := []struct {
		name     string
		call     xml_lua_binding.LuaCallRef
		wantCat  string
		wantAPI  bool
	}{
		{"ui call", xml_lua_binding.LuaCallRef{Name: "WindowSetShowing"}, "ui", true},
		{"data call", xml_lua_binding.LuaCallRef{Name: "SystemData.GetPlayerName"}, "data", true},
		{"event call", xml_lua_binding.LuaCallRef{Name: "RegisterEventHandler"}, "event", true},
		{"utility call", xml_lua_binding.LuaCallRef{Name: "towstring"}, "utility", false},
		{"prefix-matched ui", xml_lua_binding.LuaCallRef{Name: "LabelSetFont"}, "ui", true},
		{"unknown call", xml_lua_binding.LuaCallRef{Name: "MyAddon.DoStuff"}, "unknown", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ca := classifyCall(tt.call)
			if ca.Category != tt.wantCat {
				t.Errorf("classifyCall(%q): got category %q, want %q", tt.call.Name, ca.Category, tt.wantCat)
			}
			if ca.IsAPICall != tt.wantAPI {
				t.Errorf("classifyCall(%q): got IsAPICall=%v, want %v", tt.call.Name, ca.IsAPICall, tt.wantAPI)
			}
		})
	}
}

func TestHandlerCategories(t *testing.T) {
	tests := []struct {
		event    string
		wantCat  string
	}{
		{"OnInitialize", "lifecycle"},
		{"OnShutdown", "lifecycle"},
		{"OnUpdate", "lifecycle"},
		{"OnLButtonUp", "input"},
		{"OnMouseOver", "input"},
		{"OnKeyDown", "input"},
		{"OnTextChanged", "data"},
		{"OnSelChanged", "data"},
		{"OnSizeChanged", "layout"},
		{"OnEditFocusGained", "focus"},
		{"OnTooltip", "tooltip"},
		{"OnCustomEvent", "custom"},
	}

	for _, tt := range tests {
		t.Run(tt.event, func(t *testing.T) {
			got := categorizeHandler(tt.event)
			if got != tt.wantCat {
				t.Errorf("categorizeHandler(%q) = %q, want %q", tt.event, got, tt.wantCat)
			}
		})
	}
}

func TestNilInputs(t *testing.T) {
	// Should not panic with nil bindings
	corpus := BuildSemanticCorpus(nil, nil)
	if corpus == nil {
		t.Fatal("expected non-nil corpus")
	}
	if len(corpus.Functions) != 0 {
		t.Errorf("expected empty functions map")
	}
}

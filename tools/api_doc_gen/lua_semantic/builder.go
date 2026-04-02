package lua_semantic

import (
	"strings"

	"roraddons/tools/api_doc_gen/xml_lua_binding"
)

// knownXMLHandlerArgs maps XML event names to their known handler argument signatures.
// This is the canonical source of truth for handler callback parameters.
// Events are derived from observed WAR addon XML usage and engine documentation.
var knownXMLHandlerArgs = map[string][]CallbackParam{
	// Lifecycle hooks – no engine-supplied arguments.
	"OnInitialize": {},
	"OnShutdown":   {},
	"OnShow":       {},
	"OnHide":       {},
	"OnShown":      {},
	"OnHidden":     {},

	// Per-frame update hook – elapsed time in seconds.
	"OnUpdate": {
		{Position: 0, Name: "elapsed", Type: "number", Role: "time_delta", Source: "xml_handler"},
	},

	// Mouse button events — flags + cursor position.
	"OnLButtonUp": {
		{Position: 0, Name: "flags", Type: "number", Role: "modifier_flags", Source: "xml_handler"},
		{Position: 1, Name: "x", Type: "number", Role: "mouse_x", Source: "xml_handler"},
		{Position: 2, Name: "y", Type: "number", Role: "mouse_y", Source: "xml_handler"},
	},
	"OnLButtonDown": {
		{Position: 0, Name: "flags", Type: "number", Role: "modifier_flags", Source: "xml_handler"},
		{Position: 1, Name: "x", Type: "number", Role: "mouse_x", Source: "xml_handler"},
		{Position: 2, Name: "y", Type: "number", Role: "mouse_y", Source: "xml_handler"},
	},
	"OnLButtonDblClk": {
		{Position: 0, Name: "flags", Type: "number", Role: "modifier_flags", Source: "xml_handler"},
		{Position: 1, Name: "x", Type: "number", Role: "mouse_x", Source: "xml_handler"},
		{Position: 2, Name: "y", Type: "number", Role: "mouse_y", Source: "xml_handler"},
	},
	"OnRButtonUp": {
		{Position: 0, Name: "flags", Type: "number", Role: "modifier_flags", Source: "xml_handler"},
		{Position: 1, Name: "x", Type: "number", Role: "mouse_x", Source: "xml_handler"},
		{Position: 2, Name: "y", Type: "number", Role: "mouse_y", Source: "xml_handler"},
	},
	"OnRButtonDown": {
		{Position: 0, Name: "flags", Type: "number", Role: "modifier_flags", Source: "xml_handler"},
		{Position: 1, Name: "x", Type: "number", Role: "mouse_x", Source: "xml_handler"},
		{Position: 2, Name: "y", Type: "number", Role: "mouse_y", Source: "xml_handler"},
	},
	"OnMButtonUp": {
		{Position: 0, Name: "flags", Type: "number", Role: "modifier_flags", Source: "xml_handler"},
		{Position: 1, Name: "x", Type: "number", Role: "mouse_x", Source: "xml_handler"},
		{Position: 2, Name: "y", Type: "number", Role: "mouse_y", Source: "xml_handler"},
	},
	"OnMButtonDown": {
		{Position: 0, Name: "flags", Type: "number", Role: "modifier_flags", Source: "xml_handler"},
		{Position: 1, Name: "x", Type: "number", Role: "mouse_x", Source: "xml_handler"},
		{Position: 2, Name: "y", Type: "number", Role: "mouse_y", Source: "xml_handler"},
	},
	"OnMouseDown": {
		{Position: 0, Name: "button", Type: "number", Role: "button_id", Source: "xml_handler"},
	},
	"OnMouseUp": {
		{Position: 0, Name: "button", Type: "number", Role: "button_id", Source: "xml_handler"},
	},
	"OnClick":       {},
	"OnDoubleClick": {},

	// Mouse motion / hover events.
	"OnMouseOver":    {},
	"OnMouseOut":     {},
	"OnMouseOverEnd": {},
	"OnMouseWheel": {
		{Position: 0, Name: "x", Type: "number", Role: "mouse_x", Source: "xml_handler"},
		{Position: 1, Name: "y", Type: "number", Role: "mouse_y", Source: "xml_handler"},
		{Position: 2, Name: "delta", Type: "number", Role: "wheel_delta", Source: "xml_handler"},
	},

	// Drag events.
	"OnMouseDrag": {
		{Position: 0, Name: "flags", Type: "number", Role: "modifier_flags", Source: "xml_handler"},
		{Position: 1, Name: "x", Type: "number", Role: "mouse_x", Source: "xml_handler"},
		{Position: 2, Name: "y", Type: "number", Role: "mouse_y", Source: "xml_handler"},
	},
	"OnDragStart":   {},
	"OnDragStop":    {},
	"OnReceiveDrag": {},

	// Keyboard events.
	"OnEnterPressed":  {},
	"OnEscapePressed": {},
	"OnKeyDown": {
		{Position: 0, Name: "key", Type: "number", Role: "key_code", Source: "xml_handler"},
	},
	"OnKeyUp": {
		{Position: 0, Name: "key", Type: "number", Role: "key_code", Source: "xml_handler"},
	},

	// Value / text change events.
	"OnTextChanged": {
		{Position: 0, Name: "text", Type: "wstring", Role: "current_text", Source: "xml_handler"},
	},
	"OnSelChanged": {
		{Position: 0, Name: "index", Type: "number", Role: "selected_index", Source: "xml_handler"},
	},
	"OnSelectionChanged": {
		{Position: 0, Name: "selectedRow", Type: "number", Role: "selected_row", Source: "xml_handler"},
	},
	"OnScrollPosChanged": {
		{Position: 0, Name: "scrollPos", Type: "number", Role: "scroll_position", Source: "xml_handler"},
	},
	"OnScrollChanged": {},
	"OnValueChanged": {
		{Position: 0, Name: "value", Type: "number", Role: "current_value", Source: "xml_handler"},
	},

	// Resize / move events.
	"OnSizeChanged": {
		{Position: 0, Name: "width", Type: "number", Role: "new_width", Source: "xml_handler"},
		{Position: 1, Name: "height", Type: "number", Role: "new_height", Source: "xml_handler"},
	},
	"OnMoveStop":    {},
	"OnResizeBegin": {},
	"OnResizeEnd":   {},
	"OnMove":        {},

	// Edit box events.
	"OnCursorChanged":          {},
	"OnInputLanguageChanged":   {},
	"OnTabPressed":             {},
	"OnEditFocusGained":        {},
	"OnEditFocusLost":          {},

	// Tooltip / interaction events.
	"OnTooltip":       {},
	"OnAlphaChanged":  {},
	"OnTintChanged":   {},

	// Color picker events.
	"OnColorSelected": {
		{Position: 0, Name: "r", Type: "number", Role: "red_component", Source: "xml_handler"},
		{Position: 1, Name: "g", Type: "number", Role: "green_component", Source: "xml_handler"},
		{Position: 2, Name: "b", Type: "number", Role: "blue_component", Source: "xml_handler"},
	},

	// Slider events.
	"OnSliderChanged": {
		{Position: 0, Name: "value", Type: "number", Role: "slider_value", Source: "xml_handler"},
	},
}

// windowAPICalls lists function names that are known UI/Window API calls.
// This is a comprehensive list from observed WAR addon usage.
var windowAPICalls = map[string]bool{
	// Core Window API
	"WindowSetShowing":              true,
	"WindowGetShowing":              true,
	"WindowSetAlpha":                true,
	"WindowGetAlpha":                true,
	"WindowSetDimensions":           true,
	"WindowGetDimensions":           true,
	"WindowSetTintColor":            true,
	"WindowGetTintColor":            true,
	"WindowGetId":                   true,
	"WindowSetId":                   true,
	"WindowSetScale":                true,
	"WindowGetScale":                true,
	"WindowSetHandleInput":          true,
	"WindowSetLayer":                true,
	"WindowGetLayer":                true,
	"WindowSetMovable":              true,
	"WindowSetResizable":            true,
	"WindowSetOffsetFromParent":     true,
	"WindowGetOffsetFromParent":     true,
	"WindowGetParent":               true,
	"WindowSetParent":               true,
	"WindowGetScreenPosition":       true,
	"WindowSetFontAlpha":            true,
	"WindowForceProcessAnchors":     true,

	// Label API
	"LabelSetText":                  true,
	"LabelGetText":                  true,
	"LabelSetTextColor":             true,
	"LabelGetTextColor":             true,
	"LabelSetFont":                  true,
	"LabelGetFont":                  true,

	// Button API
	"ButtonSetText":                 true,
	"ButtonGetText":                 true,
	"ButtonSetStayDownFlag":         true,
	"ButtonGetStayDownFlag":         true,
	"ButtonSetDisabledFlag":         true,
	"ButtonGetDisabledFlag":         true,
	"ButtonSetPressedFlag":          true,
	"ButtonGetPressedFlag":          true,

	// Image / Texture API
	"DynamicImageSetTexture":        true,
	"DynamicImageSetTextureSlice":   true,
	"DynamicImageSetTextureDimensions": true,
	"DynamicImageSetTextureScale":   true,
	"DynamicImageSetRotation":       true,

	// TextEditBox API
	"TextEditBoxGetText":            true,
	"TextEditBoxSetText":            true,
	"TextEditBoxSetFocus":           true,
	"TextEditBoxInsertText":         true,
	"TextEditBoxSetMaxLength":       true,

	// ComboBox API
	"ComboBoxSetSelectedIndex":      true,
	"ComboBoxGetSelectedIndex":      true,
	"ComboBoxClearMenuItems":        true,
	"ComboBoxAddMenuItem":           true,

	// ScrollWindow API
	"ScrollWindowSetOffset":         true,
	"ScrollWindowGetOffset":         true,
	"ScrollWindowUpdateScrollRect":  true,
	"ScrollWindowSetScrollbar":      true,

	// StatusBar API
	"StatusBarSetCurrentValue":      true,
	"StatusBarGetCurrentValue":      true,
	"StatusBarSetMaximumValue":      true,
	"StatusBarGetMaximumValue":      true,
	"StatusBarSetMinimumValue":      true,

	// Window creation/destruction
	"CreateWindowFromTemplate":      true,
	"CreateWindow":                  true,
	"DestroyWindow":                 true,
	"DoesWindowExist":               true,

	// Window anchoring
	"WindowClearAnchors":            true,
	"WindowAddAnchor":               true,

	// SliderBar API
	"SliderBarSetCurrentPosition":   true,
	"SliderBarGetCurrentPosition":   true,
	"SliderBarSetMinMaxValues":      true,

	// ListBox API
	"ListBoxSetDisplay":             true,
	"ListBoxGetItemCount":           true,
	"ListBoxGetSelectedRow":         true,
	"ListBoxSetSelectedRow":         true,

	// Map API
	"MapSetMapView":                 true,
	"MapSetZoomFactor":              true,

	// Cursor API
	"Cursor.Clear":                  true,
	"Cursor.Set":                    true,
	"Cursor.PopEntry":               true,

	// Event registration API
	"RegisterEventHandler":          true,
	"UnregisterEventHandler":        true,
	"BroadcastEvent":                true,
	"WindowRegisterEventHandler":    true,
	"WindowRegisterCoreEventHandler": true,
}

// knownDataAPIPrefixes lists namespace prefixes for platform data API calls.
var knownDataAPIPrefixes = []string{
	"SystemData.",
	"GameData.",
	"Player.",
	"PlayerUnit.",
	"AuctionData.",
	"BattlefieldData.",
	"CraftingSystem.",
	"EquipmentData.",
	"GuildVaultData.",
	"MailData.",
	"MarketplaceData.",
	"QuestData.",
	"SocialData.",
	"TradeSkillData.",
	"Cursor.",
}

// knownEventAPIs lists function names related to event handling.
var knownEventAPIs = map[string]bool{
	"RegisterEventHandler":          true,
	"UnregisterEventHandler":        true,
	"BroadcastEvent":                true,
	"WindowRegisterEventHandler":    true,
	"WindowRegisterCoreEventHandler": true,
}

// knownUtilityAPIs lists common utility functions.
var knownUtilityAPIs = map[string]bool{
	"towstring":          true,
	"tostring":           true,
	"tonumber":           true,
	"d":                  true, // debug print
	"table.insert":       true,
	"table.remove":       true,
	"table.sort":         true,
	"pairs":              true,
	"ipairs":             true,
	"math.floor":         true,
	"math.mod":           true,
	"math.ceil":          true,
	"math.abs":           true,
	"math.max":           true,
	"math.min":           true,
	"string.format":      true,
	"string.find":        true,
	"string.sub":         true,
	"string.gsub":        true,
	"string.len":         true,
	"wstring.format":     true,
	"wstring.len":        true,
	"wstring.sub":        true,
	"wstring.find":       true,
	"type":               true,
	"pcall":              true,
	"error":              true,
	"assert":             true,
	"unpack":             true,
	"select":             true,
	"GetStringFromTable": true,
	"GetStringFormat":    true,
	"GetStringFormatFromTable": true,
}

// BuildSemanticCorpus performs Phase 3 analysis: deep semantic analysis of
// Lua functions that are correlated with XML elements.
func BuildSemanticCorpus(bindings *xml_lua_binding.XMLLuaBindingSet, luaIndex *xml_lua_binding.LuaFunctionIndex) *LuaSemanticCorpus {
	corpus := &LuaSemanticCorpus{
		Functions:       make(map[string]*FunctionAnalysis),
		HandlerAnalyses: make(map[string]*HandlerAnalysis),
	}

	if bindings == nil {
		return corpus
	}

	// Phase 3a: Analyze all resolved handler functions
	for _, binding := range bindings.HandlerBindings {
		if !binding.Resolved {
			continue
		}

		// Build or retrieve function analysis
		fa := corpus.ensureFunction(binding.LuaQualifiedName, binding.LuaFunction, luaIndex)

		// Build handler-specific analysis
		handlerKey := binding.Addon + "|" + binding.FrameName + "|" + binding.Event
		ha := &HandlerAnalysis{
			FunctionAnalysis: fa,
			FrameName:        binding.FrameName,
			FrameType:        binding.FrameType,
			Event:            binding.Event,
			Addon:            binding.Addon,
		}

		// Infer handler arguments from known event signatures
		if knownArgs, ok := knownXMLHandlerArgs[binding.Event]; ok {
			ha.InferredHandlerArgs = knownArgs
			if len(knownArgs) > 0 {
				ha.HandlerConfidence = "MEDIUM"
			} else {
				ha.HandlerConfidence = "HIGH" // We're confident there are no args
			}
		} else {
			// Unknown event — attempt to infer from actual function params
			ha.InferredHandlerArgs = inferArgsFromFunctionParams(fa, binding.Event)
			ha.HandlerConfidence = "LOW"
		}

		// Categorize handler
		ha.HandlerCategory = categorizeHandler(binding.Event)

		corpus.HandlerAnalyses[handlerKey] = ha
	}

	// Phase 3b: Analyze handler functions referenced in unresolved handlers
	// (these exist in the binding set but we couldn't find their Lua source —
	// still worth recording what event they're supposed to handle).
	for _, unresolved := range bindings.UnresolvedHandlers {
		handlerKey := unresolved.Addon + "|" + unresolved.FrameName + "|" + unresolved.Event
		if _, exists := corpus.HandlerAnalyses[handlerKey]; exists {
			continue // already analyzed via resolved path
		}
		ha := &HandlerAnalysis{
			FunctionAnalysis: &FunctionAnalysis{
				Name:               unresolved.LuaFunction,
				AnalysisConfidence: "LOW",
				Notes:              []string{"Function definition not found in Lua source"},
			},
			FrameName: unresolved.FrameName,
			FrameType: unresolved.FrameType,
			Event:     unresolved.Event,
			Addon:     unresolved.Addon,
		}
		if knownArgs, ok := knownXMLHandlerArgs[unresolved.Event]; ok {
			ha.InferredHandlerArgs = knownArgs
			ha.HandlerConfidence = "LOW" // We know the event args but can't verify the function
		} else {
			ha.HandlerConfidence = "LOW"
		}
		ha.HandlerCategory = categorizeHandler(unresolved.Event)
		corpus.HandlerAnalyses[handlerKey] = ha
	}

	return corpus
}

// inferArgsFromFunctionParams attempts to infer handler arguments from the
// actual Lua function parameter names when no canonical signature is known.
func inferArgsFromFunctionParams(fa *FunctionAnalysis, event string) []CallbackParam {
	if fa == nil || len(fa.Parameters) == 0 {
		return nil
	}
	var params []CallbackParam
	for _, p := range fa.Parameters {
		if p.Name == "self" || p.Name == "..." {
			continue // skip self and varargs
		}
		params = append(params, CallbackParam{
			Position: p.Position,
			Name:     p.Name,
			Type:     p.Type,
			Role:     p.Role,
			Source:   "inferred_from_lua",
		})
	}
	return params
}

// ensureFunction builds or retrieves the analysis for a function.
func (c *LuaSemanticCorpus) ensureFunction(qualifiedName, shortName string, luaIndex *xml_lua_binding.LuaFunctionIndex) *FunctionAnalysis {
	if fa, ok := c.Functions[qualifiedName]; ok {
		return fa
	}

	fa := &FunctionAnalysis{
		Name:          shortName,
		QualifiedName: qualifiedName,
	}

	// Look up the function definition for detailed analysis
	if luaIndex == nil {
		fa.AnalysisConfidence = "LOW"
		fa.Notes = append(fa.Notes, "No Lua index available")
		c.Functions[qualifiedName] = fa
		return fa
	}

	if def, ok := luaIndex.ByQualifiedName[qualifiedName]; ok {
		fa.Addon = def.Addon
		fa.File = def.File
		fa.Line = def.Line
		fa.EndLine = def.EndLine

		// Analyze parameters
		for i, param := range def.Params {
			pa := ParameterAnalysis{
				Name:       param,
				Position:   i,
				Role:       inferParamRole(param),
				Type:       inferParamType(param),
				Confidence: inferParamConfidence(param),
			}
			fa.Parameters = append(fa.Parameters, pa)
		}

		// Analyze downstream calls
		for _, call := range def.Calls {
			ca := classifyCall(call)
			fa.Calls = append(fa.Calls, ca)
		}

		// Determine overall confidence from evidence quality
		fa.EvidenceCount = len(fa.Parameters) + len(fa.Calls)
		if fa.EvidenceCount >= 3 {
			fa.AnalysisConfidence = "MEDIUM"
		} else if fa.EvidenceCount > 0 {
			fa.AnalysisConfidence = "LOW"
		} else {
			fa.AnalysisConfidence = "LOW"
			fa.Notes = append(fa.Notes, "Function has no parameters and makes no calls")
		}
	} else {
		fa.AnalysisConfidence = "LOW"
		fa.Notes = append(fa.Notes, "Function definition not found in Lua source")
	}

	c.Functions[qualifiedName] = fa
	return fa
}

// classifyCall categorises a call reference as UI, data, event, utility, or unknown.
func classifyCall(call xml_lua_binding.LuaCallRef) CallAnalysis {
	ca := CallAnalysis{
		Callee:    call.Name,
		Arguments: call.Arguments,
		Line:      call.Line,
	}

	// Check event APIs first (before window APIs, since some like
	// RegisterEventHandler are in both maps)
	if knownEventAPIs[call.Name] {
		ca.IsAPICall = true
		ca.Category = "event"
		return ca
	}

	if windowAPICalls[call.Name] {
		ca.IsWindowCall = true
		ca.IsAPICall = true
		ca.Category = "ui"
		return ca
	}

	if isDataAPICall(call.Name) {
		ca.IsAPICall = true
		ca.Category = "data"
		return ca
	}

	if knownUtilityAPIs[call.Name] {
		ca.Category = "utility"
		return ca
	}

	// Check if it looks like a window API by prefix matching
	for prefix := range windowFunctionPrefixSet {
		if strings.HasPrefix(call.Name, prefix) {
			ca.IsWindowCall = true
			ca.IsAPICall = true
			ca.Category = "ui"
			return ca
		}
	}

	ca.Category = "unknown"
	return ca
}

// windowFunctionPrefixSet groups window API function prefixes.
var windowFunctionPrefixSet = map[string]bool{
	"Window":        true,
	"Label":         true,
	"Button":        true,
	"DynamicImage":  true,
	"TextEditBox":   true,
	"ComboBox":      true,
	"ScrollWindow":  true,
	"StatusBar":     true,
	"SliderBar":     true,
	"ListBox":       true,
	"FullResizeImage": true,
	"AnimatedImage": true,
	"MapDisplay":    true,
	"HorizontalScrollbar": true,
	"VerticalScrollbar": true,
}

// inferParamRole infers a parameter's role from its name.
func inferParamRole(name string) string {
	lower := strings.ToLower(name)
	switch {
	case lower == "self":
		return "self"
	case lower == "elapsed" || lower == "timepassed" || lower == "dt":
		return "time_delta"
	case lower == "x":
		return "mouse_x"
	case lower == "y":
		return "mouse_y"
	case lower == "delta":
		return "wheel_delta"
	case lower == "button":
		return "button_id"
	case lower == "key":
		return "key_code"
	case lower == "flags":
		return "modifier_flags"
	case strings.Contains(lower, "callback") || strings.Contains(lower, "func") || strings.Contains(lower, "handler"):
		return "callback"
	case strings.Contains(lower, "flag") || strings.Contains(lower, "bool") || strings.Contains(lower, "enable"):
		return "flag"
	case strings.Contains(lower, "name") || strings.Contains(lower, "id") || strings.Contains(lower, "key"):
		return "identifier"
	case strings.Contains(lower, "data") || strings.Contains(lower, "table") || strings.Contains(lower, "list"):
		return "data"
	case strings.Contains(lower, "index") || strings.Contains(lower, "count") || strings.Contains(lower, "num") || strings.Contains(lower, "row") || strings.Contains(lower, "col"):
		return "number"
	case strings.Contains(lower, "text") || strings.Contains(lower, "str") || strings.Contains(lower, "label") || strings.Contains(lower, "title"):
		return "text"
	case strings.Contains(lower, "color") || strings.Contains(lower, "tint"):
		return "color"
	case strings.Contains(lower, "alpha") || strings.Contains(lower, "opacity"):
		return "alpha"
	case strings.Contains(lower, "width") || strings.Contains(lower, "height") || strings.Contains(lower, "size"):
		return "dimension"
	default:
		return "unknown"
	}
}

// inferParamType infers a parameter's type from its name.
func inferParamType(name string) string {
	lower := strings.ToLower(name)
	switch {
	case lower == "self":
		return "table"
	case strings.Contains(lower, "callback") || strings.Contains(lower, "func") || strings.Contains(lower, "handler"):
		return "function"
	case strings.Contains(lower, "flag") || strings.Contains(lower, "bool") || strings.Contains(lower, "enable") || strings.Contains(lower, "visible") || strings.Contains(lower, "showing"):
		return "boolean"
	case strings.Contains(lower, "index") || strings.Contains(lower, "count") || strings.Contains(lower, "num") || strings.Contains(lower, "row") || strings.Contains(lower, "col") ||
		lower == "x" || lower == "y" || lower == "elapsed" || lower == "delta" || lower == "button" || lower == "key" ||
		strings.Contains(lower, "width") || strings.Contains(lower, "height") || strings.Contains(lower, "alpha") ||
		strings.Contains(lower, "value") || strings.Contains(lower, "flags") || strings.Contains(lower, "offset"):
		return "number"
	case strings.Contains(lower, "text") || strings.Contains(lower, "label") || strings.Contains(lower, "title") || strings.Contains(lower, "tooltip"):
		return "wstring"
	case strings.Contains(lower, "name") || strings.Contains(lower, "id") || strings.Contains(lower, "str") || strings.Contains(lower, "path"):
		return "string"
	case strings.Contains(lower, "data") || strings.Contains(lower, "table") || strings.Contains(lower, "list") || strings.Contains(lower, "info"):
		return "table"
	case strings.Contains(lower, "color") || strings.Contains(lower, "tint"):
		return "table"
	default:
		return "unknown"
	}
}

// inferParamConfidence returns a confidence level based on how strongly the
// parameter name indicates its role.
func inferParamConfidence(name string) string {
	lower := strings.ToLower(name)
	// Strong signals — these names are highly conventional
	switch lower {
	case "self", "elapsed", "x", "y", "delta", "button", "key", "flags", "index", "value", "text":
		return "MEDIUM"
	}
	if strings.Contains(lower, "callback") || strings.Contains(lower, "handler") {
		return "MEDIUM"
	}
	return "LOW"
}

func isDataAPICall(name string) bool {
	for _, prefix := range knownDataAPIPrefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

func categorizeHandler(event string) string {
	switch event {
	case "OnInitialize", "OnShutdown", "OnShow", "OnHide", "OnShown", "OnHidden":
		return "lifecycle"
	case "OnUpdate":
		return "lifecycle"
	case "OnLButtonUp", "OnLButtonDown", "OnLButtonDblClk",
		"OnRButtonUp", "OnRButtonDown",
		"OnMButtonUp", "OnMButtonDown",
		"OnMouseDown", "OnMouseUp",
		"OnMouseOver", "OnMouseOverEnd", "OnMouseOut",
		"OnMouseWheel", "OnMouseDrag",
		"OnClick", "OnDoubleClick",
		"OnDragStart", "OnDragStop", "OnReceiveDrag":
		return "input"
	case "OnKeyDown", "OnKeyUp", "OnEnterPressed", "OnEscapePressed", "OnTabPressed":
		return "input"
	case "OnTextChanged", "OnSelChanged", "OnSelectionChanged",
		"OnScrollPosChanged", "OnScrollChanged",
		"OnValueChanged", "OnSliderChanged", "OnColorSelected":
		return "data"
	case "OnSizeChanged", "OnMoveStop", "OnResizeBegin", "OnResizeEnd", "OnMove":
		return "layout"
	case "OnEditFocusGained", "OnEditFocusLost":
		return "focus"
	case "OnTooltip":
		return "tooltip"
	default:
		return "custom"
	}
}

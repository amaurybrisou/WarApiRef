package lua_semantic

import (
	"strings"

	"roraddons/tools/api_doc_gen/xml_lua_binding"
)

// knownXMLHandlerArgs maps XML event names to their known handler argument signatures.
// This is the canonical source of truth for handler callback parameters.
var knownXMLHandlerArgs = map[string][]CallbackParam{
	"OnInitialize": {},
	"OnShutdown":   {},
	"OnUpdate": {
		{Position: 0, Name: "elapsed", Type: "number", Role: "time_delta", Source: "xml_handler"},
	},
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
	"OnMouseOver": {},
	"OnMouseOverEnd": {},
	"OnShown": {},
	"OnHidden": {},
	"OnMoveStop": {},
	"OnResizeBegin": {},
	"OnResizeEnd": {},
	"OnTextChanged": {
		{Position: 0, Name: "text", Type: "wstring", Role: "current_text", Source: "xml_handler"},
	},
	"OnSelChanged": {
		{Position: 0, Name: "index", Type: "number", Role: "selected_index", Source: "xml_handler"},
	},
	"OnScrollPosChanged": {
		{Position: 0, Name: "scrollPos", Type: "number", Role: "scroll_position", Source: "xml_handler"},
	},
}

// windowAPICalls lists function names that are known UI/Window API calls.
var windowAPICalls = map[string]bool{
	"WindowSetShowing":          true,
	"WindowGetShowing":          true,
	"WindowSetAlpha":            true,
	"WindowGetAlpha":            true,
	"WindowSetDimensions":       true,
	"WindowGetDimensions":       true,
	"WindowSetTintColor":        true,
	"WindowGetId":               true,
	"WindowSetId":               true,
	"LabelSetText":              true,
	"LabelGetText":              true,
	"LabelSetTextColor":         true,
	"ButtonSetText":             true,
	"ButtonSetStayDownFlag":     true,
	"ButtonGetStayDownFlag":     true,
	"DynamicImageSetTexture":    true,
	"DynamicImageSetTextureSlice": true,
	"TextEditBoxGetText":        true,
	"TextEditBoxSetText":        true,
	"ComboBoxSetSelectedIndex":  true,
	"ComboBoxGetSelectedIndex":  true,
	"ScrollWindowSetOffset":     true,
	"ScrollWindowGetOffset":     true,
	"StatusBarSetCurrentValue":  true,
	"StatusBarSetMaximumValue":  true,
	"CreateWindowFromTemplate":  true,
	"CreateWindow":              true,
	"DestroyWindow":             true,
}

// BuildSemanticCorpus performs Phase 3 analysis: deep semantic analysis of
// Lua functions that are correlated with XML elements.
func BuildSemanticCorpus(bindings *xml_lua_binding.XMLLuaBindingSet, luaIndex *xml_lua_binding.LuaFunctionIndex) *LuaSemanticCorpus {
	corpus := &LuaSemanticCorpus{
		Functions:       make(map[string]*FunctionAnalysis),
		HandlerAnalyses: make(map[string]*HandlerAnalysis),
	}

	// Analyze all resolved handler functions
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
			ha.HandlerConfidence = "MEDIUM"
		} else {
			ha.HandlerConfidence = "LOW"
		}

		// Categorize handler
		ha.HandlerCategory = categorizeHandler(binding.Event)

		corpus.HandlerAnalyses[handlerKey] = ha
	}

	return corpus
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
				Confidence: "LOW",
			}
			fa.Parameters = append(fa.Parameters, pa)
		}

		// Analyze downstream calls
		for _, call := range def.Calls {
			ca := CallAnalysis{
				Callee:    call.Name,
				Arguments: call.Arguments,
				Line:      call.Line,
			}
			if windowAPICalls[call.Name] {
				ca.IsWindowCall = true
				ca.IsAPICall = true
				ca.Category = "ui"
			} else if isDataAPICall(call.Name) {
				ca.IsAPICall = true
				ca.Category = "data"
			} else if isEventAPICall(call.Name) {
				ca.IsAPICall = true
				ca.Category = "event"
			} else {
				ca.Category = "unknown"
			}
			fa.Calls = append(fa.Calls, ca)
		}

		// Determine overall confidence
		fa.AnalysisConfidence = "MEDIUM"
		if len(fa.Parameters) > 0 || len(fa.Calls) > 0 {
			fa.EvidenceCount = len(fa.Parameters) + len(fa.Calls)
		}
	} else {
		fa.AnalysisConfidence = "LOW"
		fa.Notes = append(fa.Notes, "Function definition not found in Lua source")
	}

	c.Functions[qualifiedName] = fa
	return fa
}

// inferParamRole infers a parameter's role from its name.
func inferParamRole(name string) string {
	lower := strings.ToLower(name)
	switch {
	case lower == "self":
		return "self"
	case strings.Contains(lower, "callback") || strings.Contains(lower, "func") || strings.Contains(lower, "handler"):
		return "callback"
	case strings.Contains(lower, "flag") || strings.Contains(lower, "bool") || strings.Contains(lower, "enable"):
		return "flag"
	case strings.Contains(lower, "name") || strings.Contains(lower, "id") || strings.Contains(lower, "key"):
		return "identifier"
	case strings.Contains(lower, "data") || strings.Contains(lower, "table") || strings.Contains(lower, "list"):
		return "data"
	case strings.Contains(lower, "index") || strings.Contains(lower, "count") || strings.Contains(lower, "num"):
		return "number"
	case strings.Contains(lower, "text") || strings.Contains(lower, "str") || strings.Contains(lower, "label"):
		return "text"
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
	case strings.Contains(lower, "flag") || strings.Contains(lower, "bool") || strings.Contains(lower, "enable") || strings.Contains(lower, "visible"):
		return "boolean"
	case strings.Contains(lower, "index") || strings.Contains(lower, "count") || strings.Contains(lower, "num") || lower == "x" || lower == "y" || lower == "elapsed":
		return "number"
	case strings.Contains(lower, "text") || strings.Contains(lower, "name") || strings.Contains(lower, "label") || strings.Contains(lower, "title"):
		return "wstring"
	case strings.Contains(lower, "data") || strings.Contains(lower, "table") || strings.Contains(lower, "list"):
		return "table"
	default:
		return "unknown"
	}
}

func isDataAPICall(name string) bool {
	return strings.HasPrefix(name, "SystemData.") || strings.HasPrefix(name, "GameData.")
}

func isEventAPICall(name string) bool {
	lower := strings.ToLower(name)
	return strings.Contains(lower, "registerevent") || strings.Contains(lower, "unregisterevent") ||
		strings.Contains(lower, "broadcastevent") || strings.Contains(lower, "fireevent")
}

func categorizeHandler(event string) string {
	switch event {
	case "OnInitialize", "OnShutdown":
		return "lifecycle"
	case "OnUpdate":
		return "lifecycle"
	case "OnLButtonUp", "OnLButtonDown", "OnRButtonUp", "OnRButtonDown",
		"OnMouseOver", "OnMouseOverEnd", "OnMouseWheel":
		return "input"
	case "OnTextChanged", "OnSelChanged", "OnScrollPosChanged":
		return "data"
	default:
		return "custom"
	}
}

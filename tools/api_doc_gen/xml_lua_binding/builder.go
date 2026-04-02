package xml_lua_binding

import (
	"strings"

	"roraddons/tools/api_doc_gen/xml_structure"
)

// LuaFunctionIndex provides fast lookup of Lua functions by name.
// This is built from the parsed source model's function data.
type LuaFunctionIndex struct {
	// ByName maps normalised function name → definitions
	ByName map[string][]LuaFunctionDef
	// ByQualifiedName maps "addon.functionName" → definition
	ByQualifiedName map[string]LuaFunctionDef
}

// LuaFunctionDef is the minimal Lua function info needed for correlation.
type LuaFunctionDef struct {
	Name          string
	QualifiedName string // "addon.name" or "module.name"
	Addon         string
	File          string
	Line          int
	EndLine       int
	Params        []string
	Local         bool
	Calls         []LuaCallRef // Downstream calls
}

// LuaCallRef is a simplified call reference.
type LuaCallRef struct {
	Name      string
	Arguments []string
	Line      int
}

// BuildBindings correlates XML structure with Lua functions to produce
// a complete binding set. This is the main entry point for Phase 2.
func BuildBindings(xmlCorpus *xml_structure.XMLCorpus, luaIndex *LuaFunctionIndex) *XMLLuaBindingSet {
	result := &XMLLuaBindingSet{
		ElementTypeBindings: make(map[string]*ElementTypeBinding),
	}

	// Phase 2a: Resolve XML handler declarations to Lua functions
	for _, tree := range xmlCorpus.Trees {
		for _, node := range tree.AllNodes {
			if len(node.Handlers) == 0 {
				continue
			}
			for _, handler := range node.Handlers {
				binding := resolveHandlerBinding(node, handler, luaIndex)
				result.HandlerBindings = append(result.HandlerBindings, binding)

				if !binding.Resolved {
					result.UnresolvedHandlers = append(result.UnresolvedHandlers, &UnresolvedHandler{
						Addon:       node.Addon,
						FrameName:   node.Name,
						FrameType:   node.Tag,
						Event:       handler.Event,
						LuaFunction: handler.Function,
						XMLFile:     node.File,
						XMLLine:     handler.Line,
					})
				}

				// Aggregate into element type bindings
				etb := result.ensureElementTypeBinding(node.Tag)
				eb := etb.ensureEventBinding(handler.Event)
				eb.TotalCount++
				eb.LuaFunctions[handler.Function]++
				eb.ByAddon[node.Addon]++
				if binding.Resolved {
					eb.Resolved++
				} else {
					eb.Unresolved++
				}
			}
		}
	}

	// Phase 2b: Find Lua functions that manipulate XML frames by name
	// Look for patterns like CreateWindowFromTemplate("FrameName", ...),
	// WindowGetId("FrameName"), etc.
	for _, tree := range xmlCorpus.Trees {
		for _, node := range tree.AllNodes {
			if !node.IsNamed || node.Name == "" {
				continue
			}
			manipulations := findFrameManipulations(node, luaIndex)
			result.FrameManipulations = append(result.FrameManipulations, manipulations...)

			for _, m := range manipulations {
				etb := result.ensureElementTypeBinding(node.Tag)
				if !containsStr(etb.ManipulatingFunctions, m.LuaFunction) {
					etb.ManipulatingFunctions = append(etb.ManipulatingFunctions, m.LuaFunction)
				}
				etb.ManipulationPatterns[m.CallPattern]++
			}
		}
	}

	return result
}

// resolveHandlerBinding attempts to find the Lua function referenced by an XML handler.
func resolveHandlerBinding(node *xml_structure.XMLElement, handler xml_structure.XMLHandlerDecl, luaIndex *LuaFunctionIndex) *HandlerBinding {
	binding := &HandlerBinding{
		Addon:       node.Addon,
		FrameName:   node.Name,
		FrameType:   node.Tag,
		Event:       handler.Event,
		XMLFile:     node.File,
		XMLLine:     handler.Line,
		LuaFunction: handler.Function,
	}

	if handler.Function == "" {
		binding.Confidence = "LOW"
		return binding
	}

	// Try exact qualified match first: "addon.functionName"
	qualifiedKey := node.Addon + "." + handler.Function
	if def, ok := luaIndex.ByQualifiedName[qualifiedKey]; ok {
		fillBindingFromDef(binding, &def)
		binding.Confidence = "HIGH"
		return binding
	}

	// Try normalized name match
	normalizedName := normalizeFunc(handler.Function)
	if defs, ok := luaIndex.ByName[normalizedName]; ok {
		// Prefer same-addon match
		for _, def := range defs {
			if def.Addon == node.Addon {
				fillBindingFromDef(binding, &def)
				binding.Confidence = "HIGH"
				return binding
			}
		}
		// Accept any addon match
		if len(defs) > 0 {
			fillBindingFromDef(binding, &defs[0])
			binding.Confidence = "MEDIUM"
			return binding
		}
	}

	binding.Confidence = "LOW"
	return binding
}

func fillBindingFromDef(binding *HandlerBinding, def *LuaFunctionDef) {
	binding.Resolved = true
	binding.LuaFile = def.File
	binding.LuaLine = def.Line
	binding.LuaParams = def.Params
	binding.LuaIsLocal = def.Local
	binding.LuaQualifiedName = def.QualifiedName
}

// findFrameManipulations searches the Lua function index for functions that
// reference a specific frame name in their call arguments.
func findFrameManipulations(node *xml_structure.XMLElement, luaIndex *LuaFunctionIndex) []*FrameManipulation {
	var result []*FrameManipulation

	// Window manipulation API patterns that take a frame name as argument
	windowAPIs := map[string]bool{
		"CreateWindowFromTemplate": true,
		"CreateWindow":             true,
		"DestroyWindow":            true,
		"WindowSetShowing":         true,
		"WindowGetId":              true,
		"WindowSetId":              true,
	}

	for _, def := range luaIndex.ByQualifiedName {
		for _, call := range def.Calls {
			if !windowAPIs[call.Name] {
				continue
			}
			// Check if any argument references this frame name
			for _, arg := range call.Arguments {
				cleaned := strings.Trim(arg, "\"' ")
				if cleaned == node.Name || strings.Contains(cleaned, node.Name) {
					result = append(result, &FrameManipulation{
						LuaFunction: def.QualifiedName,
						LuaFile:     def.File,
						LuaLine:     call.Line,
						FrameName:   node.Name,
						FrameType:   node.Tag,
						Addon:       def.Addon,
						CallPattern: call.Name,
						Confidence:  "MEDIUM",
					})
				}
			}
		}
	}

	return result
}

// ensureElementTypeBinding returns or creates an ElementTypeBinding for the given tag.
func (bs *XMLLuaBindingSet) ensureElementTypeBinding(tag string) *ElementTypeBinding {
	if etb, ok := bs.ElementTypeBindings[tag]; ok {
		return etb
	}
	etb := &ElementTypeBinding{
		Tag:                  tag,
		EventBindings:        make(map[string]*EventBinding),
		ManipulationPatterns: make(map[string]int),
	}
	bs.ElementTypeBindings[tag] = etb
	return etb
}

// ensureEventBinding returns or creates an EventBinding for the given event.
func (etb *ElementTypeBinding) ensureEventBinding(event string) *EventBinding {
	if eb, ok := etb.EventBindings[event]; ok {
		return eb
	}
	eb := &EventBinding{
		Event:        event,
		LuaFunctions: make(map[string]int),
		ByAddon:      make(map[string]int),
	}
	etb.EventBindings[event] = eb
	return eb
}

func normalizeFunc(name string) string {
	trimmed := strings.TrimSpace(name)
	parts := strings.Split(trimmed, ".")
	return strings.ToLower(parts[len(parts)-1])
}

func containsStr(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

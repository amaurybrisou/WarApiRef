// Package lua_ast provides real Lua source file parsing for the phased pipeline.
//
// It bridges the output of [lua_parser.ParseFile] — which uses a proper
// tokenizer and span-based parser, not regex or line-by-line heuristics —
// into the [xml_lua_binding.LuaFunctionDef] format consumed by Phase 2 of the
// XML↔Lua analysis pipeline.
//
// # Design note
//
// The underlying [lua_parser] tokenizes Lua source into a typed token stream
// and then performs span-based extraction of function definitions, parameter
// lists, local/global scope, and downstream call chains. This is a genuine
// parse step on real source files, not a reconstruction from flattened docs.
//
// Callers should prefer [ExtractFromFile] over the synthetic reconstruction
// in platform/phase_integration.go, which is the degraded fallback path.
package lua_ast

import (
	"fmt"

	"roraddons/tools/api_doc_gen/graph"
	"roraddons/tools/api_doc_gen/lua_parser"
	"roraddons/tools/api_doc_gen/xml_lua_binding"
)

// ExtractFromFile parses a real Lua source file and returns
// [xml_lua_binding.LuaFunctionDef] entries for use in the Phase 2 pipeline.
//
// The extraction is backed by [lua_parser.ParseFile], which tokenizes and
// parses the Lua source to locate function definitions, parameter lists, and
// call chains. Regex and string-matching heuristics are not used.
func ExtractFromFile(addonName, filePath string, manifest graph.Manifest) ([]xml_lua_binding.LuaFunctionDef, error) {
	result, err := lua_parser.ParseFile(addonName, filePath, manifest)
	if err != nil {
		return nil, fmt.Errorf("lua_ast: parse %s: %w", filePath, err)
	}
	return convertFunctions(result.Functions, addonName, filePath), nil
}

// convertFunctions maps a slice of [graph.Function] (the lua_parser output
// model) to [xml_lua_binding.LuaFunctionDef] (the Phase 2 input model).
func convertFunctions(fns []graph.Function, addonName, filePath string) []xml_lua_binding.LuaFunctionDef {
	defs := make([]xml_lua_binding.LuaFunctionDef, 0, len(fns))
	for _, fn := range fns {
		def := xml_lua_binding.LuaFunctionDef{
			Name:          fn.Name,
			QualifiedName: buildQualifiedName(fn, addonName),
			Addon:         firstNonEmpty(fn.Addon, addonName),
			File:          firstNonEmpty(fn.File, filePath),
			Line:          fn.Line,
			EndLine:       fn.EndLine,
			Local:         fn.Local,
			Params:        fn.Params,
		}
		for _, call := range fn.Calls {
			def.Calls = append(def.Calls, xml_lua_binding.LuaCallRef{
				Name:      call.Name,
				Arguments: call.Arguments,
				Line:      call.Line,
			})
		}
		defs = append(defs, def)
	}
	return defs
}

// buildQualifiedName constructs the dotted qualified name for a function
// (e.g. "AddonName.FunctionName"). If the function name already contains a
// dot (e.g. "Module.Fn") it is returned as-is.
func buildQualifiedName(fn graph.Function, defaultAddon string) string {
	if containsDot(fn.Name) {
		return fn.Name
	}
	addon := firstNonEmpty(fn.Addon, defaultAddon)
	if addon != "" && fn.Name != "" {
		return addon + "." + fn.Name
	}
	return fn.Name
}

func containsDot(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return true
		}
	}
	return false
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

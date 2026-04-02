// Package semantic_merge provides the pipeline orchestrator that runs all four
// phases of the XML↔Lua analysis in sequence and produces the final enriched
// catalog.
package semantic_merge

import (
	"roraddons/tools/api_doc_gen/lua_semantic"
	"roraddons/tools/api_doc_gen/xml_lua_binding"
	"roraddons/tools/api_doc_gen/xml_structure"
)

// PipelineInput contains everything needed to run the full 4-phase pipeline.
type PipelineInput struct {
	// XMLTrees from Phase 1 parsing (one per XML file)
	XMLTrees []*xml_structure.XMLTree

	// Lua function definitions for Phase 2 correlation
	LuaFunctions []xml_lua_binding.LuaFunctionDef
}

// PipelineOutput contains the results of all four phases.
type PipelineOutput struct {
	// Phase 1 output
	XMLCorpus *xml_structure.XMLCorpus

	// Phase 2 output
	Bindings *xml_lua_binding.XMLLuaBindingSet

	// Phase 3 output
	LuaSemantic *lua_semantic.LuaSemanticCorpus

	// Phase 4 output (final merged catalog)
	Catalog *EnrichedElementCatalog
}

// RunPipeline executes the complete 4-phase XML↔Lua analysis pipeline.
//
//	Phase 1: XML Structure — build XMLCorpus from parsed trees
//	Phase 2: Lua Correlation — correlate XML handlers with Lua functions
//	Phase 3: Lua Semantic Analysis — deep analysis of correlated functions
//	Phase 4: Semantic Merge — merge all findings into EnrichedElementCatalog
func RunPipeline(input *PipelineInput) *PipelineOutput {
	output := &PipelineOutput{}

	// Phase 1: Build XML corpus with structural profiles
	output.XMLCorpus = xml_structure.BuildCorpus(input.XMLTrees)

	// Phase 2: Build Lua function index and correlate with XML
	luaIndex := buildLuaIndex(input.LuaFunctions)
	output.Bindings = xml_lua_binding.BuildBindings(output.XMLCorpus, luaIndex)

	// Phase 3: Deep semantic analysis of correlated Lua functions
	output.LuaSemantic = lua_semantic.BuildSemanticCorpus(output.Bindings, luaIndex)

	// Phase 4: Merge all findings into enriched element catalog
	output.Catalog = BuildEnrichedCatalog(output.XMLCorpus, output.Bindings, output.LuaSemantic)

	return output
}

// buildLuaIndex creates the LuaFunctionIndex from a slice of definitions.
func buildLuaIndex(defs []xml_lua_binding.LuaFunctionDef) *xml_lua_binding.LuaFunctionIndex {
	index := &xml_lua_binding.LuaFunctionIndex{
		ByName:          make(map[string][]xml_lua_binding.LuaFunctionDef),
		ByQualifiedName: make(map[string]xml_lua_binding.LuaFunctionDef),
	}

	for _, def := range defs {
		index.ByQualifiedName[def.QualifiedName] = def

		// Normalize and index by short name
		normalised := normalizeFunc(def.Name)
		index.ByName[normalised] = append(index.ByName[normalised], def)
	}

	return index
}

func normalizeFunc(name string) string {
	// Extract the last segment after any dots
	parts := splitDot(name)
	last := parts[len(parts)-1]
	result := make([]byte, len(last))
	for i, c := range []byte(last) {
		if c >= 'A' && c <= 'Z' {
			result[i] = c + 32
		} else {
			result[i] = c
		}
	}
	return string(result)
}

func splitDot(s string) []string {
	var parts []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			parts = append(parts, s[start:i])
			start = i + 1
		}
	}
	parts = append(parts, s[start:])
	return parts
}

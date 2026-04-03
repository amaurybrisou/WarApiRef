// Package semantic_merge provides the pipeline orchestrator that runs all four
// phases of the XML↔Lua analysis in sequence and produces the final enriched
// catalog.
package semantic_merge

import (
	"roraddons/tools/api_doc_gen/lua_semantic"
	"roraddons/tools/api_doc_gen/mod_semantic"
	"roraddons/tools/api_doc_gen/xml_lua_binding"
	"roraddons/tools/api_doc_gen/xml_structure"
)

// PipelineInput contains everything needed to run the full 4-phase pipeline.
type PipelineInput struct {
	// XMLTrees from Phase 1 parsing (one per XML file)
	XMLTrees []*xml_structure.XMLTree

	// Lua function definitions for Phase 2 correlation
	LuaFunctions []xml_lua_binding.LuaFunctionDef

	// PreResolvedBindings provides pre-existing XML↔Lua binding resolution data
	// from the platform contract-derived binding set, enabling Phase 2 to start
	// with already-known binding evidence.
	PreResolvedBindings []PreResolvedBinding

	// ModSemantics carries the semantic analysis results derived from each
	// addon's .mod manifest tree.  It is nil for addons with a .toc manifest.
	// The pipeline passes this through to PipelineOutput unchanged so that
	// callers can apply .mod-level facts (lifecycle entrypoints, startup
	// windows, etc.) to the enriched model.
	ModSemantics []*mod_semantic.ModuleSemantic
}

// PreResolvedBinding is a pre-existing XML↔Lua binding from contract-derived
// platform inputs, providing additional evidence for Phase 2 correlation.
type PreResolvedBinding struct {
	Addon       string
	Frame       string
	Event       string
	XMLFunction string
	LuaFunction string
	Resolved    bool
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

	// ModSemantics carries the .mod semantic analysis results passed in via
	// PipelineInput.ModSemantics.  The pipeline does not transform this data;
	// it passes it through so that callers (e.g. platform layer) can apply
	// .mod lifecycle facts to the element-type enrichment step.
	ModSemantics []*mod_semantic.ModuleSemantic
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

	// Inject pre-resolved bindings from contract-derived inputs as additional evidence
	injectPreResolvedBindings(output.Bindings, input.PreResolvedBindings, luaIndex)

	// Phase 3: Deep semantic analysis of correlated Lua functions
	output.LuaSemantic = lua_semantic.BuildSemanticCorpus(output.Bindings, luaIndex)

	// Phase 4: Merge all findings into enriched element catalog
	output.Catalog = BuildEnrichedCatalog(output.XMLCorpus, output.Bindings, output.LuaSemantic)

	// Phase 4 mod enrichment: merge .mod lifecycle semantics into the catalog.
	// Build a frame-name → element-type-tag lookup from the parsed XML trees so
	// that CreateWindow targets can be associated with their element type.
	frameTypeByName := buildFrameTypeByName(input.XMLTrees)
	enrichCatalogFromModSemantics(output.Catalog, input.ModSemantics, frameTypeByName)

	// Pass .mod semantic analysis through to the caller so that the platform
	// layer can apply additional element-type enrichments (e.g. StartupWindowFacts).
	output.ModSemantics = input.ModSemantics

	return output
}

// injectPreResolvedBindings adds evidence from platform-level binding data
// that the Phase 2 tree walk may have missed (e.g., bindings established
// via Lua RegisterEventHandler calls rather than XML EventHandler declarations).
func injectPreResolvedBindings(bindings *xml_lua_binding.XMLLuaBindingSet, preResolved []PreResolvedBinding, luaIndex *xml_lua_binding.LuaFunctionIndex) {
	if len(preResolved) == 0 || bindings == nil {
		return
	}

	// Build a set of already-known bindings to avoid duplicates
	existing := make(map[string]bool)
	for _, b := range bindings.HandlerBindings {
		key := b.Addon + "|" + b.FrameName + "|" + b.Event + "|" + b.LuaFunction
		existing[key] = true
	}

	for _, prb := range preResolved {
		key := prb.Addon + "|" + prb.Frame + "|" + prb.Event + "|" + prb.LuaFunction
		if existing[key] {
			continue
		}

		binding := &xml_lua_binding.HandlerBinding{
			Addon:       prb.Addon,
			FrameName:   prb.Frame,
			Event:       prb.Event,
			LuaFunction: prb.LuaFunction,
			Resolved:    prb.Resolved,
			Confidence:  "MEDIUM", // pre-resolved from platform analysis
		}

		// Try to resolve the function in the lua index for additional data
		if !binding.Resolved {
			qualifiedKey := prb.Addon + "." + prb.LuaFunction
			if def, ok := luaIndex.ByQualifiedName[qualifiedKey]; ok {
				binding.Resolved = true
				binding.LuaFile = def.File
				binding.LuaLine = def.Line
				binding.LuaParams = def.Params
				binding.LuaIsLocal = def.Local
				binding.LuaQualifiedName = def.QualifiedName
				binding.Confidence = "MEDIUM"
			}
		}

		bindings.HandlerBindings = append(bindings.HandlerBindings, binding)

		// Aggregate into element type bindings
		// Frame type is unknown from pre-resolved data, so skip element type aggregation
		// (it will be picked up if the frame exists in the XML corpus)
	}
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

// buildFrameTypeByName constructs a lookup from named frame name to element
// type tag across all parsed XML trees.  It is used by
// enrichCatalogFromModSemantics to resolve CreateWindow targets to element types.
func buildFrameTypeByName(trees []*xml_structure.XMLTree) map[string]string {
	m := make(map[string]string)
	for _, tree := range trees {
		for _, node := range tree.AllNodes {
			if node.IsNamed && node.Name != "" && node.Tag != "" {
				m[node.Name] = node.Tag
			}
		}
	}
	return m
}

// enrichCatalogFromModSemantics merges .mod lifecycle semantics into the
// EnrichedElementCatalog.  For each ModuleSemantic it:
//
//   - Populates EnrichedElementCatalog.AddonLifecycleSemantics with a
//     CatalogAddonLifecycleSemantic summarising hooks, actions, and saved vars.
//   - Populates EnrichedElementCatalog.FunctionLifecycleFacts so that callers
//     can look up lifecycle roles by normalised function name.
//   - Adds ModLifecycleWindowFact records to the LifecycleWindowFacts field of
//     every EnrichedElement whose type tag matches a CreateWindow target.
func enrichCatalogFromModSemantics(
	catalog *EnrichedElementCatalog,
	modSemantics []*mod_semantic.ModuleSemantic,
	frameTypeByName map[string]string,
) {
	if catalog == nil || len(modSemantics) == 0 {
		return
	}

	if catalog.FunctionLifecycleFacts == nil {
		catalog.FunctionLifecycleFacts = make(map[string][]ModLifecycleFunctionFact)
	}

	for _, sem := range modSemantics {
		addonSem := CatalogAddonLifecycleSemantic{
			AddonName:      sem.AddonName,
			SavedVariables: sem.SavedVariables,
		}

		hookKindSeen := make(map[string]bool)
		hookTagSeen := make(map[string]bool)

		for _, hook := range sem.LifecycleHooks {
			hookKind := string(hook.Kind)
			hookTag := hook.Tag

			if !hookKindSeen[hookKind] {
				hookKindSeen[hookKind] = true
				addonSem.HookKinds = append(addonSem.HookKinds, hookKind)
			}
			if !hookTagSeen[hookTag] {
				hookTagSeen[hookTag] = true
				addonSem.HookTags = append(addonSem.HookTags, hookTag)
			}

			for _, action := range hook.Actions {
				switch action.Kind {
				case mod_semantic.ActionKindCallFunction:
					fact := ModLifecycleFunctionFact{
						RefName:      action.Name,
						MatchedFuncs: action.MatchedLuaFunctions,
						AddonName:    sem.AddonName,
						HookKind:     hookKind,
						HookTag:      hookTag,
						ActionTag:    action.Tag,
						ActionIndex:  action.Index,
						Resolution:   string(action.Resolution),
						Confidence:   action.Confidence,
					}
					if hook.Kind == mod_semantic.HookKindUnknown {
						addonSem.UnknownActions = append(addonSem.UnknownActions, fact)
					} else {
						addonSem.FunctionFacts = append(addonSem.FunctionFacts, fact)
					}
					// Index by each matched qualified function name (normalised)
					for _, qname := range action.MatchedLuaFunctions {
						norm := normalizeFunc(qname)
						catalog.FunctionLifecycleFacts[norm] = append(catalog.FunctionLifecycleFacts[norm], fact)
					}
					// Also index by normalised reference name for unresolved lookups
					if len(action.MatchedLuaFunctions) == 0 {
						norm := normalizeFunc(action.Name)
						catalog.FunctionLifecycleFacts[norm] = append(catalog.FunctionLifecycleFacts[norm], fact)
					}

				case mod_semantic.ActionKindCreateWindow:
					windowFact := ModLifecycleWindowFact{
						FrameName:   action.Name,
						AddonName:   sem.AddonName,
						HookKind:    hookKind,
						HookTag:     hookTag,
						ActionTag:   action.Tag,
						ActionIndex: action.Index,
						Resolution:  string(action.Resolution),
						Confidence:  action.Confidence,
					}
					if hook.Kind == mod_semantic.HookKindUnknown {
						// Store as unknown action (reuse function fact type with RefName)
						addonSem.UnknownActions = append(addonSem.UnknownActions, ModLifecycleFunctionFact{
							RefName:     action.Name,
							AddonName:   sem.AddonName,
							HookKind:    hookKind,
							HookTag:     hookTag,
							ActionTag:   action.Tag,
							ActionIndex: action.Index,
							Resolution:  string(action.Resolution),
							Confidence:  action.Confidence,
						})
					} else {
						addonSem.WindowFacts = append(addonSem.WindowFacts, windowFact)
					}
					// Add to the EnrichedElement for the matching element type
					for _, frameName := range action.MatchedXMLNames {
						elemTag, ok := frameTypeByName[frameName]
						if !ok {
							continue
						}
						if elem, ok := catalog.Elements[elemTag]; ok {
							elem.LifecycleWindowFacts = append(elem.LifecycleWindowFacts, windowFact)
						}
					}

				default: // ActionKindUnknown
					addonSem.UnknownActions = append(addonSem.UnknownActions, ModLifecycleFunctionFact{
						RefName:     action.Name,
						AddonName:   sem.AddonName,
						HookKind:    hookKind,
						HookTag:     hookTag,
						ActionTag:   action.Tag,
						ActionIndex: action.Index,
						Resolution:  string(action.Resolution),
						Confidence:  action.Confidence,
					})
				}
			}
		}

		catalog.AddonLifecycleSemantics = append(catalog.AddonLifecycleSemantics, addonSem)
	}
}

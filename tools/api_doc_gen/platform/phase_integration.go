package platform

import (
	"fmt"
	"os"

	"roraddons/tools/api_doc_gen/lua_ast"
	"roraddons/tools/api_doc_gen/semantic_merge"
	"roraddons/tools/api_doc_gen/source_scan"
	"roraddons/tools/api_doc_gen/xml_lua_binding"
	"roraddons/tools/api_doc_gen/xml_structure"
)

// runPhasedPipeline is the pipeline dispatch function. It selects the
// source-first path when a sourceRoot is available, or falls back to the
// degraded docs-based path.
//
//   - When sourceRoot != "": calls [runPhasedPipelineFromSources] (primary path).
//   - When sourceRoot == "": calls [runPhasedPipelineFromDocs] (degraded fallback).
func runPhasedPipeline(elementTypes []ElementTypeSymbol, source SourceModel, sourceRoot string) []ElementTypeSymbol {
	if sourceRoot != "" {
		return runPhasedPipelineFromSources(elementTypes, sourceRoot, source)
	}
	return runPhasedPipelineFromDocs(elementTypes, source)
}

// runPhasedPipelineFromSources is the PRIMARY pipeline path.
//
// It discovers addon XML and Lua files from sourceRoot, parses them using
// real parsers (etree for XML, tokenizer-based lua_parser for Lua), and feeds
// the resulting phase inputs to the 4-phase semantic pipeline.
//
// The pre-resolved bindings from the SourceModel are injected as supplementary
// evidence, but the primary structural data always comes from source files.
func runPhasedPipelineFromSources(elementTypes []ElementTypeSymbol, sourceRoot string, source SourceModel) []ElementTypeSymbol {
	addonSources, err := source_scan.DiscoverAddonSources(sourceRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"warning: source-first pipeline: discover sources failed (%v); falling back to degraded docs path\n", err)
		return runPhasedPipelineFromDocs(elementTypes, source)
	}
	if len(addonSources) == 0 {
		fmt.Fprintf(os.Stderr,
			"warning: source-first pipeline: no addon sources found in %q; falling back to degraded docs path\n", sourceRoot)
		return runPhasedPipelineFromDocs(elementTypes, source)
	}

	// Phase 1: parse real XML source files.
	var trees []*xml_structure.XMLTree
	for _, addon := range addonSources {
		for _, xmlFile := range addon.XMLFiles {
			tree, parseErr := xml_structure.ParseFile(addon.AddonName, xmlFile)
			if parseErr != nil {
				fmt.Fprintf(os.Stderr,
					"warning: source-first pipeline: skipping XML file %s: %v\n", xmlFile, parseErr)
				continue
			}
			trees = append(trees, tree)
		}
	}

	// Phase 2 input: parse real Lua source files.
	var luaDefs []xml_lua_binding.LuaFunctionDef
	for _, addon := range addonSources {
		for _, luaFile := range addon.LuaFiles {
			defs, parseErr := lua_ast.ExtractFromFile(addon.AddonName, luaFile, addon.Manifest)
			if parseErr != nil {
				fmt.Fprintf(os.Stderr,
					"warning: source-first pipeline: skipping Lua file %s: %v\n", luaFile, parseErr)
				continue
			}
			luaDefs = append(luaDefs, defs...)
		}
	}

	// Supplement with pre-resolved bindings from the SourceModel as additional
	// evidence (they do not replace source-parsed hierarchy data).
	var preResolved []semantic_merge.PreResolvedBinding
	for _, b := range source.Bindings {
		preResolved = append(preResolved, semantic_merge.PreResolvedBinding{
			Addon:       b.Addon,
			Frame:       b.Frame,
			Event:       b.Event,
			XMLFunction: b.XMLFunction,
			LuaFunction: b.LuaFunction,
			Resolved:    b.Resolved,
		})
	}

	output := semantic_merge.RunPipeline(&semantic_merge.PipelineInput{
		XMLTrees:            trees,
		LuaFunctions:        luaDefs,
		PreResolvedBindings: preResolved,
	})
	return enrichElementTypesFromCatalog(elementTypes, output.Catalog)
}

// runPhasedPipelineFromDocs is the DEGRADED FALLBACK path.
//
// It reconstructs XML trees and Lua definitions from pre-flattened
// SourceModel data (API reference docs). This path is lossy: parent/child
// relationships are inferred, not parsed; Lua AST facts are unavailable;
// the hierarchy is partial.
//
// # WARNING
//
// This path is provided only for backwards compatibility when no source root
// is available. It must NOT be used as the primary pipeline input. Callers
// should prefer [runPhasedPipelineFromSources].
func runPhasedPipelineFromDocs(elementTypes []ElementTypeSymbol, source SourceModel) []ElementTypeSymbol {
	fmt.Fprintln(os.Stderr,
		"[DEGRADED] platform pipeline: running from flattened docs (no source root). "+
			"Hierarchy will be partial; provide a source root via BuildOptions.SourceRoot for source-first analysis.")

	trees := synthesizeXMLTrees(source)
	luaDefs := synthesizeLuaDefs(source)

	var preResolved []semantic_merge.PreResolvedBinding
	for _, b := range source.Bindings {
		preResolved = append(preResolved, semantic_merge.PreResolvedBinding{
			Addon:       b.Addon,
			Frame:       b.Frame,
			Event:       b.Event,
			XMLFunction: b.XMLFunction,
			LuaFunction: b.LuaFunction,
			Resolved:    b.Resolved,
		})
	}

	output := semantic_merge.RunPipeline(&semantic_merge.PipelineInput{
		XMLTrees:            trees,
		LuaFunctions:        luaDefs,
		PreResolvedBindings: preResolved,
	})
	return enrichElementTypesFromCatalog(elementTypes, output.Catalog)
}

// synthesizeXMLTrees converts SourceModel.Frames into XMLTrees that the Phase 1
// corpus builder can aggregate. Each frame becomes an XMLElement; handlers from
// FrameDoc.Handlers and SourceModel.Handlers are attached.
func synthesizeXMLTrees(source SourceModel) []*xml_structure.XMLTree {
	// Group frames by file
	byFile := make(map[string][]*xml_structure.XMLElement)
	addonByFile := make(map[string]string)

	for _, frame := range source.Frames {
		fileKey := frame.Source
		if fileKey == "" {
			fileKey = frame.Addon + "/_synthetic"
		}
		addonByFile[fileKey] = frame.Addon

		elem := &xml_structure.XMLElement{
			Tag:             frame.Type,
			Name:            frame.Name,
			RawName:         frame.Name,
			Addon:           frame.Addon,
			File:            fileKey,
			IsNamed:         frame.Name != "",
			IsTemplate:      frame.Template,
			Inherits:        frame.Inherits,
			Attributes:      frame.Attributes,
			ParentFrameName: frame.Parent,
			ParentFrameType: frame.ParentType,
		}
		if elem.Attributes == nil {
			elem.Attributes = make(map[string]string)
		}

		// Add handlers from the frame doc
		for _, h := range frame.Handlers {
			elem.Handlers = append(elem.Handlers, xml_structure.XMLHandlerDecl{
				Event:    h.Event,
				Function: h.Function,
			})
		}

		// Synthesize structural children
		for _, childType := range frame.StructuralChildTypes {
			childElem := &xml_structure.XMLElement{
				Tag:             childType,
				Addon:           frame.Addon,
				File:            fileKey,
				Parent:          elem,
				IsStructural:    true,
				Attributes:      make(map[string]string),
				ParentFrameName: frame.Name,
				ParentFrameType: frame.Type,
			}
			// Carry over attribute keys
			if attrKeys, ok := frame.StructuralChildAttrKeys[childType]; ok {
				for _, k := range attrKeys {
					childElem.Attributes[k] = ""
				}
			}
			elem.Children = append(elem.Children, childElem)
		}

		// Synthesize named child element types
		for _, childType := range frame.ChildElementTypes {
			childElem := &xml_structure.XMLElement{
				Tag:             childType,
				Addon:           frame.Addon,
				File:            fileKey,
				Parent:          elem,
				IsNamed:         true, // Named child placeholder
				Attributes:      make(map[string]string),
				ParentFrameName: frame.Name,
				ParentFrameType: frame.Type,
			}
			elem.Children = append(elem.Children, childElem)
		}

		byFile[fileKey] = append(byFile[fileKey], elem)
	}

	// Also attach handlers from source.Handlers to the synthesized elements
	elemByName := make(map[string]*xml_structure.XMLElement)
	for _, elems := range byFile {
		for _, elem := range elems {
			if elem.IsNamed {
				elemByName[elem.Addon+"|"+elem.Name] = elem
			}
		}
	}
	for _, handler := range source.Handlers {
		key := handler.Addon + "|" + handler.Frame
		if elem, ok := elemByName[key]; ok {
			// Check if handler already exists
			exists := false
			for _, h := range elem.Handlers {
				if h.Event == handler.Event && h.Function == handler.Function {
					exists = true
					break
				}
			}
			if !exists {
				elem.Handlers = append(elem.Handlers, xml_structure.XMLHandlerDecl{
					Event:    handler.Event,
					Function: handler.Function,
				})
			}
		}
	}

	// Build trees
	var trees []*xml_structure.XMLTree
	for file, elems := range byFile {
		tree := &xml_structure.XMLTree{
			Addon: addonByFile[file],
			File:  file,
			Root:  elems,
		}
		// Build flat AllNodes index
		var allNodes []*xml_structure.XMLElement
		var walk func(e *xml_structure.XMLElement)
		walk = func(e *xml_structure.XMLElement) {
			allNodes = append(allNodes, e)
			for _, c := range e.Children {
				walk(c)
			}
		}
		for _, root := range elems {
			walk(root)
		}
		tree.AllNodes = allNodes
		trees = append(trees, tree)
	}

	return trees
}

// synthesizeLuaDefs converts SourceModel.Functions into LuaFunctionDefs for the
// Phase 2 Lua correlation step.
func synthesizeLuaDefs(source SourceModel) []xml_lua_binding.LuaFunctionDef {
	var defs []xml_lua_binding.LuaFunctionDef
	for _, fn := range source.Functions {
		def := xml_lua_binding.LuaFunctionDef{
			Name:          fn.Name,
			QualifiedName: fn.Addon + "." + fn.Name,
			Addon:         fn.Addon,
			File:          fn.Source,
			Local:         fn.Local,
			Params:        fn.Parameters,
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

// enrichElementTypesFromCatalog applies the Phase 4 EnrichedElementCatalog to the
// existing ElementTypeSymbol slice, adding structured attribute profiles,
// structural child profiles, Lua API call aggregations, handler argument
// patterns, relationship data with counts, inheritance info, binding stats,
// per-event API calls, and event categories.
func enrichElementTypesFromCatalog(elementTypes []ElementTypeSymbol, catalog *semantic_merge.EnrichedElementCatalog) []ElementTypeSymbol {
	if catalog == nil {
		return elementTypes
	}

	for i, sym := range elementTypes {
		enriched, ok := catalog.Elements[sym.Name]
		if !ok {
			continue
		}

		// Attribute profiles
		for _, attr := range enriched.AttributeProfiles {
			elementTypes[i].AttributeProfiles = append(elementTypes[i].AttributeProfiles, AttributeProfileEntry{
				Name:         attr.Name,
				IsRequired:   attr.IsRequired,
				SampleValues: attr.SampleValues,
				Count:        attr.Count,
				TotalCount:   attr.TotalCount,
			})
		}

		// Structural child profiles
		for _, sc := range enriched.StructuralChildren {
			profile := StructuralChildProfile{
				Tag:   sc.Tag,
				Count: sc.Count,
			}
			for _, attr := range sc.Attributes {
				profile.Attributes = append(profile.Attributes, AttributeProfileEntry{
					Name:         attr.Name,
					IsRequired:   attr.IsRequired,
					SampleValues: attr.SampleValues,
					Count:        attr.Count,
					TotalCount:   attr.TotalCount,
				})
			}
			elementTypes[i].StructuralChildProfiles = append(elementTypes[i].StructuralChildProfiles, profile)
		}

		// Lua API calls from handlers — now including Category
		for _, call := range enriched.LuaAPICalls {
			elementTypes[i].LuaAPICallsFromHandlers = append(elementTypes[i].LuaAPICallsFromHandlers, LuaAPICallEntry{
				Function:   call.Function,
				Count:      call.Count,
				Category:   call.Category,
				FromEvents: call.FromEvents,
			})
		}

		// Handler argument patterns
		for _, pattern := range enriched.HandlerArgPatterns {
			entry := HandlerArgPatternEntry{
				Event:      pattern.Event,
				Confidence: pattern.Confidence,
			}
			for _, p := range pattern.ExpectedParams {
				entry.Params = append(entry.Params, HandlerExpectedParam{
					Position: p.Position,
					Name:     p.Name,
					Type:     p.Type,
					Role:     p.Role,
				})
			}
			elementTypes[i].HandlerArgPatterns = append(elementTypes[i].HandlerArgPatterns, entry)
		}

		// Lua manipulators
		for _, m := range enriched.LuaManipulators {
			elementTypes[i].LuaManipulators = append(elementTypes[i].LuaManipulators, m.Function)
		}

		// Enriched parent references with counts and confidence
		for _, p := range enriched.Parents {
			elementTypes[i].ParentRefs = append(elementTypes[i].ParentRefs, ElementRelRef{
				Tag:        p.Tag,
				Count:      p.Count,
				Confidence: p.Confidence,
			})
		}

		// Enriched named child references with counts and confidence
		for _, c := range enriched.Children {
			elementTypes[i].ChildRefs = append(elementTypes[i].ChildRefs, ElementRelRef{
				Tag:        c.Tag,
				Count:      c.Count,
				Confidence: c.Confidence,
			})
		}

		// Enriched structural child references with counts
		for _, sc := range enriched.StructuralChildren {
			elementTypes[i].StructuralChildRefs = append(elementTypes[i].StructuralChildRefs, ElementRelRef{
				Tag:        sc.Tag,
				Count:      sc.Count,
				Confidence: sc.Confidence,
			})
		}

		// Enrich parents/children flat lists if empty (backward compat)
		if len(enriched.Parents) > 0 && len(sym.CommonParentTypes) == 0 {
			for _, p := range enriched.Parents {
				elementTypes[i].CommonParentTypes = append(elementTypes[i].CommonParentTypes, p.Tag)
			}
		}
		if len(enriched.Children) > 0 && len(sym.CommonChildElementTypes) == 0 {
			for _, c := range enriched.Children {
				elementTypes[i].CommonChildElementTypes = append(elementTypes[i].CommonChildElementTypes, c.Tag)
			}
		}
		if len(enriched.StructuralChildren) > 0 && len(sym.CommonChildTypes) == 0 {
			for _, sc := range enriched.StructuralChildren {
				elementTypes[i].CommonChildTypes = append(elementTypes[i].CommonChildTypes, sc.Tag)
			}
		}

		// Inheritance data
		if len(enriched.InheritsBases) > 0 {
			elementTypes[i].InheritsBases = enriched.InheritsBases
		}
		elementTypes[i].IsTemplate = enriched.IsTemplate

		// Composition snippet (if available from Phase 4 and not already set)
		if enriched.CompositionSnippet != "" && elementTypes[i].CompositionSnippet == "" {
			elementTypes[i].CompositionSnippet = enriched.CompositionSnippet
		}

		// Enrich event bindings with per-event Lua API calls and categories
		enrichEventBindingsFromCatalog(&elementTypes[i], enriched)

		// Binding resolution statistics
		enrichBindingStats(&elementTypes[i], enriched)
	}

	return elementTypes
}

// enrichEventBindingsFromCatalog transfers per-event Lua API calls and category
// data from the Phase 4 enriched event bindings into the platform model's
// XMLEventBinding entries.
func enrichEventBindingsFromCatalog(sym *ElementTypeSymbol, enriched *semantic_merge.EnrichedElement) {
	if len(enriched.EventBindings) == 0 || len(sym.XMLEventBindings) == 0 {
		return
	}

	// Build lookup from enriched event bindings
	enrichedByEvent := make(map[string]*semantic_merge.EnrichedEventBinding, len(enriched.EventBindings))
	for i := range enriched.EventBindings {
		enrichedByEvent[enriched.EventBindings[i].Event] = &enriched.EventBindings[i]
	}

	for i := range sym.XMLEventBindings {
		eb, ok := enrichedByEvent[sym.XMLEventBindings[i].Event]
		if !ok {
			continue
		}

		// Transfer per-event Lua API calls
		if len(eb.LuaAPICalls) > 0 {
			sym.XMLEventBindings[i].LuaAPICalls = eb.LuaAPICalls
		}

		// Transfer handler category
		if eb.Category != "" {
			sym.XMLEventBindings[i].Category = eb.Category
		}

		// Upgrade args confidence if enriched data has better analysis
		if eb.ArgsConfidence != "" && confidenceRank(eb.ArgsConfidence) > confidenceRank(sym.XMLEventBindings[i].ArgsConfidence) {
			sym.XMLEventBindings[i].InferredArgs = eb.InferredArgs
			sym.XMLEventBindings[i].ArgsConfidence = eb.ArgsConfidence
		}
	}
}

// enrichBindingStats computes binding resolution statistics from the enriched
// event bindings. Counts how many handler bindings were resolved to Lua functions.
func enrichBindingStats(sym *ElementTypeSymbol, enriched *semantic_merge.EnrichedElement) {
	totalHandlers := 0
	resolvedCount := 0
	for _, eb := range enriched.EventBindings {
		for _, bf := range eb.LuaFunctions {
			totalHandlers += bf.Count
			if bf.Resolved {
				resolvedCount += bf.Count
			}
		}
	}
	sym.BindingTotalHandlers = totalHandlers
	sym.BindingResolvedCount = resolvedCount
	if totalHandlers > 0 {
		sym.BindingResolvedPct = resolvedCount * 100 / totalHandlers
	}
}

func confidenceRank(c string) int {
	switch c {
	case "HIGH":
		return 3
	case "MEDIUM":
		return 2
	case "LOW":
		return 1
	default:
		return 0
	}
}

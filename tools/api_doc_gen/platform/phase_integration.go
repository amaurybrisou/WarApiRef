package platform

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"roraddons/tools/api_doc_gen/lua_ast"
	"roraddons/tools/api_doc_gen/mod_semantic"
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
func runPhasedPipeline(elementTypes []ElementTypeSymbol, source SourceModel, sourceRoot string) phasedPipelineResult {
	if sourceRoot != "" {
		return runPhasedPipelineFromSources(elementTypes, sourceRoot, source)
	}
	return runPhasedPipelineFromDocs(elementTypes, source)
}

// phasedPipelineResult bundles all structured outputs from the analysis
// pipeline, including element-type enrichments and .mod lifecycle semantics.
type phasedPipelineResult struct {
	ElementTypes            []ElementTypeSymbol
	AddonLifecycleSemantics []AddonLifecycleSemantic
	FunctionLifecycleRoles  []FunctionLifecycleRole
}

// runPhasedPipelineFromSources is the PRIMARY pipeline path.
//
// It discovers addon XML and Lua files from sourceRoot, parses them using
// real parsers (etree for XML, tokenizer-based lua_parser for Lua), and feeds
// the resulting phase inputs to the 4-phase semantic pipeline.
//
// The pre-resolved bindings from the SourceModel are injected as supplementary
// evidence, but the primary structural data always comes from source files.
func runPhasedPipelineFromSources(elementTypes []ElementTypeSymbol, sourceRoot string, source SourceModel) phasedPipelineResult {
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

	// Build indexes for .mod semantic correlation from the already-collected
	// Lua definitions and XML trees.
	luaQualifiedNames := make([]string, 0, len(luaDefs))
	for _, def := range luaDefs {
		luaQualifiedNames = append(luaQualifiedNames, def.QualifiedName)
	}
	luaNameIndex := mod_semantic.BuildLuaNameIndex(luaQualifiedNames)

	xmlFrameNameIndex := buildXMLFrameNameIndex(trees)

	// Build a lookup from named frame name to element type tag for startup
	// window annotation (used in enrichElementTypesFromModSemantics).
	frameTypeByName := buildFrameTypeByName(trees)

	// Perform .mod semantic analysis for every addon that has a ModuleTree.
	// Addons with a .toc manifest have ModuleTree == nil and are skipped.
	var modSemantics []*mod_semantic.ModuleSemantic
	for _, addon := range addonSources {
		if addon.ModuleTree == nil {
			continue
		}
		sem := mod_semantic.AnalyzeTree(addon.AddonName, addon.ModuleTree, luaNameIndex, xmlFrameNameIndex)
		modSemantics = append(modSemantics, sem)
	}

	output := semantic_merge.RunPipeline(&semantic_merge.PipelineInput{
		XMLTrees:            trees,
		LuaFunctions:        luaDefs,
		PreResolvedBindings: preResolved,
		ModSemantics:        modSemantics,
	})
	enriched := enrichElementTypesFromCatalog(elementTypes, output.Catalog)
	enriched = enrichElementTypesFromModSemantics(enriched, output.ModSemantics, frameTypeByName)
	return phasedPipelineResult{
		ElementTypes:            enriched,
		AddonLifecycleSemantics: buildAddonLifecycleSemantics(output.Catalog),
		FunctionLifecycleRoles:  buildFunctionLifecycleRoles(output.Catalog),
	}
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
func runPhasedPipelineFromDocs(elementTypes []ElementTypeSymbol, source SourceModel) phasedPipelineResult {
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
	return phasedPipelineResult{
		ElementTypes: enrichElementTypesFromCatalog(elementTypes, output.Catalog),
	}
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
			RawName:         frame.RawName,
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
		if elem.RawName == "" {
			if rawAttr := strings.TrimSpace(elem.Attributes["name"]); rawAttr != "" {
				elem.RawName = rawAttr
			} else {
				elem.RawName = frame.Name
			}
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
	typeIndex := make(map[string]int, len(elementTypes))
	for i, sym := range elementTypes {
		typeIndex[sym.Name] = i
	}

	for i := range elementTypes {
		enriched, ok := catalog.Elements[elementTypes[i].Name]
		if !ok {
			continue
		}
		applyEnrichedElementType(&elementTypes[i], enriched)
	}

	// Promote element types discovered from source-first XML that are not
	// already present in the platform symbol set seeded from parsed docs.
	missingTags := make([]string, 0)
	for tag := range catalog.Elements {
		if _, ok := typeIndex[tag]; ok {
			continue
		}
		missingTags = append(missingTags, tag)
	}
	sort.Strings(missingTags)

	for _, tag := range missingTags {
		enriched := catalog.Elements[tag]
		promoted := ElementTypeSymbol{
			Name:        tag,
			Confidence:  confidenceFromCatalogString(enriched.Confidence),
			RawScore:    enriched.Score,
			Score:       enriched.Score,
			Description: strings.TrimSpace(enriched.Description),
			SeenIn:      append([]string(nil), enriched.SeenIn...),
		}
		if promoted.Description == "" {
			promoted.Description = describeElement(tag, len(enriched.SeenIn))
		}
		applyEnrichedElementType(&promoted, enriched)
		elementTypes = append(elementTypes, promoted)
	}

	sort.Slice(elementTypes, func(i, j int) bool {
		return elementTypes[i].Name < elementTypes[j].Name
	})

	return elementTypes
}

func applyEnrichedElementType(sym *ElementTypeSymbol, enriched *semantic_merge.EnrichedElement) {
	if sym.Description == "" {
		if strings.TrimSpace(enriched.Description) != "" {
			sym.Description = strings.TrimSpace(enriched.Description)
		} else {
			sym.Description = describeElement(sym.Name, len(enriched.SeenIn))
		}
	}
	if sym.Score == 0 {
		sym.Score = enriched.Score
	}
	if sym.RawScore == 0 {
		sym.RawScore = enriched.Score
	}
	if sym.Confidence == "" {
		sym.Confidence = confidenceFromCatalogString(enriched.Confidence)
	}
	if len(sym.SeenIn) == 0 {
		sym.SeenIn = append(sym.SeenIn, enriched.SeenIn...)
	}

	// Attribute profiles
	for _, attr := range enriched.AttributeProfiles {
		sym.AttributeProfiles = append(sym.AttributeProfiles, AttributeProfileEntry{
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
		sym.StructuralChildProfiles = append(sym.StructuralChildProfiles, profile)
	}

	// Lua API calls from handlers — now including Category
	for _, call := range enriched.LuaAPICalls {
		sym.LuaAPICallsFromHandlers = append(sym.LuaAPICallsFromHandlers, LuaAPICallEntry{
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
		sym.HandlerArgPatterns = append(sym.HandlerArgPatterns, entry)
	}

	// Lua manipulators
	for _, m := range enriched.LuaManipulators {
		sym.LuaManipulators = append(sym.LuaManipulators, m.Function)
	}

	// Enriched parent references with counts and confidence
	for _, p := range enriched.Parents {
		sym.ParentRefs = append(sym.ParentRefs, ElementRelRef{
			Tag:        p.Tag,
			Count:      p.Count,
			Confidence: p.Confidence,
		})
	}

	// Enriched named child references with counts and confidence
	for _, c := range enriched.Children {
		sym.ChildRefs = append(sym.ChildRefs, ElementRelRef{
			Tag:        c.Tag,
			Count:      c.Count,
			Confidence: c.Confidence,
		})
	}

	// Enriched structural child references with counts
	for _, sc := range enriched.StructuralChildren {
		sym.StructuralChildRefs = append(sym.StructuralChildRefs, ElementRelRef{
			Tag:        sc.Tag,
			Count:      sc.Count,
			Confidence: sc.Confidence,
		})
	}

	// Enrich parents/children flat lists if empty (backward compat)
	if len(enriched.Parents) > 0 && len(sym.CommonParentTypes) == 0 {
		for _, p := range enriched.Parents {
			sym.CommonParentTypes = append(sym.CommonParentTypes, p.Tag)
		}
	}
	if len(enriched.Children) > 0 && len(sym.CommonChildElementTypes) == 0 {
		for _, c := range enriched.Children {
			sym.CommonChildElementTypes = append(sym.CommonChildElementTypes, c.Tag)
		}
	}
	if len(enriched.StructuralChildren) > 0 && len(sym.CommonChildTypes) == 0 {
		for _, sc := range enriched.StructuralChildren {
			sym.CommonChildTypes = append(sym.CommonChildTypes, sc.Tag)
		}
	}

	if len(sym.CommonAttributes) == 0 {
		for _, attr := range enriched.AttributeProfiles {
			sym.CommonAttributes = append(sym.CommonAttributes, attr.Name)
		}
	}

	// Inheritance data
	if len(enriched.InheritsBases) > 0 {
		sym.InheritsBases = enriched.InheritsBases
	}
	if len(sym.CommonInherits) == 0 && len(enriched.InheritsBases) > 0 {
		sym.CommonInherits = append(sym.CommonInherits, enriched.InheritsBases...)
	}
	sym.IsTemplate = enriched.IsTemplate

	// Composition snippet (if available from Phase 4 and not already set)
	if enriched.CompositionSnippet != "" && sym.CompositionSnippet == "" {
		sym.CompositionSnippet = enriched.CompositionSnippet
	}

	for _, eb := range enriched.EventBindings {
		xmlBinding := XMLEventBinding{
			Event:          eb.Event,
			InferredArgs:   eb.InferredArgs,
			ArgsConfidence: eb.ArgsConfidence,
			Category:       eb.Category,
			LuaAPICalls:    append([]string(nil), eb.LuaAPICalls...),
		}
		for _, bf := range eb.LuaFunctions {
			xmlBinding.LuaFunctions = append(xmlBinding.LuaFunctions, bf.Name)
		}
		sym.XMLEventBindings = append(sym.XMLEventBindings, xmlBinding)
	}

	if len(sym.CommonHandlers) == 0 {
		for _, eb := range enriched.EventBindings {
			sym.CommonHandlers = append(sym.CommonHandlers, eb.Event)
		}
	}
	if len(sym.CommonHandlerFunctions) == 0 {
		for _, eb := range enriched.EventBindings {
			for _, bf := range eb.LuaFunctions {
				sym.CommonHandlerFunctions = appendUniqueString(sym.CommonHandlerFunctions, bf.Name)
			}
		}
	}

	// Enrich event bindings with per-event Lua API calls and categories
	enrichEventBindingsFromCatalog(sym, enriched)

	// Binding resolution statistics
	enrichBindingStats(sym, enriched)

	// .mod lifecycle window facts — map ModLifecycleWindowFact records from
	// the catalog's EnrichedElement into platform WindowLifecycleSemantic.
	for _, wf := range enriched.LifecycleWindowFacts {
		sym.StartupWindowFacts = append(sym.StartupWindowFacts, WindowLifecycleSemantic{
			FrameName:   wf.FrameName,
			ElementType: sym.Name,
			HookKind:    wf.HookKind,
			Resolution:  wf.Resolution,
			Confidence:  wf.Confidence,
			Provenance: ModProvenance{
				AddonName:   wf.AddonName,
				HookKind:    wf.HookKind,
				HookTag:     wf.HookTag,
				ActionTag:   wf.ActionTag,
				ActionIndex: wf.ActionIndex,
				Resolution:  wf.Resolution,
				Confidence:  wf.Confidence,
			},
		})
	}

	if len(enriched.Notes) > 0 {
		for _, note := range enriched.Notes {
			sym.Notes = appendUniqueString(sym.Notes, note)
		}
	}

	sort.Strings(sym.CommonAttributes)
	sort.Strings(sym.CommonHandlers)
	sort.Strings(sym.CommonInherits)
	sort.Strings(sym.CommonChildTypes)
	sort.Strings(sym.CommonChildElementTypes)
	sort.Strings(sym.CommonParentTypes)
}

func confidenceFromCatalogString(value string) Confidence {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "HIGH":
		return ConfidenceHigh
	case "MEDIUM":
		return ConfidenceMedium
	case "LOW":
		return ConfidenceLow
	default:
		return ConfidenceLow
	}
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

// buildXMLFrameNameIndex collects every named frame name from parsed XML trees
// and returns a presence index suitable for mod_semantic.AnalyzeTree.
func buildXMLFrameNameIndex(trees []*xml_structure.XMLTree) map[string]bool {
	names := make(map[string]bool)
	for _, tree := range trees {
		for _, node := range tree.AllNodes {
			if node.IsNamed && node.Name != "" {
				names[node.Name] = true
			}
		}
	}
	return names
}

// buildFrameTypeByName builds a lookup from named frame name to element type
// tag across all parsed XML trees.  This is used by
// enrichElementTypesFromModSemantics to resolve which element type corresponds
// to a startup-created window.
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

// enrichElementTypesFromModSemantics applies .mod lifecycle facts to the
// enriched element types.  For each CreateWindow action that was exactly
// resolved to an XML frame name, the element type of that frame (looked up
// from frameTypeByName) receives:
//
//   - A structured WindowLifecycleSemantic in StartupWindowFacts (primary
//     representation with full provenance).
//   - A human-readable note in Notes (derived from the structured fact).
//
// StartupWindowFacts populated via enrichElementTypesFromCatalog (from the
// catalog's EnrichedElement.LifecycleWindowFacts) already cover exact matches.
// This function adds ambiguous and previously-unseen facts to ensure all
// resolved and partially-resolved references are visible, and re-derives Notes
// so that presentation always reflects the structured data.
func enrichElementTypesFromModSemantics(elementTypes []ElementTypeSymbol, modSemantics []*mod_semantic.ModuleSemantic, frameTypeByName map[string]string) []ElementTypeSymbol {
	if len(modSemantics) == 0 {
		return elementTypes
	}

	// Build an index from element type tag → slice position for fast update.
	typeIndex := make(map[string]int, len(elementTypes))
	for i, sym := range elementTypes {
		typeIndex[sym.Name] = i
	}

	for _, sem := range modSemantics {
		for _, hook := range sem.LifecycleHooks {
			for _, action := range hook.Actions {
				if action.Kind != mod_semantic.ActionKindCreateWindow {
					continue
				}
				// Process both exact and ambiguous resolutions.
				if action.Resolution == mod_semantic.ResolutionUnresolved {
					continue
				}
				prov := ModProvenance{
					AddonName:   sem.AddonName,
					HookKind:    string(hook.Kind),
					HookTag:     hook.Tag,
					ActionTag:   action.Tag,
					ActionIndex: action.Index,
					Resolution:  string(action.Resolution),
					Confidence:  action.Confidence,
				}
				for _, frameName := range action.MatchedXMLNames {
					elemTag, ok := frameTypeByName[frameName]
					if !ok {
						continue
					}
					idx, found := typeIndex[elemTag]
					if !found {
						continue
					}
					// Check whether this structured fact was already added by
					// enrichElementTypesFromCatalog to avoid duplicates.
					alreadyPresent := false
					for _, existing := range elementTypes[idx].StartupWindowFacts {
						if existing.FrameName == frameName && existing.Provenance.AddonName == sem.AddonName {
							alreadyPresent = true
							break
						}
					}
					if !alreadyPresent {
						elementTypes[idx].StartupWindowFacts = append(elementTypes[idx].StartupWindowFacts, WindowLifecycleSemantic{
							FrameName:   frameName,
							ElementType: elemTag,
							HookKind:    string(hook.Kind),
							Resolution:  string(action.Resolution),
							Confidence:  action.Confidence,
							Provenance:  prov,
						})
					}
				}
			}
		}
	}

	// Derive Notes from the authoritative StartupWindowFacts.
	// This ensures Notes always reflect the structured data and prevents
	// the two representations from diverging.
	for i := range elementTypes {
		noteSet := make(map[string]bool, len(elementTypes[i].Notes))
		for _, n := range elementTypes[i].Notes {
			noteSet[n] = true
		}
		for _, wf := range elementTypes[i].StartupWindowFacts {
			hookLabel := wf.HookKind
			if hookLabel == "" {
				hookLabel = "unknown lifecycle hook"
			}
			note := "Startup-created window: " + wf.FrameName + " (from " + wf.Provenance.AddonName + " " + hookLabel + ")"
			if !noteSet[note] {
				noteSet[note] = true
				elementTypes[i].Notes = append(elementTypes[i].Notes, note)
			}
		}
	}

	return elementTypes
}

// buildAddonLifecycleSemantics converts the catalog's CatalogAddonLifecycleSemantic
// records into platform-level AddonLifecycleSemantic values.
func buildAddonLifecycleSemantics(catalog *semantic_merge.EnrichedElementCatalog) []AddonLifecycleSemantic {
	if catalog == nil || len(catalog.AddonLifecycleSemantics) == 0 {
		return nil
	}

	result := make([]AddonLifecycleSemantic, 0, len(catalog.AddonLifecycleSemantics))
	for _, cas := range catalog.AddonLifecycleSemantics {
		as := AddonLifecycleSemantic{
			AddonName:      cas.AddonName,
			HookKinds:      cas.HookKinds,
			SavedVariables: cas.SavedVariables,
		}

		// Distribute FunctionFacts into lifecycle-phase buckets.
		for _, ff := range cas.FunctionFacts {
			rec := lifecycleActionRecordFromFunctionFact(ff, "CallFunction")
			switch ff.HookKind {
			case "OnInitialize":
				as.StartupActions = append(as.StartupActions, rec)
			case "OnShutdown":
				as.ShutdownActions = append(as.ShutdownActions, rec)
			case "OnUpdate":
				as.UpdateActions = append(as.UpdateActions, rec)
			default:
				as.UnknownActions = append(as.UnknownActions, rec)
			}
			if ff.Resolution == "unresolved" {
				as.UnresolvedRefs = append(as.UnresolvedRefs, rec)
			}
		}

		// CreateWindow facts go into the same phase buckets with "CreateWindow" kind.
		for _, wf := range cas.WindowFacts {
			rec := LifecycleActionRecord{
				ActionKind:   "CreateWindow",
				ActionTag:    wf.ActionTag,
				Name:         wf.FrameName,
				HookKind:     wf.HookKind,
				HookTag:      wf.HookTag,
				Resolution:   wf.Resolution,
				Confidence:   wf.Confidence,
				MatchedNames: []string{wf.FrameName},
			}
			switch wf.HookKind {
			case "OnInitialize":
				as.StartupActions = append(as.StartupActions, rec)
			case "OnShutdown":
				as.ShutdownActions = append(as.ShutdownActions, rec)
			case "OnUpdate":
				as.UpdateActions = append(as.UpdateActions, rec)
			default:
				as.UnknownActions = append(as.UnknownActions, rec)
			}
			if wf.Resolution == "unresolved" {
				as.UnresolvedRefs = append(as.UnresolvedRefs, rec)
			}
		}

		// Unknown-section actions.
		for _, uf := range cas.UnknownActions {
			rec := lifecycleActionRecordFromFunctionFact(uf, uf.ActionTag)
			as.UnknownActions = append(as.UnknownActions, rec)
			if uf.Resolution == "unresolved" {
				as.UnresolvedRefs = append(as.UnresolvedRefs, rec)
			}
		}

		result = append(result, as)
	}
	return result
}

// lifecycleActionRecordFromFunctionFact converts a ModLifecycleFunctionFact
// to a LifecycleActionRecord.
func lifecycleActionRecordFromFunctionFact(ff semantic_merge.ModLifecycleFunctionFact, actionKind string) LifecycleActionRecord {
	return LifecycleActionRecord{
		ActionKind:   actionKind,
		ActionTag:    ff.ActionTag,
		Name:         ff.RefName,
		HookKind:     ff.HookKind,
		HookTag:      ff.HookTag,
		Resolution:   ff.Resolution,
		Confidence:   ff.Confidence,
		MatchedNames: ff.MatchedFuncs,
	}
}

// buildFunctionLifecycleRoles converts the catalog's FunctionLifecycleFacts
// index into platform-level FunctionLifecycleRole values.
func buildFunctionLifecycleRoles(catalog *semantic_merge.EnrichedElementCatalog) []FunctionLifecycleRole {
	if catalog == nil || len(catalog.FunctionLifecycleFacts) == 0 {
		return nil
	}

	// Deduplicate by (funcName, addonName, hookKind) to avoid emitting the
	// same role multiple times when a function appears in several lifecycle hooks.
	seen := make(map[string]bool)
	var result []FunctionLifecycleRole

	for _, facts := range catalog.FunctionLifecycleFacts {
		for _, ff := range facts {
			role := hookKindToRole(ff.HookKind, ff.Resolution)

			// Prefer qualified matched names; fall back to ref name.
			funcNames := ff.MatchedFuncs
			if len(funcNames) == 0 {
				funcNames = []string{ff.RefName}
			}

			for _, fname := range funcNames {
				key := fname + "|" + ff.AddonName + "|" + ff.HookKind + "|" + role
				if seen[key] {
					continue
				}
				seen[key] = true

				result = append(result, FunctionLifecycleRole{
					FuncName: fname,
					RefName:  ff.RefName,
					Role:     role,
					Provenance: ModProvenance{
						AddonName:   ff.AddonName,
						HookKind:    ff.HookKind,
						HookTag:     ff.HookTag,
						ActionTag:   ff.ActionTag,
						ActionIndex: ff.ActionIndex,
						Resolution:  ff.Resolution,
						Confidence:  ff.Confidence,
					},
				})
			}
		}
	}
	return result
}

// hookKindToRole converts a .mod lifecycle hook kind and resolution status
// to a human-readable FunctionLifecycleRole.Role value.
func hookKindToRole(hookKind, resolution string) string {
	if resolution == "unresolved" {
		return "unresolved_lifecycle_ref"
	}
	switch hookKind {
	case "OnInitialize":
		return "startup_entrypoint"
	case "OnShutdown":
		return "shutdown_entrypoint"
	case "OnUpdate":
		return "update_callback"
	default:
		return "unknown_lifecycle_ref"
	}
}

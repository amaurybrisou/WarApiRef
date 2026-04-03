package platform

import (
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"roraddons/tools/api_doc_gen/confidence"
	"roraddons/tools/api_doc_gen/graph"
	md "roraddons/tools/api_doc_gen/templates"
)

func Generate(outputRoot string, corpus Corpus) error {
	if err := prepareOutput(outputRoot); err != nil {
		return err
	}
	if err := writeREADME(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeBrowserIndex(outputRoot); err != nil {
		return err
	}
	if err := writeIndex(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeGlobalFunctionDocs(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeGlobalTables(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeConstants(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeWindowFunctions(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeWindowLifecycle(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeWindowPatterns(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeElementTypes(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeXMLHandlers(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeXMLSchema(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeEvents(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeEventPatterns(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeSystemData(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeGameData(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeSlashCommands(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeLifecycle(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeAddonLifecycleSemantics(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeFunctionLifecycleRoles(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeExamples(outputRoot, corpus); err != nil {
		return err
	}
	if err := writePatternDocs(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeConventions(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeNavigationArtifacts(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeGraphArtifacts(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeInferenceRules(outputRoot, corpus); err != nil {
		return err
	}
	return writeCoverage(outputRoot, corpus)
}

func prepareOutput(outputRoot string) error {
	managed := []string{"globals", "window_api", "xml", "events", "systemdata", "gamedata", "slash_commands", "lifecycle", "examples", "patterns", "navigation", "graph", "meta"}
	for _, name := range managed {
		if err := os.RemoveAll(filepath.Join(outputRoot, name)); err != nil {
			return fmt.Errorf("reset output %s: %w", name, err)
		}
	}
	directories := []string{
		outputRoot,
		filepath.Join(outputRoot, "globals", "functions"),
		filepath.Join(outputRoot, "globals", "tables"),
		filepath.Join(outputRoot, "globals", "constants"),
		filepath.Join(outputRoot, "window_api", "functions"),
		filepath.Join(outputRoot, "xml", "element_types"),
		filepath.Join(outputRoot, "xml", "handlers"),
		filepath.Join(outputRoot, "events", "game_events"),
		filepath.Join(outputRoot, "events", "window_events"),
		filepath.Join(outputRoot, "systemdata", "fields"),
		filepath.Join(outputRoot, "systemdata"),
		filepath.Join(outputRoot, "gamedata", "fields"),
		filepath.Join(outputRoot, "gamedata"),
		filepath.Join(outputRoot, "slash_commands"),
		filepath.Join(outputRoot, "lifecycle"),
		filepath.Join(outputRoot, "examples"),
		filepath.Join(outputRoot, "patterns"),
		filepath.Join(outputRoot, "navigation"),
		filepath.Join(outputRoot, "graph"),
		filepath.Join(outputRoot, "meta"),
	}
	for _, directory := range directories {
		if err := os.MkdirAll(directory, 0o755); err != nil {
			return fmt.Errorf("create directory %s: %w", directory, err)
		}
	}
	return nil
}

func writeREADME(outputRoot string, corpus Corpus) error {
	content := "# WAR Addon Development API Reference\n\n"
	content += "This documentation is the platform-facing second layer derived from the generated addon corpus under docs/addon-api. It attempts to answer a narrower question: which APIs, events, XML hooks, tables, and lifecycle patterns are shared across WAR addons.\n\n"
	content += md.Section("Method",
		"- Input: generated markdown under docs/addon-api only.",
		"- Output: canonical platform symbols with explicit confidence score, level, evidence signals, and rationale.",
		"- Scope: WAR-exposed functions, window APIs, XML handlers, runtime events, SystemData, GameData, lifecycle, and slash-command patterns.",
	)
	content += md.Section("Confidence",
		"- Score range: 0 to 100 after clamping the weighted raw signal total.",
		"- HIGH: 70 to 100 and promoted automatically into the canonical platform surface.",
		"- MEDIUM: 40 to 69 and promoted only when supported by namespace, default UI, event-binding, or direct XML evidence.",
		"- LOW: 0 to 39 and retained as non-canonical evidence only.",
	)
	content += md.Section("Start Here",
		"- [Overview](index.md)",
		"- [Global functions](globals/functions/index.md)",
		"- [Window API](window_api/functions/index.md)",
		"- [XML handlers](xml/handlers/index.md)",
		"- [Game events](events/game_events/index.md)",
		"- [Semantic patterns](patterns/index.md)",
		"- [Navigation tree](navigation/tree.json)",
		"- [API graph](graph/api_graph.json)",
		"- [Coverage report](meta/coverage_report.md)",
	)
	return writeFile(filepath.Join(outputRoot, "README.md"), content)
}

func writeIndex(outputRoot string, corpus Corpus) error {
	rows := [][]string{
		{"Global functions", fmt.Sprintf("%d", len(corpus.GlobalFunctions))},
		{"Window functions", fmt.Sprintf("%d", len(corpus.WindowFunctions))},
		{"Global tables", fmt.Sprintf("%d", len(corpus.GlobalTables))},
		{"Constants", fmt.Sprintf("%d", len(corpus.Constants))},
		{"XML element types", fmt.Sprintf("%d", len(corpus.ElementTypes))},
		{"XML handlers", fmt.Sprintf("%d", len(corpus.XMLHandlers))},
		{"Game events", fmt.Sprintf("%d", len(corpus.GameEvents))},
		{"Window events", fmt.Sprintf("%d", len(corpus.WindowEvents))},
		{"SystemData fields", fmt.Sprintf("%d", len(corpus.SystemDataFields))},
		{"GameData fields", fmt.Sprintf("%d", len(corpus.GameDataFields))},
	}
	candidateRows := [][]string{
		{"High confidence platform", fmt.Sprintf("%d", len(corpus.Coverage.HighConfidencePlatform))},
		{"Medium confidence candidates", fmt.Sprintf("%d", len(corpus.Coverage.MediumConfidenceCandidates))},
		{"Low confidence symbols", fmt.Sprintf("%d", len(corpus.Coverage.LowConfidenceSymbols))},
		{"Rejected addon-local", fmt.Sprintf("%d", len(corpus.Coverage.RejectedAddonLocal))},
	}
	content := "# WAR Addon Development API Reference\n\n"
	content += fmt.Sprintf("Generated from addon-api rooted at `%s`. The source corpus contained %d function docs, %d frame docs, %d handler docs, and %d event docs.\n\n", corpus.SourceRoot, corpus.Coverage.SourceCounts["function_docs"], corpus.Coverage.SourceCounts["frame_docs"], corpus.Coverage.SourceCounts["handler_docs"], corpus.Coverage.SourceCounts["event_docs"])
	content += md.Section("Coverage", md.Table([]string{"Category", "Count"}, rows))
	content += md.Section("Candidate Outcomes", md.Table([]string{"Outcome", "Count"}, candidateRows))
	content += md.Section("Sections",
		"- [Global functions](globals/functions/index.md)",
		"- [Global tables](globals/tables/index.md)",
		"- [Global constants](globals/constants/index.md)",
		"- [Window API functions](window_api/functions/index.md)",
		"- [Window lifecycle](window_api/lifecycle.md)",
		"- [Window patterns](window_api/patterns.md)",
		"- [XML element types](xml/element_types/index.md)",
		"- [XML handlers](xml/handlers/index.md)",
		"- [XML schema](xml/schema.md)",
		"- [Game events](events/game_events/index.md)",
		"- [Window events](events/window_events/index.md)",
		"- [Event patterns](events/patterns.md)",
		"- [SystemData](systemdata/index.md)",
		"- [GameData](gamedata/index.md)",
		"- [Slash commands](slash_commands/index.md)",
		"- [Addon lifecycle](lifecycle/addon_lifecycle.md)",
		"- [Semantic patterns](patterns/index.md)",
		"- [Examples](examples/usage_by_addon.md)",
		"- [Navigation tree](navigation/tree.json)",
		"- [Sidebar](navigation/sidebar.json)",
		"- [API graph](graph/api_graph.json)",
		"- [Relations](graph/relations.json)",
		"- [Conventions](meta/conventions.md)",
		"- [Inference rules](meta/inference_rules.md)",
		"- [Coverage report](meta/coverage_report.md)",
	)
	return writeFile(filepath.Join(outputRoot, "index.md"), content)
}

func writeGlobalFunctionDocs(outputRoot string, corpus Corpus) error {
	links := []string{}
	for _, symbol := range corpus.GlobalFunctions {
		currentPath := slashPath("globals", "functions", docName("global", symbol.Name))
		content := renderFunctionSymbol(corpus, symbol, currentPath)
		path := filepath.Join(outputRoot, filepath.FromSlash(currentPath))
		if err := writeFile(path, content); err != nil {
			return err
		}
		links = append(links, "- "+markdownLink(symbol.Name, docName("global", symbol.Name))+" ("+formatConfidenceShort(symbol.Confidence, symbol.Score)+")")
	}
	index := "# Global Functions\n\n"
	index += md.Section("Functions", bulletOrNone(links))
	return writeFile(filepath.Join(outputRoot, "globals", "functions", "index.md"), index)
}

func renderFunctionSymbol(corpus Corpus, symbol FunctionSymbol, currentPath string) string {
	parameterRows := [][]string{}
	for _, parameter := range symbol.Parameters {
		parameterRows = append(parameterRows, []string{parameter.Name, safeValue(parameter.Role), safeValue(parameter.Evidence)})
	}
	content := "# " + symbol.Name + "\n\n"
	content += "- Category: " + symbol.Category + "\n"
	content += "- Confidence level: " + string(symbol.Confidence) + "\n"
	content += fmt.Sprintf("- Confidence score: %d/100\n", symbol.Score)
	content += fmt.Sprintf("- Seen in: %d addons\n\n", len(symbol.SeenIn))
	content += renderConfidenceSections(symbol.Confidence, symbol.Score, symbol.RawScore, symbol.Rationale, symbol.Signals, symbol.Evidence)
	content += md.Section("Signature (inferred)", "```lua\n"+symbol.Signature+"\n```")
	content += md.Section("Description", symbol.Description)
	content += md.Section("Parameters", md.Table([]string{"Name", "Role", "Evidence"}, parameterRows))
	content += renderAdvancedReturnSection(symbol)
	content += md.Section("Side Effects", md.BulletList(symbol.SideEffects))
	content += md.Section("Seen In", md.BulletList(symbol.SeenIn))
	content += md.Section("Examples", md.BulletList(formatUsageExamples(symbol.Examples)))
	content += renderSemanticSections(currentPath, corpus.SymbolLinks[functionSymbolID(symbol.Category, symbol.Name)])
	content += md.Section("Notes", md.BulletList(symbol.Notes))
	return content
}

func renderAdvancedReturnSection(symbol FunctionSymbol) string {
	if len(symbol.ReturnPositions) == 0 && len(symbol.ReturnVariants) == 0 {
		return md.Section("Returns", md.BulletList(symbol.Returns))
	}

	lines := []string{}
	if symbol.InferredReturnArity > 0 {
		lines = append(lines, fmt.Sprintf("Observed return arity: `%v`", symbol.ObservedReturnArity))
		lines = append(lines, fmt.Sprintf("Inferred return arity: `%d`", symbol.InferredReturnArity))
	}
	if symbol.ReturnRationale != "" {
		lines = append(lines, "Rationale: "+symbol.ReturnRationale)
	}
	if symbol.BranchSensitive {
		lines = append(lines, "Branch sensitive: yes")
	}
	if symbol.WrapperAffected {
		lines = append(lines, "Wrapper affected: yes")
	}

	for _, pos := range symbol.ReturnPositions {
		head := fmt.Sprintf("### Return %d", pos.Position)
		lines = append(lines, head)
		lines = append(lines, fmt.Sprintf("- inferred type: %s", safeValue(pos.InferredType)))
		lines = append(lines, fmt.Sprintf("- inferred role: %s", safeValue(pos.InferredRole)))
		lines = append(lines, fmt.Sprintf("- confidence: %d (%s)", pos.ConfidenceScore, safeValue(pos.ConfidenceLevel)))
		if pos.Optional {
			lines = append(lines, "- optional: yes")
		}
		if pos.Stable {
			lines = append(lines, "- appears stable across observed contexts")
		}
		if len(pos.Evidence) > 0 {
			lines = append(lines, "- evidence:")
			for _, ev := range pos.Evidence {
				lines = append(lines, "  - "+safeValue(ev))
			}
		}
	}

	if len(symbol.ReturnVariants) > 0 {
		lines = append(lines, "### Observed Return Variants")
		for _, variant := range symbol.ReturnVariants {
			shape := strings.Join(variant.Shape, ", ")
			if shape == "" {
				shape = "unknown"
			}
			lines = append(lines, fmt.Sprintf("- %s: %s (arity %d, confidence %d)", safeValue(variant.Label), shape, variant.Arity, variant.Confidence))
			for _, note := range variant.Notes {
				lines = append(lines, "  - "+safeValue(note))
			}
		}
	}

	if len(symbol.ReturnUncertainty) > 0 {
		lines = append(lines, "### Uncertainty Notes")
		for _, note := range symbol.ReturnUncertainty {
			lines = append(lines, "- "+safeValue(note))
		}
	}

	return md.Section("Returns", strings.Join(lines, "\n"))
}

func writeGlobalTables(outputRoot string, corpus Corpus) error {
	links := []string{}
	for _, symbol := range corpus.GlobalTables {
		currentPath := slashPath("globals", "tables", docName("table", symbol.Name))
		content := "# " + symbol.Name + "\n\n"
		content += "- Category: Global Table\n"
		content += "- Confidence level: " + string(symbol.Confidence) + "\n"
		content += fmt.Sprintf("- Confidence score: %d/100\n\n", symbol.Score)
		content += renderConfidenceSections(symbol.Confidence, symbol.Score, symbol.RawScore, symbol.Rationale, symbol.Signals, symbol.Evidence)
		content += md.Section("Description", symbol.Description)
		content += md.Section("Functions", md.BulletList(symbol.Functions))
		content += md.Section("Observed Members", md.BulletList(symbol.ObservedMembers))
		content += md.Section("Seen In", md.BulletList(symbol.SeenIn))
		content += md.Section("Examples", md.BulletList(formatUsageExamples(symbol.RepresentativeUsages)))
		content += renderSemanticSections(currentPath, corpus.SymbolLinks[tableID(symbol.Name)])
		content += md.Section("Notes", md.BulletList(symbol.Notes))
		path := filepath.Join(outputRoot, filepath.FromSlash(currentPath))
		if err := writeFile(path, content); err != nil {
			return err
		}
		links = append(links, "- "+markdownLink(symbol.Name, docName("table", symbol.Name))+" ("+formatConfidenceShort(symbol.Confidence, symbol.Score)+")")
	}
	index := "# Global Tables\n\n"
	index += md.Section("Tables", bulletOrNone(links))
	return writeFile(filepath.Join(outputRoot, "globals", "tables", "index.md"), index)
}

func writeConstants(outputRoot string, corpus Corpus) error {
	links := []string{}
	for _, symbol := range corpus.Constants {
		currentPath := slashPath("globals", "constants", docName("constant", symbol.Name))
		content := "# " + symbol.Name + "\n\n"
		content += "- Category: Constant\n"
		content += "- Confidence level: " + string(symbol.Confidence) + "\n"
		content += fmt.Sprintf("- Confidence score: %d/100\n\n", symbol.Score)
		content += renderConfidenceSections(symbol.Confidence, symbol.Score, symbol.RawScore, symbol.Rationale, symbol.Signals, symbol.Evidence)
		content += md.Section("Description", symbol.Description)
		content += md.Section("Seen In", md.BulletList(symbol.SeenIn))
		content += md.Section("Used By", md.BulletList(symbol.UsedBy))
		content += renderSemanticSections(currentPath, corpus.SymbolLinks[constantID(symbol.Name)])
		content += md.Section("Notes", md.BulletList(symbol.Notes))
		path := filepath.Join(outputRoot, filepath.FromSlash(currentPath))
		if err := writeFile(path, content); err != nil {
			return err
		}
		links = append(links, "- "+markdownLink(symbol.Name, docName("constant", symbol.Name))+" ("+formatConfidenceShort(symbol.Confidence, symbol.Score)+")")
	}
	index := "# Global Constants\n\n"
	index += md.Section("Constants", bulletOrNone(links))
	return writeFile(filepath.Join(outputRoot, "globals", "constants", "index.md"), index)
}

func writeWindowFunctions(outputRoot string, corpus Corpus) error {
	links := []string{}
	for _, symbol := range corpus.WindowFunctions {
		currentPath := slashPath("window_api", "functions", docName("window", symbol.Name))
		path := filepath.Join(outputRoot, filepath.FromSlash(currentPath))
		if err := writeFile(path, renderFunctionSymbol(corpus, symbol, currentPath)); err != nil {
			return err
		}
		links = append(links, "- "+markdownLink(symbol.Name, docName("window", symbol.Name))+" ("+formatConfidenceShort(symbol.Confidence, symbol.Score)+")")
	}
	index := "# Window API Functions\n\n"
	index += md.Section("Functions", bulletOrNone(links))
	return writeFile(filepath.Join(outputRoot, "window_api", "functions", "index.md"), index)
}

func writeWindowLifecycle(outputRoot string, corpus Corpus) error {
	content := "# Window Lifecycle\n\n"
	content += "This page reconstructs the window-facing lifecycle from observed function calls, XML hooks, and lifecycle flow evidence.\n\n"
	content += md.Section("Creation APIs", md.BulletList(findFunctionNames(corpus.GlobalFunctions, "CreateWindow", "CreateWindowFromTemplate")))
	content += md.Section("Visibility Hooks", md.BulletList(findEventNames(corpus.WindowEvents, "OnInitialize", "OnShown", "OnHidden")))
	content += md.Section("Destruction APIs", md.BulletList(findFunctionNames(corpus.GlobalFunctions, "DestroyWindow", "DoesWindowExist")))
	content += md.Section("Observed Flow", md.BulletList(formatLifecycle(corpus.Lifecycle)))
	return writeFile(filepath.Join(outputRoot, "window_api", "lifecycle.md"), content)
}

func writeWindowPatterns(outputRoot string, corpus Corpus) error {
	content := "# Window Patterns\n\n"
	for _, pattern := range corpus.WindowPatterns {
		content += md.Section(pattern.Name,
			"- Confidence: "+string(pattern.Confidence),
			"- Description: "+pattern.Description,
			"- Evidence:",
			indentBullets(pattern.Evidence),
		)
	}
	return writeFile(filepath.Join(outputRoot, "window_api", "patterns.md"), content)
}

func writeElementTypes(outputRoot string, corpus Corpus) error {
	// Build set of element type names that will have their own pages in this run.
	// Used to decide whether CommonChildTypes entries can be rendered as links.
	elementTypePageNames := map[string]bool{}
	elementTypeByName := map[string]ElementTypeSymbol{}
	for _, s := range corpus.ElementTypes {
		elementTypePageNames[s.Name] = true
		elementTypeByName[s.Name] = s
	}

	// Build set of XMLHandler event names that have dedicated pages so we can
	// produce links from element event lists to those handler pages.
	xmlHandlerPageNames := map[string]bool{}
	for _, h := range corpus.XMLHandlers {
		xmlHandlerPageNames[h.Name] = true
	}

	links := []string{}
	for _, symbol := range corpus.ElementTypes {
		currentPath := slashPath("xml", "element_types", docName("element", symbol.Name))
		content := "# " + symbol.Name + "\n\n"
		content += "- Category: XML Element Type\n"
		content += "- Confidence level: " + string(symbol.Confidence) + "\n"
		content += fmt.Sprintf("- Confidence score: %d/100\n\n", symbol.Score)
		content += renderConfidenceSections(symbol.Confidence, symbol.Score, symbol.RawScore, symbol.Rationale, symbol.Signals, symbol.Evidence)
		content += md.Section("Description", symbol.Description)
		content += md.Section("Common Attributes", md.BulletList(symbol.CommonAttributes))
		// Render "Common Handlers" with links to XMLHandler pages where available.
		if len(symbol.CommonHandlers) > 0 {
			handlerLinks := make([]string, 0, len(symbol.CommonHandlers))
			for _, evt := range symbol.CommonHandlers {
				if xmlHandlerPageNames[evt] {
					handlerLinks = append(handlerLinks, markdownLink(evt, "../handlers/"+docName("handler", evt)))
				} else {
					handlerLinks = append(handlerLinks, evt)
				}
			}
			content += md.Section("Common Handlers", md.BulletList(handlerLinks))
		}
		if len(symbol.CommonHandlerFunctions) > 0 {
			content += md.Section("Common Handler Functions", md.BulletList(symbol.CommonHandlerFunctions))
		}
		// XML Event Bindings: per-event table showing which Lua functions are
		// commonly bound and what the expected callback signature looks like.
		if len(symbol.XMLEventBindings) > 0 {
			content += renderXMLEventBindings(symbol.XMLEventBindings, xmlHandlerPageNames)
		}
		content += md.Section("Common Inherits", md.BulletList(symbol.CommonInherits))
		// Enriched parent relationships with counts and confidence
		if len(symbol.ParentRefs) > 0 {
			parentLinks := make([]string, 0, len(symbol.ParentRefs))
			for _, pr := range symbol.ParentRefs {
				label := pr.Tag
				if elementTypePageNames[pr.Tag] {
					label = markdownLink(pr.Tag, docName("element", pr.Tag))
				}
				parentLinks = append(parentLinks, fmt.Sprintf("%s — %d× (%s)", label, pr.Count, pr.Confidence))
			}
			content += md.Section("Common Parent Elements", md.BulletList(parentLinks))
		} else if len(symbol.CommonParentTypes) > 0 {
			parentLinks := make([]string, 0, len(symbol.CommonParentTypes))
			for _, pt := range symbol.CommonParentTypes {
				if elementTypePageNames[pt] {
					parentLinks = append(parentLinks, markdownLink(pt, docName("element", pt)))
				} else {
					parentLinks = append(parentLinks, pt)
				}
			}
			content += md.Section("Common Parent Elements", md.BulletList(parentLinks))
		}
		if len(symbol.ChildRefs) > 0 {
			namedChildLinks := make([]string, 0, len(symbol.ChildRefs))
			for _, cr := range symbol.ChildRefs {
				label := cr.Tag
				if elementTypePageNames[cr.Tag] {
					label = markdownLink(cr.Tag, docName("element", cr.Tag))
				}
				namedChildLinks = append(namedChildLinks, fmt.Sprintf("%s — %d× (%s)", label, cr.Count, cr.Confidence))
			}
			content += md.Section("Common Named Child Elements", md.BulletList(namedChildLinks))
		} else if len(symbol.CommonChildElementTypes) > 0 {
			namedChildLinks := make([]string, 0, len(symbol.CommonChildElementTypes))
			for _, ct := range symbol.CommonChildElementTypes {
				if elementTypePageNames[ct] {
					namedChildLinks = append(namedChildLinks, markdownLink(ct, docName("element", ct)))
				} else {
					namedChildLinks = append(namedChildLinks, ct)
				}
			}
			content += md.Section("Common Named Child Elements", md.BulletList(namedChildLinks))
		}
		if len(symbol.StructuralChildRefs) > 0 {
			childLinks := make([]string, 0, len(symbol.StructuralChildRefs))
			for _, scr := range symbol.StructuralChildRefs {
				label := scr.Tag
				if elementTypePageNames[scr.Tag] {
					childDocName := docName("element", scr.Tag)
					label = markdownLink(scr.Tag, childDocName)
				}
				childLinks = append(childLinks, fmt.Sprintf("%s — %d× (%s)", label, scr.Count, scr.Confidence))
			}
			content += md.Section("Common Structural Child Elements", md.BulletList(childLinks))
		} else if len(symbol.CommonChildTypes) > 0 {
			childLinks := make([]string, 0, len(symbol.CommonChildTypes))
			for _, childType := range symbol.CommonChildTypes {
				if elementTypePageNames[childType] {
					childDocName := docName("element", childType)
					childLinks = append(childLinks, markdownLink(childType, childDocName))
				} else {
					childLinks = append(childLinks, childType)
				}
			}
			content += md.Section("Common Structural Child Elements", md.BulletList(childLinks))
		}
		// Template inheritance
		if len(symbol.InheritsBases) > 0 {
			content += md.Section("Common Template Bases", md.BulletList(symbol.InheritsBases))
		}
		if symbol.IsTemplate {
			content += "\n> **Note**: This element type commonly acts as a template base.\n\n"
		}
		if symbol.CompositionSnippet != "" {
			content += md.Section("Typical XML Structure", "```xml\n"+symbol.CompositionSnippet+"\n```\n")
		}
		// Phase 4 enrichment: Attribute Reference table
		if len(symbol.AttributeProfiles) > 0 {
			content += renderAttributeProfiles(symbol.AttributeProfiles)
		}
		// Phase 4 enrichment: Structural Sub-Elements with attribute data
		if len(symbol.StructuralChildProfiles) > 0 {
			content += renderStructuralChildProfiles(symbol.StructuralChildProfiles, elementTypePageNames)
		}
		content += renderRecursiveHierarchy(symbol, elementTypeByName, elementTypePageNames)
		// Phase 4 enrichment: Lua API Usage from Handlers
		if len(symbol.LuaAPICallsFromHandlers) > 0 {
			content += renderLuaAPICallsFromHandlers(symbol.LuaAPICallsFromHandlers)
		}
		// Phase 4 enrichment: Handler Argument Patterns
		if len(symbol.HandlerArgPatterns) > 0 {
			content += renderHandlerArgPatterns(symbol.HandlerArgPatterns)
		}
		// Phase 4 enrichment: Lua Manipulators
		if len(symbol.LuaManipulators) > 0 {
			content += md.Section("Lua Functions Manipulating This Type", md.BulletList(symbol.LuaManipulators))
		}
		// Phase 4 enrichment: Binding resolution statistics
		if symbol.BindingTotalHandlers > 0 {
			content += "\n## Binding Resolution\n\n"
			content += fmt.Sprintf("- Total handler declarations: %d\n", symbol.BindingTotalHandlers)
			content += fmt.Sprintf("- Resolved to Lua functions: %d (%d%%)\n\n", symbol.BindingResolvedCount, symbol.BindingResolvedPct)
		}
		// .mod lifecycle enrichment: startup-created window facts (structured)
		if len(symbol.StartupWindowFacts) > 0 {
			content += renderStartupWindowFacts(symbol.StartupWindowFacts)
		}
		content += md.Section("Seen In", md.BulletList(symbol.SeenIn))
		content += md.Section("Examples", md.BulletList(formatUsageExamples(symbol.Examples)))
		content += renderSemanticSections(currentPath, corpus.SymbolLinks[elementTypeID(symbol.Name)])
		path := filepath.Join(outputRoot, filepath.FromSlash(currentPath))
		if err := writeFile(path, content); err != nil {
			return err
		}
		links = append(links, "- "+markdownLink(symbol.Name, docName("element", symbol.Name))+" ("+formatConfidenceShort(symbol.Confidence, symbol.Score)+")")
	}
	index := "# XML Element Types\n\n"
	index += md.Section("Element Types", bulletOrNone(links))
	return writeFile(filepath.Join(outputRoot, "xml", "element_types", "index.md"), index)
}

func writeXMLHandlers(outputRoot string, corpus Corpus) error {
	links := []string{}
	for _, symbol := range corpus.XMLHandlers {
		currentPath := slashPath("xml", "handlers", docName("handler", symbol.Name))
		content := "# " + symbol.Name + "\n\n"
		content += "- Type: XML Handler\n"
		content += "- Confidence level: " + string(symbol.Confidence) + "\n"
		content += fmt.Sprintf("- Confidence score: %d/100\n\n", symbol.Score)
		content += renderConfidenceSections(symbol.Confidence, symbol.Score, symbol.RawScore, symbol.Rationale, symbol.Signals, symbol.Evidence)
		content += md.Section("Description", symbol.Description)
		content += md.Section("Expected Lua Binding", "```lua\n"+symbol.ExpectedBinding+"\n```")
		content += md.Section("Element Types", md.BulletList(symbol.ElementTypes))
		content += md.Section("Seen In", md.BulletList(symbol.SeenIn))
		content += md.Section("Examples", md.BulletList(formatUsageExamples(symbol.Examples)))
		content += renderSemanticSections(currentPath, corpus.SymbolLinks[xmlHandlerID(symbol.Name)])
		content += md.Section("Notes", md.BulletList(symbol.Notes))
		path := filepath.Join(outputRoot, filepath.FromSlash(currentPath))
		if err := writeFile(path, content); err != nil {
			return err
		}
		links = append(links, "- "+markdownLink(symbol.Name, docName("handler", symbol.Name))+" ("+formatConfidenceShort(symbol.Confidence, symbol.Score)+")")
	}
	index := "# XML Handlers\n\n"
	index += md.Section("Handlers", bulletOrNone(links))
	return writeFile(filepath.Join(outputRoot, "xml", "handlers", "index.md"), index)
}

func writeXMLSchema(outputRoot string, corpus Corpus) error {
	rows := [][]string{}
	for _, symbol := range corpus.ElementTypes {
		rows = append(rows, []string{symbol.Name, fmt.Sprintf("%d", len(symbol.SeenIn)), strings.Join(firstStrings(symbol.CommonHandlers, 4), ", "), strings.Join(firstStrings(symbol.CommonAttributes, 4), ", ")})
	}
	content := "# XML Schema\n\n"
	content += "Observed schema information derived from generated frame pages. This is a usage-oriented schema, not an engine XML XSD.\n\n"
	content += md.Section("Element Types", md.Table([]string{"Element", "Addons", "Handlers", "Attributes"}, rows))
	content += md.Section("Shared Inherits Constants", md.BulletList(constantNames(corpus.Constants)))
	return writeFile(filepath.Join(outputRoot, "xml", "schema.md"), content)
}

func writeEvents(outputRoot string, corpus Corpus) error {
	if err := writeEventGroup(outputRoot, filepath.Join(outputRoot, "events", "game_events"), "Game Events", corpus.GameEvents, "game_event", corpus); err != nil {
		return err
	}
	return writeEventGroup(outputRoot, filepath.Join(outputRoot, "events", "window_events"), "Window Events", corpus.WindowEvents, "window_event", corpus)
}

func writeEventGroup(outputRoot string, directory string, title string, events []EventSymbol, namespace string, corpus Corpus) error {
	links := []string{}
	for _, symbol := range events {
		currentPath := slashPath(strings.TrimPrefix(strings.TrimPrefix(directory, outputRoot), string(filepath.Separator)), docName(namespace, symbol.Name))
		currentPath = strings.TrimPrefix(filepath.ToSlash(currentPath), "/")
		content := "# " + symbol.Name + "\n\n"
		content += "- Category: " + symbol.Category + "\n"
		content += "- Confidence level: " + string(symbol.Confidence) + "\n"
		content += fmt.Sprintf("- Confidence score: %d/100\n\n", symbol.Score)
		content += renderConfidenceSections(symbol.Confidence, symbol.Score, symbol.RawScore, symbol.Rationale, symbol.Signals, symbol.Evidence)
		content += md.Section("Description", symbol.Description)
		content += md.Section("Handler Pattern", symbol.HandlerPattern)
		content += md.Section("Payload", md.BulletList(symbol.Payload))
		content += md.Section("Seen In", md.BulletList(symbol.SeenIn))
		content += md.Section("Registrars And Handlers", md.BulletList(symbol.Registrars))
		content += md.Section("Examples", md.BulletList(formatUsageExamples(symbol.Examples)))
		content += renderSemanticSections(currentPath, corpus.SymbolLinks[eventSymbolID(symbol.Category, symbol.Name)])
		content += md.Section("Notes", md.BulletList(symbol.Notes))
		path := filepath.Join(outputRoot, filepath.FromSlash(currentPath))
		if err := writeFile(path, content); err != nil {
			return err
		}
		links = append(links, "- "+markdownLink(symbol.Name, docName(namespace, symbol.Name))+" ("+formatConfidenceShort(symbol.Confidence, symbol.Score)+")")
	}
	index := "# " + title + "\n\n"
	index += md.Section("Events", bulletOrNone(links))
	return writeFile(filepath.Join(directory, "index.md"), index)
}

func writeEventPatterns(outputRoot string, corpus Corpus) error {
	content := "# Event Patterns\n\n"
	for _, pattern := range corpus.EventPatterns {
		content += md.Section(pattern.Name,
			"- Confidence: "+string(pattern.Confidence),
			"- Description: "+pattern.Description,
			"- Evidence:",
			indentBullets(pattern.Evidence),
		)
	}
	return writeFile(filepath.Join(outputRoot, "events", "patterns.md"), content)
}

func writeSystemData(outputRoot string, corpus Corpus) error {
	return writeFieldDocs(outputRoot, "SystemData", "systemdata", "SystemData", corpus.SystemDataFields, corpus)
}

func writeGameData(outputRoot string, corpus Corpus) error {
	return writeFieldDocs(outputRoot, "GameData", "gamedata", "GameData", corpus.GameDataFields, corpus)
}

func writeFieldIndex(path string, title string, values []FieldSymbol) error {
	rows := [][]string{}
	for _, symbol := range values {
		rows = append(rows, []string{symbol.Name, fmt.Sprintf("%d", symbol.Score), string(symbol.Confidence), fmt.Sprintf("%d", len(symbol.SeenIn)), summarizeEvidence(symbol.Evidence), symbol.Rationale})
	}
	content := "# " + title + "\n\n"
	content += md.Section("Observed Fields", md.Table([]string{"Field", "Score", "Level", "Addons", "Evidence", "Rationale"}, rows))
	return writeFile(path, content)
}

func writeSlashCommands(outputRoot string, corpus Corpus) error {
	content := "# Slash Commands\n\n"
	if len(corpus.SlashCommandPatterns) == 0 {
		content += "No shared slash-command registration pattern could be inferred from the current addon-api corpus.\n"
		return writeFile(filepath.Join(outputRoot, "slash_commands", "index.md"), content)
	}
	for _, pattern := range corpus.SlashCommandPatterns {
		content += md.Section(pattern.Name,
			"- Confidence: "+string(pattern.Confidence),
			"- Description: "+pattern.Description,
			"- Evidence:",
			indentBullets(pattern.Evidence),
		)
	}
	return writeFile(filepath.Join(outputRoot, "slash_commands", "index.md"), content)
}

func writeLifecycle(outputRoot string, corpus Corpus) error {
	content := "# Addon Lifecycle\n\n"
	for _, phase := range corpus.Lifecycle {
		content += md.Section(strings.Title(phase.Phase),
			"- Seen in addons: "+fmt.Sprintf("%d", phase.SeenInCount),
			"- Description: "+phase.Description,
			"- Evidence:",
			indentBullets(phase.Evidence),
		)
	}
	return writeFile(filepath.Join(outputRoot, "lifecycle", "addon_lifecycle.md"), content)
}

func writeExamples(outputRoot string, corpus Corpus) error {
	rows := [][]string{}
	for _, links := range corpus.Contracts.Links {
		for _, link := range links.HandlerLinks {
			frame := firstNonEmpty(link.XML.ResolvedName, link.XML.RawName)
			luaFn := firstNonEmpty(link.Lua.QualifiedName, link.Lua.DeclaredName, link.Lua.ShortName)
			rows = append(rows, []string{links.Addon, frame, link.XML.Event, luaFn})
			if len(rows) == 80 {
				break
			}
		}
		if len(rows) == 80 {
			break
		}
	}
	content := "# Usage By Addon\n\n"
	content += md.Section("XML To Lua Examples", md.Table([]string{"Addon", "Frame", "Event", "Lua Function"}, rows))
	content += md.Section("Representative Function Usage", md.BulletList(representativeExamples(corpus.GlobalFunctions, corpus.WindowFunctions, 20)))
	return writeFile(filepath.Join(outputRoot, "examples", "usage_by_addon.md"), content)
}

func writeConventions(outputRoot string, corpus Corpus) error {
	content := "# Conventions\n\n"
	for _, convention := range corpus.Conventions {
		content += md.Section(convention.Name,
			"- Confidence: "+string(convention.Confidence),
			"- Description: "+convention.Description,
			"- Evidence:",
			indentBullets(convention.Evidence),
		)
	}
	return writeFile(filepath.Join(outputRoot, "meta", "conventions.md"), content)
}

func writeInferenceRules(outputRoot string, corpus Corpus) error {
	content := "# Inference Rules\n\n"
	weightRows := [][]string{}
	for _, row := range confidence.WeightRows() {
		weightRows = append(weightRows, []string{row.Condition, formatSignedWeight(row.Weight), row.Category})
	}
	content += md.Section("Thresholds",
		"- Score clamp: 0 to 100.",
		"- HIGH: 70 to 100.",
		"- MEDIUM: 40 to 69.",
		"- LOW: 0 to 39.",
	)
	content += md.Section("Promotion Rules",
		"- HIGH symbols are promoted automatically.",
		"- MEDIUM symbols are promoted only when they also have known namespace, default UI, event-binding, direct XML attribute, or XML-plus-Lua evidence.",
		"- LOW symbols remain evidence only and do not enter canonical symbol pages.",
	)
	content += md.Section("Special Rules",
		"- Known namespace override: a known engine namespace is lifted to at least MEDIUM unless strong contradictory local-only evidence is present.",
		"- XML privilege: direct XML handler attribute evidence is treated as a strong platform signal.",
		"- Addon-prefix penalty: addon-specific naming lowers confidence unless cross-addon or namespace evidence outweighs it.",
		"- Wrapper detection: thin local wrappers over a stronger primitive API are penalized and usually remain non-canonical.",
		"- Cross-source reinforcement: symbols corroborated by multiple generated source types receive additional weight.",
	)
	content += md.Section("Signal Weights", md.Table([]string{"Condition", "Weight", "Category"}, weightRows))
	content += md.Section("Pipeline Heuristics", md.BulletList(corpus.InferenceRules))
	return writeFile(filepath.Join(outputRoot, "meta", "inference_rules.md"), content)
}

func writeCoverage(outputRoot string, corpus Corpus) error {
	sourceRows := [][]string{}
	for _, key := range sortedKeys(corpus.Coverage.SourceCounts) {
		sourceRows = append(sourceRows, []string{key, fmt.Sprintf("%d", corpus.Coverage.SourceCounts[key])})
	}
	symbolRows := [][]string{}
	for _, key := range sortedKeys(corpus.Coverage.SymbolCounts) {
		symbolRows = append(symbolRows, []string{key, fmt.Sprintf("%d", corpus.Coverage.SymbolCounts[key])})
	}
	confidenceRows := [][]string{}
	for _, confidence := range []Confidence{ConfidenceHigh, ConfidenceMedium, ConfidenceLow} {
		confidenceRows = append(confidenceRows, []string{string(confidence), fmt.Sprintf("%d", corpus.Coverage.ConfidenceCounts[confidence])})
	}
	classificationRows := [][]string{
		{"High confidence platform", fmt.Sprintf("%d", len(corpus.Coverage.HighConfidencePlatform))},
		{"Medium confidence candidates", fmt.Sprintf("%d", len(corpus.Coverage.MediumConfidenceCandidates))},
		{"Low confidence symbols", fmt.Sprintf("%d", len(corpus.Coverage.LowConfidenceSymbols))},
		{"Rejected addon-local", fmt.Sprintf("%d", len(corpus.Coverage.RejectedAddonLocal))},
	}
	matrixRows := [][]string{}
	for _, row := range corpus.Coverage.ExplanationMatrix {
		matrixRows = append(matrixRows, []string{row.Name, row.Category, fmt.Sprintf("%d", row.Score), string(row.Confidence), formatSignedWeight(row.CrossAddon), formatSignedWeight(row.Namespace), formatSignedWeight(row.XMLExposure), formatSignedWeight(row.EventLinkage), formatSignedWeight(row.DefaultUI), formatSignedWeight(row.LocalPenalty), formatSignedWeight(row.Signature)})
	}
	content := "# Coverage Report\n\n"
	content += md.Section("Source Counts", md.Table([]string{"Input", "Count"}, sourceRows))
	content += md.Section("Symbol Counts", md.Table([]string{"Symbol Type", "Count"}, symbolRows))
	content += md.Section("Confidence", md.Table([]string{"Level", "Count"}, confidenceRows))
	content += md.Section("Candidate Outcomes", md.Table([]string{"Outcome", "Count"}, classificationRows))
	content += md.Section("Spread",
		"- Symbols seen once: "+fmt.Sprintf("%d", corpus.Coverage.SeenOnceCount),
		"- Symbols seen in multiple addons: "+fmt.Sprintf("%d", corpus.Coverage.SeenManyCount),
	)
	content += md.Section("High Confidence Platform", candidateTable(corpus.Coverage.HighConfidencePlatform))
	content += md.Section("Medium Confidence Candidates", candidateTable(corpus.Coverage.MediumConfidenceCandidates))
	content += md.Section("Low Confidence Symbols", candidateTable(corpus.Coverage.LowConfidenceSymbols))
	content += md.Section("Rejected Addon Local", candidateTable(corpus.Coverage.RejectedAddonLocal))
	content += md.Section("Confidence Explanation Matrix", md.Table([]string{"Name", "Category", "Score", "Level", "Cross-addon", "Namespace", "XML", "Event", "Default UI", "Local penalty", "Signature"}, matrixRows))
	content += md.Section("Uncertain Areas", md.BulletList(corpus.Coverage.UncertainSymbols))
	return writeFile(filepath.Join(outputRoot, "meta", "coverage_report.md"), content)
}

func writeFile(path string, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(strings.TrimSpace(content)+"\n"), 0o644)
}

func markdownLink(label string, target string) string {
	return "[" + label + "](" + target + ")"
}

func docName(namespace string, name string) string {
	slug := graph.Slug(namespace + "_" + name)
	if len(slug) > 96 {
		hasher := fnv.New32a()
		_, _ = hasher.Write([]byte(slug))
		slug = slug[:80] + "_" + fmt.Sprintf("%08x", hasher.Sum32())
	}
	return slug + ".md"
}

func formatUsageExamples(examples []UsageExample) []string {
	lines := []string{}
	for _, example := range examples {
		line := example.Addon + ": "
		if example.Caller != "" {
			line += example.Caller + " -> "
		}
		line += example.Snippet
		lines = append(lines, line)
	}
	return lines
}

// renderXMLEventBindings produces a ## XML Event Bindings section that shows,
// for each observed XML handler event on this element type, which Lua functions
// are commonly bound, the expected callback signature, handler category, and
// per-event downstream Lua API calls.
func renderXMLEventBindings(bindings []XMLEventBinding, handlerPages map[string]bool) string {
	if len(bindings) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("\n## XML Event Bindings\n\n")
	b.WriteString("| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |\n")
	b.WriteString("|-------|----------|---------------------|-------------------|-----------------|\n")
	for _, binding := range bindings {
		eventCell := binding.Event
		if handlerPages[binding.Event] {
			eventCell = markdownLink(binding.Event, "../handlers/"+docName("handler", binding.Event))
		}
		luaCell := strings.Join(binding.LuaFunctions, ", ")
		if luaCell == "" {
			luaCell = "-"
		}
		argsCell := "`" + strings.ReplaceAll(binding.InferredArgs, "`", "\\`") + "`"
		catCell := binding.Category
		if catCell == "" {
			catCell = "-"
		}
		b.WriteString("| " + eventCell + " | " + catCell + " | " + luaCell + " | " + argsCell + " | " + binding.ArgsConfidence + " |\n")
	}
	b.WriteString("\n")

	// Render per-event API call details for events that have them
	hasPerEventCalls := false
	for _, binding := range bindings {
		if len(binding.LuaAPICalls) > 0 {
			hasPerEventCalls = true
			break
		}
	}
	if hasPerEventCalls {
		b.WriteString("### Per-Event Lua API Calls\n\n")
		for _, binding := range bindings {
			if len(binding.LuaAPICalls) == 0 {
				continue
			}
			b.WriteString(fmt.Sprintf("**%s** handlers call: %s\n\n", binding.Event, strings.Join(formatAPICalls(binding.LuaAPICalls), ", ")))
		}
	}

	return b.String()
}

func bulletOrNone(lines []string) string {
	if len(lines) == 0 {
		return "- none"
	}
	return strings.Join(lines, "\n")
}

func indentBullets(values []string) string {
	if len(values) == 0 {
		return "  - none"
	}
	lines := []string{}
	for _, value := range values {
		lines = append(lines, "  - "+value)
	}
	return strings.Join(lines, "\n")
}

func safeValue(value string) string {
	if strings.TrimSpace(value) == "" {
		return "unknown"
	}
	return value
}

func findFunctionNames(values []FunctionSymbol, names ...string) []string {
	lookup := map[string]FunctionSymbol{}
	for _, value := range values {
		lookup[value.Name] = value
	}
	result := []string{}
	for _, name := range names {
		if value, ok := lookup[name]; ok {
			result = append(result, value.Name+" ("+string(value.Confidence)+")")
		}
	}
	if len(result) == 0 {
		return []string{"none"}
	}
	return result
}

func findEventNames(values []EventSymbol, names ...string) []string {
	lookup := map[string]EventSymbol{}
	for _, value := range values {
		lookup[value.Name] = value
	}
	result := []string{}
	for _, name := range names {
		if value, ok := lookup[name]; ok {
			result = append(result, value.Name+" ("+string(value.Confidence)+")")
		}
	}
	if len(result) == 0 {
		return []string{"none"}
	}
	return result
}

func formatLifecycle(values []LifecyclePhase) []string {
	lines := []string{}
	for _, value := range values {
		lines = append(lines, value.Phase+": "+value.Description)
	}
	return lines
}

func constantNames(values []ConstantSymbol) []string {
	result := []string{}
	for _, value := range values {
		result = append(result, value.Name)
		if len(result) == 24 {
			break
		}
	}
	if len(result) == 0 {
		return []string{"none"}
	}
	return result
}

func representativeExamples(global []FunctionSymbol, window []FunctionSymbol, limit int) []string {
	lines := []string{}
	collect := func(values []FunctionSymbol) {
		for _, value := range values {
			for _, example := range value.Examples {
				lines = append(lines, value.Name+": "+example.Addon+" -> "+example.Snippet)
				if len(lines) == limit {
					return
				}
			}
			if len(lines) == limit {
				return
			}
		}
	}
	collect(global)
	if len(lines) < limit {
		collect(window)
	}
	if len(lines) == 0 {
		return []string{"none"}
	}
	return lines
}

func sortedKeys(values map[string]int) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func renderConfidenceSections(level Confidence, score int, rawScore int, rationale string, signals []confidence.Signal, evidence confidence.Evidence) string {
	content := ""
	if rawScore != score {
		content += md.Section("Confidence Assessment",
			"- Level: "+string(level),
			fmt.Sprintf("- Final score: %d/100", score),
			fmt.Sprintf("- Raw weighted score: %d", rawScore),
			"- Rationale: "+safeValue(rationale),
		)
	} else {
		content += md.Section("Confidence Assessment",
			"- Level: "+string(level),
			fmt.Sprintf("- Score: %d/100", score),
			"- Rationale: "+safeValue(rationale),
		)
	}
	content += md.Section("Evidence Signals", md.BulletList(renderSignalLines(signals)))
	content += md.Section("Evidence Summary", md.Table([]string{"Evidence", "Value"}, evidenceRows(evidence)))
	return content
}

func renderSignalLines(signals []confidence.Signal) []string {
	if len(signals) == 0 {
		return []string{"none"}
	}
	lines := make([]string, 0, len(signals))
	for _, signal := range signals {
		line := formatSignedWeight(signal.Weight) + " " + signal.Label
		if detail := strings.TrimSpace(signal.Detail); detail != "" {
			line += ": " + detail
		}
		lines = append(lines, line)
	}
	return lines
}

func evidenceRows(evidence confidence.Evidence) [][]string {
	rows := [][]string{
		{"Addons seen in", joinOrNone(firstStrings(evidence.AddonsSeenIn, 8))},
		{"Files seen in", joinOrNone(firstStrings(evidence.FilesSeenIn, 8))},
		{"Namespaces detected", joinOrNone(firstStrings(evidence.NamespacesDetected, 6))},
		{"Source kinds", joinOrNone(firstStrings(evidence.SourceKinds, 6))},
		{"Example locations", joinOrNone(firstStrings(evidence.ExampleLocations, 6))},
		{"XML usage count", fmt.Sprintf("%d", evidence.XMLUsageCount)},
		{"XML attribute usage count", fmt.Sprintf("%d", evidence.XMLAttributeUsageCount)},
		{"Lua usage count", fmt.Sprintf("%d", evidence.LuaUsageCount)},
		{"Global usage count", fmt.Sprintf("%d", evidence.GlobalUsageCount)},
		{"Local definition count", fmt.Sprintf("%d", evidence.LocalDefinitionCount)},
		{"Documentation references", fmt.Sprintf("%d", evidence.DocumentationReferenceCount)},
		{"Initialization flow references", fmt.Sprintf("%d", evidence.TOCInitReferenceCount)},
		{"Known engine namespace", yesNo(evidence.KnownEngineNamespace)},
		{"Default UI presence", yesNo(evidence.DefaultUIPresence)},
		{"Event binding presence", yesNo(evidence.EventBindingPresence)},
		{"Observed in XML and Lua", yesNo(evidence.UsedByXMLAndLua)},
		{"Consistent role", yesNo(evidence.ConsistentRole)},
		{"Consistent arguments", yesNo(evidence.ConsistentArguments)},
		{"Consistent returns", yesNo(evidence.ConsistentReturns)},
		{"Slash command presence", yesNo(evidence.SlashCommandPresence)},
		{"Weak usage only", yesNo(evidence.WeakUsageOnly)},
		{"Project-specific name", yesNo(evidence.ProjectSpecificName)},
		{"Placeholder or computed name", yesNo(evidence.PlaceholderLikeName)},
		{"Conflicting signatures", yesNo(evidence.ConflictingSignatures)},
		{"Conflicting roles", yesNo(evidence.ConflictingRoles)},
		{"Wrapper likely", yesNo(evidence.WrapperLikely)},
		{"Never outside local graph", yesNo(evidence.NeverOutsideLocalGraph)},
		{"Local helper only", yesNo(evidence.LocalHelperOnly)},
	}
	filtered := [][]string{}
	for _, row := range rows {
		if row[1] == "none" && (row[0] == "Files seen in" || row[0] == "Namespaces detected" || row[0] == "Example locations") {
			continue
		}
		filtered = append(filtered, row)
	}
	return filtered
}

func formatConfidenceShort(level Confidence, score int) string {
	return string(level) + " " + fmt.Sprintf("%d/100", score)
}

func formatSignedWeight(value int) string {
	if value > 0 {
		return fmt.Sprintf("+%d", value)
	}
	return fmt.Sprintf("%d", value)
}

func candidateTable(values []CandidateSummary) string {
	if len(values) == 0 {
		return "- none"
	}
	rows := [][]string{}
	for _, value := range values {
		rows = append(rows, []string{value.Name, value.Category, fmt.Sprintf("%d", value.Score), string(value.Confidence), summarizeEvidence(value.Evidence), value.Rationale})
	}
	return md.Table([]string{"Name", "Category", "Score", "Level", "Evidence", "Rationale"}, rows)
}

func summarizeEvidence(evidence confidence.Evidence) string {
	parts := []string{}
	if len(evidence.AddonsSeenIn) > 0 {
		parts = append(parts, fmt.Sprintf("addons=%d", len(evidence.AddonsSeenIn)))
	}
	if evidence.XMLAttributeUsageCount > 0 {
		parts = append(parts, fmt.Sprintf("xml=%d", evidence.XMLAttributeUsageCount))
	}
	if evidence.GlobalUsageCount > 0 {
		parts = append(parts, fmt.Sprintf("global=%d", evidence.GlobalUsageCount))
	}
	if evidence.LocalDefinitionCount > 0 {
		parts = append(parts, fmt.Sprintf("local=%d", evidence.LocalDefinitionCount))
	}
	if len(evidence.SourceKinds) > 0 {
		parts = append(parts, "sources="+strings.Join(firstStrings(evidence.SourceKinds, 3), ","))
	}
	if len(parts) == 0 {
		return "none"
	}
	return strings.Join(parts, "; ")
}

func joinOrNone(values []string) string {
	if len(values) == 0 {
		return "none"
	}
	return strings.Join(values, ", ")
}

func yesNo(value bool) string {
	if value {
		return "yes"
	}
	return "no"
}

// ---------------------------------------------------------------------------
// Phase 4 rendering helpers for enriched element type data
// ---------------------------------------------------------------------------

// renderAttributeProfiles generates a markdown table showing attribute reference
// data with required/optional status and sample values.
func renderAttributeProfiles(profiles []AttributeProfileEntry) string {
	if len(profiles) == 0 {
		return ""
	}
	rows := make([][]string, 0, len(profiles))
	for _, p := range profiles {
		required := "optional"
		if p.IsRequired {
			required = "**required**"
		}
		samples := ""
		if len(p.SampleValues) > 0 {
			vals := p.SampleValues
			if len(vals) > 4 {
				vals = vals[:4]
			}
			samples = strings.Join(vals, ", ")
			if len(p.SampleValues) > 4 {
				samples += ", ..."
			}
		}
		pct := ""
		if p.TotalCount > 0 {
			pct = fmt.Sprintf("%d%%", p.Count*100/p.TotalCount)
		}
		rows = append(rows, []string{"`" + p.Name + "`", required, pct, samples})
	}
	content := "## Attribute Reference\n\n"
	content += md.Table([]string{"Attribute", "Required", "Usage %", "Sample Values"}, rows)
	return content + "\n"
}

// renderStructuralChildProfiles generates detailed documentation for unnamed
// structural child elements, including their attributes.
func renderStructuralChildProfiles(profiles []StructuralChildProfile, elementTypePages map[string]bool) string {
	if len(profiles) == 0 {
		return ""
	}
	content := "## Structural Sub-Elements\n\n"
	for _, sc := range profiles {
		heading := sc.Tag
		if elementTypePages[sc.Tag] {
			heading = markdownLink(sc.Tag, docName("element", sc.Tag))
		}
		content += fmt.Sprintf("### %s\n\n", heading)
		content += fmt.Sprintf("Observed %d times as an unnamed child.\n\n", sc.Count)

		if len(sc.Attributes) > 0 {
			rows := make([][]string, 0, len(sc.Attributes))
			for _, attr := range sc.Attributes {
				required := "optional"
				if attr.IsRequired {
					required = "**required**"
				}
				samples := strings.Join(firstStrings(attr.SampleValues, 4), ", ")
				rows = append(rows, []string{"`" + attr.Name + "`", required, samples})
			}
			content += md.Table([]string{"Attribute", "Required", "Sample Values"}, rows)
			content += "\n"
		}
	}
	return content
}

type hierarchyEdge struct {
	Tag        string
	Kind       string
	Count      int
	Confidence string
}

// renderRecursiveHierarchy writes a deterministic descendant tree for this
// element type using source-derived named/structural relationships.
func renderRecursiveHierarchy(root ElementTypeSymbol, byName map[string]ElementTypeSymbol, elementTypePages map[string]bool) string {
	edges := elementHierarchyEdges(root)
	if len(edges) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("## Recursive Hierarchy\n\n")
	b.WriteString("- Root: ")
	b.WriteString(formatHierarchyNode(root.Name, elementTypePages))
	b.WriteString("\n")

	visited := map[string]bool{root.Name: true}
	renderHierarchyEdges(&b, edges, 0, visited, byName, elementTypePages)
	b.WriteString("\n")
	return b.String()
}

func renderHierarchyEdges(
	b *strings.Builder,
	edges []hierarchyEdge,
	depth int,
	path map[string]bool,
	byName map[string]ElementTypeSymbol,
	elementTypePages map[string]bool,
) {
	const maxDepth = 8
	if depth >= maxDepth {
		indent := strings.Repeat("  ", depth)
		b.WriteString(indent + "- ... (depth limit reached)\n")
		return
	}

	for _, edge := range edges {
		indent := strings.Repeat("  ", depth)
		b.WriteString(indent + "- ")
		b.WriteString(formatHierarchyNode(edge.Tag, elementTypePages))
		b.WriteString(" ")
		b.WriteString(formatHierarchyMeta(edge))
		b.WriteString("\n")

		if path[edge.Tag] {
			b.WriteString(indent + "  - (cycle)\n")
			continue
		}

		child, ok := byName[edge.Tag]
		if !ok {
			continue
		}
		childEdges := elementHierarchyEdges(child)
		if len(childEdges) == 0 {
			continue
		}

		path[edge.Tag] = true
		renderHierarchyEdges(b, childEdges, depth+1, path, byName, elementTypePages)
		delete(path, edge.Tag)
	}
}

func elementHierarchyEdges(symbol ElementTypeSymbol) []hierarchyEdge {
	byKey := make(map[string]hierarchyEdge)

	for _, child := range symbol.ChildRefs {
		key := "named|" + child.Tag
		byKey[key] = hierarchyEdge{
			Tag:        child.Tag,
			Kind:       "named",
			Count:      child.Count,
			Confidence: child.Confidence,
		}
	}
	for _, child := range symbol.StructuralChildRefs {
		key := "structural|" + child.Tag
		byKey[key] = hierarchyEdge{
			Tag:        child.Tag,
			Kind:       "structural",
			Count:      child.Count,
			Confidence: child.Confidence,
		}
	}

	// Backward compatibility: when refs are unavailable, fall back to flat lists.
	if len(symbol.ChildRefs) == 0 {
		for _, tag := range symbol.CommonChildElementTypes {
			key := "named|" + tag
			if _, exists := byKey[key]; exists {
				continue
			}
			byKey[key] = hierarchyEdge{Tag: tag, Kind: "named"}
		}
	}
	if len(symbol.StructuralChildRefs) == 0 {
		for _, tag := range symbol.CommonChildTypes {
			key := "structural|" + tag
			if _, exists := byKey[key]; exists {
				continue
			}
			byKey[key] = hierarchyEdge{Tag: tag, Kind: "structural"}
		}
	}

	result := make([]hierarchyEdge, 0, len(byKey))
	for _, edge := range byKey {
		if strings.TrimSpace(edge.Tag) == "" {
			continue
		}
		result = append(result, edge)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Kind != result[j].Kind {
			return result[i].Kind < result[j].Kind
		}
		return result[i].Tag < result[j].Tag
	})

	return result
}

func formatHierarchyNode(tag string, elementTypePages map[string]bool) string {
	if elementTypePages[tag] {
		return markdownLink(tag, docName("element", tag))
	}
	return tag
}

func formatHierarchyMeta(edge hierarchyEdge) string {
	parts := []string{edge.Kind}
	if edge.Count > 0 {
		parts = append(parts, fmt.Sprintf("%d×", edge.Count))
	}
	if strings.TrimSpace(edge.Confidence) != "" {
		parts = append(parts, edge.Confidence)
	}
	return "(" + strings.Join(parts, ", ") + ")"
}

// renderLuaAPICallsFromHandlers generates a section showing API calls commonly
// made from XML event handler functions on this element type.
func renderLuaAPICallsFromHandlers(calls []LuaAPICallEntry) string {
	if len(calls) == 0 {
		return ""
	}
	content := "## Lua API Usage (from Handlers)\n\n"
	content += "API functions commonly called from event handler Lua functions on this element type:\n\n"
	rows := make([][]string, 0, len(calls))
	for _, c := range calls {
		events := strings.Join(c.FromEvents, ", ")
		cat := c.Category
		if cat == "" {
			cat = "-"
		}
		rows = append(rows, []string{"`" + c.Function + "`", cat, fmt.Sprintf("%d", c.Count), events})
	}
	content += md.Table([]string{"API Function", "Category", "Call Count", "From Events"}, rows)
	return content + "\n"
}

// formatAPICalls formats API call names for inline display.
func formatAPICalls(calls []string) []string {
	result := make([]string, 0, len(calls))
	for _, call := range calls {
		result = append(result, "`"+call+"`")
	}
	return result
}

// renderHandlerArgPatterns generates a section showing expected callback argument
// patterns for XML event handler functions on this element type.
func renderHandlerArgPatterns(patterns []HandlerArgPatternEntry) string {
	if len(patterns) == 0 {
		return ""
	}
	content := "## Handler Callback Signatures\n\n"
	content += "Expected callback argument patterns for event handlers on this element type:\n\n"
	for _, p := range patterns {
		content += fmt.Sprintf("### %s\n\n", p.Event)
		content += fmt.Sprintf("Confidence: %s\n\n", p.Confidence)
		if len(p.Params) > 0 {
			rows := make([][]string, 0, len(p.Params))
			for _, param := range p.Params {
				rows = append(rows, []string{
					fmt.Sprintf("%d", param.Position),
					"`" + param.Name + "`",
					param.Type,
					param.Role,
				})
			}
			content += md.Table([]string{"Position", "Name", "Type", "Role"}, rows)
			content += "\n"
		}
	}
	return content
}

// renderStartupWindowFacts generates a structured section for element types
// that are instantiated as startup windows via .mod lifecycle hooks.
func renderStartupWindowFacts(facts []WindowLifecycleSemantic) string {
	if len(facts) == 0 {
		return ""
	}
	content := "## .mod Lifecycle: Startup Windows\n\n"
	content += "This element type is instantiated as a startup window by the following .mod addon(s):\n\n"
	rows := make([][]string, 0, len(facts))
	for _, f := range facts {
		hook := f.HookKind
		if hook == "" {
			hook = "(unknown)"
		}
		rows = append(rows, []string{
			f.FrameName,
			f.Provenance.AddonName,
			hook,
			f.Resolution,
			f.Confidence,
		})
	}
	content += md.Table([]string{"Frame Name", "Addon", "Hook", "Resolution", "Confidence"}, rows)
	content += "\n"
	return content
}

// writeAddonLifecycleSemantics writes addon-level lifecycle semantic pages
// derived from .mod manifest analysis.
func writeAddonLifecycleSemantics(outputRoot string, corpus Corpus) error {
	if len(corpus.AddonLifecycleSemantics) == 0 {
		return nil
	}
	dir := filepath.Join(outputRoot, "lifecycle", "addons")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create addon lifecycle dir: %w", err)
	}

	links := []string{}
	for _, addon := range corpus.AddonLifecycleSemantics {
		content := "# " + addon.AddonName + " Lifecycle\n\n"
		content += "> Source: `.mod` manifest semantic analysis\n\n"

		if len(addon.HookKinds) > 0 {
			content += md.Section("Lifecycle Hooks", md.BulletList(addon.HookKinds))
		}
		if len(addon.SavedVariables) > 0 {
			content += md.Section("Saved Variables", md.BulletList(addon.SavedVariables))
		}

		content += renderLifecycleActionTable("Startup Actions (OnInitialize)", addon.StartupActions)
		content += renderLifecycleActionTable("Shutdown Actions (OnShutdown)", addon.ShutdownActions)
		content += renderLifecycleActionTable("Update Actions (OnUpdate)", addon.UpdateActions)
		content += renderLifecycleActionTable("Unknown / Custom Hook Actions", addon.UnknownActions)

		if len(addon.UnresolvedRefs) > 0 {
			content += renderLifecycleActionTable("Unresolved References", addon.UnresolvedRefs)
		}

		slug := strings.ToLower(strings.ReplaceAll(addon.AddonName, " ", "_"))
		path := filepath.Join(dir, slug+".md")
		if err := writeFile(path, content); err != nil {
			return err
		}
		links = append(links, "- "+markdownLink(addon.AddonName, "addons/"+slug))
	}

	index := "# Addon Lifecycle Semantics\n\n"
	index += "> Structured lifecycle facts derived from `.mod` manifest analysis.\n\n"
	index += md.Section("Addons", bulletOrNone(links))
	return writeFile(filepath.Join(outputRoot, "lifecycle", "addon_lifecycle.md"), index)
}

// renderLifecycleActionTable renders a table of LifecycleActionRecord entries
// under the given section heading.  Returns an empty string when the slice is
// empty so that callers can safely call it unconditionally.
func renderLifecycleActionTable(heading string, actions []LifecycleActionRecord) string {
	if len(actions) == 0 {
		return ""
	}
	content := "## " + heading + "\n\n"
	rows := make([][]string, 0, len(actions))
	for _, a := range actions {
		matched := strings.Join(a.MatchedNames, ", ")
		if matched == "" {
			matched = "—"
		}
		rows = append(rows, []string{
			a.ActionKind,
			a.Name,
			a.Resolution,
			a.Confidence,
			matched,
		})
	}
	content += md.Table([]string{"Kind", "Name", "Resolution", "Confidence", "Matched"}, rows)
	content += "\n"
	return content
}

// writeFunctionLifecycleRoles writes a summary page for function-level
// lifecycle roles derived from .mod manifest analysis.
func writeFunctionLifecycleRoles(outputRoot string, corpus Corpus) error {
	if len(corpus.FunctionLifecycleRoles) == 0 {
		return nil
	}
	dir := filepath.Join(outputRoot, "lifecycle")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create lifecycle dir: %w", err)
	}

	// Group by role for readability.
	byRole := make(map[string][]FunctionLifecycleRole)
	for _, r := range corpus.FunctionLifecycleRoles {
		byRole[r.Role] = append(byRole[r.Role], r)
	}

	roleOrder := []string{
		"startup_entrypoint",
		"shutdown_entrypoint",
		"update_callback",
		"unresolved_lifecycle_ref",
		"unknown_lifecycle_ref",
	}
	roleLabels := map[string]string{
		"startup_entrypoint":       "Startup Entrypoints (OnInitialize)",
		"shutdown_entrypoint":      "Shutdown Entrypoints (OnShutdown)",
		"update_callback":          "Update Callbacks (OnUpdate)",
		"unresolved_lifecycle_ref": "Unresolved Lifecycle References",
		"unknown_lifecycle_ref":    "Unknown / Custom Lifecycle Refs",
	}

	content := "# Function Lifecycle Roles\n\n"
	content += "> Lua functions with lifecycle roles derived from `.mod` manifest analysis.\n\n"

	for _, role := range roleOrder {
		roles, ok := byRole[role]
		if !ok {
			continue
		}
		label := roleLabels[role]
		if label == "" {
			label = role
		}
		content += "## " + label + "\n\n"
		rows := make([][]string, 0, len(roles))
		for _, r := range roles {
			rows = append(rows, []string{
				"`" + r.FuncName + "`",
				r.RefName,
				r.Provenance.AddonName,
				r.Provenance.Resolution,
				r.Provenance.Confidence,
			})
		}
		content += md.Table([]string{"Function", "Ref Name", "Addon", "Resolution", "Confidence"}, rows)
		content += "\n"
	}

	return writeFile(filepath.Join(dir, "function_lifecycle_roles.md"), content)
}

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
	content += fmt.Sprintf("Generated from addon-api rooted at `%s`. The source corpus contained %d function docs, %d frame docs, %d handler docs, and %d event docs.\n\n", corpus.SourceRoot, len(corpus.Source.Functions), len(corpus.Source.Frames), len(corpus.Source.Handlers), len(corpus.Source.Events))
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
		content += md.Section("Common Handlers", md.BulletList(symbol.CommonHandlers))
		content += md.Section("Common Inherits", md.BulletList(symbol.CommonInherits))
		if len(symbol.CommonChildTypes) > 0 {
			content += md.Section("Common Structural Child Elements", md.BulletList(symbol.CommonChildTypes))
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
	for _, example := range corpus.Source.Examples {
		rows = append(rows, []string{example.Addon, example.Frame, example.Event, example.LuaFunction})
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

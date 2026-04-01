package platform

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"roraddons/tools/api_doc_gen/graph"
)

func ParseAPIRef(root string) (SourceModel, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return SourceModel{}, fmt.Errorf("resolve addon-api root: %w", err)
	}
	if info, err := os.Stat(absRoot); err != nil || !info.IsDir() {
		if err != nil {
			return SourceModel{}, fmt.Errorf("stat addon-api root: %w", err)
		}
		return SourceModel{}, fmt.Errorf("addon-api root is not a directory: %s", absRoot)
	}

	model := SourceModel{Root: graph.NormalizePath(absRoot), LoadedAt: time.Now().UTC()}

	if model.Functions, err = parseFunctionDocs(filepath.Join(absRoot, "lua", "functions")); err != nil {
		return SourceModel{}, err
	}
	if model.Frames, err = parseFrameDocs(filepath.Join(absRoot, "xml", "frames")); err != nil {
		return SourceModel{}, err
	}
	if model.Handlers, err = parseHandlerDocs(filepath.Join(absRoot, "xml", "handlers")); err != nil {
		return SourceModel{}, err
	}
	if model.Events, err = parseEventDocs(filepath.Join(absRoot, "lua", "events")); err != nil {
		return SourceModel{}, err
	}
	if model.Bindings, err = parseBindings(filepath.Join(absRoot, "bindings", "xml_to_lua.md")); err != nil {
		return SourceModel{}, err
	}
	if model.Flows, err = parseFlows(filepath.Join(absRoot, "flows", "execution.md")); err != nil {
		return SourceModel{}, err
	}
	if model.Examples, err = parseExamples(filepath.Join(absRoot, "examples", "index.md")); err != nil {
		return SourceModel{}, err
	}
	if model.Globals, err = parseGlobals(filepath.Join(absRoot, "lua", "globals.md")); err != nil {
		return SourceModel{}, err
	}

	return model, nil
}

func parseFunctionDocs(directory string) ([]FunctionDoc, error) {
	files, err := readMarkdownFiles(directory)
	if err != nil {
		return nil, err
	}
	result := make([]FunctionDoc, 0, len(files))
	for _, path := range files {
		doc, err := parseFunctionDoc(path)
		if err != nil {
			return nil, err
		}
		result = append(result, doc)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Addon == result[j].Addon {
			return result[i].Name < result[j].Name
		}
		return result[i].Addon < result[j].Addon
	})
	return result, nil
}

func parseFunctionDoc(path string) (FunctionDoc, error) {
	lines, err := readMarkdownLines(path)
	if err != nil {
		return FunctionDoc{}, err
	}
	if len(lines) == 0 || !strings.HasPrefix(lines[0], "# Function ") {
		return FunctionDoc{}, fmt.Errorf("unexpected function page format: %s", graph.NormalizePath(path))
	}
	sectionsStart := firstSectionIndex(lines)
	meta := parseBulletMetadata(lines[1:sectionsStart])
	sections := parseSections(lines[sectionsStart:])

	callRows := parseTableRows(sections["Calls"])
	calls := make([]CallDoc, 0, len(callRows))
	for _, row := range callRows {
		line, _ := strconv.Atoi(strings.TrimSpace(row["Line"]))
		rawArguments := strings.TrimSpace(row["Arguments"])
		calls = append(calls, CallDoc{
			Name:      strings.TrimSpace(row["Call"]),
			Line:      line,
			Arguments: splitArgumentList(rawArguments),
			Raw:       rawArguments,
		})
	}

	eventRows := parseTableRows(sections["Event Registrations"])
	events := make([]EventRegistrationDoc, 0, len(eventRows))
	for _, row := range eventRows {
		events = append(events, EventRegistrationDoc{
			Event:   strings.TrimSpace(row["Event"]),
			Scope:   strings.TrimSpace(row["Scope"]),
			Handler: strings.TrimSpace(row["Handler"]),
		})
	}

	return FunctionDoc{
		Name:               strings.TrimSpace(strings.TrimPrefix(lines[0], "# Function ")),
		Addon:              strings.TrimSpace(meta["Addon"]),
		Kind:               strings.TrimSpace(meta["Kind"]),
		Module:             normalizeNone(meta["Module"]),
		Local:              parseBoolish(meta["Local"]),
		Source:             strings.TrimSpace(meta["Source"]),
		Parameters:         parseBulletList(sections["Parameters"]),
		Aliases:            parseBulletList(sections["Aliases"]),
		Calls:              calls,
		EventRegistrations: events,
		StateWrites:        parseBulletList(sections["State Writes"]),
	}, nil
}

func parseFrameDocs(directory string) ([]FrameDoc, error) {
	files, err := readMarkdownFiles(directory)
	if err != nil {
		return nil, err
	}
	result := make([]FrameDoc, 0, len(files))
	for _, path := range files {
		doc, err := parseFrameDoc(path)
		if err != nil {
			return nil, err
		}
		result = append(result, doc)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Addon == result[j].Addon {
			return result[i].Name < result[j].Name
		}
		return result[i].Addon < result[j].Addon
	})
	return result, nil
}

func parseFrameDoc(path string) (FrameDoc, error) {
	lines, err := readMarkdownLines(path)
	if err != nil {
		return FrameDoc{}, err
	}
	if len(lines) == 0 || !strings.HasPrefix(lines[0], "# Frame ") {
		return FrameDoc{}, fmt.Errorf("unexpected frame page format: %s", graph.NormalizePath(path))
	}
	sectionsStart := firstSectionIndex(lines)
	meta := parseBulletMetadata(lines[1:sectionsStart])
	sections := parseSections(lines[sectionsStart:])
	attributes := map[string]string{}
	for _, item := range parseBulletList(sections["Attributes"]) {
		key, value := splitKeyValue(item)
		if key == "" {
			continue
		}
		attributes[key] = value
	}
	handlerRows := parseTableRows(sections["Handlers"])
	handlers := make([]FrameHandlerDoc, 0, len(handlerRows))
	for _, row := range handlerRows {
		handlers = append(handlers, FrameHandlerDoc{Event: strings.TrimSpace(row["Event"]), Function: strings.TrimSpace(row["Function"])})
	}
	return FrameDoc{
		Name:       strings.TrimSpace(strings.TrimPrefix(lines[0], "# Frame ")),
		Addon:      strings.TrimSpace(meta["Addon"]),
		Type:       strings.TrimSpace(meta["Type"]),
		Parent:     normalizeNone(meta["Parent"]),
		Inherits:   normalizeNone(meta["Inherits"]),
		Template:   parseBoolish(meta["Template"]),
		Source:     strings.TrimSpace(meta["Source"]),
		Children:   parseBulletList(sections["Children"]),
		Attributes: attributes,
		Handlers:   handlers,
	}, nil
}

func parseHandlerDocs(directory string) ([]HandlerDoc, error) {
	files, err := readMarkdownFiles(directory)
	if err != nil {
		return nil, err
	}
	result := make([]HandlerDoc, 0, len(files))
	for _, path := range files {
		doc, err := parseHandlerDoc(path)
		if err != nil {
			return nil, err
		}
		result = append(result, doc)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Addon == result[j].Addon {
			if result[i].Frame == result[j].Frame {
				return result[i].Event < result[j].Event
			}
			return result[i].Frame < result[j].Frame
		}
		return result[i].Addon < result[j].Addon
	})
	return result, nil
}

func parseHandlerDoc(path string) (HandlerDoc, error) {
	lines, err := readMarkdownLines(path)
	if err != nil {
		return HandlerDoc{}, err
	}
	if len(lines) == 0 || !strings.HasPrefix(lines[0], "# XML Handler ") {
		return HandlerDoc{}, fmt.Errorf("unexpected handler page format: %s", graph.NormalizePath(path))
	}
	sectionsStart := firstSectionIndex(lines)
	meta := parseBulletMetadata(lines[1:sectionsStart])
	return HandlerDoc{
		Addon:    strings.TrimSpace(meta["Addon"]),
		Frame:    strings.TrimSpace(meta["Frame"]),
		Event:    strings.TrimSpace(meta["Event"]),
		Function: strings.TrimSpace(meta["Function"]),
		Source:   strings.TrimSpace(meta["Source"]),
	}, nil
}

func parseEventDocs(directory string) ([]EventDoc, error) {
	files, err := readMarkdownFiles(directory)
	if err != nil {
		return nil, err
	}
	result := make([]EventDoc, 0, len(files))
	for _, path := range files {
		doc, err := parseEventDoc(path)
		if err != nil {
			return nil, err
		}
		result = append(result, doc)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}

func parseEventDoc(path string) (EventDoc, error) {
	lines, err := readMarkdownLines(path)
	if err != nil {
		return EventDoc{}, err
	}
	if len(lines) == 0 || !strings.HasPrefix(lines[0], "# Event ") {
		return EventDoc{}, fmt.Errorf("unexpected event page format: %s", graph.NormalizePath(path))
	}
	sectionsStart := firstSectionIndex(lines)
	sections := parseSections(lines[sectionsStart:])
	luaRows := parseTableRows(sections["Lua Registrations"])
	registrations := make([]EventUsageDoc, 0, len(luaRows))
	for _, row := range luaRows {
		registrations = append(registrations, EventUsageDoc{
			Addon:     strings.TrimSpace(row["Addon"]),
			Registrar: strings.TrimSpace(row["Registrar"]),
			Handler:   strings.TrimSpace(row["Handler"]),
			Scope:     strings.TrimSpace(row["Scope"]),
		})
	}
	xmlRows := parseTableRows(sections["XML Handlers"])
	xmlHandlers := make([]EventHandlerUsageDoc, 0, len(xmlRows))
	for _, row := range xmlRows {
		xmlHandlers = append(xmlHandlers, EventHandlerUsageDoc{
			Addon:    strings.TrimSpace(row["Addon"]),
			Frame:    strings.TrimSpace(row["Frame"]),
			Function: strings.TrimSpace(row["Function"]),
		})
	}
	return EventDoc{
		Name:             strings.TrimSpace(strings.TrimPrefix(lines[0], "# Event ")),
		LuaRegistrations: registrations,
		XMLHandlers:      xmlHandlers,
		TriggeredBy:      parseBulletList(sections["Triggered By"]),
	}, nil
}

func parseBindings(path string) ([]BindingDoc, error) {
	lines, err := readMarkdownLines(path)
	if err != nil {
		return nil, err
	}
	rows := parseTableRows(lines[1:])
	result := make([]BindingDoc, 0, len(rows))
	for _, row := range rows {
		result = append(result, BindingDoc{
			Addon:       strings.TrimSpace(row["Addon"]),
			Frame:       strings.TrimSpace(row["Frame"]),
			Event:       strings.TrimSpace(row["Event"]),
			XMLFunction: strings.TrimSpace(row["XML Function"]),
			LuaFunction: strings.TrimSpace(row["Lua Function"]),
			Resolved:    parseBoolish(row["Resolved"]),
		})
	}
	return result, nil
}

func parseFlows(path string) ([]FlowDoc, error) {
	lines, err := readMarkdownLines(path)
	if err != nil {
		return nil, err
	}
	sections := parseSections(lines[1:])
	result := make([]FlowDoc, 0, len(sections))
	for addon, sectionLines := range sections {
		steps := make([]FlowStepDoc, 0, len(sectionLines))
		for _, item := range parseBulletList(sectionLines) {
			phase, remainder := splitKeyValue(item)
			if phase == "" {
				continue
			}
			detail, evidence := splitTrailingEvidence(remainder)
			steps = append(steps, FlowStepDoc{Phase: phase, Detail: detail, Evidence: evidence})
		}
		result = append(result, FlowDoc{Addon: addon, Steps: steps})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Addon < result[j].Addon
	})
	return result, nil
}

func parseExamples(path string) ([]ExampleDoc, error) {
	lines, err := readMarkdownLines(path)
	if err != nil {
		return nil, err
	}
	sections := parseSections(lines[1:])
	items := parseBulletList(sections["XML to Lua"])
	result := make([]ExampleDoc, 0, len(items))
	for _, item := range items {
		parts := strings.SplitN(item, ":", 2)
		if len(parts) != 2 {
			continue
		}
		addon := strings.TrimSpace(parts[0])
		mapping := strings.TrimSpace(parts[1])
		segments := strings.SplitN(mapping, " -> ", 2)
		if len(segments) != 2 {
			continue
		}
		left := strings.TrimSpace(segments[0])
		luaFunction := strings.TrimSpace(segments[1])
		frameEvent := strings.SplitN(left, ".", 2)
		if len(frameEvent) != 2 {
			continue
		}
		result = append(result, ExampleDoc{
			Addon:       addon,
			Frame:       strings.TrimSpace(frameEvent[0]),
			Event:       strings.TrimSpace(frameEvent[1]),
			LuaFunction: luaFunction,
		})
	}
	return result, nil
}

func parseGlobals(path string) (GlobalsDoc, error) {
	lines, err := readMarkdownLines(path)
	if err != nil {
		return GlobalsDoc{}, err
	}
	sections := parseSections(lines[1:])
	namespaceRows := parseTableRows(sections["Global Namespaces"])
	namespaces := make([]GlobalNamespaceDoc, 0, len(namespaceRows))
	for _, row := range namespaceRows {
		namespaces = append(namespaces, GlobalNamespaceDoc{Addon: strings.TrimSpace(row["Addon"]), Name: strings.TrimSpace(row["Global"])})
	}
	savedRows := parseTableRows(sections["Saved Variables"])
	saved := make([]SavedVariableDoc, 0, len(savedRows))
	for _, row := range savedRows {
		saved = append(saved, SavedVariableDoc{Addon: strings.TrimSpace(row["Addon"]), Name: strings.TrimSpace(row["Variable"])})
	}
	return GlobalsDoc{Namespaces: namespaces, SavedVariables: saved}, nil
}

func readMarkdownFiles(directory string) ([]string, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("read markdown directory %s: %w", graph.NormalizePath(directory), err)
	}
	files := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(strings.ToLower(entry.Name()), ".md") {
			continue
		}
		if strings.EqualFold(entry.Name(), "index.md") {
			continue
		}
		files = append(files, filepath.Join(directory, entry.Name()))
	}
	sort.Strings(files)
	return files, nil
}

func readMarkdownLines(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read markdown file %s: %w", graph.NormalizePath(path), err)
	}
	content := strings.ReplaceAll(string(bytes), "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	trimmed := strings.TrimSpace(content)
	if trimmed == "" {
		return []string{}, nil
	}
	return strings.Split(trimmed, "\n"), nil
}

func firstSectionIndex(lines []string) int {
	for index, line := range lines {
		if strings.HasPrefix(line, "## ") {
			return index
		}
	}
	return len(lines)
}

func parseBulletMetadata(lines []string) map[string]string {
	meta := map[string]string{}
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "- ") {
			continue
		}
		key, value := splitKeyValue(strings.TrimPrefix(trimmed, "- "))
		if key == "" {
			continue
		}
		meta[key] = value
	}
	return meta
}

func parseSections(lines []string) map[string][]string {
	sections := map[string][]string{}
	current := ""
	for _, line := range lines {
		if strings.HasPrefix(line, "## ") {
			current = strings.TrimSpace(strings.TrimPrefix(line, "## "))
			sections[current] = []string{}
			continue
		}
		if current == "" {
			continue
		}
		sections[current] = append(sections[current], line)
	}
	return sections
}

func parseBulletList(lines []string) []string {
	items := []string{}
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "- ") {
			continue
		}
		item := strings.TrimSpace(strings.TrimPrefix(trimmed, "- "))
		if item == "" || strings.EqualFold(item, "none") {
			continue
		}
		items = append(items, item)
	}
	return items
}

func parseTableRows(lines []string) []map[string]string {
	tableLines := []string{}
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "|") {
			tableLines = append(tableLines, trimmed)
		} else if len(tableLines) > 0 {
			break
		}
	}
	if len(tableLines) < 2 {
		return nil
	}
	headers := parseTableRow(tableLines[0])
	rows := []map[string]string{}
	for _, line := range tableLines[2:] {
		cells := parseTableRow(line)
		row := map[string]string{}
		for index, header := range headers {
			value := ""
			if index < len(cells) {
				value = cells[index]
			}
			row[header] = value
		}
		rows = append(rows, row)
	}
	return rows
}

func parseTableRow(line string) []string {
	trimmed := strings.TrimSpace(line)
	trimmed = strings.TrimPrefix(trimmed, "|")
	trimmed = strings.TrimSuffix(trimmed, "|")
	parts := strings.Split(trimmed, "|")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		result = append(result, strings.TrimSpace(strings.ReplaceAll(part, "\\|", "|")))
	}
	return result
}

func splitKeyValue(value string) (string, string) {
	parts := strings.SplitN(value, ":", 2)
	if len(parts) != 2 {
		return "", strings.TrimSpace(value)
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func splitTrailingEvidence(value string) (string, []string) {
	trimmed := strings.TrimSpace(value)
	if !strings.HasSuffix(trimmed, "]") {
		return trimmed, nil
	}
	index := strings.LastIndex(trimmed, "[")
	if index <= 0 {
		return trimmed, nil
	}
	detail := strings.TrimSpace(trimmed[:index])
	evidenceBody := strings.TrimSuffix(strings.TrimSpace(trimmed[index+1:]), "]")
	if evidenceBody == "" {
		return detail, nil
	}
	parts := strings.Split(evidenceBody, ",")
	evidence := make([]string, 0, len(parts))
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			evidence = append(evidence, item)
		}
	}
	return detail, evidence
}

func splitArgumentList(raw string) []string {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return nil
	}
	parts := []string{}
	var builder strings.Builder
	parenDepth := 0
	braceDepth := 0
	bracketDepth := 0
	inSingle := false
	inDouble := false
	escaped := false
	for _, runeValue := range trimmed {
		char := byte(runeValue)
		if escaped {
			builder.WriteByte(char)
			escaped = false
			continue
		}
		if char == '\\' {
			builder.WriteByte(char)
			escaped = true
			continue
		}
		switch char {
		case '\'':
			if !inDouble {
				inSingle = !inSingle
			}
		case '"':
			if !inSingle {
				inDouble = !inDouble
			}
		case '(':
			if !inSingle && !inDouble {
				parenDepth++
			}
		case ')':
			if !inSingle && !inDouble && parenDepth > 0 {
				parenDepth--
			}
		case '{':
			if !inSingle && !inDouble {
				braceDepth++
			}
		case '}':
			if !inSingle && !inDouble && braceDepth > 0 {
				braceDepth--
			}
		case '[':
			if !inSingle && !inDouble {
				bracketDepth++
			}
		case ']':
			if !inSingle && !inDouble && bracketDepth > 0 {
				bracketDepth--
			}
		case ',':
			if !inSingle && !inDouble && parenDepth == 0 && braceDepth == 0 && bracketDepth == 0 {
				part := strings.TrimSpace(builder.String())
				if part != "" {
					parts = append(parts, part)
				}
				builder.Reset()
				continue
			}
		}
		builder.WriteByte(char)
	}
	if part := strings.TrimSpace(builder.String()); part != "" {
		parts = append(parts, part)
	}
	return parts
}

func parseBoolish(value string) bool {
	trimmed := strings.TrimSpace(strings.ToLower(value))
	return trimmed == "yes" || trimmed == "true"
}

func normalizeNone(value string) string {
	trimmed := strings.TrimSpace(value)
	if strings.EqualFold(trimmed, "none") {
		return ""
	}
	return trimmed
}

package generator

import (
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"roraddons/tools/api_doc_gen/graph"
	md "roraddons/tools/api_doc_gen/templates"
)

func Generate(outputRoot string, corpus graph.Corpus) error {
	if err := prepareOutput(outputRoot); err != nil {
		return err
	}

	if err := writeIndex(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeGlobals(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeModules(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeFunctions(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeEvents(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeState(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeFrames(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeTemplates(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeMixins(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeHandlers(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeBindings(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeExecution(outputRoot, corpus); err != nil {
		return err
	}
	if err := writePatterns(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeConventions(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeExamples(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeLuaAnalysisArtifacts(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeXMLTreeArtifacts(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeXMLLuaLinkArtifacts(outputRoot); err != nil {
		return err
	}
	return writeSectionIndexes(outputRoot, corpus)
}

func prepareOutput(outputRoot string) error {
	managed := []string{"lua", "lua-analysis", "xml", "xml-tree", "xml-lua-links", "bindings", "flows", "patterns", "examples", "meta"}
	for _, name := range managed {
		if err := os.RemoveAll(filepath.Join(outputRoot, name)); err != nil {
			return fmt.Errorf("reset output %s: %w", name, err)
		}
	}
	directories := []string{
		outputRoot,
		filepath.Join(outputRoot, "lua", "modules"),
		filepath.Join(outputRoot, "lua", "functions"),
		filepath.Join(outputRoot, "lua", "events"),
		filepath.Join(outputRoot, "lua", "state"),
		filepath.Join(outputRoot, "xml", "frames"),
		filepath.Join(outputRoot, "xml", "templates"),
		filepath.Join(outputRoot, "xml", "mixins"),
		filepath.Join(outputRoot, "xml", "handlers"),
		filepath.Join(outputRoot, "lua-analysis"),
		filepath.Join(outputRoot, "xml-tree"),
		filepath.Join(outputRoot, "xml-lua-links"),
		filepath.Join(outputRoot, "bindings"),
		filepath.Join(outputRoot, "flows"),
		filepath.Join(outputRoot, "patterns"),
		filepath.Join(outputRoot, "examples"),
		filepath.Join(outputRoot, "meta"),
	}
	for _, directory := range directories {
		if err := os.MkdirAll(directory, 0o755); err != nil {
			return fmt.Errorf("create directory %s: %w", directory, err)
		}
	}
	return nil
}

func writeIndex(outputRoot string, corpus graph.Corpus) error {
	rows := [][]string{}
	for _, addon := range corpus.Addons {
		rows = append(rows, []string{
			addon.Name,
			fmt.Sprintf("%d", len(addon.Functions)),
			fmt.Sprintf("%d", len(addon.Frames)),
			fmt.Sprintf("%d", len(addon.Events)),
			fmt.Sprintf("%d", len(addon.State)),
			addon.Manifest.Type,
		})
	}

	content := "# API Reference\n\n"
	content += fmt.Sprintf("Generated from %d addons rooted at `%s`.\n\n", len(corpus.Addons), graph.NormalizePath(corpus.SourceRoot))
	content += md.Section("Coverage", md.Table([]string{"Addon", "Functions", "Frames", "Events", "State", "Manifest"}, rows))
	content += md.Section("Sections",
		"- [Lua modules](lua/modules/index.md)",
		"- [Lua functions](lua/functions/index.md)",
		"- [Lua events](lua/events/index.md)",
		"- [Lua state](lua/state/index.md)",
		"- [Lua globals](lua/globals.md)",
		"- [XML frames](xml/frames/index.md)",
		"- [XML templates](xml/templates/index.md)",
		"- [XML mixins](xml/mixins/index.md)",
		"- [XML handlers](xml/handlers/index.md)",
		"- [XML to Lua bindings](bindings/xml_to_lua.md)",
		"- [Execution flow](flows/execution.md)",
		"- [Patterns](patterns/addon_patterns.md)",
		"- [Conventions](meta/conventions.md)",
		"- [Examples](examples/index.md)",
	)
	return writeFile(filepath.Join(outputRoot, "index.md"), content)
}

func writeSectionIndexes(outputRoot string, corpus graph.Corpus) error {
	if err := writeLuaOverview(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeXMLOverview(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeModuleIndex(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeFunctionIndex(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeEventIndex(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeStateIndex(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeFrameIndex(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeTemplateIndex(outputRoot, corpus); err != nil {
		return err
	}
	if err := writeMixinIndex(outputRoot, corpus); err != nil {
		return err
	}
	return writeHandlerIndex(outputRoot, corpus)
}

func writeLuaOverview(outputRoot string, corpus graph.Corpus) error {
	rows := [][]string{}
	moduleCount := 0
	functionCount := 0
	eventCount := 0
	stateCount := 0
	for _, addon := range corpus.Addons {
		moduleCount += len(addon.Modules)
		functionCount += len(addon.Functions)
		eventCount += len(addon.Events)
		stateCount += len(addon.State)
		rows = append(rows, []string{addon.Name, fmt.Sprintf("%d", len(addon.Modules)), fmt.Sprintf("%d", len(addon.Functions)), fmt.Sprintf("%d", len(addon.Events)), fmt.Sprintf("%d", len(addon.State))})
	}
	content := "# Lua API\n\n"
	content += fmt.Sprintf("Generated %d module pages, %d function pages, %d event pages, and %d state pages.\n\n", moduleCount, functionCount, eventCount, stateCount)
	content += md.Section("Sections",
		"- [Modules](modules/index.md)",
		"- [Functions](functions/index.md)",
		"- [Events](events/index.md)",
		"- [State](state/index.md)",
		"- [Globals](globals.md)",
	)
	content += md.Section("Coverage", md.Table([]string{"Addon", "Modules", "Functions", "Events", "State"}, rows))
	return writeFile(filepath.Join(outputRoot, "lua", "index.md"), content)
}

func writeXMLOverview(outputRoot string, corpus graph.Corpus) error {
	rows := [][]string{}
	frameCount := 0
	templateCount := 0
	handlerCount := 0
	for _, addon := range corpus.Addons {
		templateTotal := 0
		for _, frame := range addon.Frames {
			if frame.Template {
				templateTotal++
			}
		}
		frameCount += len(addon.Frames)
		templateCount += templateTotal
		handlerCount += len(addon.Handlers)
		rows = append(rows, []string{addon.Name, fmt.Sprintf("%d", len(addon.Frames)), fmt.Sprintf("%d", templateTotal), fmt.Sprintf("%d", len(addon.Handlers))})
	}
	content := "# XML API\n\n"
	content += fmt.Sprintf("Generated %d frame pages, %d template pages, and %d handler pages.\n\n", frameCount, templateCount, handlerCount)
	content += md.Section("Sections",
		"- [Frames](frames/index.md)",
		"- [Templates](templates/index.md)",
		"- [Mixins](mixins/index.md)",
		"- [Handlers](handlers/index.md)",
	)
	content += md.Section("Coverage", md.Table([]string{"Addon", "Frames", "Templates", "Handlers"}, rows))
	return writeFile(filepath.Join(outputRoot, "xml", "index.md"), content)
}

func writeModuleIndex(outputRoot string, corpus graph.Corpus) error {
	sections := []string{}
	for _, addon := range corpus.Addons {
		links := []string{}
		for _, module := range addon.Modules {
			links = append(links, "- "+markdownLink(module.Name, docName(module.Addon, module.Name)))
		}
		sections = append(sections, md.Section(addon.Name, strings.Join(links, "\n")))
	}
	content := "# Lua Modules\n\n"
	content += strings.Join(sections, "")
	return writeFile(filepath.Join(outputRoot, "lua", "modules", "index.md"), content)
}

func writeFunctionIndex(outputRoot string, corpus graph.Corpus) error {
	sections := []string{}
	for _, addon := range corpus.Addons {
		links := []string{}
		for _, function := range addon.Functions {
			links = append(links, "- "+markdownLink(function.Name, docName(function.Addon, function.Name)))
		}
		sections = append(sections, md.Section(addon.Name, strings.Join(links, "\n")))
	}
	content := "# Lua Functions\n\n"
	content += strings.Join(sections, "")
	return writeFile(filepath.Join(outputRoot, "lua", "functions", "index.md"), content)
}

func writeEventIndex(outputRoot string, corpus graph.Corpus) error {
	events := map[string]bool{}
	for _, addon := range corpus.Addons {
		for _, event := range addon.Events {
			events[event.Event] = true
		}
		for _, handler := range addon.Handlers {
			events[handler.Event] = true
		}
	}
	links := []string{}
	for _, eventName := range mapKeys(events) {
		links = append(links, "- "+markdownLink(eventName, docName("event", eventName)))
	}
	content := "# Lua Events\n\n"
	content += md.Section("Events", strings.Join(links, "\n"))
	return writeFile(filepath.Join(outputRoot, "lua", "events", "index.md"), content)
}

func writeStateIndex(outputRoot string, corpus graph.Corpus) error {
	sections := []string{}
	for _, addon := range corpus.Addons {
		links := []string{}
		for _, variable := range addon.State {
			links = append(links, "- "+markdownLink(variable.Name, docName(variable.Addon, variable.Name)))
		}
		sections = append(sections, md.Section(addon.Name, strings.Join(links, "\n")))
	}
	content := "# Lua State\n\n"
	content += strings.Join(sections, "")
	return writeFile(filepath.Join(outputRoot, "lua", "state", "index.md"), content)
}

func writeFrameIndex(outputRoot string, corpus graph.Corpus) error {
	sections := []string{}
	for _, addon := range corpus.Addons {
		links := []string{}
		for _, frame := range addon.Frames {
			links = append(links, "- "+markdownLink(frame.Name, docName(frame.Addon, frame.Name)))
		}
		sections = append(sections, md.Section(addon.Name, strings.Join(links, "\n")))
	}
	content := "# XML Frames\n\n"
	content += strings.Join(sections, "")
	return writeFile(filepath.Join(outputRoot, "xml", "frames", "index.md"), content)
}

func writeTemplateIndex(outputRoot string, corpus graph.Corpus) error {
	sections := []string{}
	for _, addon := range corpus.Addons {
		links := []string{}
		for _, frame := range addon.Frames {
			if !frame.Template {
				continue
			}
			links = append(links, "- "+markdownLink(frame.Name, docName(frame.Addon, frame.Name)))
		}
		sections = append(sections, md.Section(addon.Name, strings.Join(links, "\n")))
	}
	content := "# XML Templates\n\n"
	content += strings.Join(sections, "")
	return writeFile(filepath.Join(outputRoot, "xml", "templates", "index.md"), content)
}

func writeMixinIndex(outputRoot string, corpus graph.Corpus) error {
	targets := map[string]bool{}
	for _, addon := range corpus.Addons {
		for _, frame := range addon.Frames {
			if frame.Inherits == "" {
				continue
			}
			targets[frame.Inherits] = true
		}
	}
	links := []string{}
	for _, mixin := range mapKeys(targets) {
		links = append(links, "- "+markdownLink(mixin, docName("mixin", mixin)))
	}
	content := "# XML Mixins\n\n"
	content += md.Section("Mixins", strings.Join(links, "\n"))
	return writeFile(filepath.Join(outputRoot, "xml", "mixins", "index.md"), content)
}

func writeHandlerIndex(outputRoot string, corpus graph.Corpus) error {
	sections := []string{}
	for _, addon := range corpus.Addons {
		links := []string{}
		for _, handler := range addon.Handlers {
			label := handler.Frame + " / " + handler.Event + " / " + handler.Function
			links = append(links, "- "+markdownLink(label, docName(handler.Addon, handler.Frame+"_"+handler.Event+"_"+handler.Function)))
		}
		sections = append(sections, md.Section(addon.Name, strings.Join(links, "\n")))
	}
	content := "# XML Handlers\n\n"
	content += strings.Join(sections, "")
	return writeFile(filepath.Join(outputRoot, "xml", "handlers", "index.md"), content)
}

func writeGlobals(outputRoot string, corpus graph.Corpus) error {
	rows := [][]string{}
	savedRows := [][]string{}
	for _, addon := range corpus.Addons {
		roots := graph.UniqueStrings(append([]string{addon.Name}, rootModuleNames(addon.Modules)...))
		for _, root := range roots {
			rows = append(rows, []string{addon.Name, root})
		}
		for _, variable := range addon.Manifest.SavedVariables {
			savedRows = append(savedRows, []string{addon.Name, variable})
		}
	}
	content := "# Lua Globals\n\n"
	content += md.Section("Global Namespaces", md.Table([]string{"Addon", "Global"}, rows))
	content += md.Section("Saved Variables", md.Table([]string{"Addon", "Variable"}, savedRows))
	return writeFile(filepath.Join(outputRoot, "lua", "globals.md"), content)
}

func writeModules(outputRoot string, corpus graph.Corpus) error {
	for _, addon := range corpus.Addons {
		functionByModule := map[string][]graph.Function{}
		stateByOwner := map[string][]graph.StateVariable{}
		for _, function := range addon.Functions {
			owner := function.Module
			if owner == "" {
				owner = addon.Name
			}
			functionByModule[owner] = append(functionByModule[owner], function)
		}
		for _, variable := range addon.State {
			owner := variable.Owner
			if owner == "" {
				owner = graph.RootSegment(variable.Name)
			}
			if owner == "" {
				owner = addon.Name
			}
			stateByOwner[owner] = append(stateByOwner[owner], variable)
		}

		for _, module := range addon.Modules {
			content := "# Module " + module.Name + "\n\n"
			content += "- Addon: " + module.Addon + "\n"
			content += "- Kind: " + safeValue(module.Kind) + "\n"
			content += "- Source: `" + module.File + ":" + fmt.Sprintf("%d", module.Line) + "`\n\n"
			content += md.Section("Functions", md.BulletList(functionNames(functionByModule[module.Name])))
			content += md.Section("State", md.BulletList(stateNames(stateByOwner[module.Name])))
			content += md.Section("Aliases", md.BulletList(module.Aliases))
			path := filepath.Join(outputRoot, "lua", "modules", docName(module.Addon, module.Name))
			if err := writeFile(path, content); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeFunctions(outputRoot string, corpus graph.Corpus) error {
	for _, addon := range corpus.Addons {
		for _, function := range addon.Functions {
			rows := [][]string{}
			for _, call := range function.Calls {
				rows = append(rows, []string{call.Name, fmt.Sprintf("%d", call.Line), strings.Join(call.Arguments, ", ")})
			}
			eventRows := [][]string{}
			for _, event := range function.Events {
				eventRows = append(eventRows, []string{event.Event, event.Scope, event.Handler})
			}
			content := "# Function " + function.Name + "\n\n"
			content += "- Addon: " + function.Addon + "\n"
			content += "- Kind: " + safeValue(function.Kind) + "\n"
			content += "- Module: " + safeValue(function.Module) + "\n"
			content += "- Local: " + boolLabel(function.Local) + "\n"
			content += "- Source: `" + function.File + ":" + fmt.Sprintf("%d", function.Line) + "`\n\n"
			content += md.Section("Parameters", md.BulletList(function.Params))
			content += md.Section("Aliases", md.BulletList(function.Aliases))
			content += md.Section("Calls", md.Table([]string{"Call", "Line", "Arguments"}, rows))
			content += md.Section("Event Registrations", md.Table([]string{"Event", "Scope", "Handler"}, eventRows))
			content += md.Section("State Writes", md.BulletList(function.StateWrites))
			path := filepath.Join(outputRoot, "lua", "functions", docName(function.Addon, function.Name))
			if err := writeFile(path, content); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeEvents(outputRoot string, corpus graph.Corpus) error {
	eventMap := map[string][]graph.EventRegistration{}
	handlerMap := map[string][]graph.XMLHandler{}
	triggerMap := map[string][]string{}
	for _, addon := range corpus.Addons {
		for _, event := range addon.Events {
			eventMap[event.Event] = append(eventMap[event.Event], event)
		}
		for _, handler := range addon.Handlers {
			handlerMap[handler.Event] = append(handlerMap[handler.Event], handler)
		}
	}
	for _, edge := range corpus.Graph.Edges {
		if edge.Type != graph.EdgeTriggered {
			continue
		}
		eventName := strings.TrimPrefix(edge.To, "event:")
		triggerMap[eventName] = append(triggerMap[eventName], strings.TrimPrefix(edge.From, "function:"))
	}

	allEvents := graph.UniqueStrings(append(mapKeys(eventMap), mapKeys(handlerMap)...))
	for _, eventName := range allEvents {
		registrations := eventMap[eventName]
		handlers := handlerMap[eventName]
		registrationRows := [][]string{}
		for _, event := range registrations {
			registrationRows = append(registrationRows, []string{event.Addon, event.Registrar, event.Handler, event.Scope})
		}
		handlerRows := [][]string{}
		for _, handler := range handlers {
			handlerRows = append(handlerRows, []string{handler.Addon, handler.Frame, handler.Function})
		}
		content := "# Event " + eventName + "\n\n"
		content += md.Section("Lua Registrations", md.Table([]string{"Addon", "Registrar", "Handler", "Scope"}, registrationRows))
		content += md.Section("XML Handlers", md.Table([]string{"Addon", "Frame", "Function"}, handlerRows))
		content += md.Section("Triggered By", md.BulletList(graph.UniqueStrings(triggerMap[eventName])))
		path := filepath.Join(outputRoot, "lua", "events", docName("event", eventName))
		if err := writeFile(path, content); err != nil {
			return err
		}
	}
	return nil
}

func writeState(outputRoot string, corpus graph.Corpus) error {
	for _, addon := range corpus.Addons {
		for _, variable := range addon.State {
			content := "# State " + variable.Name + "\n\n"
			content += "- Addon: " + variable.Addon + "\n"
			content += "- Owner: " + safeValue(variable.Owner) + "\n"
			content += "- Type: " + safeValue(variable.ValueType) + "\n"
			content += "- Saved: " + boolLabel(variable.Saved) + "\n"
			content += "- Scope: " + safeValue(variable.Scope) + "\n"
			content += "- Source: `" + variable.File + ":" + fmt.Sprintf("%d", variable.Line) + "`\n"
			path := filepath.Join(outputRoot, "lua", "state", docName(variable.Addon, variable.Name))
			if err := writeFile(path, content); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeFrames(outputRoot string, corpus graph.Corpus) error {
	handlerByFrame := map[string][]graph.XMLHandler{}
	for _, addon := range corpus.Addons {
		for _, handler := range addon.Handlers {
			key := addon.Name + "|" + handler.Frame
			handlerByFrame[key] = append(handlerByFrame[key], handler)
		}
	}
	for _, addon := range corpus.Addons {
		for _, frame := range addon.Frames {
			content := "# Frame " + frame.Name + "\n\n"
			content += "- Addon: " + frame.Addon + "\n"
			content += "- Resolved Name: " + frame.Name + "\n"
			if strings.TrimSpace(frame.RawName) != "" && frame.RawName != frame.Name {
				content += "- Raw Name: " + frame.RawName + "\n"
			}
			content += "- Type: " + frame.Type + "\n"
			content += "- Parent: " + safeValue(frame.Parent) + "\n"
			content += "- Parent Type: " + safeValue(frame.ParentType) + "\n"
			content += "- Inherits: " + safeValue(frame.Inherits) + "\n"
			content += "- Template: " + boolLabel(frame.Template) + "\n"
			content += "- Source: `" + frame.File + ":" + fmt.Sprintf("%d", frame.Line) + "`\n\n"
			content += md.Section("Children", md.BulletList(frame.Children))
			content += md.Section("Child Element Types", md.BulletList(frame.ChildElementTypes))
			structuralChildLines := make([]string, 0, len(frame.StructuralChildTypes))
			for _, childType := range frame.StructuralChildTypes {
				if attrKeys, ok := frame.StructuralChildAttrKeys[childType]; ok && len(attrKeys) > 0 {
					structuralChildLines = append(structuralChildLines, childType+": "+strings.Join(attrKeys, ", "))
				} else {
					structuralChildLines = append(structuralChildLines, childType)
				}
			}
			content += md.Section("Structural Child Types", md.BulletList(structuralChildLines))
			if frame.CompositionSnippet != "" {
				content += md.Section("Composition Pattern", "```xml\n"+frame.CompositionSnippet+"\n```\n")
			}
			content += md.Section("Attributes", md.BulletList(attributeLines(frame.Attributes)))
			content += md.Section("Handlers", md.Table([]string{"Event", "Function"}, handlerRows(handlerByFrame[addon.Name+"|"+frame.Name])))
			path := filepath.Join(outputRoot, "xml", "frames", docName(frame.Addon, frame.Name))
			if err := writeFile(path, content); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeTemplates(outputRoot string, corpus graph.Corpus) error {
	for _, addon := range corpus.Addons {
		for _, frame := range addon.Frames {
			if !frame.Template {
				continue
			}
			content := "# Template " + frame.Name + "\n\n"
			content += "- Addon: " + frame.Addon + "\n"
			content += "- Resolved Name: " + frame.Name + "\n"
			if strings.TrimSpace(frame.RawName) != "" && frame.RawName != frame.Name {
				content += "- Raw Name: " + frame.RawName + "\n"
			}
			content += "- Type: " + frame.Type + "\n"
			content += "- Inherits: " + safeValue(frame.Inherits) + "\n"
			content += "- Source: `" + frame.File + ":" + fmt.Sprintf("%d", frame.Line) + "`\n\n"
			content += md.Section("Attributes", md.BulletList(attributeLines(frame.Attributes)))
			content += md.Section("Children", md.BulletList(frame.Children))
			structuralChildLines := make([]string, 0, len(frame.StructuralChildTypes))
			for _, childType := range frame.StructuralChildTypes {
				if attrKeys, ok := frame.StructuralChildAttrKeys[childType]; ok && len(attrKeys) > 0 {
					structuralChildLines = append(structuralChildLines, childType+": "+strings.Join(attrKeys, ", "))
				} else {
					structuralChildLines = append(structuralChildLines, childType)
				}
			}
			content += md.Section("Structural Child Types", md.BulletList(structuralChildLines))
			if frame.CompositionSnippet != "" {
				content += md.Section("Composition Pattern", "```xml\n"+frame.CompositionSnippet+"\n```\n")
			}
			path := filepath.Join(outputRoot, "xml", "templates", docName(frame.Addon, frame.Name))
			if err := writeFile(path, content); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeMixins(outputRoot string, corpus graph.Corpus) error {
	targets := map[string][]string{}
	for _, addon := range corpus.Addons {
		for _, frame := range addon.Frames {
			if frame.Inherits == "" {
				continue
			}
			targets[frame.Inherits] = append(targets[frame.Inherits], addon.Name+": "+frame.Name)
		}
	}
	for mixin, users := range targets {
		content := "# Mixin " + mixin + "\n\n"
		content += md.Section("Used By", md.BulletList(graph.UniqueStrings(users)))
		path := filepath.Join(outputRoot, "xml", "mixins", docName("mixin", mixin))
		if err := writeFile(path, content); err != nil {
			return err
		}
	}
	return nil
}

func writeHandlers(outputRoot string, corpus graph.Corpus) error {
	for _, addon := range corpus.Addons {
		for _, handler := range addon.Handlers {
			content := "# XML Handler " + handler.Frame + " / " + handler.Event + "\n\n"
			content += "- Addon: " + handler.Addon + "\n"
			content += "- Frame: " + safeValue(handler.Frame) + "\n"
			content += "- Event: " + safeValue(handler.Event) + "\n"
			content += "- Function: " + safeValue(handler.Function) + "\n"
			content += "- Source: `" + handler.File + ":" + fmt.Sprintf("%d", handler.Line) + "`\n"
			path := filepath.Join(outputRoot, "xml", "handlers", docName(handler.Addon, handler.Frame+"_"+handler.Event+"_"+handler.Function))
			if err := writeFile(path, content); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeBindings(outputRoot string, corpus graph.Corpus) error {
	rows := [][]string{}
	for _, binding := range corpus.Bindings {
		rows = append(rows, []string{binding.Addon, binding.Frame, binding.Event, binding.XMLFunction, binding.LuaFunction, boolLabel(binding.Resolved)})
	}
	content := "# XML to Lua Bindings\n\n"
	content += md.Table([]string{"Addon", "Frame", "Event", "XML Function", "Lua Function", "Resolved"}, rows)
	return writeFile(filepath.Join(outputRoot, "bindings", "xml_to_lua.md"), content)
}

func writeExecution(outputRoot string, corpus graph.Corpus) error {
	content := "# Execution Flow\n\n"
	for _, flow := range corpus.Normalized.ExecutionFlows {
		lines := []string{}
		for _, step := range flow.Steps {
			line := "- " + step.Phase + ": " + step.Detail
			if len(step.Evidence) > 0 {
				line += " [" + strings.Join(step.Evidence, ", ") + "]"
			}
			lines = append(lines, line)
		}
		content += md.Section(flow.Addon, strings.Join(lines, "\n"))
	}
	return writeFile(filepath.Join(outputRoot, "flows", "execution.md"), content)
}

func writePatterns(outputRoot string, corpus graph.Corpus) error {
	content := "# Addon Patterns\n\n"
	for _, pattern := range corpus.Normalized.Patterns {
		lines := []string{"- Category: " + safeValue(pattern.Category), "- Summary: " + safeValue(pattern.Summary), "- Evidence:"}
		for _, evidence := range pattern.Evidence {
			lines = append(lines, "  - "+evidence)
		}
		content += md.Section(pattern.Name, strings.Join(lines, "\n"))
	}
	content += md.Section("Normalized Concepts", md.Table([]string{"Concept", "Kind", "Members", "Addons", "Summary"}, conceptRows(corpus.Normalized.ConceptGroups)))
	return writeFile(filepath.Join(outputRoot, "patterns", "addon_patterns.md"), content)
}

func writeConventions(outputRoot string, corpus graph.Corpus) error {
	content := "# Conventions\n\n"
	rows := [][]string{}
	for _, convention := range corpus.Normalized.NamingConventions {
		rows = append(rows, []string{convention.Rule, strings.Join(convention.Evidence, ", ")})
	}
	content += md.Section("Naming", md.Table([]string{"Rule", "Evidence"}, rows))
	content += md.Section("Common Events", md.BulletList(corpus.Normalized.CommonEvents))
	content += md.Section("Common Modules", md.BulletList(corpus.Normalized.CommonModules))
	content += md.Section("Common State Owners", md.BulletList(corpus.Normalized.CommonStateOwners))
	content += md.Section("Common Frame Parents", md.BulletList(corpus.Normalized.CommonFrameParents))
	return writeFile(filepath.Join(outputRoot, "meta", "conventions.md"), content)
}

func writeExamples(outputRoot string, corpus graph.Corpus) error {
	lines := []string{}
	for _, binding := range firstNBindings(corpus.Bindings, 12) {
		lines = append(lines, "- "+binding.Addon+": "+binding.Frame+"."+binding.Event+" -> "+binding.LuaFunction)
	}
	content := "# Examples\n\n"
	content += md.Section("XML to Lua", strings.Join(lines, "\n"))
	return writeFile(filepath.Join(outputRoot, "examples", "index.md"), content)
}

func markdownLink(label string, target string) string {
	return "[" + label + "](" + target + ")"
}

func writeFile(path string, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(strings.TrimSpace(content)+"\n"), 0o644)
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

func rootModuleNames(modules []graph.Module) []string {
	result := []string{}
	for _, module := range modules {
		result = append(result, graph.RootSegment(module.Name))
	}
	return graph.UniqueStrings(result)
}

func functionNames(functions []graph.Function) []string {
	result := []string{}
	for _, function := range functions {
		result = append(result, function.Name)
	}
	return graph.UniqueStrings(result)
}

func stateNames(state []graph.StateVariable) []string {
	result := []string{}
	for _, variable := range state {
		result = append(result, variable.Name)
	}
	return graph.UniqueStrings(result)
}

func attributeLines(attributes map[string]string) []string {
	if len(attributes) == 0 {
		return nil
	}
	keys := make([]string, 0, len(attributes))
	for key := range attributes {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	result := make([]string, 0, len(keys))
	for _, key := range keys {
		result = append(result, key+": "+attributes[key])
	}
	return result
}

func handlerRows(handlers []graph.XMLHandler) [][]string {
	rows := [][]string{}
	for _, handler := range handlers {
		rows = append(rows, []string{handler.Event, handler.Function})
	}
	return rows
}

func conceptRows(groups []graph.ConceptGroup) [][]string {
	rows := [][]string{}
	for _, group := range groups {
		rows = append(rows, []string{group.Canonical, group.Kind, strings.Join(group.Members, ", "), strings.Join(group.Addons, ", "), group.Summary})
	}
	return rows
}

func firstNBindings(bindings []graph.Binding, limit int) []graph.Binding {
	if len(bindings) <= limit {
		return bindings
	}
	return bindings[:limit]
}

func mapKeys[T any](values map[string]T) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func safeValue(value string) string {
	if strings.TrimSpace(value) == "" {
		return "none"
	}
	return value
}

func boolLabel(value bool) string {
	if value {
		return "yes"
	}
	return "no"
}

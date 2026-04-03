package platform

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"roraddons/tools/api_doc_gen/deep_analysis"
	"roraddons/tools/api_doc_gen/graph"
)

type symbolCatalog struct {
	entries map[string]DocLink
	byName  map[string][]string
}

type graphEdgeAccumulator struct {
	From     string
	To       string
	Type     string
	Weight   int
	Evidence map[string]bool
}

type relationAccumulator struct {
	LeftID   string
	RightID  string
	Weight   int
	Evidence map[string]bool
}

type eventFlowAccumulator struct {
	EventID    string
	HandlerIDs map[string]bool
	UIIDs      map[string]bool
	Evidence   map[string]bool
	Weight     int
}

// maxCombinationContextSize caps the number of symbol IDs considered per context
// in recordCombination. Without a cap, a function with N context symbols produces
// N*(N-1)/2 relation pairs (e.g. 20 symbols → 190 pairs), inflating commonly_used_with edges.
const maxCombinationContextSize = 6

// minCommonlyUsedWithWeight is the minimum accumulator weight to emit a commonly_used_with edge.
// Raised from 2 to 4 to suppress low-signal combinatorial pairs.
const minCommonlyUsedWithWeight = 4

// maxCommonlyUsedWithPerNode caps the number of commonly_used_with edges emitted per
// source node after weight-based sorting. Prevents hub nodes from generating dozens of
// low-quality associations.
const maxCommonlyUsedWithPerNode = 8

func enrichSemanticArtifacts(corpus *Corpus) {
	catalog := buildSymbolCatalog(*corpus)

	// Apply deep analysis inference to symbols (Priority 2: return types, Priority 3: arguments)
	enrichSymbolsWithAnalysis(corpus, corpus.Source)

	graphData, relations, links := buildSemanticGraph(*corpus, catalog)
	patternPages := buildPatternPages(*corpus, catalog, relations)
	corpus.SymbolLinks = links
	corpus.APIGraph = graphData
	corpus.Relations = relations
	corpus.PatternPages = patternPages
	corpus.NavigationTree = buildNavigationTree(*corpus, catalog, patternPages)
	corpus.Sidebar = buildSidebar(*corpus, catalog, patternPages)
}

func buildSymbolCatalog(corpus Corpus) symbolCatalog {
	catalog := symbolCatalog{entries: map[string]DocLink{}, byName: map[string][]string{}}
	add := func(entry DocLink, aliases ...string) {
		catalog.entries[entry.ID] = entry
		catalog.byName[normalizeLookupKey(entry.Label)] = appendUniqueString(catalog.byName[normalizeLookupKey(entry.Label)], entry.ID)
		for _, alias := range aliases {
			if strings.TrimSpace(alias) == "" {
				continue
			}
			catalog.byName[normalizeLookupKey(alias)] = appendUniqueString(catalog.byName[normalizeLookupKey(alias)], entry.ID)
		}
	}
	for _, symbol := range corpus.GlobalFunctions {
		add(DocLink{
			ID:         globalFunctionID(symbol.Name),
			Label:      symbol.Name,
			Type:       "function",
			Category:   symbol.Category,
			Path:       slashPath("globals", "functions", docName("global", symbol.Name)),
			Confidence: symbol.Confidence,
			Score:      symbol.Score,
			Summary:    symbol.Description,
		}, symbol.Aliases...)
	}
	for _, symbol := range corpus.WindowFunctions {
		add(DocLink{
			ID:         windowFunctionID(symbol.Name),
			Label:      symbol.Name,
			Type:       "function",
			Category:   symbol.Category,
			Path:       slashPath("window_api", "functions", docName("window", symbol.Name)),
			Confidence: symbol.Confidence,
			Score:      symbol.Score,
			Summary:    symbol.Description,
		}, symbol.Aliases...)
	}
	for _, symbol := range corpus.GlobalTables {
		add(DocLink{
			ID:         tableID(symbol.Name),
			Label:      symbol.Name,
			Type:       "data_structure",
			Category:   "Global Table",
			Path:       slashPath("globals", "tables", docName("table", symbol.Name)),
			Confidence: symbol.Confidence,
			Score:      symbol.Score,
			Summary:    symbol.Description,
		})
	}
	for _, symbol := range corpus.Constants {
		add(DocLink{
			ID:         constantID(symbol.Name),
			Label:      symbol.Name,
			Type:       "data_structure",
			Category:   "Constant",
			Path:       slashPath("globals", "constants", docName("constant", symbol.Name)),
			Confidence: symbol.Confidence,
			Score:      symbol.Score,
			Summary:    symbol.Description,
		})
	}
	for _, symbol := range corpus.ElementTypes {
		add(DocLink{
			ID:         elementTypeID(symbol.Name),
			Label:      symbol.Name,
			Type:       "ui_element",
			Category:   "XML Element Type",
			Path:       slashPath("xml", "element_types", docName("element", symbol.Name)),
			Confidence: symbol.Confidence,
			Score:      symbol.Score,
			Summary:    symbol.Description,
		})
	}
	for _, symbol := range corpus.XMLHandlers {
		add(DocLink{
			ID:         xmlHandlerID(symbol.Name),
			Label:      symbol.Name,
			Type:       "xml_event",
			Category:   "XML Event",
			Path:       slashPath("xml", "handlers", docName("handler", symbol.Name)),
			Confidence: symbol.Confidence,
			Score:      symbol.Score,
			Summary:    symbol.Description,
		})
	}
	for _, symbol := range corpus.GameEvents {
		add(DocLink{
			ID:         gameEventID(symbol.Name),
			Label:      symbol.Name,
			Type:       "event",
			Category:   symbol.Category,
			Path:       slashPath("events", "game_events", docName("game_event", symbol.Name)),
			Confidence: symbol.Confidence,
			Score:      symbol.Score,
			Summary:    symbol.Description,
		})
	}
	for _, symbol := range corpus.WindowEvents {
		add(DocLink{
			ID:         windowEventID(symbol.Name),
			Label:      symbol.Name,
			Type:       "event",
			Category:   symbol.Category,
			Path:       slashPath("events", "window_events", docName("window_event", symbol.Name)),
			Confidence: symbol.Confidence,
			Score:      symbol.Score,
			Summary:    symbol.Description,
		})
	}
	for _, symbol := range corpus.SystemDataFields {
		add(DocLink{
			ID:         systemFieldID(symbol.Name),
			Label:      symbol.Name,
			Type:       "data_structure",
			Category:   "SystemData Field",
			Path:       slashPath("systemdata", "fields", docName("systemdata", symbol.Name)),
			Confidence: symbol.Confidence,
			Score:      symbol.Score,
			Summary:    symbol.Description,
		})
	}
	for _, symbol := range corpus.GameDataFields {
		add(DocLink{
			ID:         gameFieldID(symbol.Name),
			Label:      symbol.Name,
			Type:       "data_structure",
			Category:   "GameData Field",
			Path:       slashPath("gamedata", "fields", docName("gamedata", symbol.Name)),
			Confidence: symbol.Confidence,
			Score:      symbol.Score,
			Summary:    symbol.Description,
		})
	}
	return catalog
}

func buildSemanticGraph(corpus Corpus, catalog symbolCatalog) (APIGraph, RelationReport, map[string]SemanticLinks) {
	frameTypes := map[string]string{}
	sourceFunctionsByKey := map[string]FunctionDoc{}
	invokersByCaller := map[string][]DocLink{}
	uiLinksByCaller := map[string][]DocLink{}
	for _, frame := range corpus.Source.Frames {
		frameTypes[frameLookupKey(frame.Addon, frame.Name)] = frame.Type
	}
	for _, doc := range corpus.Source.Functions {
		sourceFunctionsByKey[callerKey(doc.Addon, doc.Name)] = doc
	}

	edgeAccumulators := map[string]*graphEdgeAccumulator{}
	relationAccumulators := map[string]*relationAccumulator{}
	eventFlows := map[string]*eventFlowAccumulator{}

	addEdge := func(fromID string, toID string, edgeType string, evidence string) {
		if fromID == "" || toID == "" || fromID == toID {
			return
		}
		key := edgeType + "|" + fromID + "|" + toID
		accumulator, ok := edgeAccumulators[key]
		if !ok {
			accumulator = &graphEdgeAccumulator{From: fromID, To: toID, Type: edgeType, Evidence: map[string]bool{}}
			edgeAccumulators[key] = accumulator
		}
		accumulator.Weight++
		if strings.TrimSpace(evidence) != "" {
			accumulator.Evidence[evidence] = true
		}
	}

	recordCombination := func(symbolIDs []string, evidence string) {
		uniqueIDs := graph.UniqueStrings(symbolIDs)
		if len(uniqueIDs) < 2 {
			return
		}
		// Cap context size to prevent combinatorial pair explosion.
		if len(uniqueIDs) > maxCombinationContextSize {
			uniqueIDs = uniqueIDs[:maxCombinationContextSize]
		}
		sort.Strings(uniqueIDs)
		for leftIndex := 0; leftIndex < len(uniqueIDs); leftIndex++ {
			for rightIndex := leftIndex + 1; rightIndex < len(uniqueIDs); rightIndex++ {
				key := uniqueIDs[leftIndex] + "|" + uniqueIDs[rightIndex]
				accumulator, ok := relationAccumulators[key]
				if !ok {
					accumulator = &relationAccumulator{LeftID: uniqueIDs[leftIndex], RightID: uniqueIDs[rightIndex], Evidence: map[string]bool{}}
					relationAccumulators[key] = accumulator
				}
				accumulator.Weight++
				if strings.TrimSpace(evidence) != "" {
					accumulator.Evidence[evidence] = true
				}
			}
		}
	}

	ensureEventFlow := func(eventID string) *eventFlowAccumulator {
		accumulator, ok := eventFlows[eventID]
		if !ok {
			accumulator = &eventFlowAccumulator{EventID: eventID, HandlerIDs: map[string]bool{}, UIIDs: map[string]bool{}, Evidence: map[string]bool{}}
			eventFlows[eventID] = accumulator
		}
		return accumulator
	}

	appendCallerLink := func(target map[string][]DocLink, addon string, caller string, link DocLink) {
		key := callerKey(addon, caller)
		target[key] = appendUniqueLink(target[key], link)
	}

	for _, doc := range corpus.Source.Functions {
		sourceEntry, ok := catalog.lookup(doc.Name, "function")
		if !ok {
			continue
		}
		contextIDs := []string{sourceEntry.ID}
		for _, call := range doc.Calls {
			targetEntry, found := catalog.lookup(call.Name, "function")
			if !found {
				// Qualified call names (e.g. "SystemData.GetCurrentCareer") may not match
				// the catalog entry label directly. Try the last dot-component as a fallback.
				if dotIdx := strings.LastIndex(call.Name, "."); dotIdx >= 0 {
					targetEntry, found = catalog.lookup(call.Name[dotIdx+1:], "function")
				}
			}
			if !found {
				continue
			}
			addEdge(sourceEntry.ID, targetEntry.ID, "calls", fmt.Sprintf("%s: %s -> %s", doc.Addon, doc.Name, safeValue(call.Raw)))
			contextIDs = append(contextIDs, targetEntry.ID)
		}
		for _, dataEntry := range extractDataLinks(doc, catalog) {
			addEdge(sourceEntry.ID, dataEntry.ID, "reads_data", fmt.Sprintf("%s: %s reads %s", doc.Addon, doc.Name, dataEntry.Label))
			contextIDs = append(contextIDs, dataEntry.ID)
		}
		for _, stateEntry := range extractStateLinks(doc.StateWrites, catalog) {
			addEdge(sourceEntry.ID, stateEntry.ID, "writes_state", fmt.Sprintf("%s: %s writes %s", doc.Addon, doc.Name, stateEntry.Label))
			contextIDs = append(contextIDs, stateEntry.ID)
		}
		recordCombination(contextIDs, doc.Addon+": "+doc.Name)
	}

	for _, eventDoc := range corpus.Source.Events {
		eventEntry, ok := catalog.lookup(eventDoc.Name, "event")
		if !ok {
			continue
		}
		contextIDs := []string{eventEntry.ID}
		for _, usage := range eventDoc.LuaRegistrations {
			handlerEntry, found := catalog.lookup(usage.Handler, "function")
			if found {
				evidence := fmt.Sprintf("%s: %s -> %s", usage.Addon, usage.Registrar, usage.Handler)
				addEdge(handlerEntry.ID, eventEntry.ID, "triggered_by", evidence)
				contextIDs = append(contextIDs, handlerEntry.ID)
				flow := ensureEventFlow(eventEntry.ID)
				flow.HandlerIDs[handlerEntry.ID] = true
				flow.Evidence[evidence] = true
				flow.Weight++
			}
			appendCallerLink(invokersByCaller, usage.Addon, usage.Handler, eventEntry)
			registrarEntry, found := catalog.lookup(usage.Registrar, "function")
			if found {
				contextIDs = append(contextIDs, registrarEntry.ID)
			}
		}
		for _, xmlUsage := range eventDoc.XMLHandlers {
			handlerEntry, found := catalog.lookup(xmlUsage.Function, "function")
			if found {
				evidence := fmt.Sprintf("%s: %s -> %s", xmlUsage.Addon, eventDoc.Name, xmlUsage.Function)
				addEdge(handlerEntry.ID, eventEntry.ID, "triggered_by", evidence)
				contextIDs = append(contextIDs, handlerEntry.ID)
				flow := ensureEventFlow(eventEntry.ID)
				flow.HandlerIDs[handlerEntry.ID] = true
				flow.Evidence[evidence] = true
				flow.Weight++
			}
			appendCallerLink(invokersByCaller, xmlUsage.Addon, xmlUsage.Function, eventEntry)
		}
		for _, trigger := range eventDoc.TriggeredBy {
			triggerEntry, found := catalog.lookup(trigger, "function")
			if !found {
				continue
			}
			evidence := fmt.Sprintf("%s triggered by %s", eventDoc.Name, triggerEntry.Label)
			addEdge(eventEntry.ID, triggerEntry.ID, "triggered_by", evidence)
			contextIDs = append(contextIDs, triggerEntry.ID)
		}
		recordCombination(contextIDs, eventDoc.Name)
	}

	for _, binding := range corpus.Source.Bindings {
		xmlHandlerEntry, hasXMLHandler := catalog.lookup(binding.Event, "xml_event")
		functionEntry, hasFunction := catalog.lookup(binding.LuaFunction, "function")
		eventEntry, hasEvent := catalog.lookup(binding.Event, "xml_event", "event")
		elementEntry, hasElement := catalog.lookup(frameTypes[frameLookupKey(binding.Addon, binding.Frame)], "ui_element")
		evidence := fmt.Sprintf("%s: %s.%s -> %s", binding.Addon, binding.Frame, binding.Event, binding.LuaFunction)
		contextIDs := []string{}
		if hasXMLHandler {
			contextIDs = append(contextIDs, xmlHandlerEntry.ID)
			appendCallerLink(invokersByCaller, binding.Addon, binding.LuaFunction, xmlHandlerEntry)
		}
		if hasFunction {
			contextIDs = append(contextIDs, functionEntry.ID)
		}
		if hasEvent {
			contextIDs = append(contextIDs, eventEntry.ID)
		}
		if hasElement {
			contextIDs = append(contextIDs, elementEntry.ID)
			uiLinksByCaller[callerKey(binding.Addon, binding.LuaFunction)] = appendUniqueLink(uiLinksByCaller[callerKey(binding.Addon, binding.LuaFunction)], elementEntry)
		}
		if hasXMLHandler && hasFunction {
			addEdge(xmlHandlerEntry.ID, functionEntry.ID, "calls", evidence)
			addEdge(functionEntry.ID, xmlHandlerEntry.ID, "bound_from_xml", evidence)
		}
		if hasFunction && hasEvent {
			addEdge(functionEntry.ID, eventEntry.ID, "triggered_by", evidence)
			flow := ensureEventFlow(eventEntry.ID)
			flow.HandlerIDs[functionEntry.ID] = true
			flow.Evidence[evidence] = true
			flow.Weight++
		}
		if hasFunction && hasElement {
			addEdge(functionEntry.ID, elementEntry.ID, "updates_ui", evidence)
		}
		if hasEvent && hasElement {
			flow := ensureEventFlow(eventEntry.ID)
			flow.UIIDs[elementEntry.ID] = true
			flow.Evidence[evidence] = true
			flow.Weight++
		}
		recordCombination(contextIDs, evidence)
	}

	for _, handler := range corpus.Source.Handlers {
		xmlHandlerEntry, hasXMLHandler := catalog.lookup(handler.Event, "xml_event")
		functionEntry, hasFunction := catalog.lookup(handler.Function, "function")
		eventEntry, hasEvent := catalog.lookup(handler.Event, "xml_event", "event")
		elementEntry, hasElement := catalog.lookup(frameTypes[frameLookupKey(handler.Addon, handler.Frame)], "ui_element")
		evidence := fmt.Sprintf("%s: %s.%s -> %s", handler.Addon, handler.Frame, handler.Event, handler.Function)
		contextIDs := []string{}
		if hasXMLHandler {
			contextIDs = append(contextIDs, xmlHandlerEntry.ID)
			appendCallerLink(invokersByCaller, handler.Addon, handler.Function, xmlHandlerEntry)
		}
		if hasFunction {
			contextIDs = append(contextIDs, functionEntry.ID)
		}
		if hasEvent {
			contextIDs = append(contextIDs, eventEntry.ID)
		}
		if hasElement {
			contextIDs = append(contextIDs, elementEntry.ID)
			uiLinksByCaller[callerKey(handler.Addon, handler.Function)] = appendUniqueLink(uiLinksByCaller[callerKey(handler.Addon, handler.Function)], elementEntry)
		}
		if hasXMLHandler && hasFunction {
			addEdge(xmlHandlerEntry.ID, functionEntry.ID, "calls", evidence)
			addEdge(functionEntry.ID, xmlHandlerEntry.ID, "bound_from_xml", evidence)
		}
		if hasFunction && hasEvent {
			addEdge(functionEntry.ID, eventEntry.ID, "triggered_by", evidence)
			flow := ensureEventFlow(eventEntry.ID)
			flow.HandlerIDs[functionEntry.ID] = true
			flow.Evidence[evidence] = true
			flow.Weight++
		}
		if hasFunction && hasElement {
			addEdge(functionEntry.ID, elementEntry.ID, "updates_ui", evidence)
		}
		if hasEvent && hasElement {
			flow := ensureEventFlow(eventEntry.ID)
			flow.UIIDs[elementEntry.ID] = true
			flow.Evidence[evidence] = true
			flow.Weight++
		}
		recordCombination(contextIDs, evidence)
	}

	for _, example := range corpus.Source.Examples {
		functionEntry, hasFunction := catalog.lookup(example.LuaFunction, "function")
		eventEntry, hasEvent := catalog.lookup(example.Event, "xml_event", "event")
		xmlHandlerEntry, hasXMLHandler := catalog.lookup(example.Event, "xml_event")
		elementEntry, hasElement := catalog.lookup(frameTypes[frameLookupKey(example.Addon, example.Frame)], "ui_element")
		contextIDs := []string{}
		evidence := fmt.Sprintf("%s: %s.%s -> %s", example.Addon, example.Frame, example.Event, example.LuaFunction)
		if hasFunction {
			contextIDs = append(contextIDs, functionEntry.ID)
		}
		if hasEvent {
			contextIDs = append(contextIDs, eventEntry.ID)
		}
		if hasXMLHandler {
			contextIDs = append(contextIDs, xmlHandlerEntry.ID)
		}
		if hasElement {
			contextIDs = append(contextIDs, elementEntry.ID)
		}
		recordCombination(contextIDs, evidence)
	}

	for _, symbol := range corpus.ElementTypes {
		entry, found := catalog.entries[elementTypeID(symbol.Name)]
		if !found {
			continue
		}
		contextIDs := []string{entry.ID}

		for _, parent := range symbol.ParentRefs {
			parentEntry, matched := catalog.lookup(parent.Tag, "ui_element")
			if !matched {
				continue
			}
			evidence := fmt.Sprintf("%s contains %s (%d observations)", parent.Tag, symbol.Name, parent.Count)
			addEdge(parentEntry.ID, entry.ID, "contains", evidence)
			contextIDs = append(contextIDs, parentEntry.ID)
		}
		if len(symbol.ParentRefs) == 0 {
			for _, parentType := range symbol.CommonParentTypes {
				parentEntry, matched := catalog.lookup(parentType, "ui_element")
				if !matched {
					continue
				}
				evidence := fmt.Sprintf("%s commonly contains %s", parentType, symbol.Name)
				addEdge(parentEntry.ID, entry.ID, "contains", evidence)
				contextIDs = append(contextIDs, parentEntry.ID)
			}
		}
		for _, child := range symbol.ChildRefs {
			childEntry, matched := catalog.lookup(child.Tag, "ui_element")
			if !matched {
				continue
			}
			evidence := fmt.Sprintf("%s contains named child %s (%d observations)", symbol.Name, child.Tag, child.Count)
			addEdge(entry.ID, childEntry.ID, "contains", evidence)
			contextIDs = append(contextIDs, childEntry.ID)
		}
		for _, child := range symbol.StructuralChildRefs {
			childEntry, matched := catalog.lookup(child.Tag, "ui_element")
			if !matched {
				continue
			}
			evidence := fmt.Sprintf("%s contains structural child %s (%d observations)", symbol.Name, child.Tag, child.Count)
			addEdge(entry.ID, childEntry.ID, "contains", evidence)
			contextIDs = append(contextIDs, childEntry.ID)
		}

		handlerNames := append([]string{}, symbol.CommonHandlers...)
		for _, binding := range symbol.XMLEventBindings {
			handlerNames = append(handlerNames, binding.Event)
		}
		for _, handlerName := range graph.UniqueStrings(handlerNames) {
			if handlerName == "" {
				continue
			}
			if eventEntry, matched := catalog.lookup(handlerName, "xml_event", "event"); matched {
				evidence := fmt.Sprintf("%s declares XML event %s", symbol.Name, handlerName)
				addEdge(entry.ID, eventEntry.ID, "handled_by", evidence)
				contextIDs = append(contextIDs, eventEntry.ID)
			}
			if handlerEntry, matched := catalog.lookup(handlerName, "xml_event"); matched {
				evidence := fmt.Sprintf("%s declares XML handler %s", symbol.Name, handlerName)
				addEdge(entry.ID, handlerEntry.ID, "handled_by", evidence)
				contextIDs = append(contextIDs, handlerEntry.ID)
			}
		}
		recordCombination(contextIDs, symbol.Name+": XML element relationships")
	}

	for _, symbol := range append(append([]FunctionSymbol{}, corpus.GlobalFunctions...), corpus.WindowFunctions...) {
		entry, found := catalog.entries[functionSymbolID(symbol.Category, symbol.Name)]
		if !found {
			continue
		}
		for _, example := range symbol.Examples {
			lookupKey := callerKey(example.Addon, example.Caller)
			evidence := example.Addon + ": " + example.Caller + " -> " + example.Snippet
			contextIDs := []string{entry.ID}
			for _, link := range invokersByCaller[lookupKey] {
				if link.Type == "event" {
					addEdge(entry.ID, link.ID, "triggered_by", evidence)
				} else if link.Type == "xml_event" {
					addEdge(entry.ID, link.ID, "bound_from_xml", evidence)
				}
				contextIDs = append(contextIDs, link.ID)
			}
			for _, link := range uiLinksByCaller[lookupKey] {
				addEdge(entry.ID, link.ID, "updates_ui", evidence)
				contextIDs = append(contextIDs, link.ID)
			}
			callerDoc, found := sourceFunctionsByKey[lookupKey]
			if found {
				for _, call := range callerDoc.Calls {
					relatedEntry, matched := catalog.lookup(call.Name, "function")
					if !matched || relatedEntry.ID == entry.ID {
						continue
					}
					contextIDs = append(contextIDs, relatedEntry.ID)
				}
				for _, dataEntry := range extractDataLinks(callerDoc, catalog) {
					addEdge(entry.ID, dataEntry.ID, "reads_data", evidence)
					contextIDs = append(contextIDs, dataEntry.ID)
				}
				for _, stateEntry := range extractStateLinks(callerDoc.StateWrites, catalog) {
					addEdge(entry.ID, stateEntry.ID, "writes_state", evidence)
					contextIDs = append(contextIDs, stateEntry.ID)
				}
			}
			recordCombination(contextIDs, evidence)
		}
		if strings.Contains(strings.ToLower(symbol.Description), "window") || symbol.Category == "Window Function" {
			windowElement, matched := catalog.lookup("Window", "ui_element")
			if matched {
				addEdge(entry.ID, windowElement.ID, "updates_ui", symbol.Name+": description and category indicate window-facing behavior")
			}
		}
	}

	// Emit commonly_used_with edges directly into edgeAccumulators so that the
	// final graph edge carries the real relation-accumulator weight (>= minCommonlyUsedWithWeight)
	// and the actual co-occurrence evidence strings.  Using addEdge() here would always
	// produce weight=1 (one call = one increment), hiding the true co-occurrence count
	// from downstream enrichment and cap logic.
	for _, relation := range materializeCombinationRelations(relationAccumulators, catalog) {
		if relation.Weight < minCommonlyUsedWithWeight {
			continue
		}
		if len(relation.Participants) != 2 {
			continue
		}
		leftID := relation.Participants[0].ID
		rightID := relation.Participants[1].ID
		if leftID == "" || rightID == "" || leftID == rightID {
			continue
		}
		key := "commonly_used_with|" + leftID + "|" + rightID
		acc := &graphEdgeAccumulator{
			From:     leftID,
			To:       rightID,
			Type:     "commonly_used_with",
			Weight:   relation.Weight,
			Evidence: map[string]bool{},
		}
		for _, ev := range relation.Evidence {
			if strings.TrimSpace(ev) != "" {
				acc.Evidence[ev] = true
			}
		}
		edgeAccumulators[key] = acc
	}
	edges := materializeGraphEdges(edgeAccumulators)
	edges = capCommonlyUsedWithEdges(edges, maxCommonlyUsedWithPerNode)

	// Enrich edges with deep analysis findings (confidence scores, rationales, missing edge types)
	edges = enrichEdgesWithDeepAnalysis(edges, corpus, catalog)

	graphData := APIGraph{GeneratedAt: time.Now().UTC(), Nodes: materializeGraphNodes(catalog), Edges: edges}
	relations := buildRelationReport(catalog, edges, relationAccumulators, eventFlows)
	links := buildSemanticLinks(catalog, edges)
	return graphData, relations, links
}

func buildPatternPages(corpus Corpus, catalog symbolCatalog, relations RelationReport) []PatternPage {
	pages := []PatternPage{}
	appendPages := func(category string, docs []PatternDoc) {
		for _, doc := range docs {
			slug := graph.Slug(category + "_" + doc.Name)
			involved := extractPatternLinks(doc, catalog)
			pages = append(pages, PatternPage{
				Name:        doc.Name,
				Slug:        slug,
				Category:    category,
				Confidence:  doc.Confidence,
				Description: doc.Description,
				Involved:    involved,
				Flow:        inferPatternFlow(doc, involved, relations),
				ExampleCode: selectPatternExample(doc),
				Evidence:    firstStrings(doc.Evidence, 12),
				Path:        slashPath("patterns", slug+".md"),
			})
		}
	}
	appendPages("window", corpus.WindowPatterns)
	appendPages("events", corpus.EventPatterns)
	appendPages("slash_commands", corpus.SlashCommandPatterns)
	appendPages("conventions", corpus.Conventions)
	sort.Slice(pages, func(leftIndex, rightIndex int) bool {
		if pages[leftIndex].Category == pages[rightIndex].Category {
			return pages[leftIndex].Name < pages[rightIndex].Name
		}
		return pages[leftIndex].Category < pages[rightIndex].Category
	})
	return pages
}

func buildNavigationTree(corpus Corpus, catalog symbolCatalog, patternPages []PatternPage) NavigationNode {
	return NavigationNode{
		ID:    "war_api_ref",
		Label: "WAR API Reference",
		Type:  "root",
		Path:  "index.md",
		Children: []NavigationNode{
			makeCategoryNode("globals", "Globals", []NavigationNode{
				makeSectionNode("globals_functions", "Functions", "globals/functions/index.md", linksToNavigationNodes(linksForCategory(catalog, "Global Function"))),
				makeSectionNode("globals_tables", "Tables", "globals/tables/index.md", linksToNavigationNodes(linksForCategory(catalog, "Global Table"))),
				makeSectionNode("globals_constants", "Constants", "globals/constants/index.md", linksToNavigationNodes(linksForCategory(catalog, "Constant"))),
			}),
			makeCategoryNode("window_api", "Window API", []NavigationNode{
				makeSectionNode("window_api_functions", "Functions", "window_api/functions/index.md", linksToNavigationNodes(linksForCategory(catalog, "Window Function"))),
				makeSectionNode("window_api_lifecycle", "Lifecycle", "window_api/lifecycle.md", nil),
				makeSectionNode("window_api_patterns", "Patterns", "window_api/patterns.md", nil),
			}),
			makeCategoryNode("xml", "XML", []NavigationNode{
				makeSectionNode("xml_element_types", "Element Types", "xml/element_types/index.md", linksToNavigationNodes(linksForType(catalog, "ui_element"))),
				makeSectionNode("xml_handlers", "Handlers", "xml/handlers/index.md", linksToNavigationNodes(linksForType(catalog, "xml_event"))),
				makeSectionNode("xml_schema", "Schema", "xml/schema.md", nil),
			}),
			makeCategoryNode("events", "Events", []NavigationNode{
				makeSectionNode("events_game", "Game Events", "events/game_events/index.md", linksToNavigationNodes(linksForCategory(catalog, "Game Event"))),
				makeSectionNode("events_window", "Window Events", "events/window_events/index.md", linksToNavigationNodes(linksForCategory(catalog, "Window Event"))),
				makeSectionNode("events_patterns", "Patterns", "events/patterns.md", nil),
			}),
			makeCategoryNode("systemdata", "SystemData", []NavigationNode{
				makeSectionNode("systemdata_index", "Fields", "systemdata/index.md", linksToNavigationNodes(linksForCategory(catalog, "SystemData Field"))),
			}),
			makeCategoryNode("gamedata", "GameData", []NavigationNode{
				makeSectionNode("gamedata_index", "Fields", "gamedata/index.md", linksToNavigationNodes(linksForCategory(catalog, "GameData Field"))),
			}),
			makeCategoryNode("lifecycle", "Lifecycle", []NavigationNode{
				makeSectionNode("addon_lifecycle", "Addon Lifecycle", "lifecycle/addon_lifecycle.md", nil),
			}),
			makeCategoryNode("patterns", "Patterns", []NavigationNode{
				makeSectionNode("patterns_index", "Overview", "patterns/index.md", patternPagesToNavigationNodes(patternPages)),
			}),
		},
	}
}

func buildSidebar(corpus Corpus, catalog symbolCatalog, patternPages []PatternPage) []SidebarItem {
	root := buildNavigationTree(corpus, catalog, patternPages)
	items := []SidebarItem{}
	for index, child := range root.Children {
		items = append(items, navigationNodeToSidebarItem(child, index+1))
	}
	return items
}

func materializeGraphNodes(catalog symbolCatalog) []GraphNode {
	nodes := make([]GraphNode, 0, len(catalog.entries))
	for _, entry := range catalog.entries {
		nodes = append(nodes, GraphNode{
			ID:         entry.ID,
			Label:      entry.Label,
			Type:       entry.Type,
			Category:   entry.Category,
			Path:       entry.Path,
			Confidence: entry.Confidence,
			Score:      entry.Score,
			Summary:    entry.Summary,
		})
	}
	sort.Slice(nodes, func(leftIndex, rightIndex int) bool {
		if nodes[leftIndex].Type == nodes[rightIndex].Type {
			return nodes[leftIndex].Label < nodes[rightIndex].Label
		}
		return nodes[leftIndex].Type < nodes[rightIndex].Type
	})
	return nodes
}

func materializeGraphEdges(accumulators map[string]*graphEdgeAccumulator) []GraphEdge {
	edges := make([]GraphEdge, 0, len(accumulators))
	for _, accumulator := range accumulators {
		edges = append(edges, GraphEdge{
			From:     accumulator.From,
			To:       accumulator.To,
			Type:     accumulator.Type,
			Weight:   accumulator.Weight,
			Evidence: firstStrings(mapKeys(accumulator.Evidence), 6),
		})
	}
	sort.Slice(edges, func(leftIndex, rightIndex int) bool {
		if edges[leftIndex].Type == edges[rightIndex].Type {
			if edges[leftIndex].From == edges[rightIndex].From {
				return edges[leftIndex].To < edges[rightIndex].To
			}
			return edges[leftIndex].From < edges[rightIndex].From
		}
		return edges[leftIndex].Type < edges[rightIndex].Type
	})
	return edges
}

// capCommonlyUsedWithEdges limits the number of commonly_used_with edges per source
// node to maxPerNode (sorted by descending weight). This prevents high-degree hub
// nodes from dominating the graph with many low-quality association edges.
func capCommonlyUsedWithEdges(edges []GraphEdge, maxPerNode int) []GraphEdge {
	// Separate commonly_used_with from other edges.
	var cwEdges []GraphEdge
	var otherEdges []GraphEdge
	for _, e := range edges {
		if e.Type == "commonly_used_with" {
			cwEdges = append(cwEdges, e)
		} else {
			otherEdges = append(otherEdges, e)
		}
	}
	if len(cwEdges) == 0 {
		return otherEdges
	}
	// Sort by descending weight so each node keeps its strongest associations.
	sort.Slice(cwEdges, func(i, j int) bool {
		return cwEdges[i].Weight > cwEdges[j].Weight
	})
	// Count how many edges have been kept per source node.
	keptPerNode := make(map[string]int)
	for _, e := range cwEdges {
		if keptPerNode[e.From] < maxPerNode {
			otherEdges = append(otherEdges, e)
			keptPerNode[e.From]++
		}
	}
	return otherEdges
}

// enrichEdgesWithDeepAnalysis applies deep_analysis inference to enrich edges with confidence scores
func enrichEdgesWithDeepAnalysis(edges []GraphEdge, corpus Corpus, catalog symbolCatalog) []GraphEdge {
	if len(corpus.Source.Functions) == 0 {
		return edges
	}

	// Create enricher and analyze function sources
	enricher := deep_analysis.NewEdgeEnricher()
	for _, fnDoc := range corpus.Source.Functions {
		if fnDoc.Source == "" {
			continue
		}
		// Determine if this is an init or update function based on name patterns
		isInit := strings.Contains(fnDoc.Name, "Init") || strings.Contains(fnDoc.Name, "Load")
		isUpdate := strings.Contains(fnDoc.Name, "Update") || strings.Contains(fnDoc.Name, "Refresh") || strings.Contains(fnDoc.Name, "OnUpdate")

		// Analyze function source
		fullPath := fnDoc.Addon + "." + fnDoc.Name
		enricher.EdgeInference.AnalyzeFunctionSource(fullPath, fnDoc.Source, isInit, isUpdate, false, "")
		enricher.ReturnInference.AnalyzeReturns(fullPath, fnDoc.Source)
		enricher.ReturnInference.AnalyzeCallSites(fullPath, fnDoc.Source)
		enricher.ReturnInference.AnalyzeFieldAccess(fullPath, fnDoc.Source)

		// Analyze arguments
		paramNames := enricher.ArgumentInference.AnalyzeParameters(fullPath, fnDoc.Source)
		enricher.ArgumentInference.AnalyzeParameterUsage(fullPath, fnDoc.Source, paramNames)
		enricher.ArgumentInference.AnalyzeCallSites(fullPath, fnDoc.Source, paramNames)
	}

	// Score existing edges with confidence based on evidence
	enrichedEdges := make([]GraphEdge, 0, len(edges))
	for _, edge := range edges {
		enrichedEdge := edge

		// Calculate confidence based on edge type and evidence count
		confidence := calculateEdgeConfidence(edge.Type, edge.Weight, len(edge.Evidence))
		enrichedEdge.ConfidenceScore = confidence
		enrichedEdge.EvidenceCount = len(edge.Evidence)
		enrichedEdge.EvidenceSources = edge.Evidence // Map evidence strings to sources

		// Generate rationale
		enrichedEdge.Rationale = generateEdgeRationale(edge.Type, edge.From, edge.To, edge.Weight)

		enrichedEdges = append(enrichedEdges, enrichedEdge)
	}

	// Build missing edges from deep analysis and score them
	missingEdges := buildMissingEdgesFromAnalysis(enricher, catalog)
	filterThreshold := 50 // Include edges with 50%+ confidence
	for _, mEdge := range missingEdges {
		if mEdge.ConfidenceScore >= filterThreshold {
			enrichedEdges = append(enrichedEdges, GraphEdge{
				From:            mEdge.From,
				To:              mEdge.To,
				Type:            mEdge.Type,
				Weight:          1, // Missing edges have minimal weight
				ConfidenceScore: mEdge.ConfidenceScore,
				EvidenceCount:   mEdge.EvidenceCount,
				EvidenceSources: mEdge.EvidenceSources,
				Rationale:       mEdge.Rationale,
				Evidence:        mEdge.EvidenceSources,
			})
		}
	}

	// Re-sort edges after enrichment
	sort.Slice(enrichedEdges, func(leftIndex, rightIndex int) bool {
		if enrichedEdges[leftIndex].Type == enrichedEdges[rightIndex].Type {
			if enrichedEdges[leftIndex].From == enrichedEdges[rightIndex].From {
				return enrichedEdges[leftIndex].To < enrichedEdges[rightIndex].To
			}
			return enrichedEdges[leftIndex].From < enrichedEdges[rightIndex].From
		}
		return enrichedEdges[leftIndex].Type < enrichedEdges[rightIndex].Type
	})

	return enrichedEdges
}

// calculateEdgeConfidence returns confidence score for an edge based on type and evidence
func calculateEdgeConfidence(edgeType string, weight int, evidenceCount int) int {
	baseScore := 50

	switch edgeType {
	case "calls":
		baseScore = 85 + (weight-1)*2
	case "handled_by", "triggered_by":
		baseScore = 80 + (weight-1)*2
	case "contains":
		baseScore = 78 + (weight-1)*2
	case "defined_in":
		baseScore = 95
	case "reads_data", "writes_state":
		baseScore = 70 + (weight-1)*2
	case "updates_ui", "bound_from_xml":
		baseScore = 72 + (weight-1)*2
	case "commonly_used_with":
		baseScore = 55 + (weight-1)*1
	}

	// Evidence multiplier
	if evidenceCount > 5 {
		baseScore += 10
	}
	if evidenceCount > 10 {
		baseScore += 5
	}

	if baseScore > 95 {
		baseScore = 95
	}
	return baseScore
}

// generateEdgeRationale creates human-readable explanation for why an edge exists
func generateEdgeRationale(edgeType, from, to string, weight int) string {
	switch edgeType {
	case "calls":
		return fmt.Sprintf("%s calls %s (%d observations)", from, to, weight)
	case "handled_by":
		return fmt.Sprintf("%s is handled by %s", from, to)
	case "triggered_by":
		return fmt.Sprintf("%s triggers %s", from, to)
	case "defined_in":
		return fmt.Sprintf("%s is defined in %s", from, to)
	case "contains":
		return fmt.Sprintf("%s contains %s", from, to)
	case "reads_data":
		return fmt.Sprintf("%s reads data from %s", from, to)
	case "writes_state":
		return fmt.Sprintf("%s updates state in %s", from, to)
	case "updates_ui":
		return fmt.Sprintf("%s updates UI element %s", from, to)
	case "bound_from_xml":
		return fmt.Sprintf("%s is bound from XML handler %s", from, to)
	case "commonly_used_with":
		return fmt.Sprintf("%s and %s are commonly used together", from, to)
	default:
		return fmt.Sprintf("%s relates to %s via %s", from, to, edgeType)
	}
}

// buildMissingEdgesFromAnalysis discovers new edges using deep_analysis.
// Only edges whose target resolves to a real catalog entry are emitted; phantom
// fallback IDs (e.g. "systemdata_collective", "ui_visibility") are dropped entirely
// to avoid unresolvable graph nodes.
func buildMissingEdgesFromAnalysis(enricher *deep_analysis.EdgeEnricher, catalog symbolCatalog) []deep_analysis.EnrichedEdge {
	enrichedEdges := []deep_analysis.EnrichedEdge{}

	// Discover data access edges
	for functionPath, fa := range enricher.EdgeInference.AllFunctions {
		// reads_systemdata — only emit when catalog resolves "SystemData"
		if enricher.EdgeInference.InferReadsSystemData(functionPath) {
			systemDataEntry, ok := catalog.lookup("SystemData", "data_structure")
			if ok {
				if fromID, fromOK := findFunctionID(functionPath, catalog); fromOK {
					enrichedEdges = append(enrichedEdges, deep_analysis.EnrichedEdge{
						From:            fromID,
						To:              systemDataEntry.ID,
						Type:            "reads_systemdata",
						ConfidenceScore: 72,
						EvidenceCount:   len(fa.AccessedGlobals),
						EvidenceSources: []string{"static_analysis:global_access"},
						Rationale:       fmt.Sprintf("%s reads SystemData (%d access points)", functionPath, len(fa.AccessedGlobals)),
					})
				}
			}
		}

		// reads_gamedata — only emit when catalog resolves "GameData"
		if enricher.EdgeInference.InferReadsGameData(functionPath) {
			gameDataEntry, ok := catalog.lookup("GameData", "data_structure")
			if ok {
				if fromID, fromOK := findFunctionID(functionPath, catalog); fromOK {
					enrichedEdges = append(enrichedEdges, deep_analysis.EnrichedEdge{
						From:            fromID,
						To:              gameDataEntry.ID,
						Type:            "reads_gamedata",
						ConfidenceScore: 72,
						EvidenceCount:   len(fa.AccessedGlobals),
						EvidenceSources: []string{"static_analysis:global_access"},
						Rationale:       fmt.Sprintf("%s reads GameData (%d access points)", functionPath, len(fa.AccessedGlobals)),
					})
				}
			}
		}

		// updates_ui — only emit when catalog resolves "Window"
		if enricher.EdgeInference.InferUpdatesUI(functionPath) {
			windowEntry, ok := catalog.lookup("Window", "ui_element")
			if ok {
				if fromID, fromOK := findFunctionID(functionPath, catalog); fromOK {
					enrichedEdges = append(enrichedEdges, deep_analysis.EnrichedEdge{
						From:            fromID,
						To:              windowEntry.ID,
						Type:            "updates_ui",
						ConfidenceScore: 70,
						EvidenceCount:   len(fa.CallsSites),
						EvidenceSources: []string{"static_analysis:ui_patterns"},
						Rationale:       fmt.Sprintf("%s performs UI updates", functionPath),
					})
				}
			}
		}

		// toggles_visibility and updates_layout are dropped: they previously pointed to
		// "ui_visibility" / "ui_layout" — placeholder strings that are not catalog entries.
		// Re-enable if a canonical ui_element entry for these concepts is added to the catalog.

		// initializes and refreshes are dropped: they previously pointed to
		// "addon_init_phase" / "addon_update_phase" — phantom IDs with no catalog backing.
		// Re-enable if lifecycle phase nodes are added as proper catalog entries.
	}

	return enrichedEdges
}

// findFunctionID looks up a function by its full path in the catalog
func findFunctionID(functionPath string, catalog symbolCatalog) (string, bool) {
	parts := strings.Split(functionPath, ".")
	if len(parts) > 0 {
		funcName := parts[len(parts)-1]
		if entries, ok := catalog.byName[normalizeLookupKey(funcName)]; ok && len(entries) > 0 {
			return entries[0], true
		}
	}
	return "", false
}

func materializeCombinationRelations(accumulators map[string]*relationAccumulator, catalog symbolCatalog) []RelationChain {
	relations := []RelationChain{}
	for key, accumulator := range accumulators {
		leftEntry, leftFound := catalog.entries[accumulator.LeftID]
		rightEntry, rightFound := catalog.entries[accumulator.RightID]
		if !leftFound || !rightFound {
			continue
		}
		relations = append(relations, RelationChain{
			ID:           "combination:" + graph.Slug(key),
			Type:         "common_combination",
			Title:        leftEntry.Label + " + " + rightEntry.Label,
			Description:  "Observed together in the same inferred usage contexts.",
			Participants: []DocLink{leftEntry, rightEntry},
			Flow:         []string{leftEntry.Label + " <-> " + rightEntry.Label},
			Evidence:     firstStrings(mapKeys(accumulator.Evidence), 6),
			Weight:       accumulator.Weight,
		})
	}
	sort.Slice(relations, func(leftIndex, rightIndex int) bool {
		if relations[leftIndex].Weight == relations[rightIndex].Weight {
			return relations[leftIndex].Title < relations[rightIndex].Title
		}
		return relations[leftIndex].Weight > relations[rightIndex].Weight
	})
	return relations
}

func buildRelationReport(catalog symbolCatalog, edges []GraphEdge, combinations map[string]*relationAccumulator, eventFlows map[string]*eventFlowAccumulator) RelationReport {
	report := RelationReport{GeneratedAt: time.Now().UTC()}
	for _, edge := range edges {
		if edge.Type != "calls" || edge.Weight < 2 {
			continue
		}
		fromEntry, fromFound := catalog.entries[edge.From]
		toEntry, toFound := catalog.entries[edge.To]
		if !fromFound || !toFound {
			continue
		}
		report.FrequentCallChains = append(report.FrequentCallChains, RelationChain{
			ID:           "call_chain:" + graph.Slug(edge.From+"_"+edge.To),
			Type:         "call_chain",
			Title:        fromEntry.Label + " -> " + toEntry.Label,
			Description:  "Repeated call relationship extracted from canonical symbol usage contexts.",
			Participants: []DocLink{fromEntry, toEntry},
			Flow:         []string{fromEntry.Label + " -> " + toEntry.Label},
			Evidence:     edge.Evidence,
			Weight:       edge.Weight,
		})
	}
	report.CommonCombinations = firstRelationChains(materializeCombinationRelations(combinations, catalog), 60)
	for _, flow := range eventFlows {
		eventEntry, found := catalog.entries[flow.EventID]
		if !found {
			continue
		}
		handlers := linksForIDs(catalog, mapKeys(flow.HandlerIDs))
		uiElements := linksForIDs(catalog, mapKeys(flow.UIIDs))
		participants := []DocLink{eventEntry}
		participants = append(participants, handlers...)
		participants = append(participants, uiElements...)
		flowLines := []string{eventEntry.Label}
		if len(handlers) > 0 {
			flowLines = append(flowLines, "  -> handlers: "+joinLabels(handlers))
		}
		if len(uiElements) > 0 {
			flowLines = append(flowLines, "  -> ui: "+joinLabels(uiElements))
		}
		report.EventHandlerUIFlows = append(report.EventHandlerUIFlows, RelationChain{
			ID:           "event_flow:" + graph.Slug(flow.EventID),
			Type:         "event_handler_ui_flow",
			Title:        eventEntry.Label,
			Description:  "Observed event to handler to UI-element flow reconstructed from bindings, handlers, and examples.",
			Participants: participants,
			Flow:         flowLines,
			Evidence:     firstStrings(mapKeys(flow.Evidence), 8),
			Weight:       flow.Weight,
		})
	}
	sort.Slice(report.FrequentCallChains, func(leftIndex, rightIndex int) bool {
		if report.FrequentCallChains[leftIndex].Weight == report.FrequentCallChains[rightIndex].Weight {
			return report.FrequentCallChains[leftIndex].Title < report.FrequentCallChains[rightIndex].Title
		}
		return report.FrequentCallChains[leftIndex].Weight > report.FrequentCallChains[rightIndex].Weight
	})
	sort.Slice(report.EventHandlerUIFlows, func(leftIndex, rightIndex int) bool {
		if report.EventHandlerUIFlows[leftIndex].Weight == report.EventHandlerUIFlows[rightIndex].Weight {
			return report.EventHandlerUIFlows[leftIndex].Title < report.EventHandlerUIFlows[rightIndex].Title
		}
		return report.EventHandlerUIFlows[leftIndex].Weight > report.EventHandlerUIFlows[rightIndex].Weight
	})
	report.FrequentCallChains = firstRelationChains(report.FrequentCallChains, 60)
	report.EventHandlerUIFlows = firstRelationChains(report.EventHandlerUIFlows, 60)
	return report
}

func buildSemanticLinks(catalog symbolCatalog, edges []GraphEdge) map[string]SemanticLinks {
	links := map[string]SemanticLinks{}
	appendLink := func(symbolID string, kind string, link DocLink) {
		current := links[symbolID]
		switch kind {
		case "related":
			current.RelatedAPIs = appendUniqueLink(current.RelatedAPIs, link)
		case "used_with":
			current.UsedWith = appendUniqueLink(current.UsedWith, link)
		case "triggered_by":
			current.TriggeredBy = appendUniqueLink(current.TriggeredBy, link)
		case "affects":
			current.Affects = appendUniqueLink(current.Affects, link)
		}
		links[symbolID] = current
	}
	for _, edge := range edges {
		fromEntry, fromFound := catalog.entries[edge.From]
		toEntry, toFound := catalog.entries[edge.To]
		if !fromFound || !toFound {
			continue
		}
		switch edge.Type {
		case "calls":
			appendLink(edge.From, "related", toEntry)
			appendLink(edge.To, "related", fromEntry)
		case "defined_in", "contains":
			appendLink(edge.From, "related", toEntry)
			appendLink(edge.To, "related", fromEntry)
		case "commonly_used_with":
			appendLink(edge.From, "used_with", toEntry)
			appendLink(edge.To, "used_with", fromEntry)
		case "handled_by", "triggered_by":
			appendLink(edge.From, "triggered_by", toEntry)
			appendLink(edge.To, "related", fromEntry)
		case "bound_from_xml":
			appendLink(edge.From, "triggered_by", toEntry)
			appendLink(edge.To, "related", fromEntry)
		case "reads_data", "writes_state", "updates_ui":
			appendLink(edge.From, "affects", toEntry)
			appendLink(edge.To, "related", fromEntry)
		}
	}
	for symbolID, value := range links {
		value.RelatedAPIs = sortDocLinks(value.RelatedAPIs)
		value.UsedWith = sortDocLinks(value.UsedWith)
		value.TriggeredBy = sortDocLinks(value.TriggeredBy)
		value.Affects = sortDocLinks(value.Affects)
		links[symbolID] = value
	}
	return links
}

func extractDataLinks(doc FunctionDoc, catalog symbolCatalog) []DocLink {
	textParts := []string{doc.Name, doc.Module}
	for _, call := range doc.Calls {
		textParts = append(textParts, call.Name, call.Raw)
		textParts = append(textParts, call.Arguments...)
	}
	for _, registration := range doc.EventRegistrations {
		textParts = append(textParts, registration.Event, registration.Handler)
	}
	textParts = append(textParts, doc.StateWrites...)
	result := []DocLink{}
	seen := map[string]bool{}
	for _, part := range textParts {
		for _, token := range namespaceTokenPattern.FindAllString(part, -1) {
			entry, found := catalog.lookup(token, "data_structure")
			if !found || seen[entry.ID] {
				continue
			}
			seen[entry.ID] = true
			result = append(result, entry)
		}
	}
	return sortDocLinks(result)
}

func extractStateLinks(stateWrites []string, catalog symbolCatalog) []DocLink {
	result := []DocLink{}
	seen := map[string]bool{}
	for _, write := range stateWrites {
		entry, found := catalog.lookup(write, "data_structure")
		if !found {
			root := graph.RootSegment(write)
			entry, found = catalog.lookup(root, "data_structure")
		}
		if !found || seen[entry.ID] {
			continue
		}
		seen[entry.ID] = true
		result = append(result, entry)
	}
	return sortDocLinks(result)
}

func extractPatternLinks(doc PatternDoc, catalog symbolCatalog) []DocLink {
	result := []DocLink{}
	seen := map[string]bool{}
	texts := []string{doc.Name, doc.Description}
	texts = append(texts, doc.Evidence...)
	for _, entry := range sortDocLinks(catalogEntries(catalog)) {
		for _, text := range texts {
			if !strings.Contains(text, entry.Label) {
				continue
			}
			if !seen[entry.ID] {
				seen[entry.ID] = true
				result = append(result, entry)
			}
			break
		}
	}
	return sortDocLinks(result)
}

func inferPatternFlow(doc PatternDoc, involved []DocLink, relations RelationReport) []string {
	if best := bestMatchingRelation(involved, relations); len(best.Flow) > 0 {
		return best.Flow
	}
	if len(involved) >= 2 {
		labels := []string{}
		for _, link := range involved {
			labels = append(labels, link.Label)
		}
		return []string{strings.Join(labels, " -> ")}
	}
	if len(doc.Evidence) > 0 {
		return []string{doc.Evidence[0]}
	}
	return []string{"No stable multi-symbol flow was inferable from the current corpus."}
}

func selectPatternExample(doc PatternDoc) string {
	for _, evidence := range doc.Evidence {
		trimmed := strings.TrimSpace(evidence)
		if trimmed == "" {
			continue
		}
		return trimmed
	}
	return ""
}

func normalizeLookupKey(value string) string {
	trimmed := strings.TrimSpace(value)
	trimmed = strings.Trim(trimmed, "\"'")
	return strings.ToLower(trimmed)
}

func callerKey(addon string, caller string) string {
	return normalizeLookupKey(addon) + "|" + normalizeLookupKey(caller)
}

func frameLookupKey(addon string, frame string) string {
	return normalizeLookupKey(addon) + "|" + normalizeLookupKey(frame)
}

func (catalog symbolCatalog) lookup(name string, preferredTypes ...string) (DocLink, bool) {
	ids := catalog.byName[normalizeLookupKey(name)]
	if len(ids) == 0 {
		return DocLink{}, false
	}
	preferred := map[string]int{}
	for index, preferredType := range preferredTypes {
		preferred[preferredType] = index
	}
	best := DocLink{}
	bestFound := false
	bestRank := len(preferredTypes) + 1
	for _, identifier := range ids {
		entry := catalog.entries[identifier]
		rank := len(preferredTypes)
		if value, ok := preferred[entry.Type]; ok {
			rank = value
		}
		if !bestFound || rank < bestRank || (rank == bestRank && entry.Score > best.Score) || (rank == bestRank && entry.Score == best.Score && entry.Label < best.Label) {
			best = entry
			bestFound = true
			bestRank = rank
		}
		if bestFound && bestRank == 0 {
			continue
		}
	}
	return best, bestFound
}

func appendUniqueLink(target []DocLink, link DocLink) []DocLink {
	for _, existing := range target {
		if existing.ID == link.ID {
			return target
		}
	}
	return append(target, link)
}

func sortDocLinks(values []DocLink) []DocLink {
	sort.Slice(values, func(leftIndex, rightIndex int) bool {
		if values[leftIndex].Score == values[rightIndex].Score {
			return values[leftIndex].Label < values[rightIndex].Label
		}
		return values[leftIndex].Score > values[rightIndex].Score
	})
	return values
}

func linksForCategory(catalog symbolCatalog, category string) []DocLink {
	result := []DocLink{}
	for _, entry := range catalog.entries {
		if entry.Category == category {
			result = append(result, entry)
		}
	}
	return sortDocLinks(result)
}

func linksForType(catalog symbolCatalog, linkType string) []DocLink {
	result := []DocLink{}
	for _, entry := range catalog.entries {
		if entry.Type == linkType {
			result = append(result, entry)
		}
	}
	return sortDocLinks(result)
}

func linksForIDs(catalog symbolCatalog, ids []string) []DocLink {
	result := []DocLink{}
	for _, identifier := range ids {
		entry, found := catalog.entries[identifier]
		if found {
			result = append(result, entry)
		}
	}
	return sortDocLinks(result)
}

func catalogEntries(catalog symbolCatalog) []DocLink {
	entries := make([]DocLink, 0, len(catalog.entries))
	for _, entry := range catalog.entries {
		entries = append(entries, entry)
	}
	return entries
}

func bestMatchingRelation(involved []DocLink, relations RelationReport) RelationChain {
	best := RelationChain{}
	bestOverlap := 0
	relationSets := [][]RelationChain{relations.FrequentCallChains, relations.CommonCombinations, relations.EventHandlerUIFlows}
	for _, group := range relationSets {
		for _, relation := range group {
			overlap := 0
			for _, participant := range relation.Participants {
				for _, link := range involved {
					if participant.ID == link.ID {
						overlap++
						break
					}
				}
			}
			if overlap > bestOverlap || (overlap == bestOverlap && relation.Weight > best.Weight) {
				best = relation
				bestOverlap = overlap
			}
		}
	}
	return best
}

func makeCategoryNode(identifier string, label string, children []NavigationNode) NavigationNode {
	return NavigationNode{ID: identifier, Label: label, Type: "category", Children: children}
}

func makeSectionNode(identifier string, label string, path string, children []NavigationNode) NavigationNode {
	return NavigationNode{ID: identifier, Label: label, Type: "section", Path: path, Children: children}
}

func linksToNavigationNodes(links []DocLink) []NavigationNode {
	result := make([]NavigationNode, 0, len(links))
	for _, link := range links {
		result = append(result, NavigationNode{ID: link.ID, Label: link.Label, Type: link.Type, Path: link.Path, Importance: link.Score})
	}
	return result
}

func patternPagesToNavigationNodes(patternPages []PatternPage) []NavigationNode {
	result := make([]NavigationNode, 0, len(patternPages))
	for _, page := range patternPages {
		result = append(result, NavigationNode{ID: "pattern:" + page.Slug, Label: page.Name, Type: "pattern", Path: page.Path, Importance: patternImportance(page.Confidence)})
	}
	return result
}

func navigationNodeToSidebarItem(node NavigationNode, order int) SidebarItem {
	item := SidebarItem{Label: node.Label, Type: node.Type, Path: node.Path, Order: order}
	for index, child := range node.Children {
		item.Items = append(item.Items, navigationNodeToSidebarItem(child, index+1))
	}
	return item
}

func patternImportance(confidence Confidence) int {
	switch confidence {
	case ConfidenceHigh:
		return 100
	case ConfidenceMedium:
		return 60
	default:
		return 20
	}
}

func firstRelationChains(values []RelationChain, limit int) []RelationChain {
	if len(values) <= limit {
		return values
	}
	return values[:limit]
}

func joinLabels(links []DocLink) string {
	labels := []string{}
	for _, link := range links {
		labels = append(labels, link.Label)
	}
	return strings.Join(labels, ", ")
}

func slashPath(parts ...string) string {
	return strings.ReplaceAll(strings.Join(parts, "/"), "\\", "/")
}

func globalFunctionID(name string) string { return "function:global:" + name }
func windowFunctionID(name string) string { return "function:window:" + name }
func tableID(name string) string          { return "data:table:" + name }
func constantID(name string) string       { return "data:constant:" + name }
func elementTypeID(name string) string    { return "ui:" + name }
func xmlHandlerID(name string) string     { return "xml_event:" + name }
func gameEventID(name string) string      { return "event:game:" + name }
func windowEventID(name string) string    { return "event:window:" + name }
func systemFieldID(name string) string    { return "data:system:" + name }
func gameFieldID(name string) string      { return "data:game:" + name }

func functionSymbolID(category string, name string) string {
	if category == "Window Function" {
		return windowFunctionID(name)
	}
	return globalFunctionID(name)
}

func eventSymbolID(category string, name string) string {
	if category == "Window Event" {
		return windowEventID(name)
	}
	return gameEventID(name)
}

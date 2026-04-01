package normalizer

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"

	"roraddons/tools/api_doc_gen/graph"
)

func Normalize(corpus graph.Corpus) graph.NormalizedData {
	data := graph.NormalizedData{
		GeneratedAt: time.Now().UTC(),
	}

	concepts := map[string]*graph.ConceptGroup{}
	events := map[string]map[string]bool{}
	modules := map[string]map[string]bool{}
	stateOwners := map[string]map[string]bool{}
	frameParents := map[string]map[string]bool{}
	patternEvidence := map[string][]string{}
	patternCategories := map[string]string{}

	namingRules := map[string][]string{}

	for _, addon := range corpus.Addons {
		for _, function := range addon.Functions {
			conceptName := canonicalConcept(function.Name)
			if conceptName != "" {
				group, ok := concepts[conceptName]
				if !ok {
					group = &graph.ConceptGroup{Canonical: conceptName, Kind: "function", Members: []string{}, Addons: []string{}}
					concepts[conceptName] = group
				}
				group.Members = append(group.Members, function.Name)
				group.Addons = append(group.Addons, addon.Name)
			}

			last := graph.LastSegment(function.Name)
			if strings.HasPrefix(last, "On") {
				namingRules["Handler functions use the On<Event> prefix"] = append(namingRules["Handler functions use the On<Event> prefix"], function.Name)
			}
			if strings.Contains(function.Name, ".") {
				namingRules["Lua logic is namespaced through module-qualified function names"] = append(namingRules["Lua logic is namespaced through module-qualified function names"], function.Name)
			}
			for _, call := range function.Calls {
				if call.Name == "LayoutEditor.RegisterWindow" {
					patternCategories["LayoutEditor movable overlays"] = "ui"
					patternEvidence["LayoutEditor movable overlays"] = append(patternEvidence["LayoutEditor movable overlays"], addon.Name+": "+function.Name)
				}
				if call.Name == "CreateWindowFromTemplate" {
					patternCategories["XML template instantiation"] = "ui"
					patternEvidence["XML template instantiation"] = append(patternEvidence["XML template instantiation"], addon.Name+": "+function.Name)
				}
				if call.Name == "RegisterEventHandler" || call.Name == "WindowRegisterEventHandler" || call.Name == "WindowRegisterCoreEventHandler" {
					patternCategories["Event-driven controllers"] = "events"
					patternEvidence["Event-driven controllers"] = append(patternEvidence["Event-driven controllers"], addon.Name+": "+function.Name)
				}
			}
		}

		for _, event := range addon.Events {
			addToSetMap(events, event.Event, addon.Name)
			if strings.Contains(event.Event, "LOADING_END") || strings.Contains(event.Event, "RELOAD_INTERFACE") || strings.Contains(event.Event, "PLAYER_CAREER_LINE_UPDATED") {
				patternCategories["Deferred post-load initialization"] = "architecture"
				patternEvidence["Deferred post-load initialization"] = append(patternEvidence["Deferred post-load initialization"], addon.Name+": "+event.Event+" -> "+event.Handler)
			}
		}

		for _, module := range addon.Modules {
			addToSetMap(modules, module.Name, addon.Name)
			if strings.HasSuffix(module.Name, ".Settings") || strings.HasSuffix(module.Name, "Settings") {
				namingRules["Settings state is grouped under Settings-named modules or globals"] = append(namingRules["Settings state is grouped under Settings-named modules or globals"], module.Name)
			}
		}

		for _, variable := range addon.State {
			if variable.Owner != "" {
				addToSetMap(stateOwners, variable.Owner, addon.Name)
			}
			if variable.Saved {
				patternCategories["Saved-variable settings roots"] = "state"
				patternEvidence["Saved-variable settings roots"] = append(patternEvidence["Saved-variable settings roots"], addon.Name+": "+variable.Name)
			}
		}

		for _, frame := range addon.Frames {
			if frame.Parent != "" {
				addToSetMap(frameParents, frame.Parent, addon.Name)
			}
			if strings.HasSuffix(frame.Name, "Window") {
				namingRules["Top-level XML frames commonly use the Window suffix"] = append(namingRules["Top-level XML frames commonly use the Window suffix"], frame.Name)
			}
			if frame.Template || strings.Contains(frame.Name, "Template") {
				namingRules["Reusable XML elements use the Template suffix"] = append(namingRules["Reusable XML elements use the Template suffix"], frame.Name)
				patternCategories["XML template instantiation"] = "ui"
				patternEvidence["XML template instantiation"] = append(patternEvidence["XML template instantiation"], addon.Name+": "+frame.Name)
			}
			if frame.Inherits != "" {
				patternCategories["Inheritance-based XML composition"] = "ui"
				patternEvidence["Inheritance-based XML composition"] = append(patternEvidence["Inheritance-based XML composition"], addon.Name+": "+frame.Name+" inherits "+frame.Inherits)
			}
		}

		if len(addon.Manifest.InitializeCalls) > 0 || len(addon.Manifest.UpdateCalls) > 0 || len(addon.Manifest.ShutdownCalls) > 0 {
			patternCategories["Manifest lifecycle bootstrap"] = "architecture"
			patternEvidence["Manifest lifecycle bootstrap"] = append(patternEvidence["Manifest lifecycle bootstrap"], addon.Name+": "+strings.Join(append(append([]string{}, addon.Manifest.InitializeCalls...), addon.Manifest.UpdateCalls...), ", "))
		}

		data.ExecutionFlows = append(data.ExecutionFlows, buildExecutionFlow(addon))
	}

	for _, group := range concepts {
		group.Members = graph.UniqueStrings(group.Members)
		group.Addons = graph.UniqueStrings(group.Addons)
		group.Summary = fmt.Sprintf("%d implementations across %d addons", len(group.Members), len(group.Addons))
		if len(group.Members) == 1 {
			group.Summary = "Single explicit implementation"
		}
		data.ConceptGroups = append(data.ConceptGroups, *group)
	}
	sort.Slice(data.ConceptGroups, func(i, j int) bool {
		return data.ConceptGroups[i].Canonical < data.ConceptGroups[j].Canonical
	})

	for name, evidence := range patternEvidence {
		data.Patterns = append(data.Patterns, graph.Pattern{
			Name:     name,
			Category: patternCategories[name],
			Summary:  buildPatternSummary(name, evidence),
			Evidence: graph.UniqueStrings(evidence),
		})
	}
	sort.Slice(data.Patterns, func(i, j int) bool {
		if data.Patterns[i].Category == data.Patterns[j].Category {
			return data.Patterns[i].Name < data.Patterns[j].Name
		}
		return data.Patterns[i].Category < data.Patterns[j].Category
	})

	for rule, evidence := range namingRules {
		data.NamingConventions = append(data.NamingConventions, graph.NamingConvention{
			Rule:     rule,
			Evidence: firstN(graph.UniqueStrings(evidence), 8),
		})
	}
	sort.Slice(data.NamingConventions, func(i, j int) bool {
		return data.NamingConventions[i].Rule < data.NamingConventions[j].Rule
	})

	data.CommonEvents = topSharedKeys(events)
	data.CommonModules = topSharedKeys(modules)
	data.CommonStateOwners = topSharedKeys(stateOwners)
	data.CommonFrameParents = topSharedKeys(frameParents)

	sort.Slice(data.ExecutionFlows, func(i, j int) bool {
		return data.ExecutionFlows[i].Addon < data.ExecutionFlows[j].Addon
	})
	return data
}

func buildExecutionFlow(addon graph.AddonModel) graph.ExecutionFlow {
	steps := []graph.ExecutionStep{}
	if len(addon.Manifest.Files) > 0 {
		steps = append(steps, graph.ExecutionStep{
			Phase:    "manifest",
			Detail:   fmt.Sprintf("Loads %d files in manifest order", len(addon.Manifest.Files)),
			Evidence: firstN(addon.Manifest.Files, 8),
		})
	}
	if len(addon.Manifest.CreateWindows) > 0 || len(addon.Manifest.InitializeCalls) > 0 {
		evidence := append([]string{}, addon.Manifest.CreateWindows...)
		evidence = append(evidence, addon.Manifest.InitializeCalls...)
		steps = append(steps, graph.ExecutionStep{
			Phase:    "initialize",
			Detail:   fmt.Sprintf("Creates %d windows and calls %d initialize hooks", len(addon.Manifest.CreateWindows), len(addon.Manifest.InitializeCalls)),
			Evidence: firstN(graph.UniqueStrings(evidence), 10),
		})
	}
	if len(addon.Frames) > 0 || len(addon.Handlers) > 0 {
		steps = append(steps, graph.ExecutionStep{
			Phase:    "xml",
			Detail:   fmt.Sprintf("Defines %d XML frames and %d bound handlers", len(addon.Frames), len(addon.Handlers)),
			Evidence: firstN(frameNames(addon.Frames), 8),
		})
	}
	if len(addon.Events) > 0 {
		steps = append(steps, graph.ExecutionStep{
			Phase:    "runtime",
			Detail:   fmt.Sprintf("Registers %d runtime events", len(addon.Events)),
			Evidence: firstN(eventNames(addon.Events), 8),
		})
	}
	if len(addon.Manifest.UpdateCalls) > 0 {
		steps = append(steps, graph.ExecutionStep{
			Phase:    "update",
			Detail:   fmt.Sprintf("Runs %d update hooks from the manifest", len(addon.Manifest.UpdateCalls)),
			Evidence: addon.Manifest.UpdateCalls,
		})
	}
	if len(addon.Manifest.ShutdownCalls) > 0 {
		steps = append(steps, graph.ExecutionStep{
			Phase:    "shutdown",
			Detail:   fmt.Sprintf("Executes %d shutdown hooks", len(addon.Manifest.ShutdownCalls)),
			Evidence: addon.Manifest.ShutdownCalls,
		})
	}
	return graph.ExecutionFlow{Addon: addon.Name, Steps: steps}
}

func canonicalConcept(name string) string {
	segment := graph.LastSegment(name)
	tokens := splitIdentifier(segment)
	if len(tokens) == 0 {
		return ""
	}
	joined := strings.Join(tokens, " ")
	switch {
	case strings.Contains(joined, "refresh") || strings.Contains(joined, "update"):
		return "Refresh"
	case strings.Contains(joined, "initialize") || strings.HasPrefix(joined, "init"):
		return "Initialize"
	case strings.Contains(joined, "load") && strings.Contains(joined, "settings"):
		return "Load Settings"
	case strings.Contains(joined, "save") && strings.Contains(joined, "settings"):
		return "Save Settings"
	case strings.HasPrefix(segment, "On"):
		return "Event Handler"
	case strings.Contains(joined, "register"):
		return "Register"
	case strings.Contains(joined, "show"):
		return "Show"
	case strings.Contains(joined, "hide"):
		return "Hide"
	case strings.Contains(joined, "toggle"):
		return "Toggle"
	case strings.Contains(joined, "create"):
		return "Create"
	case strings.Contains(joined, "destroy"):
		return "Destroy"
	case strings.Contains(joined, "apply"):
		return "Apply"
	case strings.Contains(joined, "set"):
		return "Set"
	case strings.Contains(joined, "get"):
		return "Get"
	default:
		return strings.Title(tokens[0])
	}
}

func splitIdentifier(value string) []string {
	if value == "" {
		return nil
	}
	value = strings.ReplaceAll(value, ".", " ")
	value = strings.ReplaceAll(value, "_", " ")
	var builder strings.Builder
	for index, r := range value {
		if index > 0 && unicode.IsUpper(r) {
			previous := rune(value[index-1])
			if unicode.IsLower(previous) || unicode.IsDigit(previous) {
				builder.WriteRune(' ')
			}
		}
		builder.WriteRune(unicode.ToLower(r))
	}
	return strings.Fields(builder.String())
}

func buildPatternSummary(name string, evidence []string) string {
	count := len(graph.UniqueStrings(evidence))
	if count == 1 {
		return "Observed in a single addon path"
	}
	return fmt.Sprintf("Observed in %d extracted code paths", count)
}

func topSharedKeys(values map[string]map[string]bool) []string {
	scored := []scoredKey{}
	for key, addons := range values {
		if len(addons) < 2 {
			continue
		}
		scored = append(scored, scoredKey{name: key, score: len(addons)})
	}
	sort.Slice(scored, func(i, j int) bool {
		if scored[i].score == scored[j].score {
			return scored[i].name < scored[j].name
		}
		return scored[i].score > scored[j].score
	})
	result := []string{}
	for _, item := range firstNScored(scored, 12) {
		result = append(result, item.name)
	}
	return result
}

func addToSetMap(values map[string]map[string]bool, key string, addon string) {
	trimmed := strings.TrimSpace(key)
	if trimmed == "" {
		return
	}
	if _, ok := values[trimmed]; !ok {
		values[trimmed] = map[string]bool{}
	}
	values[trimmed][addon] = true
}

func firstN(values []string, limit int) []string {
	if len(values) <= limit {
		return values
	}
	return values[:limit]
}

type scoredKey struct {
	name  string
	score int
}

func firstNScored(values []scoredKey, limit int) []scoredKey {
	if len(values) <= limit {
		return values
	}
	return values[:limit]
}

func frameNames(frames []graph.Frame) []string {
	result := make([]string, 0, len(frames))
	for _, frame := range frames {
		result = append(result, frame.Name)
	}
	return graph.UniqueStrings(result)
}

func eventNames(events []graph.EventRegistration) []string {
	result := make([]string, 0, len(events))
	for _, event := range events {
		result = append(result, event.Event)
	}
	return graph.UniqueStrings(result)
}

package semantic_merge

import (
	"sort"
	"strings"

	"roraddons/tools/api_doc_gen/lua_semantic"
	"roraddons/tools/api_doc_gen/xml_lua_binding"
	"roraddons/tools/api_doc_gen/xml_structure"
)

// BuildEnrichedCatalog performs Phase 4: merges XML structural data (Phase 1),
// XML↔Lua bindings (Phase 2), and Lua semantic analysis (Phase 3) into a
// unified EnrichedElementCatalog suitable for documentation generation.
func BuildEnrichedCatalog(
	xmlCorpus *xml_structure.XMLCorpus,
	bindings *xml_lua_binding.XMLLuaBindingSet,
	luaSemantic *lua_semantic.LuaSemanticCorpus,
) *EnrichedElementCatalog {
	catalog := &EnrichedElementCatalog{
		Elements: make(map[string]*EnrichedElement),
	}

	// Step 1: Seed from XML structural profiles (Phase 1)
	for tag, profile := range xmlCorpus.ElementTypes {
		elem := &EnrichedElement{
			Tag:        tag,
			TotalCount: profile.TotalCount,
		}

		// Seen in addons
		for addon := range profile.ByAddon {
			elem.SeenIn = append(elem.SeenIn, addon)
		}
		sort.Strings(elem.SeenIn)

		// Parents
		for parentTag, count := range profile.ParentTags {
			elem.Parents = append(elem.Parents, ElementRef{
				Tag:   parentTag,
				Count: count,
			})
		}
		sortElementRefs(elem.Parents)

		// Named children
		for childTag, count := range profile.NamedChildren {
			elem.Children = append(elem.Children, ElementRef{
				Tag:   childTag,
				Count: count,
			})
		}
		sortElementRefs(elem.Children)

		// Structural children
		for childTag, count := range profile.StructuralChildren {
			scr := StructuralChildRef{
				Tag:   childTag,
				Count: count,
			}
			// Carry over attribute profiles from the structural child's own profile
			if childProfile, ok := xmlCorpus.ElementTypes[childTag]; ok {
				for attrName, attrProfile := range childProfile.Attributes {
					scr.Attributes = append(scr.Attributes, AttributeRef{
						Name:         attrName,
						IsRequired:   attrProfile.IsRequired,
						SampleValues: attrProfile.SampleValues,
						Count:        attrProfile.TotalCount,
						TotalCount:   childProfile.TotalCount,
					})
				}
				sortAttributeRefs(scr.Attributes)
			}
			elem.StructuralChildren = append(elem.StructuralChildren, scr)
		}
		sortStructuralChildRefs(elem.StructuralChildren)

		// Attribute profiles
		for attrName, attrProfile := range profile.Attributes {
			elem.AttributeProfiles = append(elem.AttributeProfiles, AttributeRef{
				Name:         attrName,
				IsRequired:   attrProfile.IsRequired,
				SampleValues: attrProfile.SampleValues,
				Count:        attrProfile.TotalCount,
				TotalCount:   profile.TotalCount,
			})
		}
		sortAttributeRefs(elem.AttributeProfiles)

		// Inheritance
		for base, count := range profile.InheritsBases {
			if count > 0 {
				elem.InheritsBases = append(elem.InheritsBases, base)
			}
		}
		sort.Strings(elem.InheritsBases)
		elem.IsTemplate = profile.IsTemplate

		// Composition snippets
		if len(profile.CompositionSnippets) > 0 {
			elem.CompositionSnippet = bestSnippet(profile.CompositionSnippets)
		}

		catalog.Elements[tag] = elem
	}

	// Step 2: Merge XML↔Lua bindings (Phase 2)
	if bindings != nil {
		for tag, etb := range bindings.ElementTypeBindings {
			elem := catalog.ensureElement(tag)

			// Event bindings
			for event, eb := range etb.EventBindings {
				enrichedBinding := EnrichedEventBinding{
					Event:      event,
					TotalCount: eb.TotalCount,
				}

				for funcName, count := range eb.LuaFunctions {
					bf := BoundFunction{
						Name:     funcName,
						Count:    count,
						Resolved: eb.Resolved > 0,
					}
					enrichedBinding.LuaFunctions = append(enrichedBinding.LuaFunctions, bf)
				}
				sortBoundFunctions(enrichedBinding.LuaFunctions)

				elem.EventBindings = append(elem.EventBindings, enrichedBinding)
			}
			sortEventBindings(elem.EventBindings)

			// Manipulating functions — build from the binding set's frame manipulations.
			// Each manipulation already carries a CallPattern and the function name,
			// so we associate them directly without blind pattern assignment.
			manipByFunc := make(map[string]*LuaManipulatorRef)
			for _, funcName := range etb.ManipulatingFunctions {
				manipByFunc[funcName] = &LuaManipulatorRef{
					Function: funcName,
				}
			}
			// Now fill in patterns/counts from the binding set
			if bindings != nil {
				for _, m := range bindings.FrameManipulations {
					if ref, ok := manipByFunc[m.LuaFunction]; ok {
						ref.CallPattern = m.CallPattern
						ref.Count++
						ref.Confidence = m.Confidence
					}
				}
			}
			for _, ref := range manipByFunc {
				elem.LuaManipulators = append(elem.LuaManipulators, *ref)
			}
		}
	}

	// Step 3: Merge Lua semantic analysis (Phase 3)
	if luaSemantic != nil {
		// Aggregate API calls from handlers per element type
		handlerCallsByType := make(map[string]map[string]int) // tag → api_call → count
		handlerCallEventsByType := make(map[string]map[string]map[string]bool) // tag → api_call → events

		for key, ha := range luaSemantic.HandlerAnalyses {
			parts := strings.SplitN(key, "|", 3)
			if len(parts) < 3 {
				continue
			}
			event := parts[2]

			// Find what element type this handler belongs to
			elemTag := ha.FrameType
			if elemTag == "" {
				continue
			}

			if handlerCallsByType[elemTag] == nil {
				handlerCallsByType[elemTag] = make(map[string]int)
			}
			if handlerCallEventsByType[elemTag] == nil {
				handlerCallEventsByType[elemTag] = make(map[string]map[string]bool)
			}

			// Collect API calls from this handler
			if ha.FunctionAnalysis != nil {
				for _, call := range ha.FunctionAnalysis.Calls {
					if call.IsAPICall {
						handlerCallsByType[elemTag][call.Callee]++
						if handlerCallEventsByType[elemTag][call.Callee] == nil {
							handlerCallEventsByType[elemTag][call.Callee] = make(map[string]bool)
						}
						handlerCallEventsByType[elemTag][call.Callee][event] = true
					}
				}
			}

			// Enrich event bindings with handler argument patterns
			if elem, ok := catalog.Elements[elemTag]; ok {
				for i, eb := range elem.EventBindings {
					if eb.Event == event {
						// Set inferred args from handler analysis
						if len(ha.InferredHandlerArgs) > 0 {
							argParts := make([]string, 0, len(ha.InferredHandlerArgs))
							for _, param := range ha.InferredHandlerArgs {
								argParts = append(argParts, param.Name)
							}
							elem.EventBindings[i].InferredArgs = strings.Join(argParts, ", ")
							elem.EventBindings[i].ArgsConfidence = ha.HandlerConfidence
						}
						elem.EventBindings[i].Category = ha.HandlerCategory

						// Mark bound functions as having analysis
						for j, bf := range elem.EventBindings[i].LuaFunctions {
							if _, ok := luaSemantic.Functions[ha.QualifiedName]; ok {
								elem.EventBindings[i].LuaFunctions[j] = BoundFunction{
									Name:        bf.Name,
									Count:       bf.Count,
									Resolved:    bf.Resolved,
									HasAnalysis: true,
								}
							}
						}
						break
					}
				}

				// Store handler argument patterns
				if len(ha.InferredHandlerArgs) > 0 {
					pattern := HandlerArgPattern{
						Event:      event,
						Confidence: ha.HandlerConfidence,
						Source:     "xml_handler",
					}
					for _, param := range ha.InferredHandlerArgs {
						pattern.ExpectedParams = append(pattern.ExpectedParams, ExpectedParam{
							Position: param.Position,
							Name:     param.Name,
							Type:     param.Type,
							Role:     param.Role,
						})
					}
					elem.HandlerArgPatterns = append(elem.HandlerArgPatterns, pattern)
				}
			}
		}

		// Write aggregated API calls to elements
		for tag, calls := range handlerCallsByType {
			if elem, ok := catalog.Elements[tag]; ok {
				for apiCall, count := range calls {
					var fromEvents []string
					if eventMap, ok := handlerCallEventsByType[tag][apiCall]; ok {
						for evt := range eventMap {
							fromEvents = append(fromEvents, evt)
						}
						sort.Strings(fromEvents)
					}
					elem.LuaAPICalls = append(elem.LuaAPICalls, LuaAPICallRef{
						Function:   apiCall,
						Count:      count,
						FromEvents: fromEvents,
					})
				}
				sortLuaAPICallRefs(elem.LuaAPICalls)
			}
		}
	}

	// Step 4: Compute confidence scores
	for _, elem := range catalog.Elements {
		elem.Score, elem.Confidence = computeConfidence(elem)
	}

	return catalog
}

func (c *EnrichedElementCatalog) ensureElement(tag string) *EnrichedElement {
	if elem, ok := c.Elements[tag]; ok {
		return elem
	}
	elem := &EnrichedElement{Tag: tag}
	c.Elements[tag] = elem
	return elem
}

func computeConfidence(elem *EnrichedElement) (int, string) {
	score := 0

	// Cross-addon presence
	switch {
	case len(elem.SeenIn) >= 4:
		score += 25
	case len(elem.SeenIn) >= 2:
		score += 15
	}

	// Total occurrences
	if elem.TotalCount >= 10 {
		score += 15
	} else if elem.TotalCount >= 3 {
		score += 10
	}

	// Has event bindings
	if len(elem.EventBindings) > 0 {
		score += 20
	}

	// Has structural children
	if len(elem.StructuralChildren) > 0 {
		score += 10
	}

	// Has Lua API calls
	if len(elem.LuaAPICalls) > 0 {
		score += 10
	}

	// Clamp
	if score > 100 {
		score = 100
	}

	switch {
	case score >= 60:
		return score, "HIGH"
	case score >= 30:
		return score, "MEDIUM"
	default:
		return score, "LOW"
	}
}

func bestSnippet(snippets []string) string {
	if len(snippets) == 0 {
		return ""
	}
	best := snippets[0]
	for _, s := range snippets[1:] {
		if len(s) > len(best) {
			best = s
		}
	}
	return best
}

// Sort helpers

func sortElementRefs(refs []ElementRef) {
	sort.Slice(refs, func(i, j int) bool {
		if refs[i].Count != refs[j].Count {
			return refs[i].Count > refs[j].Count
		}
		return refs[i].Tag < refs[j].Tag
	})
}

func sortStructuralChildRefs(refs []StructuralChildRef) {
	sort.Slice(refs, func(i, j int) bool {
		if refs[i].Count != refs[j].Count {
			return refs[i].Count > refs[j].Count
		}
		return refs[i].Tag < refs[j].Tag
	})
}

func sortAttributeRefs(refs []AttributeRef) {
	sort.Slice(refs, func(i, j int) bool {
		// Required first, then by count
		if refs[i].IsRequired != refs[j].IsRequired {
			return refs[i].IsRequired
		}
		if refs[i].Count != refs[j].Count {
			return refs[i].Count > refs[j].Count
		}
		return refs[i].Name < refs[j].Name
	})
}

func sortBoundFunctions(fns []BoundFunction) {
	sort.Slice(fns, func(i, j int) bool {
		if fns[i].Count != fns[j].Count {
			return fns[i].Count > fns[j].Count
		}
		return fns[i].Name < fns[j].Name
	})
}

func sortEventBindings(bindings []EnrichedEventBinding) {
	sort.Slice(bindings, func(i, j int) bool {
		if bindings[i].TotalCount != bindings[j].TotalCount {
			return bindings[i].TotalCount > bindings[j].TotalCount
		}
		return bindings[i].Event < bindings[j].Event
	})
}

func sortLuaAPICallRefs(refs []LuaAPICallRef) {
	sort.Slice(refs, func(i, j int) bool {
		if refs[i].Count != refs[j].Count {
			return refs[i].Count > refs[j].Count
		}
		return refs[i].Function < refs[j].Function
	})
}

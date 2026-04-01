package graph

import (
	"sort"
	"strings"
	"time"
)

func BuildCorpus(sourceRoot string, addons []AddonModel) Corpus {
	corpus := Corpus{
		SourceRoot: sourceRoot,
		Addons:     addons,
	}
	corpus.Graph, corpus.Bindings = BuildGraph(addons)
	corpus.Normalized.GeneratedAt = time.Now().UTC()
	return corpus
}

func BuildGraph(addons []AddonModel) (DocumentationGraph, []Binding) {
	nodes := []Node{}
	edges := []Edge{}
	bindings := []Binding{}
	nodeSeen := map[string]bool{}
	edgeSeen := map[string]bool{}
	functionIndex := map[string]Function{}

	addNode := func(node Node) {
		if node.ID == "" || nodeSeen[node.ID] {
			return
		}
		nodeSeen[node.ID] = true
		nodes = append(nodes, node)
	}

	addEdge := func(edge Edge) {
		key := string(edge.Type) + "|" + edge.From + "|" + edge.To + "|" + edge.Metadata["addon"] + "|" + edge.Metadata["event"]
		if edge.From == "" || edge.To == "" || edgeSeen[key] {
			return
		}
		edgeSeen[key] = true
		edges = append(edges, edge)
	}

	for _, addon := range addons {
		for _, function := range addon.Functions {
			functionIndex[function.Name] = function
			for _, alias := range function.Aliases {
				functionIndex[alias] = function
			}
		}
	}

	ensureImplicitFunction := func(addonName, name string, file string, line int) Function {
		if existing, ok := functionIndex[name]; ok {
			return existing
		}
		implicit := Function{
			Addon: addonName,
			Name:  name,
			File:  file,
			Line:  line,
			Kind:  "implicit",
		}
		functionIndex[name] = implicit
		addNode(Node{
			ID:    functionNodeID(addonName, name),
			Type:  NodeFunction,
			Name:  name,
			Addon: addonName,
			File:  file,
			Line:  line,
			Metadata: map[string]string{
				"kind": "implicit",
			},
		})
		return implicit
	}

	for _, addon := range addons {
		rootModuleName := addon.Name
		addNode(Node{
			ID:    moduleNodeID(addon.Name, rootModuleName),
			Type:  NodeModule,
			Name:  rootModuleName,
			Addon: addon.Name,
			File:  addon.Manifest.Path,
			Line:  1,
			Metadata: map[string]string{
				"kind": "addon-root",
			},
		})

		for _, module := range addon.Modules {
			addNode(Node{
				ID:    moduleNodeID(module.Addon, module.Name),
				Type:  NodeModule,
				Name:  module.Name,
				Addon: module.Addon,
				File:  module.File,
				Line:  module.Line,
				Metadata: map[string]string{
					"kind": module.Kind,
				},
			})
			owner := OwnerName(module.Name)
			if owner == "" {
				owner = rootModuleName
			}
			addEdge(Edge{
				From: moduleNodeID(module.Addon, module.Name),
				To:   moduleNodeID(module.Addon, owner),
				Type: EdgeDefinedIn,
				Metadata: map[string]string{
					"addon": addon.Name,
				},
			})
		}

		for _, function := range addon.Functions {
			addNode(Node{
				ID:    functionNodeID(function.Addon, function.Name),
				Type:  NodeFunction,
				Name:  function.Name,
				Addon: function.Addon,
				File:  function.File,
				Line:  function.Line,
				Metadata: map[string]string{
					"kind":  function.Kind,
					"local": boolString(function.Local),
				},
			})
			owner := function.Module
			if owner == "" {
				owner = rootModuleName
			}
			addEdge(Edge{
				From: functionNodeID(function.Addon, function.Name),
				To:   moduleNodeID(function.Addon, owner),
				Type: EdgeDefinedIn,
				Metadata: map[string]string{
					"addon": addon.Name,
				},
			})

			for _, call := range function.Calls {
				callee, ok := functionIndex[call.Name]
				if !ok {
					continue
				}
				addEdge(Edge{
					From: functionNodeID(function.Addon, function.Name),
					To:   functionNodeID(callee.Addon, callee.Name),
					Type: EdgeCalls,
					Metadata: map[string]string{
						"addon": addon.Name,
					},
				})

				if strings.EqualFold(call.Name, "BroadcastEvent") || strings.HasSuffix(call.Name, ".TriggerEvent") {
					if len(call.Arguments) > 0 {
						eventName := NormalizeEventName(call.Arguments[0])
						addNode(Node{
							ID:    eventNodeID(eventName),
							Type:  NodeEvent,
							Name:  eventName,
							Addon: addon.Name,
							File:  function.File,
							Line:  call.Line,
							Metadata: map[string]string{
								"kind": "triggered",
							},
						})
						addEdge(Edge{
							From: functionNodeID(function.Addon, function.Name),
							To:   eventNodeID(eventName),
							Type: EdgeTriggered,
							Metadata: map[string]string{
								"addon": addon.Name,
								"event": eventName,
							},
						})
					}
				}
			}
		}

		for _, frame := range addon.Frames {
			metadata := map[string]string{}
			if frame.Inherits != "" {
				metadata["inherits"] = frame.Inherits
			}
			if frame.Template {
				metadata["template"] = "true"
			}
			addNode(Node{
				ID:       frameNodeID(frame.Addon, frame.Name),
				Type:     NodeFrame,
				Name:     frame.Name,
				Addon:    frame.Addon,
				File:     frame.File,
				Line:     frame.Line,
				Metadata: metadata,
			})
			addEdge(Edge{
				From: frameNodeID(frame.Addon, frame.Name),
				To:   moduleNodeID(frame.Addon, rootModuleName),
				Type: EdgeDefinedIn,
				Metadata: map[string]string{
					"addon": addon.Name,
				},
			})
		}

		for _, event := range addon.Events {
			eventName := NormalizeEventName(event.Event)
			addNode(Node{
				ID:    eventNodeID(eventName),
				Type:  NodeEvent,
				Name:  eventName,
				Addon: addon.Name,
				File:  event.File,
				Line:  event.Line,
				Metadata: map[string]string{
					"registrar": event.Registrar,
					"scope":     event.Scope,
				},
			})
			handler := ensureImplicitFunction(addon.Name, event.Handler, event.File, event.Line)
			addEdge(Edge{
				From: eventNodeID(eventName),
				To:   functionNodeID(handler.Addon, handler.Name),
				Type: EdgeHandledBy,
				Metadata: map[string]string{
					"addon": addon.Name,
					"event": eventName,
				},
			})
		}

		for _, handler := range addon.Handlers {
			functionName := handler.Function
			if functionName == "" {
				continue
			}
			resolvedFunction := ensureImplicitFunction(addon.Name, functionName, handler.File, handler.Line)
			eventName := NormalizeEventName(handler.Event)
			addNode(Node{
				ID:    eventNodeID(eventName),
				Type:  NodeEvent,
				Name:  eventName,
				Addon: addon.Name,
				File:  handler.File,
				Line:  handler.Line,
				Metadata: map[string]string{
					"kind": "xml-handler",
				},
			})
			addEdge(Edge{
				From: eventNodeID(eventName),
				To:   functionNodeID(resolvedFunction.Addon, resolvedFunction.Name),
				Type: EdgeHandledBy,
				Metadata: map[string]string{
					"addon": addon.Name,
					"event": eventName,
				},
			})
			addEdge(Edge{
				From: frameNodeID(addon.Name, handler.Frame),
				To:   functionNodeID(resolvedFunction.Addon, resolvedFunction.Name),
				Type: EdgeHandledBy,
				Metadata: map[string]string{
					"addon": addon.Name,
					"event": eventName,
				},
			})
			bindings = append(bindings, Binding{
				Addon:       addon.Name,
				Frame:       handler.Frame,
				Event:       handler.Event,
				XMLFunction: handler.Function,
				LuaFunction: resolvedFunction.Name,
				File:        handler.File,
				Line:        handler.Line,
				Resolved:    resolvedFunction.Kind != "implicit",
			})
		}
	}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Type == nodes[j].Type {
			if nodes[i].Addon == nodes[j].Addon {
				return nodes[i].Name < nodes[j].Name
			}
			return nodes[i].Addon < nodes[j].Addon
		}
		return nodes[i].Type < nodes[j].Type
	})
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].Type == edges[j].Type {
			if edges[i].From == edges[j].From {
				return edges[i].To < edges[j].To
			}
			return edges[i].From < edges[j].From
		}
		return edges[i].Type < edges[j].Type
	})

	return DocumentationGraph{Nodes: nodes, Edges: edges}, UniqueSortedBindings(bindings)
}

func moduleNodeID(addon, name string) string {
	return "module:" + addon + ":" + name
}

func functionNodeID(addon, name string) string {
	return "function:" + addon + ":" + name
}

func frameNodeID(addon, name string) string {
	return "frame:" + addon + ":" + name
}

func eventNodeID(name string) string {
	return "event:" + name
}

func boolString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

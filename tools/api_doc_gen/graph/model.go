package graph

import (
	"regexp"
	"sort"
	"strings"
	"time"
)

type Manifest struct {
	Name            string
	Path            string
	Type            string
	Version         string
	Files           []string
	SavedVariables  []string
	CreateWindows   []string
	InitializeCalls []string
	UpdateCalls     []string
	ShutdownCalls   []string
	Metadata        map[string]string
}

type AddonSpec struct {
	Name     string
	Path     string
	Manifest Manifest
}

type Call struct {
	Name      string
	Line      int
	Arguments []string
}

type EventRegistration struct {
	Addon          string
	Registrar      string
	Event          string
	Handler        string
	Window         string
	Scope          string
	SourceFunction string
	File           string
	Line           int
}

type StateVariable struct {
	Addon     string
	Name      string
	Owner     string
	File      string
	Line      int
	ValueType string
	Saved     bool
	Scope     string
}

type Module struct {
	Addon   string
	Name    string
	File    string
	Line    int
	Kind    string
	Members []string
	Aliases []string
}

type Function struct {
	Addon       string
	Name        string
	Aliases     []string
	Module      string
	File        string
	Line        int
	EndLine     int
	Params      []string
	Calls       []Call
	Events      []EventRegistration
	StateWrites []string
	Local       bool
	Kind        string
}

type Frame struct {
	Addon                string
	Name                 string
	Type                 string
	Parent               string
	File                 string
	Line                 int
	Children             []string
	StructuralChildTypes []string // element type tags that appear as unnamed structural children (e.g. ListData, ListColumns)
	Inherits             string
	Attributes           map[string]string
	Template             bool
}

type XMLHandler struct {
	Addon    string
	Frame    string
	Event    string
	Function string
	File     string
	Line     int
}

type LuaFileResult struct {
	Functions []Function
	Modules   []Module
	Events    []EventRegistration
	State     []StateVariable
}

type XMLFileResult struct {
	Frames   []Frame
	Handlers []XMLHandler
}

type AddonModel struct {
	Name      string
	Root      string
	Manifest  Manifest
	Functions []Function
	Modules   []Module
	Events    []EventRegistration
	State     []StateVariable
	Frames    []Frame
	Handlers  []XMLHandler
}

type NodeType string

const (
	NodeFunction NodeType = "function"
	NodeFrame    NodeType = "frame"
	NodeEvent    NodeType = "event"
	NodeModule   NodeType = "module"
)

type EdgeType string

const (
	EdgeCalls     EdgeType = "calls"
	EdgeHandledBy EdgeType = "handled_by"
	EdgeTriggered EdgeType = "triggered_by"
	EdgeDefinedIn EdgeType = "defined_in"
)

type Node struct {
	ID       string
	Type     NodeType
	Name     string
	Addon    string
	File     string
	Line     int
	Metadata map[string]string
}

type Edge struct {
	From     string
	To       string
	Type     EdgeType
	Metadata map[string]string
}

type DocumentationGraph struct {
	Nodes []Node
	Edges []Edge
}

type Binding struct {
	Addon       string
	Frame       string
	Event       string
	XMLFunction string
	LuaFunction string
	File        string
	Line        int
	Resolved    bool
}

type ConceptGroup struct {
	Canonical string
	Kind      string
	Members   []string
	Addons    []string
	Summary   string
}

type Pattern struct {
	Name     string
	Category string
	Summary  string
	Evidence []string
}

type NamingConvention struct {
	Rule     string
	Evidence []string
}

type ExecutionStep struct {
	Phase    string
	Detail   string
	Evidence []string
}

type ExecutionFlow struct {
	Addon string
	Steps []ExecutionStep
}

type NormalizedData struct {
	ConceptGroups      []ConceptGroup
	Patterns           []Pattern
	NamingConventions  []NamingConvention
	ExecutionFlows     []ExecutionFlow
	CommonEvents       []string
	CommonModules      []string
	CommonStateOwners  []string
	CommonFrameParents []string
	GeneratedAt        time.Time
}

type Corpus struct {
	SourceRoot string
	Addons     []AddonModel
	Graph      DocumentationGraph
	Bindings   []Binding
	Normalized NormalizedData
}

var slugSanitizer = regexp.MustCompile(`[^A-Za-z0-9._-]+`)

func Slug(value string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return "unnamed"
	}
	trimmed = strings.ReplaceAll(trimmed, "::", ".")
	trimmed = strings.ReplaceAll(trimmed, ":", ".")
	trimmed = strings.ReplaceAll(trimmed, "\\", ".")
	trimmed = strings.ReplaceAll(trimmed, "/", ".")
	trimmed = slugSanitizer.ReplaceAllString(trimmed, "_")
	trimmed = strings.Trim(trimmed, "._-")
	if trimmed == "" {
		return "unnamed"
	}
	return trimmed
}

func OwnerName(name string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return ""
	}
	trimmed = strings.ReplaceAll(trimmed, ":", ".")
	lastDot := strings.LastIndex(trimmed, ".")
	if lastDot < 0 {
		return ""
	}
	return trimmed[:lastDot]
}

func LastSegment(name string) string {
	trimmed := strings.TrimSpace(name)
	trimmed = strings.ReplaceAll(trimmed, ":", ".")
	if trimmed == "" {
		return ""
	}
	parts := strings.Split(trimmed, ".")
	return parts[len(parts)-1]
}

func RootSegment(name string) string {
	trimmed := strings.TrimSpace(name)
	trimmed = strings.ReplaceAll(trimmed, ":", ".")
	if trimmed == "" {
		return ""
	}
	parts := strings.Split(trimmed, ".")
	return parts[0]
}

func UniqueStrings(values []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" || seen[trimmed] {
			continue
		}
		seen[trimmed] = true
		result = append(result, trimmed)
	}
	sort.Strings(result)
	return result
}

func UniqueSortedModules(modules []Module) []Module {
	seen := map[string]bool{}
	result := make([]Module, 0, len(modules))
	for _, module := range modules {
		key := module.Addon + "|" + module.Name
		if seen[key] {
			continue
		}
		seen[key] = true
		module.Members = UniqueStrings(module.Members)
		module.Aliases = UniqueStrings(module.Aliases)
		result = append(result, module)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Addon == result[j].Addon {
			return result[i].Name < result[j].Name
		}
		return result[i].Addon < result[j].Addon
	})
	return result
}

func UniqueSortedFunctions(functions []Function) []Function {
	seen := map[string]bool{}
	result := make([]Function, 0, len(functions))
	for _, function := range functions {
		key := function.Addon + "|" + function.Name + "|" + function.File + "|" + Slug(strings.TrimSpace(function.Kind))
		if seen[key] {
			continue
		}
		seen[key] = true
		function.Aliases = UniqueStrings(function.Aliases)
		function.StateWrites = UniqueStrings(function.StateWrites)
		result = append(result, function)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Addon == result[j].Addon {
			return result[i].Name < result[j].Name
		}
		return result[i].Addon < result[j].Addon
	})
	return result
}

func UniqueSortedEvents(events []EventRegistration) []EventRegistration {
	seen := map[string]bool{}
	result := make([]EventRegistration, 0, len(events))
	for _, event := range events {
		key := strings.Join([]string{
			event.Addon,
			event.Registrar,
			event.Event,
			event.Handler,
			event.Window,
			event.SourceFunction,
			event.File,
		}, "|")
		if seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, event)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Event == result[j].Event {
			if result[i].Addon == result[j].Addon {
				return result[i].Handler < result[j].Handler
			}
			return result[i].Addon < result[j].Addon
		}
		return result[i].Event < result[j].Event
	})
	return result
}

func UniqueSortedState(state []StateVariable) []StateVariable {
	seen := map[string]bool{}
	result := make([]StateVariable, 0, len(state))
	for _, item := range state {
		key := strings.Join([]string{item.Addon, item.Name, item.File, item.Scope}, "|")
		if seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, item)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Addon == result[j].Addon {
			return result[i].Name < result[j].Name
		}
		return result[i].Addon < result[j].Addon
	})
	return result
}

func UniqueSortedFrames(frames []Frame) []Frame {
	seen := map[string]bool{}
	result := make([]Frame, 0, len(frames))
	for _, frame := range frames {
		key := frame.Addon + "|" + frame.Name
		if seen[key] {
			continue
		}
		seen[key] = true
		frame.Children = UniqueStrings(frame.Children)
		result = append(result, frame)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Addon == result[j].Addon {
			return result[i].Name < result[j].Name
		}
		return result[i].Addon < result[j].Addon
	})
	return result
}

func UniqueSortedHandlers(handlers []XMLHandler) []XMLHandler {
	seen := map[string]bool{}
	result := make([]XMLHandler, 0, len(handlers))
	for _, handler := range handlers {
		key := strings.Join([]string{handler.Addon, handler.Frame, handler.Event, handler.Function, handler.File}, "|")
		if seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, handler)
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
	return result
}

func UniqueSortedBindings(bindings []Binding) []Binding {
	seen := map[string]bool{}
	result := make([]Binding, 0, len(bindings))
	for _, binding := range bindings {
		key := strings.Join([]string{binding.Addon, binding.Frame, binding.Event, binding.XMLFunction, binding.LuaFunction, binding.File}, "|")
		if seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, binding)
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
	return result
}

func NormalizePath(value string) string {
	return strings.ReplaceAll(value, "\\", "/")
}

func NormalizeEventName(value string) string {
	trimmed := strings.TrimSpace(value)
	trimmed = strings.Trim(trimmed, "\"'")
	trimmed = strings.ReplaceAll(trimmed, " ", "")
	if trimmed == "" {
		return "unknown-event"
	}
	return trimmed
}

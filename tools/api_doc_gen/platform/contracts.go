package platform

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"roraddons/tools/api_doc_gen/graph"
)

type ContractModel struct {
	Root        string
	XMLTrees    []contractXMLTree
	LuaAnalyses []contractLuaAnalysis
	Links       []contractXMLLuaLinks
	LoadedAt    time.Time
}

// contractSemanticInput is the contract-native internal semantic dataset used by
// build/semantic/enrichment internals. It intentionally excludes markdown-only
// artifacts such as flows/examples/saved variable docs.
type contractSemanticInput struct {
	Root      string
	Functions []FunctionDoc
	Frames    []FrameDoc
	Handlers  []HandlerDoc
	Events    []EventDoc
	Bindings  []BindingDoc
}

type contractXMLTree struct {
	SchemaVersion string            `json:"schema_version"`
	Addon         string            `json:"addon"`
	SourceFile    string            `json:"source_file"`
	Roots         []string          `json:"roots"`
	Nodes         []contractXMLNode `json:"nodes"`
}

type contractXMLNode struct {
	NodeID       string             `json:"node_id"`
	ParentID     *string            `json:"parent_id"`
	Tag          string             `json:"tag"`
	RawName      string             `json:"raw_name"`
	ResolvedName string             `json:"resolved_name"`
	Attributes   map[string]string  `json:"attributes"`
	Flags        contractNodeFlags  `json:"flags"`
	Source       contractNodeSource `json:"source"`
}

type contractNodeFlags struct {
	IsNamed      bool `json:"is_named"`
	IsStructural bool `json:"is_structural"`
	IsIgnored    bool `json:"is_ignored"`
	IsTemplate   bool `json:"is_template"`
}

type contractNodeSource struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

type contractLuaAnalysis struct {
	SchemaVersion string                    `json:"schema_version"`
	Addon         string                    `json:"addon"`
	SourceFile    string                    `json:"source_file"`
	Functions     []contractLuaFunction     `json:"functions"`
	Registrations []contractLuaRegistration `json:"registrations"`
}

type contractLuaFunction struct {
	FunctionID      string              `json:"function_id"`
	DeclaredName    string              `json:"declared_name"`
	QualifiedName   string              `json:"qualified_name"`
	ShortName       string              `json:"short_name"`
	ScopeKind       string              `json:"scope_kind"`
	DeclarationKind string              `json:"declaration_kind"`
	Namespace       string              `json:"namespace"`
	Source          contractLuaFnSource `json:"source"`
	Params          []string            `json:"params"`
	Calls           []contractLuaCall   `json:"calls"`
}

type contractLuaFnSource struct {
	LineStart int `json:"line_start"`
	LineEnd   int `json:"line_end"`
}

type contractLuaCall struct {
	CalleeRaw      string              `json:"callee_raw"`
	CalleeResolved string              `json:"callee_resolved"`
	Source         contractLuaCallLine `json:"source"`
	Arguments      []string            `json:"arguments"`
}

type contractLuaCallLine struct {
	Line int `json:"line"`
}

type contractLuaRegistration struct {
	Registrar       string `json:"registrar"`
	EventResolved   string `json:"event_resolved"`
	EventRaw        string `json:"event_raw"`
	HandlerResolved string `json:"handler_resolved"`
	HandlerRaw      string `json:"handler_raw"`
	Scope           string `json:"scope"`
}

type contractXMLLuaLinks struct {
	SchemaVersion string                `json:"schema_version"`
	Addon         string                `json:"addon"`
	HandlerLinks  []contractHandlerLink `json:"handler_links"`
}

type contractHandlerLink struct {
	XML contractHandlerXMLRef `json:"xml"`
	Lua contractHandlerLuaRef `json:"lua"`
}

type contractHandlerXMLRef struct {
	SourceFile   string `json:"source_file"`
	ParentNodeID string `json:"parent_node_id"`
	Event        string `json:"event"`
	FunctionRaw  string `json:"function_raw"`
	ResolvedName string `json:"resolved_name"`
	RawName      string `json:"raw_name"`
}

type contractHandlerLuaRef struct {
	QualifiedName string `json:"qualified_name"`
	DeclaredName  string `json:"declared_name"`
	ShortName     string `json:"short_name"`
}

func LoadContractInputs(addonAPIRoot string) (ContractModel, error) {
	root, err := filepath.Abs(addonAPIRoot)
	if err != nil {
		return ContractModel{}, fmt.Errorf("resolve addon-api root: %w", err)
	}
	if err := ensureDirectory(root); err != nil {
		return ContractModel{}, err
	}

	xmlTreeDir := filepath.Join(root, "xml-tree")
	luaAnalysisDir := filepath.Join(root, "lua-analysis")
	xmlLuaLinksDir := filepath.Join(root, "xml-lua-links")

	if err := ensureDirectory(xmlTreeDir); err != nil {
		return ContractModel{}, fmt.Errorf("xml-tree: %w", err)
	}
	if err := ensureDirectory(luaAnalysisDir); err != nil {
		return ContractModel{}, fmt.Errorf("lua-analysis: %w", err)
	}
	if err := ensureDirectory(xmlLuaLinksDir); err != nil {
		return ContractModel{}, fmt.Errorf("xml-lua-links: %w", err)
	}

	xmlTreeFiles, err := collectJSONFiles(xmlTreeDir)
	if err != nil {
		return ContractModel{}, fmt.Errorf("xml-tree file discovery: %w", err)
	}
	if len(xmlTreeFiles) == 0 {
		return ContractModel{}, fmt.Errorf("xml-tree: no JSON files found in %s", graph.NormalizePath(xmlTreeDir))
	}

	luaAnalysisFiles, err := collectJSONFiles(luaAnalysisDir)
	if err != nil {
		return ContractModel{}, fmt.Errorf("lua-analysis file discovery: %w", err)
	}
	if len(luaAnalysisFiles) == 0 {
		return ContractModel{}, fmt.Errorf("lua-analysis: no JSON files found in %s", graph.NormalizePath(luaAnalysisDir))
	}

	xmlLuaLinkFiles, err := collectJSONFiles(xmlLuaLinksDir)
	if err != nil {
		return ContractModel{}, fmt.Errorf("xml-lua-links file discovery: %w", err)
	}
	if len(xmlLuaLinkFiles) == 0 {
		return ContractModel{}, fmt.Errorf("xml-lua-links: no JSON files found in %s", graph.NormalizePath(xmlLuaLinksDir))
	}

	model := ContractModel{Root: graph.NormalizePath(root), LoadedAt: time.Now().UTC()}
	for _, path := range xmlTreeFiles {
		var artifact contractXMLTree
		if err := decodeJSONFile(path, &artifact); err != nil {
			return ContractModel{}, fmt.Errorf("xml-tree %s: %w", graph.NormalizePath(path), err)
		}
		if err := validateXMLTreeArtifact(path, artifact); err != nil {
			return ContractModel{}, err
		}
		model.XMLTrees = append(model.XMLTrees, artifact)
	}
	for _, path := range luaAnalysisFiles {
		var artifact contractLuaAnalysis
		if err := decodeJSONFile(path, &artifact); err != nil {
			return ContractModel{}, fmt.Errorf("lua-analysis %s: %w", graph.NormalizePath(path), err)
		}
		if err := validateLuaAnalysisArtifact(path, artifact); err != nil {
			return ContractModel{}, err
		}
		model.LuaAnalyses = append(model.LuaAnalyses, artifact)
	}
	for _, path := range xmlLuaLinkFiles {
		var artifact contractXMLLuaLinks
		if err := decodeJSONFile(path, &artifact); err != nil {
			return ContractModel{}, fmt.Errorf("xml-lua-links %s: %w", graph.NormalizePath(path), err)
		}
		if err := validateXMLLuaLinksArtifact(path, artifact); err != nil {
			return ContractModel{}, err
		}
		model.Links = append(model.Links, artifact)
	}

	return model, nil
}

func semanticInputFromContracts(model ContractModel) contractSemanticInput {
	input := contractSemanticInput{Root: model.Root}

	for _, analysis := range model.LuaAnalyses {
		for _, fn := range analysis.Functions {
			name := firstNonEmpty(fn.ShortName, fn.DeclaredName, fn.QualifiedName)
			qualified := firstNonEmpty(fn.QualifiedName, fn.DeclaredName)
			if strings.TrimSpace(name) == "" {
				continue
			}

			aliases := make([]string, 0, 2)
			if qualified != "" && qualified != name {
				aliases = append(aliases, qualified)
			}
			if fn.DeclaredName != "" && fn.DeclaredName != name && fn.DeclaredName != qualified {
				aliases = append(aliases, fn.DeclaredName)
			}

			calls := make([]CallDoc, 0, len(fn.Calls))
			for _, call := range fn.Calls {
				callName := firstNonEmpty(call.CalleeResolved, call.CalleeRaw)
				if strings.TrimSpace(callName) == "" {
					continue
				}
				calls = append(calls, CallDoc{
					Name:      callName,
					Line:      call.Source.Line,
					Arguments: append([]string(nil), call.Arguments...),
					Raw:       strings.Join(call.Arguments, ", "),
				})
			}

			input.Functions = append(input.Functions, FunctionDoc{
				Name:               name,
				Addon:              analysis.Addon,
				Kind:               fn.DeclarationKind,
				Module:             fn.Namespace,
				Local:              strings.EqualFold(strings.TrimSpace(fn.ScopeKind), "local"),
				Source:             analysis.SourceFile,
				Parameters:         append([]string(nil), fn.Params...),
				Aliases:            aliases,
				Calls:              calls,
				EventRegistrations: nil,
			})
		}

		for _, reg := range analysis.Registrations {
			eventName := graph.NormalizeEventName(firstNonEmpty(reg.EventResolved, reg.EventRaw))
			if strings.TrimSpace(eventName) == "" {
				continue
			}
			handler := firstNonEmpty(reg.HandlerResolved, reg.HandlerRaw)
			sourceEventReg := EventRegistrationDoc{
				Event:   eventName,
				Scope:   reg.Scope,
				Handler: handler,
			}
			registrar := strings.TrimSpace(reg.Registrar)
			if registrar != "" {
				for i := range input.Functions {
					if input.Functions[i].Addon == analysis.Addon && symbolMatches(input.Functions[i], registrar) {
						input.Functions[i].EventRegistrations = append(input.Functions[i].EventRegistrations, sourceEventReg)
					}
				}
			}
		}
	}

	eventByName := map[string]*EventDoc{}
	for _, fn := range input.Functions {
		for _, reg := range fn.EventRegistrations {
			event := ensureEventDoc(eventByName, reg.Event)
			event.LuaRegistrations = append(event.LuaRegistrations, EventUsageDoc{
				Addon:     fn.Addon,
				Registrar: fn.Name,
				Handler:   reg.Handler,
				Scope:     reg.Scope,
			})
		}
	}

	frameByID := map[string]FrameDoc{}
	for _, tree := range model.XMLTrees {
		nodesByID := make(map[string]contractXMLNode, len(tree.Nodes))
		childrenByParent := map[string][]contractXMLNode{}
		for _, node := range tree.Nodes {
			nodesByID[node.NodeID] = node
			if node.ParentID != nil && strings.TrimSpace(*node.ParentID) != "" {
				childrenByParent[*node.ParentID] = append(childrenByParent[*node.ParentID], node)
			}
		}

		for _, node := range tree.Nodes {
			if !node.Flags.IsNamed && strings.TrimSpace(node.ResolvedName) == "" && strings.TrimSpace(node.RawName) == "" {
				continue
			}
			name := firstNonEmpty(node.ResolvedName, node.RawName, node.Attributes["name"])
			if strings.TrimSpace(name) == "" {
				continue
			}
			parentName, parentType := resolveNamedParent(node, nodesByID)

			childNames := make([]string, 0)
			childElementTypes := make([]string, 0)
			structuralChildTypes := make([]string, 0)
			structuralChildAttrs := map[string][]string{}
			handlers := make([]FrameHandlerDoc, 0)

			for _, child := range childrenByParent[node.NodeID] {
				if strings.EqualFold(strings.TrimSpace(child.Tag), "EventHandler") {
					eventName := graph.NormalizeEventName(child.Attributes["event"])
					fn := strings.TrimSpace(child.Attributes["function"])
					if eventName != "" && fn != "" {
						handlers = append(handlers, FrameHandlerDoc{Event: eventName, Function: fn})
					}
					continue
				}
				childName := firstNonEmpty(child.ResolvedName, child.RawName, child.Attributes["name"])
				if strings.TrimSpace(childName) != "" {
					childNames = appendUniqueString(childNames, childName)
					childElementTypes = appendUniqueString(childElementTypes, child.Tag)
					continue
				}
				if strings.TrimSpace(child.Tag) != "" {
					structuralChildTypes = appendUniqueString(structuralChildTypes, child.Tag)
					attrKeys := mapKeys(child.Attributes)
					sort.Strings(attrKeys)
					for _, key := range attrKeys {
						structuralChildAttrs[child.Tag] = appendUniqueString(structuralChildAttrs[child.Tag], key)
					}
				}
			}

			frameKey := tree.Addon + "|" + name
			frameByID[frameKey] = FrameDoc{
				Name:                    name,
				RawName:                 firstNonEmpty(node.RawName, name),
				Addon:                   tree.Addon,
				Type:                    node.Tag,
				Parent:                  parentName,
				ParentType:              parentType,
				Inherits:                firstNonEmpty(node.Attributes["inherits"], node.Attributes["template"]),
				Template:                node.Flags.IsTemplate,
				Source:                  tree.SourceFile,
				Children:                childNames,
				ChildElementTypes:       childElementTypes,
				StructuralChildTypes:    structuralChildTypes,
				StructuralChildAttrKeys: structuralChildAttrs,
				Attributes:              cloneStringMap(node.Attributes),
				Handlers:                handlers,
			}
		}
	}

	for _, frame := range frameByID {
		input.Frames = append(input.Frames, frame)
		for _, h := range frame.Handlers {
			input.Handlers = append(input.Handlers, HandlerDoc{
				Addon:    frame.Addon,
				Frame:    frame.Name,
				Event:    h.Event,
				Function: h.Function,
				Source:   frame.Source,
			})
			event := ensureEventDoc(eventByName, h.Event)
			event.XMLHandlers = append(event.XMLHandlers, EventHandlerUsageDoc{
				Addon:    frame.Addon,
				Frame:    frame.Name,
				Function: h.Function,
			})
		}
	}

	for _, links := range model.Links {
		for _, link := range links.HandlerLinks {
			eventName := graph.NormalizeEventName(link.XML.Event)
			if strings.TrimSpace(eventName) == "" {
				continue
			}
			luaFn := firstNonEmpty(link.Lua.QualifiedName, link.Lua.DeclaredName, link.Lua.ShortName)
			if strings.TrimSpace(luaFn) == "" {
				continue
			}

			input.Bindings = append(input.Bindings, BindingDoc{
				Addon:       links.Addon,
				Frame:       firstNonEmpty(link.XML.ResolvedName, link.XML.RawName),
				Event:       eventName,
				XMLFunction: strings.TrimSpace(link.XML.FunctionRaw),
				LuaFunction: luaFn,
				Resolved:    true,
			})
		}
	}

	for _, event := range eventByName {
		input.Events = append(input.Events, *event)
	}

	sort.Slice(input.Functions, func(i, j int) bool {
		if input.Functions[i].Addon == input.Functions[j].Addon {
			return input.Functions[i].Name < input.Functions[j].Name
		}
		return input.Functions[i].Addon < input.Functions[j].Addon
	})
	sort.Slice(input.Frames, func(i, j int) bool {
		if input.Frames[i].Addon == input.Frames[j].Addon {
			return input.Frames[i].Name < input.Frames[j].Name
		}
		return input.Frames[i].Addon < input.Frames[j].Addon
	})
	sort.Slice(input.Handlers, func(i, j int) bool {
		if input.Handlers[i].Addon == input.Handlers[j].Addon {
			if input.Handlers[i].Frame == input.Handlers[j].Frame {
				return input.Handlers[i].Event < input.Handlers[j].Event
			}
			return input.Handlers[i].Frame < input.Handlers[j].Frame
		}
		return input.Handlers[i].Addon < input.Handlers[j].Addon
	})
	sort.Slice(input.Events, func(i, j int) bool { return input.Events[i].Name < input.Events[j].Name })

	return input
}

func ensureDirectory(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("missing directory %s", graph.NormalizePath(path))
		}
		return fmt.Errorf("stat %s: %w", graph.NormalizePath(path), err)
	}
	if !info.IsDir() {
		return fmt.Errorf("expected directory, got file %s", graph.NormalizePath(path))
	}
	return nil
}

func collectJSONFiles(root string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		if strings.EqualFold(filepath.Ext(d.Name()), ".json") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(files)
	return files, nil
}

func decodeJSONFile(path string, out any) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(out); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	if decoder.More() {
		return fmt.Errorf("invalid JSON: trailing tokens")
	}
	return nil
}

func validateXMLTreeArtifact(path string, artifact contractXMLTree) error {
	if artifact.SchemaVersion != "phase1-tree/v1" {
		return fmt.Errorf("xml-tree %s: schema_version must be phase1-tree/v1", graph.NormalizePath(path))
	}
	if strings.TrimSpace(artifact.Addon) == "" {
		return fmt.Errorf("xml-tree %s: missing addon", graph.NormalizePath(path))
	}
	if strings.TrimSpace(artifact.SourceFile) == "" {
		return fmt.Errorf("xml-tree %s: missing source_file", graph.NormalizePath(path))
	}
	if artifact.Nodes == nil || len(artifact.Nodes) == 0 {
		return fmt.Errorf("xml-tree %s: missing nodes", graph.NormalizePath(path))
	}
	for _, node := range artifact.Nodes {
		if strings.TrimSpace(node.NodeID) == "" {
			return fmt.Errorf("xml-tree %s: node missing node_id", graph.NormalizePath(path))
		}
		if strings.TrimSpace(node.Tag) == "" {
			return fmt.Errorf("xml-tree %s: node %s missing tag", graph.NormalizePath(path), node.NodeID)
		}
	}
	return nil
}

func validateLuaAnalysisArtifact(path string, artifact contractLuaAnalysis) error {
	if artifact.SchemaVersion != "phase2-lua/v1" {
		return fmt.Errorf("lua-analysis %s: schema_version must be phase2-lua/v1", graph.NormalizePath(path))
	}
	if strings.TrimSpace(artifact.Addon) == "" {
		return fmt.Errorf("lua-analysis %s: missing addon", graph.NormalizePath(path))
	}
	if strings.TrimSpace(artifact.SourceFile) == "" {
		return fmt.Errorf("lua-analysis %s: missing source_file", graph.NormalizePath(path))
	}
	if artifact.Functions == nil {
		return fmt.Errorf("lua-analysis %s: missing functions", graph.NormalizePath(path))
	}
	if artifact.Registrations == nil {
		return fmt.Errorf("lua-analysis %s: missing registrations", graph.NormalizePath(path))
	}
	for _, fn := range artifact.Functions {
		if strings.TrimSpace(fn.FunctionID) == "" {
			return fmt.Errorf("lua-analysis %s: function missing function_id", graph.NormalizePath(path))
		}
		if strings.TrimSpace(firstNonEmpty(fn.QualifiedName, fn.DeclaredName, fn.ShortName)) == "" {
			return fmt.Errorf("lua-analysis %s: function %s missing identity", graph.NormalizePath(path), fn.FunctionID)
		}
		if fn.Calls == nil {
			return fmt.Errorf("lua-analysis %s: function %s missing calls", graph.NormalizePath(path), fn.FunctionID)
		}
	}
	return nil
}

func validateXMLLuaLinksArtifact(path string, artifact contractXMLLuaLinks) error {
	if artifact.SchemaVersion != "phase3-xml-lua-link/v1" {
		return fmt.Errorf("xml-lua-links %s: schema_version must be phase3-xml-lua-link/v1", graph.NormalizePath(path))
	}
	if strings.TrimSpace(artifact.Addon) == "" {
		return fmt.Errorf("xml-lua-links %s: missing addon", graph.NormalizePath(path))
	}
	if artifact.HandlerLinks == nil {
		return fmt.Errorf("xml-lua-links %s: missing handler_links", graph.NormalizePath(path))
	}
	return nil
}

func resolveNamedParent(node contractXMLNode, nodesByID map[string]contractXMLNode) (string, string) {
	parentID := node.ParentID
	for parentID != nil {
		parent, ok := nodesByID[*parentID]
		if !ok {
			break
		}
		name := firstNonEmpty(parent.ResolvedName, parent.RawName, parent.Attributes["name"])
		if strings.TrimSpace(name) != "" {
			return name, strings.TrimSpace(parent.Tag)
		}
		parentID = parent.ParentID
	}
	return "", ""
}

func cloneStringMap(input map[string]string) map[string]string {
	if input == nil {
		return map[string]string{}
	}
	out := make(map[string]string, len(input))
	for k, v := range input {
		out[k] = v
	}
	return out
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if trimmed := strings.TrimSpace(v); trimmed != "" {
			return trimmed
		}
	}
	return ""
}

func ensureEventDoc(index map[string]*EventDoc, eventName string) *EventDoc {
	name := graph.NormalizeEventName(eventName)
	if strings.TrimSpace(name) == "" {
		name = strings.TrimSpace(eventName)
	}
	existing, ok := index[name]
	if ok {
		return existing
	}
	created := &EventDoc{Name: name}
	index[name] = created
	return created
}

func symbolMatches(doc FunctionDoc, symbol string) bool {
	if strings.EqualFold(strings.TrimSpace(doc.Name), strings.TrimSpace(symbol)) {
		return true
	}
	for _, alias := range doc.Aliases {
		if strings.EqualFold(strings.TrimSpace(alias), strings.TrimSpace(symbol)) {
			return true
		}
	}
	return false
}

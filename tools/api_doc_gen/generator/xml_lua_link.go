package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

const xmlLuaLinkSchemaVersion = "phase3-xml-lua-link/v1"

type xmlLuaLinkArtifact struct {
	SchemaVersion     string                   `json:"schema_version"`
	Addon             string                   `json:"addon"`
	Sources           xmlLuaLinkSources        `json:"sources"`
	HandlerLinks      []xmlLuaHandlerLink      `json:"handler_links"`
	UIUsageLinks      []xmlLuaUIUsageLink      `json:"ui_usage_links"`
	RegistrationLinks []xmlLuaRegistrationLink `json:"registration_links"`
	Unresolved        xmlLuaUnresolved         `json:"unresolved"`
}

type xmlLuaLinkSources struct {
	XMLTreeFiles     []string `json:"xml_tree_files"`
	LuaAnalysisFiles []string `json:"lua_analysis_files"`
}

type xmlLuaHandlerLink struct {
	LinkID     string              `json:"link_id"`
	MatchType  string              `json:"match_type"`
	MatchField string              `json:"match_field"`
	XML        xmlLuaHandlerXMLRef `json:"xml"`
	Lua        xmlLuaFunctionRef   `json:"lua"`
}

type xmlLuaHandlerXMLRef struct {
	SourceFile   string `json:"source_file"`
	NodeID       string `json:"node_id"`
	ParentNodeID string `json:"parent_node_id"`
	Event        string `json:"event"`
	FunctionRaw  string `json:"function_raw"`
	Tag          string `json:"tag"`
	ResolvedName string `json:"resolved_name"`
	RawName      string `json:"raw_name"`
}

type xmlLuaFunctionRef struct {
	SourceFile        string `json:"source_file"`
	FunctionID        string `json:"function_id"`
	QualifiedName     string `json:"qualified_name"`
	DeclaredName      string `json:"declared_name"`
	ShortName         string `json:"short_name"`
	ScopeKind         string `json:"scope_kind"`
	DeclarationKind   string `json:"declaration_kind"`
	ReceiverName      string `json:"receiver_name"`
	ReceiverSeparator string `json:"receiver_separator"`
	LineStart         int    `json:"line_start"`
	LineEnd           int    `json:"line_end"`
}

type xmlLuaUIUsageLink struct {
	LinkID    string            `json:"link_id"`
	MatchType string            `json:"match_type"`
	XML       xmlLuaNamedXMLRef `json:"xml"`
	Lua       xmlLuaCallRef     `json:"lua"`
}

type xmlLuaNamedXMLRef struct {
	SourceFile   string `json:"source_file"`
	NodeID       string `json:"node_id"`
	Tag          string `json:"tag"`
	ResolvedName string `json:"resolved_name"`
	RawName      string `json:"raw_name"`
}

type xmlLuaCallRef struct {
	SourceFile      string `json:"source_file"`
	FunctionID      string `json:"function_id"`
	CallID          string `json:"call_id"`
	CalleeResolved  string `json:"callee_resolved"`
	Line            int    `json:"line"`
	MatchedArgument string `json:"matched_argument"`
}

type xmlLuaRegistrationLink struct {
	LinkID          string                `json:"link_id"`
	MatchType       string                `json:"match_type"`
	XML             xmlLuaNamedXMLRef     `json:"xml"`
	Registration    xmlLuaRegistrationRef `json:"registration"`
	HandlerFunction *xmlLuaFunctionRef    `json:"handler_function,omitempty"`
}

type xmlLuaRegistrationRef struct {
	SourceFile      string `json:"source_file"`
	RegistrationID  string `json:"registration_id"`
	Registrar       string `json:"registrar"`
	Scope           string `json:"scope"`
	EventRaw        string `json:"event_raw"`
	EventResolved   string `json:"event_resolved"`
	HandlerRaw      string `json:"handler_raw"`
	HandlerResolved string `json:"handler_resolved"`
	WindowRaw       string `json:"window_raw"`
	WindowResolved  string `json:"window_resolved"`
	Line            int    `json:"line"`
}

type xmlLuaUnresolved struct {
	HandlerBindings []xmlLuaUnresolvedHandler      `json:"handler_bindings"`
	UIUsages        []xmlLuaUnresolvedUIUsage      `json:"ui_usages"`
	Registrations   []xmlLuaUnresolvedRegistration `json:"registrations"`
}

type xmlLuaUnresolvedHandler struct {
	SourceFile  string `json:"source_file"`
	NodeID      string `json:"node_id"`
	Event       string `json:"event"`
	FunctionRaw string `json:"function_raw"`
	Reason      string `json:"reason"`
}

type xmlLuaUnresolvedUIUsage struct {
	SourceFile     string `json:"source_file"`
	FunctionID     string `json:"function_id"`
	CallID         string `json:"call_id"`
	CalleeResolved string `json:"callee_resolved"`
	Literal        string `json:"literal"`
	Reason         string `json:"reason"`
}

type xmlLuaUnresolvedRegistration struct {
	SourceFile      string `json:"source_file"`
	RegistrationID  string `json:"registration_id"`
	Reason          string `json:"reason"`
	WindowResolved  string `json:"window_resolved"`
	HandlerResolved string `json:"handler_resolved"`
}

type luaFunctionRecord struct {
	SourceFile string
	Entry      luaAnalysisFunctionEntry
}

type luaCallRecord struct {
	SourceFile string
	FunctionID string
	Call       luaAnalysisCallEntry
}

type luaRegistrationRecord struct {
	SourceFile   string
	Registration luaAnalysisRegistration
}

type handlerLookupResult struct {
	function   luaFunctionRecord
	matchType  string
	matchField string
	reason     string
}

func writeXMLLuaLinkArtifacts(outputRoot string) error {
	xmlTreeRoot := filepath.Join(outputRoot, "xml-tree")
	luaAnalysisRoot := filepath.Join(outputRoot, "lua-analysis")
	linkRoot := filepath.Join(outputRoot, "xml-lua-links")

	addons, err := collectLinkAddons(xmlTreeRoot, luaAnalysisRoot)
	if err != nil {
		return err
	}

	for _, addon := range addons {
		xmlArtifacts, xmlFiles, err := loadXMLTreeArtifactsForAddon(xmlTreeRoot, addon)
		if err != nil {
			return err
		}
		luaArtifacts, luaFiles, err := loadLuaAnalysisArtifactsForAddon(luaAnalysisRoot, addon)
		if err != nil {
			return err
		}

		artifact := buildXMLLuaLinkArtifact(addon, xmlFiles, luaFiles, xmlArtifacts, luaArtifacts)
		outputPath := filepath.Join(linkRoot, addon, "links.json")
		if err := writeJSONFile(outputPath, artifact); err != nil {
			return fmt.Errorf("write xml-lua link artifact %s: %w", outputPath, err)
		}
	}

	return nil
}

func collectLinkAddons(xmlTreeRoot string, luaAnalysisRoot string) ([]string, error) {
	addonSet := map[string]bool{}
	for _, root := range []string{xmlTreeRoot, luaAnalysisRoot} {
		entries, err := os.ReadDir(root)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, fmt.Errorf("list link input root %s: %w", root, err)
		}
		for _, entry := range entries {
			if entry.IsDir() {
				addonSet[entry.Name()] = true
			}
		}
	}
	addons := make([]string, 0, len(addonSet))
	for addon := range addonSet {
		addons = append(addons, addon)
	}
	sort.Strings(addons)
	return addons, nil
}

func loadXMLTreeArtifactsForAddon(root, addon string) ([]phase1TreeArtifact, []string, error) {
	addonRoot := filepath.Join(root, addon)
	files, err := collectFilesWithSuffix(addonRoot, ".tree.json")
	if err != nil {
		return nil, nil, err
	}
	artifacts := make([]phase1TreeArtifact, 0, len(files))
	for _, file := range files {
		var artifact phase1TreeArtifact
		if err := readJSONFile(file, &artifact); err != nil {
			return nil, nil, fmt.Errorf("read xml-tree artifact %s: %w", file, err)
		}
		artifacts = append(artifacts, artifact)
	}
	return artifacts, toSlashPathsRelativeTo(addonRoot, files), nil
}

func loadLuaAnalysisArtifactsForAddon(root, addon string) ([]luaAnalysisArtifact, []string, error) {
	addonRoot := filepath.Join(root, addon)
	files, err := collectFilesWithSuffix(addonRoot, ".analysis.json")
	if err != nil {
		return nil, nil, err
	}
	artifacts := make([]luaAnalysisArtifact, 0, len(files))
	for _, file := range files {
		var artifact luaAnalysisArtifact
		if err := readJSONFile(file, &artifact); err != nil {
			return nil, nil, fmt.Errorf("read lua-analysis artifact %s: %w", file, err)
		}
		artifacts = append(artifacts, artifact)
	}
	return artifacts, toSlashPathsRelativeTo(addonRoot, files), nil
}

func collectFilesWithSuffix(root string, suffix string) ([]string, error) {
	if _, err := os.Stat(root); err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, fmt.Errorf("stat directory %s: %w", root, err)
	}
	files := []string{}
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(d.Name()), strings.ToLower(suffix)) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walk directory %s: %w", root, err)
	}
	sort.Strings(files)
	return files, nil
}

func toSlashPathsRelativeTo(root string, files []string) []string {
	relative := make([]string, 0, len(files))
	for _, file := range files {
		rel, err := filepath.Rel(root, file)
		if err != nil {
			relative = append(relative, filepath.ToSlash(file))
			continue
		}
		relative = append(relative, filepath.ToSlash(rel))
	}
	sort.Strings(relative)
	return relative
}

func readJSONFile(path string, target any) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, target)
}

func writeJSONFile(path string, value any) error {
	content, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json %s: %w", path, err)
	}
	content = append(content, '\n')
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create json directory %s: %w", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, content, 0o644); err != nil {
		return fmt.Errorf("write json file %s: %w", path, err)
	}
	return nil
}

func buildXMLLuaLinkArtifact(addon string, xmlFiles, luaFiles []string, xmlArtifacts []phase1TreeArtifact, luaArtifacts []luaAnalysisArtifact) xmlLuaLinkArtifact {
	artifact := xmlLuaLinkArtifact{
		SchemaVersion: xmlLuaLinkSchemaVersion,
		Addon:         addon,
		Sources: xmlLuaLinkSources{
			XMLTreeFiles:     nonNilStrings(xmlFiles),
			LuaAnalysisFiles: nonNilStrings(luaFiles),
		},
		HandlerLinks:      []xmlLuaHandlerLink{},
		UIUsageLinks:      []xmlLuaUIUsageLink{},
		RegistrationLinks: []xmlLuaRegistrationLink{},
		Unresolved: xmlLuaUnresolved{
			HandlerBindings: []xmlLuaUnresolvedHandler{},
			UIUsages:        []xmlLuaUnresolvedUIUsage{},
			Registrations:   []xmlLuaUnresolvedRegistration{},
		},
	}

	xmlNodesByID := map[string]xmlLuaNamedXMLRef{}
	xmlNodesByResolved := map[string][]xmlLuaNamedXMLRef{}
	xmlNodesByResolvedNormalized := map[string][]xmlLuaNamedXMLRef{}
	xmlSourceByNodeID := map[string]string{}
	xmlNodeParent := map[string]string{}

	type handlerCandidate struct {
		xmlNode xmlLuaNamedXMLRef
		source  string
		nodeID  string
		event   string
		fnRaw   string
	}
	handlerCandidates := []handlerCandidate{}

	for _, tree := range xmlArtifacts {
		sourceFile := normalizeRelativeXMLPath(tree.SourceFile)
		for _, node := range tree.Nodes {
			nodeRef := xmlLuaNamedXMLRef{
				SourceFile:   sourceFile,
				NodeID:       node.NodeID,
				Tag:          node.Tag,
				ResolvedName: node.ResolvedName,
				RawName:      node.RawName,
			}
			xmlNodesByID[node.NodeID] = nodeRef
			xmlSourceByNodeID[node.NodeID] = sourceFile
			if node.ParentID != nil {
				xmlNodeParent[node.NodeID] = *node.ParentID
			}

			if node.Flags.IsNamed && strings.TrimSpace(node.ResolvedName) != "" {
				key := strings.TrimSpace(node.ResolvedName)
				xmlNodesByResolved[key] = append(xmlNodesByResolved[key], nodeRef)
				xmlNodesByResolvedNormalized[normalizeSymbol(key)] = append(xmlNodesByResolvedNormalized[normalizeSymbol(key)], nodeRef)
			}

			if strings.EqualFold(node.Tag, "EventHandler") {
				fnRaw := attributeValue(node.Attributes, "function")
				if strings.TrimSpace(fnRaw) == "" {
					continue
				}
				handlerCandidates = append(handlerCandidates, handlerCandidate{
					xmlNode: nodeRef,
					source:  sourceFile,
					nodeID:  node.NodeID,
					event:   attributeValue(node.Attributes, "event"),
					fnRaw:   fnRaw,
				})
			}
		}
	}

	sort.Slice(handlerCandidates, func(i, j int) bool {
		if handlerCandidates[i].source == handlerCandidates[j].source {
			return handlerCandidates[i].nodeID < handlerCandidates[j].nodeID
		}
		return handlerCandidates[i].source < handlerCandidates[j].source
	})

	luaFunctions := []luaFunctionRecord{}
	luaCalls := []luaCallRecord{}
	luaRegistrations := []luaRegistrationRecord{}

	byQualifiedExact := map[string][]luaFunctionRecord{}
	byDeclaredExact := map[string][]luaFunctionRecord{}
	byShortExact := map[string][]luaFunctionRecord{}
	byQualifiedNorm := map[string][]luaFunctionRecord{}
	byDeclaredNorm := map[string][]luaFunctionRecord{}
	byShortNorm := map[string][]luaFunctionRecord{}

	for _, luaArtifact := range luaArtifacts {
		source := normalizeRelativeLuaPath(luaArtifact.SourceFile)
		for _, fn := range luaArtifact.Functions {
			record := luaFunctionRecord{SourceFile: source, Entry: fn}
			luaFunctions = append(luaFunctions, record)
			appendByKey(byQualifiedExact, strings.TrimSpace(fn.QualifiedName), record)
			appendByKey(byDeclaredExact, strings.TrimSpace(fn.DeclaredName), record)
			appendByKey(byShortExact, strings.TrimSpace(fn.ShortName), record)
			appendByKey(byQualifiedNorm, normalizeSymbol(fn.QualifiedName), record)
			appendByKey(byDeclaredNorm, normalizeSymbol(fn.DeclaredName), record)
			appendByKey(byShortNorm, normalizeSymbol(fn.ShortName), record)
			for _, call := range fn.Calls {
				luaCalls = append(luaCalls, luaCallRecord{SourceFile: source, FunctionID: fn.FunctionID, Call: call})
			}
		}
		for _, reg := range luaArtifact.Registrations {
			luaRegistrations = append(luaRegistrations, luaRegistrationRecord{SourceFile: source, Registration: reg})
		}
	}

	sort.Slice(luaFunctions, func(i, j int) bool {
		if luaFunctions[i].SourceFile == luaFunctions[j].SourceFile {
			return luaFunctions[i].Entry.FunctionID < luaFunctions[j].Entry.FunctionID
		}
		return luaFunctions[i].SourceFile < luaFunctions[j].SourceFile
	})
	sort.Slice(luaCalls, func(i, j int) bool {
		if luaCalls[i].SourceFile == luaCalls[j].SourceFile {
			if luaCalls[i].FunctionID == luaCalls[j].FunctionID {
				return luaCalls[i].Call.CallID < luaCalls[j].Call.CallID
			}
			return luaCalls[i].FunctionID < luaCalls[j].FunctionID
		}
		return luaCalls[i].SourceFile < luaCalls[j].SourceFile
	})
	sort.Slice(luaRegistrations, func(i, j int) bool {
		if luaRegistrations[i].SourceFile == luaRegistrations[j].SourceFile {
			return luaRegistrations[i].Registration.RegistrationID < luaRegistrations[j].Registration.RegistrationID
		}
		return luaRegistrations[i].SourceFile < luaRegistrations[j].SourceFile
	})

	lookupHandler := func(raw string) handlerLookupResult {
		clean := cleanSymbol(raw)
		if clean == "" {
			return handlerLookupResult{reason: "empty_handler_name"}
		}
		if found, reason := uniqueFunctionMatch(byQualifiedExact[clean]); reason == "" {
			return handlerLookupResult{function: found, matchType: "exact", matchField: "qualified_name"}
		} else if reason == "ambiguous" {
			return handlerLookupResult{reason: "ambiguous_exact_qualified"}
		}
		if found, reason := uniqueFunctionMatch(byDeclaredExact[clean]); reason == "" {
			return handlerLookupResult{function: found, matchType: "exact", matchField: "declared_name"}
		} else if reason == "ambiguous" {
			return handlerLookupResult{reason: "ambiguous_exact_declared"}
		}
		if found, reason := uniqueFunctionMatch(byShortExact[clean]); reason == "" {
			return handlerLookupResult{function: found, matchType: "exact", matchField: "short_name"}
		} else if reason == "ambiguous" {
			return handlerLookupResult{reason: "ambiguous_exact_short"}
		}

		norm := normalizeSymbol(clean)
		if found, reason := uniqueFunctionMatch(byQualifiedNorm[norm]); reason == "" {
			return handlerLookupResult{function: found, matchType: "normalized", matchField: "qualified_name"}
		} else if reason == "ambiguous" {
			return handlerLookupResult{reason: "ambiguous_normalized_qualified"}
		}
		if found, reason := uniqueFunctionMatch(byDeclaredNorm[norm]); reason == "" {
			return handlerLookupResult{function: found, matchType: "normalized", matchField: "declared_name"}
		} else if reason == "ambiguous" {
			return handlerLookupResult{reason: "ambiguous_normalized_declared"}
		}
		if found, reason := uniqueFunctionMatch(byShortNorm[norm]); reason == "" {
			return handlerLookupResult{function: found, matchType: "normalized", matchField: "short_name"}
		} else if reason == "ambiguous" {
			return handlerLookupResult{reason: "ambiguous_normalized_short"}
		}

		return handlerLookupResult{reason: "no_lua_function_match"}
	}

	for _, candidate := range handlerCandidates {
		lookup := lookupHandler(candidate.fnRaw)
		if lookup.reason != "" {
			artifact.Unresolved.HandlerBindings = append(artifact.Unresolved.HandlerBindings, xmlLuaUnresolvedHandler{
				SourceFile:  candidate.source,
				NodeID:      candidate.nodeID,
				Event:       candidate.event,
				FunctionRaw: candidate.fnRaw,
				Reason:      lookup.reason,
			})
			continue
		}
		parentID := xmlNodeParent[candidate.nodeID]
		artifact.HandlerLinks = append(artifact.HandlerLinks, xmlLuaHandlerLink{
			MatchType:  lookup.matchType,
			MatchField: lookup.matchField,
			XML: xmlLuaHandlerXMLRef{
				SourceFile:   candidate.source,
				NodeID:       candidate.nodeID,
				ParentNodeID: parentID,
				Event:        candidate.event,
				FunctionRaw:  candidate.fnRaw,
				Tag:          candidate.xmlNode.Tag,
				ResolvedName: candidate.xmlNode.ResolvedName,
				RawName:      candidate.xmlNode.RawName,
			},
			Lua: functionRefFromRecord(lookup.function),
		})
	}

	for _, call := range luaCalls {
		if !isUIWindowCall(call.Call.CalleeResolved) {
			continue
		}
		literals := extractStringLiterals(call.Call.Arguments)
		if len(literals) == 0 {
			continue
		}
		for _, literal := range literals {
			if strings.TrimSpace(literal) == "" {
				continue
			}
			matched, matchType, reason := matchXMLNodeByResolvedName(literal, xmlNodesByResolved, xmlNodesByResolvedNormalized)
			if reason != "" {
				artifact.Unresolved.UIUsages = append(artifact.Unresolved.UIUsages, xmlLuaUnresolvedUIUsage{
					SourceFile:     call.SourceFile,
					FunctionID:     call.FunctionID,
					CallID:         call.Call.CallID,
					CalleeResolved: call.Call.CalleeResolved,
					Literal:        literal,
					Reason:         reason,
				})
				continue
			}
			artifact.UIUsageLinks = append(artifact.UIUsageLinks, xmlLuaUIUsageLink{
				MatchType: matchType,
				XML:       matched,
				Lua: xmlLuaCallRef{
					SourceFile:      call.SourceFile,
					FunctionID:      call.FunctionID,
					CallID:          call.Call.CallID,
					CalleeResolved:  call.Call.CalleeResolved,
					Line:            call.Call.Source.Line,
					MatchedArgument: literal,
				},
			})
		}
	}

	for _, registration := range luaRegistrations {
		reg := registration.Registration
		windowValue := strings.TrimSpace(reg.WindowResolved)
		if windowValue == "" {
			artifact.Unresolved.Registrations = append(artifact.Unresolved.Registrations, xmlLuaUnresolvedRegistration{
				SourceFile:      registration.SourceFile,
				RegistrationID:  reg.RegistrationID,
				Reason:          "empty_window_resolved",
				WindowResolved:  reg.WindowResolved,
				HandlerResolved: reg.HandlerResolved,
			})
			continue
		}

		xmlNode, matchType, reason := matchXMLNodeByResolvedName(windowValue, xmlNodesByResolved, xmlNodesByResolvedNormalized)
		if reason != "" {
			artifact.Unresolved.Registrations = append(artifact.Unresolved.Registrations, xmlLuaUnresolvedRegistration{
				SourceFile:      registration.SourceFile,
				RegistrationID:  reg.RegistrationID,
				Reason:          reason,
				WindowResolved:  reg.WindowResolved,
				HandlerResolved: reg.HandlerResolved,
			})
			continue
		}

		handlerLookup := lookupHandler(reg.HandlerResolved)
		var handlerRef *xmlLuaFunctionRef
		if handlerLookup.reason == "" {
			resolved := functionRefFromRecord(handlerLookup.function)
			handlerRef = &resolved
		} else {
			artifact.Unresolved.Registrations = append(artifact.Unresolved.Registrations, xmlLuaUnresolvedRegistration{
				SourceFile:      registration.SourceFile,
				RegistrationID:  reg.RegistrationID,
				Reason:          "handler_" + handlerLookup.reason,
				WindowResolved:  reg.WindowResolved,
				HandlerResolved: reg.HandlerResolved,
			})
		}

		artifact.RegistrationLinks = append(artifact.RegistrationLinks, xmlLuaRegistrationLink{
			MatchType: matchType,
			XML:       xmlNode,
			Registration: xmlLuaRegistrationRef{
				SourceFile:      registration.SourceFile,
				RegistrationID:  reg.RegistrationID,
				Registrar:       reg.Registrar,
				Scope:           reg.Scope,
				EventRaw:        reg.EventRaw,
				EventResolved:   reg.EventResolved,
				HandlerRaw:      reg.HandlerRaw,
				HandlerResolved: reg.HandlerResolved,
				WindowRaw:       reg.WindowRaw,
				WindowResolved:  reg.WindowResolved,
				Line:            reg.Source.Line,
			},
			HandlerFunction: handlerRef,
		})
	}

	sort.Slice(artifact.HandlerLinks, func(i, j int) bool {
		if artifact.HandlerLinks[i].XML.SourceFile == artifact.HandlerLinks[j].XML.SourceFile {
			return artifact.HandlerLinks[i].XML.NodeID < artifact.HandlerLinks[j].XML.NodeID
		}
		return artifact.HandlerLinks[i].XML.SourceFile < artifact.HandlerLinks[j].XML.SourceFile
	})
	sort.Slice(artifact.UIUsageLinks, func(i, j int) bool {
		if artifact.UIUsageLinks[i].Lua.SourceFile == artifact.UIUsageLinks[j].Lua.SourceFile {
			if artifact.UIUsageLinks[i].Lua.FunctionID == artifact.UIUsageLinks[j].Lua.FunctionID {
				return artifact.UIUsageLinks[i].Lua.CallID < artifact.UIUsageLinks[j].Lua.CallID
			}
			return artifact.UIUsageLinks[i].Lua.FunctionID < artifact.UIUsageLinks[j].Lua.FunctionID
		}
		return artifact.UIUsageLinks[i].Lua.SourceFile < artifact.UIUsageLinks[j].Lua.SourceFile
	})
	sort.Slice(artifact.RegistrationLinks, func(i, j int) bool {
		if artifact.RegistrationLinks[i].Registration.SourceFile == artifact.RegistrationLinks[j].Registration.SourceFile {
			return artifact.RegistrationLinks[i].Registration.RegistrationID < artifact.RegistrationLinks[j].Registration.RegistrationID
		}
		return artifact.RegistrationLinks[i].Registration.SourceFile < artifact.RegistrationLinks[j].Registration.SourceFile
	})
	sort.Slice(artifact.Unresolved.HandlerBindings, func(i, j int) bool {
		if artifact.Unresolved.HandlerBindings[i].SourceFile == artifact.Unresolved.HandlerBindings[j].SourceFile {
			return artifact.Unresolved.HandlerBindings[i].NodeID < artifact.Unresolved.HandlerBindings[j].NodeID
		}
		return artifact.Unresolved.HandlerBindings[i].SourceFile < artifact.Unresolved.HandlerBindings[j].SourceFile
	})
	sort.Slice(artifact.Unresolved.UIUsages, func(i, j int) bool {
		if artifact.Unresolved.UIUsages[i].SourceFile == artifact.Unresolved.UIUsages[j].SourceFile {
			if artifact.Unresolved.UIUsages[i].FunctionID == artifact.Unresolved.UIUsages[j].FunctionID {
				return artifact.Unresolved.UIUsages[i].CallID < artifact.Unresolved.UIUsages[j].CallID
			}
			return artifact.Unresolved.UIUsages[i].FunctionID < artifact.Unresolved.UIUsages[j].FunctionID
		}
		return artifact.Unresolved.UIUsages[i].SourceFile < artifact.Unresolved.UIUsages[j].SourceFile
	})
	sort.Slice(artifact.Unresolved.Registrations, func(i, j int) bool {
		if artifact.Unresolved.Registrations[i].SourceFile == artifact.Unresolved.Registrations[j].SourceFile {
			return artifact.Unresolved.Registrations[i].RegistrationID < artifact.Unresolved.Registrations[j].RegistrationID
		}
		return artifact.Unresolved.Registrations[i].SourceFile < artifact.Unresolved.Registrations[j].SourceFile
	})

	for i := range artifact.HandlerLinks {
		artifact.HandlerLinks[i].LinkID = fmt.Sprintf("h%d", i+1)
	}
	for i := range artifact.UIUsageLinks {
		artifact.UIUsageLinks[i].LinkID = fmt.Sprintf("u%d", i+1)
	}
	for i := range artifact.RegistrationLinks {
		artifact.RegistrationLinks[i].LinkID = fmt.Sprintf("r%d", i+1)
	}

	_ = xmlNodesByID
	_ = xmlSourceByNodeID
	return artifact
}

func appendByKey(index map[string][]luaFunctionRecord, key string, value luaFunctionRecord) {
	trimmed := strings.TrimSpace(key)
	if trimmed == "" {
		return
	}
	index[trimmed] = append(index[trimmed], value)
}

func uniqueFunctionMatch(matches []luaFunctionRecord) (luaFunctionRecord, string) {
	if len(matches) == 0 {
		return luaFunctionRecord{}, "none"
	}
	if len(matches) > 1 {
		return luaFunctionRecord{}, "ambiguous"
	}
	return matches[0], ""
}

func functionRefFromRecord(record luaFunctionRecord) xmlLuaFunctionRef {
	entry := record.Entry
	return xmlLuaFunctionRef{
		SourceFile:        record.SourceFile,
		FunctionID:        entry.FunctionID,
		QualifiedName:     entry.QualifiedName,
		DeclaredName:      entry.DeclaredName,
		ShortName:         entry.ShortName,
		ScopeKind:         entry.ScopeKind,
		DeclarationKind:   entry.DeclarationKind,
		ReceiverName:      entry.ReceiverName,
		ReceiverSeparator: entry.ReceiverSeparator,
		LineStart:         entry.Source.LineStart,
		LineEnd:           entry.Source.LineEnd,
	}
}

func attributeValue(attributes sortedStringMap, key string) string {
	if value, ok := attributes[key]; ok {
		return value
	}
	for attrKey, attrValue := range attributes {
		if strings.EqualFold(attrKey, key) {
			return attrValue
		}
	}
	return ""
}

func cleanSymbol(value string) string {
	trimmed := strings.TrimSpace(value)
	trimmed = strings.Trim(trimmed, "\"'")
	return strings.TrimSpace(trimmed)
}

func normalizeSymbol(value string) string {
	cleaned := cleanSymbol(value)
	if cleaned == "" {
		return ""
	}
	var b strings.Builder
	b.Grow(len(cleaned))
	for _, r := range cleaned {
		switch {
		case unicode.IsLetter(r), unicode.IsDigit(r):
			b.WriteRune(unicode.ToLower(r))
		case r == '_', r == '.', r == ':':
			if r == ':' {
				b.WriteRune('.')
			} else {
				b.WriteRune(r)
			}
		}
	}
	return b.String()
}

func extractStringLiterals(arguments []string) []string {
	literals := []string{}
	for _, argument := range arguments {
		for _, quote := range []rune{'\'', '"'} {
			start := -1
			for i, r := range argument {
				if r == quote {
					if start >= 0 {
						literal := strings.TrimSpace(argument[start+1 : i])
						if literal != "" {
							literals = append(literals, literal)
						}
						start = -1
					} else {
						start = i
					}
				}
			}
		}
	}
	return literals
}

func matchXMLNodeByResolvedName(value string, exact map[string][]xmlLuaNamedXMLRef, normalized map[string][]xmlLuaNamedXMLRef) (xmlLuaNamedXMLRef, string, string) {
	if matches := exact[strings.TrimSpace(value)]; len(matches) == 1 {
		return matches[0], "exact", ""
	} else if len(matches) > 1 {
		return xmlLuaNamedXMLRef{}, "", "ambiguous_exact_resolved_name"
	}
	normalizedValue := normalizeSymbol(value)
	if normalizedValue == "" {
		return xmlLuaNamedXMLRef{}, "", "empty_match_value"
	}
	if matches := normalized[normalizedValue]; len(matches) == 1 {
		return matches[0], "normalized", ""
	} else if len(matches) > 1 {
		return xmlLuaNamedXMLRef{}, "", "ambiguous_normalized_resolved_name"
	}
	return xmlLuaNamedXMLRef{}, "", "no_xml_resolved_name_match"
}

func isUIWindowCall(callee string) bool {
	if strings.TrimSpace(callee) == "" {
		return false
	}
	if strings.HasPrefix(callee, "Window") || strings.HasPrefix(callee, "Label") || strings.HasPrefix(callee, "Button") || strings.HasPrefix(callee, "DynamicImage") || strings.HasPrefix(callee, "TextEditBox") || strings.HasPrefix(callee, "ComboBox") || strings.HasPrefix(callee, "ScrollWindow") || strings.HasPrefix(callee, "StatusBar") || strings.HasPrefix(callee, "SliderBar") || strings.HasPrefix(callee, "ListBox") {
		return true
	}
	return callee == "CreateWindow" || callee == "CreateWindowFromTemplate" || callee == "DestroyWindow" || callee == "DoesWindowExist"
}

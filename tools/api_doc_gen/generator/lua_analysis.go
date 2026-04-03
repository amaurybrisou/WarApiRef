package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"roraddons/tools/api_doc_gen/graph"
	"roraddons/tools/api_doc_gen/lua_parser"
)

const luaAnalysisSchemaVersion = "phase2-lua/v1"

type luaAnalysisArtifact struct {
	SchemaVersion string                     `json:"schema_version"`
	Addon         string                     `json:"addon"`
	SourceFile    string                     `json:"source_file"`
	Functions     []luaAnalysisFunctionEntry `json:"functions"`
	Namespaces    []string                   `json:"namespaces"`
	Registrations []luaAnalysisRegistration  `json:"registrations"`
}

type luaAnalysisFunctionEntry struct {
	FunctionID        string                 `json:"function_id"`
	DeclaredName      string                 `json:"declared_name"`
	QualifiedName     string                 `json:"qualified_name"`
	ShortName         string                 `json:"short_name"`
	ScopeKind         string                 `json:"scope_kind"`
	DeclarationKind   string                 `json:"declaration_kind"`
	ReceiverName      string                 `json:"receiver_name"`
	ReceiverSeparator string                 `json:"receiver_separator"`
	Namespace         string                 `json:"namespace"`
	Source            luaAnalysisSourceRange `json:"source"`
	Params            []string               `json:"params"`
	Calls             []luaAnalysisCallEntry `json:"calls"`
}

type luaAnalysisSourceRange struct {
	LineStart int `json:"line_start"`
	LineEnd   int `json:"line_end"`
}

type luaAnalysisCallEntry struct {
	CallID         string            `json:"call_id"`
	CalleeRaw      string            `json:"callee_raw"`
	CalleeResolved string            `json:"callee_resolved"`
	Source         luaAnalysisSource `json:"source"`
	Arguments      []string          `json:"arguments"`
	IsUIAPI        bool              `json:"is_ui_api"`
	IsEventAPI     bool              `json:"is_event_api"`
}

type luaAnalysisSource struct {
	Line int `json:"line"`
}

type luaAnalysisRegistration struct {
	RegistrationID   string            `json:"registration_id"`
	SourceFunctionID string            `json:"source_function_id"`
	Registrar        string            `json:"registrar"`
	EventRaw         string            `json:"event_raw"`
	EventResolved    string            `json:"event_resolved"`
	HandlerRaw       string            `json:"handler_raw"`
	HandlerResolved  string            `json:"handler_resolved"`
	Scope            string            `json:"scope"`
	WindowRaw        string            `json:"window_raw"`
	WindowResolved   string            `json:"window_resolved"`
	Source           luaAnalysisSource `json:"source"`
}

func writeLuaAnalysisArtifacts(outputRoot string, corpus graph.Corpus) error {
	for _, addon := range corpus.Addons {
		for _, relativeLuaPath := range resolveLuaInputs(addon) {
			normalizedRelative := normalizeRelativeLuaPath(relativeLuaPath)
			absolutePath := filepath.Join(addon.Root, filepath.FromSlash(normalizedRelative))
			parsed, err := lua_parser.ParseFile(addon.Name, absolutePath, addon.Manifest)
			if err != nil {
				fmt.Fprintf(os.Stderr, "warning: skipping lua analysis %s/%s: %v\n", addon.Name, normalizedRelative, err)
				continue
			}

			artifact := buildLuaAnalysisArtifact(addon.Name, normalizedRelative, parsed)
			relativeArtifactPath := filepath.Join("lua-analysis", addon.Name, filepath.FromSlash(normalizedRelative)+".analysis.json")
			outputPath := filepath.Join(outputRoot, relativeArtifactPath)
			if err := writeLuaAnalysisArtifact(outputPath, artifact); err != nil {
				return err
			}
		}
	}
	return nil
}

func resolveLuaInputs(addon graph.AddonModel) []string {
	manifestLua := make([]string, 0)
	for _, manifestFile := range addon.Manifest.Files {
		if strings.ToLower(filepath.Ext(manifestFile)) != ".lua" {
			continue
		}
		normalizedRelative := normalizeRelativeLuaPath(manifestFile)
		absolutePath := filepath.Join(addon.Root, filepath.FromSlash(normalizedRelative))
		if _, err := os.Stat(absolutePath); err != nil {
			continue
		}
		manifestLua = append(manifestLua, normalizedRelative)
	}
	manifestLua = uniqueSortedStrings(manifestLua)
	if len(manifestLua) > 0 {
		return manifestLua
	}

	return discoverLuaInputsOnDisk(addon.Root)
}

func discoverLuaInputsOnDisk(addonRoot string) []string {
	relativePaths := make([]string, 0)
	_ = filepath.WalkDir(addonRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			if strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.EqualFold(filepath.Ext(path), ".lua") {
			return nil
		}
		relative, relErr := filepath.Rel(addonRoot, path)
		if relErr != nil {
			return nil
		}
		relativePaths = append(relativePaths, normalizeRelativeLuaPath(relative))
		return nil
	})
	return uniqueSortedStrings(relativePaths)
}

func uniqueSortedStrings(values []string) []string {
	if len(values) == 0 {
		return []string{}
	}
	set := make(map[string]bool, len(values))
	for _, v := range values {
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			continue
		}
		set[trimmed] = true
	}
	result := make([]string, 0, len(set))
	for v := range set {
		result = append(result, v)
	}
	sort.Strings(result)
	return result
}

func writeLuaAnalysisArtifact(outputPath string, artifact luaAnalysisArtifact) error {
	content, err := json.MarshalIndent(artifact, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal lua analysis artifact %s: %w", outputPath, err)
	}
	content = append(content, '\n')
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("create lua analysis directory %s: %w", filepath.Dir(outputPath), err)
	}
	if err := os.WriteFile(outputPath, content, 0o644); err != nil {
		return fmt.Errorf("write lua analysis artifact %s: %w", outputPath, err)
	}
	return nil
}

func buildLuaAnalysisArtifact(addon, sourceFile string, parsed graph.LuaFileResult) luaAnalysisArtifact {
	functions := make([]graph.Function, 0, len(parsed.Functions))
	functions = append(functions, parsed.Functions...)
	sort.Slice(functions, func(i, j int) bool {
		if functions[i].Line == functions[j].Line {
			if functions[i].DeclarationOrder == functions[j].DeclarationOrder {
				return functions[i].Name < functions[j].Name
			}
			return functions[i].DeclarationOrder < functions[j].DeclarationOrder
		}
		return functions[i].Line < functions[j].Line
	})

	artifact := luaAnalysisArtifact{
		SchemaVersion: luaAnalysisSchemaVersion,
		Addon:         addon,
		SourceFile:    normalizeRelativeLuaPath(sourceFile),
		Functions:     make([]luaAnalysisFunctionEntry, 0, len(functions)),
		Namespaces:    make([]string, 0),
		Registrations: make([]luaAnalysisRegistration, 0),
	}

	namespaceSet := make(map[string]bool)
	for _, module := range parsed.Modules {
		if strings.TrimSpace(module.Name) == "" {
			continue
		}
		namespaceSet[module.Name] = true
	}

	functionIDByName := make(map[string]string)
	for index, fn := range functions {
		functionID := fmt.Sprintf("f%d", index+1)
		functionIDByName[fn.Name] = functionID

		calls := make([]luaAnalysisCallEntry, 0, len(fn.Calls))
		for callIndex, call := range fn.Calls {
			calls = append(calls, luaAnalysisCallEntry{
				CallID:         fmt.Sprintf("c%d", callIndex+1),
				CalleeRaw:      defaultString(call.CalleeRaw, call.Name),
				CalleeResolved: defaultString(call.CalleeResolved, call.Name),
				Source: luaAnalysisSource{
					Line: call.Line,
				},
				Arguments:  nonNilStrings(call.Arguments),
				IsUIAPI:    isLuaUIAPI(call.Name),
				IsEventAPI: isLuaEventAPI(call.Name),
			})
		}

		artifact.Functions = append(artifact.Functions, luaAnalysisFunctionEntry{
			FunctionID:        functionID,
			DeclaredName:      defaultString(fn.DeclaredName, fn.Name),
			QualifiedName:     fn.Name,
			ShortName:         defaultString(fn.ShortName, graph.LastSegment(fn.Name)),
			ScopeKind:         defaultString(fn.ScopeKind, scopeKindFromBool(fn.Local)),
			DeclarationKind:   defaultString(fn.DeclarationKind, "function_decl"),
			ReceiverName:      fn.ReceiverName,
			ReceiverSeparator: fn.ReceiverSeparator,
			Namespace:         defaultString(fn.Module, graph.OwnerName(fn.Name)),
			Source: luaAnalysisSourceRange{
				LineStart: fn.Line,
				LineEnd:   fn.EndLine,
			},
			Params: nonNilStrings(fn.Params),
			Calls:  calls,
		})
	}

	namespaces := make([]string, 0, len(namespaceSet))
	for ns := range namespaceSet {
		namespaces = append(namespaces, ns)
	}
	sort.Strings(namespaces)
	artifact.Namespaces = namespaces

	registrations := make([]graph.EventRegistration, 0, len(parsed.Events))
	registrations = append(registrations, parsed.Events...)
	sort.Slice(registrations, func(i, j int) bool {
		if registrations[i].Line == registrations[j].Line {
			if registrations[i].Registrar == registrations[j].Registrar {
				return registrations[i].Event < registrations[j].Event
			}
			return registrations[i].Registrar < registrations[j].Registrar
		}
		return registrations[i].Line < registrations[j].Line
	})

	for index, reg := range registrations {
		artifact.Registrations = append(artifact.Registrations, luaAnalysisRegistration{
			RegistrationID:   fmt.Sprintf("r%d", index+1),
			SourceFunctionID: functionIDByName[reg.SourceFunction],
			Registrar:        reg.Registrar,
			EventRaw:         defaultString(reg.EventRaw, reg.Event),
			EventResolved:    defaultString(reg.EventResolved, reg.Event),
			HandlerRaw:       defaultString(reg.HandlerRaw, reg.Handler),
			HandlerResolved:  defaultString(reg.HandlerResolved, reg.Handler),
			Scope:            reg.Scope,
			WindowRaw:        reg.WindowRaw,
			WindowResolved:   defaultString(reg.WindowResolved, reg.Window),
			Source: luaAnalysisSource{
				Line: reg.Line,
			},
		})
	}

	return artifact
}

func normalizeRelativeLuaPath(path string) string {
	normalized := filepath.ToSlash(strings.TrimSpace(path))
	normalized = strings.TrimPrefix(normalized, "./")
	for strings.HasPrefix(normalized, "../") {
		normalized = strings.TrimPrefix(normalized, "../")
	}
	return normalized
}

func nonNilStrings(values []string) []string {
	if values == nil {
		return []string{}
	}
	return values
}

func defaultString(value string, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}

func scopeKindFromBool(local bool) string {
	if local {
		return "local"
	}
	return "global"
}

func isLuaUIAPI(name string) bool {
	if windowAPICallsForLuaContract[name] {
		return true
	}
	for prefix := range windowFunctionPrefixSetForLuaContract {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

func isLuaEventAPI(name string) bool {
	return knownEventAPIsForLuaContract[name]
}

var knownEventAPIsForLuaContract = map[string]bool{
	"RegisterEventHandler":           true,
	"UnregisterEventHandler":         true,
	"BroadcastEvent":                 true,
	"WindowRegisterEventHandler":     true,
	"WindowRegisterCoreEventHandler": true,
}

var windowAPICallsForLuaContract = map[string]bool{
	"WindowSetShowing":            true,
	"WindowGetShowing":            true,
	"WindowSetAlpha":              true,
	"WindowGetAlpha":              true,
	"WindowSetDimensions":         true,
	"WindowGetDimensions":         true,
	"WindowSetTintColor":          true,
	"WindowGetTintColor":          true,
	"WindowGetId":                 true,
	"WindowSetId":                 true,
	"WindowSetScale":              true,
	"WindowSetLayer":              true,
	"WindowSetHandleInput":        true,
	"WindowSetParent":             true,
	"WindowGetParent":             true,
	"WindowGetScreenPosition":     true,
	"WindowClearAnchors":          true,
	"WindowAddAnchor":             true,
	"WindowForceProcessAnchors":   true,
	"LabelSetText":                true,
	"LabelGetText":                true,
	"ButtonSetText":               true,
	"ButtonGetText":               true,
	"DynamicImageSetTexture":      true,
	"TextEditBoxGetText":          true,
	"TextEditBoxSetText":          true,
	"ComboBoxSetSelectedIndex":    true,
	"ComboBoxGetSelectedIndex":    true,
	"ScrollWindowSetOffset":       true,
	"ScrollWindowGetOffset":       true,
	"StatusBarSetCurrentValue":    true,
	"StatusBarSetMaximumValue":    true,
	"SliderBarSetCurrentPosition": true,
	"ListBoxSetDisplay":           true,
	"ListBoxGetItemCount":         true,
}

var windowFunctionPrefixSetForLuaContract = map[string]bool{
	"Window":       true,
	"Label":        true,
	"Button":       true,
	"DynamicImage": true,
	"TextEditBox":  true,
	"ComboBox":     true,
	"ScrollWindow": true,
	"StatusBar":    true,
	"SliderBar":    true,
	"ListBox":      true,
}

// orderedJSONMap is retained to ensure object-key stability for future map fields.
type orderedJSONMap map[string]string

func (m orderedJSONMap) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("{}"), nil
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var b bytes.Buffer
	b.WriteByte('{')
	for i, key := range keys {
		if i > 0 {
			b.WriteByte(',')
		}
		encodedKey, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		encodedValue, err := json.Marshal(m[key])
		if err != nil {
			return nil, err
		}
		b.Write(encodedKey)
		b.WriteByte(':')
		b.Write(encodedValue)
	}
	b.WriteByte('}')
	return b.Bytes(), nil
}

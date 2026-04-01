package graph

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var skippedSourceDirectories = map[string]struct{}{
	"api_ref":     {},
	"war_api_ref": {},
	"docs":        {},
	"data":        {},
	"infra":       {},
	"tools":       {},
}

type moduleFileXML struct {
	UIMod uiModXML `xml:"UiMod"`
}

type uiModXML struct {
	Name          string             `xml:"name,attr"`
	Version       string             `xml:"version,attr"`
	Date          string             `xml:"date,attr"`
	Author        authorXML          `xml:"Author"`
	Description   descriptionXML     `xml:"Description"`
	Files         []namedFileXML     `xml:"Files>File"`
	SavedVars     []savedVariableXML `xml:"SavedVariables>SavedVariable"`
	CreateWindows []createWindowXML  `xml:"OnInitialize>CreateWindow"`
	InitCalls     []callFunctionXML  `xml:"OnInitialize>CallFunction"`
	UpdateCalls   []callFunctionXML  `xml:"OnUpdate>CallFunction"`
	ShutdownCalls []callFunctionXML  `xml:"OnShutdown>CallFunction"`
}

type authorXML struct {
	Name string `xml:"name,attr"`
}

type descriptionXML struct {
	Text string `xml:"text,attr"`
}

type namedFileXML struct {
	Name string `xml:"name,attr"`
}

type savedVariableXML struct {
	Name   string `xml:"name,attr"`
	Global string `xml:"global,attr"`
}

type createWindowXML struct {
	Name string `xml:"name,attr"`
	Show string `xml:"show,attr"`
}

type callFunctionXML struct {
	Name string `xml:"name,attr"`
}

func DiscoverAddons(sourceRoot string, filters []string) ([]AddonSpec, error) {
	entries, err := os.ReadDir(sourceRoot)
	if err != nil {
		return nil, fmt.Errorf("read source root: %w", err)
	}

	filterSet := map[string]bool{}
	for _, filter := range filters {
		filterSet[strings.ToLower(strings.TrimSpace(filter))] = true
	}

	addons := make([]AddonSpec, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, ".") || shouldSkipSourceDirectory(name) {
			continue
		}
		if len(filterSet) > 0 && !filterSet[strings.ToLower(name)] {
			continue
		}

		addonPath := filepath.Join(sourceRoot, name)
		manifest, ok, parseErr := discoverManifest(addonPath)
		if parseErr != nil {
			return nil, parseErr
		}
		if !ok {
			continue
		}

		addons = append(addons, AddonSpec{
			Name:     manifest.Name,
			Path:     addonPath,
			Manifest: manifest,
		})
	}

	sort.Slice(addons, func(i, j int) bool {
		return addons[i].Name < addons[j].Name
	})
	return addons, nil
}

func shouldSkipSourceDirectory(name string) bool {
	_, found := skippedSourceDirectories[strings.ToLower(name)]
	return found
}

func discoverManifest(addonPath string) (Manifest, bool, error) {
	modFiles, err := filepath.Glob(filepath.Join(addonPath, "*.mod"))
	if err != nil {
		return Manifest{}, false, fmt.Errorf("glob manifests: %w", err)
	}
	if len(modFiles) > 0 {
		manifest, err := ParseModManifest(modFiles[0])
		return manifest, true, err
	}

	tocFiles, err := filepath.Glob(filepath.Join(addonPath, "*.toc"))
	if err != nil {
		return Manifest{}, false, fmt.Errorf("glob toc manifests: %w", err)
	}
	if len(tocFiles) > 0 {
		manifest, err := ParseTOCManifest(tocFiles[0])
		return manifest, true, err
	}

	return Manifest{}, false, nil
}

func ParseModManifest(path string) (Manifest, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return Manifest{}, fmt.Errorf("read mod manifest %s: %w", path, err)
	}
	bytes = sanitizeXMLDeclaration(bytes)
	bytes = sanitizeXMLAmpersands(bytes)
	bytes = sanitizeXMLComments(bytes)

	var document moduleFileXML
	if err := xml.Unmarshal(bytes, &document); err != nil {
		manifest, fallbackErr := parseModManifestFallback(string(bytes), path)
		if fallbackErr != nil {
			return Manifest{}, fmt.Errorf("parse mod manifest %s: %w", path, err)
		}
		return manifest, nil
	}

	manifest := Manifest{
		Name:            strings.TrimSpace(document.UIMod.Name),
		Path:            NormalizePath(path),
		Type:            "mod",
		Version:         strings.TrimSpace(document.UIMod.Version),
		Files:           make([]string, 0, len(document.UIMod.Files)),
		SavedVariables:  make([]string, 0, len(document.UIMod.SavedVars)),
		CreateWindows:   make([]string, 0, len(document.UIMod.CreateWindows)),
		InitializeCalls: make([]string, 0, len(document.UIMod.InitCalls)),
		UpdateCalls:     make([]string, 0, len(document.UIMod.UpdateCalls)),
		ShutdownCalls:   make([]string, 0, len(document.UIMod.ShutdownCalls)),
		Metadata: map[string]string{
			"author":      strings.TrimSpace(document.UIMod.Author.Name),
			"description": strings.TrimSpace(document.UIMod.Description.Text),
			"date":        strings.TrimSpace(document.UIMod.Date),
		},
	}

	for _, fileEntry := range document.UIMod.Files {
		trimmed := strings.TrimSpace(fileEntry.Name)
		if trimmed == "" {
			continue
		}
		manifest.Files = append(manifest.Files, NormalizePath(trimmed))
	}
	for _, saved := range document.UIMod.SavedVars {
		trimmed := strings.TrimSpace(saved.Name)
		if trimmed == "" {
			continue
		}
		manifest.SavedVariables = append(manifest.SavedVariables, trimmed)
	}
	for _, action := range document.UIMod.CreateWindows {
		trimmed := strings.TrimSpace(action.Name)
		if trimmed == "" {
			continue
		}
		manifest.CreateWindows = append(manifest.CreateWindows, trimmed)
	}
	for _, action := range document.UIMod.InitCalls {
		trimmed := strings.TrimSpace(action.Name)
		if trimmed == "" {
			continue
		}
		manifest.InitializeCalls = append(manifest.InitializeCalls, trimmed)
	}
	for _, action := range document.UIMod.UpdateCalls {
		trimmed := strings.TrimSpace(action.Name)
		if trimmed == "" {
			continue
		}
		manifest.UpdateCalls = append(manifest.UpdateCalls, trimmed)
	}
	for _, action := range document.UIMod.ShutdownCalls {
		trimmed := strings.TrimSpace(action.Name)
		if trimmed == "" {
			continue
		}
		manifest.ShutdownCalls = append(manifest.ShutdownCalls, trimmed)
	}

	if manifest.Name == "" {
		manifest.Name = filepath.Base(filepath.Dir(path))
	}
	return manifest, nil
}

func ParseTOCManifest(path string) (Manifest, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return Manifest{}, fmt.Errorf("read toc manifest %s: %w", path, err)
	}

	manifest := Manifest{
		Name:            strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Path:            NormalizePath(path),
		Type:            "toc",
		Files:           []string{},
		SavedVariables:  []string{},
		InitializeCalls: []string{},
		UpdateCalls:     []string{},
		ShutdownCalls:   []string{},
		Metadata:        map[string]string{},
	}

	for _, rawLine := range strings.Split(string(bytes), "\n") {
		line := strings.TrimSpace(strings.TrimSuffix(rawLine, "\r"))
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "##") {
			content := strings.TrimSpace(strings.TrimPrefix(line, "##"))
			parts := strings.SplitN(content, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				manifest.Metadata[strings.ToLower(key)] = value
				if strings.EqualFold(key, "SavedVariables") || strings.EqualFold(key, "SavedVariablesPerCharacter") {
					for _, token := range strings.Split(value, ",") {
						trimmed := strings.TrimSpace(token)
						if trimmed != "" {
							manifest.SavedVariables = append(manifest.SavedVariables, trimmed)
						}
					}
				}
			}
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		manifest.Files = append(manifest.Files, NormalizePath(line))
	}

	if version := strings.TrimSpace(manifest.Metadata["version"]); version != "" {
		manifest.Version = version
	}
	return manifest, nil
}

func sanitizeXMLDeclaration(content []byte) []byte {
	trimmed := strings.TrimLeft(string(content), "\ufeff \t\r\n")
	if strings.HasPrefix(trimmed, "<? version") {
		trimmed = strings.Replace(trimmed, "<? version", "<?xml version", 1)
		return []byte(trimmed)
	}
	return content
}

func sanitizeXMLAmpersands(content []byte) []byte {
	input := string(content)
	var builder strings.Builder
	builder.Grow(len(input))
	for index := 0; index < len(input); index++ {
		if input[index] != '&' {
			builder.WriteByte(input[index])
			continue
		}
		if looksLikeEntity(input[index:]) {
			builder.WriteByte('&')
			continue
		}
		builder.WriteString("&amp;")
	}
	return []byte(builder.String())
}

func looksLikeEntity(value string) bool {
	if len(value) < 4 || value[0] != '&' {
		return false
	}
	semicolon := strings.IndexByte(value, ';')
	if semicolon < 0 || semicolon > 12 {
		return false
	}
	body := value[1:semicolon]
	if body == "" {
		return false
	}
	if strings.HasPrefix(body, "#x") || strings.HasPrefix(body, "#X") {
		if len(body) == 2 {
			return false
		}
		for _, r := range body[2:] {
			if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
				return false
			}
		}
		return true
	}
	if strings.HasPrefix(body, "#") {
		if len(body) == 1 {
			return false
		}
		for _, r := range body[1:] {
			if r < '0' || r > '9' {
				return false
			}
		}
		return true
	}
	for _, r := range body {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

func sanitizeXMLComments(content []byte) []byte {
	result := make([]byte, 0, len(content))
	for index := 0; index < len(content); {
		if hasPrefixAt(content, index, "<!--") {
			result = append(result, '<', '!', '-', '-')
			index += 4
			for index < len(content) {
				if hasPrefixAt(content, index, "-->") {
					result = append(result, '-', '-', '>')
					index += 3
					break
				}
				if hasPrefixAt(content, index, "--") {
					result = append(result, '_', '_')
					index += 2
					continue
				}
				result = append(result, content[index])
				index++
			}
			continue
		}
		result = append(result, content[index])
		index++
	}
	return result
}

func hasPrefixAt(content []byte, index int, prefix string) bool {
	if index+len(prefix) > len(content) {
		return false
	}
	for offset := 0; offset < len(prefix); offset++ {
		if content[index+offset] != prefix[offset] {
			return false
		}
	}
	return true
}

var attrPattern = regexp.MustCompile(`([A-Za-z_:][A-Za-z0-9_:\-]*)\s*=\s*"([^"]*)"`)

func parseModManifestFallback(content string, path string) (Manifest, error) {
	uiModTag := findTag(content, "UiMod")
	attrs := parseAttributes(uiModTag)
	manifest := Manifest{
		Name:            strings.TrimSpace(attrs["name"]),
		Path:            NormalizePath(path),
		Type:            "mod",
		Version:         strings.TrimSpace(attrs["version"]),
		Files:           findAttributeValues(content, "File", "name"),
		SavedVariables:  findAttributeValues(content, "SavedVariable", "name"),
		CreateWindows:   findAttributeValues(findSection(content, "OnInitialize"), "CreateWindow", "name"),
		InitializeCalls: findAttributeValues(findSection(content, "OnInitialize"), "CallFunction", "name"),
		UpdateCalls:     findAttributeValues(findSection(content, "OnUpdate"), "CallFunction", "name"),
		ShutdownCalls:   findAttributeValues(findSection(content, "OnShutdown"), "CallFunction", "name"),
		Metadata: map[string]string{
			"author":      strings.TrimSpace(parseAttributes(findTag(content, "Author"))["name"]),
			"description": strings.TrimSpace(parseAttributes(findTag(content, "Description"))["text"]),
			"date":        strings.TrimSpace(attrs["date"]),
		},
	}
	if manifest.Name == "" {
		manifest.Name = filepath.Base(filepath.Dir(path))
	}
	if len(manifest.Files) == 0 && len(manifest.InitializeCalls) == 0 && len(manifest.CreateWindows) == 0 {
		return Manifest{}, fmt.Errorf("fallback parser could not recover manifest shape")
	}
	return manifest, nil
}

func findTag(content string, tagName string) string {
	open := "<" + tagName
	start := strings.Index(content, open)
	if start < 0 {
		return ""
	}
	end := strings.Index(content[start:], ">")
	if end < 0 {
		return ""
	}
	return content[start : start+end+1]
}

func findSection(content string, tagName string) string {
	open := "<" + tagName
	start := strings.Index(content, open)
	if start < 0 {
		return ""
	}
	openEnd := strings.Index(content[start:], ">")
	if openEnd < 0 {
		return ""
	}
	openText := content[start : start+openEnd+1]
	if strings.HasSuffix(strings.TrimSpace(openText), "/>") {
		return openText
	}
	close := "</" + tagName + ">"
	end := strings.Index(content[start+openEnd+1:], close)
	if end < 0 {
		return content[start:]
	}
	return content[start : start+openEnd+1+end+len(close)]
}

func parseAttributes(tag string) map[string]string {
	result := map[string]string{}
	for _, match := range attrPattern.FindAllStringSubmatch(tag, -1) {
		if len(match) == 3 {
			result[match[1]] = match[2]
		}
	}
	return result
}

func findAttributeValues(content string, tagName string, attrName string) []string {
	pattern := regexp.MustCompile(`(?s)<` + regexp.QuoteMeta(tagName) + `\b([^>]*)>`)
	values := []string{}
	for _, match := range pattern.FindAllStringSubmatch(content, -1) {
		if len(match) != 2 {
			continue
		}
		attributes := parseAttributes(match[1])
		if value := strings.TrimSpace(attributes[attrName]); value != "" {
			values = append(values, NormalizePath(value))
		}
	}
	return UniqueStrings(values)
}

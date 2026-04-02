package xml_parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/beevik/etree"
	"roraddons/tools/api_doc_gen/graph"
)

type elementContext struct {
	tag       string
	frameName string
}

var ignoredFrameTags = map[string]bool{
	"Interface":     true,
	"Assets":        true,
	"Windows":       true,
	"Anchors":       true,
	"Anchor":        true,
	"Size":          true,
	"Sizes":         true,
	"AbsPoint":      true,
	"Color":         true,
	"TexDims":       true,
	"TexCoords":     true,
	"TopLeft":       true,
	"TopCenter":     true,
	"TopRight":      true,
	"MiddleLeft":    true,
	"MiddleCenter":  true,
	"MiddleRight":   true,
	"BottomLeft":    true,
	"BottomCenter":  true,
	"BottomRight":   true,
	"EventHandlers": true,
	"EventHandler":  true,
	"Texture":       true,
	"Font":          true,
	"Layer":         true,
}

func ParseFile(addonName string, filePath string) (graph.XMLFileResult, error) {
	bytesContent, err := os.ReadFile(filePath)
	if err != nil {
		return graph.XMLFileResult{}, fmt.Errorf("read xml file %s: %w", filePath, err)
	}
	bytesContent = sanitizeXMLDeclaration(bytesContent)
	bytesContent = sanitizeXMLAmpersands(bytesContent)
	bytesContent = sanitizeXMLComments(bytesContent)

	decoder := xml.NewDecoder(bytes.NewReader(bytesContent))
	lineOffsets := buildLineOffsets(bytesContent)
	stack := []elementContext{}
	frames := []graph.Frame{}
	handlers := []graph.XMLHandler{}
	frameIndexes := map[string]int{}
	normalizedPath := graph.NormalizePath(filePath)

	for {
		token, err := decoder.RawToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			return graph.XMLFileResult{}, fmt.Errorf("parse xml file %s: %w", filePath, err)
		}

		switch typed := token.(type) {
		case xml.StartElement:
			line := lineNumberForOffset(lineOffsets, int(decoder.InputOffset()))
			parentFrame := nearestFrame(stack)
			attributes := attributesMap(typed.Attr)
			name := expandParentReference(attributes["name"], parentFrame)
			if name != "" {
				attributes["name"] = name
			}
			if value, ok := attributes["relativeTo"]; ok {
				attributes["relativeTo"] = expandParentReference(value, parentFrame)
			}
			if value, ok := attributes["inherits"]; ok {
				attributes["inherits"] = expandParentReference(value, parentFrame)
			}

			currentFrame := ""
			if shouldCaptureFrame(typed.Name.Local, attributes) {
				currentFrame = attributes["name"]
				frame := graph.Frame{
					Addon:      addonName,
					Name:       currentFrame,
					Type:       typed.Name.Local,
					Parent:     parentFrame,
					File:       normalizedPath,
					Line:       line,
					Children:   []string{},
					Inherits:   attributes["inherits"],
					Attributes: attributes,
					Template:   isTemplateFrame(currentFrame, attributes["inherits"]),
				}
				frameIndexes[currentFrame] = len(frames)
				frames = append(frames, frame)
				if parentFrame != "" {
					if parentIndex, ok := frameIndexes[parentFrame]; ok {
						frames[parentIndex].Children = append(frames[parentIndex].Children, currentFrame)
					}
				}
			} else if !ignoredFrameTags[typed.Name.Local] && strings.TrimSpace(attributes["name"]) == "" && parentFrame != "" {
				// Unnamed structural element inside a named frame (e.g. ListData, ListColumns, ListColumn).
				// Track its type and attribute keys so element-type documentation can surface sub-structure.
				if parentIndex, ok := frameIndexes[parentFrame]; ok {
					frames[parentIndex].StructuralChildTypes = append(frames[parentIndex].StructuralChildTypes, typed.Name.Local)
					// Collect attribute keys (not values) for this structural child occurrence.
					attrKeys := make([]string, 0, len(attributes))
					for k := range attributes {
						k = strings.TrimSpace(k)
						if k != "" && k != "name" {
							attrKeys = append(attrKeys, k)
						}
					}
					sort.Strings(attrKeys)
					if frames[parentIndex].StructuralChildAttrKeys == nil {
						frames[parentIndex].StructuralChildAttrKeys = map[string][]string{}
					}
					existing := frames[parentIndex].StructuralChildAttrKeys[typed.Name.Local]
					frames[parentIndex].StructuralChildAttrKeys[typed.Name.Local] = UniqueStringsSorted(append(existing, attrKeys...))
				}
			}

			if typed.Name.Local == "EventHandler" {
				frameName := parentFrame
				handlers = append(handlers, graph.XMLHandler{
					Addon:    addonName,
					Frame:    frameName,
					Event:    strings.TrimSpace(attributes["event"]),
					Function: strings.TrimSpace(attributes["function"]),
					File:     normalizedPath,
					Line:     line,
				})
			}

			stack = append(stack, elementContext{tag: typed.Name.Local, frameName: currentFrame})
		case xml.EndElement:
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}

	for index := range frames {
		frames[index].Children = graph.UniqueStrings(frames[index].Children)
		frames[index].StructuralChildTypes = graph.UniqueStrings(frames[index].StructuralChildTypes)
	}

	// Second pass: use etree to extract structural hierarchy (nesting of unnamed children).
	// This produces CompositionSnippet per named frame.
	snippets := extractCompositionSnippets(bytesContent)
	for index := range frames {
		if snippet, ok := snippets[frames[index].Name]; ok {
			frames[index].CompositionSnippet = snippet
		}
	}

	return graph.XMLFileResult{
		Frames:   graph.UniqueSortedFrames(frames),
		Handlers: graph.UniqueSortedHandlers(handlers),
	}, nil
}

func shouldCaptureFrame(tag string, attributes map[string]string) bool {
	if ignoredFrameTags[tag] {
		return false
	}
	if strings.TrimSpace(attributes["name"]) == "" {
		return false
	}
	return true
}

func isTemplateFrame(name string, inherits string) bool {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return false
	}
	if strings.Contains(trimmedName, "Template") {
		return true
	}
	trimmedInherits := strings.TrimSpace(inherits)
	return strings.Contains(trimmedInherits, "Template")
}

func nearestFrame(stack []elementContext) string {
	for index := len(stack) - 1; index >= 0; index-- {
		if stack[index].frameName != "" {
			return stack[index].frameName
		}
	}
	return ""
}

func attributesMap(attributes []xml.Attr) map[string]string {
	result := map[string]string{}
	for _, attribute := range attributes {
		result[attribute.Name.Local] = strings.TrimSpace(attribute.Value)
	}
	return result
}

func expandParentReference(value string, parent string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" || parent == "" {
		return trimmed
	}
	return strings.ReplaceAll(trimmed, "$parent", parent)
}

func buildLineOffsets(content []byte) []int {
	offsets := []int{0}
	for index, value := range content {
		if value == '\n' {
			offsets = append(offsets, index+1)
		}
	}
	return offsets
}

func lineNumberForOffset(offsets []int, offset int) int {
	index := sort.Search(len(offsets), func(position int) bool {
		return offsets[position] > offset
	})
	if index == 0 {
		return 1
	}
	return index
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


// UniqueStringsSorted returns a deduplicated, sorted copy of the input slice.
func UniqueStringsSorted(input []string) []string {
seen := map[string]bool{}
result := make([]string, 0, len(input))
for _, s := range input {
if s != "" && !seen[s] {
seen[s] = true
result = append(result, s)
}
}
sort.Strings(result)
return result
}

// ---------------------------------------------------------------------------
// etree-based structural composition extraction
// ---------------------------------------------------------------------------

// extractCompositionSnippets parses the XML file a second time using etree to
// build a proper in-memory element tree. For each named frame element it finds,
// it generates a CompositionSnippet — a compact XML representation that shows
// the hierarchy of unnamed structural children (e.g. ListBox → ListColumns →
// ListColumn), not just the flat list that the token parser produces.
//
// The result is keyed by frame name (value of the "name" attribute).
func extractCompositionSnippets(content []byte) map[string]string {
doc := etree.NewDocument()
if err := doc.ReadFromBytes(content); err != nil {
return map[string]string{}
}

result := map[string]string{}

// Find all elements that have a "name" attribute — these are the named frames.
for _, elem := range doc.FindElements("//*[@name]") {
name := strings.TrimSpace(elem.SelectAttrValue("name", ""))
if name == "" {
continue
}
// Only build snippets for element types we'd capture as frames.
if ignoredFrameTags[elem.Tag] {
continue
}
// Only bother when the element has structural (unnamed) children.
if !hasStructuralChildren(elem) {
continue
}
snippet := buildCompositionSnippet(elem, 0)
if snippet != "" {
result[name] = snippet
}
}

return result
}

// hasStructuralChildren returns true when elem contains at least one child
// element that is not in ignoredFrameTags and has no name attribute.
func hasStructuralChildren(elem *etree.Element) bool {
for _, child := range elem.ChildElements() {
if ignoredFrameTags[child.Tag] {
continue
}
if strings.TrimSpace(child.SelectAttrValue("name", "")) != "" {
continue // named child = separate frame
}
return true
}
return false
}

// buildCompositionSnippet recursively renders elem and its unnamed structural
// children as a compact XML snippet, indented by depth.
func buildCompositionSnippet(elem *etree.Element, depth int) string {
indent := strings.Repeat("  ", depth)

// Collect representative attributes for the outer element.
attrStr := compositionAttrs(elem, depth == 0)
tag := elem.Tag

// Collect structural (unnamed, non-ignored) children.
var childLines []string
for _, child := range elem.ChildElements() {
if ignoredFrameTags[child.Tag] {
continue
}
// Named children are separate documented frames — skip them here.
if strings.TrimSpace(child.SelectAttrValue("name", "")) != "" {
continue
}
childLines = append(childLines, buildCompositionSnippet(child, depth+1))
}

if len(childLines) == 0 {
return indent + "<" + tag + attrStr + "/>"
}
var b strings.Builder
b.WriteString(indent + "<" + tag + attrStr + ">")
for _, cl := range childLines {
b.WriteString("\n")
b.WriteString(cl)
}
b.WriteString("\n" + indent + "</" + tag + ">")
return b.String()
}

// compositionAttrs builds a concise attribute string for a snippet element.
// For the root element (isRoot=true) we use placeholder values; for inner
// elements we use the real attribute values from the parsed XML.
func compositionAttrs(elem *etree.Element, isRoot bool) string {
var parts []string
for _, attr := range elem.Attr {
key := attr.Key
if key == "name" {
if isRoot {
parts = append(parts, `name="..."`)
}
continue
}
val := strings.TrimSpace(attr.Value)
if val == "" {
continue
}
// Truncate long values so the snippet stays readable.
if len(val) > 32 {
val = val[:29] + "..."
}
parts = append(parts, key+`="`+val+`"`)
}
if len(parts) == 0 {
return ""
}
return " " + strings.Join(parts, " ")
}

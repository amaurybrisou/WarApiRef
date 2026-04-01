package xml_parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

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

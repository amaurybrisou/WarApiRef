package xml_structure

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/beevik/etree"
)

// ignoredTags contains element tags that are layout/anchoring primitives
// and should be marked as non-semantic (IsIgnored = true).
var ignoredTags = map[string]bool{
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

// maxSampleValues limits the number of distinct sample values stored per attribute.
const maxSampleValues = 8

// ParseFile parses a single XML file into an XMLTree using a real tree approach.
// It uses beevik/etree for proper tree construction, then extracts all structural
// metadata including parent/child relationships, handlers, and attribute data.
func ParseFile(addonName, filePath string) (*XMLTree, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	content = sanitizeXML(content)
	normalizedPath := normalizePath(filePath)

	tree := &XMLTree{
		Addon: addonName,
		File:  normalizedPath,
	}

	// Primary parse: build real tree with etree for proper hierarchy
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(content); err != nil {
		// Fallback: use token-based parsing if etree fails
		return parseWithTokens(addonName, normalizedPath, content)
	}

	// Walk the etree and build our XMLElement tree
	for _, child := range doc.ChildElements() {
		elem := buildElementTree(child, nil, addonName, normalizedPath)
		tree.Root = append(tree.Root, elem)
	}

	// Build flat index
	tree.AllNodes = collectAllNodes(tree.Root)

	// Resolve parent frame references and handlers
	resolveFrameContext(tree)

	return tree, nil
}

// buildElementTree recursively converts an etree.Element into our XMLElement tree.
func buildElementTree(etreeElem *etree.Element, parent *XMLElement, addon, file string) *XMLElement {
	attrs := make(map[string]string, len(etreeElem.Attr))
	for _, a := range etreeElem.Attr {
		attrs[a.Key] = strings.TrimSpace(a.Value)
	}

	// Determine nearest named ancestor for $parent expansion
	parentFrameName := ""
	parentFrameType := ""
	for p := parent; p != nil; p = p.Parent {
		if p.IsNamed {
			parentFrameName = p.Name
			parentFrameType = p.Tag
			break
		}
	}

	rawName := attrs["name"]
	resolvedName := expandParentRef(rawName, parentFrameName)

	isIgnored := ignoredTags[etreeElem.Tag]
	isNamed := strings.TrimSpace(resolvedName) != "" && !isIgnored

	elem := &XMLElement{
		Tag:             etreeElem.Tag,
		Name:            resolvedName,
		RawName:         rawName,
		Addon:           addon,
		File:            file,
		Line:            0, // etree doesn't expose line numbers; we'll fill from token pass
		Parent:          parent,
		Attributes:      attrs,
		IsNamed:         isNamed,
		IsStructural:    !isNamed && !isIgnored && parentFrameName != "",
		IsTemplate:      isTemplateElement(resolvedName, attrs["inherits"]),
		IsIgnored:       isIgnored,
		Inherits:        expandParentRef(attrs["inherits"], parentFrameName),
		ParentFrameName: parentFrameName,
		ParentFrameType: parentFrameType,
	}

	// Process children
	for _, childEtree := range etreeElem.ChildElements() {
		child := buildElementTree(childEtree, elem, addon, file)
		elem.Children = append(elem.Children, child)
	}

	return elem
}

// resolveFrameContext does a second pass to extract EventHandler declarations
// and assign them to their owning named elements.
// Note: etree does not expose source line numbers, so Line fields remain 0
// for etree-based parses. Line numbers are only available when the fallback
// token-based parser is used.
func resolveFrameContext(tree *XMLTree) {
	// Walk all nodes and collect EventHandler elements. Handlers can appear
	// either as direct children of a frame element or inside an <EventHandlers>
	// container. We process only <EventHandlers> containers (which is the
	// canonical XML pattern) to avoid double-counting.
	for _, node := range tree.AllNodes {
		if node.Tag != "EventHandlers" {
			continue
		}
		// Find the owning frame: the nearest named ancestor of the EventHandlers container.
		target := findNearestNamedAncestor(node)
		if target == nil {
			continue
		}
		for _, handlerElem := range node.Children {
			if handlerElem.Tag != "EventHandler" {
				continue
			}
			handler := XMLHandlerDecl{
				Event:    strings.TrimSpace(handlerElem.Attributes["event"]),
				Function: strings.TrimSpace(handlerElem.Attributes["function"]),
				Line:     handlerElem.Line,
			}
			if handler.Event != "" {
				target.Handlers = append(target.Handlers, handler)
			}
		}
	}
}

// findNearestNamedAncestor walks up the parent chain to find the nearest named element.
func findNearestNamedAncestor(elem *XMLElement) *XMLElement {
	for p := elem; p != nil; p = p.Parent {
		if p.IsNamed {
			return p
		}
	}
	return nil
}

// collectAllNodes returns all XMLElements in the tree in document order.
func collectAllNodes(roots []*XMLElement) []*XMLElement {
	var result []*XMLElement
	var walk func(elem *XMLElement)
	walk = func(elem *XMLElement) {
		result = append(result, elem)
		for _, child := range elem.Children {
			walk(child)
		}
	}
	for _, root := range roots {
		walk(root)
	}
	return result
}

// parseWithTokens is the fallback parser when etree fails. It uses the Go
// stdlib xml.Decoder with RawToken() similar to the original xml_parser.
func parseWithTokens(addon, file string, content []byte) (*XMLTree, error) {
	decoder := xml.NewDecoder(bytes.NewReader(content))
	lineOffsets := buildLineOffsets(content)

	tree := &XMLTree{Addon: addon, File: file}
	var stack []*XMLElement

	for {
		token, err := decoder.RawToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			break // Best effort
		}

		switch typed := token.(type) {
		case xml.StartElement:
			line := lineNumberForOffset(lineOffsets, int(decoder.InputOffset()))
			attrs := make(map[string]string, len(typed.Attr))
			for _, a := range typed.Attr {
				attrs[a.Name.Local] = strings.TrimSpace(a.Value)
			}

			// Find parent frame context
			parentFrameName := ""
			parentFrameType := ""
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i].IsNamed {
					parentFrameName = stack[i].Name
					parentFrameType = stack[i].Tag
					break
				}
			}

			rawName := attrs["name"]
			resolvedName := expandParentRef(rawName, parentFrameName)
			isIgnored := ignoredTags[typed.Name.Local]
			isNamed := strings.TrimSpace(resolvedName) != "" && !isIgnored

			var parent *XMLElement
			if len(stack) > 0 {
				parent = stack[len(stack)-1]
			}

			elem := &XMLElement{
				Tag:             typed.Name.Local,
				Name:            resolvedName,
				RawName:         rawName,
				Addon:           addon,
				File:            file,
				Line:            line,
				Parent:          parent,
				Attributes:      attrs,
				IsNamed:         isNamed,
				IsStructural:    !isNamed && !isIgnored && parentFrameName != "",
				IsTemplate:      isTemplateElement(resolvedName, attrs["inherits"]),
				IsIgnored:       isIgnored,
				Inherits:        expandParentRef(attrs["inherits"], parentFrameName),
				ParentFrameName: parentFrameName,
				ParentFrameType: parentFrameType,
			}

			if parent != nil {
				parent.Children = append(parent.Children, elem)
			} else {
				tree.Root = append(tree.Root, elem)
			}

			stack = append(stack, elem)

		case xml.EndElement:
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}

	tree.AllNodes = collectAllNodes(tree.Root)
	resolveFrameContext(tree)
	return tree, nil
}

// BuildCorpus aggregates multiple XMLTrees into an XMLCorpus with indexes
// and element type profiles.
func BuildCorpus(trees []*XMLTree) *XMLCorpus {
	corpus := &XMLCorpus{
		Trees:        trees,
		ByName:       make(map[string][]*XMLElement),
		ByTag:        make(map[string][]*XMLElement),
		ByAddon:      make(map[string][]*XMLTree),
		ByFile:       make(map[string]*XMLTree),
		ElementTypes: make(map[string]*ElementTypeProfile),
	}

	for _, tree := range trees {
		corpus.ByAddon[tree.Addon] = append(corpus.ByAddon[tree.Addon], tree)
		corpus.ByFile[tree.File] = tree

		for _, node := range tree.AllNodes {
			if node.IsIgnored {
				continue
			}

			// Index by name
			if node.IsNamed {
				corpus.ByName[node.Name] = append(corpus.ByName[node.Name], node)
			}

			// Index by tag
			corpus.ByTag[node.Tag] = append(corpus.ByTag[node.Tag], node)

			// Build element type profiles
			profile := corpus.ensureProfile(node.Tag)
			profile.TotalCount++
			if node.IsNamed {
				profile.NamedCount++
			} else {
				profile.UnnamedCount++
			}
			profile.ByAddon[node.Addon]++

			// Aggregate attributes
			for key, val := range node.Attributes {
				if key == "name" {
					continue
				}
				attrProfile := profile.ensureAttribute(key)
				attrProfile.TotalCount++
				attrProfile.ByAddon[node.Addon]++
				if len(attrProfile.SampleValues) < maxSampleValues {
					if !containsString(attrProfile.SampleValues, val) && val != "" {
						attrProfile.SampleValues = append(attrProfile.SampleValues, val)
					}
				}
			}

			// Structural relationships
			if node.Parent != nil && !node.Parent.IsIgnored {
				profile.ParentTags[node.Parent.Tag]++
			}
			for _, child := range node.Children {
				if child.IsIgnored {
					continue
				}
				profile.ChildTags[child.Tag]++
				if child.IsNamed {
					profile.NamedChildren[child.Tag]++
				} else if child.IsStructural {
					profile.StructuralChildren[child.Tag]++
				}
			}

			// Inheritance
			if node.Inherits != "" {
				profile.InheritsBases[node.Inherits]++
			}
			if node.IsTemplate {
				profile.IsTemplate = true
			}

			// Handler events
			for _, h := range node.Handlers {
				profile.HandlerEvents[h.Event]++
				if h.Function != "" {
					if profile.HandlerBindings[h.Event] == nil {
						profile.HandlerBindings[h.Event] = make(map[string]int)
					}
					profile.HandlerBindings[h.Event][h.Function]++
				}
			}
		}
	}

	// Compute attribute required flags
	for _, profile := range corpus.ElementTypes {
		for _, attr := range profile.Attributes {
			if profile.TotalCount > 0 {
				ratio := float64(attr.TotalCount) / float64(profile.TotalCount)
				attr.IsRequired = ratio > 0.8
			}
		}
	}

	return corpus
}

// ensureProfile returns or creates an ElementTypeProfile for the given tag.
func (c *XMLCorpus) ensureProfile(tag string) *ElementTypeProfile {
	if p, ok := c.ElementTypes[tag]; ok {
		return p
	}
	p := &ElementTypeProfile{
		Tag:                tag,
		ByAddon:            make(map[string]int),
		Attributes:         make(map[string]*AttributeProfile),
		ParentTags:         make(map[string]int),
		ChildTags:          make(map[string]int),
		NamedChildren:      make(map[string]int),
		StructuralChildren: make(map[string]int),
		InheritsBases:      make(map[string]int),
		HandlerEvents:      make(map[string]int),
		HandlerBindings:    make(map[string]map[string]int),
	}
	c.ElementTypes[tag] = p
	return p
}

// ensureAttribute returns or creates an AttributeProfile for the given key.
func (p *ElementTypeProfile) ensureAttribute(name string) *AttributeProfile {
	if a, ok := p.Attributes[name]; ok {
		return a
	}
	a := &AttributeProfile{
		Name:    name,
		ByAddon: make(map[string]int),
	}
	p.Attributes[name] = a
	return a
}

// CompositionSnippet builds a representative XML structure snippet for an element
// type using the etree-based approach.
func BuildCompositionSnippet(elem *XMLElement, maxDepth int) string {
	return buildSnippet(elem, 0, maxDepth)
}

func buildSnippet(elem *XMLElement, depth, maxDepth int) string {
	if depth > maxDepth {
		return ""
	}
	indent := strings.Repeat("  ", depth)
	tag := elem.Tag

	attrStr := snippetAttrs(elem, depth == 0)

	var childLines []string
	for _, child := range elem.Children {
		if child.IsIgnored {
			continue
		}
		if child.IsNamed {
			continue // Named children are separate documented frames
		}
		line := buildSnippet(child, depth+1, maxDepth)
		if line != "" {
			childLines = append(childLines, line)
		}
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

func snippetAttrs(elem *XMLElement, isRoot bool) string {
	var parts []string
	// Sort keys for deterministic output
	keys := make([]string, 0, len(elem.Attributes))
	for k := range elem.Attributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if key == "name" {
			if isRoot {
				parts = append(parts, `name="..."`)
			}
			continue
		}
		val := elem.Attributes[key]
		if val == "" {
			continue
		}
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

// -----------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------

func expandParentRef(value, parentName string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" || parentName == "" {
		return trimmed
	}
	return strings.ReplaceAll(trimmed, "$parent", parentName)
}

func isTemplateElement(name, inherits string) bool {
	return strings.Contains(name, "Template") || strings.Contains(inherits, "Template")
}

func normalizePath(p string) string {
	return strings.ReplaceAll(p, "\\", "/")
}

func containsString(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

func sanitizeXML(content []byte) []byte {
	content = sanitizeDeclaration(content)
	content = sanitizeAmpersands(content)
	content = sanitizeComments(content)
	return content
}

func sanitizeDeclaration(content []byte) []byte {
	trimmed := strings.TrimLeft(string(content), "\ufeff \t\r\n")
	if strings.HasPrefix(trimmed, "<? version") {
		trimmed = strings.Replace(trimmed, "<? version", "<?xml version", 1)
		return []byte(trimmed)
	}
	return content
}

func sanitizeAmpersands(content []byte) []byte {
	input := string(content)
	var b strings.Builder
	b.Grow(len(input))
	for i := 0; i < len(input); i++ {
		if input[i] != '&' {
			b.WriteByte(input[i])
			continue
		}
		if looksLikeEntity(input[i:]) {
			b.WriteByte('&')
			continue
		}
		b.WriteString("&amp;")
	}
	return []byte(b.String())
}

func looksLikeEntity(value string) bool {
	if len(value) < 4 || value[0] != '&' {
		return false
	}
	semi := strings.IndexByte(value, ';')
	if semi < 0 || semi > 12 {
		return false
	}
	body := value[1:semi]
	if body == "" {
		return false
	}
	if strings.HasPrefix(body, "#x") || strings.HasPrefix(body, "#X") {
		for _, r := range body[2:] {
			if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
				return false
			}
		}
		return len(body) > 2
	}
	if strings.HasPrefix(body, "#") {
		for _, r := range body[1:] {
			if r < '0' || r > '9' {
				return false
			}
		}
		return len(body) > 1
	}
	for _, r := range body {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

func sanitizeComments(content []byte) []byte {
	result := make([]byte, 0, len(content))
	for i := 0; i < len(content); {
		if i+4 <= len(content) && string(content[i:i+4]) == "<!--" {
			result = append(result, '<', '!', '-', '-')
			i += 4
			for i < len(content) {
				if i+3 <= len(content) && string(content[i:i+3]) == "-->" {
					result = append(result, '-', '-', '>')
					i += 3
					break
				}
				if i+2 <= len(content) && string(content[i:i+2]) == "--" {
					result = append(result, '_', '_')
					i += 2
					continue
				}
				result = append(result, content[i])
				i++
			}
			continue
		}
		result = append(result, content[i])
		i++
	}
	return result
}

func buildLineOffsets(content []byte) []int {
	offsets := []int{0}
	for i, b := range content {
		if b == '\n' {
			offsets = append(offsets, i+1)
		}
	}
	return offsets
}

func lineNumberForOffset(offsets []int, offset int) int {
	idx := sort.Search(len(offsets), func(i int) bool {
		return offsets[i] > offset
	})
	if idx == 0 {
		return 1
	}
	return idx
}

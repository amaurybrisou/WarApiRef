// Package xml_structure provides conversion helpers from XMLTree/XMLCorpus
// into graph-facing DTOs used by addon documentation generation.
package xml_structure

import (
	"sort"
	"strings"
)

// Frame is a minimal representation matching graph.Frame fields for bridging.
type Frame struct {
	Addon                   string
	Name                    string
	RawName                 string
	Type                    string
	Parent                  string
	ParentType              string
	File                    string
	Line                    int
	Children                []string
	ChildElementTypes       []string
	StructuralChildTypes    []string
	StructuralChildAttrKeys map[string][]string
	CompositionSnippet      string
	Inherits                string
	Attributes              map[string]string
	Template                bool
}

// Handler is a minimal representation matching graph.XMLHandler for bridging.
type Handler struct {
	Addon    string
	Frame    string
	Event    string
	Function string
	File     string
	Line     int
}

// ToFramesAndHandlers converts an XMLTree into flat Frame/Handler lists used
// by addon docs generation.
func (tree *XMLTree) ToFramesAndHandlers() ([]Frame, []Handler) {
	var frames []Frame
	var handlers []Handler

	frameIndex := make(map[string]int) // name → index in frames

	for _, node := range tree.AllNodes {
		if node.IsIgnored {
			continue
		}

		// Extract handlers first (they apply to named frames)
		for _, h := range node.Handlers {
			handlers = append(handlers, Handler{
				Addon:    tree.Addon,
				Frame:    node.Name,
				Event:    h.Event,
				Function: h.Function,
				File:     tree.File,
				Line:     h.Line,
			})
		}

		// Only named elements become Frames
		if !node.IsNamed {
			continue
		}

		frame := Frame{
			Addon:      tree.Addon,
			Name:       node.Name,
			RawName:    node.RawName,
			Type:       node.Tag,
			Parent:     node.ParentFrameName,
			ParentType: node.ParentFrameType,
			File:       tree.File,
			Line:       node.Line,
			Inherits:   node.Inherits,
			Attributes: node.Attributes,
			Template:   node.IsTemplate,
		}

		// Collect named children and structural children
		for _, child := range node.Children {
			if child.IsNamed {
				frame.Children = append(frame.Children, child.Name)
				frame.ChildElementTypes = append(frame.ChildElementTypes, child.Tag)
			} else if child.IsStructural {
				frame.StructuralChildTypes = append(frame.StructuralChildTypes, child.Tag)
				// Collect attribute keys for structural children
				attrKeys := make([]string, 0)
				for k := range child.Attributes {
					if k != "name" && strings.TrimSpace(k) != "" {
						attrKeys = append(attrKeys, k)
					}
				}
				sort.Strings(attrKeys)
				if frame.StructuralChildAttrKeys == nil {
					frame.StructuralChildAttrKeys = make(map[string][]string)
				}
				existing := frame.StructuralChildAttrKeys[child.Tag]
				frame.StructuralChildAttrKeys[child.Tag] = uniqueSorted(append(existing, attrKeys...))
			}
		}

		// Deduplicate
		frame.Children = uniqueSorted(frame.Children)
		frame.ChildElementTypes = uniqueSorted(frame.ChildElementTypes)
		frame.StructuralChildTypes = uniqueSorted(frame.StructuralChildTypes)

		// Build composition snippet for structural children
		if hasStructuralChildElements(node) {
			frame.CompositionSnippet = BuildCompositionSnippet(node, 4)
		}

		frameIndex[node.Name] = len(frames)
		frames = append(frames, frame)
	}

	return frames, handlers
}

// hasStructuralChildElements returns true if elem has at least one structural child.
func hasStructuralChildElements(elem *XMLElement) bool {
	for _, child := range elem.Children {
		if child.IsStructural {
			return true
		}
	}
	return false
}

func uniqueSorted(input []string) []string {
	if len(input) == 0 {
		return nil
	}
	seen := make(map[string]bool, len(input))
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

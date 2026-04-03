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
	"roraddons/tools/api_doc_gen/xml_structure"
)

const phase1TreeSchemaVersion = "phase1-tree/v1"

type phase1TreeArtifact struct {
	SchemaVersion string           `json:"schema_version"`
	Addon         string           `json:"addon"`
	SourceFile    string           `json:"source_file"`
	Roots         []string         `json:"roots"`
	Nodes         []phase1TreeNode `json:"nodes"`
}

type phase1TreeNode struct {
	NodeID       string                 `json:"node_id"`
	ParentID     *string                `json:"parent_id"`
	ChildIndex   int                    `json:"child_index"`
	Tag          string                 `json:"tag"`
	RawName      string                 `json:"raw_name"`
	ResolvedName string                 `json:"resolved_name"`
	Attributes   sortedStringMap        `json:"attributes"`
	Flags        phase1TreeFlags        `json:"flags"`
	Source       phase1TreeSource       `json:"source"`
	Ancestor     phase1TreeAncestorInfo `json:"ancestor"`
}

type phase1TreeFlags struct {
	IsNamed      bool `json:"is_named"`
	IsStructural bool `json:"is_structural"`
	IsIgnored    bool `json:"is_ignored"`
	IsTemplate   bool `json:"is_template"`
}

type phase1TreeSource struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

type phase1TreeAncestorInfo struct {
	NearestNamedName string `json:"nearest_named_name"`
	NearestNamedTag  string `json:"nearest_named_tag"`
}

// sortedStringMap ensures deterministic key ordering in JSON output.
type sortedStringMap map[string]string

func (m sortedStringMap) MarshalJSON() ([]byte, error) {
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

func writeXMLTreeArtifacts(outputRoot string, corpus graph.Corpus) error {
	for _, addon := range corpus.Addons {
		for _, manifestFile := range addon.Manifest.Files {
			ext := strings.ToLower(filepath.Ext(manifestFile))
			if ext != ".xml" {
				continue
			}

			normalizedRelative := normalizeRelativeXMLPath(manifestFile)
			absolutePath := filepath.Join(addon.Root, filepath.FromSlash(normalizedRelative))
			tree, err := xml_structure.ParseFile(addon.Name, absolutePath)
			if err != nil {
				return fmt.Errorf("parse xml tree %s/%s: %w", addon.Name, normalizedRelative, err)
			}

			artifact := buildPhase1TreeArtifact(addon.Name, normalizedRelative, tree)
			relativeArtifactPath := filepath.Join("xml-tree", addon.Name, filepath.FromSlash(normalizedRelative)+".tree.json")
			outputPath := filepath.Join(outputRoot, relativeArtifactPath)
			if err := writePhase1TreeArtifact(outputPath, artifact); err != nil {
				return err
			}
		}
	}
	return nil
}

func writePhase1TreeArtifact(outputPath string, artifact phase1TreeArtifact) error {
	content, err := json.MarshalIndent(artifact, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal xml tree artifact %s: %w", outputPath, err)
	}
	content = append(content, '\n')
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("create xml tree directory %s: %w", filepath.Dir(outputPath), err)
	}
	if err := os.WriteFile(outputPath, content, 0o644); err != nil {
		return fmt.Errorf("write xml tree artifact %s: %w", outputPath, err)
	}
	return nil
}

func buildPhase1TreeArtifact(addon, sourceFile string, tree *xml_structure.XMLTree) phase1TreeArtifact {
	artifact := phase1TreeArtifact{
		SchemaVersion: phase1TreeSchemaVersion,
		Addon:         addon,
		SourceFile:    normalizeRelativeXMLPath(sourceFile),
		Roots:         make([]string, 0, len(tree.Root)),
		Nodes:         make([]phase1TreeNode, 0, len(tree.AllNodes)),
	}

	counter := 0
	var walk func(node *xml_structure.XMLElement, parentID *string, childIndex int)
	walk = func(node *xml_structure.XMLElement, parentID *string, childIndex int) {
		counter++
		nodeID := fmt.Sprintf("n%d", counter)

		attributes := make(sortedStringMap, len(node.Attributes))
		for key, value := range node.Attributes {
			attributes[key] = value
		}

		nodeSourceFile := artifact.SourceFile
		if strings.TrimSpace(nodeSourceFile) == "" {
			nodeSourceFile = normalizeRelativeXMLPath(node.File)
		}

		artifact.Nodes = append(artifact.Nodes, phase1TreeNode{
			NodeID:       nodeID,
			ParentID:     parentID,
			ChildIndex:   childIndex,
			Tag:          node.Tag,
			RawName:      node.RawName,
			ResolvedName: node.Name,
			Attributes:   attributes,
			Flags: phase1TreeFlags{
				IsNamed:      node.IsNamed,
				IsStructural: node.IsStructural,
				IsIgnored:    node.IsIgnored,
				IsTemplate:   node.IsTemplate,
			},
			Source: phase1TreeSource{
				File: nodeSourceFile,
				Line: node.Line,
			},
			Ancestor: phase1TreeAncestorInfo{
				NearestNamedName: node.ParentFrameName,
				NearestNamedTag:  node.ParentFrameType,
			},
		})

		for i, child := range node.Children {
			parent := nodeID
			walk(child, &parent, i)
		}
	}

	for i, root := range tree.Root {
		preorderIndex := len(artifact.Nodes) + 1
		rootID := fmt.Sprintf("n%d", preorderIndex)
		artifact.Roots = append(artifact.Roots, rootID)
		walk(root, nil, i)
	}

	return artifact
}

func normalizeRelativeXMLPath(path string) string {
	normalized := filepath.ToSlash(strings.TrimSpace(path))
	normalized = strings.TrimPrefix(normalized, "./")
	for strings.HasPrefix(normalized, "../") {
		normalized = strings.TrimPrefix(normalized, "../")
	}
	return normalized
}

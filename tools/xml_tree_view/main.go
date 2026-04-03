package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type artifact struct {
	SchemaVersion string     `json:"schema_version"`
	Addon         string     `json:"addon"`
	SourceFile    string     `json:"source_file"`
	Roots         []string   `json:"roots"`
	Nodes         []treeNode `json:"nodes"`
}

type treeNode struct {
	NodeID       string            `json:"node_id"`
	ParentID     *string           `json:"parent_id"`
	ChildIndex   int               `json:"child_index"`
	Tag          string            `json:"tag"`
	RawName      string            `json:"raw_name"`
	ResolvedName string            `json:"resolved_name"`
	Attributes   map[string]string `json:"attributes"`
	Flags        nodeFlags         `json:"flags"`
}

type nodeFlags struct {
	IsNamed      bool `json:"is_named"`
	IsStructural bool `json:"is_structural"`
	IsIgnored    bool `json:"is_ignored"`
	IsTemplate   bool `json:"is_template"`
}

type indexedTree struct {
	byID     map[string]*treeNode
	children map[string][]*treeNode
}

func main() {
	maxDepth := flag.Int("max-depth", -1, "Maximum depth to render (-1 for unlimited)")
	flag.Parse()

	paths := flag.Args()
	if len(paths) == 0 {
		fmt.Fprintln(os.Stderr, "usage: go run ./tools/xml_tree_view --max-depth 4 <tree.json> [tree.json...]")
		os.Exit(2)
	}

	for i, path := range paths {
		if i > 0 {
			fmt.Println()
		}
		if err := renderArtifact(path, *maxDepth); err != nil {
			fmt.Fprintf(os.Stderr, "error rendering %s: %v\n", path, err)
			os.Exit(1)
		}
	}
}

func renderArtifact(path string, maxDepth int) error {
	parsed, err := readArtifact(path)
	if err != nil {
		return err
	}
	indexed := buildIndex(parsed)

	fmt.Printf("=== %s ===\n", filepath.ToSlash(path))
	fmt.Printf("schema=%s addon=%s source=%s roots=%d nodes=%d\n", parsed.SchemaVersion, parsed.Addon, parsed.SourceFile, len(parsed.Roots), len(parsed.Nodes))

	for i, rootID := range parsed.Roots {
		root := indexed.byID[rootID]
		if root == nil {
			return fmt.Errorf("missing root node %s", rootID)
		}
		lastRoot := i == len(parsed.Roots)-1
		renderNode(root, indexed, "", lastRoot, 0, maxDepth)
	}

	return nil
}

func readArtifact(path string) (*artifact, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var parsed artifact
	if err := json.Unmarshal(content, &parsed); err != nil {
		return nil, err
	}
	return &parsed, nil
}

func buildIndex(parsed *artifact) indexedTree {
	byID := make(map[string]*treeNode, len(parsed.Nodes))
	children := make(map[string][]*treeNode)

	for i := range parsed.Nodes {
		node := &parsed.Nodes[i]
		byID[node.NodeID] = node
		if node.ParentID != nil {
			parentID := *node.ParentID
			children[parentID] = append(children[parentID], node)
		}
	}

	for parentID := range children {
		sort.Slice(children[parentID], func(i, j int) bool {
			if children[parentID][i].ChildIndex == children[parentID][j].ChildIndex {
				return children[parentID][i].NodeID < children[parentID][j].NodeID
			}
			return children[parentID][i].ChildIndex < children[parentID][j].ChildIndex
		})
	}

	return indexedTree{byID: byID, children: children}
}

func renderNode(node *treeNode, indexed indexedTree, prefix string, isLast bool, depth int, maxDepth int) {
	branch := "|- "
	nextPrefix := prefix + "|  "
	if isLast {
		branch = "`- "
		nextPrefix = prefix + "   "
	}
	if depth == 0 {
		branch = ""
	}

	fmt.Printf("%s%s%s\n", prefix, branch, formatLabel(node))

	if maxDepth >= 0 && depth >= maxDepth {
		return
	}

	kids := indexed.children[node.NodeID]
	for i, child := range kids {
		renderNode(child, indexed, nextPrefix, i == len(kids)-1, depth+1, maxDepth)
	}
}

func formatLabel(node *treeNode) string {
	flags := make([]string, 0, 2)
	if node.Flags.IsNamed {
		flags = append(flags, "named")
	}
	if node.Flags.IsStructural {
		flags = append(flags, "structural")
	}
	if len(flags) == 0 {
		flags = append(flags, "plain")
	}

	namePart := node.ResolvedName
	if strings.TrimSpace(namePart) == "" {
		namePart = "-"
	}
	rawPart := ""
	if strings.TrimSpace(node.RawName) != "" && node.RawName != node.ResolvedName {
		rawPart = fmt.Sprintf(" raw=%s", node.RawName)
	}

	return fmt.Sprintf("%s name=%s%s [%s]", node.Tag, namePart, rawPart, strings.Join(flags, ","))
}

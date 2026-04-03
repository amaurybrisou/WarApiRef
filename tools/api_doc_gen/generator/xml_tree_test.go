package generator

import (
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"testing"

	"roraddons/tools/api_doc_gen/xml_structure"
)

func repoRoot(t *testing.T) string {
	t.Helper()
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to locate test file path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(currentFile), "..", "..", ".."))
}

func firstExistingFile(root string, candidates ...string) string {
	for _, candidate := range candidates {
		full := filepath.Join(root, filepath.FromSlash(candidate))
		if _, err := os.Stat(full); err == nil {
			return candidate
		}
	}
	return ""
}

func findNodeByID(nodes []phase1TreeNode, id string) *phase1TreeNode {
	for i := range nodes {
		if nodes[i].NodeID == id {
			return &nodes[i]
		}
	}
	return nil
}

func TestBuildPhase1TreeArtifactPartyCastHierarchy(t *testing.T) {
	root := repoRoot(t)
	sourceFile := firstExistingFile(root,
		"addons/PartyCast/PartyCast.xml",
		"all_addons/PartyCast/PartyCast.xml",
	)
	if sourceFile == "" {
		t.Skip("PartyCast XML not present in workspace")
	}

	tree, err := xml_structure.ParseFile("PartyCast", filepath.Join(root, filepath.FromSlash(sourceFile)))
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	artifact := buildPhase1TreeArtifact("PartyCast", sourceFile, tree)
	if artifact.SchemaVersion != phase1TreeSchemaVersion {
		t.Fatalf("unexpected schema version: %s", artifact.SchemaVersion)
	}
	wantSource := filepath.ToSlash(sourceFile)
	if artifact.SourceFile != wantSource {
		t.Fatalf("unexpected source file: %s", artifact.SourceFile)
	}
	if len(artifact.Roots) == 0 {
		t.Fatal("expected roots")
	}

	for i, node := range artifact.Nodes {
		want := "n" + strconv.Itoa(i+1)
		if node.NodeID != want {
			t.Fatalf("node id mismatch at index %d: got %s want %s", i, node.NodeID, want)
		}
	}

	var windowsID, lolID string
	for _, node := range artifact.Nodes {
		if node.Tag == "Windows" && node.ParentID != nil {
			parent := findNodeByID(artifact.Nodes, *node.ParentID)
			if parent != nil && parent.Tag == "Interface" {
				windowsID = node.NodeID
			}
		}
		if node.Tag == "HorizontalResizeImage" && node.ResolvedName == "LOL_BAR" {
			lolID = node.NodeID
		}
	}

	if windowsID == "" {
		t.Fatal("expected Windows node under Interface")
	}
	if lolID == "" {
		t.Fatal("expected LOL_BAR node")
	}

	lolNode := findNodeByID(artifact.Nodes, lolID)
	if lolNode.ParentID == nil || *lolNode.ParentID != windowsID {
		t.Fatalf("expected LOL_BAR parent %s, got %v", windowsID, lolNode.ParentID)
	}

	if lolNode.Attributes["texture"] != "EA_Training_Specialization" {
		t.Fatalf("missing LOL_BAR texture attribute: %v", lolNode.Attributes)
	}
	if lolNode.Attributes["textureScale"] != "0.89" {
		t.Fatalf("missing LOL_BAR textureScale attribute: %v", lolNode.Attributes)
	}

	type childRecord struct {
		tag        string
		childIndex int
	}
	children := make([]childRecord, 0)
	for _, node := range artifact.Nodes {
		if node.ParentID != nil && *node.ParentID == lolID {
			children = append(children, childRecord{tag: node.Tag, childIndex: node.ChildIndex})
		}
	}
	sort.Slice(children, func(i, j int) bool { return children[i].childIndex < children[j].childIndex })

	if len(children) < 3 {
		t.Fatalf("expected at least 3 children for LOL_BAR, got %v", children)
	}
	if children[0].tag != "Size" || children[1].tag != "Sizes" || children[2].tag != "TexCoords" {
		t.Fatalf("unexpected LOL_BAR child order: %v", children)
	}

	for _, node := range artifact.Nodes {
		if node.ParentID != nil && *node.ParentID == lolID {
			if !node.Flags.IsStructural {
				t.Fatalf("expected LOL_BAR child %s to be structural", node.Tag)
			}
			if node.Source.File != wantSource {
				t.Fatalf("expected node source file %s, got %s", wantSource, node.Source.File)
			}
		}
	}
}

func TestBuildPhase1TreeArtifactAdvancedRenownRawResolved(t *testing.T) {
	root := repoRoot(t)
	sourceFile := "addons/advancedrenowntrainer/AdvancedRenownTrainingPresets.xml"
	fullPath := filepath.Join(root, filepath.FromSlash(sourceFile))
	if _, err := os.Stat(fullPath); err != nil {
		t.Skip("AdvancedRenownTrainingPresets.xml not present in workspace")
	}

	tree, err := xml_structure.ParseFile("advancedrenowntrainer", fullPath)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	artifact := buildPhase1TreeArtifact("advancedrenowntrainer", sourceFile, tree)

	var parentID string
	for _, node := range artifact.Nodes {
		if node.Tag == "Window" && node.ResolvedName == "AdvancedRenownTrainingPresetsWindow" {
			parentID = node.NodeID
			break
		}
	}
	if parentID == "" {
		t.Fatal("expected AdvancedRenownTrainingPresetsWindow node")
	}

	var bg *phase1TreeNode
	for i := range artifact.Nodes {
		node := &artifact.Nodes[i]
		if node.Tag == "Window" && node.RawName == "$parentBackground" && node.Ancestor.NearestNamedName == "AdvancedRenownTrainingPresetsWindow" {
			bg = node
			break
		}
	}
	if bg == nil {
		t.Fatal("expected $parentBackground node")
	}
	if bg.ResolvedName != "AdvancedRenownTrainingPresetsWindowBackground" {
		t.Fatalf("unexpected resolved name: %s", bg.ResolvedName)
	}
	if bg.ParentID == nil {
		t.Fatal("expected parent id for $parentBackground node")
	}
	parentNode := findNodeByID(artifact.Nodes, *bg.ParentID)
	if parentNode == nil || parentNode.Tag != "Windows" {
		t.Fatalf("expected parent tag Windows, got %+v", parentNode)
	}
	if bg.Source.File != filepath.ToSlash(sourceFile) {
		t.Fatalf("unexpected node source file: %s", bg.Source.File)
	}
	if bg.Ancestor.NearestNamedTag != "Window" {
		t.Fatalf("unexpected nearest named tag: %s", bg.Ancestor.NearestNamedTag)
	}
}

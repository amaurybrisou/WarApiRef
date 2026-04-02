// Package source_scan provides source file discovery for addon XML and Lua files.
//
// This is the source-first entry point for the phased pipeline. Rather than
// consuming pre-flattened API-reference docs, callers use [DiscoverAddonSources]
// to obtain a concrete inventory of XML and Lua files that the pipeline can
// parse directly.
package source_scan

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"roraddons/tools/api_doc_gen/graph"
)

// AddonSource represents the raw source files of a single addon.
type AddonSource struct {
	// AddonName is the canonical addon identifier from the manifest.
	AddonName string

	// AddonPath is the root directory of the addon on disk.
	AddonPath string

	// XMLFiles contains the absolute paths to all XML source files listed in
	// the manifest and confirmed to exist on disk.
	XMLFiles []string

	// LuaFiles contains the absolute paths to all Lua source files listed in
	// the manifest and confirmed to exist on disk.
	LuaFiles []string

	// Manifest is the parsed addon manifest.
	Manifest graph.Manifest

	// ModuleTree is the full discovery-oriented tree of the .mod manifest.
	// It is non-nil only when the addon uses a .mod manifest (not a .toc
	// file).  The tree preserves every XML node and attribute in the manifest,
	// including sections that are not recognised by the current parser.
	ModuleTree *graph.ModuleTree
}

// DiscoverAddonSources enumerates all addon source files under sourceRoot.
//
// It uses the addon manifests (*.mod or *.toc files) found inside each
// sub-directory of sourceRoot to resolve XML and Lua file paths. Only files
// that are both listed in the manifest and present on disk are included.
//
// The returned slice is sorted by addon name.
func DiscoverAddonSources(sourceRoot string) ([]AddonSource, error) {
	specs, err := graph.DiscoverAddons(sourceRoot, nil)
	if err != nil {
		return nil, fmt.Errorf("discover addon sources: %w", err)
	}

	sources := make([]AddonSource, 0, len(specs))
	for _, spec := range specs {
		src := resolveAddonSource(spec)
		if len(src.XMLFiles)+len(src.LuaFiles) > 0 {
			sources = append(sources, src)
		}
	}
	return sources, nil
}

// resolveAddonSource builds an AddonSource from a discovered AddonSpec by
// resolving each manifest file entry to an absolute path and filtering by
// extension. Files that do not exist on disk are silently skipped.
//
// When the manifest is a .mod file, the full discovery-oriented ModuleTree is
// also parsed and attached to the returned AddonSource so that downstream
// pipeline stages can inspect any section of the manifest without loss.
func resolveAddonSource(spec graph.AddonSpec) AddonSource {
	src := AddonSource{
		AddonName: spec.Name,
		AddonPath: spec.Path,
		Manifest:  spec.Manifest,
	}

	if spec.Manifest.Type == "mod" {
		if tree, err := graph.ParseModManifestTree(spec.Manifest.Path); err == nil {
			src.ModuleTree = &tree
		}
	}

	for _, manifestFile := range spec.Manifest.Files {
		absPath := filepath.Join(
			spec.Path,
			filepath.FromSlash(strings.ReplaceAll(manifestFile, "\\", "/")),
		)
		if _, err := os.Stat(absPath); err != nil {
			// File listed in manifest but not present on disk — skip silently.
			continue
		}
		switch strings.ToLower(filepath.Ext(absPath)) {
		case ".xml":
			src.XMLFiles = append(src.XMLFiles, absPath)
		case ".lua":
			src.LuaFiles = append(src.LuaFiles, absPath)
		}
	}
	return src
}

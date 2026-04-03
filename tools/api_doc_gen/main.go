package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"roraddons/tools/api_doc_gen/generator"
	"roraddons/tools/api_doc_gen/graph"
	"roraddons/tools/api_doc_gen/lua_parser"
	"roraddons/tools/api_doc_gen/normalizer"
	"roraddons/tools/api_doc_gen/platform"
	"roraddons/tools/api_doc_gen/site"
	"roraddons/tools/api_doc_gen/xml_structure"
)

const sourceRootManifestFile = ".api_doc_gen_source_root"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "api_doc_gen:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) > 0 && strings.EqualFold(strings.TrimSpace(args[0]), "generate") {
		return runGenerateCommand(args[1:])
	}
	return runLegacyMode(args)
}

func runGenerateCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: go run tools/api_doc_gen/main.go generate <addon|platform|site> <inputRoot> <outputRoot> [addonName ...]")
	}
	mode := strings.ToLower(strings.TrimSpace(args[0]))
	positional := args[1:]
	switch mode {
	case "addon":
		return runAddonMode(positional)
	case "platform":
		return runPlatformMode(positional)
	case "site":
		return runSiteMode(positional)
	default:
		return fmt.Errorf("unsupported generate mode %q", mode)
	}
}

func runLegacyMode(args []string) error {
	flags := flag.NewFlagSet("api_doc_gen", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	mode := flags.String("mode", "addon", "generation mode: addon, platform, or site")
	if err := flags.Parse(args); err != nil {
		return err
	}
	positional := flags.Args()

	switch strings.ToLower(strings.TrimSpace(*mode)) {
	case "site":
		return runSiteMode(positional)
	case "platform":
		return runPlatformMode(positional)
	case "addon", "":
		return runAddonMode(positional)
	default:
		return fmt.Errorf("unsupported mode %q", *mode)
	}
}

func runAddonMode(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: go run tools/api_doc_gen/main.go generate addon <sourceRoot> <outputRoot> [addonName ...]")
	}

	sourceRoot, err := resolveSourceRoot(args[0])
	if err != nil {
		return err
	}
	outputRoot, err := filepath.Abs(args[1])
	if err != nil {
		return fmt.Errorf("resolve output root: %w", err)
	}
	filters := args[2:]

	addonSpecs, err := graph.DiscoverAddons(sourceRoot, filters)
	if err != nil {
		return err
	}
	if len(addonSpecs) == 0 {
		return fmt.Errorf("no addons found in %s", sourceRoot)
	}

	addons := make([]graph.AddonModel, 0, len(addonSpecs))
	for _, spec := range addonSpecs {
		addon, err := parseAddon(spec)
		if err != nil {
			return err
		}
		addons = append(addons, addon)
	}

	sort.Slice(addons, func(i, j int) bool {
		return addons[i].Name < addons[j].Name
	})

	corpus := graph.BuildCorpus(sourceRoot, addons)
	corpus.Normalized = normalizer.Normalize(corpus)
	return generator.Generate(outputRoot, corpus)
}

func runPlatformMode(args []string) error {
	positional, sourceRootArg, err := extractPlatformArgs(args)
	if err != nil {
		return err
	}
	if len(positional) < 2 {
		return fmt.Errorf("usage: go run tools/api_doc_gen/main.go generate platform <addonApiRoot> <outputRoot> [--source-root <sourceRoot>]")
	}
	apiRefRoot, err := resolveExistingDirectory(positional[0])
	if err != nil {
		return err
	}
	outputRoot, err := filepath.Abs(positional[1])
	if err != nil {
		return fmt.Errorf("resolve output root: %w", err)
	}

	if sourceRootArg == "" {
		return fmt.Errorf("platform generation requires --source-root in contract-only mode")
	}
	sourceRoot, err := resolveExistingDirectory(sourceRootArg)
	if err != nil {
		return fmt.Errorf("resolve source root: %w", err)
	}

	contracts, err := platform.LoadContractInputs(apiRefRoot)
	if err != nil {
		return fmt.Errorf("load contract inputs: %w", err)
	}
	corpus := platform.BuildWithOptions(contracts, platform.BuildOptions{SourceRoot: sourceRoot})
	return platform.Generate(outputRoot, corpus)
}

func extractPlatformArgs(args []string) ([]string, string, error) {
	positional := make([]string, 0, len(args))
	sourceRoot := ""

	for i := 0; i < len(args); i++ {
		arg := strings.TrimSpace(args[i])
		if arg == "" {
			continue
		}

		if arg == "--source-root" {
			if i+1 >= len(args) {
				return nil, "", fmt.Errorf("--source-root requires a value")
			}
			i++
			sourceRoot = strings.TrimSpace(args[i])
			if sourceRoot == "" {
				return nil, "", fmt.Errorf("--source-root requires a non-empty value")
			}
			continue
		}

		if strings.HasPrefix(arg, "--source-root=") {
			sourceRoot = strings.TrimSpace(strings.TrimPrefix(arg, "--source-root="))
			if sourceRoot == "" {
				return nil, "", fmt.Errorf("--source-root requires a non-empty value")
			}
			continue
		}

		positional = append(positional, arg)
	}

	return positional, sourceRoot, nil
}

func runSiteMode(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: go run tools/api_doc_gen/main.go generate site <warApiRoot> <outputRoot>")
	}
	warAPIRoot, err := resolveExistingDirectory(args[0])
	if err != nil {
		return err
	}
	outputRoot, err := filepath.Abs(args[1])
	if err != nil {
		return fmt.Errorf("resolve output root: %w", err)
	}
	return site.Generate(warAPIRoot, outputRoot)
}

func parseAddon(spec graph.AddonSpec) (graph.AddonModel, error) {
	addon := graph.AddonModel{
		Name:     spec.Name,
		Root:     graph.NormalizePath(spec.Path),
		Manifest: spec.Manifest,
	}

	luaFiles := resolveAddonLuaFiles(spec)
	luaSeen := make(map[string]bool, len(luaFiles))
	for _, luaFile := range luaFiles {
		luaSeen[strings.ToLower(graph.NormalizePath(luaFile))] = true
	}

	for _, manifestFile := range spec.Manifest.Files {
		absolutePath := filepath.Join(spec.Path, filepath.FromSlash(strings.ReplaceAll(manifestFile, "\\", "/")))
		normalizedAbsolute := strings.ToLower(graph.NormalizePath(absolutePath))
		extension := strings.ToLower(filepath.Ext(absolutePath))
		switch extension {
		case ".lua":
			if !luaSeen[normalizedAbsolute] {
				continue
			}
			delete(luaSeen, normalizedAbsolute)
			parsed, err := lua_parser.ParseFile(spec.Name, absolutePath, spec.Manifest)
			if err != nil {
				fmt.Fprintf(os.Stderr, "warning: skipping lua file %s: %v\n", graph.NormalizePath(absolutePath), err)
				continue
			}
			addon.Functions = append(addon.Functions, parsed.Functions...)
			addon.Modules = append(addon.Modules, parsed.Modules...)
			addon.Events = append(addon.Events, parsed.Events...)
			addon.State = append(addon.State, parsed.State...)
		case ".xml":
			parsed, err := parseXMLFile(spec.Name, absolutePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "warning: skipping xml file %s: %v\n", graph.NormalizePath(absolutePath), err)
				continue
			}
			addon.Frames = append(addon.Frames, parsed.Frames...)
			addon.Handlers = append(addon.Handlers, parsed.Handlers...)
		case ".mod":
			continue
		default:
			continue
		}
	}

	for _, luaFile := range luaFiles {
		normalizedAbsolute := strings.ToLower(graph.NormalizePath(luaFile))
		if _, remaining := luaSeen[normalizedAbsolute]; !remaining {
			continue
		}
		delete(luaSeen, normalizedAbsolute)
		parsed, err := lua_parser.ParseFile(spec.Name, luaFile, spec.Manifest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: skipping lua file %s: %v\n", graph.NormalizePath(luaFile), err)
			continue
		}
		addon.Functions = append(addon.Functions, parsed.Functions...)
		addon.Modules = append(addon.Modules, parsed.Modules...)
		addon.Events = append(addon.Events, parsed.Events...)
		addon.State = append(addon.State, parsed.State...)
	}

	addon.Functions = graph.UniqueSortedFunctions(addon.Functions)
	addon.Modules = graph.UniqueSortedModules(addon.Modules)
	addon.Events = graph.UniqueSortedEvents(addon.Events)
	addon.State = graph.UniqueSortedState(addon.State)
	addon.Frames = graph.UniqueSortedFrames(addon.Frames)
	addon.Handlers = graph.UniqueSortedHandlers(addon.Handlers)
	return addon, nil
}

func resolveAddonLuaFiles(spec graph.AddonSpec) []string {
	manifestLua := make([]string, 0)
	for _, manifestFile := range spec.Manifest.Files {
		if !strings.EqualFold(filepath.Ext(manifestFile), ".lua") {
			continue
		}
		absolutePath := filepath.Join(spec.Path, filepath.FromSlash(strings.ReplaceAll(manifestFile, "\\", "/")))
		if _, err := os.Stat(absolutePath); err != nil {
			continue
		}
		manifestLua = append(manifestLua, graph.NormalizePath(absolutePath))
	}
	manifestLua = uniqueSortedNormalizedPaths(manifestLua)
	if len(manifestLua) > 0 {
		return manifestLua
	}

	fallback := make([]string, 0)
	_ = filepath.WalkDir(spec.Path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			if strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.EqualFold(filepath.Ext(path), ".lua") {
			fallback = append(fallback, graph.NormalizePath(path))
		}
		return nil
	})
	return uniqueSortedNormalizedPaths(fallback)
}

func uniqueSortedNormalizedPaths(values []string) []string {
	if len(values) == 0 {
		return []string{}
	}
	set := make(map[string]bool, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue
		}
		set[graph.NormalizePath(trimmed)] = true
	}
	result := make([]string, 0, len(set))
	for value := range set {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func parseXMLFile(addonName string, filePath string) (graph.XMLFileResult, error) {
	tree, err := xml_structure.ParseFile(addonName, filePath)
	if err != nil {
		return graph.XMLFileResult{}, err
	}

	frames, handlers := tree.ToFramesAndHandlers()
	graphFrames := make([]graph.Frame, 0, len(frames))
	for _, frame := range frames {
		graphFrames = append(graphFrames, graph.Frame{
			Addon:                   frame.Addon,
			Name:                    frame.Name,
			RawName:                 frame.RawName,
			Type:                    frame.Type,
			Parent:                  frame.Parent,
			ParentType:              frame.ParentType,
			File:                    frame.File,
			Line:                    frame.Line,
			Children:                frame.Children,
			ChildElementTypes:       frame.ChildElementTypes,
			StructuralChildTypes:    frame.StructuralChildTypes,
			StructuralChildAttrKeys: frame.StructuralChildAttrKeys,
			CompositionSnippet:      frame.CompositionSnippet,
			Inherits:                frame.Inherits,
			Attributes:              frame.Attributes,
			Template:                frame.Template,
		})
	}

	graphHandlers := make([]graph.XMLHandler, 0, len(handlers))
	for _, handler := range handlers {
		graphHandlers = append(graphHandlers, graph.XMLHandler{
			Addon:    handler.Addon,
			Frame:    handler.Frame,
			Event:    handler.Event,
			Function: handler.Function,
			File:     handler.File,
			Line:     handler.Line,
		})
	}

	return graph.XMLFileResult{
		Frames:   graph.UniqueSortedFrames(graphFrames),
		Handlers: graph.UniqueSortedHandlers(graphHandlers),
	}, nil
}

func resolveSourceRoot(argument string) (string, error) {
	absolutePath, err := resolveExistingDirectory(argument)
	if err == nil {
		return resolveMappedSourceRoot(absolutePath)
	}

	if strings.EqualFold(filepath.Base(argument), "addons") {
		workingDirectory, cwdErr := os.Getwd()
		if cwdErr == nil && directoryLooksLikeAddonRoot(workingDirectory) {
			return workingDirectory, nil
		}
	}

	if err != nil {
		return "", fmt.Errorf("resolve source root: %w", err)
	}
	return "", fmt.Errorf("source root does not exist: %s", argument)
}

func resolveMappedSourceRoot(absolutePath string) (string, error) {
	manifestPath := filepath.Join(absolutePath, sourceRootManifestFile)
	contents, err := os.ReadFile(manifestPath)
	if err != nil {
		if os.IsNotExist(err) {
			return absolutePath, nil
		}
		return "", fmt.Errorf("read source root manifest: %w", err)
	}
	target := strings.TrimSpace(string(contents))
	if target == "" {
		return "", fmt.Errorf("source root manifest %s is empty", manifestPath)
	}
	if !filepath.IsAbs(target) {
		target = filepath.Join(absolutePath, target)
	}
	resolved, err := resolveExistingDirectory(target)
	if err != nil {
		return "", fmt.Errorf("resolve mapped source root: %w", err)
	}
	return resolved, nil
}

func resolveExistingDirectory(argument string) (string, error) {
	absolutePath, err := filepath.Abs(argument)
	if err != nil {
		return "", fmt.Errorf("resolve source root: %w", err)
	}
	info, statErr := os.Stat(absolutePath)
	if statErr != nil {
		return "", statErr
	}
	if !info.IsDir() {
		return "", fmt.Errorf("path is not a directory: %s", argument)
	}
	return absolutePath, nil
}

func directoryLooksLikeAddonRoot(path string) bool {
	entries, err := os.ReadDir(path)
	if err != nil {
		return false
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, ".") || shouldSkipSourceEntry(name) {
			continue
		}
		addonPath := filepath.Join(path, name)
		modMatches, _ := filepath.Glob(filepath.Join(addonPath, "*.mod"))
		if len(modMatches) > 0 {
			return true
		}
		tocMatches, _ := filepath.Glob(filepath.Join(addonPath, "*.toc"))
		if len(tocMatches) > 0 {
			return true
		}
	}
	return false
}

func shouldSkipSourceEntry(name string) bool {
	for _, candidate := range []string{"API_Ref", "WAR_API_Ref", "docs", "data", "infra", "tools"} {
		if strings.EqualFold(name, candidate) {
			return true
		}
	}
	return false
}

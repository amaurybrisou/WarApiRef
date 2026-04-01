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
	"roraddons/tools/api_doc_gen/xml_parser"
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
	if len(args) < 2 {
		return fmt.Errorf("usage: go run tools/api_doc_gen/main.go generate platform <addonApiRoot> <outputRoot>")
	}
	apiRefRoot, err := resolveExistingDirectory(args[0])
	if err != nil {
		return err
	}
	outputRoot, err := filepath.Abs(args[1])
	if err != nil {
		return fmt.Errorf("resolve output root: %w", err)
	}
	source, err := platform.ParseAPIRef(apiRefRoot)
	if err != nil {
		return err
	}
	corpus := platform.Build(source)
	return platform.Generate(outputRoot, corpus)
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

	for _, manifestFile := range spec.Manifest.Files {
		absolutePath := filepath.Join(spec.Path, filepath.FromSlash(strings.ReplaceAll(manifestFile, "\\", "/")))
		extension := strings.ToLower(filepath.Ext(absolutePath))
		switch extension {
		case ".lua":
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
			parsed, err := xml_parser.ParseFile(spec.Name, absolutePath)
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

	addon.Functions = graph.UniqueSortedFunctions(addon.Functions)
	addon.Modules = graph.UniqueSortedModules(addon.Modules)
	addon.Events = graph.UniqueSortedEvents(addon.Events)
	addon.State = graph.UniqueSortedState(addon.State)
	addon.Frames = graph.UniqueSortedFrames(addon.Frames)
	addon.Handlers = graph.UniqueSortedHandlers(addon.Handlers)
	return addon, nil
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

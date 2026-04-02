package mod_semantic

import (
	"strings"

	"roraddons/tools/api_doc_gen/graph"
)

// knownLifecycleHooks maps tag names to HookKind for sections that are
// known to contain executable lifecycle actions.
var knownLifecycleHooks = map[string]HookKind{
	"OnInitialize": HookKindOnInitialize,
	"OnUpdate":     HookKindOnUpdate,
	"OnShutdown":   HookKindOnShutdown,
}

// knownActions maps tag names to ActionKind for recognised action elements.
var knownActions = map[string]ActionKind{
	"CallFunction": ActionKindCallFunction,
	"CreateWindow": ActionKindCreateWindow,
}

// metadataSections is the set of top-level .mod sections that hold
// declarative metadata rather than executable lifecycle actions.  They are
// excluded from lifecycle hook discovery so that, e.g., a <Files> section
// with <File> children is never misidentified as an unknown hook.
var metadataSections = map[string]bool{
	"Files":          true,
	"SavedVariables": true,
	"Author":         true,
	"Description":    true,
}

// AnalyzeTree performs semantic analysis on a ModuleTree.
//
// It classifies lifecycle hooks and their actions, correlates CallFunction
// references against luaByName, and correlates CreateWindow references
// against xmlFrameNames.
//
// Parameters:
//   - addonName: canonical addon identifier.
//   - tree: the ModuleTree to analyse; may be nil (returns empty result).
//   - luaByName: maps normalised function name to qualified Lua names.
//     Build with [BuildLuaNameIndex].
//   - xmlFrameNames: maps XML frame/window name to true (existence index).
//     Build with [BuildXMLFrameNameIndex].
//
// The raw tree is never mutated.
func AnalyzeTree(addonName string, tree *graph.ModuleTree, luaByName map[string][]string, xmlFrameNames map[string]bool) *ModuleSemantic {
	sem := &ModuleSemantic{
		AddonName: addonName,
		RawTree:   tree,
	}
	if tree == nil || tree.Root == nil {
		return sem
	}

	// Saved variables from the flat Manifest (already extracted from tree).
	sem.SavedVariables = append(sem.SavedVariables, tree.Manifest.SavedVariables...)

	for _, child := range tree.Root.Children {
		hookKind, isKnown := knownLifecycleHooks[child.Tag]
		if isKnown {
			hook := buildHook(child, hookKind, luaByName, xmlFrameNames)
			sem.LifecycleHooks = append(sem.LifecycleHooks, hook)
			continue
		}
		// Preserve unknown non-metadata sections that carry action-like
		// children as unknown lifecycle hooks.
		if !metadataSections[child.Tag] && len(child.Children) > 0 {
			hook := buildHook(child, HookKindUnknown, luaByName, xmlFrameNames)
			sem.LifecycleHooks = append(sem.LifecycleHooks, hook)
		}
	}

	return sem
}

// buildHook constructs a LifecycleHook from a ModNode, classifying and
// correlating every child action node.
func buildHook(node *graph.ModNode, kind HookKind, luaByName map[string][]string, xmlFrameNames map[string]bool) LifecycleHook {
	hook := LifecycleHook{
		Kind:    kind,
		Tag:     node.Tag,
		RawNode: node,
	}
	for i, child := range node.Children {
		action := buildAction(child, i, node.Tag, luaByName, xmlFrameNames)
		hook.Actions = append(hook.Actions, action)
	}
	return hook
}

// buildAction constructs a LifecycleAction from a ModNode, classifying the
// action kind and attempting correlation with known Lua/XML symbols.
func buildAction(node *graph.ModNode, index int, sourceHook string, luaByName map[string][]string, xmlFrameNames map[string]bool) LifecycleAction {
	kind := knownActions[node.Tag] // empty string when unrecognised
	name := strings.TrimSpace(node.Attrs["name"])

	action := LifecycleAction{
		Index:          index,
		Tag:            node.Tag,
		Kind:           kind,
		Attrs:          node.Attrs,
		Name:           name,
		SourceHook:     sourceHook,
		RawNode:        node,
		EvidenceSource: ".mod",
		Resolution:     ResolutionUnresolved,
	}

	switch kind {
	case ActionKindCallFunction:
		action = resolveCallFunction(action, luaByName)
	case ActionKindCreateWindow:
		action = resolveCreateWindow(action, xmlFrameNames)
	}

	return action
}

// resolveCallFunction attempts to match a CallFunction action's name against
// the Lua function index.
//
//   - Exact match (1 candidate): ResolutionExact / HIGH confidence.
//   - Ambiguous (>1 candidates): ResolutionAmbiguous / MEDIUM confidence.
//   - No match: ResolutionUnresolved (action kept; reference visible).
func resolveCallFunction(action LifecycleAction, luaByName map[string][]string) LifecycleAction {
	if action.Name == "" || len(luaByName) == 0 {
		return action
	}
	normalized := normalizeFuncName(action.Name)
	matches, ok := luaByName[normalized]
	if !ok || len(matches) == 0 {
		return action
	}
	action.MatchedLuaFunctions = matches
	if len(matches) == 1 {
		action.Resolution = ResolutionExact
		action.Confidence = "HIGH"
	} else {
		action.Resolution = ResolutionAmbiguous
		action.Confidence = "MEDIUM"
	}
	return action
}

// resolveCreateWindow attempts to match a CreateWindow action's name against
// the XML frame name index.
//
//   - Exact match: ResolutionExact / HIGH confidence.
//   - No match: ResolutionUnresolved (action kept; reference visible).
func resolveCreateWindow(action LifecycleAction, xmlFrameNames map[string]bool) LifecycleAction {
	if action.Name == "" || len(xmlFrameNames) == 0 {
		return action
	}
	if xmlFrameNames[action.Name] {
		action.MatchedXMLNames = []string{action.Name}
		action.Resolution = ResolutionExact
		action.Confidence = "HIGH"
	}
	return action
}

// normalizeFuncName normalises a Lua function name for lookup: it takes the
// last dot-separated segment and lower-cases it, matching the normalisation
// used by semantic_merge.normalizeFunc.
func normalizeFuncName(name string) string {
	if idx := strings.LastIndex(name, "."); idx >= 0 {
		name = name[idx+1:]
	}
	return strings.ToLower(name)
}

// BuildLuaNameIndex builds a lookup map from normalised function name to
// qualified Lua function names, suitable for passing to [AnalyzeTree].
//
// qualifiedNames is a slice of fully-qualified function names such as
// "MyAddon.Initialize" or "MyAddon.Update".
func BuildLuaNameIndex(qualifiedNames []string) map[string][]string {
	idx := make(map[string][]string, len(qualifiedNames))
	for _, qn := range qualifiedNames {
		key := normalizeFuncName(qn)
		idx[key] = append(idx[key], qn)
	}
	return idx
}

// BuildXMLFrameNameIndex builds a presence index from a slice of XML frame
// names, suitable for passing to [AnalyzeTree].
func BuildXMLFrameNameIndex(names []string) map[string]bool {
	idx := make(map[string]bool, len(names))
	for _, n := range names {
		if n != "" {
			idx[n] = true
		}
	}
	return idx
}

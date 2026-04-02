// Package mod_semantic implements semantic analysis of .mod manifest trees.
//
// It builds typed analysis results on top of a [graph.ModuleTree] without
// mutating the raw tree.  The raw tree is always preserved as-is; the
// semantic layer adds classification, correlation, and resolution data on
// top.
//
// Design principles:
//   - ModuleTree is the source of truth; Manifest is compatibility only.
//   - Unknown lifecycle sections and actions are preserved, not dropped.
//   - Hook and action order from the .mod file is maintained.
//   - Every semantic node keeps a pointer back to its originating ModNode.
//   - Unresolved and ambiguous references remain explicitly visible.
package mod_semantic

import "roraddons/tools/api_doc_gen/graph"

// HookKind classifies a top-level lifecycle section of a .mod manifest.
// The empty string (HookKindUnknown) means the section was preserved but
// not matched against any known lifecycle hook name.
type HookKind string

const (
	HookKindOnInitialize HookKind = "OnInitialize"
	HookKindOnUpdate     HookKind = "OnUpdate"
	HookKindOnShutdown   HookKind = "OnShutdown"
	HookKindUnknown      HookKind = ""
)

// ActionKind classifies an action element inside a lifecycle hook.
// The empty string (ActionKindUnknown) means the element was preserved but
// not matched against any known action tag.
type ActionKind string

const (
	ActionKindCallFunction ActionKind = "CallFunction"
	ActionKindCreateWindow ActionKind = "CreateWindow"
	ActionKindUnknown      ActionKind = ""
)

// ResolutionStatus records how well an action reference was resolved against
// the addon's Lua function definitions or XML frame names.
type ResolutionStatus string

const (
	// ResolutionExact means exactly one match was found with high confidence.
	ResolutionExact ResolutionStatus = "exact"
	// ResolutionAmbiguous means multiple candidates exist; auto-resolution
	// was not performed.
	ResolutionAmbiguous ResolutionStatus = "ambiguous"
	// ResolutionUnresolved means no matching Lua function or XML frame was
	// found.  The reference is retained and remains visible.
	ResolutionUnresolved ResolutionStatus = "unresolved"
)

// LifecycleAction is a single action element inside a lifecycle hook,
// combining the raw .mod node with semantic classification and resolution.
//
// Every action retains its RawNode pointer so callers can always access
// the unmodified .mod tree node.
type LifecycleAction struct {
	// Index is the 0-based position of this action within the enclosing hook.
	Index int

	// Tag is the raw XML element name (e.g. "CallFunction", "CreateWindow",
	// or an unrecognised custom tag).
	Tag string

	// Kind classifies the action type.  ActionKindUnknown is used for tags
	// that are not in the set of known action elements.
	Kind ActionKind

	// Attrs holds every attribute from the .mod node, keyed by attribute name.
	Attrs map[string]string

	// Name is the value of the "name" attribute (convenience accessor).
	// It is empty when the attribute is absent.
	Name string

	// SourceHook is the tag name of the enclosing lifecycle hook.
	SourceHook string

	// RawNode is an unmodified reference back to the originating ModNode.
	// It is never nil for actions derived from a real tree node.
	RawNode *graph.ModNode

	// Resolution tracks how the action's name reference was resolved.
	// It is always set; ResolutionUnresolved is the initial / default state.
	Resolution ResolutionStatus

	// EvidenceSource identifies the file type that produced this evidence.
	// It is always ".mod" for actions derived from a .mod manifest.
	EvidenceSource string

	// MatchedLuaFunctions holds the qualified Lua function names that were
	// found when Kind == ActionKindCallFunction.  It is non-nil only when
	// Resolution != ResolutionUnresolved.
	MatchedLuaFunctions []string

	// MatchedXMLNames holds the XML frame/window names that were found when
	// Kind == ActionKindCreateWindow.  It is non-nil only when
	// Resolution != ResolutionUnresolved.
	MatchedXMLNames []string

	// Confidence is the resolution confidence: "HIGH", "MEDIUM", or "LOW".
	// It is empty when the action is unresolved.
	Confidence string
}

// LifecycleHook is a lifecycle section from the .mod file, preserving both
// the raw ModNode and the semantically classified actions inside it.
//
// Unknown hook sections (HookKindUnknown) are kept so that future or custom
// lifecycle containers are never silently discarded.
type LifecycleHook struct {
	// Kind classifies the hook type.  HookKindUnknown is used for sections
	// that are not in the set of known lifecycle hook names.
	Kind HookKind

	// Tag is the raw XML element name (identical to Kind for known hooks,
	// but preserved generically for unknown hooks).
	Tag string

	// Actions are the classified actions within this hook, in document order.
	// Mixed known and unknown actions are all retained.
	Actions []LifecycleAction

	// RawNode is an unmodified reference back to the originating ModNode.
	RawNode *graph.ModNode
}

// ModuleSemantic is the semantic analysis result for a single .mod addon.
//
// It is produced by [AnalyzeTree] and must never mutate the raw ModuleTree.
// All raw tree nodes remain accessible via the RawNode pointers on each
// LifecycleHook and LifecycleAction.
type ModuleSemantic struct {
	// AddonName is the canonical addon identifier derived from the manifest.
	AddonName string

	// LifecycleHooks contains every lifecycle section found in the .mod file,
	// in document order.  Both known and unknown hook sections are included.
	LifecycleHooks []LifecycleHook

	// SavedVariables lists the declared saved variable names, extracted from
	// the SavedVariables section of the .mod tree.
	SavedVariables []string

	// RawTree is the original ModuleTree.  It is never mutated by analysis.
	RawTree *graph.ModuleTree
}

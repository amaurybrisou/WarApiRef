package xmlsemantic

import (
	"sort"
	"strings"
)

// maxLuaAPICallsPerEvent is the maximum number of Lua API call records kept per
// XML event binding. Keeping this focused prevents low-signal calls from
// overwhelming the per-event detail.
const maxLuaAPICallsPerEvent = 10

// Build constructs ElementTypeSchema records for all element types observed in
// the provided frames and handlers. It uses Lua function call data to populate
// HandlerEventBinding.LuaAPICalls with real call-graph-derived API usage — not
// graph co-occurrence approximations.
//
// The result is keyed by element type name (e.g. "ListBox", "Button").
func Build(
	frames []FrameInput,
	handlers []HandlerInput,
	funcs []FunctionCallInput,
) map[string]*ElementTypeSchema {
	// ── frame-type lookup: "addon|frame-name" → element type ─────────────
	frameTypeLookup := make(map[string]string, len(frames))
	for _, f := range frames {
		if f.Type != "" {
			frameTypeLookup[f.Addon+"|"+f.Name] = f.Type
		}
	}

	// ── function call lookup: function name → set of called API names ────
	funcCallLookup := make(map[string]map[string]bool, len(funcs))
	for _, fn := range funcs {
		if fn.Name == "" {
			continue
		}
		if _, ok := funcCallLookup[fn.Name]; !ok {
			funcCallLookup[fn.Name] = make(map[string]bool, len(fn.Calls))
		}
		for _, call := range fn.Calls {
			call = strings.TrimSpace(call)
			if call != "" {
				funcCallLookup[fn.Name][call] = true
			}
		}
	}

	// ── schema + intermediate data stores ────────────────────────────────
	schemas := make(map[string]*ElementTypeSchema)
	// attrsBySchema[typeName][attrName] → {count, values}
	type attrData struct {
		count  int
		values map[string]int
	}
	attrsBySchema := make(map[string]map[string]*attrData)
	addonsBySchema := make(map[string]map[string]bool)

	ensureSchema := func(name string) *ElementTypeSchema {
		if s, ok := schemas[name]; ok {
			return s
		}
		s := &ElementTypeSchema{
			Name:               name,
			StructuralChildren: make(map[string]int),
			NamedChildTypes:    make(map[string]int),
			ParentTypes:        make(map[string]int),
		}
		schemas[name] = s
		return s
	}

	ensureAttrData := func(schemaName, attrName string) *attrData {
		if _, ok := attrsBySchema[schemaName]; !ok {
			attrsBySchema[schemaName] = make(map[string]*attrData)
		}
		if _, ok := attrsBySchema[schemaName][attrName]; !ok {
			attrsBySchema[schemaName][attrName] = &attrData{values: make(map[string]int)}
		}
		return attrsBySchema[schemaName][attrName]
	}

	// ── pass 1: process frames ────────────────────────────────────────────
	for _, f := range frames {
		if strings.TrimSpace(f.Type) == "" {
			continue
		}
		s := ensureSchema(f.Type)
		s.TotalInstances++

		if _, ok := addonsBySchema[f.Type]; !ok {
			addonsBySchema[f.Type] = make(map[string]bool)
		}
		addonsBySchema[f.Type][f.Addon] = true

		// Attribute profiles: track key presence and observed values.
		for key, val := range f.Attributes {
			key = strings.TrimSpace(key)
			if key == "" {
				continue
			}
			ad := ensureAttrData(f.Type, key)
			ad.count++
			val = strings.TrimSpace(val)
			// Record concrete values but skip name/inherits (too instance-specific)
			// and very long values (e.g. lengthy function names).
			if val != "" && key != "name" && key != "inherits" && len(val) <= 48 {
				ad.values[val]++
			}
		}

		// Structural (unnamed) children.
		for _, childType := range f.StructuralChildTypes {
			if childType != "" {
				s.StructuralChildren[childType]++
			}
		}

		// Named child element types.
		for _, childType := range f.NamedChildTypes {
			if childType != "" {
				s.NamedChildTypes[childType]++
			}
		}

		// Parent types.
		if f.ParentType != "" {
			s.ParentTypes[f.ParentType]++
		}
	}

	// ── convert attrsBySchema to AttributeProfile slices ─────────────────
	for schemaName, attrs := range attrsBySchema {
		s := ensureSchema(schemaName)
		total := s.TotalInstances
		profiles := make([]AttributeProfile, 0, len(attrs))
		for attrName, ad := range attrs {
			isRequired := total > 0 && float64(ad.count)/float64(total) >= 0.75
			profiles = append(profiles, AttributeProfile{
				Name:       attrName,
				Count:      ad.count,
				Total:      total,
				TopValues:  topValuesByCount(ad.values, 5),
				IsRequired: isRequired,
			})
		}
		sort.Slice(profiles, func(i, j int) bool {
			if profiles[i].Count != profiles[j].Count {
				return profiles[i].Count > profiles[j].Count // descending by frequency
			}
			return profiles[i].Name < profiles[j].Name
		})
		s.AttributeProfiles = profiles
	}

	// ── populate addon lists ──────────────────────────────────────────────
	for schemaName, addons := range addonsBySchema {
		s := ensureSchema(schemaName)
		names := make([]string, 0, len(addons))
		for name := range addons {
			names = append(names, name)
		}
		sort.Strings(names)
		s.Addons = names
	}

	// ── pass 2: process handlers with Lua call-chain derivation ──────────
	//
	// Intermediate: per (schema, event) → direct func counts + API call counts.
	type eventKey struct{ schema, event string }
	type handlerIntermediate struct {
		directFuncs map[string]int
		apiCalls    map[string]int // api function name → count
	}
	handlerData := make(map[eventKey]*handlerIntermediate)

	ensureHandlerData := func(schema, event string) *handlerIntermediate {
		k := eventKey{schema, event}
		if _, ok := handlerData[k]; !ok {
			handlerData[k] = &handlerIntermediate{
				directFuncs: make(map[string]int),
				apiCalls:    make(map[string]int),
			}
		}
		return handlerData[k]
	}

	for _, h := range handlers {
		elemType := frameTypeLookup[h.Addon+"|"+h.Frame]
		if elemType == "" {
			continue
		}
		event := strings.TrimSpace(h.Event)
		fn := strings.TrimSpace(h.Function)
		if event == "" || fn == "" {
			continue
		}

		hd := ensureHandlerData(elemType, event)
		hd.directFuncs[fn]++

		// Derive API calls from this handler function's call graph.
		for apiCall := range funcCallLookup[fn] {
			if isExcludedCall(apiCall) {
				continue
			}
			hd.apiCalls[apiCall]++
		}
	}

	// ── convert handler intermediates to HandlerEventBinding slices ───────
	for key, hd := range handlerData {
		s := ensureSchema(key.schema)
		inferred, argsConf := inferHandlerArgs(key.event)

		luaAPICalls := make([]LuaAPICall, 0, len(hd.apiCalls))
		for fn, cnt := range hd.apiCalls {
			luaAPICalls = append(luaAPICalls, LuaAPICall{
				FunctionName: fn,
				ViaEvent:     key.event,
				Count:        cnt,
			})
		}
		sort.Slice(luaAPICalls, func(i, j int) bool {
			if luaAPICalls[i].Count != luaAPICalls[j].Count {
				return luaAPICalls[i].Count > luaAPICalls[j].Count
			}
			return luaAPICalls[i].FunctionName < luaAPICalls[j].FunctionName
		})
		if len(luaAPICalls) > maxLuaAPICallsPerEvent {
			luaAPICalls = luaAPICalls[:maxLuaAPICallsPerEvent]
		}

		s.HandlerBindings = append(s.HandlerBindings, HandlerEventBinding{
			Event:          key.event,
			DirectFuncs:    topKeysByCount(hd.directFuncs, 6),
			LuaAPICalls:    luaAPICalls,
			InferredArgs:   inferred,
			ArgsConfidence: argsConf,
		})
	}

	// Sort handler bindings by event name for stable, reproducible output.
	for _, s := range schemas {
		sort.Slice(s.HandlerBindings, func(i, j int) bool {
			return s.HandlerBindings[i].Event < s.HandlerBindings[j].Event
		})
	}

	return schemas
}

// ── helpers ───────────────────────────────────────────────────────────────────

// isExcludedCall returns true for Lua standard library calls and other
// non-WAR-API symbols that would add noise to LuaAPICall records.
func isExcludedCall(name string) bool {
	if name == "" {
		return true
	}
	// Lua standard library prefixes.
	for _, prefix := range []string{"string.", "table.", "math.", "io.", "os.", "coroutine.", "debug."} {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	// Exact standard library names.
	switch name {
	case "tonumber", "tostring", "pairs", "ipairs", "next", "type",
		"print", "error", "assert", "pcall", "xpcall", "rawequal",
		"select", "unpack", "rawget", "rawset", "setmetatable", "getmetatable",
		"require", "dofile", "load", "loadstring", "loadfile":
		return true
	// WAR UI framework string utilities — excluded because they appear in
	// almost every addon and add noise without meaningful semantic signal.
	case "towstring", "wstring", "L":
		return true
	}
	return false
}

// topKeysByCount returns up to n keys from m sorted by descending count.
func topKeysByCount(m map[string]int, n int) []string {
	type pair struct {
		key string
		cnt int
	}
	pairs := make([]pair, 0, len(m))
	for k, v := range m {
		k = strings.TrimSpace(k)
		if k != "" {
			pairs = append(pairs, pair{k, v})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].cnt != pairs[j].cnt {
			return pairs[i].cnt > pairs[j].cnt
		}
		return pairs[i].key < pairs[j].key
	})
	result := make([]string, 0, n)
	for i, p := range pairs {
		if i >= n {
			break
		}
		result = append(result, p.key)
	}
	return result
}

// topValuesByCount is an alias for topKeysByCount used on attribute-value maps.
func topValuesByCount(m map[string]int, n int) []string {
	return topKeysByCount(m, n)
}

// knownHandlerArgs maps well-known WAR XML handler event names to their expected
// Lua callback signature. Confidence is MEDIUM because these are observed
// patterns, not engine-source-verified contracts.
var knownHandlerArgs = map[string]string{
	// Lifecycle hooks — no engine-supplied arguments.
	"OnInitialize":           "function()",
	"OnShutdown":             "function()",
	"OnShow":                 "function()",
	"OnHide":                 "function()",
	"OnShown":                "function()",
	"OnHidden":               "function()",
	// Mouse / pointer events.
	"OnMouseOver":            "function()",
	"OnMouseOut":             "function()",
	"OnMouseOverEnd":         "function()",
	"OnMouseDown":            "function(button)",
	"OnMouseUp":              "function(button)",
	"OnLButtonUp":            "function()",
	"OnLButtonDown":          "function()",
	"OnRButtonUp":            "function()",
	"OnRButtonDown":          "function()",
	"OnClick":                "function()",
	"OnDoubleClick":          "function()",
	"OnMouseWheel":           "function(delta)",
	// Keyboard events.
	"OnEnterPressed":         "function()",
	"OnEscapePressed":        "function()",
	"OnKeyDown":              "function(key)",
	"OnKeyUp":                "function(key)",
	// Value / text change events.
	"OnTextChanged":          "function()",
	"OnValueChanged":         "function(value)",
	// Drag events.
	"OnDragStart":            "function()",
	"OnDragStop":             "function()",
	"OnReceiveDrag":          "function()",
	// Resize / move events.
	"OnSizeChanged":          "function(width, height)",
	"OnMove":                 "function()",
	// Per-frame update hook — elapsed time in seconds.
	"OnUpdate":               "function(elapsed)",
	// Edit box events.
	"OnCursorChanged":        "function()",
	"OnInputLanguageChanged": "function()",
	// Selection / list events.
	"OnSelectionChanged":     "function(selectedRow)",
	"OnScrollChanged":        "function()",
}

// inferHandlerArgs returns the best-effort callback signature and confidence
// level for the given XML handler event name.
func inferHandlerArgs(event string) (sig, confidence string) {
	if s, ok := knownHandlerArgs[event]; ok {
		return s, "MEDIUM"
	}
	return "function(...)", "LOW"
}

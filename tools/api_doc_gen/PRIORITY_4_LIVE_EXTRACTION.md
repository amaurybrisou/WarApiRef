# Priority 4: Live Game Extraction & Runtime Evidence

**Status:** 📋 SPECIFICATION PHASE (Ready for Implementation)

**Date:** 2025-01-16 (Session 4, Phase C)

**Objective:** Boost inference confidence from 50-95% to 99%+ by capturing real function behavior during addon execution.

---

## 1. Strategic Rationale

### Current Confidence Limits (Priorities 1-3)
- Static analysis confidence: 50-95%
- Evidence sources: source code patterns, call sites, literals
- Limitation: Cannot see actual runtime behavior
- Gap: Unknown calls, dynamic arguments, runtime state

### Runtime Evidence Advantage
- **Confidence Boost:** 95% → 99%+ (observed fact)
- **Evidence Type:** Runtime observation (highest precedence)
- **Latency:** Single execution captures evidence
- **Completeness:** Covers dynamic behavior static analysis misses

### Target Improvements
| Inference | Static Confidence | Runtime Confidence | Gain |
|-----------|-------------------|-------------------|------|
| Return types | 60-85% | 99%+ | +14-39% |
| Edge relationships | 70-90% | 99%+ | +9-29% |
| Argument types | 50-75% | 99%+ | +24-49% |
| Function roles | 65-85% | 99%+ | +14-34% |

---

## 2. Technical Architecture

### 2.1 Probe Infrastructure

**Design Pattern:** Lua hooks + event capture

```lua
-- Pseudo-code structure for instrumentation
ProbeManager = {
  hooks = {},
  events = {},
  observations = {}
}

function ProbeManager:InstallProbes()
  -- Install hooks on function entry/exit
  -- Track call contexts and return values
  -- Record state changes
end

function ProbeManager:CaptureObservation(context)
  -- context = {function, args, returns, state_before, state_after}
  -- Store with timestamp and execution context
end
```

### 2.2 Probe Types

#### Probe Type 1: Function Entry/Exit
**Purpose:** Capture actual arguments and return values

```
Hook on function entry:
  - Capture: argument names, types, values
  - Record: who called it (context)
  
Hook on function exit:
  - Capture: return values
  - Record: state changes made by function
  
Evidence Generated:
  - Run-time-observed-arguments (99% confidence)
  - Run-time-observed-returns (99% confidence)
  - Run-time-sequence-ordering (99% confidence)
```

**Example Observation:**
```json
{
  "function": "function:global:BroadcastEvent",
  "call_count": 42,
  "argument_patterns": [
    {"position": 1, "observed_types": ["integer", "string", "table"], "frequency": [40, 1, 1]},
    {"position": 2, "observed_types": ["table"], "frequency": [42]}
  ],
  "return_type": "nil",
  "return_frequency": 42,
  "state_changes": ["SystemData.Events queue updated"],
  "observation_source": "runtime:function_hook",
  "confidence_boost": 0.39  // Increase from ~60% to 99%
}
```

#### Probe Type 2: Event Handlers
**Purpose:** Track event flow and handler relationships

```
Hook on event system:
  - Intercept: Event registration
  - Intercept: Event broadcasts
  - Intercept: Handler execution
  
Evidence Generated:
  - event-triggers-handler (99% confidence)
  - handler-updates-state (99% confidence)
  - event-propagation-chain (99% confidence)
```

**Example Observation:**
```json
{
  "event": "SCENARIO_STOP_UPDATING_PLAYERS_STATS",
  "handlers": [
    "function:global:BroadcastEvent",
    "function:global:UpdateScenarioInfo"
  ],
  "execution_order": [0, 1],
  "frequency": 8,
  "observation_source": "runtime:event_hook",
  "confidence_boost": 0.29
}
```

#### Probe Type 3: UI State Changes
**Purpose:** Track UI element modifications

```
Hook on UI manipulation:
  - Intercept: Show/Hide calls
  - Intercept: Layout updates
  - Intercept: Property changes
  
Evidence Generated:
  - function-updates-ui-element (99% confidence)
  - element-visibility-sequence (99% confidence)
  - layout-modification-patterns (99% confidence)
```

**Example Observation:**
```json
{
  "ui_operation": "show",
  "element": "xml_frame:AlertFrame",
  "caller": "function:global:AlertTextWindow.AddLine",
  "preconditions": ["AlertFrame exists", "AlertFrame.hidden"],
  "postconditions": ["AlertFrame.visible", "AlertFrame updated"],
  "call_count": 87,
  "success_rate": 1.0,
  "observation_source": "runtime:ui_hook",
  "confidence_boost": 0.24
}
```

#### Probe Type 4: Data Access
**Purpose:** Track reads/writes to system data

```
Hook on data access:
  - Intercept: SystemData reads
  - Intercept: GameData reads
  - Intercept: State modifications
  
Evidence Generated:
  - function-reads-systemdata (99% confidence)
  - function-writes-gamedata (99% confidence)
  - data-dependency-graph (99% confidence)
```

**Example Observation:**
```json
{
  "data_entity": "SystemData.Events",
  "accessor": "function:global:BroadcastEvent",
  "operation": "read",
  "frequency": 42,
  "timing": "function_entry",
  "observation_source": "runtime:data_hook",
  "confidence_boost": 0.29
}
```

---

## 3. Probe Installation Points

### 3.1 Instrumentation Trigger
**When:** Plugin initialization / addon startup

**Mechanism:**
```lua
-- In AdvancedPetAssist.lua or decorator addon
function InstallInstrumentationProbes()
  -- Call runtime probe installer
  ProbeManager:InstallProbes()
  
  -- Configure observation filters
  ProbeManager:SetSamplingRate(1.0)  -- 100% capture or 0.1 for 10% sampling
  
  -- Register collector callback
  ProbeManager:OnObservationCaptured(function(obs)
    LogObservation(obs)
  end)
end
```

### 3.2 Addon Integration Points

**Option A: Decorator Addon (Recommended)**
- Create separate `InstrumentationProbe` addon
- Installs probes on top of existing addons
- Zero impact on target addon functionality
- Easy to enable/disable

**Option B: Embedded Instrumentation**
- Add probe calls to existing addons (AdvancedPetAssist, etc.)
- Tighter integration
- Risk of impacting addon behavior
- Requires coordination with addon authors

**Option C: Engine-Level Hooks**
- Modify game engine API binding layer
- Intercept all function calls globally
- Highest coverage, highest risk
- Requires war/reckoning source modification

**Recommendation:** Option A (Decorator addon)

### 3.3 Execution Context

**Addon Execution Scenarios to Capture:**
1. **Login to character** — Initialization probe fire
2. **Zone transitions** — Event handler chains
3. **Combat sequences** — High-frequency function calls
4. **UI interactions** — Click handlers, visibility changes
5. **State updates** — Event broadcasts, data modifications
6. **Campaign/faction activities** — Domain-specific flows

---

## 4. Evidence Collection & Storage

### 4.1 Live Collection Format

**Evidence File Structure:**
```json
{
  "capture_session": {
    "timestamp": "2025-01-16T12:00:00Z",
    "addon": "AdvancedPetAssist",
    "version": "1.0",
    "character": "Sigurg",
    "realm": "Martyr's Square",
    "session_duration_seconds": 3600,
    "observation_count": 847
  },
  "observations": [
    {
      "type": "function_call",
      "function": "function:global:GetPetAbility",
      "arguments": [
        {"position": 1, "value": 2, "type": "integer"},
        {"position": 2, "value": null, "type": "nil"}
      ],
      "return_value": {"type": "table", "fields": ["name", "description", "icon"]},
      "timestamp": "2025-01-16T12:00:01.123Z",
      "context": {"caller": "APACore.OnUpdate", "event": null}
    },
    ...
  ]
}
```

### 4.2 Aggregation Format

**Aggregated Evidence (for submission to API reference):**
```json
{
  "runtime_evidence": {
    "session_count": 5,
    "observation_count": 4231,
    "function_data": {
      "function:global:GetPetAbility": {
        "call_count": 87,
        "argument_patterns": [
          {
            "position": 1,
            "observed_types": ["integer"],
            "frequency": [87],
            "type_confidence": 0.99
          }
        ],
        "return_type_observed": "table",
        "return_type_confidence": 0.99,
        "state_changes": ["PetAbility state modified"],
        "related_functions": ["SetPetAbility", "UpdatePetAbility"],
        "edge_discoveries": ["GetPetAbility calls SetPetAbility"]
      }
    },
    "edge_data": {
      "event_triggers": [
        {
          "event": "COMBAT_START",
          "handlers": ["function:global:OnCombatStart", "function:global:UpdateCombatUI"],
          "observations": 12
        }
      ]
    }
  }
}
```

---

## 5. Integration with Confidence System

### 5.1 Evidence Weighting

**Deep_analysis confidence weights (current):**
```
direct_observation: 0.95,
runtime_observed: 0.99,  // ← Runtime evidence highest
explicit_annotation: 1.0,  // Perfect but rare
call_site: 0.75,
static_analysis: 0.70,
pattern_match: 0.60,
inferred_from_usage: 0.55,
heuristic: 0.50
```

**Score Calculation with Runtime Evidence:**
```go
// From deep_analysis/confidence.go
finalScore := 0.0

// Static evidence: 70-80% likely
staticScore := 0.75

// Runtime evidence: 99% definitive
if hasRuntimeObservation {
  // Weight heavily toward runtime observation
  finalScore = (staticScore * 0.3) + (0.99 * 0.7)  // Result: ~95%
  
  // Or replace entirely if high coverage
  if runtimeObservationCount >= 10 {
    finalScore = 0.99  // High confidence from repeated observation
  }
}
```

### 5.2 Updated Graph Data

**Before runtime evidence (Priority 1):**
```json
{
  "from": "function:global:BroadcastEvent",
  "to": "xml_handler:OnHidden",
  "type": "bound_from_xml",
  "confidence_score": 72,
  "evidence_sources": ["static_analysis", "pattern_match"]
}
```

**After runtime evidence (Priority 4):**
```json
{
  "from": "function:global:BroadcastEvent",
  "to": "xml_handler:OnHidden",
  "type": "bound_from_xml",
  "confidence_score": 99,
  "evidence_sources": ["runtime_observed", "static_analysis", "pattern_match"],
  "runtime_observation_count": 8,
  "rationale": "Observed 8 times during gameplay in event handling context"
}
```

---

## 6. Implementation Roadmap

### Phase 4A: Probe Framework (Week 1)
**Scope:** Create base instrumentation system

**Deliverables:**
1. Create `ProbeManager` Lua module
   - Hook installation system
   - Event capture infrastructure
   - Observation storage API
   - File: `tools/instrumentation/probe_manager.lua`

2. Create `FunctionProbe` 
   - Function entry/exit detection
   - Argument capture
   - Return value recording
   - File: `tools/instrumentation/function_probe.lua`

3. Create `EventProbe`
   - Event registration interception
   - Handler tracking
   - Event flow recording
   - File: `tools/instrumentation/event_probe.lua`

4. Create observation logger
   - Write observations to file
   - Compress for storage
   - File: `tools/instrumentation/observation_logger.lua`

**Testing:** Manual verification that probes fire and log data

### Phase 4B: Decorator Addon (Week 2)
**Scope:** Package probes as installable addon

**Deliverables:**
1. Create `InstrumentationProbe` addon
   - Minimal UI
   - Enable/disable toggle
   - Observation export button
   - File: `InstrumentationProbe/InstrumentationProbe.lua`
   - File: `InstrumentationProbe/InstrumentationProbe.xml`
   - File: `InstrumentationProbe/InstrumentationProbe.mod`

2. Test with target addons
   - AdvancedPetAssist
   - QuickWarReport
   - Sample output verification

### Phase 4C: Evidence Aggregation (Week 3)
**Scope:** Convert raw observations to confidence boosters

**Deliverables:**
1. Create aggregation script
   - Read observation files (multiple sessions)
   - Merge and deduplicate
   - Calculate statistics
   - File: `tools/instrumentation/aggregate_evidence.go`

2. Create boost generator
   - Input: aggregated evidence
   - Output: JSON with inferred types + confidence
   - File: `tools/instrumentation/generate_evidence_boost.go`

3. Modify deep_analysis to accept boost
   - Load runtime evidence
   - Apply confidence adjustments
   - Update final graph
   - File: `deep_analysis/runtime_evidence.go`

### Phase 4D: Integration & Deployment (Week 4)
**Scope:** Wire runtime evidence into generation pipeline

**Deliverables:**
1. Modify `semantic.go` to load runtime evidence
2. Update confidence calculation to use runtime data
3. Regenerate graph with runtime-boosted confidence
4. Document addon setup and usage

---

## 7. Success Criteria

### Probe Functionality
- ✅ Probes install without crashing addon
- ✅ Observations captured for 100+ function calls
- ✅ Argument types recorded accurately
- ✅ Return types matched to static analysis results
- ✅ Event handlers properly tracked

### Evidence Quality
- ✅ >90% of observations decode successfully
- ✅ No false positives in type inference
- ✅ Event chains reconstructible from observations
- ✅ UI operation sequences captured correctly

### Confidence Improvement
- ✅ Edge confidence increases from 70-80% to 95%+
- ✅ Return type confidence increases from 60-80% to 99%
- ✅ Argument role confidence increases from 50-70% to 98%+
- ✅ New edge types discovered (low-frequency relationships)

### Integration
- ✅ Runtime evidence loadable by semantic.go
- ✅ Confidence scores recalculated with runtime boost
- ✅ Final graph JSON includes runtime evidence indicators
- ✅ Documentation explains confidence sources

---

## 8. Risks & Mitigations

| Risk | Probability | Impact | Mitigation |
|------|-----------|--------|-----------|
| Probes impact addon performance | Medium | High | Use sampling, decorator addon, optional |
| Observation format evolution | Low | Medium | Version obs format, handle upgrades |
| Privacy/data concerns | Low | High | Local capture only, user opt-in |
| Coverage gaps | Medium | Medium | Multiple test characters, iterate |
| False positives | High | High | Validation against source, expert review |

---

## 9. Comparison: Static vs. Runtime

### Capability Matrix

| Capability | Static Analysis | Runtime Probe | Winner |
|-----------|-----------------|---------------|--------|
| Completeness | ~70% coverage | ~95% coverage | Runtime |
| Confidence | 50-95% | 99%+ | Runtime |
| Speed | Immediate | Requires execution | Static |
| Code changes | No | Minimal (decorator) | Static |
| Precision | Good | Excellent | Runtime |
| Edge discovery | 70% | 98% | Runtime |
| Dynamic behavior | Poor | Excellent | Runtime |
| Initial setup | Easy | Medium | Static |

**Recommendation:** Use both together
- Static: Fast baseline (5 min to graph)
- Runtime: Confidence boost (add proof in 1 hour of gameplay)

---

## 10. Example: Complete Workflow

### Workflow Step-by-Step

**1. Initial State**
```
Developer wants API docs with high-confidence edge types
Current state: ~70% confidence on edges, some relationships missing
```

**2. Run Static Analysis (Priority 1-3)**
```bash
docker compose run --rm api-doc-gen generate platform docs/addon-api docs/war-api
# Result: api_graph.json with 70-85% confidence edges
# Time: ~2 minutes
```

**3. Install Instrumentation Probe**
```
Copy InstrumentationProbe addon to game directory
Launch game, enable probe in addon settings
Play for 1-2 hours in various content
```

**4. Export Observations**
```
Click "Export Evidence" button in InstrumentationProbe
File: evidence_capture_session_1.json (stored locally)
Observations captured: 847 function calls, 34 events, 156 UI ops
```

**5. Aggregate Evidence**
```bash
go run tools/instrumentation/aggregate_evidence.go \
  --input evidence_capture_session_1.json \
  --output evidence_aggregated.json
# Result: Merged statistics across all observations
```

**6. Generate Runtime Boost**
```bash
go run tools/instrumentation/generate_evidence_boost.go \
  --evidence evidence_aggregated.json \
  --output runtime_boost.json
# Result: Confidence adjustments for edges/returns/args
```

**7. Regenerate Graph with Runtime Evidence**
```bash
docker compose run --rm api-doc-gen generate platform \
  docs/addon-api \
  docs/war-api \
  --runtime-evidence runtime_boost.json
# Result: api_graph.json with 95%+ confidence + runtime indicators
# Time: ~2 minutes
```

**8. Publish Enhanced Documentation**
```
Final graph includes:
- Original 1,247 edges (70-85% confidence)
- Enhanced with runtime evidence (99% confidence)
- 34 new edges discovered only in runtime (98% confidence)
- ~200 return types inferred (95%+ confidence)
- ~450 arguments types inferred (95%+ confidence)
```

---

## 11. Files to Create

**New files (Phase 4A-D):**
1. `tools/instrumentation/probe_manager.lua` — Probe infrastructure
2. `tools/instrumentation/function_probe.lua` — Function interception
3. `tools/instrumentation/event_probe.lua` — Event tracking
4. `tools/instrumentation/observation_logger.lua` — Log storage
5. `tools/instrumentation/aggregate_evidence.go` — Evidence merging
6. `tools/instrumentation/generate_evidence_boost.go` — Confidence adjustment
7. `InstrumentationProbe/InstrumentationProbe.lua` — Addon main
8. `InstrumentationProbe/InstrumentationProbe.xml` — UI definition
9. `InstrumentationProbe/InstrumentationProbe.mod` — Addon manifest
10. `deep_analysis/runtime_evidence.go` — Runtime evidence handler
11. `PRIORITY_4_EXECUTION_STATUS.md` — Implementation progress tracker

---

## 12. Summary

Priority 4 completes the data quality initiative by adding runtime evidence layer:

- **Priorities 1-3:** 50-95% confidence via static analysis, covers ~70% of relationships
- **Priority 4:** 99%+ confidence via runtime observation, captures dynamic behavior
- **Combined:** Highest-confidence API reference documentation in Warhammer Online community

**Next:** Begin Phase 4A implementation of probe framework.


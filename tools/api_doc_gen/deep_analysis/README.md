# Deep Analysis Module

## Overview

The `deep_analysis` module provides evidence-driven inference engines for improving the WAR API reference data quality. It discovers missing relationships, infers function return types and argument roles, and calculates confidence scores based on actual evidence.

## Components

### 1. Evidence Tracking (`evidence.go`)

**Purpose:** Central evidence accumulation and confidence calculation system.

**Key Types:**
- `Evidence` — Represents a collection of observations about a symbol with full provenance
  - `Sources[]` — Where the evidence came from (file, line, type, description)
  - `Count` — Number of independent observations
  - `Patterns[]` — Identified patterns
  - `Weight` — Calculated importance score
  
- `EvidenceSource` — Single observation point
  - `File` — Source file path
  - `Line` — Line number
  - `Type` — Type of evidence (call_site, return_location, xml_binding, etc.)
  - `Description` — Human-readable details

- `ConfidenceLevel` — Confidence assessment with explanation
  - `Score` — 0-100 integer
  - `Level` — HIGH, MEDIUM, LOW, UNCERTAIN
  - `Basis` — Why we have this confidence
  - `Caveats` — Known limitations

- `Inference` — Complete inferred fact with full tracking
  - `Subject` — What we're inferring about
  - `Property` — What property (return_type, etc.)
  - `Value` — The inferred value
  - `Evidence` — Supporting evidence
  - `Confidence` — Confidence assessment
  - `Observed` — Direct observation vs inferred
  - `Uncertain` — If too uncertain to include

**Key Functions:**
- `NewEvidence()` — Create evidence with initial source
- `(e *Evidence) AddSource()` — Add another observation
- `(e *Evidence) AddPattern()` — Mark discovered pattern
- `(e *Evidence) UpdateWeight()` — Recalculate importance
- `(e *Evidence) UniqueSourceLocations()` — Get unique file:line references
- `(e *Evidence) SourcesByType()` — Group evidence by type
- `ScoreConfidence()` — Calculate confidence from evidence

### 2. Edge Inference (`edge_inference.go`)

**Purpose:** Discover relationship patterns between symbols using static analysis.

**Key Types:**
- `FunctionAnalysis` — Extracted information about a function
  - `IsInitFunction` — Called at addon initialization
  - `IsUpdateFunction` — Called at periodic update
  - `IsEventHandler` — Handles specific event
  - `AccessedGlobals` — SystemData, GameData access
  - `CallsSites[]` — Functions it calls
  - `UIOperations[]` — UI manipulation operations
  - `AssignmentTargets[]` — What variables it assigns to

- `EdgeInference` — Analysis engine
  - `AllFunctions` — Analyzed function repository
  - `CallGraph` — Who calls whom
  - `GlobalAccess` — Data access patterns
  - `EventChains` — Event handler mappings

**Key Methods for Discovering Missing Edge Types:**
- `InferReadsSystemData()` — Functions accessing SystemData parameters
- `InferReadsGameData()` — Functions accessing GameData
- `InferUpdatesUI()` — UI refresh/update patterns
- `InferTogglesVisibility()` — Show/Hide patterns (SetShowing, Show, Hide)
- `InferUpdatesLayout()` — Layout modification (SetAnchor, SetSize, SetPosition)
- `InferInitializes()` — Init-phase detection
- `InferRefreshes()` — Update/refresh phase detection
- `InferCommonlyUsedWith()` — Function co-occurrence
- `InferAppearsInSameFlow()` — Event/execution flow membership

### 3. Return Type Inference (`returns.go`)

**Purpose:** Infer function return types from multiple evidence sources.

**Key Types:**
- `ReturnAnalysis` — Analysis of return types for a function
  - `ExplicitReturns[]` — Types from @returns annotations
  - `InferredTypes[]` — Types inferred from return statements
  - `LiteralReturns[]` — Actual values returned
  - `AssignmentContexts[]` — Variables assigned return values
  - `ComparisonContexts[]` — How values are compared
  - `Evidence[]` — Supporting evidence with sources

**Key Methods:**
- `AnalyzeReturns()` — Extract return statements and literal values from source
- `AnalyzeCallSites()` — Infer types from where function results are used
- `AnalyzeFieldAccess()` — Infer table structure from field access patterns
- `BestGuess()` — Return most likely type with confidence

**Evidence Types Tracked:**
- `explicit_annotation` — @returns comment (100% confidence)
- `literal` — Literal value like true/false/123/"string" (75% confidence)
- `assignment_target` — Variable assigned and used for type hints (60% confidence)
- `comparison` — Value used in boolean/numeric/string context (70-80% confidence)
- `call_result` — Result of function call (50% confidence)
- `table_expected` — Used with field access pattern (70% confidence)

### 4. Argument Inference (`arguments.go`)

**Purpose:** Infer function parameter types and roles from usage patterns.

**Key Types:**
- `ArgumentAnalysis` — Analysis of a single parameter
  - `ParameterName` — Parameter name
  - `Usage[]` — How it's used (indexing, calling, arithm,etic)
  - `TypeSignals[]` — Multiple type observations
  - `IsOptional` — Has default value
  - `IsVariadic` — Accepts variable arguments
  - `Evidence[]` — Call site evidence

- `TypeSignal` — Observation about parameter type
  - `Signal` — Type of signal (indexed, called, numeric, etc.)
  - `Confidence` — Confidence in this signal
  - `Details` — Explanation

**Key Methods:**
- `AnalyzeParameters()` — Extract parameter names from function signature
- `AnalyzeParameterUsage()` — Examine how parameters are used in function body
- `AnalyzeCallSites()` — Look at actual calls to infer argument types
- `InferArgumentRole()` — Determine semantic role (collection, object, callback, etc.)

**Type Signals Tracked:**
- `indexed` — Used with [key] operator → table/array
- `called` — Invoked with () → callback/function
- `numeric` — Used in arithmetic → number
- `has_methods` — Called with : syntax → object pattern
- `truthy_check` — Used in if statement → flag/boolean
- `array_like` — Iterated with ipairs() → array
- `table_like` — Iterated with pairs() → table/dictionary
- `has_fields` — Accessed with . operator → object/table
- `has_length` — Used with # operator → table/string/array
- `comparable_to_X` — Compared with specific type

### 5. Confidence Scoring (`confidence.go`)

**Purpose:** Calculate evidence-weighted confidence scores for all inferences.

**Key Types:**
- `ConfidenceScorer` — Scoring engine with evidence weights

**Key Methods:**
- `ScoreEvidence()` — Calculate confidence from evidence collection
- `ScoreReturnTypeConfidence()` — Score return type inference
- `ScoreArgumentTypeConfidence()` — Score argument inference
- `ScoreEdgeConfidence()` — Score relationship inference
- `CombineConfidences()` — Merge multiple confidence scores
- `AggregateEvidence()` — Combine evidence from multiple sources
- `ShouldIncludeInference()` — Determine if inference is reliable enough

**Evidence Weights (Configurable):**
- Direct observation: 0.95 (95%)
- Runtime observed: 0.99 (99%)
- Explicit annotation: 1.0 (100%)
- Call site: 0.75 (75%)
- Static analysis: 0.70 (70%)
- Pattern match: 0.60 (60%)
- Inferred from usage: 0.55 (55%)
- Heuristic: 0.50 (50%)

**Confidence Levels:**
- **DEFINITE** (95-100%) — Definite and should be trusted
- **VERY_HIGH** (85-94%) — Very likely correct
- **HIGH** (75-84%) — Probably correct
- **MEDIUM** (60-74%) — Reasonably likely, with some uncertainty
- **LOW** (45-59%) — Should be treated tentatively
- **VERY_LOW** (30-44%) — Speculative and uncertain
- **UNCERTAIN** (0-29%) — Insufficient evidence

### 6. Usage Clustering (`usage_clusters.go`)

**Purpose:** Identify natural groupings of related symbols through co-occurrence analysis.

**Key Types:**
- `UsageCluster` — Group of frequently-used-together symbols
  - `Name` — Cluster identifier
  - `Symbols[]` — Members
  - `Frequency` — Co-occurrence count
  - `Patterns[]` — Identified patterns
  - `Confidence` — 0-100 cluster strength

- `UsageClustering` — Co-occurrence analyzer
  - `CoOccurrenceMatrix` — Symbol pair frequency counts
  - `SymbolFrequency` — How often each symbol appears
  - `Clusters[]` — Discovered clusters

**Key Methods:**
- `RecordCoOccurrence()` — Track two symbols used together
- `AnalyzeSymbolCalls()` — Record co-occurrences from call patterns
- `BuildClusters()` — Discover natural groupings
- `GetRelatedSymbols()` — Find symbols related to given one
- `IdentifyPatterns()` — Detect cluster types (UI, events, data, utility)
- `FindCommonlyUsedWith()` — Generate edges from clustering

### 7. Integration (`integration.go`)

**Purpose:** Bridge deep_analysis with existing semantic graph building.

**Key Types:**
- `EnrichedEdge` — Edge with full confidence and evidence
  - `ConfidenceScore` — 0-100
  - `EvidenceCount` — Number of observations
  - `EvidenceSources[]` — file:line locations
  - `Rationale` — Explanation
  - `AnalysisMethod` — How inferred

- `EdgeEnricher` — Integrator joining all analysis engines

**Key Methods:**
- `EnrichEdge()` — Add confidence and evidence to basic edge
- `BuildMissingEdges()` — Discover edges from deep analysis
- `EnrichReturnTypes()` — Apply return inference to symbols
- `EnrichArgumentRoles()` — Apply argument inference to parameters
- `SummarizeNewEdges()` — Generate report of discoveries

## Workflow

### Priority 1: Graph Edge Extraction (Current Phase)

1. **Analyze function sources** — `EdgeInference.AnalyzeFunctionSource()`
2. **Detect patterns** — `InferReadsSystemData()`, `InferUpdatesUI()`, etc.
3. **Score confidence** — `ConfidenceScorer.ScoreEdgeConfidence()`
4. **Enrich edges** — `EdgeEnricher.EnrichEdge()` with confidence scores
5. **Export to graph** — Convert to platform/model.py `GraphEdge` with new fields

**Missing edge types to discover:**
- reads_systemdata (data access patterns)
- reads_gamedata (data access patterns)
- commonly_used_with (co-occurrence)
- appears_in_same_flow (sequence patterns)
- initializes (init phase)
- refreshes (update phase)
- updates_ui (UI operations)
- toggles_visibility (Show/Hide patterns)
- updates_layout (Position/Size operations)

### Priority 2: Return Types (Next Phase)

1. **Analyze return statements** — `ReturnInference.AnalyzeReturns()`
2. **Study call sites** — `AnalyzeCallSites()` for usage context
3. **Examine field access** — `AnalyzeFieldAccess()` for structure hints
4. **Score best guess** — `BestGuess()` with confidence

### Priority 3: Arguments (Next Phase)

1. **Extract parameters** — `ArgumentInference.AnalyzeParameters()`
2. **Study usage patterns** — `AnalyzeParameterUsage()` in function body
3. **Analyze call sites** — `AnalyzeCallSites()` to see what's passed
4. **Determine roles** — `InferArgumentRole()` for semantic meaning

### Priority 4: Live Game Extraction (Future)

Spec and probe scripts to extract actual runtime evidence from live game clients.

## Integration with Existing Platform

### Modifying `platform/semantic.go`

The `buildSemanticGraph()` function should:

1. Create `EdgeEnricher` after initial graph is built
2. Run analysis on collected function sources
3. Call `BuildMissingEdges()` to discover new relationships
4. Score existing edges with confidence
5. Update graph with enriched edges before serialization

```go
// In buildSemanticGraph():
enricher := deep_analysis.NewEdgeEnricher()

// Analyze collected sources
for _, fn := range functions {
    enricher.EdgeInference.AnalyzeFunctionSource(fn.Path, fn.Source, ...)
    enricher.ReturnInference.AnalyzeReturns(fn.Path, fn.Source)
    enricher.ArgumentInference.AnalyzeParameters(fn.Path, fn.Source)
}

// Discover missing edges
missingEdges := enricher.BuildMissingEdges()
for _, edge := range missingEdges {
    if enricher.ConfidenceScorer.ShouldIncludeInference(edge.ConfidenceScore, "edge") {
        // Add to graph
    }
}
```

### Data Model Updates

`platform/model.go` `GraphEdge` struct now includes:
- `ConfidenceScore` (int) — 0-100
- `EvidenceCount` (int) — number of observations
- `EvidenceSources` ([]string) — file:line references
- `Rationale` (string) — explanation

Previous fields (`Weight`, `Evidence`) retained for backward compatibility.

## Example Usage

```go
// Create enricher
enricher := deep_analysis.NewEdgeEnricher()

// Analyze a function
enricher.EdgeInference.AnalyzeFunctionSource("Module.Function", source, false, true, false, "")

// Enrich an edge
edge := enricher.EnrichEdge("Module.UpdateUI", "Window", "updates_ui",
    []string{"static_analysis:refresh_call", "static_analysis:invalidate_call"})

// Get confidence and rationale
fmt.Printf("Confidence: %d%%\n", edge.ConfidenceScore)
fmt.Printf("Rationale: %s\n", edge.Rationale)

// Score inferred return type
retType, conf, explanation := enricher.EnrichReturnTypes("Module.GetValue")
fmt.Printf("Return type: %s (%d%%)\n", retType, conf)

// Discover missing relationships
missingEdges := enricher.BuildMissingEdges()
for _, edge := range missingEdges {
    if conf >= 60 {  // Include medium+ confidence
        fmt.Printf("Found: %s → %s [%d%%]\n", edge.From, edge.To, edge.ConfidenceScore)
    }
}
```

## Confidence Interpretation Guide

| Score | Rating | Interpretation | Action |
|-------|--------|---|---|
| 95-100 | DEFINITE | Proven true | Include, mark as definite |
| 85-94 | VERY_HIGH | Very likely correct | Include, mark as highly confident |
| 75-84 | HIGH | Probably correct | Include, mark as confident |
| 60-74 | MEDIUM | Reasonably likely | Include, mark as tentative |
| 45-59 | LOW | Should be tentative | Include carefully, mark as uncertain |
| 30-44 | VERY_LOW | Speculative | Optional, mark as experimental |
| 0-29 | UNCERTAIN | Insufficient evidence | Exclude or mark as placeholder |

## Evidence Quality Standards

### High Quality Evidence (80%+ confidence)
- Direct observations (static analysis, test results)
- Multiple independent sources
- Explicit declarations (comments, type hints)
- Runtime observations
- Pattern matches from known conventions

### Medium Quality Evidence (50-80% confidence)
- Inferred from usage patterns
- Single or few sources
- Based on naming conventions
- Pattern matches that might have exceptions

### Low Quality Evidence (< 50% confidence)
- Heuristic guesses
- Single weak signal
- Contradictory evidence
- Speculative patterns

## Performance Considerations

- Analyze only modified functions incrementally
- Cache analysis results for reuse
- Limit clustering to relevant symbols
- Score only high-probability edges
- Run full analysis only on release builds

## Future Enhancements

1. **Runtime extraction** — Live game probes for definitive evidence
2. **Type annotations** — Parse embedded type hints more thoroughly
3. **Semantic patterns** — Learn common addon patterns for better inference
4. **Contradation resolution** — Handle conflicting evidence intelligently
5. **User feedback loop** — Incorporate manual corrections into confidence
6. **Cross-addon analysis** — Find patterns across multiple addons
7. **Trend analysis** — Track confidence improvements over time

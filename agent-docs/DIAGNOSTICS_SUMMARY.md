# Tightened Graph Diagnostics - Executive Summary
**Date: 2026-04-03 | Status: DIAGNOSTIC FRAMEWORK COMPLETE**

---

## Problem Identified

**User Observation**: UI shows isolated nodes  
**Previous Diagnostic**: "0 dangling nodes"  
**Contradiction Type**: Measurement mismatch between artifact and visualization

---

## Root Cause: Incomplete Diagnostic Definition

The previous "0 dangling nodes" result was checking **only edge validity**:
```
✓ All 409 edges have valid source/target nodeIDs
✓ No edges point to non-existent nodes
→ Conclusion: No dangling nodes
```

This was **insufficient** because it didn't check:
```
✗ Do all 315 nodes have at least one edge?
✗ What is the in-degree and out-degree of each node?
✗ Which nodes appear in relations.json but not in edges?
✗ How does the 300-node cap filter remove edges?
✗ Which nodes become isolated due to UI filtering?
```

---

## Five Categories of Isolated Nodes

### Category A: Truly Isolated (Graph Defect)
- **Definition**: Node with in-degree=0 AND out-degree=0
- **Location**: Exists in api_graph.json but appears in zero edges
- **Cause**: Never referenced, never references anything
- **Is it a bug?**: Possibly - may indicate incomplete generation or orphaned symbol
- **Example**: Unused constant or unreferenced data field

### Category B: Weakly Connected (Design Pattern)
- **Definition**: Node with (in-degree=0 XOR out-degree=0)
- **Subtypes**:
  - B1: Incoming-only — called/referenced but doesn't call anything
  - B2: Outgoing-only — calls things but isn't called
- **Cause**: Normal for sinks (events, data) and sources (entry points, setters)
- **Example**: OnButtonClick event (listeners call it, it doesn't call anything)

### Category C: Relations-Only Connected (Data Layer)
- **Definition**: Node has degree=0 in api_graph.json BUT appears in relations.json
- **Location**: In "participants" array of common_api_combinations
- **Cause**: Semantic grouping (conceptually related) not functional linking
- **Is it isolated?**: Technically yes in graph, but semantically connected in relations
- **Example**: XML UI element appearing in "Button + OnLButtonUp" combination

### Category D: Hidden by Filter Cap (UI Artifact)
- **Definition**: Node has degree>0 in api_graph but appears isolated in viewport
- **Cause**: neighbors filtered out by 300-node cap, edges can't render
- **Location**: UI filtering at line 506-511 of app.js
- **Is it dangling?**: No - connected in artifact but hidden by filtering
- **Example**: Function called by 10 others, but those 10 not in current 300-node view

### Category E: Lost Edges (Rendering Artifact)
- **Definition**: Edge exists in api_graph but doesn't appear in UI
- **Cause**: Both or one endpoint filtered out by cap
- **Effect**: Node X appears disconnected from node Y even though X → Y exists
- **Impact**: False isolation for visible nodes

---

## Diagnostic Framework Created

Three documents now define the complete diagnostic:

1. **TIGHTENED_DIAGNOSTICS.md**
   - Problem statement and analysis path
   - Five category definitions with precise criteria
   - Why previous diagnostic was incomplete

2. **DIAGNOSTIC_FRAMEWORK.md**  
   - Formal implementation checklist
   - Metrics to capture for each category
   - Expected findings and implications

3. **UI_RENDERING_PATH.md**
   - Complete trace from JSON through rendering
   - Code-level analysis of app.js
   - Why each isolation source occurs

---

## What's Now Known

### Artifact-Level Facts
```
api_graph.json:
  - 315 nodes (all types)
  - 409 edges (all "from"→"to" format)
  - Structural integrity: ✓ All edges valid

relations.json:
  - 60 common_api_combinations
  - Each has 2+ "participants"
  - Separate semantic layer from edges
```

### Filtering Logic (Verified from Code)
```
Step 1: Load api_graph.nodes and api_graph.edges
Step 2: Apply kind filter (user selection)
Step 3: Apply text search filter
Step 4: Cap to 300 nodes via capNodesDiverse()
        - Uses round-robin across types
        - No connectivity optimization
Step 5: Filter edges: (source in visible AND target in visible)
        - Removes edges where either endpoint filtered
Step 6: Calculate degree from visible edges only
Step 7: Render Cytoscape with visible nodes + visible edges
```

### Isolation Sources (Identified)
- **True isolated**: Nodes with degree=0 in api_graph
- **Weakly connected**: Nodes with asymmetric degree (B1/B2)  
- **Relations-only**: Degree=0 in graph but in relations
- **Filter-hidden**: Degree>0 in graph but edges removed by cap
- **Rendering artifact**: Edges that exist but don't display

---

## What's NOT Yet Measured

The framework is complete, but counts are still needed:

| Category | Count | Status |
|----------|-------|--------|
| A: Truly Isolated | ? | To measure |
| B1: Incoming-only | ? | To measure |
| B2: Outgoing-only | ? | To measure |
| C: Relations-only | ? | To measure |
| D: Hidden by cap | ? | To measure |
| E: Lost edges | ? | To measure |

---

## Why This Matters

### For Understanding
- "Dangling node" now has precise definition
- Mismatch between artifact and UI is explained
- Each isolation source has different cause and remedy

### For Fixing
- Category A needs evidence emission
- Category B needs visual distinction
- Category C needs relations UI layer
- Category D/E need filtering strategy change

### For Trust
- Previous "0 dangling" was technically true but misleading
- New diagnostics separate structural from rendering issues
- All isolation types documented and categorizable

---

## Next Steps: Strict Audit

To complete the measurement, execute:

```
FOR each node in api_graph.nodes:
  inDegree = count(edges where to == node.id)
  outDegree = count(edges where from == node.id)
  
  IF inDegree == 0 AND outDegree == 0:
    Category.A.add(node)
  ELSE IF inDegree == 0 OR outDegree == 0:
    Category.B.add(node, inDegree, outDegree)

FOR each node in Category.A:
  IF node.id in relations.participants:
    Category.C.add(node)  # Move to relations-only
  ELSE:
    keep in Category.A   # Truly isolated

FOR each node in visibleNodes (after 300-cap):
  degree_visible = count(visibleLinks touching this node)
  degree_full = inDegree[node.id] + outDegree[node.id]
  IF degree_visible == 0 AND degree_full > 0:
    Category.D.add(node)

FOR each edge in allLinks:
  IF source not in visibleIds OR target not in visibleIds:
    Category.E.add(edge)
```

Result: Complete categorization of all 315 nodes

---

## Rules Applied Throughout

- ✓ No guessing — only artifact inspection and code analysis
- ✓ Current generated artifacts only — Jan period, pipeline validated
- ✓ Precise definitions — before analysis, not after
- ✓ Five distinct categories — all isolation modes identified
- ✓ Code-backed — references to actual rendering logic (app.js lines)
- ✓ No fixes proposed — audit framework only
- ✓ Aligns with UI — explains what users actually see

---

## Documents Generated

| Document | Purpose | Location |
|----------|---------|----------|
| TIGHTENED_DIAGNOSTICS.md | Root cause analysis | /repo root |
| DIAGNOSTIC_FRAMEWORK.md | Implementation plan | /repo root |
| UI_RENDERING_PATH.md | Complete rendering trace | /repo root |

---

## Conclusion

**Problem**: "0 dangling nodes" contradicts visible isolated nodes

**Solution**: Tightened diagnostic framework with 5 precise categories

**Status**: Framework complete, ready for audit execution

**Next**: Run strict measurements to populate counts and generate categorized report

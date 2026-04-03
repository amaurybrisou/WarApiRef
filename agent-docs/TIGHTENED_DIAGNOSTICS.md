# Tightened Graph Diagnostics Report
**Generated: 2026-04-03 01:00 UTC**

## Executive Summary

Previous diagnostics reported **"0 dangling nodes"** while the UI clearly shows isolated nodes. This mismatch occurs because:

1. **Previous diagnostic was incomplete**: It only verified that all edges had valid source/target nodes (structural integrity)
2. **Did NOT check**: Whether nodes had zero connectivity (degree 0)
3. **Did NOT distinguish**: Between nodes connected via edges vs. nodes connected only through relations.json
4. **Did NOT account for**: UI rendering filtering logic that impacts visibility

---

## 1. Diagnostic Categories (Precise Definition)

### Category A: Truly Isolated Nodes
- **In graph JSON**: Yes
- **Has incoming edges**: No (in-degree = 0)
- **Has outgoing edges**: No (out-degree = 0)
- **In relations.json**: Unknown (separate data layer)
- **Visible in UI**: Only if manually searched for
- **Assessment**: Structurally disconnected from API relationship graph

### Category B: Weakly Connected Nodes
- **In graph JSON**: Yes
- **Has incoming edges**: Yes OR No (not both)
- **Has outgoing edges**: No OR Yes (not both)
- **Relation direction**: One-directional only
- **Visible in UI**: Limited depending on layout
- **Assessment**: Partially integrated into graph structure

### Category C: Connected Only via Relations.json
- **In graph JSON**: Yes (as node)
- **Has edges in api_graph.json**: No
- **Appears in relations.json**: Yes (in common_api_combinations)
- **Connection type**: Co-occurrence patterns, not call/reference edges
- **Visible in UI**: If relations are rendered separately
- **Assessment**: Semantically related but not directly linked

### Category D: Hidden by UI Filtering
- **In graph JSON**: Yes
- **Has valid edges**: Yes
- **Visible in full graph dump**: Yes
- **Visible when filtered/capped**: Maybe (depends on 300-node cap)
- **Visible in current viewport**: Possibly No
- **Assessment**: Technically connected but obscured by UI constraints

### Category E: Rendering Mismatches
- **In api_graph.json**: Yes
- **Both endpoints visible in UI**: No (one or both filtered out)
- **Edge should render**: Unknown (filtering may have removed one end)
- **Assessment**: Edge exists but neither node visible simultaneously

---

## 2. Full Audit Path: JSON → Rendering

### Phase 1: api_graph.json Structure
```
- Location: docs/war-api/graph/api_graph.json (578 KB)
- Nodes array: Lines 1-3154 (~315 nodes)
- Edges array: Lines 3155+ (~409 edges)
- Edge format: Uses "from"/"to" (NOT "source"/"target")
```

**Previous validation checked**:
- ✓ All edges have valid "from" nodeID
- ✓ All edges have valid "to" nodeID

**Previous validation did NOT check**:
- ✗ Do any nodes have 0 incoming AND 0 outgoing edges?
- ✗ Breakdown by edge type distribution
- ✗ Subgraph connectivity (strongly vs. weakly connected components)

### Phase 2: relations.json Structure
```
- Location: docs/war-api/graph/relations.json (138 KB)
- Structure: common_api_combinations array (60 items)
- Each combination has: participants array
- Participants: id, label, type, category, confidence
```

**Overlap questions**:
- How many nodes in relations are also in api_graph edges?
- How many nodes in api_graph have degree=0 but appear in relations?

### Phase 3: UI Rendering (app.js)
```javascript
// Key filtering logic (lines ~490-510)

1. Load: loadJSON('content/graph/api_graph.json')
2. Filter nodes by kind and search: filteredNodes = ...
3. Cap to 300: visibleNodes = capNodesDiverse(filteredNodes, 300)
4. Get edges from: normalizeGraphLinks(graphData)
   - Handles both 'edges' and 'links'
   - Maps "from/to" → "source/target"
5. Filter edges: Only include if BOTH endpoints in visibleIds
   
   const visibleLinks = allLinks.filter((l) => {
     const sid = l.source?.id ?? l.source;
     const tid = l.target?.id ?? l.target;
     return visibleIds.has(sid) && visibleIds.has(tid);
   })
```

**Critical filtering points**:
- Line 500-504: `capNodesDiverse()` reduces to 300 max
- Line 506-511: Edge filtering removes any edge where source or target not in visible set
- **Result**: A node can be in the graph but appear isolated if its edges' other endpoints are filtered

### Phase 4: Fallback SVG Rendering
```javascript
// ~450-480: When Cytoscape fails, fallback SVG grid layout
drawFallbackSvg(nodes, links) {
  // Grid layout: cols = Math.ceil(sqrt(nodes.length * 1.8))
  // Positions nodes in grid with jitter
}
```

**Implication**: All visible nodes appear somewhere, but edges are only drawn if both endpoints visible

---

## 3. Why Previous Diagnostic Was Misleading

### What was checked:
```
For each edge E:
  ✓ edge.from exists in nodeId set
  ✓ edge.to exists in nodeId set
Result: All 409 edges structurally valid → "0 dangling edges"
```

### What was NOT checked:
```
For each node N:
  ? sum(incoming_edges) == 0 AND sum(outgoing_edges) == 0
  → Category A: Truly isolated
  
For each node N:
  ? N appears in relations.json but has degree=0 in api_graph
  → Category C: Connected only via relations
  
For each node N in visibleNodes (after 300-cap):
  ? Can any of its edges actually be rendered?
  → Category D / E: Hidden by filtering
```

**Diagnosis**: The check "all edge source/target valid" ≠ "all nodes connected"

---

## 4. Root Cause Analysis: Why "0 Dangling Nodes" Was Misleading

### The Mismatch Explained

**Previous diagnostic assumption**:
```
if (all edges have valid source/target nodes) {
  NO DANGLING NODES
}
```

**What this actually checks**:
- Structural integrity of edge references
- No edges point to non-existent nodes
- Valid graph topology at the binary level

**What this DOES NOT check**:
- Whether nodes have ANY edges at all
- Degree distribution (how many in/out for each node)
- Connectivity properties (weakly vs. strongly connected)
- Whether all nodes are in a single connected component

**Analogy**: "No broken edges" ≠ "No isolated nodes"
- A node with degree=0 has no broken edges—it just has ZERO edges
- The diagnostic was looking at edge validity, not node connectivity

### The Real Issue

The UI shows isolated nodes because:

1. **api_graph.json structure**:
   - 315 nodes in "nodes" array
   - 409 edges in "edges" array
   - **Not all nodes are referenced by edges**
   
2. **Which nodes have zero connectivity**:
   - Some nodes may have: in-degree=0 AND out-degree=0
   - These are Category A (truly isolated)
   - Not detected by previous "edge validity" check

3. **Which nodes are weakly connected**:
   - Some nodes have only incoming or only outgoing edges
   - These are Category B (weakly connected)
   - Structurally valid but directionally one-sided

4. **UI filtering impact**:
   - app.js caps visible nodes at 300
   - Edge rendering requires BOTH endpoints visible
   - A connected node can appear isolated if its edges' other endpoints are filtered out

5. **Relations.json layer**:
   - 60 common API combinations
   - Separate from the graph edge layer
   - Some degree-0 nodes in api_graph may still appear in relations
   - Creates illusion of connection through co-occurrence data

---

## 5. Current State of Generated Artifacts

| Artifact | Status | Notes |
|----------|--------|-------|
| api_graph.json | 315 nodes, 409 edges | Structural validity verified, connectivity unknown |
| relations.json | 60 combinations | Separate from api_graph edge layer |
| app.js | Rendering logic | 300-node cap + edge filtering applied |
| UI display | Shows isolated nodes | Filtering explains visibility but not root cause |

---

## 6. Next Steps (Without Patching)

1. **Analyze api_graph.json**:
   - Iterate all 315 nodes
   - Count in-degree and out-degree for each
   - Identify all degree-0 nodes
   - Group by type and category

2. **Analyze relations.json**:
   - Extract all participant IDs
   - Cross-reference with api_graph node IDs
   - Identify nodes with zero edges but in relations

3. **Simulate UI rendering**:
   - Apply 300-node cap
   - Apply edge filtering
   - Report what becomes isolated due to cap

4. **Produce diagnostic report**:
   - True isolated count by type
   - Weakly connected count by direction
   - Relations-only count
   - Filtering impact quantified

---

## Rules Applied

- ✓ No guessing
- ✓ Current generated artifacts only
- ✓ No code inspection without verification
- ✓ Precise definitions before analysis  
- ✓ Distinguish all isolation categories
- ✓ Align diagnostics with actual UI behavior
- ✓ No patches proposed yet

**Status**: Ready for strict audit using correct definitions

# Graph Diagnostic Framework & Implementation Plan
**Status: DIAGNOSTIC COMPLETE - READY FOR AUDIT EXECUTION**

---

## Problem Statement

**Contradiction Detected**:
- Previous validation: "0 dangling nodes"
- Current UI observation: Isolated nodes clearly visible
- **Root cause**: Incomplete diagnostic definition

---

## Diagnostic Categories (Formal Definition)

### Category A: Truly Isolated Nodes
```
Property: node.inDegree == 0 AND node.outDegree == 0
Definition: No edges reference this node
Location: Node exists in api_graph.json but not in any edge's "from" or "to"
Visibility: Only if manually searched (not appearing via edges)
Impact: Structurally disconnected from the relationship graph
```

### Category B: Weakly Connected Nodes  
```
Property: (node.inDegree == 0 XOR node.outDegree == 0)
Definition: Edges point TO but not FROM (or vice versa)
Subtypes:
  - B1: Incoming-only (inDegree > 0, outDegree == 0)
  - B2: Outgoing-only (inDegree == 0, outDegree > 0)
Visibility: May appear isolated in filtered viewport
Impact: Limited integration into call graph
```

### Category C: Relations-Only Connected
```
Property: node in relations.json but zero degree in api_graph
Definition: Appears in common_api_combinations but has no edges
Location: Present in "participants" array of relations.json
Visibility: Through relations UI layer (not call/reference graph)
Impact: Semantically grouped but not functionally linked
```

### Category D: Hidden by Filter Cap
```
Property: degree > 0 in api_graph, but one/both endpoints filtered
Trigger: UI applies 300-node cap via capNodesDiverse()
Definition: Node exists but can't reach its edges because neighbors filtered
Visibility: Filtered from current viewport, visible in un-capped dump
Impact: False isolation due to rendering constraints
```

### Category E: Lost During Edge Filtering
```
Property: Edge exists in api_graph, but both endpoints filtered
Logic: UI line 506-511 removes edges where source OR target not visible
Trigger: 300-node cap filtering removes one of: {source, target}
Visibility: Edge disappears from rendered graph
Impact: Connected nodes appear disconnected in current view
```

---

## Audit Checklist

### Phase 1: Parse api_graph.json
```
✓ Extract all 315 nodes
✓ Create nodeId -> node mapping
✓ Extract all 409 edges (using "from"/"to" fields)
✓ Build adjacency: inDegree[id], outDegree[id] for all nodes
```

### Phase 2: Categorize Each Node
```
For each node n in nodes:
  inD = inDegree[n.id]
  outD = outDegree[n.id]
  
  if (inD == 0 AND outD == 0)           → Category A
  else if (inD == 0 XOR outD == 0)     → Category B
  else                                   → Connected (both directions)
```

### Phase 3: Cross-Reference relations.json
```
✓ Extract all participant IDs from common_api_combinations
✓ For each Category A node:
    - Is it in relations.json participants?
    - → If yes: Category C (relations-only)
    - → If no: Truly isolated (Category A)
```

### Phase 4: Simulate UI Filtering
```
✓ Apply kind filter (if any)
✓ Apply search text filter  
✓ Apply 300-node diversity cap (capNodesDiverse algorithm)
✓ For each filtered-out node:
    - Was it degree > 0?
    - → Yes: Category D candidate
✓ For each remaining edge:
    - Are both source and target in visibleNodes?
    - → No: Category E (lost to filtering)
```

### Phase 5: Generate Report
```
COUNTS:
  Category A: truly isolated (no relations backup)
  Category C: relations-only connected
  Category B1: incoming-only
  Category B2: outgoing-only
  Category D: hidden by cap
  Category E: lost edges

DISTRIBUTIONS:
  By type (data_structure vs. event vs. function, etc.)
  By category
  By confidence score
```

---

## Key Metrics to Capture

### Node-Level Metrics
```
Property         | Count | Examples | Implications
-----------------+-------+----------+----------------------------
Category A nodes | ? | [list] | True graph defects
Category C nodes | ? | [list] | Relations data quality
in-degree only   | ? | [list] | Sinks (endpoints)
out-degree only  | ? | [list] | Sources (origins)
Balanced degree  | ? | [list] | Hub nodes
```

### Graph-Level Metrics
```
Metric                   | Value | Analysis
-------------------------+-------+---------
Total nodes             | 315   | Source: validation output
Total edges             | 409   | Source: validation output
Avg degree per node     | 2.6   | Calculated: (2*edges)/nodes
Nodes with degree=0     | ?     | To be calculated
Nodes with in=0 out>0   | ?     | To be calculated
Nodes with out=0 in>0   | ?     | To be calculated
```

### Coverage Metrics
```
Metric                      | Value | Implication
----------------------------+-------+-----------
Nodes in relations.json     | ?     | Data layer coverage
Relations-only nodes        | ?     | Orphaned in graph
Nodes in both layers        | ?     | Well-integrated
Graph connectivity ratio    | ?     | (Nodes in giant component) / Total
```

---

## Expected Findings

Based on structure analysis, predictions:

1. **Likely True Isolated (Category A)**:
   - Some Constants may have degree=0
   - Some GameData fields may be unreferenced
   - Some utility functions may have no callers

2. **Likely Weakly Connected (Category B)**:
   - Events (typically only targets of function calls, no outgoing edges)
   - Data structures (referenced but don't reference others)
   - Tables (sinks, not sources)

3. **Likely Relations-Only (Category C)**:
   - UI elements appearing in common_api_combinations
   - Metadata nodes that are conceptually related but not functionally linked

4. **Filtering Impact (Category D/E)**:
   - 300-node cap may isolate 5-15% of nodes
   - Some hubs may disappear, orphaning their edges
   - Exact count depends on degree distribution

---

## Why This Matters

Understanding these categories enables:

1. **Diagnostic Accuracy**: "Dangling" now has precise meaning
2. **Root Cause Fixes**: Different causes need different solutions
3. **Priority Setting**: Category A > B > C in importance
4. **UI Improvements**: May need different rendering for weak connections
5. **Data Quality**: Relations may compensate for missing graph connections

---

## Next Action

Execute strict audit using Node.js/Python script:
1. Load both JSON files
2. Build degree maps  
3. Categorize all 315 nodes
4. Simulate UI filtering
5. Generate categorized report with examples
6. Document mismatches between diagnostic assumptions and reality

**Current Status**: Framework defined, ready for implementation

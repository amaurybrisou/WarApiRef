# UI Rendering Path Analysis
**Complete trace from JSON artifacts to screen**

---

##  1. Data Sources (Generated Artifacts)

### docs/war-api/graph/api_graph.json
```json
{
  "generated_at": "2026-04-03T00:59:34Z",
  "nodes": [
    { "id": "...", "type": "...", "label": "...", ... },
    ...  // 315 total nodes
  ],
  "edges": [
    { "from": "source-id", "to": "target-id", "type": "...", ... },
    ...  // 409 total edges
  ]
}
```

**Key observation**: Uses "from"/"to" (NOT "source"/"target")

### docs/war-api/graph/relations.json
```json
{
  "common_api_combinations": [
    {
      "participants": [
        { "id": "ui:Label", ... },
        { "id": "xml_event:OnHyperLinkRButtonUp", ... }
      ],
      ...
    },
    ...  // 60 total combinations
  ]
}
```

**Key observation**: Separate from api_graph edges, contains semantic pairings

---

## 2. App.js Loading & Normalization

### Step 2A: Load JSON (Line 326-340)
```javascript
async function loadGraph() {
  const candidates = [
    "content/graph/api_graph.json",
    "graph/api_graph.json"
  ];
  // Try both paths, assign to global graphData
  graphData = await loadJSON(path);
}
```

**Output**: Global `graphData` object with nodes[] and edges[]

### Step 2B: Normalize Edges (Line 389-400)
```javascript
function normalizeGraphLinks(data) {
  if (!data) return [];
  if (Array.isArray(data.links) && data.links.length > 0) return data.links;
  if (Array.isArray(data.edges) && data.edges.length > 0) {
    return data.edges.map((e) => ({
      source: e.source?.id ?? e.source ?? e.from,         // ← Handles both formats
      target: e.target?.id ?? e.target ?? e.to,           // ← Handles both formats
      type: e.type || "related",
      weight: e.weight || 1,
    }));
  }
  return [];
}
```

**Key mapping**:
- Input: `edge.from` / `edge.to`
- Output: `link.source` / `link.target`

**Output**: Normalized array of links with {source, target, type, weight}

---

## 3. Drawing Graph (Line 481+)

### Step 3A: Apply Filters (Line 497-510)
```javascript
function drawGraph(filterText, filterKind) {
  // 1. Filter nodes by kind and search text
  const filteredNodes = graphData.nodes.filter((n) => {
    if (!matchesKind(n.type, filterKind)) return false;  // Filter by type
    const label = String(n.label || n.id || "").toLowerCase();
    if (filterText && !label.includes(filterText)) return false;  // Search filter
    return true;
  });

  // 2. Cap to 300 diverse nodes
  const visibleNodes = capNodesDiverse(filteredNodes, 300);
  
  // 3. Get edges (normalized from api_graph)
  const allLinks = normalizeGraphLinks(graphData);
  
  // 4. CRITICAL FILTERING: Only render edges where BOTH endpoints visible
  const visibleIds = new Set(visibleNodes.map((n) => n.id));
  const visibleLinks = allLinks.filter((l) => {
    const sid = l.source?.id ?? l.source;       // Extract source ID
    const tid = l.target?.id ?? l.target;       // Extract target ID
    return visibleIds.has(sid) && visibleIds.has(tid);  // ← Only if BOTH in visibleIds
  }).slice(0, 800);
  
  // 5. Calculate node degrees based on VISIBLE links
  const degMap = new Map();
  visibleLinks.forEach((l) => {
    const s = l.source?.id ?? l.source;
    const t = l.target?.id ?? l.target;
    degMap.set(s, (degMap.get(s) || 0) + 1);
    degMap.set(t, (degMap.get(t) || 0) + 1);
  });
}
```

**Critical logic at line 506-511**:
```javascript
const visibleLinks = allLinks.filter((l) => {
  const sid = l.source?.id ?? l.source;
  const tid = l.target?.id ?? l.target;
  return visibleIds.has(sid) && visibleIds.has(tid);  // ← BOTH must be visible
})
```

**Implication**: An edge is rendered only if:
- Its source node is in `visibleNodes` (300-capped)
- AND its target node is in `visibleNodes` (300-capped)

**Node Isolation Mechanism**:
- If a node survives the cap but its neighbor doesn't
- The edge between them is deleted
- The node appears isolated

### Step 3B: capNodesDiverse Algorithm (Line 364-383)
```javascript
function capNodesDiverse(nodes, maxNodes) {
  if (nodes.length <= maxNodes) return nodes;
  
  const grouped = new Map();
  // Group nodes by type
  nodes.forEach((n) => {
    const bucket = grouped.get(n.type) || [];
    bucket.push(n);
    grouped.set(n.type, bucket);
  });
  
  const types = [...grouped.keys()];
  if (types.length <= 1) return nodes.slice(0, maxNodes);
  
  // Round-robin selection: evenly distribute types
  const result = [];
  let remaining = maxNodes;
  let active = types.filter((t) => (grouped.get(t) || []).length > 0);
  
  while (remaining > 0 && active.length > 0) {
    for (let i = 0; i < active.length && remaining > 0; i++) {
      const arr = grouped.get(active[i]);
      if (!arr || arr.length === 0) continue;
      result.push(arr.shift());  // Pop from front
      remaining--;
    }
    active = active.filter((t) => (grouped.get(t) || []).length > 0);
  }
  return result;
}
```

**Algorithm**:
- Groups nodes by type
- Round-robin selects from each type to form diverse set
- Result: 300 nodes distributed across all types

**Effect**: Biased toward including diverse types, not preserving connectivity

---

## 4. Rendering (Cytoscape or Fallback SVG)

### Step 4A: Build Cytoscape Elements (Line 532-570)
```javascript
const cyElements = [];

// Add nodes
visibleNodes.forEach((n) => {
  const deg = degMap.get(n.id) || 0;  // Degree from VISIBLE edges only
  const size = 18 + Math.min(deg * 2, 22);
  cyElements.push({
    group: "nodes",
    data: { id: n.id, label: ..., ... }
  });
});

// Add edges
visibleLinks.forEach((l, i) => {
  const sid = l.source?.id ?? l.source;
  const tid = l.target?.id ?? l.target;
  if (!visibleIds.has(sid) || !visibleIds.has(tid)) return;  // Extra safety check
  cyElements.push({
    group: "edges",
    data: { id: "e" + i, source: sid, target: tid, ... }
  });
});
```

**Key point**: Degree map only counts visible edges
- A node with degree=0 in full graph appears with degree=0 (correct)
- A node with degree>0 in full graph but isolated neighbors appears with degree=0 (filtered effect)

### Step 4B: Fallback SVG (Line 439-478)
```javascript
if (typeof cytoscape !== "function") {
  drawFallbackSvg(visibleNodes, visibleLinks);
  return;
}
```

**Fallback layout**:
- Grid-based positioning: `cols = Math.ceil(Math.sqrt(nodes.length * 1.8))`
- All 300 visible nodes appear, arranged in grid
- Edges still only rendered if both endpoints visible
- Result: Isolated nodes appear in grid with no lines

---

## 5. Why Isolated Nodes Appear in UI

### Scenario: Node N with degree > 0 in api_graph.json
```
api_graph.json:
  Nodes: [..., N, ..., neighbor1, ..., neighbor2, ...]
  Edges: [..., (N → neighbor1), (N → neighbor2), ...]

Phase 1 - Initial nodes loaded: 315 nodes
Phase 2 - Apply filters: Same, no filter applied (full set)
Phase 3 - Cap to 300: 
  - capNodesDiverse selected 300 nodes
  - N selected: YES
  - neighbor1 selected: NO (filtered out)
  - neighbor2 selected: NO (filtered out)

Phase 4 - Create visibleIds:
  - visibleIds = {id(N), ...other 299 nodes..., but NOT neighbors}

Phase 5 - Filter edges:
  - (N → neighbor1): neighbor1 not in visibleIds → REMOVE
  - (N → neighbor2): neighbor2 not in visibleIds → REMOVE

Phase 6 - Render N in Cytoscape/SVG:
  - N appears with degree=0 of visible edges
  - N appears isolated in visualization
  - But actually has degree>0 in api_graph.json
```

---

## 6. Why Diagnostic Was Incomplete

### What Was Checked
```javascript
// Previous diagnostic
for (each edge e in api_graph.edges) {
  if (nodeIds has e.from && nodeIds has e.to)
    → valid
  else
    → dangling
}
// Result: No dangling edges
```

### What Was NOT Checked
```javascript
// Missing diagnostic
for (each node n in api_graph.nodes) {
  inDegree[n.id] = count of edges where n.id == to
  outDegree[n.id] = count of edges where n.id == from
  
  if (inDegree[n.id] == 0 && outDegree[n.id] == 0)
    → Category A: Truly isolated
  
  else if (inDegree[n.id] == 0 || outDegree[n.id] == 0)
    → Category B: Weakly connected
}
```

---

## 7. Full Isolation Sources (Summary)

### Source 1: True Graph Isolation (Category A)
```
Structure: Node exists but has zero in/out edges
Trigger: api_graph.json has node with no edges referencing it
Example: An unused constant or unreferenced data field
Fix: Emit evidence for why node was included
```

### Source 2: Filtering (Category D)
```
Structure: Node exists, degree > 0, but edges filtered
Trigger: capNodesDiverse removes neighbors
Example: A function called by 5 functions, but those 5 filtered out
Fix: Adjust cap strategy or show multi-layer edges
```

### Source 3: Weak Connection (Category B)
```
Structure: Node has direction bias (in>0, out=0 or vice versa)
Trigger: Normal for events and data structures
Example: Event listened to but never fired, or called but never calls
Fix: Visual distinction for weak edges
```

### Source 4: Relations-Only (Category C)
```
Structure: Degree=0 in api_graph, but in relations.json
Trigger: Semantic grouping vs. functional linking
Example: XML element appears in common combinations but no edges
Fix: Expose relations layer separately when edges missing
```

---

## 8. Root Cause: Design Assumption

**Current UI Assumption**:
> "If something isn't connected in api_graph.json edges, it doesn't exist in the graph"

**Reality**:
1. Some nodes have degree=0 (truly isolated)
2. Some nodes have degree>0 but get filtered (hidden by cap)
3. Some nodes appear in relations but not edges (separate layer)
4. Filter priority chooses diversity over connectivity

---

## Conclusion

The "0 dangling nodes" diagnostic was checking **edge validity**, not **node connectivity**. 

The UI shows isolated nodes because:
1. Some nodes truly have degree=0
2. Some nodes are isolated by the 300-node cap filter
3. Edges are only rendered when BOTH endpoints pass filtering
4. Diagnostic definitions were incomplete

To fix diagnostics, must track:
- **Nodes**: in-degree, out-degree, category
- **Filters**: what gets removed and why
- **Visibility**: what appears in UI vs. artifact
- **Mismatch**: between what exists and what's visible

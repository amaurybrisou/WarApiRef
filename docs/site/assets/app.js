// --- DOM refs -----------------------------------------------------------------
const navRoot         = document.getElementById("nav");
const docRoot         = document.getElementById("doc");
const searchInput     = document.getElementById("search");
const resultsRoot     = document.getElementById("results");
const breadcrumb      = document.getElementById("breadcrumb");
const graphToggle     = document.getElementById("graphToggle");
const graphPanel      = document.getElementById("graphPanel");
const graphSvg        = document.getElementById("graphSvg");
const graphTooltip    = document.getElementById("graphTooltip");
const graphSearch     = document.getElementById("graphSearch");
const graphKindFilter = document.getElementById("graphKindFilter");
const graphReset      = document.getElementById("graphReset");
const mcpEndpoint     = document.getElementById("mcpEndpoint");
const mcpInitBtn      = document.getElementById("mcpInitBtn");
const mcpToolsBtn     = document.getElementById("mcpToolsBtn");
const mcpConsoleOut   = document.getElementById("mcpConsoleOutput");

// --- App state ----------------------------------------------------------------
const DEFAULT_PATH = "index.md";
let searchEntries  = [];
let graphData      = null;
let cyInstance     = null;   // Cytoscape instance

// --- Routing helpers ----------------------------------------------------------
function activePath() {
  return window.location.hash.replace(/^#/, "").trim() || DEFAULT_PATH;
}

function toContentPath(p) {
  return "content/" + p.replace(/^\/+/, "");
}

// --- Fetch helper -------------------------------------------------------------
async function loadJSON(path) {
  const res = await fetch(path);
  if (!res.ok) throw new Error(`${res.status} ${path}`);
  return res.json();
}

// --- MCP console -------------------------------------------------------------
function renderMcpConsole(payload) {
  if (!mcpConsoleOut) return;
  if (typeof payload === "string") {
    mcpConsoleOut.textContent = payload;
    return;
  }
  mcpConsoleOut.textContent = JSON.stringify(payload, null, 2);
}

async function postMcp(method, params) {
  if (!mcpEndpoint) return;
  const endpoint = mcpEndpoint.value.trim();
  if (!endpoint) {
    renderMcpConsole("Please provide an MCP endpoint URL.");
    return;
  }

  const request = {
    jsonrpc: "2.0",
    id: Date.now(),
    method,
    params: params || {},
  };

  renderMcpConsole(`Sending ${method} to ${endpoint} ...`);

  try {
    const response = await fetch(endpoint, {
      method: "POST",
      headers: { "content-type": "application/json" },
      body: JSON.stringify(request),
    });

    const text = await response.text();
    let decoded = null;
    try {
      decoded = JSON.parse(text);
    } catch {
      decoded = { status: response.status, raw: text };
    }

    renderMcpConsole(decoded);
  } catch (err) {
    renderMcpConsole({
      error: "request_failed",
      message: err && err.message ? err.message : String(err),
      endpoint,
      method,
    });
  }
}

// --- Navigation tree  (collapsible via <details>/<summary>) ------------------
const INDENT = ["", "nav-depth-1", "nav-depth-2", "nav-depth-3"];

function buildTree(items, parent, depth) {
  depth = depth || 0;
  items.forEach((item) => {
    const children = item.children || item.items;
    const hasKids  = Array.isArray(children) && children.length > 0;

    if (hasKids) {
      // Collapsible group
      const details = document.createElement("details");

      const summary = document.createElement("summary");
      summary.textContent = item.label;
      details.appendChild(summary);

      // If the group itself has a page, add a direct link inside
      if (item.path) {
        const self = makeNavLink(item.path, "-> " + item.label, INDENT[Math.min(depth, 3)]);
        details.appendChild(self);
      }

      buildTree(children, details, depth + 1);
      parent.appendChild(details);
    } else if (item.path) {
      parent.appendChild(makeNavLink(item.path, item.label, INDENT[Math.min(depth, 3)]));
    } else {
      // Label-only (separator)
      const lbl = document.createElement("div");
      lbl.className = "nav-group-title " + (INDENT[Math.min(depth, 3)] || "");
      lbl.textContent = item.label;
      parent.appendChild(lbl);
    }
  });
}

function makeNavLink(path, label, depthCls) {
  const a = document.createElement("a");
  a.href      = "#" + path;
  a.textContent = label;
  a.className = "nav-link" + (depthCls ? " " + depthCls : "");
  if (activePath() === path) a.classList.add("active");
  return a;
}

function updateActiveNavLink() {
  const cur = activePath();
  navRoot.querySelectorAll(".nav-link").forEach((a) => {
    a.classList.toggle("active", a.getAttribute("href").replace(/^#/, "") === cur);
  });
}

function generateBreadcrumb(path) {
  breadcrumb.innerHTML = "";
  
  // Root link
  const rootLink = document.createElement("a");
  rootLink.href = "#";
  rootLink.textContent = "API";
  breadcrumb.appendChild(rootLink);
  
  if (path === DEFAULT_PATH) return;
  
  // Display current page path as text-only breadcrumb (no intermediate links since they don't exist)
  const parts = path.replace(/\.md$/, "").split("/").filter(Boolean);
  parts.forEach((part) => {
    const sep = document.createElement("span");
    sep.className = "breadcrumb-sep";
    sep.textContent = "/";
    breadcrumb.appendChild(sep);
    
    const segment = document.createElement("span");
    segment.className = "breadcrumb-segment";
    segment.textContent = part.replace(/[-_]/g, " ");
    breadcrumb.appendChild(segment);
  });
}

function updateActiveNavLink() {
  const cur = activePath();
  navRoot.querySelectorAll(".nav-link").forEach((a) => {
    a.classList.toggle("active", a.getAttribute("href").replace(/^#/, "") === cur);
  });
}

async function loadNavigation() {
  navRoot.innerHTML = "<small>Loading...</small>";
  try {
    const tree  = await loadJSON("content/navigation/tree.json");
    navRoot.innerHTML = "";
    buildTree(tree.children || tree.items || [], navRoot, 0);
  } catch {
    navRoot.innerHTML = "<small>Navigation unavailable.</small>";
  }
}

// --- Search -------------------------------------------------------------------
async function loadSearchIndex() {
  try {
    const data = await loadJSON("content/search/index.json");
    searchEntries = Array.isArray(data.entries) ? data.entries : [];
  } catch {
    searchEntries = [];
  }
}

function runSearch(raw) {
  const q = raw.trim().toLowerCase();
  resultsRoot.innerHTML = "";
  if (!q) return;

  const terms = q.split(/\s+/).filter(Boolean);
  const hits = searchEntries
    .map((e) => {
      const hay = [e.title, e.summary || "", ...(e.keywords || []), e.content].join(" ").toLowerCase();
      let score = 0;
      terms.forEach((t) => {
        if (hay.includes(t)) score += 1;
        if (e.title.toLowerCase().includes(t)) score += 2;
      });
      return { e, score };
    })
    .filter((r) => r.score > 0)
    .sort((a, b) => b.score - a.score)
    .slice(0, 12);

  hits.forEach(({ e }) => {
    const a = document.createElement("a");
    a.href        = "#" + e.path;
    a.className   = "result-link";
    a.textContent = e.section ? `${e.title} - ${e.section}` : e.title;
    resultsRoot.appendChild(a);
  });
}

// --- Doc rendering ------------------------------------------------------------
async function loadDoc(path) {
  generateBreadcrumb(path);
  docRoot.innerHTML = "<p><small>Loading...</small></p>";

  const res = await fetch(toContentPath(path));
  if (!res.ok) {
    docRoot.innerHTML = `<p><strong>404</strong> - cannot load <code>${path}</code>.</p>`;
    return;
  }

  docRoot.innerHTML = marked.parse(await res.text());

  // Rewrite relative links to hash-based navigation
  docRoot.querySelectorAll("a[href]").forEach((a) => {
    const href = a.getAttribute("href");
    if (!href || href.startsWith("http") || href.startsWith("#")) return;
    if (href.endsWith(".md")) {
      const abs = new URL(href, "https://local/" + path).pathname.replace(/^\//, "");
      a.setAttribute("href", "#" + abs);
    } else if (href.endsWith(".json")) {
      const abs = new URL(href, "https://local/content/" + path).pathname.replace(/^\//, "");
      a.setAttribute("href", abs);
    }
  });
}

// --- Graph (Cytoscape.js) ----------------------------------------------------
async function loadGraph() {
  try {
    graphData = await loadJSON("content/graph/api_graph.json");
  } catch {
    graphData = null;
  }
}

const KIND_COLOR = {
  function:       "#c4673a",
  event:          "#2ea87e",
  data_structure: "#6b83c0",
  xml_element:    "#a08030",
  xml_handler:    "#a08030",
  ui_element:     "#a08030",
};
const KIND_COLOR_DEFAULT = "#7a7060";

function kindColor(k) { return KIND_COLOR[k] || KIND_COLOR_DEFAULT; }

function matchesKind(nodeType, selectedKind) {
  if (!selectedKind) return true;
  if (selectedKind === "xml") {
    return nodeType === "xml_element" || nodeType === "xml_handler" || nodeType === "ui_element";
  }
  return nodeType === selectedKind;
}

function capNodesDiverse(nodes, maxNodes) {
  if (nodes.length <= maxNodes) return nodes;
  const grouped = new Map();
  nodes.forEach((n) => {
    const bucket = grouped.get(n.type) || [];
    bucket.push(n);
    grouped.set(n.type, bucket);
  });
  const types = [...grouped.keys()];
  if (types.length <= 1) return nodes.slice(0, maxNodes);
  const result = [];
  let remaining = maxNodes;
  let active = types.filter((t) => (grouped.get(t) || []).length > 0);
  while (remaining > 0 && active.length > 0) {
    for (let i = 0; i < active.length && remaining > 0; i++) {
      const arr = grouped.get(active[i]);
      if (!arr || arr.length === 0) continue;
      result.push(arr.shift());
      remaining--;
    }
    active = active.filter((t) => (grouped.get(t) || []).length > 0);
  }
  return result;
}

function normalizeGraphLinks(data) {
  if (!data) return [];
  if (Array.isArray(data.links) && data.links.length > 0) return data.links;
  if (Array.isArray(data.edges) && data.edges.length > 0) {
    return data.edges.map((e) => ({
      source: e.source?.id ?? e.source ?? e.from,
      target: e.target?.id ?? e.target ?? e.to,
      type:   e.type || "related",
      weight: e.weight || 1,
    }));
  }
  return [];
}

const EDGE_LINE_COLOR = {
  commonly_used_with: "#c8a070",
  triggers:           "#2ea87e",
  structural_child_of:"#a08030",
};

function drawGraph(filterText, filterKind) {
  if (!graphData || !Array.isArray(graphData.nodes)) return;

  const q = (filterText || "").toLowerCase().trim();
  const k = filterKind || "";

  const filteredNodes = graphData.nodes.filter((n) => {
    if (!matchesKind(n.type, k)) return false;
    if (q && !n.label.toLowerCase().includes(q)) return false;
    return true;
  });

  const visibleNodes = capNodesDiverse(filteredNodes, 300);
  const allLinks     = normalizeGraphLinks(graphData);
  const visibleIds   = new Set(visibleNodes.map((n) => n.id));
  const visibleLinks = allLinks.filter((l) => {
    const sid = l.source?.id ?? l.source;
    const tid = l.target?.id ?? l.target;
    return visibleIds.has(sid) && visibleIds.has(tid);
  }).slice(0, 800);

  // Degree for label display.
  const degMap = new Map();
  visibleLinks.forEach((l) => {
    const s = l.source?.id ?? l.source;
    const t = l.target?.id ?? l.target;
    degMap.set(s, (degMap.get(s) || 0) + 1);
    degMap.set(t, (degMap.get(t) || 0) + 1);
  });
  const topLabels = new Set(
    [...degMap.entries()].sort((a, b) => b[1] - a[1]).slice(0, 30).map(([id]) => id)
  );

  // Build Cytoscape elements.
  const cyElements = [];
  visibleNodes.forEach((n) => {
    const deg  = degMap.get(n.id) || 0;
    const size = 18 + Math.min(deg * 2, 22);
    cyElements.push({
      group: "nodes",
      data: {
        id:         n.id,
        label:      topLabels.has(n.id) ? (n.label.length > 22 ? n.label.slice(0, 20) + "…" : n.label) : "",
        fullLabel:  n.label,
        type:       n.type,
        category:   n.category || n.type,
        confidence: n.confidence,
        summary:    n.summary || "",
        path:       n.path || "",
        color:      kindColor(n.type),
        size,
      },
    });
  });
  visibleLinks.forEach((l, i) => {
    const sid = l.source?.id ?? l.source;
    const tid = l.target?.id ?? l.target;
    if (!visibleIds.has(sid) || !visibleIds.has(tid)) return;
    cyElements.push({
      group: "edges",
      data: {
        id:     "e" + i,
        source: sid,
        target: tid,
        type:   l.type || "related",
        color:  EDGE_LINE_COLOR[l.type] || "#4a4030",
      },
    });
  });

  // Destroy previous instance.
  if (cyInstance) {
    cyInstance.destroy();
    cyInstance = null;
  }

  // Swap SVG for a div (Cytoscape needs a div container).
  const viewport = graphSvg.parentElement;
  let cyContainer = document.getElementById("cyContainer");
  if (!cyContainer) {
    cyContainer = document.createElement("div");
    cyContainer.id = "cyContainer";
    cyContainer.style.width  = "100%";
    cyContainer.style.height = "100%";
    cyContainer.style.background = "transparent";
    graphSvg.style.display = "none";
    viewport.appendChild(cyContainer);
  }
  cyContainer.style.display = "block";

  cyInstance = cytoscape({
    container: cyContainer,
    elements:  cyElements,
    style: [
      {
        selector: "node",
        style: {
          "width":              "data(size)",
          "height":             "data(size)",
          "background-color":   "data(color)",
          "label":              "data(label)",
          "font-size":          9,
          "font-family":        "Inter, system-ui, sans-serif",
          "color":              "#cfc8b8",
          "text-valign":        "top",
          "text-halign":        "center",
          "text-margin-y":      -3,
          "border-width":       1.2,
          "border-color":       "rgba(240,235,220,.45)",
          "cursor":             "pointer",
        },
      },
      {
        selector: "node:selected",
        style: {
          "border-width":   3,
          "border-color":   "#fff",
        },
      },
      {
        selector: "node.faded",
        style: { "opacity": 0.12 },
      },
      {
        selector: "edge",
        style: {
          "width":            1,
          "line-color":       "data(color)",
          "opacity":          0.38,
          "curve-style":      "bezier",
          "target-arrow-shape": "none",
        },
      },
      {
        selector: "edge.highlighted",
        style: {
          "width":   2,
          "opacity": 0.85,
        },
      },
      {
        selector: "edge.faded",
        style: { "opacity": 0.04 },
      },
    ],
    layout: {
      name:           "cose",
      animate:        false,
      nodeRepulsion:  8000,
      idealEdgeLength: 80,
      nodeOverlap:    10,
      padding:        20,
      randomize:      false,
    },
    userZoomingEnabled:    true,
    userPanningEnabled:    true,
    boxSelectionEnabled:   false,
    minZoom:               0.05,
    maxZoom:               8,
  });

  // -- Tooltip on mouseover ---------------------------------------------------
  cyInstance.on("mouseover", "node", (e) => {
    const d = e.target.data();
    graphTooltip.classList.remove("hidden");
    graphTooltip.innerHTML =
      `<strong>${d.fullLabel}</strong>` +
      `<div class="tt-type">${d.category}${d.confidence != null ? ` - ${d.confidence}% confidence` : ""}</div>` +
      (d.summary ? `<div class="tt-summary">${d.summary}</div>` : "") +
      (d.path ? `<div class="tt-type" style="margin-top:6px;color:#a08060">Click to open docs →</div>` : "");
  });

  cyInstance.on("mousemove", "node", (e) => {
    const r   = graphPanel.getBoundingClientRect();
    const pos = e.renderedPosition || { x: 0, y: 0 };
    const canvasRect = cyContainer.getBoundingClientRect();
    let left = (canvasRect.left - r.left) + pos.x + 14;
    const top  = (canvasRect.top  - r.top)  + pos.y - 10;
    if (left + 284 > r.width - 10) left -= 298;
    graphTooltip.style.left = left + "px";
    graphTooltip.style.top  = Math.max(4, top) + "px";
  });

  cyInstance.on("mouseout", "node", () => graphTooltip.classList.add("hidden"));

  // -- Click: highlight neighbours -------------------------------------------
  let selectedNodeId = null;

  function resetHighlight() {
    selectedNodeId = null;
    cyInstance.elements().removeClass("faded highlighted");
  }

  cyInstance.on("click", "node", (e) => {
    e.stopPropagation();
    const node = e.target;
    const id   = node.data("id");

    if (selectedNodeId === id) {
      resetHighlight();
      const path = node.data("path");
      if (path) window.location.hash = path;
      return;
    }

    selectedNodeId = id;
    const neighbourhood = node.closedNeighborhood();
    cyInstance.elements().addClass("faded");
    neighbourhood.removeClass("faded");
    neighbourhood.edges().addClass("highlighted").removeClass("faded");
  });

  cyInstance.on("click", (e) => {
    if (e.target === cyInstance) resetHighlight();
  });
}


// --- Event listeners ---------------------------------------------------------
window.addEventListener("hashchange", () => {
  updateActiveNavLink();
  loadDoc(activePath());
});

searchInput.addEventListener("input", (e) => runSearch(e.target.value));

function scrollGraphIntoView() {
  graphPanel.scrollIntoView({ behavior: "smooth", block: "start" });
}

graphToggle.addEventListener("click", () => {
  const hidden = graphPanel.classList.toggle("hidden");
  graphToggle.textContent = hidden ? "Graph" : "Close Graph";
  if (!hidden) {
    drawGraph(graphSearch.value, graphKindFilter.value);
    scrollGraphIntoView();
  }
});

graphSearch.addEventListener("input",       () => { if (!graphPanel.classList.contains("hidden")) drawGraph(graphSearch.value, graphKindFilter.value); });
graphKindFilter.addEventListener("change",  () => { if (!graphPanel.classList.contains("hidden")) drawGraph(graphSearch.value, graphKindFilter.value); });

graphReset.addEventListener("click", () => {
  graphSearch.value     = "";
  graphKindFilter.value = "";
  if (cyInstance) cyInstance.fit();
  drawGraph("", "");
});

if (mcpInitBtn) {
  mcpInitBtn.addEventListener("click", () => {
    postMcp("initialize", {
      protocolVersion: "2025-03-26",
      capabilities: {},
      clientInfo: { name: "war-api-docs-ui", version: "1.0.0" },
    });
  });
}

if (mcpToolsBtn) {
  mcpToolsBtn.addEventListener("click", () => {
    postMcp("tools/list", {});
  });
}

// --- Init ---------------------------------------------------------------------
(async function init() {
  await Promise.all([loadSearchIndex(), loadGraph(), loadNavigation()]);
  await loadDoc(activePath());
})();


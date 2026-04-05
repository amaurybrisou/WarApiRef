// --- DOM refs -----------------------------------------------------------------
const navRoot         = document.getElementById("nav");
const docRoot         = document.getElementById("doc");
const searchInput     = document.getElementById("search");
const resultsRoot     = document.getElementById("results");
const breadcrumb      = document.getElementById("breadcrumb");
const mcpPage         = document.getElementById("mcpPage");
const mcpPageBtn      = document.getElementById("mcpPageBtn");
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
const analysisToggle  = document.getElementById("analysisToggle");
const MCP_ROUTE       = "mcp";

// --- App state ----------------------------------------------------------------
const DEFAULT_PATH = "index.md";
let searchEntries  = [];
let graphData      = null;
let cyInstance     = null;   // Cytoscape instance

// --- Analysis toggle ---------------------------------------------------------
const ADVANCED_STORAGE_KEY = "war-api-show-advanced";
const LEGACY_ANALYSIS_STORAGE_KEY = "war-api-show-analysis";

// Exact H2 heading text (lower-cased) that belong to the internal analysis surface.
// Any section whose title matches is hidden by default and revealed via the toggle.
const ANALYSIS_HEADINGS = new Set([
  "confidence assessment",
  "evidence signals",
  "evidence summary",
  "binding resolution",
]);

// Exact H2 heading text (lower-cased) that are developer-facing XML structure content.
// These are intentionally kept visible in default mode.
const DEVELOPER_VISIBLE_XML_HEADINGS = new Set([
  "description",
  "common attributes",
  "common handlers",
  "xml event bindings",
  "common inherits",
  "common parent elements",
  "common structural child elements",
  "typical xml structure",
  "structural sub-elements",
  "recursive hierarchy",
]);

const INTERNAL_METRIC_COLUMNS = new Set(["score", "rationale"]);

function normalizeHeadingKey(value) {
  return String(value || "").trim().toLowerCase();
}

function classifyDocSectionHeading(value) {
  const headingKey = normalizeHeadingKey(value);
  if (ANALYSIS_HEADINGS.has(headingKey)) {
    return ["analysis-section", "internal-section"];
  }
  if (DEVELOPER_VISIBLE_XML_HEADINGS.has(headingKey)) {
    return ["developer-visible-section"];
  }
  return [];
}

function markInternalMetricColumns() {
  const tables = docRoot.querySelectorAll("table");
  tables.forEach((table) => {
    const headerRow = table.querySelector("thead tr") || table.querySelector("tr");
    if (!headerRow) return;

    const headerCells = Array.from(headerRow.querySelectorAll("th, td"));
    if (!headerCells.length) return;

    const metricColumnIndexes = [];
    headerCells.forEach((cell, index) => {
      const key = normalizeHeadingKey(cell.textContent);
      if (INTERNAL_METRIC_COLUMNS.has(key)) metricColumnIndexes.push(index);
    });
    if (!metricColumnIndexes.length) return;

    const rowList = table.querySelectorAll("tr");
    metricColumnIndexes.forEach((columnIndex) => {
      rowList.forEach((row) => {
        const rowCells = row.querySelectorAll("th, td");
        if (columnIndex < rowCells.length) {
          rowCells[columnIndex].classList.add("internal-metric-cell");
        }
      });
    });
  });
}

function applyAnalysisToggle(show) {
	document.body.classList.toggle("show-advanced", show);
  document.body.classList.toggle("show-analysis", show);
  if (analysisToggle) {
    analysisToggle.classList.toggle("secondary", show);
    analysisToggle.setAttribute("aria-pressed", show ? "true" : "false");
  }
  try { localStorage.setItem(ADVANCED_STORAGE_KEY, show ? "1" : "0"); } catch {}
}

// --- Routing helpers ----------------------------------------------------------
function activePath() {
  return window.location.hash.replace(/^#/, "").trim() || DEFAULT_PATH;
}

function toContentPath(p) {
  return "content/" + p.replace(/^\/+/, "");
}

function isMcpRoute(path) {
  return String(path || "").toLowerCase() === MCP_ROUTE;
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
const INDENT = ["", "nav-depth-1", "nav-depth-2", "nav-depth-3", "nav-depth-4"];
const STRUCTURAL_LABELS = new Set([
  "GLOBALS",
  "FUNCTIONS",
  "TABLES",
  "CONSTANTS",
]);

function isStructuralGroup(item, hasKids) {
  if (!hasKids) return false;
  const label = String(item.label || "").trim().toUpperCase();
  const path = String(item.path || "").trim().toLowerCase();
  if (!path) return true;
  if (STRUCTURAL_LABELS.has(label)) return true;
  return path.endsWith("/index.md");
}

function buildTree(items, parent, depth) {
  depth = depth || 0;
  items.forEach((item) => {
    const children = item.children || item.items;
    const hasKids  = Array.isArray(children) && children.length > 0;
    const depthCls = INDENT[Math.min(depth, 4)] || "";

    if (hasKids) {
      // Collapsible group
      const details = document.createElement("details");
      details.className = "nav-group" + (depthCls ? " " + depthCls : "");
      const structuralOnly = isStructuralGroup(item, hasKids);
      if (structuralOnly) details.classList.add("nav-structural");

      const summary = document.createElement("summary");
      summary.className = "nav-group-summary" + (depthCls ? " " + depthCls : "");
      summary.textContent = item.label;
      details.appendChild(summary);

      // If the group itself has a page, add a direct link inside
      if (item.path && !structuralOnly) {
        const self = makeNavLink(item.path, "Overview", depthCls + " nav-group-link");
        details.appendChild(self);
      }

      buildTree(children, details, depth + 1);
      parent.appendChild(details);
    } else if (item.path) {
      parent.appendChild(makeNavLink(item.path, item.label, depthCls));
    } else {
      // Label-only (separator)
      const lbl = document.createElement("div");
      lbl.className = "nav-group-title " + depthCls;
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

  makeDocSectionsCollapsible();
}

function makeDocSectionsCollapsible() {
  const nodes = Array.from(docRoot.childNodes).filter((node) => {
    return !(node.nodeType === Node.TEXT_NODE && !node.textContent.trim());
  });
  if (!nodes.length) return;

  const fragment = document.createDocumentFragment();
  let index = 0;

  while (index < nodes.length) {
    const current = nodes[index];
    const isHeading = current.nodeType === Node.ELEMENT_NODE && (current.tagName === "H1" || current.tagName === "H2");

    if (!isHeading) {
      fragment.appendChild(current);
      index += 1;
      continue;
    }

    const levelClass = current.tagName.toLowerCase();
    const details = document.createElement("details");
    details.className = `doc-section ${levelClass}`;

    const sectionClasses = classifyDocSectionHeading(current.textContent);
    if (sectionClasses.length) details.classList.add(...sectionClasses);

    const summary = document.createElement("summary");
    summary.className = "doc-section-summary";

    const title = document.createElement("span");
    title.className = "doc-section-title";
    title.textContent = current.textContent || "Section";

    const chevron = document.createElement("span");
    chevron.className = "doc-section-chevron";
    chevron.setAttribute("aria-hidden", "true");
    chevron.textContent = ">";

    summary.appendChild(title);
    summary.appendChild(chevron);
    details.appendChild(summary);

    const body = document.createElement("div");
    body.className = "doc-section-body";

    index += 1;
    while (index < nodes.length) {
      const nextNode = nodes[index];
      const nextIsHeading = nextNode.nodeType === Node.ELEMENT_NODE && (nextNode.tagName === "H1" || nextNode.tagName === "H2");
      if (nextIsHeading) break;
      body.appendChild(nextNode);
      index += 1;
    }

    details.appendChild(body);
    fragment.appendChild(details);
  }

  docRoot.innerHTML = "";
  docRoot.appendChild(fragment);
  markInternalMetricColumns();
}

// --- Graph (Cytoscape.js) ----------------------------------------------------
async function loadGraph() {
  const candidates = [
    "content/graph/api_graph.json",
    "graph/api_graph.json",
  ];

  graphData = null;
  for (const path of candidates) {
    try {
      graphData = await loadJSON(path);
      return;
    } catch {
      // Try next candidate path.
    }
  }
}

const KIND_COLOR = {
  function:       "#c4673a",
  event:          "#2ea87e",
  data_structure: "#6b83c0",
  xml_element:    "#a08030",
  xml_event:      "#3b8bb5",
  xml_handler:    "#a08030",
  ui_element:     "#a08030",
};
const KIND_COLOR_DEFAULT = "#7a7060";
const TEXT_NODE_COLOR_DARKER = "#8d6f2a";

function kindColor(k, node) {
  if (node && (node.id === "ui:Text" || String(node.label || "") === "Text")) {
    return TEXT_NODE_COLOR_DARKER;
  }
  return KIND_COLOR[k] || KIND_COLOR_DEFAULT;
}

function matchesKind(nodeType, selectedKind) {
  if (!selectedKind) return true;
  if (selectedKind === "xml") {
    return nodeType === "xml_element" || nodeType === "xml_event" || nodeType === "xml_handler" || nodeType === "ui_element";
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
    const links = data.edges.map((e) => ({
      source: e.source?.id ?? e.source ?? e.from,
      target: e.target?.id ?? e.target ?? e.to,
      type:   e.type || "related",
      weight: e.weight || 1,
    }));
    
    // Infer manipulates relationships: functions whose names start with UI element names
    // indicate they operate on that element type.
    if (Array.isArray(data.nodes)) {
      const elementsByName = new Map();
      const functionsByPrefix = new Map();
      
      // Build indexes
      data.nodes.forEach((n) => {
        if (n.type === "ui_element") {
          elementsByName.set(n.label, n);
        } else if (n.type === "function") {
          const label = String(n.label || "");
          // Extract prefix (e.g., "Window" from "WindowAddAnchor")
          for (const elemName of elementsByName.keys()) {
            if (label.startsWith(elemName)) {
              if (!functionsByPrefix.has(elemName)) {
                functionsByPrefix.set(elemName, []);
              }
              functionsByPrefix.get(elemName).push(n);
              break;
            }
          }
        }
      });
      
      // Create manipulates edges for matching pairs
      for (const [elemName, elementNode] of elementsByName) {
        const funcs = functionsByPrefix.get(elemName);
        if (!funcs) continue;
        funcs.forEach((func) => {
          const existingEdge = links.find(
            (l) => (l.source.id ?? l.source) === func.id && (l.target.id ?? l.target) === elementNode.id
          );
          if (!existingEdge) {
            links.push({
              source: func.id,
              target: elementNode.id,
              type: "manipulates",
              weight: 1,
            });
          }
        });
      }
    }
    
    return links;
  }
  return [];
}

const EDGE_LINE_COLOR = {
  commonly_used_with:  "#c8a070",
  triggers:            "#2ea87e",
  triggered_by:        "#2ea87e",
  structural_child_of: "#a08030",
  contains:            "#a08030",
  manipulates:         "#c4673a",
  calls:               "#6b83c0",
  reads_data:          "#6b83c0",
  writes_state:        "#c4673a",
  updates_ui:          "#a08030",
  bound_from_xml:      "#3b8bb5",
  inherited_from:      "#9b87a0",
};

const RENDER_OPTIONS = {
  edgeTypeFilter: localStorage.getItem("war-api-edge-types") || "",
  maxNodesCap:    parseInt(localStorage.getItem("war-api-node-cap") || "300", 10),
  maxLinksCap:    parseInt(localStorage.getItem("war-api-link-cap") || "800", 10),
  nodeSizeMin:    parseInt(localStorage.getItem("war-api-node-size-min") || "3", 10),
  nodeSizeMax:    parseInt(localStorage.getItem("war-api-node-size-max") || "32", 10),
};

function showGraphMessage(message) {
  if (!graphSvg) return;

  const viewport = graphSvg.parentElement;
  const cyContainer = document.getElementById("cyContainer");
  if (cyContainer) cyContainer.style.display = "none";

  graphSvg.style.display = "block";
  graphSvg.setAttribute("viewBox", "0 0 1000 600");
  graphSvg.innerHTML =
    `<rect x="0" y="0" width="1000" height="600" fill="#1a1812"></rect>` +
    `<text x="500" y="300" text-anchor="middle" fill="#cfc8b8" font-size="18" font-family="Inter, system-ui, sans-serif">${message}</text>`;

  if (viewport) viewport.appendChild(graphSvg);
}

function drawFallbackSvg(nodes, links) {
  if (!graphSvg) return;

  const viewport = graphSvg.parentElement;
  const cyContainer = document.getElementById("cyContainer");
  if (cyContainer) cyContainer.style.display = "none";

  graphSvg.style.display = "block";
  graphSvg.setAttribute("viewBox", "0 0 1200 700");

  const cols = Math.max(8, Math.ceil(Math.sqrt(nodes.length * 1.8)));
  const cellW = 1150 / cols;
  const rows = Math.max(1, Math.ceil(nodes.length / cols));
  const cellH = 650 / rows;
  const positions = new Map();

  nodes.forEach((n, idx) => {
    const col = idx % cols;
    const row = Math.floor(idx / cols);
    const jitterX = ((idx * 37) % 11) - 5;
    const jitterY = ((idx * 53) % 11) - 5;
    positions.set(n.id, {
      x: 25 + (col * cellW) + (cellW / 2) + jitterX,
      y: 25 + (row * cellH) + (cellH / 2) + jitterY,
    });
  });

  const degreeMap = buildDegreeMap(links);

  const edgesSvg = links.slice(0, 500).map((l) => {
    const sid = l.source?.id ?? l.source;
    const tid = l.target?.id ?? l.target;
    const s = positions.get(sid);
    const t = positions.get(tid);
    if (!s || !t) return "";
    const color = EDGE_LINE_COLOR[l.type] || "#4a4030";
    return `<line x1="${s.x}" y1="${s.y}" x2="${t.x}" y2="${t.y}" stroke="${color}" stroke-opacity="0.32" stroke-width="1"></line>`;
  }).join("");

  const nodesSvg = nodes.map((n) => {
    const p = positions.get(n.id);
    if (!p) return "";
    const label = String(n.label || n.id || "").replace(/[&<>\"]/g, (ch) => ({ "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;" }[ch]));
    const degree = degreeMap.get(n.id) || 0;
    const radius = fallbackRadiusFromDegree(degree);
    return (
      `<g>` +
      `<circle cx="${p.x}" cy="${p.y}" r="${radius}" fill="${kindColor(n.type, n)}"></circle>` +
      `<title>${label}</title>` +
      `</g>`
    );
  }).join("");

  graphSvg.innerHTML =
    `<rect x="0" y="0" width="1200" height="700" fill="#1a1812"></rect>` +
    edgesSvg +
    nodesSvg;

  if (viewport) viewport.appendChild(graphSvg);
}

function graphViewportHeightPx() {
  if (window.matchMedia && window.matchMedia("(max-width: 960px)").matches) {
    return 500;
  }
  return 700;
}

function buildDegreeMap(links) {
  const degreeMap = new Map();
  links.forEach((l) => {
    const s = l.source?.id ?? l.source;
    const t = l.target?.id ?? l.target;
    degreeMap.set(s, (degreeMap.get(s) || 0) + 1);
    degreeMap.set(t, (degreeMap.get(t) || 0) + 1);
  });
  return degreeMap;
}

function sortNodesByDegreeDesc(nodes, degreeMap) {
  return [...nodes].sort((left, right) => {
    const rightDegree = degreeMap.get(right.id) || 0;
    const leftDegree = degreeMap.get(left.id) || 0;
    if (rightDegree !== leftDegree) return rightDegree - leftDegree;

    const rightLabel = String(right.label || right.id || "");
    const leftLabel = String(left.label || left.id || "");
    return leftLabel.localeCompare(rightLabel);
  });
}

function cySizeFromDegree(deg) {
  const baseSize = 18;
  const maxBonus = Math.min(deg * 2, 22);
  return baseSize + maxBonus;
}

function fallbackRadiusFromDegree(deg) {
  const cySize = cySizeFromDegree(deg);
  const minRadius = RENDER_OPTIONS.nodeSizeMin;
  const maxRadius = RENDER_OPTIONS.nodeSizeMax;
  return Math.max(minRadius, Math.min(maxRadius, cySize / 3.1));
}

function matchesEdgeFilter(edgeType) {
  if (!RENDER_OPTIONS.edgeTypeFilter) return true;
  const types = RENDER_OPTIONS.edgeTypeFilter.split(',').map(t => t.trim()).filter(t => t);
  return types.length === 0 || types.includes(edgeType);
}

function drawGraph(filterText, filterKind) {
  if (!graphData || !Array.isArray(graphData.nodes) || graphData.nodes.length === 0) {
    showGraphMessage("Graph data unavailable");
    return;
  }

  const q = (filterText || "").toLowerCase().trim();
  const k = filterKind || "";

  const filteredNodes = graphData.nodes.filter((n) => {
    if (!matchesKind(n.type, k)) return false;
    const label = String(n.label || n.id || "").toLowerCase();
    if (q && !label.includes(q)) return false;
    return true;
  });

  const allLinks     = normalizeGraphLinks(graphData);
  const filteredIds  = new Set(filteredNodes.map((n) => n.id));
  const candidateLinks = allLinks.filter((l) => {
    const sid = l.source?.id ?? l.source;
    const tid = l.target?.id ?? l.target;
    if (!filteredIds.has(sid) || !filteredIds.has(tid)) return false;
    if (!matchesEdgeFilter(l.type)) return false;
    return true;
  });
  const candidateDegreeMap = buildDegreeMap(candidateLinks);
  const rankedNodes = sortNodesByDegreeDesc(filteredNodes, candidateDegreeMap);
  const visibleNodes = rankedNodes.slice(0, RENDER_OPTIONS.maxNodesCap);
  const visibleIds   = new Set(visibleNodes.map((n) => n.id));
  const visibleLinks = candidateLinks.filter((l) => {
    const sid = l.source?.id ?? l.source;
    const tid = l.target?.id ?? l.target;
    return visibleIds.has(sid) && visibleIds.has(tid);
  }).slice(0, RENDER_OPTIONS.maxLinksCap);

  if (typeof cytoscape !== "function") {
    drawFallbackSvg(visibleNodes, visibleLinks);
    return;
  }

  // Degree drives node sizing and label prioritization.
  const degMap = buildDegreeMap(visibleLinks);
  const topLabels = new Set(
    [...degMap.entries()].sort((a, b) => b[1] - a[1]).slice(0, 30).map(([id]) => id)
  );

  // Build Cytoscape elements.
  const cyElements = [];
  visibleNodes.forEach((n) => {
    const deg  = degMap.get(n.id) || 0;
    const size = cySizeFromDegree(deg);
    const fullLabel = String(n.label || n.id || "");
    cyElements.push({
      group: "nodes",
      data: {
        id:         n.id,
        label:      topLabels.has(n.id) ? (fullLabel.length > 22 ? fullLabel.slice(0, 20) + "..." : fullLabel) : "",
        fullLabel:  fullLabel,
        type:       n.type,
        category:   n.category || n.type,
        confidence: n.confidence,
        summary:    n.summary || "",
        path:       n.path || "",
        color:      kindColor(n.type, n),
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
    cyContainer.style.height = graphViewportHeightPx() + "px";
    cyContainer.style.background = "transparent";
    graphSvg.style.display = "none";
    viewport.appendChild(cyContainer);
  }
  cyContainer.style.height = graphViewportHeightPx() + "px";
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
    const mouseLeft = (canvasRect.left - r.left) + pos.x;
    const mouseTop  = (canvasRect.top  - r.top)  + pos.y;
    const offsetX = 8;
    const offsetY = 6;
    const tooltipWidth = graphTooltip.offsetWidth || 240;
    const tooltipHeight = graphTooltip.offsetHeight || 140;

    let left = mouseLeft + offsetX;
    let top = mouseTop + offsetY;

    if (left + tooltipWidth > r.width - 8) {
      left = mouseLeft - tooltipWidth - offsetX;
    }
    if (left < 8) {
      left = 8;
    }
    if (top + tooltipHeight > r.height - 8) {
      top = r.height - tooltipHeight - 8;
    }
    if (top < 4) {
      top = 4;
    }

    graphTooltip.style.left = left + "px";
    graphTooltip.style.top  = top + "px";
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

  // Ensure canvas dimensions match the visible container after panel toggle.
  requestAnimationFrame(() => {
    if (!cyInstance) return;
    cyInstance.resize();
    cyInstance.fit(undefined, 24);
  });
}


// --- Event listeners ---------------------------------------------------------
window.addEventListener("hashchange", () => {
  updateActiveNavLink();
  handleRoute();
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

if (mcpPageBtn) {
  mcpPageBtn.addEventListener("click", () => {
    window.location.hash = MCP_ROUTE;
  });
}

graphSearch.addEventListener("input",       () => { if (!graphPanel.classList.contains("hidden")) drawGraph(graphSearch.value, graphKindFilter.value); });
graphKindFilter.addEventListener("change",  () => { if (!graphPanel.classList.contains("hidden")) drawGraph(graphSearch.value, graphKindFilter.value); });

graphReset.addEventListener("click", () => {
  graphSearch.value     = "";
  graphKindFilter.value = "";
  if (cyInstance) cyInstance.fit();
  drawGraph("", "");
});

// --- Graph Rendering Controls -----------------------------------------------
function initGraphRenderingControls() {
  const edgeTypePanel = document.getElementById("graphEdgeTypePanel");
  if (!edgeTypePanel) {
    // Create controls if they don't exist yet
    const controlsHtml = `
      <div id="graphEdgeTypePanel" style="padding: 4px 6px; border-top: 1px solid #ddd; font-size: 0.68em; line-height: 1.1;">
        <div style="display: grid; grid-template-columns: minmax(0, 2.2fr) repeat(4, minmax(52px, 72px)); gap: 4px 6px; align-items: end;">
          <label for="graphEdgeTypeInput" style="display: block; min-width: 0;">
            <span style="display: block; margin-bottom: 1px; font-weight: 600;">Edges</span>
            <input type="text" id="graphEdgeTypeInput" placeholder="contains,calls" 
                   value="${RENDER_OPTIONS.edgeTypeFilter}" 
                   style="width: 100%; min-width: 0; padding: 1px 4px; font-size: 1em; line-height: 1.1;">
          </label>
          <label for="graphNodeCapInput" style="display: block;">
            <span style="display: block; margin-bottom: 1px;">Nodes</span>
            <input type="number" id="graphNodeCapInput" min="50" max="500" step="50" value="${RENDER_OPTIONS.maxNodesCap}" style="width: 100%; padding: 1px 3px; font-size: 1em; line-height: 1.1;">
          </label>
          <label for="graphLinkCapInput" style="display: block;">
            <span style="display: block; margin-bottom: 1px;">Links</span>
            <input type="number" id="graphLinkCapInput" min="100" max="2000" step="100" value="${RENDER_OPTIONS.maxLinksCap}" style="width: 100%; padding: 1px 3px; font-size: 1em; line-height: 1.1;">
          </label>
          <label for="graphNodeSizeMinInput" style="display: block;">
            <span style="display: block; margin-bottom: 1px;">Min</span>
            <input type="number" id="graphNodeSizeMinInput" min="1" max="15" step="1" value="${RENDER_OPTIONS.nodeSizeMin}" style="width: 100%; padding: 1px 3px; font-size: 1em; line-height: 1.1;">
          </label>
          <label for="graphNodeSizeMaxInput" style="display: block;">
            <span style="display: block; margin-bottom: 1px;">Max</span>
            <input type="number" id="graphNodeSizeMaxInput" min="16" max="50" step="2" value="${RENDER_OPTIONS.nodeSizeMax}" style="width: 100%; padding: 1px 3px; font-size: 1em; line-height: 1.1;">
          </label>
        </div>
      </div>`;
    
    const filterSection = document.querySelector("[id='graphKindFilter']")?.parentElement;
    if (filterSection) {
      filterSection.insertAdjacentHTML("afterend", controlsHtml);
    }
  }
  
  const edgeInput = document.getElementById("graphEdgeTypeInput");
  const nodeCapInput = document.getElementById("graphNodeCapInput");
  const linkCapInput = document.getElementById("graphLinkCapInput");
  const nodeSizeMinInput = document.getElementById("graphNodeSizeMinInput");
  const nodeSizeMaxInput = document.getElementById("graphNodeSizeMaxInput");
  
  if (edgeInput) {
    edgeInput.addEventListener("input", (e) => {
      RENDER_OPTIONS.edgeTypeFilter = e.target.value;
      localStorage.setItem("war-api-edge-types", e.target.value);
      if (!graphPanel.classList.contains("hidden")) drawGraph(graphSearch.value, graphKindFilter.value);
    });
  }
  
  if (nodeCapInput) {
    nodeCapInput.addEventListener("change", (e) => {
      const value = Math.max(50, Math.min(500, parseInt(e.target.value, 10) || RENDER_OPTIONS.maxNodesCap));
      e.target.value = String(value);
      RENDER_OPTIONS.maxNodesCap = value;
      localStorage.setItem("war-api-node-cap", String(value));
      if (!graphPanel.classList.contains("hidden")) drawGraph(graphSearch.value, graphKindFilter.value);
    });
  }
  
  if (linkCapInput) {
    linkCapInput.addEventListener("change", (e) => {
      const value = Math.max(100, Math.min(2000, parseInt(e.target.value, 10) || RENDER_OPTIONS.maxLinksCap));
      e.target.value = String(value);
      RENDER_OPTIONS.maxLinksCap = value;
      localStorage.setItem("war-api-link-cap", String(value));
      if (!graphPanel.classList.contains("hidden")) drawGraph(graphSearch.value, graphKindFilter.value);
    });
  }
  
  if (nodeSizeMinInput) {
    nodeSizeMinInput.addEventListener("change", (e) => {
      const value = Math.max(1, Math.min(15, parseInt(e.target.value, 10) || RENDER_OPTIONS.nodeSizeMin));
      e.target.value = String(value);
      RENDER_OPTIONS.nodeSizeMin = value;
      localStorage.setItem("war-api-node-size-min", String(value));
      if (!graphPanel.classList.contains("hidden")) drawGraph(graphSearch.value, graphKindFilter.value);
    });
  }
  
  if (nodeSizeMaxInput) {
    nodeSizeMaxInput.addEventListener("change", (e) => {
      const value = Math.max(16, Math.min(50, parseInt(e.target.value, 10) || RENDER_OPTIONS.nodeSizeMax));
      e.target.value = String(value);
      RENDER_OPTIONS.nodeSizeMax = value;
      localStorage.setItem("war-api-node-size-max", String(value));
      if (!graphPanel.classList.contains("hidden")) drawGraph(graphSearch.value, graphKindFilter.value);
    });
  }
}

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

if (analysisToggle) {
  analysisToggle.addEventListener("click", () => {
    applyAnalysisToggle(!document.body.classList.contains("show-advanced"));
  });
}

async function handleRoute() {
  const path = activePath();

  if (isMcpRoute(path)) {
    generateBreadcrumb(MCP_ROUTE);
    if (mcpPage) mcpPage.classList.remove("hidden");
    if (mcpPageBtn) mcpPageBtn.classList.add("secondary");
    docRoot.classList.add("hidden");
    graphPanel.classList.add("hidden");
    graphToggle.classList.add("hidden");
    graphToggle.textContent = "Graph";
    return;
  }

  if (mcpPage) mcpPage.classList.add("hidden");
  if (mcpPageBtn) mcpPageBtn.classList.remove("secondary");
  docRoot.classList.remove("hidden");
  graphToggle.classList.remove("hidden");
  await loadDoc(path);
}

// --- Init ---------------------------------------------------------------------
(async function init() {
  // Restore advanced-details toggle preference (default: hidden).
  let showAnalysis = false;
  try {
    const current = localStorage.getItem(ADVANCED_STORAGE_KEY);
    if (current === "1" || current === "0") {
      showAnalysis = current === "1";
    } else {
      // One-time migration from old key.
      showAnalysis = localStorage.getItem(LEGACY_ANALYSIS_STORAGE_KEY) === "1";
    }
  } catch {}
  applyAnalysisToggle(showAnalysis);

  await Promise.all([loadSearchIndex(), loadGraph(), loadNavigation()]);
  await handleRoute();
  initGraphRenderingControls();
})();


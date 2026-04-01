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
let d3sim          = null;
let d3zoom         = null;

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

// --- Graph --------------------------------------------------------------------
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

function capNodes(nodes, maxNodes) {
  if (nodes.length <= maxNodes) return nodes;

  // In all-types mode, preserve type diversity so one type does not dominate the first N entries.
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
    for (let i = 0; i < active.length && remaining > 0; i += 1) {
      const t = active[i];
      const arr = grouped.get(t);
      if (!arr || arr.length === 0) continue;
      result.push(arr.shift());
      remaining -= 1;
    }
    active = active.filter((t) => (grouped.get(t) || []).length > 0);
  }

  return result;
}

function normalizeGraphLinks(data) {
  if (!data) return [];

  if (Array.isArray(data.links) && data.links.length > 0) {
    return data.links;
  }

  if (Array.isArray(data.edges) && data.edges.length > 0) {
    return data.edges.map((e) => ({
      source: e.source?.id ?? e.source ?? e.from,
      target: e.target?.id ?? e.target ?? e.to,
      type: e.type || "related",
      weight: e.weight || 1,
    }));
  }

  return [];
}

function drawGraph(filterText, filterKind) {
  if (!graphData || !Array.isArray(graphData.nodes)) return;

  const q = (filterText || "").toLowerCase().trim();
  const k = filterKind || "";

  // -- filter & cap ----------------------------------------------------------
  const filteredNodes = graphData.nodes.filter((n) => {
    if (!matchesKind(n.type, k)) return false;
    if (q && !n.label.toLowerCase().includes(q)) return false;
    return true;
  });

  const visibleNodes = capNodes(filteredNodes, 300);
  const allLinks = normalizeGraphLinks(graphData);

  const visibleIds = new Set(visibleNodes.map((n) => n.id));
  const visibleLinks = allLinks.filter((l) => {
    const sid = l.source?.id ?? l.source;
    const tid = l.target?.id ?? l.target;
    return visibleIds.has(sid) && visibleIds.has(tid);
  }).slice(0, 800);

  // -- SVG setup -------------------------------------------------------------
  const svg = d3.select(graphSvg);
  svg.selectAll("*").remove();
  if (d3sim) { d3sim.stop(); d3sim = null; }

  const W = graphSvg.clientWidth  || 900;
  const H = graphSvg.clientHeight || 600;
  svg.attr("viewBox", `0 0 ${W} ${H}`);

  // Subtle dot-grid background
  const defs = svg.append("defs");
  const pat  = defs.append("pattern").attr("id", "dot").attr("width", 22).attr("height", 22).attr("patternUnits", "userSpaceOnUse");
  pat.append("circle").attr("cx", 11).attr("cy", 11).attr("r", 0.8).attr("fill", "rgba(160,145,110,.25)");
  svg.append("rect").attr("width", "100%").attr("height", "100%").attr("fill", "url(#dot)");

  // Zoom + pan
  const g = svg.append("g");
  d3zoom = d3.zoom().scaleExtent([0.05, 8]).on("zoom", (e) => g.attr("transform", e.transform));
  svg.call(d3zoom);

  // -- Data prep -------------------------------------------------------------
  const nodes    = visibleNodes.map((n) => ({ ...n }));
  const nodeById = new Map(nodes.map((n) => [n.id, n]));
  const links    = visibleLinks.map((l) => ({
    source: nodeById.get(l.source?.id ?? l.source) ?? (l.source?.id ?? l.source),
    target: nodeById.get(l.target?.id ?? l.target) ?? (l.target?.id ?? l.target),
    type:   l.type || "related",
  }));

  // Degree-based sizing
  const degMap = new Map();
  links.forEach((l) => {
    const s = typeof l.source === "object" ? l.source.id : l.source;
    const t = typeof l.target === "object" ? l.target.id : l.target;
    degMap.set(s, (degMap.get(s) || 0) + 1);
    degMap.set(t, (degMap.get(t) || 0) + 1);
  });
  const nodeR = (d) => 4 + Math.min((degMap.get(d.id) || 0) * 0.45, 8);

  // Top-N by degree get permanent labels
  const topLabels = new Set(
    [...degMap.entries()].sort((a, b) => b[1] - a[1]).slice(0, 28).map(([id]) => id)
  );

  // -- Render links ----------------------------------------------------------
  const linkEl = g.append("g").selectAll("line").data(links).join("line")
    .attr("stroke", "#4a4030")
    .attr("stroke-opacity", 0.38)
    .attr("stroke-width", 1);

  // -- Drag ------------------------------------------------------------------
  const drag = d3.drag()
    .on("start", (e, d) => { if (!e.active) d3sim.alphaTarget(0.3).restart(); d.fx = d.x; d.fy = d.y; })
    .on("drag",  (e, d) => { d.fx = e.x; d.fy = e.y; })
    .on("end",   (e, d) => { if (!e.active) d3sim.alphaTarget(0); d.fx = null; d.fy = null; });

  // -- Render nodes ----------------------------------------------------------
  const nodeEl = g.append("g").selectAll("circle").data(nodes).join("circle")
    .attr("r",            nodeR)
    .attr("fill",         (d) => kindColor(d.type))
    .attr("stroke",       "rgba(240,235,220,.55)")
    .attr("stroke-width", 1.2)
    .attr("cursor",       "pointer")
    .call(drag);

  // -- Permanent labels for high-degree nodes -----------------------------
  const labelEl = g.append("g").selectAll("text")
    .data(nodes.filter((n) => topLabels.has(n.id))).join("text")
    .attr("fill",             "#cfc8b8")
    .attr("font-size",        9.5)
    .attr("font-family",      "Inter, system-ui, sans-serif")
    .attr("text-anchor",      "middle")
    .attr("pointer-events",   "none")
    .text((d) => d.label.length > 22 ? d.label.slice(0, 20) + "..." : d.label);

  // -- Neighbour highlight on click ------------------------------------------
  function neighbours(id) {
    const s = new Set([id]);
    links.forEach((l) => {
      const sid = l.source?.id ?? l.source;
      const tid = l.target?.id ?? l.target;
      if (sid === id) s.add(tid);
      if (tid === id) s.add(sid);
    });
    return s;
  }

  let selectedId = null;

  function resetHighlight() {
    selectedId = null;
    nodeEl.attr("opacity", 1).attr("stroke-width", 1.2).attr("stroke", "rgba(240,235,220,.55)");
    linkEl.attr("stroke-opacity", 0.38).attr("stroke", "#4a4030");
  }

  nodeEl
    .on("mouseover", (e, d) => {
      graphTooltip.classList.remove("hidden");
      graphTooltip.innerHTML =
        `<strong>${d.label}</strong>` +
        `<div class="tt-type">${d.category || d.type}${d.confidence != null ? ` - ${d.confidence}% confidence` : ""}</div>` +
        (d.summary ? `<div class="tt-summary">${d.summary}</div>` : "") +
        (d.path    ? `<div class="tt-type" style="margin-top:6px;color:#a08060">Click to open docs -></div>` : "");
    })
    .on("mousemove", (e) => {
      const r = graphPanel.getBoundingClientRect();
      let left = e.clientX - r.left + 14;
      const top  = e.clientY - r.top  - 10;
      if (left + 284 > r.width - 10) left = e.clientX - r.left - 284 - 14;
      graphTooltip.style.left = left + "px";
      graphTooltip.style.top  = Math.max(4, top) + "px";
    })
    .on("mouseout",  () => graphTooltip.classList.add("hidden"))
    .on("click", (e, d) => {
      e.stopPropagation();
      if (selectedId === d.id) {
        resetHighlight();
        if (d.path) window.location.hash = d.path;
        return;
      }
      selectedId = d.id;
      const nb = neighbours(d.id);
      nodeEl
        .attr("opacity",      (n) => nb.has(n.id) ? 1 : 0.12)
        .attr("stroke",       (n) => n.id === d.id ? "#fff" : "rgba(240,235,220,.55)")
        .attr("stroke-width", (n) => n.id === d.id ? 3 : 1.2);
      linkEl
        .attr("stroke-opacity", (l) => {
          const s = l.source?.id ?? l.source;
          const t = l.target?.id ?? l.target;
          return (s === d.id || t === d.id) ? 0.85 : 0.04;
        })
        .attr("stroke", (l) => {
          const s = l.source?.id ?? l.source;
          const t = l.target?.id ?? l.target;
          return (s === d.id || t === d.id) ? "#c8a070" : "#4a4030";
        });
    });

  svg.on("click", resetHighlight);

  // -- Force simulation ------------------------------------------------------
  d3sim = d3.forceSimulation(nodes)
    .force("link",      d3.forceLink(links).id((d) => d.id).distance(75).strength(0.38))
    .force("charge",    d3.forceManyBody().strength(-130).distanceMax(380))
    .force("center",    d3.forceCenter(W / 2, H / 2))
    .force("collision", d3.forceCollide((d) => nodeR(d) + 3));

  d3sim.on("tick", () => {
    linkEl
      .attr("x1", (d) => d.source.x).attr("y1", (d) => d.source.y)
      .attr("x2", (d) => d.target.x).attr("y2", (d) => d.target.y);
    nodeEl
      .attr("cx", (d) => d.x).attr("cy", (d) => d.y);
    labelEl
      .attr("x", (d) => d.x).attr("y", (d) => d.y - nodeR(d) - 3);
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
  if (d3zoom) d3.select(graphSvg).transition().duration(300).call(d3zoom.transform, d3.zoomIdentity);
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


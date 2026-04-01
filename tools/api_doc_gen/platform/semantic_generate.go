package platform

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	md "roraddons/tools/api_doc_gen/templates"
)

func writeBrowserIndex(outputRoot string) error {
	content := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>WAR Addon API Reference</title>
	<link rel="preconnect" href="https://fonts.googleapis.com">
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
	<link href="https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:wght@400;500;600&family=IBM+Plex+Mono:wght@400;500&display=swap" rel="stylesheet">
	<style>
		:root {
			--bg: #f5f1e8;
			--panel: #fffdf7;
			--ink: #1d2430;
			--muted: #5f6875;
			--line: #d8d0be;
			--accent: #235a52;
			--accent-soft: #d7e7de;
			--code: #1f2937;
			--code-bg: #eef3f1;
			--shadow: 0 18px 40px rgba(29, 36, 48, 0.08);
		}
		* { box-sizing: border-box; }
		body {
			margin: 0;
			font-family: "IBM Plex Sans", sans-serif;
			color: var(--ink);
			background:
				radial-gradient(circle at top left, rgba(35, 90, 82, 0.08), transparent 28%),
				linear-gradient(180deg, #fbf8f1 0%, var(--bg) 100%);
			min-height: 100vh;
		}
		.shell {
			display: grid;
			grid-template-columns: 320px minmax(0, 1fr);
			min-height: 100vh;
		}
		.sidebar {
			border-right: 1px solid var(--line);
			background: rgba(255, 253, 247, 0.92);
			backdrop-filter: blur(12px);
			padding: 20px 18px 28px;
			position: sticky;
			top: 0;
			height: 100vh;
			overflow: auto;
		}
		.brand {
			margin-bottom: 16px;
			padding-bottom: 16px;
			border-bottom: 1px solid var(--line);
		}
		.brand h1 {
			font-size: 1.1rem;
			margin: 0 0 6px;
			letter-spacing: 0.02em;
		}
		.brand p {
			margin: 0;
			color: var(--muted);
			font-size: 0.92rem;
			line-height: 1.45;
		}
		.shortcut-row {
			display: flex;
			flex-wrap: wrap;
			gap: 8px;
			margin: 16px 0 18px;
		}
		.shortcut-row a {
			border: 1px solid var(--line);
			border-radius: 999px;
			padding: 7px 12px;
			color: var(--ink);
			text-decoration: none;
			font-size: 0.84rem;
			background: var(--panel);
		}
		.shortcut-row a:hover {
			border-color: var(--accent);
			color: var(--accent);
		}
		.tree,
		.tree ul {
			list-style: none;
			margin: 0;
			padding-left: 0;
		}
		.tree ul {
			margin-left: 14px;
			padding-left: 12px;
			border-left: 1px solid rgba(95, 104, 117, 0.18);
		}
		.tree li {
			margin: 4px 0;
		}
		.tree button,
		.tree a,
		.tree span {
			font: inherit;
			color: inherit;
		}
		.tree button {
			background: transparent;
			border: 0;
			padding: 6px 8px;
			width: 100%;
			text-align: left;
			border-radius: 10px;
			cursor: pointer;
			color: var(--muted);
		}
		.tree a {
			display: block;
			text-decoration: none;
			padding: 6px 8px;
			border-radius: 10px;
			color: var(--muted);
		}
		.tree a:hover,
		.tree button:hover,
		.tree .active {
			background: var(--accent-soft);
			color: var(--accent);
		}
		.content {
			padding: 26px;
		}
		.content-shell {
			max-width: 1040px;
			margin: 0 auto;
			background: var(--panel);
			border: 1px solid rgba(216, 208, 190, 0.9);
			border-radius: 26px;
			box-shadow: var(--shadow);
			overflow: hidden;
		}
		.content-header {
			padding: 18px 24px;
			border-bottom: 1px solid var(--line);
			background: linear-gradient(90deg, rgba(35, 90, 82, 0.08), rgba(35, 90, 82, 0));
		}
		.content-header strong {
			display: block;
			font-size: 0.82rem;
			letter-spacing: 0.08em;
			text-transform: uppercase;
			color: var(--muted);
			margin-bottom: 6px;
		}
		.content-header code {
			font-family: "IBM Plex Mono", monospace;
			font-size: 0.88rem;
			color: var(--accent);
			background: transparent;
			padding: 0;
		}
		article {
			padding: 28px 24px 40px;
			line-height: 1.7;
		}
		article h1,
		article h2,
		article h3 {
			line-height: 1.2;
			color: #13202a;
		}
		article a {
			color: var(--accent);
		}
		article code,
		article pre {
			font-family: "IBM Plex Mono", monospace;
		}
		article code {
			background: var(--code-bg);
			padding: 0.16rem 0.35rem;
			border-radius: 6px;
			color: var(--code);
		}
		article pre {
			background: #13202a;
			color: #f2f4f7;
			padding: 16px;
			border-radius: 16px;
			overflow: auto;
		}
		article table {
			width: 100%;
			border-collapse: collapse;
			margin: 18px 0;
			font-size: 0.94rem;
		}
		article th,
		article td {
			border: 1px solid var(--line);
			padding: 10px 12px;
			vertical-align: top;
		}
		article th {
			background: rgba(35, 90, 82, 0.08);
			text-align: left;
		}
		.status {
			color: var(--muted);
		}
		@media (max-width: 960px) {
			.shell { grid-template-columns: 1fr; }
			.sidebar {
				position: static;
				height: auto;
				border-right: 0;
				border-bottom: 1px solid var(--line);
			}
			.content { padding: 18px; }
		}
	</style>
</head>
<body>
	<div class="shell">
		<aside class="sidebar">
			<div class="brand">
				<h1>WAR Addon API Reference</h1>
				<p>Semantic platform reference generated from docs/addon-api and enriched with navigation and relationship data.</p>
			</div>
			<div class="shortcut-row">
				<a href="#index.md">Overview</a>
				<a href="graph/api_graph.json">Graph JSON</a>
				<a href="graph/relations.json">Relations</a>
				<a href="navigation/tree.json">Tree</a>
			</div>
			<nav id="nav" class="tree"><div class="status">Loading navigation...</div></nav>
		</aside>
		<main class="content">
			<div class="content-shell">
				<div class="content-header">
					<strong>Current Document</strong>
					<code id="current-path">index.md</code>
				</div>
				<article id="doc"><p class="status">Loading documentation...</p></article>
			</div>
		</main>
	</div>

	<script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js"></script>
	<script>
		const navRoot = document.getElementById('nav');
		const docRoot = document.getElementById('doc');
		const currentPath = document.getElementById('current-path');
		const defaultPath = 'index.md';

		function getActivePath() {
			const hash = window.location.hash.replace(/^#/, '').trim();
			return hash || defaultPath;
		}

		function resolvePath(basePath, targetPath) {
			return new URL(targetPath, 'https://local/' + basePath).pathname.replace(/^\//, '');
		}

		function navLink(path, label) {
			const anchor = document.createElement('a');
			anchor.href = '#' + path;
			anchor.textContent = label;
			if (getActivePath() === path) {
				anchor.classList.add('active');
			}
			return anchor;
		}

		function renderNavItems(items) {
			const list = document.createElement('ul');
			for (const item of items) {
				const row = document.createElement('li');
				const hasChildren = Array.isArray(item.items) && item.items.length > 0;
				if (item.path) {
					row.appendChild(navLink(item.path, item.label));
				} else if (hasChildren) {
					const button = document.createElement('button');
					button.textContent = item.label;
					row.appendChild(button);
				} else {
					const label = document.createElement('span');
					label.textContent = item.label;
					row.appendChild(label);
				}
				if (hasChildren) {
					row.appendChild(renderNavItems(item.items));
				}
				list.appendChild(row);
			}
			return list;
		}

		async function loadNavigation() {
			const response = await fetch('navigation/sidebar.json');
			const items = await response.json();
			navRoot.innerHTML = '';
			navRoot.appendChild(renderNavItems(items));
		}

		function rewriteLinks(basePath) {
			for (const anchor of docRoot.querySelectorAll('a[href]')) {
				const href = anchor.getAttribute('href');
				if (!href || href.startsWith('http://') || href.startsWith('https://') || href.startsWith('#')) {
					continue;
				}
				const resolved = resolvePath(basePath, href);
				if (resolved.endsWith('.md')) {
					anchor.setAttribute('href', '#' + resolved);
				} else {
					anchor.setAttribute('href', resolved);
				}
			}
		}

		async function loadDocument(path) {
			const response = await fetch(path);
			if (!response.ok) {
				throw new Error('Unable to load ' + path + ' (' + response.status + ')');
			}
			const markdown = await response.text();
			currentPath.textContent = path;
			docRoot.innerHTML = marked.parse(markdown);
			rewriteLinks(path);
			for (const link of navRoot.querySelectorAll('a')) {
				link.classList.toggle('active', link.getAttribute('href') === '#' + path);
			}
			window.scrollTo({ top: 0, behavior: 'smooth' });
		}

		async function boot() {
			try {
				await loadNavigation();
				await loadDocument(getActivePath());
			} catch (error) {
				docRoot.innerHTML = '<p class="status">' + error.message + '</p>';
			}
		}

		window.addEventListener('hashchange', () => {
			loadDocument(getActivePath()).catch((error) => {
				docRoot.innerHTML = '<p class="status">' + error.message + '</p>';
			});
		});

		boot();
	</script>
</body>
</html>`
	return writeFile(filepath.Join(outputRoot, "index.html"), content)
}

func writeNavigationArtifacts(outputRoot string, corpus Corpus) error {
	if err := writeJSONFile(filepath.Join(outputRoot, "navigation", "tree.json"), corpus.NavigationTree); err != nil {
		return err
	}
	return writeJSONFile(filepath.Join(outputRoot, "navigation", "sidebar.json"), corpus.Sidebar)
}

func writeGraphArtifacts(outputRoot string, corpus Corpus) error {
	if err := writeJSONFile(filepath.Join(outputRoot, "graph", "api_graph.json"), corpus.APIGraph); err != nil {
		return err
	}
	return writeJSONFile(filepath.Join(outputRoot, "graph", "relations.json"), corpus.Relations)
}

func writePatternDocs(outputRoot string, corpus Corpus) error {
	rows := [][]string{}
	for _, page := range corpus.PatternPages {
		rows = append(rows, []string{page.Category, markdownLink(page.Name, filepath.Base(page.Path)), string(page.Confidence), fmt.Sprintf("%d", len(page.Involved))})
		content := "# " + page.Name + "\n\n"
		content += "- Category: " + page.Category + "\n"
		content += "- Confidence: " + string(page.Confidence) + "\n\n"
		content += md.Section("Description", page.Description)
		content += md.Section("Involved APIs", md.BulletList(docLinkLines(page.Path, page.Involved)))
		content += md.Section("Flow Diagram", renderFlowDiagram(page.Flow))
		if strings.TrimSpace(page.ExampleCode) != "" {
			content += md.Section("Example Code", "```lua\n"+page.ExampleCode+"\n```")
		}
		content += md.Section("Evidence", md.BulletList(page.Evidence))
		if err := writeFile(filepath.Join(outputRoot, filepath.FromSlash(page.Path)), content); err != nil {
			return err
		}
	}
	index := "# Semantic Patterns\n\n"
	index += "These pages collect reusable WAR addon usage patterns inferred from the canonical platform surface and its supporting evidence.\n\n"
	index += md.Section("Patterns", md.Table([]string{"Category", "Pattern", "Confidence", "APIs"}, rows))
	return writeFile(filepath.Join(outputRoot, "patterns", "index.md"), index)
}

func writeFieldDocs(outputRoot string, title string, directory string, namespace string, values []FieldSymbol, corpus Corpus) error {
	rows := [][]string{}
	for _, symbol := range values {
		currentPath := slashPath(directory, "fields", docName(strings.ToLower(namespace), symbol.Name))
		content := renderFieldSymbol(symbol, namespace, currentPath, corpus.SymbolLinks[fieldSymbolID(namespace, symbol.Name)])
		if err := writeFile(filepath.Join(outputRoot, filepath.FromSlash(currentPath)), content); err != nil {
			return err
		}
		rows = append(rows, []string{markdownLink(symbol.Name, filepath.ToSlash(filepath.Join("fields", docName(strings.ToLower(namespace), symbol.Name)))), fmt.Sprintf("%d", symbol.Score), string(symbol.Confidence), fmt.Sprintf("%d", len(symbol.SeenIn)), summarizeEvidence(symbol.Evidence), symbol.Rationale})
	}
	content := "# " + title + "\n\n"
	content += md.Section("Observed Fields", md.Table([]string{"Field", "Score", "Level", "Addons", "Evidence", "Rationale"}, rows))
	return writeFile(filepath.Join(outputRoot, directory, "index.md"), content)
}

func renderFieldSymbol(symbol FieldSymbol, namespace string, currentPath string, links SemanticLinks) string {
	content := "# " + symbol.Name + "\n\n"
	content += "- Category: " + namespace + " Field\n"
	content += "- Confidence level: " + string(symbol.Confidence) + "\n"
	content += fmt.Sprintf("- Confidence score: %d/100\n", symbol.Score)
	content += fmt.Sprintf("- Seen in: %d addons\n\n", len(symbol.SeenIn))
	content += renderConfidenceSections(symbol.Confidence, symbol.Score, symbol.RawScore, symbol.Rationale, symbol.Signals, symbol.Evidence)
	content += md.Section("Description", symbol.Description)
	content += md.Section("Seen In", md.BulletList(symbol.SeenIn))
	content += renderSemanticSections(currentPath, links)
	content += md.Section("Notes", md.BulletList(symbol.Notes))
	return content
}

func renderSemanticSections(currentPath string, links SemanticLinks) string {
	content := ""
	content += md.Section("Related APIs", md.BulletList(docLinkLines(currentPath, links.RelatedAPIs)))
	content += md.Section("Used With", md.BulletList(docLinkLines(currentPath, links.UsedWith)))
	content += md.Section("Triggered By", md.BulletList(docLinkLines(currentPath, links.TriggeredBy)))
	content += md.Section("Affects", md.BulletList(docLinkLines(currentPath, links.Affects)))
	return content
}

func docLinkLines(currentPath string, links []DocLink) []string {
	if len(links) == 0 {
		return []string{"none"}
	}
	lines := []string{}
	for _, link := range links {
		label := markdownLink(link.Label, relativeDocPath(currentPath, link.Path))
		if link.Score > 0 {
			label += " (" + formatConfidenceShort(link.Confidence, link.Score) + ")"
		}
		if strings.TrimSpace(link.Category) != "" {
			label += " - " + link.Category
		}
		lines = append(lines, label)
	}
	return lines
}

func relativeDocPath(currentPath string, targetPath string) string {
	if strings.TrimSpace(targetPath) == "" {
		return "#"
	}
	fromDirectory := filepath.Dir(filepath.FromSlash(currentPath))
	relativePath, err := filepath.Rel(fromDirectory, filepath.FromSlash(targetPath))
	if err != nil {
		return filepath.ToSlash(targetPath)
	}
	return filepath.ToSlash(relativePath)
}

func renderFlowDiagram(lines []string) string {
	if len(lines) == 0 {
		return "```text\nNo stable semantic flow was inferable from the current corpus.\n```"
	}
	return "```text\n" + strings.Join(lines, "\n") + "\n```"
}

func writeJSONFile(path string, value any) error {
	encoded, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, string(encoded))
}

func fieldSymbolID(namespace string, name string) string {
	if strings.EqualFold(namespace, "SystemData") {
		return systemFieldID(name)
	}
	return gameFieldID(name)
}

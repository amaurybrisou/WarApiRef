package server

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
)

type graph struct {
	Nodes []struct {
		ID         string `json:"id"`
		Label      string `json:"label"`
		Type       string `json:"type"`
		Category   string `json:"category"`
		Path       string `json:"path"`
		Confidence string `json:"confidence"`
		Score      int    `json:"score"`
		Summary    string `json:"summary"`
	} `json:"nodes"`
	Edges []struct {
		From     string   `json:"from"`
		To       string   `json:"to"`
		Type     string   `json:"type"`
		Weight   int      `json:"weight"`
		Evidence []string `json:"evidence"`
	} `json:"edges"`
}

type relations struct {
	FrequentCallChains []struct {
		Title        string   `json:"title"`
		Description  string   `json:"description"`
		Flow         []string `json:"flow"`
		Evidence     []string `json:"evidence"`
		Participants []struct {
			ID string `json:"id"`
		} `json:"participants"`
	} `json:"frequent_call_chains"`
	CommonEventFlows []struct {
		Title        string   `json:"title"`
		Description  string   `json:"description"`
		Flow         []string `json:"flow"`
		Evidence     []string `json:"evidence"`
		Participants []struct {
			ID string `json:"id"`
		} `json:"participants"`
	} `json:"common_event_flows"`
}

type symbolDoc struct {
	Summary   model.SymbolSummary
	Sections  map[string][]string
	Related   []string
	UsedWith  []string
	Triggered []string
	Affects   []string
	RawPath   string
}

type Store struct {
	root       string
	symbols    map[string]*symbolDoc
	byName     map[string][]*symbolDoc
	byNormName map[string][]*symbolDoc
	byPath     map[string]*symbolDoc
	graph      graph
	relations  relations
	meta       map[string]string
}

func NewStore(docsRoot string) (*Store, error) {
	root, err := filepath.Abs(docsRoot)
	if err != nil {
		return nil, fmt.Errorf("resolve docs root: %w", err)
	}
	if info, err := os.Stat(root); err != nil || !info.IsDir() {
		if err != nil {
			return nil, fmt.Errorf("stat docs root: %w", err)
		}
		return nil, fmt.Errorf("docs root is not a directory: %s", root)
	}
	s := &Store{
		root:       root,
		symbols:    map[string]*symbolDoc{},
		byName:     map[string][]*symbolDoc{},
		byNormName: map[string][]*symbolDoc{},
		byPath:     map[string]*symbolDoc{},
		meta:       map[string]string{},
	}
	if err := s.loadGraph(); err != nil {
		return nil, err
	}
	if err := s.loadRelations(); err != nil {
		return nil, err
	}
	if err := s.loadMeta(); err != nil {
		return nil, err
	}
	if err := s.loadSymbols(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) loadGraph() error {
	b, err := os.ReadFile(filepath.Join(s.root, "graph", "api_graph.json"))
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &s.graph)
}

func (s *Store) loadRelations() error {
	b, err := os.ReadFile(filepath.Join(s.root, "graph", "relations.json"))
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &s.relations)
}

func (s *Store) loadMeta() error {
	pairs := map[string]string{
		"war-api://meta/conventions":     "meta/conventions.md",
		"war-api://meta/inference-rules": "meta/inference_rules.md",
		"war-api://meta/coverage-report": "meta/coverage_report.md",
		"war-api://navigation/tree":      "navigation/tree.json",
		"war-api://navigation/sidebar":   "navigation/sidebar.json",
	}
	for k, rel := range pairs {
		b, err := os.ReadFile(filepath.Join(s.root, rel))
		if err != nil {
			continue
		}
		s.meta[k] = string(b)
	}
	return nil
}

func (s *Store) loadSymbols() error {
	byDocPath := map[string]struct {
		ID, Type, Category, Label, Confidence, Summary string
		Score                                          int
	}{}
	for _, n := range s.graph.Nodes {
		byDocPath[normalizePath(n.Path)] = struct {
			ID, Type, Category, Label, Confidence, Summary string
			Score                                          int
		}{ID: n.ID, Type: n.Type, Category: n.Category, Label: n.Label, Confidence: n.Confidence, Summary: n.Summary, Score: n.Score}
	}
	return filepath.WalkDir(s.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			n := strings.ToLower(d.Name())
			if n == "graph" || n == "navigation" || n == "search" || n == "scripts" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}
		rel, _ := filepath.Rel(s.root, path)
		rel = normalizePath(rel)
		if strings.HasSuffix(rel, "/index.md") || rel == "index.md" || strings.EqualFold(rel, "README.md") {
			return nil
		}
		b, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		doc := parseMarkdownSymbol(string(b), rel)
		if meta, ok := byDocPath[rel]; ok {
			doc.Summary.ID = buildStableID(doc.Summary.Kind, meta.Label)
			doc.Summary.Type = meta.Type
			doc.Summary.CanonicalName = meta.Label
			doc.Summary.Title = meta.Label
			doc.Summary.ConfidenceScore = meta.Score
			doc.Summary.ConfidenceLevel = meta.Confidence
			doc.Summary.ConfidenceRationale = "Derived from canonical generated WAR API docs"
			doc.Summary.SemanticConfidenceScore = semanticScore(meta.Score)
			doc.Summary.SemanticConfidenceLevel = model.LevelFromScore(doc.Summary.SemanticConfidenceScore)
			doc.Summary.SemanticConfidenceRationale = "Derived from semantic graph and relation evidence"
			doc.Summary.Summary = firstNonEmpty(doc.Sections["Description"])
			if doc.Summary.Summary == "" {
				doc.Summary.Summary = meta.Summary
			}
		}
		if doc.Summary.CanonicalName == "" {
			return nil
		}
		s.symbols[doc.Summary.ID] = doc
		s.byName[doc.Summary.CanonicalName] = append(s.byName[doc.Summary.CanonicalName], doc)
		norm := schema.NormalizeName(doc.Summary.CanonicalName)
		s.byNormName[norm] = append(s.byNormName[norm], doc)
		s.byPath[doc.Summary.DocPath] = doc
		return nil
	})
}

func (s *Store) LookupSymbol(name, symbolType string) (bool, bool, model.SymbolSummary, []model.CandidateMatch) {
	name = strings.TrimSpace(name)
	if name == "" {
		return false, false, model.SymbolSummary{}, nil
	}
	if candidates := s.byName[name]; len(candidates) > 0 {
		if hit := pickByType(candidates, symbolType); hit != nil {
			return true, true, hit.Summary, s.suggestions(name, symbolType, 6)
		}
	}
	norm := schema.NormalizeName(name)
	if candidates := s.byNormName[norm]; len(candidates) > 0 {
		if hit := pickByType(candidates, symbolType); hit != nil {
			return false, true, hit.Summary, s.suggestions(name, symbolType, 6)
		}
	}
	fuzzy := s.Search(name, symbolType, "", 0, 6)
	alts := []model.CandidateMatch{}
	for _, m := range fuzzy {
		alts = append(alts, model.CandidateMatch{ID: m.ID, CanonicalName: m.CanonicalName, Type: m.Type, DocPath: m.DocPath})
	}
	return false, false, model.SymbolSummary{}, alts
}

func (s *Store) Search(query, symbolType, namespace string, confidenceMin, limit int) []model.SymbolSummary {
	if limit <= 0 {
		limit = 10
	}
	terms := strings.Fields(strings.ToLower(strings.TrimSpace(query)))
	type scored struct {
		score int
		doc   *symbolDoc
	}
	buf := []scored{}
	for _, doc := range s.symbols {
		sym := doc.Summary
		if symbolType != "" && !strings.Contains(strings.ToLower(sym.Type), strings.ToLower(symbolType)) {
			continue
		}
		if namespace != "" && !strings.EqualFold(sym.Namespace, namespace) {
			continue
		}
		if confidenceMin > 0 && sym.ConfidenceScore < confidenceMin {
			continue
		}
		hay := strings.ToLower(sym.CanonicalName + " " + sym.Summary + " " + sym.Type + " " + sym.Namespace)
		score := 0
		for _, t := range terms {
			if strings.Contains(hay, t) {
				score++
			}
			if strings.Contains(strings.ToLower(sym.CanonicalName), t) {
				score += 2
			}
		}
		if score > 0 || len(terms) == 0 {
			score += sym.ConfidenceScore / 25
			buf = append(buf, scored{score: score, doc: doc})
		}
	}
	sort.Slice(buf, func(i, j int) bool {
		if buf[i].score == buf[j].score {
			return buf[i].doc.Summary.CanonicalName < buf[j].doc.Summary.CanonicalName
		}
		return buf[i].score > buf[j].score
	})
	if len(buf) > limit {
		buf = buf[:limit]
	}
	out := make([]model.SymbolSummary, 0, len(buf))
	for _, b := range buf {
		out = append(out, b.doc.Summary)
	}
	return out
}

func (s *Store) Related(source model.SymbolSummary, relationTypes []string, depth, limit int) map[string][]model.SymbolSummary {
	if limit <= 0 {
		limit = 12
	}
	allow := map[string]bool{}
	for _, r := range relationTypes {
		allow[r] = true
	}
	include := func(rt string) bool { return len(allow) == 0 || allow[rt] }
	out := map[string][]model.SymbolSummary{}
	src := s.symbols[source.ID]
	if src == nil {
		return out
	}
	addLinks := func(rt string, ids []string) {
		if !include(rt) {
			return
		}
		for _, id := range ids {
			if d, ok := s.symbols[id]; ok {
				out[rt] = append(out[rt], d.Summary)
				if len(out[rt]) >= limit {
					break
				}
			}
		}
	}
	addLinks("related", src.Related)
	addLinks("used_with", src.UsedWith)
	addLinks("triggered_by", src.Triggered)
	addLinks("affects", src.Affects)
	if depth > 0 {
		for _, e := range s.graph.Edges {
			if e.From != source.ID && e.To != source.ID {
				continue
			}
			other := e.To
			if e.To == source.ID {
				other = e.From
			}
			if doc, ok := s.symbols[other]; ok && include(e.Type) {
				out[e.Type] = append(out[e.Type], doc.Summary)
			}
		}
	}
	return out
}

func (s *Store) EventFlow(eventName string, includeExamples, includeDownstream bool) model.EventFlow {
	_, _, eventSym, _ := s.LookupSymbol(eventName, "event")
	flow := model.EventFlow{Event: eventSym, Warnings: []model.Warning{}}
	for _, c := range s.relations.CommonEventFlows {
		if strings.Contains(strings.ToLower(c.Title), strings.ToLower(eventName)) {
			flow.CommonFlows = append(flow.CommonFlows, c.Flow...)
			if includeExamples {
				flow.Examples = append(flow.Examples, c.Evidence...)
			}
			for _, p := range c.Participants {
				if d, ok := s.symbols[p.ID]; ok {
					if strings.Contains(strings.ToLower(d.Summary.Type), "handler") {
						flow.ObservedHandlers = append(flow.ObservedHandlers, d.Summary)
					} else if includeDownstream {
						flow.DownstreamSymbols = append(flow.DownstreamSymbols, d.Summary)
					}
				}
			}
		}
	}
	if len(flow.CommonFlows) == 0 {
		flow.Warnings = append(flow.Warnings, model.Warning{Code: "no_observed_flow", Message: "No explicit event flow was found for this event in relations.json."})
	}
	return flow
}

func (s *Store) XMLHandlerInfo(handlerName string, includeExamples, includeRelated bool) model.XMLHandlerInfo {
	_, _, sym, _ := s.LookupSymbol(handlerName, "xml_handler")
	info := model.XMLHandlerInfo{Handler: sym}
	doc := s.symbols[sym.ID]
	if doc == nil {
		return info
	}
	info.ExpectedCallbackShape = codeBlock(doc.Sections["Expected Lua Binding"])
	if info.ExpectedCallbackShape == "" {
		info.ExpectedCallbackShape = "function(...)"
	}
	info.CommonElementTypes = parseBullets(doc.Sections["Element Types"])
	if includeExamples {
		info.Examples = firstN(parseBullets(doc.Sections["Examples"]), 10)
	}
	if includeRelated {
		related := s.Related(sym, []string{"used_with", "related"}, 1, 10)
		for _, group := range related {
			info.PairedSymbols = append(info.PairedSymbols, group...)
		}
	}
	if strings.HasPrefix(sym.CanonicalName, "On") {
		info.LifecycleRole = "UI event callback hook"
	}
	return info
}

func (s *Store) UsageExamples(symbolName, patternName, addonFilter string, limit int) []model.UsageExample {
	if limit <= 0 {
		limit = 10
	}
	out := []model.UsageExample{}
	if symbolName != "" {
		_, _, sym, _ := s.LookupSymbol(symbolName, "")
		doc := s.symbols[sym.ID]
		if doc != nil {
			for _, ex := range parseBullets(doc.Sections["Examples"]) {
				item := model.UsageExample{AddonName: addonFromExample(ex), Title: sym.CanonicalName, Snippet: ex, Explanation: "Observed generated usage example", DocPath: sym.DocPath, RelatedSymbolIDs: sym.RelatedIDs}
				if addonFilter != "" && !strings.EqualFold(item.AddonName, addonFilter) {
					continue
				}
				out = append(out, item)
				if len(out) >= limit {
					break
				}
			}
		}
	}
	if patternName != "" && len(out) < limit {
		if b, err := os.ReadFile(filepath.Join(s.root, "patterns", patternName+".md")); err == nil {
			lines := strings.Split(string(b), "\n")
			for _, l := range lines {
				if strings.HasPrefix(strings.TrimSpace(l), "- ") {
					out = append(out, model.UsageExample{AddonName: "multiple", Title: patternName, Snippet: strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(l), "- ")), Explanation: "Pattern evidence", DocPath: "patterns/" + patternName + ".md"})
					if len(out) >= limit {
						break
					}
				}
			}
		}
	}
	return out
}

func (s *Store) ExplainConfidence(symbolName string) (model.SymbolSummary, map[string]any, map[string]any, []string, []string, string) {
	_, inferred, sym, _ := s.LookupSymbol(symbolName, "")
	doc := s.symbols[sym.ID]
	if doc == nil {
		return model.SymbolSummary{}, nil, nil, nil, nil, ""
	}
	evidence := parseBullets(doc.Sections["Evidence Signals"])
	penalties := []string{}
	if sym.SemanticConfidenceLevel == "LOW" {
		penalties = append(penalties, "Low semantic confidence score")
	}
	if inferred {
		penalties = append(penalties, "Resolved by normalized name instead of exact canonical name")
	}
	cb := map[string]any{"score": sym.ConfidenceScore, "level": sym.ConfidenceLevel, "rationale": sym.ConfidenceRationale}
	scb := map[string]any{"score": sym.SemanticConfidenceScore, "level": sym.SemanticConfidenceLevel, "rationale": sym.SemanticConfidenceRationale}
	return sym, cb, scb, evidence, penalties, strings.Join(trimAndCompact(doc.Sections["Confidence Assessment"]), " ")
}

func (s *Store) ExplainUsage(symbolName string) (model.SymbolSummary, map[string]any) {
	_, _, sym, _ := s.LookupSymbol(symbolName, "")
	doc := s.symbols[sym.ID]
	if doc == nil {
		return model.SymbolSummary{}, map[string]any{}
	}
	payload := map[string]any{
		"purpose":                 firstNonEmpty(doc.Sections["Description"]),
		"usage_context":           firstNonEmpty(doc.Sections["Seen In"]),
		"parameters":              strings.Join(trimAndCompact(doc.Sections["Parameters"]), "\n"),
		"returns":                 strings.Join(trimAndCompact(doc.Sections["Returns"]), "\n"),
		"side_effects":            strings.Join(trimAndCompact(doc.Sections["Side Effects"]), "\n"),
		"preconditions":           []string{},
		"postconditions":          []string{},
		"common_patterns":         parseBullets(doc.Sections["Used With"]),
		"pitfalls_or_uncertainty": parseBullets(doc.Sections["Notes"]),
		"examples":                firstN(parseBullets(doc.Sections["Examples"]), 10),
	}
	return sym, payload
}

func (s *Store) ResolveResource(uri string) (model.ResourceEnvelope, *model.APIError) {
	if text, ok := s.meta[uri]; ok {
		return model.ResourceEnvelope{URI: uri, Kind: "meta", Title: uri, Summary: "Canonical generated metadata", ContentType: mimeFromURI(uri), StructuredData: map[string]any{"body": text}, SourcePaths: []string{sourcePathForURI(uri)}}, nil
	}
	if strings.HasPrefix(uri, "war-api://symbols/") {
		name := strings.TrimPrefix(uri, "war-api://symbols/")
		_, _, sym, suggestions := s.LookupSymbol(name, "")
		if sym.ID == "" {
			return model.ResourceEnvelope{}, &model.APIError{ErrorCode: "not_found", ErrorMessage: "symbol resource not found", Suggestion: "Use search_symbols to inspect valid canonical names", CandidateMatches: suggestions}
		}
		return model.ResourceEnvelope{URI: uri, Kind: "symbol", Title: sym.Title, Summary: sym.Summary, ContentType: "application/json", StructuredData: sym, SourcePaths: []string{sym.DocPath}}, nil
	}
	if strings.HasPrefix(uri, "war-api://events/") {
		name := strings.TrimPrefix(uri, "war-api://events/")
		flow := s.EventFlow(name, true, true)
		return model.ResourceEnvelope{URI: uri, Kind: "event", Title: flow.Event.Title, Summary: flow.Event.Summary, ContentType: "application/json", StructuredData: flow, SourcePaths: []string{flow.Event.DocPath}, Warnings: flow.Warnings}, nil
	}
	if strings.HasPrefix(uri, "war-api://xml-handlers/") {
		name := strings.TrimPrefix(uri, "war-api://xml-handlers/")
		info := s.XMLHandlerInfo(name, true, true)
		return model.ResourceEnvelope{URI: uri, Kind: "xml_handler", Title: info.Handler.Title, Summary: info.Handler.Summary, ContentType: "application/json", StructuredData: info, SourcePaths: []string{info.Handler.DocPath}, Warnings: info.Warnings}, nil
	}
	if strings.HasPrefix(uri, "war-api://patterns/") {
		name := strings.TrimPrefix(uri, "war-api://patterns/")
		p := filepath.Join(s.root, "patterns", name+".md")
		b, err := os.ReadFile(p)
		if err != nil {
			return model.ResourceEnvelope{}, &model.APIError{ErrorCode: "not_found", ErrorMessage: "pattern resource not found"}
		}
		return model.ResourceEnvelope{URI: uri, Kind: "pattern", Title: name, Summary: "Pattern document", ContentType: "application/json", StructuredData: map[string]any{"name": name, "examples": parseBullets(strings.Split(string(b), "\n"))}, SourcePaths: []string{"patterns/" + name + ".md"}}, nil
	}
	if strings.HasPrefix(uri, "war-api://graph/") {
		name := strings.TrimPrefix(uri, "war-api://graph/")
		_, _, sym, _ := s.LookupSymbol(name, "")
		related := s.Related(sym, nil, 1, 50)
		return model.ResourceEnvelope{URI: uri, Kind: "graph", Title: sym.Title, Summary: "Graph neighborhood", ContentType: "application/json", StructuredData: related, SourcePaths: []string{"graph/api_graph.json", sym.DocPath}}, nil
	}
	return model.ResourceEnvelope{}, &model.APIError{ErrorCode: "unsupported_uri", ErrorMessage: "unsupported resource URI", Suggestion: "Use resources/list to discover supported URI patterns."}
}

func (s *Store) suggestions(name, symbolType string, limit int) []model.CandidateMatch {
	results := s.Search(name, symbolType, "", 0, limit)
	out := make([]model.CandidateMatch, 0, len(results))
	for _, r := range results {
		out = append(out, model.CandidateMatch{ID: r.ID, CanonicalName: r.CanonicalName, Type: r.Type, DocPath: r.DocPath})
	}
	return out
}

func pickByType(list []*symbolDoc, symbolType string) *symbolDoc {
	if symbolType == "" && len(list) > 0 {
		return list[0]
	}
	for _, d := range list {
		if strings.Contains(strings.ToLower(d.Summary.Type), strings.ToLower(symbolType)) {
			return d
		}
	}
	return nil
}

func buildStableID(kind, canonical string) string {
	switch kind {
	case schema.KindEvent:
		return schema.BuildID(schema.KindEvent, canonical)
	case schema.KindXMLHandler:
		return schema.BuildID(schema.KindXMLHandler, canonical)
	case schema.KindPattern:
		return schema.BuildID(schema.KindPattern, canonical)
	default:
		return schema.BuildID(schema.KindSymbol, canonical)
	}
}

var linkRE = regexp.MustCompile(`\[[^\]]+\]\(([^)]+)\)`)

func parseMarkdownSymbol(raw, rel string) *symbolDoc {
	lines := strings.Split(strings.ReplaceAll(raw, "\r\n", "\n"), "\n")
	sections := map[string][]string{}
	current := ""
	for _, l := range lines {
		t := strings.TrimSpace(l)
		if strings.HasPrefix(t, "## ") {
			current = strings.TrimPrefix(t, "## ")
			sections[current] = []string{}
			continue
		}
		if current != "" {
			sections[current] = append(sections[current], l)
		}
	}
	name := ""
	if len(lines) > 0 && strings.HasPrefix(strings.TrimSpace(lines[0]), "# ") {
		name = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(lines[0]), "# "))
	}
	typeVal := ""
	kind := schema.KindSymbol
	namespace := "global"
	for _, l := range lines {
		t := strings.TrimSpace(l)
		if strings.HasPrefix(t, "- Category:") {
			typeVal = strings.TrimSpace(strings.TrimPrefix(t, "- Category:"))
		}
		if strings.HasPrefix(t, "- Type:") && typeVal == "" {
			typeVal = strings.TrimSpace(strings.TrimPrefix(t, "- Type:"))
		}
	}
	ltype := strings.ToLower(typeVal)
	switch {
	case strings.Contains(ltype, "event"):
		kind = schema.KindEvent
	case strings.Contains(ltype, "xml") && strings.Contains(ltype, "handler"):
		kind = schema.KindXMLHandler
	}
	if idx := strings.Index(name, "."); idx > 0 {
		namespace = name[:idx]
	}
	summary := model.SymbolSummary{
		Kind:           kind,
		CanonicalName:  name,
		Title:          name,
		Type:           strings.ReplaceAll(strings.ToLower(typeVal), " ", "_"),
		Namespace:      namespace,
		DocPath:        rel,
		NavigationPath: rel,
		RelatedIDs:     []string{},
		Examples:       []string{},
		Warnings:       []model.Warning{},
		Metadata:       map[string]interface{}{},
	}
	doc := &symbolDoc{Summary: summary, Sections: sections, RawPath: rel}
	doc.Related = linksToIDs(sections["Related APIs"])
	doc.UsedWith = linksToIDs(sections["Used With"])
	doc.Triggered = linksToIDs(sections["Triggered By"])
	doc.Affects = linksToIDs(sections["Affects"])
	for _, ex := range parseBullets(sections["Examples"]) {
		doc.Summary.Examples = append(doc.Summary.Examples, ex)
	}
	for _, id := range append(append(append(doc.Related, doc.UsedWith...), doc.Triggered...), doc.Affects...) {
		doc.Summary.RelatedIDs = append(doc.Summary.RelatedIDs, id)
	}
	return doc
}

func linksToIDs(lines []string) []string {
	out := []string{}
	for _, l := range lines {
		m := linkRE.FindStringSubmatch(l)
		if len(m) < 2 {
			continue
		}
		path := normalizePath(resolveRelativePath(m[1]))
		if path == "" {
			continue
		}
		base := strings.TrimSuffix(filepath.Base(path), ".md")
		base = strings.TrimPrefix(base, "global_")
		base = strings.TrimPrefix(base, "window_")
		base = strings.TrimPrefix(base, "event_")
		base = strings.TrimPrefix(base, "handler_")
		out = append(out, schema.BuildID(schema.KindSymbol, base))
	}
	return out
}

func normalizePath(v string) string {
	v = strings.ReplaceAll(v, "\\", "/")
	v = strings.TrimPrefix(v, "./")
	return v
}

func resolveRelativePath(v string) string {
	v = strings.TrimSpace(v)
	for strings.HasPrefix(v, "../") {
		v = strings.TrimPrefix(v, "../")
	}
	return strings.TrimPrefix(v, "./")
}

func semanticScore(score int) int {
	if score > 10 {
		return score - 10
	}
	return score
}

func firstNonEmpty(lines []string) string {
	for _, l := range lines {
		t := strings.TrimSpace(l)
		if t != "" && !strings.HasPrefix(t, "-") {
			return t
		}
	}
	return ""
}

func parseBullets(lines []string) []string {
	out := []string{}
	for _, l := range lines {
		t := strings.TrimSpace(l)
		if strings.HasPrefix(t, "- ") {
			out = append(out, strings.TrimSpace(strings.TrimPrefix(t, "- ")))
		}
	}
	return out
}

func trimAndCompact(lines []string) []string {
	out := []string{}
	for _, l := range lines {
		t := strings.TrimSpace(l)
		if t != "" {
			out = append(out, t)
		}
	}
	return out
}

func firstN(values []string, n int) []string {
	if len(values) <= n {
		return values
	}
	return values[:n]
}

func addonFromExample(v string) string {
	if idx := strings.Index(v, ":"); idx > 0 {
		return strings.TrimSpace(v[:idx])
	}
	return "unknown"
}

func codeBlock(lines []string) string {
	inside := false
	buf := []string{}
	for _, l := range lines {
		t := strings.TrimSpace(l)
		if strings.HasPrefix(t, "```") {
			inside = !inside
			continue
		}
		if inside {
			buf = append(buf, strings.TrimSpace(l))
		}
	}
	return strings.Join(buf, "\n")
}

func mimeFromURI(uri string) string {
	if strings.HasSuffix(uri, "tree") || strings.HasSuffix(uri, "sidebar") {
		return "application/json"
	}
	if strings.Contains(uri, "meta") {
		return "text/markdown"
	}
	return "application/json"
}

func sourcePathForURI(uri string) string {
	switch uri {
	case "war-api://meta/conventions":
		return "meta/conventions.md"
	case "war-api://meta/inference-rules":
		return "meta/inference_rules.md"
	case "war-api://meta/coverage-report":
		return "meta/coverage_report.md"
	case "war-api://navigation/tree":
		return "navigation/tree.json"
	case "war-api://navigation/sidebar":
		return "navigation/sidebar.json"
	default:
		return ""
	}
}

func parseScore(line string) int {
	idx := strings.Index(line, "/100")
	if idx < 0 {
		return 0
	}
	start := idx - 1
	for start >= 0 && line[start] >= '0' && line[start] <= '9' {
		start--
	}
	n, _ := strconv.Atoi(strings.TrimSpace(line[start+1 : idx]))
	return n
}

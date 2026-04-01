package site

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type SearchIndex struct {
	GeneratedAt string        `json:"generatedAt"`
	Entries     []SearchEntry `json:"entries"`
}

type SearchEntry struct {
	Title    string   `json:"title"`
	Path     string   `json:"path"`
	Section  string   `json:"section,omitempty"`
	Summary  string   `json:"summary,omitempty"`
	Headings []string `json:"headings,omitempty"`
	Content  string   `json:"content"`
	Keywords []string `json:"keywords,omitempty"`
}

func Generate(sourceRoot, outputRoot string) error {
	if err := os.RemoveAll(outputRoot); err != nil {
		return fmt.Errorf("reset site content output: %w", err)
	}
	if err := os.MkdirAll(outputRoot, 0o755); err != nil {
		return fmt.Errorf("create site content output: %w", err)
	}

	entries := make([]SearchEntry, 0)
	err := filepath.WalkDir(sourceRoot, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		rel, err := filepath.Rel(sourceRoot, path)
		if err != nil {
			return fmt.Errorf("resolve relative path for %s: %w", path, err)
		}
		rel = filepath.ToSlash(rel)

		if d.IsDir() {
			if rel == "." {
				return nil
			}
			if rel == "scripts" || strings.HasPrefix(rel, "scripts/") {
				return filepath.SkipDir
			}
			return os.MkdirAll(filepath.Join(outputRoot, filepath.FromSlash(rel)), 0o755)
		}

		if !shouldCopyContent(rel) {
			return nil
		}

		destination := filepath.Join(outputRoot, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
			return fmt.Errorf("create directory for %s: %w", destination, err)
		}
		if err := copyFile(path, destination); err != nil {
			return err
		}

		if strings.EqualFold(filepath.Ext(rel), ".md") {
			entry, err := indexMarkdown(path, rel)
			if err != nil {
				return err
			}
			entries = append(entries, entry)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("copy site content: %w", err)
	}

	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Path) < strings.ToLower(entries[j].Path)
	})

	index := SearchIndex{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Entries:     entries,
	}
	if err := writeJSON(filepath.Join(outputRoot, "search", "index.json"), index); err != nil {
		return err
	}
	return nil
}

func shouldCopyContent(relativePath string) bool {
	if strings.EqualFold(relativePath, "index.html") {
		return false
	}
	for _, prefix := range []string{"scripts/", "search/"} {
		if strings.HasPrefix(relativePath, prefix) {
			return false
		}
	}
	switch strings.ToLower(filepath.Ext(relativePath)) {
	case ".json", ".md":
		return true
	default:
		return false
	}
}

func copyFile(sourcePath, destinationPath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("open %s: %w", sourcePath, err)
	}
	defer source.Close()

	destination, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("create %s: %w", destinationPath, err)
	}
	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		return fmt.Errorf("copy %s to %s: %w", sourcePath, destinationPath, err)
	}
	return nil
}

func indexMarkdown(path, relativePath string) (SearchEntry, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return SearchEntry{}, fmt.Errorf("read markdown for search index: %w", err)
	}

	lines := strings.Split(string(contents), "\n")
	title := strings.TrimSuffix(filepath.Base(relativePath), filepath.Ext(relativePath))
	summary := ""
	headings := make([]string, 0)
	body := make([]string, 0, len(lines))
	keywords := map[string]struct{}{}

	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if strings.HasPrefix(line, "# ") && title == strings.TrimSuffix(filepath.Base(relativePath), filepath.Ext(relativePath)) {
			title = strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
		if strings.HasPrefix(line, "##") {
			heading := strings.TrimSpace(strings.TrimLeft(line, "#"))
			if heading != "" {
				headings = append(headings, heading)
				keywords[strings.ToLower(heading)] = struct{}{}
			}
		}
		plain := normalizeMarkdownLine(line)
		if plain == "" {
			continue
		}
		if summary == "" && !strings.HasPrefix(line, "#") {
			summary = plain
		}
		body = append(body, plain)
		for _, token := range strings.Fields(strings.ToLower(plain)) {
			if len(token) >= 4 {
				keywords[token] = struct{}{}
			}
		}
	}

	section := ""
	parts := strings.Split(relativePath, "/")
	if len(parts) > 1 {
		section = parts[0]
	}

	keywordList := make([]string, 0, len(keywords))
	for keyword := range keywords {
		keywordList = append(keywordList, keyword)
	}
	sort.Strings(keywordList)

	return SearchEntry{
		Title:    title,
		Path:     relativePath,
		Section:  section,
		Summary:  summary,
		Headings: headings,
		Content:  strings.Join(body, "\n"),
		Keywords: keywordList,
	}, nil
}

func normalizeMarkdownLine(line string) string {
	replacer := strings.NewReplacer(
		"`", "",
		"*", "",
		"_", " ",
		"[", "",
		"]", "",
		"(", "",
		")", "",
		"<", "",
		">", "",
		"|", " ",
	)
	cleaned := replacer.Replace(line)
	cleaned = strings.TrimSpace(strings.TrimLeft(cleaned, "-0123456789. "))
	return strings.Join(strings.Fields(cleaned), " ")
}

func writeJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create %s: %w", filepath.Dir(path), err)
	}
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal %s: %w", path, err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

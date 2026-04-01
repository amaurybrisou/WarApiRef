package schema

import (
	"regexp"
	"strings"
)

const (
	KindSymbol     = "symbol"
	KindEvent      = "event"
	KindXMLHandler = "xml_handler"
	KindPattern    = "pattern"
	KindPage       = "page"
)

var unsupportedChars = regexp.MustCompile(`[^a-z0-9_.:/-]+`)

// NormalizeName creates a case-insensitive key for lookup and indexing.
func NormalizeName(v string) string {
	v = strings.TrimSpace(strings.ToLower(v))
	v = strings.ReplaceAll(v, " ", "")
	v = strings.ReplaceAll(v, "-", "")
	v = strings.ReplaceAll(v, "_", "")
	v = strings.ReplaceAll(v, ".", "")
	v = strings.ReplaceAll(v, ":", "")
	return unsupportedChars.ReplaceAllString(v, "")
}

// BuildID creates stable entity IDs used across tools/resources.
func BuildID(kind, canonicalName string) string {
	canonicalName = strings.TrimSpace(canonicalName)
	return kind + ":" + canonicalName
}

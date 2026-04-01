package deep_analysis

import (
	"sort"
	"strings"
)

// UsageCluster represents a group of symbols frequently used together
type UsageCluster struct {
	Name       string   // cluster identifier
	Symbols    []string // symbols in this cluster
	Frequency  int      // how often seen together
	Patterns   []string // identified usage patterns
	Evidence   []string // supporting evidence files
	Confidence int      // 0-100 confidence this is a real cluster
}

// UsageClustering analyzes symbol co-occurrence patterns
type UsageClustering struct {
	CoOccurrenceMatrix map[string]map[string]int // symbol1 -> symbol2 -> count
	SymbolFrequency    map[string]int            // how often each symbol appears
	AllSymbols         map[string]bool           // all symbols seen
	Clusters           []UsageCluster
}

// NewUsageClustering creates clustering analyzer
func NewUsageClustering() *UsageClustering {
	return &UsageClustering{
		CoOccurrenceMatrix: make(map[string]map[string]int),
		SymbolFrequency:    make(map[string]int),
		AllSymbols:         make(map[string]bool),
	}
}

// RecordCoOccurrence records two symbols appearing together
func (uc *UsageClustering) RecordCoOccurrence(symbol1, symbol2 string) {
	// Normalize order to avoid duplicates
	if symbol1 > symbol2 {
		symbol1, symbol2 = symbol2, symbol1
	}

	if _, ok := uc.CoOccurrenceMatrix[symbol1]; !ok {
		uc.CoOccurrenceMatrix[symbol1] = make(map[string]int)
	}

	uc.CoOccurrenceMatrix[symbol1][symbol2]++
	uc.SymbolFrequency[symbol1]++
	uc.SymbolFrequency[symbol2]++
	uc.AllSymbols[symbol1] = true
	uc.AllSymbols[symbol2] = true
}

// AnalyzeSymbolCalls identifies which symbols call each other frequently
func (uc *UsageClustering) AnalyzeSymbolCalls(functionPath string, calledFunctions []string) {
	// Record pairwise co-occurrences
	for i, called1 := range calledFunctions {
		uc.SymbolFrequency[called1]++
		uc.AllSymbols[called1] = true

		// Record co-occurrence with other calls in same function
		for _, called2 := range calledFunctions[i+1:] {
			uc.RecordCoOccurrence(called1, called2)
		}
	}
}

// BuildClusters identifies natural groupings of related symbols
func (uc *UsageClustering) BuildClusters() []UsageCluster {
	clusters := []UsageCluster{}

	// Use simple threshold-based clustering
	visited := make(map[string]bool)

	for symbol1 := range uc.AllSymbols {
		if visited[symbol1] {
			continue
		}

		// Start a new cluster with this symbol
		cluster := UsageCluster{
			Name:       symbol1,
			Symbols:    []string{symbol1},
			Frequency:  uc.SymbolFrequency[symbol1],
			Confidence: 100, // seed symbol
		}
		visited[symbol1] = true

		// Find symbols that co-occur frequently with this one
		relatedPairs := uc.GetRelatedSymbols(symbol1, 3) // threshold: 3+ co-occurrences
		for _, pair := range relatedPairs {
			symbol2 := pair.Symbol
			count := pair.Count

			if !visited[symbol2] {
				cluster.Symbols = append(cluster.Symbols, symbol2)
				cluster.Frequency += count
				visited[symbol2] = true
			}
		}

		// Only include clusters with multiple symbols
		if len(cluster.Symbols) > 1 {
			// Calculate cluster confidence
			avgFrequency := cluster.Frequency / len(cluster.Symbols)
			cluster.Confidence = calculateClusterConfidence(avgFrequency, len(cluster.Symbols))
			clusters = append(clusters, cluster)
		}
	}

	uc.Clusters = clusters
	return clusters
}

// SymbolPair represents a symbol and its relationship strength
type SymbolPair struct {
	Symbol string
	Count  int
}

// GetRelatedSymbols finds symbols related to a given symbol (by co-occurrence)
func (uc *UsageClustering) GetRelatedSymbols(symbol string, minCooccur int) []SymbolPair {
	pairs := []SymbolPair{}

	// Check both directions in co-occurrence matrix
	if related, ok := uc.CoOccurrenceMatrix[symbol]; ok {
		for relatedSymbol, count := range related {
			if count >= minCooccur {
				pairs = append(pairs, SymbolPair{Symbol: relatedSymbol, Count: count})
			}
		}
	}

	// Check reverse direction
	for firstSymbol, related := range uc.CoOccurrenceMatrix {
		if count, ok := related[symbol]; ok && count >= minCooccur {
			if firstSymbol != symbol {
				pairs = append(pairs, SymbolPair{Symbol: firstSymbol, Count: count})
			}
		}
	}

	// Sort by count descending
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	return pairs
}

// IdentifyPatterns detects common patterns in how symbols are used
func (uc *UsageClustering) IdentifyPatterns() map[string][]string {
	patterns := make(map[string][]string)

	for _, cluster := range uc.Clusters {
		// Pattern: UI related cluster
		if uc.isUICluster(cluster.Symbols) {
			patterns["ui_update"] = append(patterns["ui_update"], cluster.Name)
		}

		// Pattern: Event handling cluster
		if uc.isEventCluster(cluster.Symbols) {
			patterns["event_handling"] = append(patterns["event_handling"], cluster.Name)
		}

		// Pattern: Data access cluster
		if uc.isDataCluster(cluster.Symbols) {
			patterns["data_access"] = append(patterns["data_access"], cluster.Name)
		}

		// Pattern: Utility/helper cluster
		if uc.isUtilityCluster(cluster.Symbols) {
			patterns["utility"] = append(patterns["utility"], cluster.Name)
		}
	}

	return patterns
}

// isUICluster detects if a cluster is related to UI operations
func (uc *UsageClustering) isUICluster(symbols []string) bool {
	uiKeywords := []string{
		"Show", "Hide", "SetShowing",
		"Anchor", "Position", "Size",
		"Invalidate", "Refresh",
		"Update", "Draw",
	}

	matchCount := 0
	for _, symbol := range symbols {
		for _, keyword := range uiKeywords {
			if strings.Contains(symbol, keyword) {
				matchCount++
			}
		}
	}

	return matchCount >= len(symbols)/2 // at least half match
}

// isEventCluster detects if a cluster is related to event handling
func (uc *UsageClustering) isEventCluster(symbols []string) bool {
	eventKeywords := []string{
		"Event", "Handle", "Register",
		"Subscribe", "Unsubscribe",
		"Trigger", "Emit",
	}

	matchCount := 0
	for _, symbol := range symbols {
		for _, keyword := range eventKeywords {
			if strings.Contains(symbol, keyword) {
				matchCount++
			}
		}
	}

	return matchCount >= len(symbols)/2
}

// isDataCluster detects if a cluster is related to data access
func (uc *UsageClustering) isDataCluster(symbols []string) bool {
	dataKeywords := []string{
		"Data", "Table", "Get", "Set",
		"Cache", "Store", "Load",
	}

	matchCount := 0
	for _, symbol := range symbols {
		for _, keyword := range dataKeywords {
			if strings.Contains(symbol, keyword) {
				matchCount++
			}
		}
	}

	return matchCount >= len(symbols)/2
}

// isUtilityCluster detects if a cluster is utility functions
func (uc *UsageClustering) isUtilityCluster(symbols []string) bool {
	// Utility clusters are smaller and don't fit other patterns
	return len(symbols) <= 3
}

// FindCommonlyUsedWith finds symbols that should have "commonly_used_with" edges
func (uc *UsageClustering) FindCommonlyUsedWith(minCooccur int) []EdgeRelation {
	edges := []EdgeRelation{}

	for symbol1, related := range uc.CoOccurrenceMatrix {
		for symbol2, count := range related {
			if count >= minCooccur {
				edges = append(edges, EdgeRelation{
					From:            symbol1,
					To:              symbol2,
					RelationType:    "commonly_used_with",
					ConfidenceScore: calculateCooccurrenceConfidence(count),
					EvidenceCount:   count,
					Rationale:       "Used together " + string(rune(count)) + " times",
				})
			}
		}
	}

	return edges
}

// calculateClusterConfidence determines how strong a cluster is
func calculateClusterConfidence(avgFrequency, clusterSize int) int {
	confidence := 50 // base

	// More frequent co-occurrence = higher confidence
	confidence += (avgFrequency - 1) * 10
	if confidence > 80 {
		confidence = 80
	}

	// Larger clusters are less confident
	if clusterSize > 5 {
		confidence -= 5
	}

	return confidence
}

// calculateCooccurrenceConfidence determines confidence in a co-occurrence relationship
func calculateCooccurrenceConfidence(count int) int {
	confidence := 50 // base

	// Each occurrence adds confidence
	confidence += (count - 1) * 5

	// Cap at 75 for co-occurrence (weaker signal than direct calls)
	if confidence > 75 {
		confidence = 75
	}

	return confidence
}

// SummarizeCluster provides a text summary of a cluster
func (uc *UsageClustering) SummarizeCluster(cluster UsageCluster) string {
	// Find the pattern this cluster represents
	patterns := uc.IdentifyPatterns()
	patternType := "mixed"
	for pType, clusters := range patterns {
		for _, c := range clusters {
			if c == cluster.Name {
				patternType = pType
				break
			}
		}
	}

	summary := "Symbols: "
	for i, sym := range cluster.Symbols {
		if i > 0 {
			summary += ", "
		}
		summary += sym
	}

	summary += "\n" +
		"Type: " + patternType + "\n" +
		"Co-occurrence frequency: " + string(rune(cluster.Frequency)) + "\n" +
		"Confidence: " + string(rune(cluster.Confidence)) + "%"

	return summary
}

// GetClusterMembers returns all symbols in a named cluster
func (uc *UsageClustering) GetClusterMembers(clusterName string) []string {
	for _, cluster := range uc.Clusters {
		if cluster.Name == clusterName {
			return cluster.Symbols
		}
	}
	return []string{}
}

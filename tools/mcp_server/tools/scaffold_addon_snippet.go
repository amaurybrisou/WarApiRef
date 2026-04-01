package tools

import (
	"fmt"
	"strings"

	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
)

func ScaffoldAddonSnippet(b Backend, req schema.ScaffoldAddonSnippetRequest) schema.ScaffoldAddonSnippetResponse {
	used := []model.SymbolSummary{}
	avoided := []string{}
	for _, desired := range req.DesiredAPIs {
		exact, found, sym, _ := b.LookupSymbol(desired, "")
		if found && (exact || sym.ConfidenceScore >= 70) {
			used = append(used, sym)
			continue
		}
		avoided = append(avoided, desired)
	}

	comment := ""
	if req.IncludeComments {
		comment = "-- Generated from WAR API MCP scaffold contract\n"
	}
	lua := comment + "local Addon = {}\n\nfunction Addon.Initialize()\n"
	for _, s := range used {
		lua += fmt.Sprintf("    -- Canonical API: %s (%s)\n", s.CanonicalName, s.DocPath)
		lua += fmt.Sprintf("    local _ = %s\n", scaffoldCall(s.CanonicalName))
	}
	lua += "end\n\nreturn Addon\n"
	artifacts := []schema.GeneratedArtifact{{Type: "lua", Filename: "AddonCore.lua", Content: lua}}
	if req.IncludeXML {
		xml := "<Window name=\"AddonWindow\" movable=\"true\" handleinput=\"true\" skipinput=\"true\"></Window>"
		artifacts = append(artifacts, schema.GeneratedArtifact{Type: "xml", Filename: "AddonUI.xml", Content: xml})
	}
	return schema.ScaffoldAddonSnippetResponse{
		GeneratedArtifacts:    artifacts,
		UsedDocumentedSymbols: used,
		AvoidedSymbols:        avoided,
		UncertaintyNotes: []string{
			"Only canonical or high-confidence symbols are used.",
			"Any unresolved desired API is listed in avoided_symbols.",
		},
		Rationale: "Scaffold favors documented high-confidence APIs and explicit uncertainty.",
	}
}

func scaffoldCall(name string) string {
	if strings.Contains(name, ".") {
		return name + "()"
	}
	return name
}

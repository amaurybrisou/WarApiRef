package main

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseXMLFileResolvesDerivedNamesForAdvancedRenownTraining(t *testing.T) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to resolve caller path")
	}
	xmlPath := filepath.Join(filepath.Dir(currentFile), "..", "..", "addons", "advancedrenowntrainer", "AdvancedRenownTrainingPresets.xml")

	parsed, err := parseXMLFile("advancedrenowntrainer", xmlPath)
	if err != nil {
		t.Fatalf("parseXMLFile failed: %v", err)
	}

	framesByName := map[string]graphFrameExpectation{}
	for _, frame := range parsed.Frames {
		framesByName[frame.Name] = graphFrameExpectation{rawName: frame.RawName, parent: frame.Parent}
	}

	background, ok := framesByName["AdvancedRenownTrainingPresetsWindowBackground"]
	if !ok {
		t.Fatalf("expected resolved background frame to be present")
	}
	if background.rawName != "$parentBackground" {
		t.Fatalf("expected raw name $parentBackground, got %q", background.rawName)
	}
	if background.parent != "AdvancedRenownTrainingPresetsWindow" {
		t.Fatalf("expected resolved parent frame name, got %q", background.parent)
	}

	if _, ok := framesByName["AdvancedRenownTrainingPresetsWindowSaveButton"]; !ok {
		t.Fatalf("expected resolved save button frame to be present")
	}

	resolvedHandler := false
	for _, handler := range parsed.Handlers {
		if handler.Frame == "AdvancedRenownTrainingPresetsWindowSaveButton" && handler.Event == "OnLButtonUp" && handler.Function == "AdvancedRenownTraining.SavePreset" {
			resolvedHandler = true
			break
		}
	}
	if !resolvedHandler {
		t.Fatalf("expected handler to bind against resolved save button frame name")
	}
}

type graphFrameExpectation struct {
	rawName string
	parent  string
}

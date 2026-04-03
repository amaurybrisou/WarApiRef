package platform

import "testing"

func TestRenderSemanticSectionsOmitsEmptySections(t *testing.T) {
	content := renderSemanticSections("xml/element_types/element_Text.md", SemanticLinks{})
	if content != "" {
		t.Fatalf("expected no semantic relationship sections for empty links, got:\n%s", content)
	}
}

func TestBuildSemanticLinksContainsProducesBidirectionalRelatedLinks(t *testing.T) {
	catalog := symbolCatalog{
		entries: map[string]DocLink{
			"ui:Label": {ID: "ui:Label", Label: "Label", Type: "ui_element", Path: "xml/element_types/element_Label.md"},
			"ui:Text":  {ID: "ui:Text", Label: "Text", Type: "ui_element", Path: "xml/element_types/element_Text.md"},
		},
	}
	links := buildSemanticLinks(catalog, []GraphEdge{{From: "ui:Label", To: "ui:Text", Type: "contains"}})

	if len(links["ui:Label"].RelatedAPIs) != 1 || links["ui:Label"].RelatedAPIs[0].ID != "ui:Text" {
		t.Fatalf("expected Label related links to include Text, got %+v", links["ui:Label"].RelatedAPIs)
	}
	if len(links["ui:Text"].RelatedAPIs) != 1 || links["ui:Text"].RelatedAPIs[0].ID != "ui:Label" {
		t.Fatalf("expected Text related links to include Label, got %+v", links["ui:Text"].RelatedAPIs)
	}
}

func TestBuildSemanticGraphEmitsContainsForStructuralChildEvidence(t *testing.T) {
	corpus := Corpus{
		ElementTypes: []ElementTypeSymbol{
			{
				Name:                "Label",
				StructuralChildRefs: []ElementRelRef{{Tag: "Text", Count: 60, Confidence: "HIGH"}},
			},
			{
				Name: "Text",
			},
		},
	}
	catalog := buildSymbolCatalog(corpus)
	graphData, _, _ := buildSemanticGraph(corpus, catalog, contractSemanticInput{})

	found := false
	for _, edge := range graphData.Edges {
		if edge.Type == "contains" && edge.From == "ui:Label" && edge.To == "ui:Text" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected contains edge ui:Label -> ui:Text from structural child evidence")
	}
}

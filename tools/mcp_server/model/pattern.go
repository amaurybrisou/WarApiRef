package model

type UsageExample struct {
	AddonName        string   `json:"addon_name"`
	Title            string   `json:"title"`
	Snippet          string   `json:"snippet"`
	Explanation      string   `json:"explanation"`
	DocPath          string   `json:"doc_path"`
	RelatedSymbolIDs []string `json:"related_symbol_ids"`
}

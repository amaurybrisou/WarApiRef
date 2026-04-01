package model

type XMLHandlerInfo struct {
	Handler               SymbolSummary   `json:"handler"`
	LifecycleRole         string          `json:"lifecycle_role"`
	ExpectedCallbackShape string          `json:"expected_callback_shape"`
	CommonElementTypes    []string        `json:"common_element_types"`
	PairedSymbols         []SymbolSummary `json:"paired_symbols"`
	Examples              []string        `json:"examples"`
	Warnings              []Warning       `json:"warnings"`
}

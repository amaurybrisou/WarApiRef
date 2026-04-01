package model

type EventFlow struct {
	Event             SymbolSummary   `json:"event"`
	ObservedHandlers  []SymbolSummary `json:"observed_handlers"`
	CommonFlows       []string        `json:"common_flows"`
	DownstreamSymbols []SymbolSummary `json:"downstream_symbols"`
	Examples          []string        `json:"examples"`
	Warnings          []Warning       `json:"warnings"`
}

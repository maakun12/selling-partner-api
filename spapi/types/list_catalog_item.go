package types

type ListCatalogItemResponse struct {
	Payload ListCatalogItemPayload `json:"payload"`
}

type ListCatalogItems struct {
	Identifiers   Identifiers     `json:"Identifiers"`
	AttributeSets []AttributeSets `json:"AttributeSets"`
	Relationships []Relationships `json:"Relationships"`
	SalesRankings []SalesRankings `json:"SalesRankings"`
}

type ListCatalogItemPayload struct {
	Items []ListCatalogItems `json:"Items"`
}

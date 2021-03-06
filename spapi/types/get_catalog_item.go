package types

type GetCatalogItemResponse struct {
	Payload GetCatalogItemPayload `json:"payload"`
}
type GetCatalogItemPayload struct {
	Identifiers   Identifiers     `json:"Identifiers"`
	AttributeSets []AttributeSets `json:"AttributeSets"`
	Relationships []Relationships `json:"Relationships"`
	SalesRankings []SalesRankings `json:"SalesRankings"`
}

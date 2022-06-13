package types

import "time"

type GetListingsItemResponse struct {
	Sku       string      `json:"sku"`
	Summaries []Summaries `json:"summaries"`
}
type MainImage struct {
	Link   string `json:"link"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}
type Summaries struct {
	MarketplaceID   string    `json:"marketplaceId"`
	Asin            string    `json:"asin"`
	ProductType     string    `json:"productType"`
	ConditionType   string    `json:"conditionType"`
	Status          []string  `json:"status"`
	FnSku           string    `json:"fnSku"`
	ItemName        string    `json:"itemName"`
	CreatedDate     time.Time `json:"createdDate"`
	LastUpdatedDate time.Time `json:"lastUpdatedDate"`
	MainImage       MainImage `json:"mainImage"`
}

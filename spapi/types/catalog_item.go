package types

type MarketplaceASIN struct {
	MarketplaceID string `json:"MarketplaceID"`
	Asin          string `json:"ASIN"`
}
type Identifiers struct {
	MarketplaceASIN MarketplaceASIN `json:"MarketplaceASIN"`
}
type Weight struct {
	Value float64 `json:"value"`
	Units string  `json:"Units"`
}
type ItemDimensions struct {
	Weight Weight `json:"Weight"`
}
type ListPrice struct {
	Amount       float64 `json:"Amount"`
	CurrencyCode string  `json:"CurrencyCode"`
}
type Height struct {
	Value float64 `json:"value"`
	Units string  `json:"Units"`
}
type Length struct {
	Value float64 `json:"value"`
	Units string  `json:"Units"`
}
type Width struct {
	Value float64 `json:"value"`
	Units string  `json:"Units"`
}
type PackageDimensions struct {
	Height Height `json:"Height"`
	Length Length `json:"Length"`
	Width  Width  `json:"Width"`
	Weight Weight `json:"Weight"`
}
type SmallImage struct {
	URL    string `json:"URL"`
	Height Height `json:"Height"`
	Width  Width  `json:"Width"`
}
type AttributeSets struct {
	Platform          []string          `json:"Platform"`
	Binding           string            `json:"Binding"`
	Brand             string            `json:"Brand"`
	Color             string            `json:"Color"`
	ItemDimensions    ItemDimensions    `json:"ItemDimensions"`
	Label             string            `json:"Label"`
	ListPrice         ListPrice         `json:"ListPrice"`
	Manufacturer      string            `json:"Manufacturer"`
	PackageDimensions PackageDimensions `json:"PackageDimensions"`
	PackageQuantity   int               `json:"PackageQuantity"`
	PartNumber        string            `json:"PartNumber"`
	ProductGroup      string            `json:"ProductGroup"`
	ProductTypeName   string            `json:"ProductTypeName"`
	Publisher         string            `json:"Publisher"`
	ReleaseDate       string            `json:"ReleaseDate"`
	Size              string            `json:"Size"`
	SmallImage        SmallImage        `json:"SmallImage"`
	Studio            string            `json:"Studio"`
	Title             string            `json:"Title"`
}
type Relationships struct {
	Identifiers Identifiers `json:"Identifiers"`
}
type SalesRankings struct {
	ProductCategoryID string `json:"ProductCategoryId"`
	Rank              int    `json:"Rank"`
}

package types

type GetMarketplaceParticipationsResponse struct {
	Payload []SellersPayload `json:"payload"`
}
type Marketplace struct {
	ID                  string `json:"id"`
	CountryCode         string `json:"countryCode"`
	Name                string `json:"name"`
	DefaultCurrencyCode string `json:"defaultCurrencyCode"`
	DefaultLanguageCode string `json:"defaultLanguageCode"`
	DomainName          string `json:"domainName"`
}
type Participation struct {
	IsParticipating      bool `json:"isParticipating"`
	HasSuspendedListings bool `json:"hasSuspendedListings"`
}
type SellersPayload struct {
	Marketplace   Marketplace   `json:"marketplace"`
	Participation Participation `json:"participation"`
}

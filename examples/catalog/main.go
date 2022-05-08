package main

import (
	"context"
	"fmt"

	"github.com/maakun12/selling-partner-api/spapi"
)

func main() {
	c, err := spapi.NewClient(
		&spapi.Config{
			RefreshToken:  "<RefreshToken>",
			ClientID:      "<ClientID>",
			ClientSecret:  "<ClientSecret>",
			AccessKey:     "<AccessKey>",
			SecretKey:     "<SecretKey>",
			Endpoint:      "sellingpartnerapi-fe.amazon.com",
			MarketplaceId: "A1VC38T7YXB528",
			Region:        "us-west-2",
		},
	)
	if err != nil {
		panic(err)
	}

	res, err := c.GetCatalogItem(context.Background(), "B092J9FZL2")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

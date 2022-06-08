package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/maakun12/selling-partner-api/spapi"
)

func main() {
	mode := flag.String("mode", "GetCatalogItem", "[GetCatalogItem|ListCatalogItem]")
	flag.Parse()

	c, err := spapi.NewClient(
		&spapi.Config{
			RefreshToken:  os.Getenv("LWA_REFRESH_TOKEN"),
			ClientID:      os.Getenv("LWA_CLIENT_ID"),
			ClientSecret:  os.Getenv("LWA_CLIENT_SECRET"),
			AccessKey:     os.Getenv("AWS_ACCESS_KEY"),
			SecretKey:     os.Getenv("AWS_SECRET_KEY"),
			Endpoint:      "sellingpartnerapi-fe.amazon.com",
			MarketplaceID: "A1VC38T7YXB528",
			Region:        "us-west-2",
		},
	)
	if err != nil {
		panic(err)
	}

	switch *mode {
	case "GetCatalogItem":
		res, err := c.GetCatalogItem(context.Background(), "B092J9FZL2")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	case "ListCatalogItem":
		res, err := c.ListCatalogItem(context.Background(), "Nintendo Switch")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/maakun12/selling-partner-api/spapi"
)

func main() {
	mode := flag.String("mode", "GetListingsItem", "[GetListingsItem|PutListingsItem|DeleteListingsItem]")
	flag.Parse()

	c, err := spapi.NewClient(
		&spapi.Config{
			RefreshToken:  os.Getenv("LWA_REFRESH_TOKEN"),
			ClientID:      os.Getenv("LWA_CLIENT_ID"),
			ClientSecret:  os.Getenv("LWA_CLIENT_SECRET"),
			AccessKey:     os.Getenv("AWS_ACCESS_KEY"),
			SecretKey:     os.Getenv("AWS_SECRET_KEY"),
			Endpoint:      "sandbox.sellingpartnerapi-fe.amazon.com",
			MarketplaceID: "A1VC38T7YXB528",
			Region:        "us-west-2",
		},
	)
	if err != nil {
		panic(err)
	}

	switch *mode {
	case "GetListingsItem":
		res, err := c.GetListingsItem(context.Background(), "example-seller-id", "example-sku")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	case "PutListingsItem":
		body := map[string]interface{}{
			"productType":  "KEYBOARDS",
			"requirements": "LISTING",
			"attributes": map[string]interface{}{
				"condition_type": []map[string]interface{}{
					{
						"value": "new_new",
					},
				},
				"item_name": []map[string]interface{}{
					{
						"value": "Sample Item Name",
					},
				},
			},
		}
		res, err := c.PutListingsItem(context.Background(), "example-seller-id", "example-sku", body)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	case "DeleteListingsItem":
		res, err := c.DeleteListingsItem(context.Background(), "example-seller-id", "example-sku")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}

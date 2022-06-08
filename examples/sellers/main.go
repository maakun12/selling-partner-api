package main

import (
	"context"
	"fmt"
	"os"

	"github.com/maakun12/selling-partner-api/spapi"
)

func main() {
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

	res, err := c.GetMarketplaceParticipations(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

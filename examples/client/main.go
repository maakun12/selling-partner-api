package main

import (
	"fmt"

	"github.com/maakun12/selling-partner-api/spapi"
)

func main() {
	c, err := spapi.NewClient(&spapi.Config{
		RefreshToken: "<RefreshToken>",
		ClientID:     "<ClientID>",
		ClientSecret: "<ClientSecret>",
		AccessKey:    "<AccessKey>",
		SecretKey:    "<SecretKey>",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(c)
}

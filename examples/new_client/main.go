package main

import (
	"fmt"

	"github.com/maakun12/selling-partner-api/spapi"
)

func main() {
	c, err := spapi.NewClient(&spapi.Credential{
		RefreshToken: "<RefreshToken>",
		ClientID:     "<ClientID>",
		ClientSecret: "<ClientSecret>",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(c)
}

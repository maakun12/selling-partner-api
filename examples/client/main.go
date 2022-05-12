package main

import (
	"fmt"
	"os"

	"github.com/maakun12/selling-partner-api/spapi"
)

func main() {
	c, err := spapi.NewClient(&spapi.Config{
		RefreshToken: os.Getenv("LWA_REFRESH_TOKEN"),
		ClientID:     os.Getenv("LWA_CLIENT_ID"),
		ClientSecret: os.Getenv("LWA_CLIENT_SECRET"),
		AccessKey:    os.Getenv("AWS_ACCESS_KEY"),
		SecretKey:    os.Getenv("AWS_SECRET_KEY"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(c)
}

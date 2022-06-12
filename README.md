# selling-partner-api

[![Build Status](https://github.com/maakun12/selling-partner-api/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/maakun12/selling-partner-api/actions)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/maakun12/selling-partner-api/spapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/maakun12/selling-partner-api)](https://goreportcard.com/report/github.com/maakun12/selling-partner-api)

## Installation

```sh
go get -u github.com/maakun12/selling-partner-api
```

## Configuration

```go
func main() {
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
    ...
}
```

see: https://developer-docs.amazon.com/sp-api/docs/connecting-to-the-selling-partner-api#step-1-request-a-login-with-amazon-access-token

## Example

```go
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
			Endpoint:      "sellingpartnerapi-fe.amazon.com",
			MarketplaceID: "A1VC38T7YXB528",
			Region:        "us-west-2",
		},
	)
	if err != nil {
		panic(err)
	}

	res, err := c.ListCatalogItem(context.Background(), "Nintendo Switch")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

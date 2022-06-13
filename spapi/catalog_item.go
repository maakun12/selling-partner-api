package spapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/maakun12/selling-partner-api/spapi/types"
)

func (c *Client) GetCatalogItem(ctx context.Context, asin string) (*types.GetCatalogItemResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf(APIEndpointGetCatalogItem, c.Config.Endpoint, asin),
		nil,
	)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("MarketplaceId", c.Config.MarketplaceID)
	req.URL.RawQuery = v.Encode()

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed. StatusCode=%v, Body=%v", resp.StatusCode, string(byteArray))
	}

	res := &types.GetCatalogItemResponse{}
	if err = json.Unmarshal(byteArray, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) ListCatalogItem(ctx context.Context, query string) (*types.ListCatalogItemResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf(APIEndpointListCatalogItem, c.Config.Endpoint),
		nil,
	)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("MarketplaceId", c.Config.MarketplaceID)
	v.Add("Query", query)
	req.URL.RawQuery = v.Encode()

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed. StatusCode=%v, Body=%v", resp.StatusCode, string(byteArray))
	}

	res := &types.ListCatalogItemResponse{}
	if err = json.Unmarshal(byteArray, res); err != nil {
		return nil, err
	}

	return res, nil
}

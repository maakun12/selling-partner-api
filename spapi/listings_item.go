package spapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/maakun12/selling-partner-api/spapi/types"
)

func (c *Client) GetListingsItem(ctx context.Context, sellerID, sku string) (*types.GetListingsItemResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(APIEndpointListingsItem, c.Config.Endpoint, sellerID, sku),
		nil,
	)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("marketplaceIds", c.Config.MarketplaceID)
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

	res := &types.GetListingsItemResponse{}
	if err = json.Unmarshal(byteArray, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) PutListingsItem(ctx context.Context, sellerID, sku string, body map[string]interface{}) (*types.PutListingsItemResponse, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		fmt.Sprintf(APIEndpointListingsItem, c.Config.Endpoint, sellerID, sku),
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("marketplaceIds", c.Config.MarketplaceID)
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

	res := &types.PutListingsItemResponse{}
	if err = json.Unmarshal(byteArray, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) DeleteListingsItem(ctx context.Context, sellerID, sku string) (*types.DeleteListingsItemResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf(APIEndpointListingsItem, c.Config.Endpoint, sellerID, sku),
		nil,
	)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("marketplaceIds", c.Config.MarketplaceID)
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

	res := &types.DeleteListingsItemResponse{}
	if err = json.Unmarshal(byteArray, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) PatchListingsItem(ctx context.Context, sellerID, sku string, body map[string]interface{}) (*types.PatchListingsItemResponse, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPatch,
		fmt.Sprintf(APIEndpointListingsItem, c.Config.Endpoint, sellerID, sku),
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("marketplaceIds", c.Config.MarketplaceID)
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

	res := &types.PatchListingsItemResponse{}
	if err = json.Unmarshal(byteArray, res); err != nil {
		return nil, err
	}

	return res, nil
}

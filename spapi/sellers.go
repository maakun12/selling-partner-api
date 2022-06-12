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

func (c *Client) GetMarketplaceParticipations(ctx context.Context) (*types.GetMarketplaceParticipationsResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf(APIEndpointGetMarketplaceParticipations, c.Config.Endpoint),
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

	res := &types.GetMarketplaceParticipationsResponse{}
	if err = json.Unmarshal(byteArray, res); err != nil {
		return nil, err
	}

	return res, nil
}

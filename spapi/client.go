package spapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

const (
	ContentType                             = "application/json"
	ServiceName                             = "execute-api"
	APIEndpointAuthToken                    = "https://api.amazon.com/auth/o2/token"
	APIEndpointGetCatalogItem               = "https://%s/catalog/v0/items/%s"
	APIEndpointListCatalogItem              = "https://%s/catalog/v0/items"
	APIEndpointGetMarketplaceParticipations = "https://%s/sellers/v1/marketplaceParticipations"
)

type Config struct {
	ClientID      string
	ClientSecret  string
	RefreshToken  string
	AccessKey     string
	SecretKey     string
	Endpoint      string // see https://developer-docs.amazon.com/amazon-shipping/docs/sp-api-endpoints
	MarketplaceId string // see https://developer-docs.amazon.com/sp-api/docs/marketplace-ids
	Region        string // see https://developer-docs.amazon.com/amazon-shipping/docs/sp-api-endpoints
}

type Client struct {
	Config            *Config
	Credentials       *credentials.Credentials
	AccessToken       string
	AccessTokenExpire time.Time
	Region            string
	ServiceName       string
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewClient(c *Config) (*Client, error) {
	tokenResp, err := GetAccessToken(c)
	if err != nil {
		return nil, err
	}

	client := &Client{Config: c}
	client.AccessToken = tokenResp.AccessToken
	client.AccessTokenExpire = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

	client.Credentials = credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, "")

	return client, nil
}

func GetAccessToken(c *Config) (*AccessTokenResponse, error) {
	req, _ := json.Marshal(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": c.RefreshToken,
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
	})
	resp, err := http.Post(APIEndpointAuthToken, ContentType, bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &AccessTokenResponse{}
	if err = json.Unmarshal(bodyByte, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Amz-Access-Token", c.AccessToken)

	signer := v4.NewSigner(c.Credentials)
	signer.Sign(req, nil, ServiceName, c.Config.Region, time.Now())

	client := &http.Client{}
	return client.Do(req)
}

package spapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

const (
	ServiceName               = "execute-api"
	APIEndpointAuthToken      = "https://api.amazon.com/auth/o2/token"
	APIEndpointGetCatalogItem = "https://%s/catalog/v0/items/%s"
)

type Config struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
	AccessKey    string
	SecretKey    string
	// see https://developer-docs.amazon.com/amazon-shipping/docs/sp-api-endpoints
	Endpoint string
	// see https://developer-docs.amazon.com/sp-api/docs/marketplace-ids
	MarketplaceId string
	// see https://developer-docs.amazon.com/amazon-shipping/docs/sp-api-endpoints
	Region string
}

type Client struct {
	Config            *Config
	Credentials       *credentials.Credentials
	AccessToken       string
	AccessTokenExpire time.Time
	Region            string
	ServiceName       string
}

type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewClient(c *Config) (*Client, error) {
	client := &Client{Config: c}
	if err := client.AuthToken(); err != nil {
		return nil, err
	}

	client.Credentials = credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, "")
	return client, nil
}

func (c *Client) AuthToken() error {
	req, _ := json.Marshal(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": c.Config.RefreshToken,
		"client_id":     c.Config.ClientID,
		"client_secret": c.Config.ClientSecret,
	})

	res, err := http.Post(APIEndpointAuthToken, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	authTokenResponse := &AuthTokenResponse{}
	if err = json.Unmarshal(bodyByte, authTokenResponse); err != nil {
		return err
	}

	c.AccessToken = authTokenResponse.AccessToken
	c.AccessTokenExpire = time.Now().Add(time.Duration(authTokenResponse.ExpiresIn) * time.Second)

	return nil
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	v := url.Values{}
	v.Add("MarketplaceId", c.Config.MarketplaceId)
	req.URL.RawQuery = v.Encode()

	req.Header.Add("X-Amz-Access-Token", c.AccessToken)

	signer := v4.NewSigner(c.Credentials)
	signer.Sign(req, nil, ServiceName, c.Config.Region, time.Now())

	client := &http.Client{}
	return client.Do(req)
}

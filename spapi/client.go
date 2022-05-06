package spapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	APIEndpointAuthToken = "https://api.amazon.com/auth/o2/token"
)

type Credential struct {
	AccessKey    string
	SecretKey    string
	ClientID     string
	ClientSecret string
	RefreshToken string
}

type Client struct {
	Credential        *Credential
	AccessToken       string
	AccessTokenExpire time.Time
	Region            string
	ServiceName       string
}

type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewClient(c *Credential) (*Client, error) {
	client := &Client{Credential: c}
	err := client.AuthToken()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) AuthToken() error {
	req, _ := json.Marshal(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": c.Credential.RefreshToken,
		"client_id":     c.Credential.ClientID,
		"client_secret": c.Credential.ClientSecret,
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

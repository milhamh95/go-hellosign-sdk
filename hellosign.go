package hellosign

import "net/http"

const (
	BaseURL = "https://api.hellosign.com/v3"
)

type Client struct {
	AccountAPI *AccountAPI
	common     service
	apiKey     string
	HTTPClient *http.Client
	BaseURL    string
}

type service struct {
	client *Client
}

func NewAPI(apiKey string, httpClient *http.Client) *Client {
	c := &Client{}
	c.common.client = c

	c.AccountAPI = (*AccountAPI)(&c.common)
	c.apiKey = apiKey
	c.HTTPClient = httpClient
	c.BaseURL = BaseURL
	return c
}

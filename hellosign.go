package hellosign

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
)

const (
	baseURL = "https://api.hellosign.com/v3"
)

// Client is api client for hellosign
type Client struct {
	common              service
	apiKey              string
	HTTPClient          *http.Client
	BaseURL             string
	AccountAPI          *AccountAPI
	SignatureRequestAPI *SignatureRequestAPI
}

type service struct {
	client *Client
}

// NewClient return new hellosign api client
func NewClient(apiKey string, httpClient *http.Client) *Client {
	c := &Client{}
	c.common.client = c

	c.apiKey = apiKey
	c.HTTPClient = httpClient
	c.BaseURL = baseURL
	c.AccountAPI = (*AccountAPI)(&c.common)
	c.SignatureRequestAPI = (*SignatureRequestAPI)(&c.common)
	return c
}

func (c *Client) doRequest(path string, method string, params *bytes.Buffer, w *multipart.Writer) (*http.Response, error) {
	req, err := http.NewRequest(method, path, params)
	if err != nil {
		return nil, err
	}

	if method != http.MethodGet && method != http.MethodDelete {
		req.Header.Set("Content-Type", w.FormDataContentType())
	}
	req.SetBasicAuth(c.apiKey, "")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		msg, err := prepareError(resp)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(msg)
	}

	return resp, err
}

package hellosign

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
)

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

func (c *Client) doRequest(path string, method string, params *bytes.Buffer, w *multipart.Writer) (*http.Response, error) {
	req, err := http.NewRequest(method, path, params)
	if err != nil {
		return nil, err
	}

	if method == http.MethodGet {
		req.Header.Set("Content-Type", w.FormDataContentType())
	}
	req.SetBasicAuth(c.apiKey, "")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		msg, err := prepareError(resp)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(msg)
	}

	return resp, err
}

package hellosign

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"time"
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

type requestParam struct {
	path   string
	method string
	body   io.Reader
	writer *multipart.Writer
}

// NewClient return new hellosign api client
func NewClient(apiKey string) *Client {
	c := &Client{}
	c.common.client = c

	c.apiKey = apiKey
	c.HTTPClient = &http.Client{
		Timeout: 5 * time.Second,
	}
	c.BaseURL = baseURL
	c.AccountAPI = (*AccountAPI)(&c.common)
	c.SignatureRequestAPI = (*SignatureRequestAPI)(&c.common)
	return c
}

func (c *Client) callAPI(ctx context.Context, r requestParam) (*http.Response, error) {
	req, err := c.prepareRequest(ctx, r)
	if err != nil {
		return nil, err
	}

	resp, err := c.executeRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) prepareRequest(ctx context.Context, r requestParam) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, r.method, r.path, r.body)
	if err != nil {
		return nil, err
	}

	if r.method != http.MethodGet && r.method != http.MethodDelete {
		req.Header.Set("Content-Type", r.writer.FormDataContentType())
	}

	return req, nil
}

func (c *Client) executeRequest(req *http.Request) (*http.Response, error) {
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

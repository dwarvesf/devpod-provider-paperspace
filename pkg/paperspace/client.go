package paperspace

import (
	"context"
	"net/http"
)

type RequestParams struct {
	Context context.Context   `json:"-,omitempty" url:"-,omitempty"`
	Headers map[string]string `json:"-,omitempty" url:"-,omitempty"`
}

type Client struct {
	APIKey  string
	Backend Backend
}

// client that makes requests to Gradient API
func NewClient(apiKey string) *Client {
	client := Client{
		Backend: NewAPIBackend(),
	}

	if apiKey != "" {
		client.APIKey = apiKey
	}

	return &client
}

func NewClientWithBackend(apiKey string, backend Backend) *Client {
	client := NewClient(apiKey)
	client.Backend = backend

	return client
}

func (c *Client) Request(method string, url string, params, result interface{}, requestParams RequestParams) (*http.Response, error) {
	if requestParams.Headers == nil {
		requestParams.Headers = make(map[string]string)
	}
	requestParams.Headers["x-api-key"] = c.APIKey

	return c.Backend.Request(method, url, params, result, requestParams)
}

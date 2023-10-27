package lotrsdk

import (
	"net/http"
)

const BaseURL = "https://the-one-api.dev/v2"

type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		HTTPClient: http.DefaultClient,
	}
}

func (c *Client) makeRequest(endpoint string) (*http.Request, error) {
	req, err := http.NewRequest("GET", BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.APIKey)
	return req, nil
}

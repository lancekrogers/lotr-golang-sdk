package lotrsdk

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetQuotes() (QuotesResponse, error) {
	return c.fetchQuotes("/quote")
}

func (c *Client) FilterQuotes(filter *Filter) (QuotesResponse, error) {
	endpoint := "/quote" + "?" + filter.Encode()
	return c.fetchQuotes(endpoint)
}

func (c *Client) GetQuoteByID(quoteID string) (Quote, error) {
	response, err := c.fetchQuotes(fmt.Sprintf("/quote/%s", quoteID))
	if err != nil || len(response.Docs) == 0 {
		return Quote{}, err
	}
	return response.Docs[0], err
}

func (c *Client) GetQuotesByMovieID(movieID string) (QuotesResponse, error) {
	return c.fetchQuotes(fmt.Sprintf("/movie/%s/quote", movieID))
}

func (c *Client) fetchQuotes(endpoint string) (QuotesResponse, error) {
	req, err := c.makeRequest(endpoint)
	if err != nil {
		return QuotesResponse{}, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return QuotesResponse{}, err
	}
	defer resp.Body.Close()

	var quotesResponse QuotesResponse
	err = json.NewDecoder(resp.Body).Decode(&quotesResponse)
	return quotesResponse, err
}

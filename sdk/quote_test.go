package lotrsdk

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetQuotes(t *testing.T) {
	client := NewClient("test_api_key")
	httpmock.ActivateNonDefault(client.HTTPClient)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{"docs": [{"_id": "1", "dialog": "Quote 1", "movie": "Movie 1 ID", "character": "Character 1", "id": "Q1"}, {"_id": "2", "dialog": "Quote 2", "movie": "Movie 2 ID", "character": "Character 2", "id": "Q2"}]}`
	httpmock.RegisterResponder("GET", BaseURL+"/quote",
		httpmock.NewStringResponder(200, mockResponse))

	quotes, err := client.GetQuotes()

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	quoteCount := len(quotes.Docs)
	if quoteCount != 2 {
		t.Errorf("Expected 2 quotes, got %d", quoteCount)
	}
}

func TestGetQuoteByID(t *testing.T) {
	client := NewClient("test_api_key")
	httpmock.ActivateNonDefault(client.HTTPClient)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{"docs": [{"_id": "2", "dialog": "Quote 2", "movie": "Movie 2 ID", "character": "Character 2", "id": "Q2"}]}`
	httpmock.RegisterResponder("GET", BaseURL+"/quote/2",
		httpmock.NewStringResponder(200, mockResponse))

	quote, err := client.GetQuoteByID("2")

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if quote.ID != "2" || quote.Dialog != "Quote 2" {
		t.Errorf("Unexpected quote data received: %#v", quote)
	}
}

func TestGetQuotesByMovieID(t *testing.T) {
	client := NewClient("test_api_key")
	httpmock.ActivateNonDefault(client.HTTPClient)
	defer httpmock.DeactivateAndReset()

	mockResponse := `{"docs": [{"_id": "1", "dialog": "Quote from Movie 1", "movie": "1", "character": "Character 1", "id": "Q1"}]}`
	httpmock.RegisterResponder("GET", BaseURL+"/movie/1/quote",
		httpmock.NewStringResponder(200, mockResponse))

	quotes, err := client.GetQuotesByMovieID("1")

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	quoteCount := len(quotes.Docs)
	if quoteCount != 1 {
		t.Errorf("Expected 1 quote, got %d", quoteCount)
	}
	firstQuote := quotes.Docs[0].Dialog
	if firstQuote != "Quote from Movie 1" {
		t.Errorf("Unexpected quote text received: %s", firstQuote)
	}
}

func TestFilterQuotes(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockResponse := `{
		"docs": [{"ID": "1", "Dialog": "One ring to rule them all"}]
	}`
	httpmock.RegisterResponder("GET", "https://the-one-api.dev/v2/quote?Dialog=One+ring+to+rule+them+all",
		httpmock.NewStringResponder(200, mockResponse))

	client := NewClient("YOUR_API_KEY")

	filter := &Filter{
		params: map[string]string{
			"Dialog=": "One ring to rule them all",
		},
	}

	resp, err := client.FilterQuotes(filter)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(resp.Docs) != 1 || resp.Docs[0].Dialog != "One ring to rule them all" {
		t.Fatalf("Unexpected quote data received")
	}
}

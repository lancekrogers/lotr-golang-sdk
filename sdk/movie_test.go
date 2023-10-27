package lotrsdk

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetMovies(t *testing.T) {
	client := NewClient("test_api_key")
	httpmock.ActivateNonDefault(client.HTTPClient)
	defer httpmock.DeactivateAndReset()

	// Mocking the response for GetMovies
	mockResponse := `{"docs": [{"_id": "1", "name": "Movie 1"}, {"_id": "2", "name": "Movie 2"}]}`
	httpmock.RegisterResponder("GET", BaseURL+"/movie",
		httpmock.NewStringResponder(200, mockResponse))

	movies, err := client.GetMovies()

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	moviesCount := len(movies.Docs)
	if moviesCount != 2 {
		t.Errorf("Expected 2 movies, got %d", moviesCount)
	}
}

func TestGetMovieByID(t *testing.T) {
	client := NewClient("test_api_key")
	httpmock.ActivateNonDefault(client.HTTPClient)
	defer httpmock.DeactivateAndReset()

	// Mocking the response for GetMovieByID
	mockResponse := `{"docs": [{"_id": "2", "name": "Movie 2"}]}`
	httpmock.RegisterResponder("GET", BaseURL+"/movie/2",
		httpmock.NewStringResponder(200, mockResponse))

	movie, err := client.GetMovieByID("2")

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if movie.ID != "2" || movie.Name != "Movie 2" {
		t.Errorf("Unexpected movie data received: %#v", movie)
	}
}

func TestFilterMovies(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Sample mock response for movies
	mockResponse := `{
		"docs": [{"ID": "1", "Name": "The Fellowship of the Ring"}]
	}`
	httpmock.RegisterResponder("GET", "https://the-one-api.dev/v2/movie?name=The+Fellowship+of+the+Ring",
		httpmock.NewStringResponder(200, mockResponse))

	client := NewClient("YOUR_API_KEY")

	filter := &Filter{
		params: map[string]string{
			"name=": "The Fellowship of the Ring",
		},
	}

	resp, err := client.FilterMovies(filter)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(resp.Docs) != 1 || resp.Docs[0].Name != "The Fellowship of the Ring" {
		t.Fatalf("Unexpected movie data received")
	}
}

package lotrsdk

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetMovies() (MoviesResponse, error) {
	return c.fetchMovies("/movie")
}

func (c *Client) FilterMovies(filter *Filter) (MoviesResponse, error) {
	endpoint := "/movie" + "?" + filter.Encode()
	return c.fetchMovies(endpoint)
}

func (c *Client) GetMovieByID(movieID string) (Movie, error) {
	response, err := c.fetchMovies(fmt.Sprintf("/movie/%s", movieID))
	if err != nil || len(response.Docs) == 0 {
		return Movie{}, err
	}
	return response.Docs[0], err
}

func (c *Client) fetchMovies(endpoint string) (MoviesResponse, error) {
	req, err := c.makeRequest(endpoint)
	if err != nil {
		return MoviesResponse{}, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return MoviesResponse{}, err
	}
	defer resp.Body.Close()

	var moviesResponse MoviesResponse
	err = json.NewDecoder(resp.Body).Decode(&moviesResponse)
	return moviesResponse, err
}

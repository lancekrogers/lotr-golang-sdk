# Lord of the Rings Golang SDK

The Lord of the Rings SDK provides an easy way to fetch data related to the Lord of the Rings universe using the official API found [here](https://the-one-api.dev/).  You will need to setup an API key to use this sdk.

## Installation


### Using `go get`:

To fetch the SDK using the `go get` command:

```bash
go get github.com/lotr-golang-sdk.git
```

### Manual Setup:

Modify the `go.mod` file in your project to replace the module path with the local path to the SDK on your system.

## Usage

The SDK is structured with clear functions to fetch movies and quotes:

### Initialize Client

```go
client := lotrsdk.NewClient("your_api_key")
```

### Get All Movies

```go
movies, err := client.GetMovies()
```

### Get Movie By ID

```go
movie, err := client.GetMovieByID("movie_id_here")
```

### Get All Quotes

```go
quotes, err := client.GetQuotes()
```

### Get Quote By ID

```go
quote, err := client.GetQuoteByID("quote_id_here")
```

### Get Quotes By Movie ID

```go
quotesByMovie, err := client.GetQuotesByMovieID("movie_id_here")
```
### Filter Movies and Quotes

To filter query create a filter and add params using the below methods:

- Add
- AddNotEqual
- AddGreaterThan
- AddLessThan
- AddMatchRegex
- AddNotMatchRegex
- AddFieldExist
- AddFieldNotExist
- AddGreaterThanOrEqualTo
- AddLessThanOrEqualTo


#### Filter Movies
```go
	filter := lotrsdk.NewFilter()
	filter.Add("param", "value")
	filter.AddGreaterThan("param", "value")

	quotes, err := client.FilterQuotes(filter)
```
 
#### Filter Quotes 
```go
	filter := lotrsdk.NewFilter()
	filter.Add("param", "value")
	filter.AddGreaterThan("param", "value")

	quotes, err := client.FilterQuotes(filter)
```


## Testing

To test this repo, first, clone this repository:

```bash
git clone https://github.com/lancekrogers/lotr-golang-sdk.git
cd lotr-golang-sdk
```


To run unit test run:
```bash
make test
```

## Examples

1. Before running the examples, make sure to add your API key to a `.env` file in the root directory:

```bash
echo "API_KEY=your_api_key_here" > .env
```

Replace `your_api_key_here` with your actual API key.


There are examples provided in the `examples` directory to help you understand the usage of the SDK. For instance, the `get_quotes_by_movie.go` fetches 10 random quotes from the movie "The Return of The King". These examples can be run by running `make example` and each example can be run indivdiually by runing `make` followed by one of the commands below:

- test-example-get-movie
- test-example-get-quote 
- test-example-get-movies 
- test-example-get-quotes 
- test-example-get-quotes-by-movie

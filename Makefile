.PHONY: all test build examples test-example-get-movie test-example-get-quote test-example-get-movies test-example-get-quotes test-example-get-quotes-by-movie test-filter-quotes


all: test examples 

build:
	go build ./...

test-example-get-movie:
	go run examples/get_movie/get_movie.go

test-example-get-quote:
	go run examples/get_quote/get_quote.go

test-example-get-movies:
	go run examples/get_movies/get_movies.go

test-example-get-quotes:
	go run examples/get_quotes/get_quotes.go

test-example-get-quotes-by-movie:
	go run examples/get_quotes_by_movie/get_quotes_by_movie.go

test-filter-quotes:
	go run examples/filter_quotes/filter_quotes.go

test-filter-movies:
	go run examples/filter_movies/filter_movies.go


test:
	go test ./...

examples:
	@echo -e "\n Starting Test Examples \n"
	$(MAKE) test-example-get-movie
	@echo -e "\n"
	$(MAKE) test-example-get-quote
	@echo -e "\n"
	$(MAKE) test-example-get-movies
	@echo -e "\n"
	$(MAKE) test-example-get-quotes
	@echo -e "\n"
	$(MAKE) test-example-get-quotes-by-movie
	@echo -e "\n"
	$(MAKE) test-filter-movies
	@echo -e "\n"
	$(MAKE) test-filter-quotes

package main

import (
	"fmt"
	"log"
	"lotrsdk/examples/utils"
	lotrsdk "lotrsdk/sdk"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file, %s", err)
	}

	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is not set")
	}

	client := lotrsdk.NewClient(apiKey)

	docs, err := client.GetQuotesByMovieID("5cd95395de30eff6ebccde5d")

	if err != nil {
		log.Fatalf("Error fetching quotes: %s", err)
	}

	quotes, err := utils.GetRandomDocs(docs, 2)
	if err != nil {
		log.Fatalf("Error fetching quotes: %s", err)
	}

	fmt.Printf("\nFetched random quotes from \"The Return of The King\": %v \n", utils.ToJSON(quotes))
}

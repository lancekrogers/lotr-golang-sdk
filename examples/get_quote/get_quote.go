package main

import (
	"fmt"
	"log"
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

	quote, err := client.GetQuoteByID("5cd96e05de30eff6ebcce7ec")
	if err != nil {
		log.Fatalf("Error fetching quote: %s", err)
	}
	fmt.Printf("\nFetched quote: %s (ID: %s)\n", quote.Dialog, quote.ID)
}

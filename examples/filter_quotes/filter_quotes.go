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

	filter := lotrsdk.NewFilter()
	filter.Add("character", "5cd99d4bde30eff6ebccfe19")

	docs, err := client.FilterQuotes(filter)
	if err != nil {
		fmt.Println("Error fetching filtered quotes:", err)
		return
	}

	quotes, err := utils.GetRandomDocs(docs, 2)
	if err != nil {
		log.Fatalf("Error fetching filtered quotes: %s", err)
	}

	fmt.Printf("\nFetched filtered quotes: %v \n", utils.ToJSON(quotes))
}

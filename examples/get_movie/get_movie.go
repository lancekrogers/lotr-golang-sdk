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

	movie, err := client.GetMovieByID("5cd95395de30eff6ebccde5d")
	if err != nil {
		log.Fatalf("Error fetching movie: %s", err)
	}
	fmt.Printf("\nFetched movie: %s (ID: %s)\n", movie.Name, movie.ID)
}

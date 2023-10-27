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

	movies, err := client.GetMovies()
	if err != nil {
		log.Fatalf("Error fetching movies: %s", err)
	}
	fmt.Printf("\nFetched movies: %v \n", utils.ToJSON(movies))
}

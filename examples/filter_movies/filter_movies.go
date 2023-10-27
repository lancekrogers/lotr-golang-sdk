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
	filter.AddLessThan("budgetInMillions", "100")

	movies, err := client.FilterMovies(filter)
	if err != nil {
		fmt.Println("Error fetching filtered movies:", err)
		return
	}
	fmt.Printf("\nFetched filter movie: %v \n", utils.ToJSON(movies.Docs))
}

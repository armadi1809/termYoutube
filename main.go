package main

import (
	"log"
	"os"
	"strings"

	youtubeapi "github.com/armadi1809/termYoutube/youtubeApi"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid command line arguments, provide a single search query")
	}
	searchQuery := strings.ReplaceAll(os.Args[1], " ", "+")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load environment variables from file")
	}

	apiClient := youtubeapi.New()
	apiClient.Search(searchQuery)
}

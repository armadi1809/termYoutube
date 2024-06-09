package main

import (
	"log"

	"github.com/armadi1809/termYoutube/ui"
	youtubeapi "github.com/armadi1809/termYoutube/youtubeApi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load environment variables from file")
	}

	apiClient := youtubeapi.New()
	ui.Start(&apiClient)
}

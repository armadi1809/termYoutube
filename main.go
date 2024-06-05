package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type videoItem struct {
	Snippet struct {
		Title string `json:"title"`
	} `json:"snippet"`
}

type videosSearchRes struct {
	Items []videoItem `json:"items"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid command line arguments, provide a single search query")
	}
	searchQuery := os.Args[1]
	searchQuery = strings.ReplaceAll(searchQuery, " ", "+")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("unable to load environment variables from file")
	}
	youtubeApiKey := os.Getenv("YOUTUBE_API_KEY")
	res, err := http.DefaultClient.Get(fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/search?part=snippet&maxResults=25&q=%s&key=%s", searchQuery, youtubeApiKey))
	if err != nil {
		log.Fatalf("unable to reach the youtube api: %s", err.Error())
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("unbale to read response from the youtube api")
	}

	searchResult := &videosSearchRes{}
	if err = json.Unmarshal(body, searchResult); err != nil {
		log.Fatalf("unable to parse json response from the youtube api: %v", err)
	}

	fmt.Println(searchResult)

}

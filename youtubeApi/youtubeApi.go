package youtubeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const baseapi string = "https://youtube.googleapis.com/youtube/v3/search?part=snippet&maxResults=25&q=%s&key=%s"

type YoutubeApiClient struct {
	authKey string
	client  *http.Client
}

type videoItem struct {
	Id struct {
		VideoId string `json:"videoId"`
	} `json:"id"`
	Snippet struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"snippet"`
}

type VideosSearchRes struct {
	Items []videoItem `json:"items"`
}

func New() YoutubeApiClient {
	key := os.Getenv("YOUTUBE_API_KEY")
	return YoutubeApiClient{authKey: key, client: &http.Client{}}
}

func (c *YoutubeApiClient) Search(searchQuery string) *VideosSearchRes {
	res, err := c.client.Get(fmt.Sprintf(baseapi, searchQuery, c.authKey))
	if err != nil {
		log.Fatalf("unable to reach the youtube api: %s", err.Error())
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("unbale to read response from the youtube api")
	}

	searchResult := &VideosSearchRes{}
	if err = json.Unmarshal(body, searchResult); err != nil {
		log.Fatalf("unable to parse json response from the youtube api: %v", err)
	}

	return searchResult

}

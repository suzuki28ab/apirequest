package api_request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const API_YOUTUBE_URL = "https://www.googleapis.com/youtube/v3/search?part=snippet&channelId="

const API_QUERY = "&type=video&eventType=live&key="

type apiYoutube struct {
	Items []item `json:"items"`
}

type item struct {
	ID      id      `json:"id"`
	Snippet snippet `json:"snippet"`
}

type id struct {
	VideoID string `json:"videoId"`
}

type snippet struct {
	Title string `json:"title"`
}

func GetYoutubeLiveData(id string) (isLive bool, title string, videoID string) {
	resp := getYoutubeResponse(id)
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	var api apiYoutube
	json.Unmarshal(b, &api)
	isLive = false
	if len(api.Items) != 0 {
		isLive = true
	}

	if isLive {
		title = api.Items[0].Snippet.Title
		videoID = api.Items[0].ID.VideoID
		return
	}
	title = ""
	videoID = ""
	return
}

func getYoutubeResponse(id string) *http.Response {
	resp, err := http.Get(API_YOUTUBE_URL + id + API_QUERY + os.Getenv("TUBE_KEY"))
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

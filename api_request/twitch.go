package api_request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const API_TWITCH_URL = "https://api.twitch.tv/helix/streams?user_login="

type apiTwitch struct {
	DataSlice []data `json:"data"`
}

type data struct {
	Title string `json:"title"`
}

func GetTwitchLiveData(id string) (isLive bool, title string) {
	req, err := http.NewRequest("GET", API_TWITCH_URL+id, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Client-Id", os.Getenv("TWITCH_KEY"))
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	var apiTwitch apiTwitch
	json.Unmarshal(b, &apiTwitch)
	isLive = false
	if len(apiTwitch.DataSlice) != 0 {
		isLive = true
	}

	if isLive {
		title = apiTwitch.DataSlice[0].Title
		return
	}
	title = ""
	return
}

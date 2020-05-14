package api_request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const OAUTH_TWITCH_URL = "https://id.twitch.tv/oauth2/token"
const API_TWITCH_URL = "https://api.twitch.tv/helix/streams?user_login="

type oauthTwitch struct {
	AccessToken string `json:"access_token"`
}

type apiTwitch struct {
	DataSlice []data `json:"data"`
}

type data struct {
	Title string `json:"title"`
}

func GetTwitchToken() string {
	url := OAUTH_TWITCH_URL + "?client_id=" + os.Getenv("TWITCH_KEY") + "&client_secret=" + os.Getenv("TWITCH_SECRET") + "&grant_type=client_credentials"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	var oauthTwitch oauthTwitch
	json.Unmarshal(b, &oauthTwitch)

	return oauthTwitch.AccessToken
}

func GetTwitchLiveData(id string, token string) (isLive bool, title string) {
	resp := getTwitchResponse(id, token)
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

func getTwitchResponse(id string, token string) *http.Response {
	req, err := http.NewRequest("GET", API_TWITCH_URL+id, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Client-ID", os.Getenv("TWITCH_KEY"))
	req.Header.Add("Authorization", "Bearer "+token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

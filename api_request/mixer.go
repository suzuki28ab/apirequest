package api_request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_MIXER_URL = "https://mixer.com/api/v1/channels/"

type apiMixer struct {
	Online bool   `json:"online"`
	Title  string `json:"name"`
}

func GetMixerApi(id string) (isLive bool, title string) {
	resp := getMixerResponse(id)
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	var apiMixer apiMixer
	json.Unmarshal(b, &apiMixer)
	isLive = apiMixer.Online
	if isLive {
		title = apiMixer.Title
		return
	}
	title = ""
	return
}

func getMixerResponse(id string) *http.Response {
	resp, err := http.Get(API_MIXER_URL + id)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

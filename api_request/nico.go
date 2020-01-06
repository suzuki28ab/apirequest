package api_request

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const LOGIN_NICO_URL = "https://secure.nicovideo.jp/secure/login?site=niconico"
const API_NICO_URL = "http://live.nicovideo.jp/api/getplayerstatus?v="

type apiNico struct {
	Status string `xml:"status,attr"`
	Stream stream `xml:"stream"`
}

type stream struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

func GetNicoLiveData(id string, userSession string) (isLive bool, title string, videoID string) {
	req, err := http.NewRequest("GET", API_NICO_URL+id, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Cookie", "user_session="+userSession)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	xmlData, _ := ioutil.ReadAll(resp.Body)

	var apiNico apiNico
	xml.Unmarshal(xmlData, &apiNico)
	isLive = false
	if apiNico.Status == "ok" {
		isLive = true
	}

	if isLive {
		videoID = apiNico.Stream.ID
		title = apiNico.Stream.Title
		return
	}
	title = ""
	videoID = ""
	return
}

func trim(start string, end string, str string) string {
	startIndex := strings.Index(str, start) + 1
	endIndex := strings.Index(str, end)
	return str[startIndex:endIndex]
}

func GetUserSeesion() string {
	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	values := url.Values{}
	values.Add("mail", os.Getenv("NICO_MAIL"))
	values.Add("password", os.Getenv("NICO_PASS"))

	resp, err := client.PostForm(LOGIN_NICO_URL, values)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	header := resp.Header["Set-Cookie"]
	var userSession string
	for _, str := range header {
		if strings.Contains(str, "user_session=") && strings.Contains(str, "deleted") != true {
			userSession = trim("=", ";", str)
			break
		}
	}
	return userSession
}

package api_request

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const YOUTUBE_CHANNEL_URL = "https://www.youtube.com/channel/"

const LIVE = "/live"

func GetYoutubeLiveData(id string) (isLive bool, title string) {
	liveUrl := YOUTUBE_CHANNEL_URL + id + LIVE

	doc, err := goquery.NewDocument(liveUrl)
	if err != nil {
		fmt.Println("scraiping error")
	}

	isLive = false
	title = ""

	se := doc.Find("body")
	html, _ := se.Html()
	if strings.Contains(html, "ライブ配信開始") {
		isLive = true
		title = doc.Find("title").Text()
		title = strings.Trim(title, " - YouTube")
	}

	return
}

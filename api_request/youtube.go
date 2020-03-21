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

	se := doc.Find("#player")
	html, err := se.Html()
	if err != nil {
		fmt.Println("HTML error")
	}

	isLive = true
	if strings.Contains(html, "LIVE_STREAM_OFFLINE") {
		isLive = false
	}

	title = ""
	if isLive {
		//動的なDOMのためerrorでる可能性アリ
		startIndex := strings.Index(html, "\"title") + 11
		endIndex := strings.Index(html, "lengthSeconds") - 5
		if len(html) > 10 {
			title = html[startIndex:endIndex]
		}
	}
	return
}

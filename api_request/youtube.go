package api_request

import (
	"fmt"

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

	// body の linkタグにある type="text/xml+oembed" title="hoge" という構造を信用している
	// オフラインのチャンネルにはこの構造が無い想定
	se := doc.Find("body > link")
	se.Each(func(i int, s *goquery.Selection) {
		attr, _ := s.Attr("type")
		if attr == "text/xml+oembed" {
			title, _ = s.Attr("title")
			isLive = true
		}
	})

	return
}

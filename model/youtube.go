package model

import (
	"github.com/jinzhu/gorm"
	"github.com/suzukix/apireq/api_request"
)

const LIVE_YOUTUBE_URL = "https://www.youtube.com/watch?v="

type Youtube struct {
	ID        int
	Account   string
	Title     string
	OnURL     string
	OffURL    string
	BcasterID int
}

func (y Youtube) UpdateYoutubeStatus(db *gorm.DB) (isLive bool) {
	isLive, title, onURL := getYoutubeInfo(y)
	if y.Title != title {
		db.Model(&y).Updates(map[string]interface{}{"title": title, "on_url": onURL})
	}
	return isLive
}

func getYoutubeInfo(youtube Youtube) (isLive bool, title string, onURL string) {
	isLive, title, videoID := api_request.GetYoutubeLiveData(youtube.Account)
	onURL = ""
	if isLive {
		onURL = LIVE_YOUTUBE_URL + videoID
		title = "(Youtube)" + title
	}
	return
}

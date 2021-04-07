package model

import (
	"time"

	"github.com/suzukix/apireq/api_request"
	"gorm.io/gorm"
)

const LIVE_YOUTUBE_URL = "https://www.youtube.com/watch?v="

type Youtube struct {
	ID        int
	Account   string
	Title     string
	OnURL     string
	OffURL    string
	BcasterID int
	Bcaster   *Bcaster
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (y Youtube) UpdateYoutubeStatus(db *gorm.DB) (isLive bool) {
	isLive, title, onURL := getYoutubeInfo(y)
	if y.Title != title {
		db.Model(&y).Updates(map[string]interface{}{"title": title, "on_url": onURL})
	}
	return isLive
}

func getYoutubeInfo(youtube Youtube) (isLive bool, title string, onURL string) {
	isLive, title = api_request.GetYoutubeLiveData(youtube.Account)
	onURL = ""
	if isLive {
		onURL = api_request.YOUTUBE_CHANNEL_URL + youtube.Account + api_request.LIVE
	}
	return
}

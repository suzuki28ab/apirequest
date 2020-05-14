package model

import (
	"github.com/jinzhu/gorm"
	"github.com/suzukix/apireq/api_request"
)

const TWITCH_URL = "https://www.twitch.tv/"

type Twitch struct {
	ID        int
	Account   string
	Title     string
	OnURL     string
	OffURL    string
	BcasterID int
}

func (t Twitch) UpdateTwitchStatus(db *gorm.DB, token string) (isLive bool) {
	isLive, title, onURL := getTwitchInfo(t, token)
	if t.Title != title {
		db.Model(&t).Updates(map[string]interface{}{"title": title, "on_url": onURL})
	}
	return
}

func getTwitchInfo(twitch Twitch, token string) (isLive bool, title string, onURL string) {
	isLive, title = api_request.GetTwitchLiveData(twitch.Account, token)
	onURL = ""
	if isLive {
		onURL = TWITCH_URL + twitch.Account
		title = "(Twitch)" + title
	}
	return
}

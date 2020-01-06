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
	OnUrl     string
	OffUrl    string
	BcasterID int
}

func (t Twitch) UpdateTwitchStatus(db *gorm.DB) (isLive bool) {
	db.LogMode(true)
	isLive, title, onUrl := getTwitchInfo(t)
	if t.Title != title {
		db.Model(&t).Updates(map[string]interface{}{"title": title, "on_url": onUrl})
	}
	return
}

func getTwitchInfo(twitch Twitch) (isLive bool, title string, onUrl string) {
	isLive, title = api_request.GetTwitchLiveData(twitch.Account)
	onUrl = ""
	if isLive {
		onUrl = TWITCH_URL + twitch.Account
		title = "(Twitch)" + title
	}
	return
}

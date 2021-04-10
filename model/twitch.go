package model

import (
	"time"

	"github.com/suzukix/apireq/api_request"
	"gorm.io/gorm"
)

const TWITCH_URL = "https://www.twitch.tv/"

type Twitch struct {
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

func (t Twitch) UpdateTwitchStatus(db *gorm.DB, token string) (isLive bool) {
	isLive, title, onURL := getTwitchInfo(t, token)
	if t.Title != title {
		db.Model(&t).Updates(map[string]interface{}{"title": title, "on_url": onURL})
		bcaster := Bcaster{}
		db.Where("id = ?", t.BcasterID).First(&bcaster)
		db.Model(&bcaster).Updates(Bcaster{StreamTitle: title, StreamUrl: onURL})
	}

	return
}

func getTwitchInfo(twitch Twitch, token string) (isLive bool, title string, onURL string) {
	isLive, title = api_request.GetTwitchLiveData(twitch.Account, token)
	onURL = ""
	if isLive {
		onURL = TWITCH_URL + twitch.Account
	}
	return
}

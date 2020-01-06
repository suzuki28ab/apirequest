package model

import (
	"github.com/jinzhu/gorm"
	"github.com/suzukix/apireq/api_request"
)

const MIXER_URL = "https://mixer.com/"

type Mixer struct {
	ID        int
	Account   string
	Title     string
	OnUrl     string
	OffUrl    string
	BcasterID int
}

func (m Mixer) UpdateMixerStatus(db *gorm.DB) (isLive bool) {
	db.LogMode(true)
	isLive, title, onUrl := getMixerInfo(m)
	if m.Title != title {
		db.Model(&m).Updates(map[string]interface{}{"title": title, "on_url": onUrl})
	}
	return
}

func getMixerInfo(mixer Mixer) (isLive bool, title string, onUrl string) {
	isLive, title = api_request.GetMixerApi(mixer.Account)
	onUrl = ""
	if isLive {
		onUrl = MIXER_URL + mixer.Account
		title = "(Mixer)" + title
	}
	return
}

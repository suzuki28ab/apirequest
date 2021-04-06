package model

import (
	"github.com/suzukix/apireq/api_request"
	"gorm.io/gorm"
)

const MIXER_URL = "https://mixer.com/"

type Mixer struct {
	ID        int
	Account   string
	Title     string
	OnURL     string
	OffURL    string
	BcasterID int
}

func (m Mixer) UpdateMixerStatus(db *gorm.DB) (isLive bool) {
	isLive, title, onURL := getMixerInfo(m)
	if m.Title != title {
		db.Model(&m).Updates(map[string]interface{}{"title": title, "on_url": onURL})
	}
	return
}

func getMixerInfo(mixer Mixer) (isLive bool, title string, onURL string) {
	isLive, title = api_request.GetMixerApi(mixer.Account)
	onURL = ""
	if isLive {
		onURL = MIXER_URL + mixer.Account
		title = "(Mixer)" + title
	}
	return
}

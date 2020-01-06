package model

import (
	"github.com/jinzhu/gorm"
	"github.com/suzukix/apireq/api_request"
)

const LIVE_NICO_URL = "http://live2.nicovideo.jp/watch/"

type Nico struct {
	ID        int
	Account   string
	Title     string
	OnURL     string
	OffURL    string
	BcasterID int
}

func (n Nico) UpdateNicoStatus(db *gorm.DB, userSession string) (isLive bool) {
	isLive, title, onURL := getNicoInfo(n, userSession)
	if n.Title != title {
		db.Model(&n).Updates(map[string]interface{}{"title": title, "on_url": onURL})
	}
	return
}

func getNicoInfo(nico Nico, userSession string) (isLive bool, title string, onURL string) {
	isLive, title, videoID := api_request.GetNicoLiveData(nico.Account, userSession)
	onURL = ""
	if isLive {
		onURL = LIVE_NICO_URL + videoID
		title = "(Nico)" + title
	}
	return
}

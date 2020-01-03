package model

import (
	"github.com/jinzhu/gorm"
	"github.com/suzukix/apireq/api_request"
)

const LIVE_NICO_URL = "http://live2.nicovideo.jp/watch/"

type Nico struct {
	Id        int
	Account   string
	Title     string
	OnUrl     string
	OffUrl    string
	BcasterId int
}

func (n Nico) UpdateNicoStatus(db *gorm.DB, userSession string) (isLive bool) {
	db.LogMode(true)
	isLive, title, onUrl := getNicoInfo(n, userSession)
	if n.Title != title {
		db.Model(&n).Updates(map[string]interface{}{"title": title, "on_url": onUrl})
	}
	return
}

func getNicoInfo(nico Nico, userSession string) (isLive bool, title string, onUrl string) {
	isLive, title, videoId := api_request.GetNicoLiveData(nico.Account, userSession)
	onUrl = ""
	if isLive {
		onUrl = LIVE_NICO_URL + videoId
		title = "(Nico)" + title
	}
	return
}

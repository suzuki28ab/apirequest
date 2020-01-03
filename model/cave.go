package model

import (
	"github.com/jinzhu/gorm"
	"github.com/suzukix/apireq/api_request"
)

const CAVE_URL = "https://www.cavelis.net/live/"

type Cavetube struct {
	Id        int
	Account   string
	Title     string
	OnUrl     string
	OffUrl    string
	BcasterId int
}

func (c Cavetube) UpdateCavetubeStatus(db *gorm.DB, liveData []api_request.Entry) (isLive bool) {
	db.LogMode(true)
	isLive, title, onUrl := getCaveInfo(c, liveData)
	if c.Title != title {
		db.Model(&c).Updates(map[string]interface{}{"title": title, "on_url": onUrl})
	}
	return
}

func getCaveInfo(cave Cavetube, liveData []api_request.Entry) (isLive bool, title string, onUrl string) {
	isLive = false
	title = ""
	onUrl = ""
	for _, e := range liveData {
		if cave.Account == e.Author.Name {
			isLive = true
			onUrl = CAVE_URL + cave.Account
			title = e.Title
			break
		}
	}
	return
}

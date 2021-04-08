package model

import (
	"time"

	"gorm.io/gorm"
)

const OFFLINE = 0
const ONLINE = 1

type Bcaster struct {
	ID        int
	Name      string
	Status    int
	StartFlag int
	CreatedAt time.Time
	UpdatedAt time.Time
	Nico      Nico
	Twitch    Twitch
	Youtube   Youtube
}

func (b Bcaster) RequestBcasterLive(db *gorm.DB, nicoUserSession string, twitchToken string) (startFlag int) {
	onlineCheckSlice := []bool{}
	if b.Youtube.ID != 0 {
		isLive := b.Youtube.UpdateYoutubeStatus(db)
		onlineCheckSlice = append(onlineCheckSlice, isLive)
	}
	if b.Twitch.ID != 0 {
		isLive := b.Twitch.UpdateTwitchStatus(db, twitchToken)
		onlineCheckSlice = append(onlineCheckSlice, isLive)
	}
	if b.Nico.ID != 0 {
		isLive := b.Nico.UpdateNicoStatus(db, nicoUserSession)
		onlineCheckSlice = append(onlineCheckSlice, isLive)
	}
	isOnline := false
	for _, onlineCheck := range onlineCheckSlice {
		if onlineCheck {
			isOnline = true
			break
		}
	}
	startFlag = b.updateStatus(db, isOnline)
	return
}

func (b Bcaster) Title() string {
	if b.Twitch.Title != "" {
		return b.Twitch.Title
	} else if b.Youtube.Title != "" {
		return b.Youtube.Title
	} else {
		return b.Nico.Title
	}
}

func (b Bcaster) StreamUrl() string {
	if b.Twitch.OnURL != "" {
		return b.Twitch.OnURL
	} else if b.Youtube.OnURL != "" {
		return b.Youtube.OnURL
	} else {
		return b.Nico.OnURL
	}
}

func (b Bcaster) updateStatus(db *gorm.DB, isOnline bool) (startFlag int) {
	startFlag = 0
	status := OFFLINE

	if isOnline {
		status = ONLINE
		if b.Status == OFFLINE && time.Since(b.UpdatedAt).Hours() > 1 {
			startFlag = 1
		}
	}

	if b.Status != status || b.StartFlag != startFlag {
		loc, _ := time.LoadLocation("UTC")
		db.Model(&b).UpdateColumns(map[string]interface{}{
			"status":     status,
			"start_flag": startFlag,
			"updated_at": time.Now().In(loc),
		})
	}
	return
}

func (b Bcaster) CreateDiscordNotice() string {
	return b.Title() + "\n" + b.StreamUrl()
}

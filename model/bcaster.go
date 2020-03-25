package model

import (
	"time"

	"github.com/jinzhu/gorm"
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
	Mixer     Mixer
	Nico      Nico
	Twitch    Twitch
	Youtube   Youtube
}

func (b Bcaster) RequestBcasterLive(db *gorm.DB, nicoUserSession string) (startFlag int) {
	onlineCheckSlice := []bool{}
	if b.Youtube.ID != 0 {
		isLive := b.Youtube.UpdateYoutubeStatus(db)
		onlineCheckSlice = append(onlineCheckSlice, isLive)
	}
	if b.Mixer.ID != 0 {
		isLive := b.Mixer.UpdateMixerStatus(db)
		onlineCheckSlice = append(onlineCheckSlice, isLive)
	}
	if b.Twitch.ID != 0 {
		isLive := b.Twitch.UpdateTwitchStatus(db)
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
		db.Model(&b).UpdateColumns(map[string]interface{}{
			"status":     status,
			"start_flag": startFlag,
			"updated_at": time.Now(),
		})
	}
	return
}

func (b Bcaster) CreateDiscordNotice() string {
	return b.Name + "さんが配信開始しました。\n https://daregirudojo.herokuapp.com"
}

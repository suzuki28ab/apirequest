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
	online := false
	for _, onlineCheck := range onlineCheckSlice {
		if onlineCheck {
			online = true
			break
		}
	}
	startFlag = b.updateStatus(db, online)
	return
}

func (b Bcaster) updateStatus(db *gorm.DB, online bool) (startFlag int) {
	startFlag = 0
	status := 0
	if online {
		if b.Status == ONLINE {
			startFlag = 0
		} else {
			startFlag = 1
		}
		status = ONLINE
	} else {
		status = OFFLINE
		startFlag = 0
	}
	if b.Status != status || b.StartFlag != startFlag {
		db.Model(&b).UpdateColumns(map[string]interface{}{
			"status": status,
			"start_flag": startFlag,
			"updated_at": time.Now(),
		})
	}
	return
}

func (b Bcaster) CreateDiscordNotice() string {
	return b.Name + "さんが配信開始しました。\n https://daregirudojo.herokuapp.com"
}
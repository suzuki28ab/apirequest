package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const DATE_TIME_FORMAT = "2006-01-02 15:04:05"

type Alarm struct {
	Id      int
	Content string
}

func SetApiTime(db *gorm.DB) {
	var alarm Alarm
	db.Model(&alarm).Where("id = ?", 1).Update(Alarm{Content: time.Now().Format(DATE_TIME_FORMAT)})
}

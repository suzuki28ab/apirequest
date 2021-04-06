package model

import (
	"time"

	"gorm.io/gorm"
)

const DATE_TIME_FORMAT = "2006-01-02 15:04:05"

type Alarm struct {
	ID      int
	Content string
}

func SetApiTime(db *gorm.DB) {
	var alarm Alarm
	db.Model(&alarm).Where("id = ?", 1).Update("Content", time.Now().Format(DATE_TIME_FORMAT))
}

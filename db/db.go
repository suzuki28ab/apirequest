package db

import (
	"os"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDbConnect() (*gorm.DB, error) {
	DBMS := "postgres"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := 5432
	NAME := os.Getenv("DB_NAME")

	CONNECT := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		HOST,
		PORT,
		USER,
		NAME,
		PASS,
	)
	return gorm.Open(DBMS, CONNECT)
}

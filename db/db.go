package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnect() (*gorm.DB, error) {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := "5432"
	NAME := os.Getenv("DB_NAME")

	DATABASE_URL := "postgres://" + USER + ":" + PASS + "@" + HOST + ":" + PORT + "/" + NAME

	db, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		fmt.Println(err)
	}

	return gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
}

package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDbConnect() (*gorm.DB, error) {
	DBMS := "postgres"
	USER := "nhpjceeofmcuwk"
	PASS := "55f497de9c14cf998dc0e12e7a1c40dc3c70c3b5801b313134b0b1a6d823f75f"
	HOST := "ec2-23-21-96-70.compute-1.amazonaws.com"
	PORT := 5432
	NAME := "dbrquo8208923h"

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

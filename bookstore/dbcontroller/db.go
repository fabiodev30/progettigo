package dbcontroller

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	dsn := "host=localhost user=prova password=prova dbname=prova port=5432 sslmode=disable TimeZone=Europe/Rome"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

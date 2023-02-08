package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func init() {
	dsn := "host=127.0.0.1 user=nostr_relayer password=nostr_relayer dbname=nostr_relayer port=5432 sslmode=disable TimeZone=Asia/Singapore"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database = db
}

func GenerateTables() {
	database.AutoMigrate(&Event{})
}

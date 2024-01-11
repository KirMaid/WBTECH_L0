package models

import (
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=user password=password dbname=db port=5452 sslmode=disable TimeZone=Europe/Moscow",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Order{}, &Delivery{}, &Payment{}, &Item{})
	if err != nil {
		return
	}
	DB = database
}

package db

import (
	"restAPI/CRUD/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func database() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return
}

func MigrateSchema() {
	db.AutoMigrate(&models.Movie{})
}

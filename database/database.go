package database

import (
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	var err error
	conn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	conn.AutoMigrate(&models.Actress{})
	db = conn
}

func ReturnDatabase() *gorm.DB {
	return db
}

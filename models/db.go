package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dbName string) *gorm.DB {
	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}


	DB = database

	return DB
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&User{})
} 
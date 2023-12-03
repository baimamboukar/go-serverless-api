package database

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"

	"github.com/baimamboukar/go-serverless-api/src/models"
)

func GetDatabaseInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./kengan_ashura.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.KenganPlayer{})
	return db
}

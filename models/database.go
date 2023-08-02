package models

import (
	"gorm.io/gorm"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
)

func ConnectDB(test_db bool) *gorm.DB {
	var (
		database *gorm.DB
		err      error
	)
	if !test_db {
		database, err = gorm.Open(mysql.Open("discus:discussion-forum@tcp(127.0.0.1:3306)/discus_forum_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	} else {
		database, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		database.Exec("PRAGMA foreign_keys = ON")
	}
	if err != nil {
		panic(err)
	}
	return database
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Topic{}, &Answer{}, &Comment{}, &Invited{})
	if admin := db.Where("name = ?", "admin").First(&User{}); admin.Error != nil {
		db.Create(&User{
			Name:        "admin",
			Username:    "admin",
			Password:    "admin",
			Active:      true,
			DateofBirth: time.Now(),
		})
	}
}

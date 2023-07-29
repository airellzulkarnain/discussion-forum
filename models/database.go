package models

import (
	"gorm.io/gorm"

	"time"

	"gorm.io/driver/mysql"
)

func ConnectDB() (*gorm.DB, error) {
	database, err := gorm.Open(mysql.Open("discus:discussion-forum@tcp(127.0.0.1:3306)/discus_forum_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	return database, err
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

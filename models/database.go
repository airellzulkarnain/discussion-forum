package models

import (
	"gorm.io/gorm"

	"time"

	"gorm.io/driver/mysql"
)

func ConnectDB(test_db bool) *gorm.DB {
	var (
		database *gorm.DB
		err      error
	)
	if !test_db {
		database, err = gorm.Open(mysql.Open("discus:discussion-forum@tcp(127.0.0.1:3306)/discus_forum_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	} else {
		database, err = gorm.Open(mysql.Open("discus:discussion-forum@tcp(127.0.0.1:3306)/discus_forum_test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
		MigrateDB(database)
	}
	if err != nil {
		panic(err)
	}
	return database
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Topic{}, &Answer{}, &Comment{}, &Invited{})
	if err := db.Limit(1).Where("name = ?", "admin").Find(&User{}); err.RowsAffected == 0 {
		db.Create(&User{
			Name:        "admin",
			Username:    "admin",
			Password:    "admin",
			Active:      true,
			DateofBirth: time.Now(),
		})
	}
}

func ClearTestDB() {
	db, err := gorm.Open(mysql.Open("discus:discussion-forum@tcp(127.0.0.1:3306)/discus_forum_test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Exec("DROP TABLE IF EXISTS comments")
	db.Exec("DROP TABLE IF EXISTS answers")
	db.Exec("DROP TABLE IF EXISTS topics")
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("DROP TABLE IF EXISTS inviteds")
}

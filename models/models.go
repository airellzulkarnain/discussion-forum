package models

import (
	"time"

	"gorm.io/gorm"
)

type status string

const (
	PUBLIC    status = "PUBLIC"
	PROTECTED status = "PROTECTED"
	PRIVATE   status = "PRIVATE"
)

type User struct {
	gorm.Model
	Name        string
	Username    string
	Password    string
	DateofBirth time.Time
	Topics      []Topic
	Answers     []Answer
	Comments    []Comment
}

type Topic struct {
	gorm.Model
	Title     string
	Body      string
	UpVotes   uint
	DownVotes uint
	Status    status `gorm:"default:'PUBLIC';type:ENUM('PUBLIC', 'PROTECTED', 'PRIVATE');column:status"`
	UserID    uint
}

type Answer struct {
	gorm.Model
	Body      string
	UpVotes   uint
	DownVotes uint
	UserID    uint
	TopicID   uint
}

type Comment struct {
	gorm.Model
	Body     string
	UserID   uint
	TopicID  uint
	AnswerID uint `gorm:"default:null"`
}

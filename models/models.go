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
	Name        string    `gorm:"not null"`
	Username    string    `gorm:"not null"`
	Password    string    `gorm:"not null"`
	DateofBirth time.Time `gorm:"not null"`
	Topics      []*Topic
	Answers     []*Answer
	Comments    []*Comment
}

type Topic struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Body      string `gorm:"not null"`
	UpVotes   uint   `gorm:"not null;default: 0"`
	DownVotes uint   `gorm:"not null;default: 0"`
	Status    status `gorm:"default:'PUBLIC';type:ENUM('PUBLIC', 'PROTECTED', 'PRIVATE');column:status"`
	Comments  []*Comment
	Answers   []*Answer
	UserID    uint `gorm:"not null"`
}

type Answer struct {
	gorm.Model
	Body      string `gorm:"not null"`
	UpVotes   uint   `gorm:"not null;default: 0"`
	DownVotes uint   `gorm:"not null;default: 0"`
	Comments  []*Comment
	UserID    uint `gorm:"not null"`
	TopicID   uint
}

type Comment struct {
	gorm.Model
	Body     string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	TopicID  uint   `gorm:"not null"`
	AnswerID uint
}

type Invited struct {
	UserID   uint `gorm:"primaryKey;autoIncrement:false"`
	AnswerID uint `gorm:"primaryKey;autoIncrement:false"`
}

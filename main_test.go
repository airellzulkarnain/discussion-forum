package main

import (
	"os"
	"testing"

	"github.com/airellzulkarnain/discussion-forum/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var router *gin.Engine

func TestSignIn(t *testing.T)        {}
func TestSignUp(t *testing.T)        {}
func TestCreateTopic(t *testing.T)   {}
func TestRetrieveTopic(t *testing.T) {}
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	router = setupRouter()

	router.Use(func(c *gin.Context) {
		c.Set("db", models.ConnectDB(true))
		db := c.MustGet("db").(*gorm.DB)
		models.MigrateDB(db)
		c.Next()
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

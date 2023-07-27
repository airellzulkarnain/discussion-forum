package main

import (
	"os"

	"github.com/airellzulkarnain/discussion-forum/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := models.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		models.MigrateDB(db)
	} else {
		gin.SetMode(gin.DebugMode)
		r := gin.Default()

		r.Run(":8080")
	}
}

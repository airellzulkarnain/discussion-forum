package main

import (
	"os"

	"github.com/airellzulkarnain/discussion-forum/controllers"
	"github.com/airellzulkarnain/discussion-forum/middlewares"
	"github.com/airellzulkarnain/discussion-forum/models"
	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		db, err := models.ConnectDB()
		if err != nil {
			panic("Failed to connect to database")
		}
		models.MigrateDB(db)
	} else {
		gin.SetMode(gin.DebugMode)
		r := setupRouter()
		r.Run(":8080")
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	public := r.Group("/api/v1")
	admin := r.Group("/api/v1/admin")
	user := r.Group("/api/v1")

	admin.Use(middlewares.AdminAuth())
	user.Use(middlewares.UserAuth())

	{
		public.GET("/topics/:id", controllers.RetrieveTopic)
		public.GET("/topics", controllers.SearchTopics)
		public.POST("/signin", controllers.SignIn)
		public.POST("/signup", controllers.SignUp)
	}

	{
		admin.PUT("/users/:id/verify", controllers.VerifyUser)
	}

	return r
}

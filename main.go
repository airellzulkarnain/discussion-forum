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
		db := models.ConnectDB(false)
		models.MigrateDB(db)
	} else {
		gin.SetMode(gin.DebugMode)
		r := setupRouter(false)

		r.Run(":8080")
	}
}

func setupRouter(test bool) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", models.ConnectDB(test))
		c.Next()
	})

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
		user.POST("/topics", controllers.CreateTopic)
		user.PUT("/topics/:id", controllers.UpdateTopic)
		user.DELETE("/topics/:id", controllers.DeleteTopic)
		user.PUT("/topics/:id/up", controllers.UpVoteTopic)
		user.PUT("/topics/:id/down", controllers.DownVoteTopic)

		user.POST("/answers", controllers.CreateAnswer)
		user.PUT("/answers/:id", controllers.UpdateAnswer)
		user.DELETE("/answers/:id", controllers.DeleteAnswer)
		user.PUT("/answers/:id/up", controllers.UpVoteAnswer)
		user.PUT("/answers/:id/down", controllers.DownVoteAnswer)

		user.POST("/comments", controllers.CreateComment)
		user.PUT("/comments/:id", controllers.UpdateComment)
		user.DELETE("/comments/:id", controllers.DeleteComment)

		user.PUT("/users/:id", controllers.UpdateUser)
		user.DELETE("/users/:id", controllers.DeleteUser)
	}

	{
		admin.PUT("/users/:id/verify", controllers.VerifyUser)
	}

	return r
}

package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/airellzulkarnain/discussion-forum/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SignIn(c *gin.Context) {
	db, err := models.ConnectDB()

	if err != nil {
		panic("Failed to connect to database")
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	db.Where("username = ? and password = ?", user.Username, user.Password).First(&user)
	if user.ID == 0 || !user.Active {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	secret, err := os.ReadFile(filepath.Join(".", "jwtRS256.key.pub"))
	if err != nil {
		panic(err)
	}
	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
func SignUp(c *gin.Context)         {}
func CreateTopic(c *gin.Context)    {}
func RetrieveTopic(c *gin.Context)  {}
func UpdateTopic(c *gin.Context)    {}
func DeleteTopic(c *gin.Context)    {}
func CreateAnswer(c *gin.Context)   {}
func UpdateAnswer(c *gin.Context)   {}
func DeleteAnswer(c *gin.Context)   {}
func CreateComment(c *gin.Context)  {}
func UpdateComment(c *gin.Context)  {}
func DeleteComment(c *gin.Context)  {}
func UpdateUser(c *gin.Context)     {}
func DeleteUser(c *gin.Context)     {}
func GetTopics(c *gin.Context)      {}
func UpVoteTopic(c *gin.Context)    {}
func DownVoteTopic(c *gin.Context)  {}
func UpVoteAnswer(c *gin.Context)   {}
func DownVoteAnswer(c *gin.Context) {}
func SearchTopics(c *gin.Context)   {}
func InviteUsers(c *gin.Context)    {}
func VerifyUser(c *gin.Context)     {}

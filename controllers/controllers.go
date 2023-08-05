package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/airellzulkarnain/discussion-forum/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func SignIn(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	db.Where("username = ? and password = ?", user.Username, user.Password).First(&user)
	if user.ID == 0 || !user.Active {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	role := "users"
	if user.ID == 1 && user.Name == "admin" {
		role = "admin"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"exp":  time.Now().Add(time.Minute * 1).Unix(),
		"role": role,
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

func SignUp(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := db.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "" + result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Create your account, wait for admin to activate your account :)"})
}

func CreateTopic(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var topic models.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	topic.UserID = c.MustGet("userId").(uint)
	if result := db.Create(&topic); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.Header("Location", fmt.Sprintf("/topics/%d", topic.ID))
	c.JSON(http.StatusOK, topic)
}
func RetrieveTopic(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var topic models.Topic
	if result := db.First(&topic, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, topic)
}
func UpdateTopic(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var topic models.Topic
	if result := db.First(&topic, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if topic.UserID != c.MustGet("userId").(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&topic)
	c.JSON(http.StatusNoContent, nil)
}
func DeleteTopic(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var topic models.Topic
	if result := db.First(&topic, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	if topic.UserID != c.MustGet("userId").(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	db.Delete(&topic)
	c.JSON(http.StatusNoContent, nil)
}
func CreateAnswer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var answer models.Answer
	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	answer.UserID = c.MustGet("userId").(uint)
	if result := db.Create(&answer); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.Header("Location", fmt.Sprintf("/topics/%d", answer.ID))
	c.JSON(http.StatusOK, answer)
}
func UpdateAnswer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var answer models.Answer
	if result := db.First(&answer, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	if answer.UserID != c.MustGet("userId").(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&answer)
	c.JSON(http.StatusNoContent, nil)
}
func DeleteAnswer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var answer models.Answer
	if result := db.First(&answer, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	if answer.UserID != c.MustGet("userId").(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	db.Delete(&answer)
	c.JSON(http.StatusNoContent, nil)
}
func CreateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.UserID = c.MustGet("userId").(uint)
	if result := db.Create(&comment); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.Header("Location", fmt.Sprintf("/topics/%d", comment.TopicID))
	c.JSON(http.StatusOK, comment)
}
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
func InviteUser(c *gin.Context)     {}

func VerifyUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if result := db.First(&user, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	user.Active = true
	db.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully verified accounts :)"})
}

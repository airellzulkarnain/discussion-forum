package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/airellzulkarnain/discussion-forum/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	router *gin.Engine
	token  string
)

func TestSignIn(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/v1/signin", bytes.NewBuffer([]byte(`{"username": "admin", "password": "admin"}`)))
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	var temp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &temp)
	if err != nil {
		panic(err)
	}
	assert.NotEqual(t, "", temp["token"])
	token = temp["token"]
}
func TestSignUp(t *testing.T) {
	date_of_birth := time.Date(2005, 10, 27, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	json_body := []byte(fmt.Sprintf(`{
		"name": "admin", 
		"username": "admin1", 
		"password": "admin1", 
		"date_of_birth": "%v"
		}`, date_of_birth))
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(json_body))
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	var temp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &temp)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "Successfully Create your account, wait for admin to activate your account :)", temp["message"])
}

func TestVerifyUser(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/v1/admin/users/2/verify", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	var temp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &temp)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "Successfully verified accounts :)", temp["message"])
}

func TestCreateTopic(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST",
		"/api/v1/topics",
		bytes.NewBuffer([]byte(`{
		"title": "title", 
		"description": "description", 
		"status": "public"
	}`)))
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, "", w.Header().Get("Location"))

}
func TestRetrieveTopic(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/topics/1", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestUpdateTopic(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/v1/topics/1", bytes.NewBuffer([]byte(`{
		"title": "title", 
		"description": "description", 
		"status": ("private", "protected", "public")
	}`)))
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}
func TestDeleteTopic(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/v1/topics/1", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestCreateAnswer(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/v1/answers", bytes.NewBuffer([]byte(`{
		"topicId": 0, 
		"body": "text"
	}`)))
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, "", w.Header().Get("Location"))
}
func TestUpdateAnswer(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/v1/answers/1", bytes.NewBuffer([]byte(`{
		"body": "text"
	}`)))
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}
func TestDeleteAnswer(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/v1/answers/1", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestCreateComment(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/v1/comments", bytes.NewBuffer([]byte(`{
		"body": "body", 
		"topicId": 0, 
		"answerId": nil, 
		"userId": 0
	}`)))
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestUpdateComment(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/v1/comments/1", bytes.NewBuffer([]byte(`{
		"body": "body"
	}`)))
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}
func TestDeleteComment(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/v1/comments/1", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestUpdateUser(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/v1/users/1", bytes.NewBuffer([]byte(`{
		"name": "name", 
		"dateofbirth": "dateofbirth", 
		"username": "username", 
		"password": "password"
	}`)))
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}
func TestDeleteUser(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestUpVoteTopic(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/v1/topics/1/up", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
func TestDownVoteTopic(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/v1/topics/1/down", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestUpVoteAnswer(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/v1/answers/1/up", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
func TestDownVoteAnswer(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/v1/answers/1/down", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestSearchTopics(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/search", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestSearchTopicsWithSort(t *testing.T) {
	w := httptest.NewRecorder()
	// desc|vote|latest|oldest|asc
	r, _ := http.NewRequest("GET", "/api/v1/search?sort=asc", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestSearchTopicsWithKeyword(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/search?keyword=title", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestSearchTopicsWithKeywordAndSort(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/search?keyword=title&sort=desc", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestInviteUser(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/api/v1/invite", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	models.ClearTestDB()
	router = setupRouter(true)
	exitCode := m.Run()
	os.Exit(exitCode)
}

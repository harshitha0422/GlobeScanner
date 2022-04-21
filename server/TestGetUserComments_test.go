package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUserComments(t *testing.T) {

	token, err := CreateToken("testG@gmail.com", "Guide", "123456789")
	assert.NoError(t, err)
	saveErr := CreateAuth("testG@gmail.com", "Guide", token)
	assert.NoError(t, saveErr)

	//var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Comment{})
	assert.NoError(t, err)
	fmt.Println("1")
	cmnt := Comment{
		Title:     "not bad",
		Name:      "123456789",
		Email:     "testG@gmail.com",
		PackageId: "1",
		Review:    "just ok",
	}
	fmt.Println("2")
	DB.Create(&cmnt)
	fmt.Println("3")
	r := gin.Default()
	fmt.Println("4")
	r.GET("/mycomments", TokenAuthMiddleware(), getalluserComments)
	fmt.Println("5")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/mycomments", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

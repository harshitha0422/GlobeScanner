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

func TestDeleteTouristProfile(t *testing.T) {

	token, err := CreateToken("testt1@gmail.com", "Tourist", "123456789")
	assert.NoError(t, err)
	saveErr := CreateAuth("testt1@gmail.com", "Tourist", token)
	assert.NoError(t, saveErr)

	//var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&UserProfile{})
	assert.NoError(t, err)
	fmt.Println("1")
	user := UserProfile{
		Email: "testt1@gmail.com",
		Name:  "123456789",
		Role:  "Tourist",
		About: "avail",
		// Age:      32,
		//Mobile: "9987896777",
		// Location: "gainesville",
		// Fav1:     "tampa",
		// Fav2:     "orlando",
		// Fav3:     "miami",
	}
	fmt.Println("2")
	DB.Create(&user)
	fmt.Println("3")
	r := gin.Default()
	fmt.Println("4")
	r.DELETE("/deleteProfile", TokenAuthMiddleware(), DeleteUserProfile)
	fmt.Println("5")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/deleteProfile", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

func TestDeleteGuideProfile(t *testing.T) {

	token, err := CreateToken("testg1@gmail.com", "Guide", "123456789")
	assert.NoError(t, err)
	saveErr := CreateAuth("testg1@gmail.com", "Guide", token)
	assert.NoError(t, saveErr)

	//var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&GuideProfile{})
	assert.NoError(t, err)
	fmt.Println("1")
	user := GuideProfile{
		Email:    "testG@gmail.com",
		Name:     "123456789",
		Role:     "Tourist",
		About:    "avail",
		Age:      32,
		Address:  "34th st",
		Location: "gainesville",
		Vehicle:  "tesla",
	}
	fmt.Println("2")
	DB.Create(&user)
	fmt.Println("3")
	r := gin.Default()
	fmt.Println("4")
	r.DELETE("/deleteProfile", TokenAuthMiddleware(), DeleteUserProfile)
	fmt.Println("5")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/deleteProfile", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

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

func TestGetUserProfile(t *testing.T) {

	token, err := CreateToken("testG@gmail.com", "Tourist", "123456789")
	assert.NoError(t, err)
	saveErr := CreateAuth("testG@gmail.com", "Tourist", token)
	assert.NoError(t, saveErr)

	//var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&UserProfile{})
	assert.NoError(t, err)
	fmt.Println("1")
	user := UserProfile{
		Email: "testG@gmail.com",
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
	r.GET("/userprofile", TokenAuthMiddleware(), getUserProfile)
	fmt.Println("5")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/userprofile", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

func TestGetAllUserProfile(t *testing.T) {

	token, err := CreateToken("testG@gmail.com", "Tourist", "123456789")
	assert.NoError(t, err)
	saveErr := CreateAuth("testG@gmail.com", "Tourist", token)
	assert.NoError(t, saveErr)

	//var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&UserProfile{})
	assert.NoError(t, err)
	fmt.Println("1")
	user := UserProfile{
		Email: "testG@gmail.com",
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
	r.GET("/userprofiles", TokenAuthMiddleware(), getallTouristprofile)
	fmt.Println("5")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/userprofiles", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

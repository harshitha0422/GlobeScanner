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

// func TestValidSearchPlaces(t *testing.T) {

// 	token, err := CreateToken("vishesha@gmail.com", "Tourist", "vishesha@123")
// 	assert.NoError(t, err)
// 	saveErr := CreateAuth("vishesha@gmail.com", "Tourist", token)
// 	assert.NoError(t, saveErr)

// 	router := gin.Default()

// 	router.GET("/searchPlaces/:location", TokenAuthMiddleware(), searchPlaces)

// 	w := httptest.NewRecorder()

// 	req, _ := http.NewRequest("GET", "/searchPlaces/Gainesville", nil)

// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

// 	router.ServeHTTP(w, req)

// 	assert.NoError(t, err)

// 	assert.Equal(t, 200, w.Code)

// }

func TestValidSearchPackage(t *testing.T) {

	token, err := CreateToken("testG@gmail.com", "Tourist", "123456789")
	assert.NoError(t, err)
	saveErr := CreateAuth("testG@gmail.com", "Tourist", token)
	assert.NoError(t, saveErr)

	//var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Package{})
	assert.NoError(t, err)
	newPackage := Package{
		Email:        "vishesha@gmail.com",
		Duration:     "1 week",
		Location:     "Agra",
		Accomodation: "Taj",
		Itinerary:    "absc",
		Included:     "bnn",
		Price:        "$2000",
	}
	DB.Create(newPackage)

	router := gin.Default()

	router.GET("/searchPackage/:location", TokenAuthMiddleware(), searchPackage)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/searchPackage/Agra", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	router.ServeHTTP(w, req)

	//err = json.Unmarshal(w.Body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "test@email.com", response.Email)
	// assert.Equal(t, "Test User", response.Name)

}

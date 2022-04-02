package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//func TestInvalidSearchPlaces(t *testing.T) {
// 	var response map[string]string

// 	token, err := CreateToken("megan12@gmail.com", "Tourist", "Megan")
// 	assert.NoError(t, err)

// 	router := gin.Default()

// 	router.GET("/searchPlaces/:location", TokenAuthMiddleware(), searchPlaces)

// 	w := httptest.NewRecorder()

// 	req, _ := http.NewRequest("GET", "/searchPlaces/", nil)
// 	//mt.Sprintf("Bearer %+v", token)
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

// 	router.ServeHTTP(w, req)

// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)

// 	assert.Equal(t, 400, w.Code)
// 	assert.Equal(t, "Please enter a valid location name.", response["error"])
// 	// assert.Equal(t, "Test User", response.Name)

//
// }

func TestValidSearchPackage(t *testing.T) {

	token, err := CreateToken("test@gmail.com", "Tourist", "test")
	assert.NoError(t, err)

	var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Package{})
	//assert.NoError(t, err)
	newPackage := Package{
		Email:        "test@gmail.com",
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

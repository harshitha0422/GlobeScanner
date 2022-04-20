package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestValidAddPackage(t *testing.T) {

	token, err := CreateToken("megan12@gmail.com", "Guide", "Megan")
	assert.NoError(t, err)
	saveErr := CreateAuth("megan12@gmail.com", "Guide", token)
	assert.NoError(t, saveErr)
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Package{})
	assert.NoError(t, err)

	router := gin.Default()

	router.POST("/addPackages", TokenAuthMiddleware(), addPackages)

	w := httptest.NewRecorder()

	var newPackage = []byte(`{
		"duration": "1 weeks", 
		"location": "Goa", 
		"accomodation": "Hyatt Regency", 
		"itinerary": "Day1:Red Fort, Day2: Fatehpur Sikhri, Day3: Agra", 
		"included": "Breakfast,Dinner, Sight-seeing",
		"price": "Rs. 10000"
	 }`)

	req, _ := http.NewRequest("POST", "/addPackages", bytes.NewBuffer(newPackage))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//fmt.Sprintf("Bearer %+v", token)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	router.ServeHTTP(w, req)

	//err = json.Unmarshal(w.Body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "test@email.com", response.Email)
	// assert.Equal(t, "Test User", response.Name)

}

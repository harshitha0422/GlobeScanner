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

	token, err := CreateToken("testG@gmail.com", "Guide", "123456789")
	assert.NoError(t, err)
	saveErr := CreateAuth("testG@gmail.com", "Guide", token)
	assert.NoError(t, saveErr)

	//var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Package{})
	assert.NoError(t, err)
	fmt.Println("1")
	pkg := Package{
		Email:        "testG@gmail.com",
		Duration:     "1 weeks",
		Location:     "Goa",
		Accomodation: "Hyatt Regency",
		Itinerary:    "Day1:Red Fort, Day2: Fatehpur Sikhri, Day3: Agra",
		Included:     "Breakfast,Dinner, Sight-seeing",
		Price:        "Rs. 10000",
	}
	fmt.Println("2")
	DB.Create(&pkg)
	r := gin.Default()
	r.POST("/addPackages", TokenAuthMiddleware(), addPackages)

	w := httptest.NewRecorder()

	var newPackage = []byte(`{
		"duration": "1 weeks",
		"location": "Goa",
		"accomodation": "Hyatt Regency",
		"itinerary": "Day1:Red Fort, Day2: Fatehpur Sikhri, Day3: Agra",
		"included": "Breakfast,Dinner, Sight-seeing",
		"price": "Rs. 10000"
	 }`)

	req, _ := http.NewRequest("POST", "/addPackages", bytes.NewReader(newPackage))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//fmt.Sprintf("Bearer %+v", token)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	r.ServeHTTP(w, req)

	//err = json.Unmarshal(w.Body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "test@email.com", response.Email)
	// assert.Equal(t, "Test User", response.Name)

}

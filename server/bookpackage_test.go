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

func TestInvalidBookPackage(t *testing.T) {

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
	r.POST("/bookPackages", TokenAuthMiddleware(), bookPackages)
	w := httptest.NewRecorder()

	var bkng = []byte(`{
		"packageId": "1"
	 }`)

	req, _ := http.NewRequest("POST", "/bookPackages", bytes.NewReader(bkng))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//fmt.Sprintf("Bearer %+v", token)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	r.ServeHTTP(w, req)

	//err = json.Unmarshal(w.Body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 401, w.Code)
	// assert.Equal(t, "test@email.com", response.Email)
	// assert.Equal(t, "Test User", response.Name)

}

func TestBookPackage(t *testing.T) {

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
	r.POST("/bookPackages", TokenAuthMiddleware(), bookPackages)
	w := httptest.NewRecorder()

	var bkng = []byte(`{
		"email":"testG@gmail.com",
		"packageId": "1"
	 }`)

	req, _ := http.NewRequest("POST", "/bookPackages", bytes.NewReader(bkng))
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

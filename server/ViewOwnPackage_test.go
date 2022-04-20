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

func TestViewOwnPackage(t *testing.T) {

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
	fmt.Println("3")
	r := gin.Default()
	fmt.Println("4")
	r.GET("/viewOwnPackage", TokenAuthMiddleware(), getPackage)
	fmt.Println("5")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/viewOwnPackage", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

func TestInvalidViewOwnPackage(t *testing.T) {

	token, err := CreateToken("testG@gmail.com", "Tourist", "123456789")
	assert.NoError(t, err)
	saveErr := CreateAuth("testG@gmail.com", "Tourist", token)
	assert.NoError(t, saveErr)

	//var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Package{})

	fmt.Println("1")
	r := gin.Default()
	fmt.Println("2")
	r.GET("/viewOwnPackage", TokenAuthMiddleware(), getPackage)
	fmt.Println("3")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/viewOwnPackage", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 500, w.Code)

}

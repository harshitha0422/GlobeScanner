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

func TestGetAllPackage(t *testing.T) {

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
		Email:        "testg@gmail.com",
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
	r.GET("/getAllPackage", TokenAuthMiddleware(), getallPackages)
	fmt.Println("5")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/getAllPackage", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

func TestViewBooking(t *testing.T) {

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
		Email:        "testg@gmail.com",
		Duration:     "1 weeks",
		Location:     "Goa",
		Accomodation: "Hyatt Regency",
		Itinerary:    "Day1:Red Fort, Day2: Fatehpur Sikhri, Day3: Agra",
		Included:     "Breakfast,Dinner, Sight-seeing",
		Price:        "Rs. 10000",
	}
	bkng := Booking{
		Email:     "testG@gmail.com",
		PackageId: "1",
	}
	fmt.Println("2")
	DB.Create(&pkg)
	DB.Create(&bkng)
	fmt.Println("3")
	r := gin.Default()
	fmt.Println("4")
	r.GET("/viewBooking", TokenAuthMiddleware(), viewBooking)
	fmt.Println("5")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/viewBooking", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

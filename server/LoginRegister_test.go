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

func TestSignUpTourist(t *testing.T) {
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Register{}, UserProfile{})
	assert.NoError(t, err)
	fmt.Println("1")
	r := gin.Default()
	// rgstr := Register{
	// 	Name:     "vishesha",
	// 	Email:    "vish@gmail.com",
	// 	Password: "Vishesha@123",
	// 	Role:     "Guide",
	// }
	// uspr := UserProfile{
	// 	Name:  "vishesha",
	// 	Email: "vish@gmail.com",
	// 	Role:  "Guide",
	// }
	// fmt.Println("6")

	// DB.Create(&rgstr)
	// DB.Create(&uspr)
	fmt.Println("7")
	r.POST("/signup", userRegister)
	fmt.Println("2")
	w := httptest.NewRecorder()
	fmt.Println("5")
	var reg = []byte(`{
		"name": "vvishesha",
		"email":"vvish3@gmail.com",
		"password":"Vishesha@123",
		"role":"Tourist"
	 }`)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewReader(reg))
	fmt.Println("3")
	r.ServeHTTP(w, req)
	assert.NoError(t, err)
	fmt.Println("4")
	assert.Equal(t, 200, w.Code)
}

func TestSignUpGuide(t *testing.T) {
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Register{}, GuideProfile{})
	assert.NoError(t, err)
	fmt.Println("1")
	r := gin.Default()
	// rgstr := Register{
	// 	Name:     "vishesha",
	// 	Email:    "vish@gmail.com",
	// 	Password: "Vishesha@123",
	// 	Role:     "Guide",
	// }
	// uspr := UserProfile{
	// 	Name:  "vishesha",
	// 	Email: "vish@gmail.com",
	// 	Role:  "Guide",
	// }
	// fmt.Println("6")

	// DB.Create(&rgstr)
	// DB.Create(&uspr)
	fmt.Println("7")
	r.POST("/signup", userRegister)
	fmt.Println("2")
	w := httptest.NewRecorder()
	fmt.Println("5")
	var reg = []byte(`{
		"name": "vvishesha1",
		"email":"vvish3@gmail.com",
		"password":"Vishesha@123",
		"role":"Guide"
	 }`)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewReader(reg))
	fmt.Println("3")
	r.ServeHTTP(w, req)
	assert.NoError(t, err)
	fmt.Println("4")
	assert.Equal(t, 200, w.Code)
}

func TestLoginTourist(t *testing.T) {
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Register{}, UserProfile{})
	assert.NoError(t, err)
	fmt.Println("1")
	r := gin.Default()

	fmt.Println("7")
	r.POST("/login", userLogin)
	fmt.Println("2")
	w := httptest.NewRecorder()
	fmt.Println("5")
	var reg = []byte(`{
		"name": "vvishesha",
		"email":"vvish3@gmail.com",
		"password":"Vishesha@123",
		"role":"Tourist"
	 }`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(reg))
	fmt.Println("3")
	r.ServeHTTP(w, req)
	assert.NoError(t, err)
	fmt.Println("4")
	assert.Equal(t, 200, w.Code)
}

func TestLoginGuide(t *testing.T) {
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Register{}, GuideProfile{})
	assert.NoError(t, err)
	fmt.Println("1")
	r := gin.Default()
	// rgstr := Register{
	// 	Name:     "vishesha",
	// 	Email:    "vish@gmail.com",
	// 	Password: "Vishesha@123",
	// 	Role:     "Guide",
	// }
	// uspr := UserProfile{
	// 	Name:  "vishesha",
	// 	Email: "vish@gmail.com",
	// 	Role:  "Guide",
	// }
	// fmt.Println("6")

	//DB.Create(&rgstr)
	// DB.Create(&uspr)
	fmt.Println("7")
	r.POST("/login", userLogin)
	fmt.Println("2")
	w := httptest.NewRecorder()
	fmt.Println("5")
	var reg = []byte(`{
		"name": "vvishesha1",
		"email":"vvish3@gmail.com",
		"password":"Vishesha@123",
		"role":"Guide"
	 }`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(reg))
	fmt.Println("3")
	r.ServeHTTP(w, req)
	assert.NoError(t, err)
	fmt.Println("4")
	assert.Equal(t, 200, w.Code)
}

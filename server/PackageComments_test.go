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

// func TestCreateComments(t *testing.T) {

// 	token, err := CreateToken("testG@gmail.com", "Tourist", "vish")
// 	assert.NoError(t, err)
// 	saveErr := CreateAuth("testG@gmail.com", "Tourist", token)
// 	assert.NoError(t, saveErr)

// 	//var DB *gorm.DB
// 	db, err := gorm.Open("sqlite3", "database.db")
// 	assert.NoError(t, err)
// 	DB = db
// 	DB.AutoMigrate(&Comment{})
// 	assert.NoError(t, err)
// 	fmt.Println("1")
// 	r := gin.Default()
// 	r.POST("/comments", TokenAuthMiddleware(), createComments)
// 	fmt.Println("2")
// 	w := httptest.NewRecorder()

// 	var cmnt = []byte(`{
// 		"title":     "not bad",
// 		"name":      "x123456789",
// 		"email":
// 		"packageId": "1",
// 		"review":    "just ok"
// 	 }`)

// 	req, _ := http.NewRequest("POST", "/comments", bytes.NewReader(cmnt))
// 	fmt.Println("3")
// 	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
// 	//fmt.Sprintf("Bearer %+v", token)
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
// 	fmt.Println("4")
// 	r.ServeHTTP(w, req)

// 	//err = json.Unmarshal(w.Body, &response)
// 	assert.NoError(t, err)

// 	assert.Equal(t, 200, w.Code)
// 	// assert.Equal(t, "test@email.com", response.Email)
// 	// assert.Equal(t, "Test User", response.Name)

// }

func TestCreateCommentsInvalid(t *testing.T) {

	token, err := CreateToken("testG@gmail.com", "Guide", "vish")
	assert.NoError(t, err)
	saveErr := CreateAuth("testG@gmail.com", "Guide", token)
	assert.NoError(t, saveErr)

	//var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Comment{})
	assert.NoError(t, err)
	fmt.Println("1")
	r := gin.Default()
	r.POST("/comments", TokenAuthMiddleware(), createComments)
	fmt.Println("2")
	w := httptest.NewRecorder()

	var cmnt = []byte(`{
		"title": "just ok",
		"name":"vish",
		"email":"testG@gmail.com",
		"packageId": "1",
		"review": "Average considering the price"
	 }`)

	req, _ := http.NewRequest("POST", "/comments", bytes.NewReader(cmnt))
	fmt.Println("3")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//fmt.Sprintf("Bearer %+v", token)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("4")
	r.ServeHTTP(w, req)

	//err = json.Unmarshal(w.Body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 401, w.Code)
	// assert.Equal(t, "test@email.com", response.Email)
	// assert.Equal(t, "Test User", response.Name)

}

func TestCreateComments(t *testing.T) {

	token, err := CreateToken("testG@gmail.com", "Tourist", "vish")
	assert.NoError(t, err)
	saveErr := CreateAuth("testG@gmail.com", "Tourist", token)
	assert.NoError(t, saveErr)

	//var DB *gorm.DB
	db, err := gorm.Open("sqlite3", "database.db")
	assert.NoError(t, err)
	DB = db
	DB.AutoMigrate(&Comment{})
	assert.NoError(t, err)
	fmt.Println("1")
	r := gin.Default()
	r.GET("/comments", TokenAuthMiddleware(), getallpackageComments)
	fmt.Println("2")
	w := httptest.NewRecorder()

	var pkgid = []byte(`{
		"packageId": "1"
	 }`)

	req, _ := http.NewRequest("GET", "/comments", bytes.NewReader(pkgid))
	fmt.Println("3")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//fmt.Sprintf("Bearer %+v", token)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("4")
	r.ServeHTTP(w, req)

	//err = json.Unmarshal(w.Body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "test@email.com", response.Email)
	// assert.Equal(t, "Test User", response.Name)

}

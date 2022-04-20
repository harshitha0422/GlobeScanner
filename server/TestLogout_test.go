package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogoutGuide(t *testing.T) {

	token, err := CreateToken("testG@gmail.com", "Guide", "123456789")
	assert.NoError(t, err)
	saveErr := CreateAuth("testG@gmail.com", "Guide", token)
	assert.NoError(t, saveErr)

	r := gin.Default()
	fmt.Println("4")
	r.POST("/logout", TokenAuthMiddleware(), Logout)
	fmt.Println("5")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/logout", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

func TestLogoutTourist(t *testing.T) {

	token, err := CreateToken("testt@gmail.com", "Tourist", "123456789")
	assert.NoError(t, err)
	saveErr := CreateAuth("testt@gmail.com", "Tourist", token)
	assert.NoError(t, saveErr)

	r := gin.Default()
	fmt.Println("4")
	r.POST("/logout", TokenAuthMiddleware(), Logout)
	fmt.Println("5")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/logout", nil)
	fmt.Sprintf("Bearer %s", token.AccessToken)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	fmt.Println("5")
	r.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

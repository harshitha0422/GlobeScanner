package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInvalidSearchPlaces(t *testing.T) {
	var response map[string]string

	token, err := CreateToken("vishesha@gmail.com", "Tourist", "vishesha@123")
	assert.NoError(t, err)
	saveErr := CreateAuth("vishesha@gmail.com", "Tourist", token)
	assert.NoError(t, saveErr)

	router := gin.Default()

	router.GET("/searchPlaces/:location", TokenAuthMiddleware(), searchPlaces)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/searchPlaces/", nil)
	//mt.Sprintf("Bearer %+v", token)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	router.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "Please enter a valid location name.", response["error"])
	// assert.Equal(t, "Test User", response.Name)

}

func TestValidSearchPlaces(t *testing.T) {

	token, err := CreateToken("vishesha@gmail.com", "Tourist", "vishesha@123")
	assert.NoError(t, err)
	saveErr := CreateAuth("vishesha@gmail.com", "Tourist", token)
	assert.NoError(t, saveErr)

	router := gin.Default()

	router.GET("/searchPlaces/:location", TokenAuthMiddleware(), searchPlaces)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/searchPlaces/Gainesville", nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	router.ServeHTTP(w, req)

	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)

}

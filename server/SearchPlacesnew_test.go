package main

// func TestInvalidSearchPlaces(t *testing.T) {
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

// func TestValidSearchPlaces(t *testing.T) {

// 	token, err := CreateToken("megan12@gmail.com", "Tourist", "Megan")
// 	assert.NoError(t, err)

// 	router := gin.Default()

// 	router.GET("/searchPlaces/:location", TokenAuthMiddleware(), searchPlaces)

// 	w := httptest.NewRecorder()

// 	req, _ := http.NewRequest("GET", "/searchPlaces/Gainesville", nil)
// 	//fmt.Sprintf("Bearer %+v", token)
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

// 	router.ServeHTTP(w, req)

// 	//err = json.Unmarshal(w.Body, &response)
// 	assert.NoError(t, err)

// 	assert.Equal(t, 200, w.Code)
// 	// assert.Equal(t, "test@email.com", response.Email)
// 	// assert.Equal(t, "Test User", response.Name)

// }

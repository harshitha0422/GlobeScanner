package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

// func TestDeleteProfile(t *testing.T) {
// 	//var response []SearchPlacesResponse
// 	var response string

// 	token, err := CreateToken("vishesha@gmail.com", "Tourist", "vishesha")
// 	assert.NoError(t, err)

// 	router := gin.Default()
// 	r.DELETE("/deleteProfile", TokenAuthMiddleware(), DeleteUserProfile)
// 	w := httptest.NewRecorder()

// 	req, _ := http.NewRequest("DELETE", "/deleteProfile", nil)
// 	fmt.Sprintf("Bearer %+v", token)
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

// 	router.ServeHTTP(w, req)

// 	err = json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	fmt.Println("DELETE PROFILE!")
// 	assert.Equal(t, 200, w.Code)
// 	// assert.Equal(t, "test@email.com", response.Email)
// 	// assert.Equal(t, "Test User", response.Name)

// 	// database.GlobalDB.Unscoped().Where("email = ?", user.Email).Delete(&models.User{})
// }

// func TestUserRegister(t *testing.T) {
// 	router := gin.Default()
// 	jsonData := []byte(`{
// 			"email": "vishesha@gmail.com",
// 			"password": "vishesha@123",
// 			"role": "Tourist",
// 			"name": "vishesha"
// 	}`)
// 	r.POST("/signup", userRegister)
// 	//w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonData))
// 	testHttpRequest(t, router, req, func(w *httptest.ResponseRecorder) bool {
// 		fmt.Println(w.Body)
// 		s := w.Code != 200
// 		return s
// 	})
// }

func fakeRegisterGenerate() Register {
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("can't connect to database")
	}
	DB = db
	// Create fake instructor
	type fakeRegisterStruct struct {
		Name     string `faker:"name"`
		Email    string `faker:"email"`
		Password string `faker:"password"`
		Role     string `faker:"role"`
	}
	fakeregister := fakeRegisterStruct{}
	err = faker.FakeData(&fakeregister)
	if err != nil {
		panic("Failed to create fake data")
	}
	var newregister = Register{
		Email:    fakeregister.Email,
		Password: fakeregister.Password,
		Name:     fakeregister.Name,
		Role:     fakeregister.Role,
	}
	result := DB.Create(&newregister)
	if result.Error != nil {
		panic("Failed to create an instructor")
	}
	return newregister
}

func fakeUserProfileGenerate() UserProfile {

	// Create instructor

	// Login
	newregister := fakeRegisterGenerate()
	// Create course

	// Create fake course
	type fakeProfileStruct struct {
		Email    string `faker:"email"`
		Name     string `faker:"name"`
		Role     string `faker:"role"`
		About    string `faker:"about"`
		Age      uint   `faker:"age"`
		Mobile   string `faker:"mobile"`
		Location string `faker:"location"`
		Fav1     string `faker:"fav1"`
		Fav2     string `faker:"fav2"`
		Fav3     string `faker:"fav3"`
	}
	fakeprofile := fakeProfileStruct{}
	err := faker.FakeData(&fakeprofile)
	if err != nil {
		panic("Failed to create fake profile")
	}
	newprofile := UserProfile{
		Email:    newregister.Email,
		Name:     newregister.Name,
		Role:     newregister.Role,
		About:    fakeprofile.About,
		Age:      fakeprofile.Age,
		Mobile:   fakeprofile.Mobile,
		Location: fakeprofile.Location,
		Fav1:     fakeprofile.Fav1,
		Fav2:     fakeprofile.Fav2,
		Fav3:     fakeprofile.Fav3,
	}
	result := DB.Create(&newprofile)
	if result.Error != nil {
		panic("Failed to create a course")
	}
	// courseId := int(newCourse.ID)
	return newprofile
}

func fakeGuideProfileGenerate() GuideProfile {

	// Create instructor

	// Login
	newregister := fakeRegisterGenerate()
	// Create course

	// Create fake course
	type fakeProfileStruct struct {
		Email    string `faker:"email"`
		Name     string `faker:"name"`
		Role     string `faker:"role"`
		About    string `faker:"about"`
		Age      uint   `faker:"age"`
		Address  string `faker:"address"`
		Location string `faker:"location"`
		Vehicle  string `faker:"vehicle"`
	}
	fakeprofile := fakeProfileStruct{}
	err := faker.FakeData(&fakeprofile)
	if err != nil {
		panic("Failed to create fake profile")
	}
	newprofile2 := GuideProfile{
		Email:    newregister.Email,
		Name:     newregister.Name,
		Role:     newregister.Role,
		About:    fakeprofile.About,
		Age:      fakeprofile.Age,
		Address:  fakeprofile.Address,
		Location: fakeprofile.Location,
		Vehicle:  fakeprofile.Vehicle,
	}
	result := DB.Create(&newprofile2)
	if result.Error != nil {
		panic("Failed to create a course")
	}
	// courseId := int(newCourse.ID)
	return newprofile2
}

// func fakeCommentGenerate() Comment {

// 	newregister := fakeRegisterGenerate()
// 	// Create course

// 	// Create fake course
// 	type fakeCommentStruct struct {
// 		Comment  string `faker:"comment"`
// 		Email    string `faker:"email"`
// 		Name     string `faker:"name"`
// 		Location string `faker:"location"`
// 		Rating   string `faker:"rating"`
// 	}
// 	fakecomment := fakeCommentStruct{}
// 	err := faker.FakeData(&fakecomment)
// 	if err != nil {
// 		panic("Failed to create fake profile")
// 	}
// 	comments := Comment{
// 		Comment:  fakecomment.Comment,
// 		Email:    newregister.Email,
// 		Name:     fakecomment.Name,
// 		Location: fakecomment.Location,
// 		Rating:   fakecomment.Rating,
// 	}
// 	result := DB.Create(&comments)
// 	if result.Error != nil {
// 		panic("Failed to create a course")
// 	}
// 	// courseId := int(newCourse.ID)
// 	return comments
// }

func fakePackageGenerate() Package {

	newregister := fakeRegisterGenerate()
	// Create course

	// Create fake course
	type fakePackageStruct struct {
		Email        string `faker:"email"`
		Duration     string `faker: "duration"`
		Location     string `faker:"location"`
		Accomodation string `faker:"accomodation"`
		Itinerary    string `faker:"itinerary"`
		Included     string `faker:"included"`
		Price        string `faker:"price"`
	}
	fakepackage := fakePackageStruct{}
	err := faker.FakeData(&fakepackage)
	if err != nil {
		panic("Failed to create fake package")
	}
	packages := Package{
		Email:        newregister.Email,
		Duration:     fakepackage.Duration,
		Location:     fakepackage.Location,
		Accomodation: fakepackage.Accomodation,
		Itinerary:    fakepackage.Itinerary,
		Included:     fakepackage.Included,
		Price:        fakepackage.Price,
	}
	result := DB.Create(&packages)
	if result.Error != nil {
		panic("Failed to create a course")
	}
	// courseId := int(newCourse.ID)
	return packages
}

// func TestRegistrationExist(t *testing.T) {
// 	// Create register
// 	want := true
// 	got := true
// 	register := fakeRegisterGenerate()
// 	// Check if the course exists
// 	result := DB.Where("email = ?", register.Email).Find(&register)
// 	if result.Error != nil {
// 		want = false
// 	} else {
// 		want = true
// 	}
// 	if got != want {
// 		t.Errorf("Course exists test failed")
// 	} else {
// 		fmt.Println("Course exists Test passed!")
// 	}
// }

func TestGenerateToken(t *testing.T) {

	ts, err := CreateToken("saduvish@gmail.com", "tourist", "vish")
	assert.NoError(t, err)

	os.Setenv("testToken", ts.AccessToken)
}

func TestCreateAuth(t *testing.T) {
	ts, err := CreateToken("saduvish@gmail.com", "tourist", "vish")
	assert.NoError(t, err)
	saveErr := CreateAuth("saduvish@gmail.com", "tourist", ts)
	assert.NoError(t, saveErr)
}

// func TestFetchAuth(t *testing.T) {
// 	var client *redis.Client
// 	init()
// 	as := AccessDetails{
// 		AccessUuid: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjA4MDU0YzZjLTEwMDUtNDgyNS1iZTJkLWQ4ZmQ3NDljODgxNSIsImF1dGhvcml6ZWQiOnRydWUsImVtYWlsIjoidmlzaEBnbWFpbC5jb20iLCJleHAiOjE2NDg3Mzc1MjcsIm5hbWUiOiJ2aXNoIiwicm9sZSI6IlRvdXJpc3QifQ.e5dPkMjV9Rvzm11dQ3unkW11ROTPxpbT1AT1GYen29w",
// 		Email:      "vish@gmail.com",
// 		Role:       "Tourist",
// 		Name:       "vish",
// 	}
// 	err := FetchAuth(&as)
// 	assert.NoError(t, err)
// }

func TestFetchEmail(t *testing.T) {
	want := true
	got := true
	st := ""
	as := AccessDetails{
		AccessUuid: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjA4MDU0YzZjLTEwMDUtNDgyNS1iZTJkLWQ4ZmQ3NDljODgxNSIsImF1dGhvcml6ZWQiOnRydWUsImVtYWlsIjoidmlzaEBnbWFpbC5jb20iLCJleHAiOjE2NDg3Mzc1MjcsIm5hbWUiOiJ2aXNoIiwicm9sZSI6IlRvdXJpc3QifQ.e5dPkMjV9Rvzm11dQ3unkW11ROTPxpbT1AT1GYen29w",
		Email:      "vish@gmail.com",
		Role:       "Tourist",
		Name:       "vish",
	}
	st = FetchEmail(&as)
	if st == "" {
		want = false
	}
	if got != want {
		t.Errorf("Cannot fetch email from Access details")
	} else {
		fmt.Println("Email is fetched from Access details!")
	}
}

func TestFetchRole(t *testing.T) {
	want := true
	got := true
	st := ""
	as := AccessDetails{
		AccessUuid: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjA4MDU0YzZjLTEwMDUtNDgyNS1iZTJkLWQ4ZmQ3NDljODgxNSIsImF1dGhvcml6ZWQiOnRydWUsImVtYWlsIjoidmlzaEBnbWFpbC5jb20iLCJleHAiOjE2NDg3Mzc1MjcsIm5hbWUiOiJ2aXNoIiwicm9sZSI6IlRvdXJpc3QifQ.e5dPkMjV9Rvzm11dQ3unkW11ROTPxpbT1AT1GYen29w",
		Email:      "vish@gmail.com",
		Role:       "Tourist",
		Name:       "vish",
	}
	st = FetchRole(&as)
	if st == "" {
		want = false
	}
	if got != want {
		t.Errorf("Cannot fetch email from Access details")
	} else {
		fmt.Println("Role is fetched from Access details!")
	}
}

func TestFetchName(t *testing.T) {
	want := true
	got := true
	st := ""
	as := AccessDetails{
		AccessUuid: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjA4MDU0YzZjLTEwMDUtNDgyNS1iZTJkLWQ4ZmQ3NDljODgxNSIsImF1dGhvcml6ZWQiOnRydWUsImVtYWlsIjoidmlzaEBnbWFpbC5jb20iLCJleHAiOjE2NDg3Mzc1MjcsIm5hbWUiOiJ2aXNoIiwicm9sZSI6IlRvdXJpc3QifQ.e5dPkMjV9Rvzm11dQ3unkW11ROTPxpbT1AT1GYen29w",
		Email:      "vish@gmail.com",
		Role:       "Tourist",
		Name:       "vish",
	}
	st = FetchName(&as)
	if st == "" {
		want = false
	}
	if got != want {
		t.Errorf("Cannot fetch email from Access details")
	} else {
		fmt.Println("Role is fetched from Access details!")
	}
}

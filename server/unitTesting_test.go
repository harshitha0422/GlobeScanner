package main

import (
	"github.com/bxcodec/faker"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

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

func fakeCommentGenerate() Comment {

	newregister := fakeRegisterGenerate()
	// Create course

	// Create fake course
	type fakeCommentStruct struct {
		Comment  string `faker:"comment"`
		Email    string `faker:"email"`
		Name     string `faker:"name"`
		Location string `faker:"location"`
		Rating   string `faker:"rating"`
	}
	fakecomment := fakeCommentStruct{}
	err := faker.FakeData(&fakecomment)
	if err != nil {
		panic("Failed to create fake profile")
	}
	comments := Comment{
		Comment:  fakecomment.Comment,
		Email:    newregister.Email,
		Name:     fakecomment.Name,
		Location: fakecomment.Location,
		Rating:   fakecomment.Rating,
	}
	result := DB.Create(&comments)
	if result.Error != nil {
		panic("Failed to create a course")
	}
	// courseId := int(newCourse.ID)
	return comments
}

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
// 	// Create instructor
// 	// Login
// 	// Create course

// 	c := gin.Context{}

// 	course := fakeRegisterGenerate()
// 	// Check if the course exists
// 	got := courseExist(int(course.ID), &c)
// 	want := true
// 	if got != want {
// 		t.Errorf("Course exists test failed")
// 	} else {
// 		fmt.Println("Course exists Test passed!")
// 	}
// }

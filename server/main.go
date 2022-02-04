package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
var r *gin.Engine

func main() {

	DB, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("can't connect to database")
	}
	//DB=db
	DB.AutoMigrate(&Register{}, &UserProfile{}, &GuideProfile{}, &Comment{})
	//seed(db)

	defer DB.Close()
	DB.LogMode(true)
	r := gin.Default()
	r.GET("/searchPlaces/:name", showSearchPlacesPage)
	r.POST("/sign up", userRegister)
	r.POST("/login", userLogin)
	r.GET("/users", getallUsers)
	r.GET("/userprofile", getallTouristprofile)
	r.GET("/guideprofile", getallGuideprofile)
	r.GET("/comments", getallComments)
	r.GET("/users/:email", getUser)
	r.GET("/userprofile/:email", getUserProfile)
	r.GET("/guideprofile/:email", getGuideProfile)
	r.GET("/comments/:location", getLocationComments)
	r.POST("/userprofile", createUserProfile)
	r.POST("/guideprofile", createGuideProfile)
	r.POST("/comments", createComments)
	r.PUT("/userprofile/:email", updateUserProfile)
	r.PUT("/guideprofile/:email", updateGuideProfile)
	r.DELETE("/userprofile/:email", DeleteUserProfile)
	r.DELETE("/guideprofile/:email", DeleteGuideProfile)

	//initializeRoutes()
	err = r.Run()
	if err != nil {
		panic("Failed to run the server")
	}
}

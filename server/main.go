package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
var r *gin.Engine

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Methods, Access-Control-Allow-Origin ,Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token,x-access-token,X-Access-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {

	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("can't connect to database")
	}
	DB = db
	// db.Delete(&Register{})
	// db.Delete(&UserProfile{})
	// db.Delete(&GuideProfile{})

	// db.Delete(&Comment{})
	// db.Delete(&Package{})
	// db.Delete(&Booking{})

	DB.AutoMigrate(&Register{}, &UserProfile{}, &GuideProfile{}, &Comment{}, &Package{}, &Booking{})
	//seed(db)

	db.LogMode(true)
	r := gin.Default()

	/*r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With","*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))*/
	r.Use(CORSMiddleware())

	r.GET("/searchPlaces/:location", TokenAuthMiddleware(), searchPlaces)

	r.GET("/searchPackage/:location", TokenAuthMiddleware(), searchPackage)
	r.GET("/viewOwnPackage", TokenAuthMiddleware(), getPackage)
	r.GET("/viewBooking", TokenAuthMiddleware(), viewBooking)
	r.GET("/getAllPackage", TokenAuthMiddleware(), getallPackages)
	r.POST("/addPackages", TokenAuthMiddleware(), addPackages)
	r.POST("/bookPackages", TokenAuthMiddleware(), bookPackages)

	r.POST("/signup", userRegister)
	r.POST("/login", userLogin)

	r.GET("/userprofile", TokenAuthMiddleware(), getUserProfile)
	r.DELETE("/deleteProfile", TokenAuthMiddleware(), DeleteUserProfile)

	r.PUT("/updateUserProfile", TokenAuthMiddleware(), updateUserProfile)
	r.PUT("/updateGuideProfile", TokenAuthMiddleware(), updateGuideProfile)

	r.GET("/userprofiles", TokenAuthMiddleware(), getallTouristprofile)
	r.GET("/guideprofiles", TokenAuthMiddleware(), getallGuideprofile)
	r.GET("/guideprofiles/:location", TokenAuthMiddleware(), getGuideProfileLocation)

	//r.GET("/users", getUser)
	//r.POST("/userprofile", TokenAuthMiddleware(), createUserProfile)
	//r.GET("/guideprofile/:email", getGuideProfile)

	//r.POST("/userprofiles", createTouristProfile)
	//r.POST("/guideprofiles", createGuideProfile)
	r.GET("/comments", TokenAuthMiddleware(), getallpackageComments)
	r.GET("/mycomments", TokenAuthMiddleware(), getalluserComments)
	r.POST("/comments", TokenAuthMiddleware(), createComments)
	r.GET("/comments/:location", TokenAuthMiddleware(), getLocationComments)

	r.POST("/token/refresh", Refresh)
	r.POST("/todo", TokenAuthMiddleware(), CreateTodo)
	r.POST("/logout", TokenAuthMiddleware(), Logout)

	//initializeRoutes()
	err = r.Run()
	if err != nil {
		panic("Failed to run the server")
	}

	defer db.Close()
}

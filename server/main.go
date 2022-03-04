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
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token,x-access-token,X-Access-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
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
	DB.AutoMigrate(&Register{}, &UserProfile{}, &GuideProfile{}, &Comment{})
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
	r.GET("/searchPlaces/:name", showSearchPlacesPage)
	r.POST("/signup", userRegister)
	r.POST("/login", userLogin)
	r.GET("/users", getallUsers)
	r.GET("/userprofiles", TokenAuthMiddleware(), getallTouristprofile)
	r.GET("/guideprofiles", getallGuideprofile)
	r.GET("/comments", getallComments)
	//r.GET("/users", getUser)
	r.GET("/userprofile", getTouristProfile)
	r.GET("/guideprofile/:email", getGuideProfile)
	r.GET("/comments/:location", getLocationComments)
	r.POST("/userprofiles", createTouristProfile)
	r.POST("/guideprofiles", createGuideProfile)
	r.POST("/comments", createComments)
	r.PUT("/userprofile/:email", updateTouristProfile)
	r.PUT("/guideprofile/:email", updateGuideProfile)
	r.DELETE("/userprofile/:email", DeleteTouristProfile)
	r.DELETE("/guideprofile/:email", DeleteGuideProfile)

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

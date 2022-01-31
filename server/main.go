package main

import (
	// "net/http"
	// "os"

	// "github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	pass := os.Getenv("DB_PASS")
// 	db, err := gorm.Open(
// 		"postgres",
// 		"host=students-db user=go password="+pass+" dbname=go sslmode=disable")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	app := App{
// 		db: db,
// 		r:  mux.NewRouter(),
// 	}
// 	app.start()
// }

// main.go

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	//router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	//initializeRoutes()

	router.GET("/searchPlaces/:name", showIndexPage)

	// Start serving the application
	router.Run()

}

package main

import (
	// "net/http"
	// "os"

	// "github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Initialize the routes
	initializeRoutes()

	//router.GET("/searchPlaces/:name", showIndexPage)

	// Start serving the application
	router.Run()

}

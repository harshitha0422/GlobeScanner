// handlers.article.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showSearchPlacesPage(c *gin.Context) {

	fmt.Println("Inside showIndex")
	name := c.Param("name")
	fmt.Print(name)
	articles := searchPlaces(name)

	fmt.Print(articles)

	// Call the HTML method of the Context to render a template
	// c.HTML(
	// 	// Set the HTTP status to 200 (OK)
	// 	http.StatusOK,
	// 	// Use the index.html template
	// 	 "index.html",
	// 	// Pass the data that the page uses
	// 	gin.H{
	// 		"title":   "Home Page",
	// 		"payload": articles,
	// 	},
	// )

	c.JSON(http.StatusOK, gin.H{"msg": articles})

}

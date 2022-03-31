package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
)

// fmt.Println("Inside search Places")
// fmt.Println("Place" + name)
//w.Header().Set("Content-Type", "application/json")

// calling the geoname api
//eg: https://api.opentripmap.com/0.1/en/places/geoname?apikey=5ae2e3f221c38a28845f05b6ddba4f8e8b04b5c6bbb6c180ea77c286&name=London

func searchPackage(c *gin.Context) {
	fmt.Println("inside search package")
	location := c.Param("location")
	fmt.Println(location)
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	role := FetchRole(tokenAuth)
	fmt.Println(role)
	/*if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized2")
		fmt.Println(err)
		return
	}*/

	// if role == "Guide" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "User needs to be a tourist. Please register as a tourist.",
	// 	})
	// 	return
	// }

	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please enter a valid location name.",
		})
		return

	}

	//email := c.Params.ByName("email"
	pkg := Package{}
	packages := DB.Where("location = ?", location).Find(&pkg)
	if packages.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No package currently available for this location.",
		})
		return
	} else {
		c.JSON(200, packages)
	}
}

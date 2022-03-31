package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPackage(c *gin.Context) {
	fmt.Println("inside search package")
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized1")
		return
	}
	err1 := FetchAuth(tokenAuth)
	if err1 != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized2")
		fmt.Println(err1)
		return
	}
	//email := c.Params.ByName("email")
	email := FetchEmail(tokenAuth)
	role := FetchRole(tokenAuth)
	fmt.Println(email)
	fmt.Println(role)

	if role == "Tourist" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User needs to be a Guide. Please register as a Guide.",
		})
		return
	}
	var pkg []Package
	packages := DB.Where("email = ?", email).Find(&pkg)
	if packages.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No package currently available for this user.",
		})
		return
	} else {
		c.JSON(200, packages)
	}
}

func getallPackages(c *gin.Context) {
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized1")
		return
	}
	err1 := FetchAuth(tokenAuth)
	if err1 != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized2")
		fmt.Println(err1)
		return
	}

	var pkg []Package
	if err := DB.Find(&pkg).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pkg)
	}
}

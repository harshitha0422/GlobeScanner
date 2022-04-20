package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createComments(c *gin.Context) {
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
	email := FetchEmail(tokenAuth)
	role := FetchRole(tokenAuth)
	name := FetchName(tokenAuth)

	if role == "Guide" {
		c.JSON(401, "error: User needs to be a tourist to create a comment")
		// c.JSON(http.StatusUnauthorized, gin.H{
		// 	"error": "User needs to be a tourist to create a comment",
		// })
		return
	}

	type ip struct {
		PackageId string `json:"packageId"`
		Title     string `json:"title"`
		Review    string `json:"review"`
	}

	req := ip{}
	error := c.BindJSON(&req)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}
	comment := Comment{
		Name:      name,
		Email:     email,
		Title:     req.Title,
		Review:    req.Review,
		PackageId: req.PackageId,
	}
	c.BindJSON(&comment)
	if err := DB.Create(&comment).Error; err != nil {
		fmt.Println("cannot create comment", err)
		c.AbortWithStatus(502)

	}

	c.JSON(200, comment)
}

func getallpackageComments(c *gin.Context) {

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
	type inp struct {
		PackageId string `json:"packageId"`
	}

	req := inp{}
	error := c.BindJSON(&req)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}

	var comments []Comment
	if err := DB.Where("package_id = ?", req.PackageId).Find(&comments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comments)
	}

}

func getalluserComments(c *gin.Context) {

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
	email := FetchEmail(tokenAuth)
	var comments []Comment
	if err := DB.Where("email = ?", email).Find(&comments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comments)
	}

}

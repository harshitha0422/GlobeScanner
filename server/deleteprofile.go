package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserProfile(c *gin.Context) {

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
	if role == "Tourist" {
		var userprofile UserProfile
		var register Register
		d := DB.Where("email = ?", email).Delete(&userprofile)
		fmt.Println(d)
		res := DB.Where("email = ? AND role = ?", email, role).Delete(&register)
		fmt.Println(res)
		c.JSON(200, gin.H{"user profile with Email #" + email: "deleted"})
	}
	if role == "Guide" {
		var guideprofile GuideProfile
		var register Register
		d := DB.Where("email = ?", email).Delete(&guideprofile)
		fmt.Println(d)
		res := DB.Where("email = ? AND role = ?", email, role).Delete(&register)
		fmt.Println(res)
		c.JSON(200, gin.H{"user profile with Email #" + email: "deleted"})
	}
}

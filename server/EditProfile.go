package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func updateUserProfile(c *gin.Context) {

	var profile UserProfile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid input")
		return
	}

	// tokenAuth, err := ExtractTokenMetadata(c.Request)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, "unauthorized")
	// 	return
	// }
	// email := FetchAuth(tokenAuth)
	//role := FetchRole(tokenAuth)
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
	//if role == "Tourist" {
	var userprofile UserProfile
	var register Register
	// if err := DB.Where("email = ?", email).First(&userprofile).Error; err != nil {

	userprofile = UserProfile{
		Name:     profile.Name,
		Role:     role,
		About:    profile.About,
		Age:      profile.Age,
		Mobile:   profile.Mobile,
		Location: profile.Location,
		Fav1:     profile.Fav1,
		Fav2:     profile.Fav2,
		Fav3:     profile.Fav3,
	}
	register = Register{
		//gorm.Model
		Name: profile.Name,
	}

	//DB.Save(&userprofile)
	result := DB.Model(&userprofile).Where("email = ?", email).Updates(&userprofile)
	result1 := DB.Model(&register).Where("email = ?", email).Updates(&register)

	if result.Error != nil {

		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	if result1.Error != nil {

		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"success": "Tourist data successfully updated.",
	})
}

func updateGuideProfile(c *gin.Context) {

	var profile GuideProfile
	if err := c.BindJSON(&profile); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid input")
		return
	}

	// tokenAuth, err := ExtractTokenMetadata(c.Request)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, "unauthorized")
	// 	return
	// }
	// email := FetchAuth(tokenAuth)
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
	var guideprofile GuideProfile
	var register Register

	guideprofile = GuideProfile{
		Email:    email,
		Name:     profile.Name,
		Role:     role,
		About:    profile.About,
		Age:      profile.Age,
		Address:  profile.Address,
		Location: profile.Location,
		Vehicle:  profile.Vehicle,
	}
	register = Register{
		//gorm.Model
		Name: profile.Name,
	}
	// if err := DB.Where("email = ?", email).First(&guideprofile).Error; err != nil {
	// 	c.AbortWithStatus(404)
	// 	fmt.Println(err)
	// }
	// c.BindJSON(&guideprofile)
	// DB.Save(&guideprofile)

	result := DB.Model(&guideprofile).Where("email = ?", email).Updates(&guideprofile)
	result1 := DB.Model(&register).Where("email = ?", email).Updates(&register)

	if result.Error != nil {

		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	if result1.Error != nil {

		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"success": "Guide data successfully updated.",
	})
}

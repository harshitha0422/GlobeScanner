// handlers.article.go

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//
func showSearchPlacesPage(c *gin.Context) {

	//fmt.Println("Inside showIndex")
	name := c.Param("name")
	// if len(name) == 0 {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": " Search name cannot be empty",
	// 	})
	// 	return
	// }

	//fmt.Print(name)
	places := searchPlaces(name)

	//fmt.Print(places)

	c.JSON(http.StatusOK, gin.H{"msg": places})

}

//fetch all the users (guide & tourist)
func getallUsers(c *gin.Context) {
	var register []Register
	if err := DB.Find(&register).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, register)
	}
}

// get all tourist's profiles
func getallTouristprofile(c *gin.Context) {
	var userprofile []UserProfile
	if err := DB.Find(&userprofile).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, userprofile)
	}
}

// get all travel guides profiles
func getallGuideprofile(c *gin.Context) {
	var guideprofile []GuideProfile
	if err := DB.Find(&guideprofile).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, guideprofile)
	}
}

// get all comments
func getallComments(c *gin.Context) {

	var comments []Comment
	if err := DB.Find(&comments).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, comments)
	}

}

// get one particular user(tourist and travel guide) by email
func getUser(c *gin.Context) {

	tokenAuth, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized1")
		return
	}
	email := FetchAuth(tokenAuth)
	/*if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized2")
		return
	}*/

	//email := c.Params.ByName("email")
	var register Register
	if err := DB.Where("email = ?", email).First(&register).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, register)
	}
}

// get particular tourist profile by email
func getUserProfile(c *gin.Context) {
	fmt.Println("inside get user profile")
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

		var userprofile UserProfile
		if err := DB.Where("email = ?", email).First(&userprofile).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, userprofile)
		}
	} else {
		var guideprofile GuideProfile
		if err := DB.Where("email = ?", email).First(&guideprofile).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, guideprofile)
		}
	}
}

// // get one travel guide profile by email
// func getGuideProfile(c *gin.Context) {

// 	email := c.Params.ByName("email")
// 	var guideprofile GuideProfile
// 	if err := DB.Where("email = ?", email).First(&guideprofile).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, guideprofile)
// 	}
// }

// get comments by location
func getLocationComments(c *gin.Context) {

	location := c.Params.ByName("location")
	var locationcomment []Comment
	if err := DB.Where("location = ?", location).Find(&locationcomment).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, locationcomment)
	}
}

func createUserProfile(c *gin.Context) {
	fmt.Println("inside get user profile")
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
		var userprofile UserProfile
		c.BindJSON(&userprofile)
		if err := DB.Create(&userprofile).Error; err != nil {
			fmt.Println("cannot create user profile", err)
			c.AbortWithStatus(502)

		}
		c.JSON(200, userprofile)
	} else {
		var guideprofile GuideProfile
		c.BindJSON(&guideprofile)
		if err := DB.Create(&guideprofile).Error; err != nil {
			fmt.Println("cannot create user profile", err)
			c.AbortWithStatus(502)

		}
		c.JSON(200, guideprofile)
	}

}

// create the tourist profile
func createTouristProfile(c *gin.Context) {
	var userprofile UserProfile
	c.BindJSON(&userprofile)
	if err := DB.Create(&userprofile).Error; err != nil {
		fmt.Println("cannot create user profile", err)
		c.AbortWithStatus(502)

	}
	c.JSON(200, userprofile)
}

// create the guide profile
func createGuideProfile(c *gin.Context) {
	var guideprofile GuideProfile
	c.BindJSON(&guideprofile)
	if err := DB.Create(&guideprofile).Error; err != nil {
		fmt.Println("cannot create user profile", err)
		c.AbortWithStatus(502)

	}
	c.JSON(200, guideprofile)
}

func createComments(c *gin.Context) {
	var comment Comment
	c.BindJSON(&comment)
	if err := DB.Create(&comment).Error; err != nil {
		fmt.Println("cannot create user profile", err)
		c.AbortWithStatus(502)

	}
	c.JSON(200, comment)
}

func updateUserProfile(c *gin.Context) {
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized1")
		return
	}
	email := FetchAuth(tokenAuth)
	role := FetchRole(tokenAuth)

	if role == "Tourist" {
		var userprofile UserProfile
		if err := DB.Where("email = ?", email).First(&userprofile).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": userprofile})
		var userprofile1 UserProfile
		c.BindJSON(&userprofile1)
		c.JSON(http.StatusOK, gin.H{"data": userprofile1})
		//DB.Save(&userprofile)
		DB.Model(&userprofile).Updates(userprofile1)
		c.JSON(200, gin.H{
			"success": "Tourist data successfully updated.",
		})
		c.JSON(http.StatusOK, gin.H{"data": userprofile})
	} else {
		var guideprofile GuideProfile
		if err := DB.Where("email = ?", email).First(&guideprofile).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.BindJSON(&guideprofile)
		DB.Save(&guideprofile)
		c.JSON(200, gin.H{
			"success": "Guide data successfully updated.",
		})
	}
}

func DeleteTouristProfile(c *gin.Context) {
	email := c.Params.ByName("email")
	var userprofile UserProfile
	d := DB.Where("email = ?", email).Delete(&userprofile)
	fmt.Println(d)
	c.JSON(200, gin.H{"user profile with Email #" + email: "deleted"})
}

func DeleteGuideProfile(c *gin.Context) {
	email := c.Params.ByName("email")
	var guideprofile GuideProfile
	d := DB.Where("email = ?", email).Delete(&guideprofile)
	fmt.Println(d)
	c.JSON(200, gin.H{"profile with Email #" + email: " deleted"})
}

func upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

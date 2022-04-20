package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func bookPackages(c *gin.Context) {

	fmt.Print("book package req")
	tokenAuth, err := ExtractTokenMetadata(c.Request)
	fmt.Println(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Error extracting the token")
		return
	}
	//err1 := FetchAuth(tokenAuth)

	email := FetchEmail(tokenAuth)

	role := FetchRole(tokenAuth)
	// if err1 != nil {
	// 	c.JSON(http.StatusUnauthorized, "Unauthorized")
	// 	fmt.Println(err1)
	// 	return
	// }

	if role == "Guide" {
		c.JSON(401, "User needs to be a Guide. Please register as a Guide.")
		// c.JSON(http.StatusUnauthorized, gin.H{
		// 	"error": "User needs to be a Tourist. Please register as a Tourist.",
		// })
		return
	}

	// Parse input request
	// type Req struct {
	// 	Email     string `json:"email" binding:"required,email"`
	// 	PackageId string `json:"packageId" binding:"required,packageId"`
	// }
	//var req Req
	//var newPackage Package
	req := Booking{}
	error := c.BindJSON(&req)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}
	// Check if touristemail exists
	//register := Register{}
	// tourist := UserProfile{}
	// //result := DB.Where("email = ?", newPackage.Email).First(&register)
	// result := DB.Where("email = ?", email).First(&tourist)
	// if result.Error != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Internal server error",
	// 	})
	// 	return
	// }
	// if result.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Tourist account needs to be created.",
	// 	})
	// 	return
	// }

	//checking if packageId is correct
	// pkg := Package{}
	// //result := DB.Where("email = ?", newPackage.Email).First(&register)
	// res := DB.Where("id = ?", req.PackageId).First(&pkg)
	// if res.Error != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Internal server error",
	// 	})
	// 	return
	// }
	// if res.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Package id is wrong.",
	// 	})
	// 	return
	// }
	// Insert into database
	//var newPackage Package
	newBooking := Booking{
		Email:     email,
		PackageId: req.PackageId,
	}
	bookingRes := DB.Create(&newBooking)
	if bookingRes.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(200, "msg:Booking created successfully.")
}

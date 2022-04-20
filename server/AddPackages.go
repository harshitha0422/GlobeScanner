package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPackages(c *gin.Context) {

	fmt.Print("add package req")
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

	if role == "Tourist" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User needs to be a Guide. Please register as a Guide.",
		})
		return
	}

	// Parse input request
	type Req struct {
		Email        string `json:"email"`
		Duration     string `json:"duration"`
		Location     string `json:"location"`
		Accomodation string `json:"accomodation"`
		Itinerary    string `json:"itinerary"`
		Included     string `json:"included"`
		Price        string `json:"price"`
	}
	//var req Req
	//var newPackage Package
	req := Req{}
	error := c.BindJSON(&req)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}
	// Check if guideemail exists
	//register := Register{}
	guide := GuideProfile{}
	//result := DB.Where("email = ?", newPackage.Email).First(&register)
	result := DB.Where("email = ?", email).First(&guide)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal server error",
		})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Travel guide account needs to be created.",
		})
		return
	}
	// Insert into database
	//var newPackage Package
	newPackage := Package{
		Email:        email,
		Duration:     req.Duration,
		Location:     req.Location,
		Accomodation: req.Accomodation,
		Itinerary:    req.Itinerary,
		Included:     req.Included,
		Price:        req.Price,
	}
	result = DB.Create(&newPackage)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, newPackage)
}

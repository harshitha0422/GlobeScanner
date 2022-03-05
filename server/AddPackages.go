package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPackages(c *gin.Context) {

	fmt.Print("add package req")
	// Parse input request
	type Req struct {
		GuideEmail   string `json:"guideEmail" gorm:"not null"`
		Duration     string `json: "duration"`
		Location     string `json:"location"`
		Accomodation string `json:"accomodation"`
		Itinerary    string `json:"itinerary"`
		Included     string `json:"included"`
	}
	req := Req{}
	//var newPackage Package

	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}
	// Check if guideemail exists
	//register := Register{}
	register := Register{}
	//result := DB.Where("email = ?", newPackage.Email).First(&register)
	result := DB.Where("email = ?", req.GuideEmail).First(&register)
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
		GuideEmail:   req.GuideEmail,
		Duration:     req.Duration,
		Location:     req.Location,
		Accomodation: req.Accomodation,
		Itinerary:    req.Itinerary,
		Included:     req.Included,
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

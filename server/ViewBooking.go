package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func viewBooking(c *gin.Context) {
	fmt.Println("inside viewBooking")
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

	if role == "Guide" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User needs to be a Tourist.",
		})
		return
	}
	// var bkg []Booking
	// res := DB.Select("package_id").Where("email = ?", email).Find(&bkg)
	// if res.Error != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "Error occured.Please try again.",
	// 	})
	// } else {
	// 	var pkg []Package
	// 	res := DB.Where("package_id in ?", bkg).Find(&pkg)
	// 	if res.Error != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"error": "Error occured.Please try again.",
	// 		})

	// 		return

	// 	} else {
	// 		c.JSON(200, pkg)
	// 	}
	// }
	// type ID struct {
	// 	packageId string
	// }
	// subQuery := DB.Model(&Booking{}).Select("package_id").Where("email = ?", email).Find(&ID{})
	// //db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)
	// var pkg []Package
	// res := DB.Where("Id IN ?", subQuery).Find(&pkg)
	// if res.Error != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "Error occured.Please try again.",
	// 	})

	// 	return

	// } else {
	// 	c.JSON(200, pkg)
	// }

	//dbRes := DB.Model(&Course{}).Select("courses.*").
	// Joins("inner join instructors on courses.instructor_id = instructors.id").
	// Where("instructors.email = ?", email).
	// Where("courses.id = ?", courseId).
	// First(&course)
	type PackageResponse struct {
		Id           string
		Email        string
		Duration     string
		Location     string
		Accomodation string
		Itinerary    string
		Included     string
		Price        string
	}
	pkg := []PackageResponse{}
	res := DB.Table("packages").Select("*").
		Joins("inner join bookings on packages.id = bookings.package_id").
		Where("bookings.email = ?", email).
		Find(&pkg)

	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Error occured.Please try again.",
		})

		return

	} else {
		c.JSON(200, pkg)
	}
}

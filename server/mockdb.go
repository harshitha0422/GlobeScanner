package main

import "github.com/jinzhu/gorm"

func seed(db *gorm.DB) {
	reg := []Register{
		{Email: "saduvishesha@gmail.com", Password: "vish", Role: "Tourist"},
		{Email: "visheshasadu@gmail.com", Password: "sadu", Role: "Guide"},
		{Email: "vish@gmail.com", Password: "vish", Role: "Guide"},
		{Email: "harshitha@gmail.com", Password: "harsh", Role: "Tourist"},
		{Email: "megh@gmail.com", Password: "megh", Role: "Tourist"},
		{Email: "arijit@gmail.com", Password: "arijit", Role: "Tourist"},
	}
	for _, c := range reg {
		db.Create(&c)
	}

	profiles := []UserProfile{
		{Email: "saduvishesha@gmail.com", Name: "vishesha", About: "travel freak", Age: 23, Fav1: "Gainesville", Fav2: "Orlando", Fav3: "tampa"}, //User: sv, Register: sv_reg
		{Email: "megh@gmail.com", Name: "meghamala", About: "loves travel!", Age: 24, Fav1: "Miami", Fav2: "Gainesville", Fav3: "New york"},      // User: sv, Register: sv_reg
	}
	for _, m := range profiles {
		db.Create(&m)
	}

	profiles2 := []GuideProfile{
		{Email: "visheshasadu@gmail.com", Name: "visheshaSadu", About: "travel freak", Age: 23, Location: "Gainesville", Vehicle: "Ferrari car"}, //User: sv, Register: sv_reg
		{Email: "vish@gmail.com", Name: "vish", About: "loves travel!", Age: 24, Location: "Miami", Vehicle: "Range Rover XL car"},               // User: sv, Register: sv_reg
	}
	for _, pr2 := range profiles2 {
		db.Create(&pr2)
	}

	comments := []Comment{
		{Comment: "Gainesville is peaceful and has nice greenery", Email: "visheshasadu@gmail.com", Location: "Gainesville"},   //User: sv, Register: sv_reg
		{Comment: "Orlando is well developed and has many nice places to visit", Email: "vish@gmail.com", Location: "Orlando"}, // User: sv, Register: sv_reg
	}
	for _, cmnt := range comments {
		db.Create(&cmnt)

	}
	// type Req struct {
	// 	GuideEmail   string `json:"guideEmail" gorm:"not null"`
	// 	Duration     string `json: "duration"`
	// 	Location     string `json:"location"`
	// 	Accomodation string `json:"accomodation"`
	// 	Itinerary    string `json:"itinerary"`
	// 	Included     string `json:"included"`
	// 	Price        string `json:"price"`
	// }
	packages := []Package{
		{GuideEmail: "visheshas@gmail.com", Duration: "2 weeks", Location: "Florida", Accomodation: "Hyatt Regency", Itinerary: "Day1:Miami, Day2: Fort Lauderdale, Day3: Orlando", Included: "Breakfast,Dinner, Sight-seeing", Price: "$100"},
		{GuideEmail: "arijitd@gmail.com", Duration: "1 weeks", Location: "New Delhi", Accomodation: "Hyatt Regency", Itinerary: "Day1:Red Fort, Day2: Fatehpur Sikhri, Day3: Agra", Included: "Breakfast,Dinner, Sight-seeing", Price: "Rs. 10000"},
	}
	for _, pkg := range packages {
		db.Create(&pkg)

	}
}

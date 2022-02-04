// routes.go

package main

func initializeRoutes() {

	// Handle the index route

	r.GET("/searchPlaces/:name", showSearchPlacesPage)
	r.POST("/sign up", userRegister)
	r.POST("/login", userLogin)
	r.GET("/users", getallUsers)
	r.GET("/userprofile", getallTouristprofile)
	r.GET("/guideprofile", getallGuideprofile)
	r.GET("/comments", getallComments)
	r.GET("/users/:email", getUser)
	r.GET("/userprofile/:email", getUserProfile)
	r.GET("/guideprofile/:email", getGuideProfile)
	r.GET("/comments/:location", getLocationComments)
	r.POST("/userprofile", createUserProfile)
	r.POST("/guideprofile", createGuideProfile)
	r.POST("/comments", createComments)
	r.PUT("/userprofile/:email", updateUserProfile)
	r.PUT("/guideprofile/:email", updateGuideProfile)
	r.DELETE("/userprofile/:email", DeleteUserProfile)
	r.DELETE("/guideprofile/:email", DeleteGuideProfile)
}

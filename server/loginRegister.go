package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func userRegister(c *gin.Context) {

	fmt.Print("sign up req")
	// Parse input request
	type Req struct {
		Name     string `json:"name"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6,max=20"`
		Role     string `json:"role"`
	}

	req := Req{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 6 to 20 chars",
		})
		return
	}
	if req.Role == "Tourist" {
		// Check if other user already exists with the same Email
		var userprofile UserProfile
		result := DB.Where("email = ?", req.Email).First(&userprofile)
		if result.RowsAffected == 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "this Email is already taken by other user",
			})
			return
		}
		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "interval server error",
			})
			return
		}
		// Insert into database
		user := UserProfile{
			Name:  req.Name,
			Email: req.Email,
			Role:  req.Role,
		}
		register := Register{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hashedPassword),
			Role:     req.Role,
		}
		result = DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error creating user register in user profile DB",
			})
			return
		}
		result = DB.Create(&register)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error creating user register DB",
			})
			return
		}
		c.JSON(200, user)
	}

	if req.Role == "Guide" {
		// Check if other user already exists with the same Email
		//register := Register{}
		var guideprofile GuideProfile
		result := DB.Where("email = ?", req.Email).First(&guideprofile)
		if result.RowsAffected == 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "this Email is already taken by other user",
			})
			return
		}
		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "interval server error",
			})
			return
		}
		// Insert into database
		user := GuideProfile{
			Name:  req.Name,
			Email: req.Email,
			Role:  req.Role,
		}
		register := Register{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hashedPassword),
			Role:     req.Role,
		}
		result = DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error creating register DB",
			})
			return
		}

		result = DB.Create(&register)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error creating user register DB",
			})
			return
		}
		c.JSON(200, user)
	}

}

func userLogin(c *gin.Context) {
	fmt.Println("user login")
	type Req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6,max=20"`
		Role     string `json:"role"`
	}
	req := Req{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect input parameters",
		})
		return
	}
	// Check if the user exists
	register := Register{}
	res := DB.Where("email = ? AND role = ?", req.Email, req.Role).Find(&register)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found or role does not match",
		})
		return
	}
	// res1 := DB.Where("role = ?", req.Role).First(&register)
	// if res1.Error != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "Role does not match",
	// 	})
	// 	return
	// }
	// Check if the password match
	err = bcrypt.CompareHashAndPassword([]byte(register.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "wrong password",
		})

		return
	}

	ts, err := CreateToken(register.Email, register.Role, register.Name)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := CreateAuth(register.Email, register.Role, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)

	/*c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})*/
}

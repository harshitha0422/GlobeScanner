package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func userRegister(c *gin.Context) {
	// Parse input request
	type Req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=20"`
		Role     string `json:"role"`
	}
	req := Req{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters, password should be between 8 to 20 chars",
		})
		return
	}
	// Check if other user already exists with the same Email
	register := Register{}
	result := DB.Where("email = ?", req.Email).First(&register)
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
	registers := Register{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}
	result = DB.Create(&registers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating register DB",
		})
		return
	}
	c.JSON(http.StatusOK, register)
}

func userLogin(c *gin.Context) {
	type Req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=20"`
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
	res := DB.Where("email = ?", req.Email).First(&register)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}
	// Check if the password match
	err = bcrypt.CompareHashAndPassword([]byte(register.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "wrong password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

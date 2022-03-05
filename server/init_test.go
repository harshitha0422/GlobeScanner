package main

import (
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
r := gin.Default()

func TestMain(m *testing.M) {
	var err error
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatalf("[Failed to initialize test database")
	}
	defer db.Close()
	code := m.Run()
	os.Exit(code)
}

func Testlogin(m *testing.M){

	type Req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=20"`
		Role     string `json:"role"`
	}

	user := Req{
		Username: vishesha,
		Email:    vishesha@gmail.com,
		Password: xyz,
		Role:     Tourist,
	}

	register := Register{}
	res := DB.Where("email = ?", req.Email).First(&register)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			log.Fatalf("user already exists in database")
		})
	}
	return res
}


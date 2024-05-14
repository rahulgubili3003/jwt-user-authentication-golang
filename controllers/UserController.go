package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulgubili3003/jwt-user-authentication-golang/db"
	"github.com/rahulgubili3003/jwt-user-authentication-golang/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Signup(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the request body",
		})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		log.Println("Failed to hash the password received in the request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}
	// Create User
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hashedPassword)}

	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user in the DB",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

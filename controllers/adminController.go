package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"greaterAltitudeapp/config"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"log"
)

type AuthController struct{}

func (a *AuthController) Login(c *gin.Context) {
	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Credentials"})
		return
	}

	var user models.User
	if err := config.H.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid Email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid Password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (a *AuthController) SignUp(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Credentials"})
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Password"})
	}
	newUser.Password = hashedPassword

	if err := config.H.DB.Create(&newUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, gin.H{"message": "User created"})
}

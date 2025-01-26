package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"strings"
)

type AuthController struct{}

func (a *AuthController) Login(c *gin.Context) {
	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	var user models.User
	if err := utils.H.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid Email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid Email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		utils.H.Logger.Printf("Error generating JWT: %v", err)
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (a *AuthController) SignUp(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to hash password"})
	}
	newUser.Password = hashedPassword

	if err := utils.H.DB.Create(&newUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(201, gin.H{"message": "User created"})
}

func (a *AuthController) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Authorization header is required",
		})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Bearer token is required",
		})
		return
	}

	err := utils.InvalidateJWT(token)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Failed to invalidate token",
		})
		return
	}
	c.JSON(200, gin.H{"message": "User logged out sucessfully"})
}

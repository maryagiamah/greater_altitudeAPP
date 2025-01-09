package controllers

import (
        "github.com/gin-gonic/gin"
        "gorm.io/gorm"
        "greaterAltitudeapp/config"
        "greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type AuthController struct{}

func (a *AuthController) Login(c *gin.Context) {
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(400, gin.H{"error": "Invalid Credentials"})
        return
    }

    var user models.User
    if err := config.H.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
        c.JSON(401, gin.H{"error": "Invalid Email"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
        c.JSON(401, gin.H{"error": "Invalid Password"})
        return
    }

    token, err := utils.GenerateJWT(user.ID, user.Role)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(200, gin.H{"token": token})
}

func (a *AuthController) Signup(c *gin.Context) {
    var newUser models.User

    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(400, gin.H{"error": "Invalid Credentials"})
        return
    }

    hashedPassword, err := utils.HashPassword(user.Password)
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

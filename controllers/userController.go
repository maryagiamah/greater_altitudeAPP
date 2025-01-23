package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type UserController struct{}

func (u *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	utils.H.Logger.Printf("Fetched User: %s", user.Email)
	c.JSON(200, gin.H{"user": user})
}

func (u *UserController) GetAllUsers(c *gin.Context) {
	var users []models.User

	if err := utils.H.DB.Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(users) == 0 {
		c.JSON(404, gin.H{"error": "No user found"})
		return
	}
	c.JSON(200, gin.H{"users": users})
}

func (u *UserController) GetUserProfile(c *gin.Context) {
}

func (u *UserController) GetAllStaffs(c *gin.Context) {
}

func (u *UserController) GetAllParents(c *gin.Context) {
}

func (u *UserController) GetAuthenticatedUser(c *gin.Context) {
}

func (u *UserController) CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Password"})
	}

	newUser.Password = hashedPassword
	result := utils.H.DB.Create(&newUser)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create User"})
		return
	}

	utils.H.Logger.Printf("New user Created with ID: %d", newUser.ID)
	c.JSON(201, gin.H{"message": "User created Sucessfully"})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	var updatedFields models.User

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := utils.H.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if updatedFields.Password != "" {
		hashPassword, err := utils.HashPassword(updatedFields.Password)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid Password"})
			return
		}
		updatedFields.Password = hashPassword

	}

	result := utils.H.DB.Model(&user).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update user"})
		utils.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	utils.H.Logger.Printf("Updated user with ID: %d", user.ID)
	c.JSON(200, gin.H{"ID": user.ID})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&user, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		return
	}
	utils.H.Logger.Printf("Deleted user with ID: %s", id)
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (u *UserController) ActivateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	user.IsActive = true
	if err := utils.H.DB.Save(&user).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to activate user"})
		return
	}

	c.JSON(200, gin.H{"message": "User activated successfully"})

}

func (u *UserController) DeactivateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	user.IsActive = false
	if err := utils.H.DB.Save(&user).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to deactivate user"})
		return
	}

	c.JSON(200, gin.H{"message": "User deactivated successfully"})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"net/http"
)

type UserController struct{}

func (u *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u *UserController) GetAllUsers(c *gin.Context) {
	var users []models.User

	if err := utils.H.DB.Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No user found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (u *UserController) GetUserProfile(c *gin.Context) {
}

func (u *UserController) GetAllStaffs(c *gin.Context) {
	var users []models.User

	if err := utils.H.DB.Where("role = ?", "staff").Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No staff user found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (u *UserController) GetAllParents(c *gin.Context) {
	var users []models.User

	if err := utils.H.DB.Where("role = ?", "parent").Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No parent user found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})

}

func (u *UserController) GetAuthenticatedUser(c *gin.Context) {
	id := c.GetUint("userId")
	var user models.User

	if id == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}
	if err := utils.H.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	var updatedFields models.User

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if updatedFields.Password != "" {
		hashPassword, err := utils.HashPassword(updatedFields.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
			return
		}
		updatedFields.Password = hashPassword

	}

	result := utils.H.DB.Model(&user).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":      user.ID,
		"message": "User updated",
	})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&user, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (u *UserController) ActivateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	user.IsActive = true
	if err := utils.H.DB.Save(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to activate user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User activated successfully"})

}

func (u *UserController) DeactivateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	user.IsActive = false
	if err := utils.H.DB.Save(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to deactivate user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deactivated successfully"})
}

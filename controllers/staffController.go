package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/config"
	"greaterAltitudeapp/models"
)

type StaffController struct{}

func (s *StaffController) FetchStaff(c *gin.Context) {
	id := c.Param("id")
	var staff models.Staff

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := config.H.DB.First(&staff, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Staff not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	config.H.Logger.Printf("Fetched Staff: %s %s", staff.FirstName, staff.LastName)
	c.JSON(200, gin.H{"staff": staff})
}

func (s *StaffController) CreateStaff(c *gin.Context) {
	var newStaff models.Staff

	if err := c.ShouldBindJSON(&newStaff); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := config.H.DB.Create(&newStaff)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Staff"})
		return
	}

	config.H.Logger.Printf("New staff Created with ID: %d", newStaff.ID)
	c.JSON(201, gin.H{"ID": newStaff.ID})
}

func (s *StaffController) UpdateStaff(c *gin.Context) {
	id := c.Param("id")
	var staff models.Staff
	var updatedFields models.Staff

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := config.H.DB.First(&staff, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Staff not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := config.H.DB.Model(&staff).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update staff"})
		config.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	config.H.Logger.Printf("Updated staff with ID: %d", staff.ID)
	c.JSON(200, gin.H{"ID": staff.ID})
}

func (s *StaffController) DeleteStaff(c *gin.Context) {
	id := c.Param("id")
	var staff models.Staff

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := config.H.DB.Delete(&staff, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Staff not found"})
		return
	}
	config.H.Logger.Printf("Deleted Staff with ID: %s", id)
	c.JSON(200, gin.H{"message": "Staff deleted successfully"})
}

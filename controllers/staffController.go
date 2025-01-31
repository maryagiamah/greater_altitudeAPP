package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"net/http"
)

type StaffController struct{}

func (s *StaffController) GetStaff(c *gin.Context) {
	id := c.Param("id")
	var staff models.Staff

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&staff, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"staff": staff})
}

func (s *StaffController) GetAllStaffs(c *gin.Context) {
	var staffs []models.Staff

	if err := utils.H.DB.Find(&staffs).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(staffs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No staff found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"staffs": staffs})
}

func (s *StaffController) CreateStaff(c *gin.Context) {
	var newStaff models.Staff

	if err := c.ShouldBindJSON(&newStaff); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := utils.H.DB.Create(&newStaff)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Staff"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID":      newStaff.ID,
		"message": "Staff Created",
	})
}

func (s *StaffController) UpdateStaff(c *gin.Context) {
	id := c.Param("id")
	var staff models.Staff
	var updatedFields models.Staff

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&staff, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&staff).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to update staff"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":      staff.ID,
		"message": "Staff updated",
	})
}

func (s *StaffController) DeleteStaff(c *gin.Context) {
	id := c.Param("id")
	var staff models.Staff

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&staff, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete staff"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Staff deleted successfully"})
}

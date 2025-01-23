package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type StaffController struct{}

func (s *StaffController) GetStaff(c *gin.Context) {
	id := c.Param("id")
	var staff models.Staff

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&staff, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Staff not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	utils.H.Logger.Printf("Fetched Staff: %s %s", staff.User.FirstName, staff.User.LastName)
	c.JSON(200, gin.H{"staff": staff})
}

func (s *StaffController) GetAllStaffs(c *gin.Context) {
	var staffs []models.Staff

        if err := utils.H.DB.Find(&staffs).Error; err != nil {
                c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
                return
        }

        if len(staffs) == 0 {
                c.JSON(404, gin.H{"error": "No staff found"})
                return
        }
        c.JSON(200, gin.H{"staffs": staffs})
}

func (s *StaffController) CreateStaff(c *gin.Context) {
	var newStaff models.Staff

	if err := c.ShouldBindJSON(&newStaff); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := utils.H.DB.Create(&newStaff)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Staff"})
		return
	}

	utils.H.Logger.Printf("New staff Created with ID: %d", newStaff.ID)
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

	if err := utils.H.DB.First(&staff, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Staff not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&staff).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update staff"})
		utils.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	utils.H.Logger.Printf("Updated staff with ID: %d", staff.ID)
	c.JSON(200, gin.H{"ID": staff.ID})
}

func (s *StaffController) DeleteStaff(c *gin.Context) {
	id := c.Param("id")
	var staff models.Staff

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&staff, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Staff not found"})
		return
	}
	utils.H.Logger.Printf("Deleted Staff with ID: %s", id)
	c.JSON(200, gin.H{"message": "Staff deleted successfully"})
}

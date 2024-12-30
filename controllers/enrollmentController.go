package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/config"
	"greaterAltitudeapp/models"
)

type EnrollmentController struct{}

func (e *EnrollmentController) FetchEnrollment(c *gin.Context) {
	id := c.Param("id")
	var enroll models.Enrollment

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := config.H.DB.First(&enroll, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Enrollment not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	config.H.Logger.Printf("Fetched Enrollment: %s %s", enroll.StudentID)
	c.JSON(200, gin.H{"Enrollment": enroll})
}

func (e *EnrollmentController) CreateEnrollment(c *gin.Context) {
	var newEnroll models.Enrollment

	if err := c.ShouldBindJSON(&newEnroll); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := config.H.DB.Create(&newEnroll)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Enroll"})
		return
	}

	config.H.Logger.Printf("New Enroll Created with ID: %d", newEnroll.ID)
	c.JSON(201, gin.H{"ID": newEnroll.ID})
}

func (p *EnrollmentController) UpdateEnrollment(c *gin.Context) {
	id := c.Param("id")
	var enroll models.Enrollment
	var updatedFields models.Enrollment

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := config.H.DB.First(&enroll, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Enrollment not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := config.H.DB.Model(&enroll).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update enrollment"})
		config.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	config.H.Logger.Printf("Updated enrollment with ID: %d", enroll.ID)
	c.JSON(200, gin.H{"ID": enroll.ID})
}

func (e *EnrollmentController) DeleteEnrollment(c *gin.Context) {
	id := c.Param("id")
	var enroll models.Enrollment

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := config.H.DB.Delete(&enroll, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Enrollment not found"})
		return
	}
	config.H.Logger.Printf("Deleted Enrollment with ID: %s", id)
	c.JSON(200, gin.H{"message": "Enrollment deleted successfully"})
}

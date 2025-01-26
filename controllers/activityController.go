package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type ActivityController struct{}

func (a *ActivityController) GetActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&activity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Activity not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(200, gin.H{"activity": activity})
}

func (a *ActivityController) GetAllActivities(c *gin.Context) {
	var activities []models.Activity

	if err := utils.H.DB.Find(&activities).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(activities) == 0 {
		c.JSON(404, gin.H{"error": "No events found"})
		return
	}
	c.JSON(200, gin.H{"activities": activities})
}

func (a *ActivityController) CreateActivity(c *gin.Context) {
	var newActivity models.Activity

	if err := c.ShouldBindJSON(&newActivity); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := utils.H.DB.Create(&newActivity)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Event created successfully",
		"ID":      newActivity.ID,
	})
}

func (a *ActivityController) UpdateActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity
	var updatedFields models.Activity

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&activity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Activity not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&activity).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"ID":      activity.ID,
		"message": "Activity updated successfully",
	})
}

func (a *ActivityController) DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&activity, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Activity not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Activity deleted successfully"})
}

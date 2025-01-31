package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"net/http"
)

type ActivityController struct{}

func (a *ActivityController) GetActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&activity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"activity": activity})
}

func (a *ActivityController) GetAllActivities(c *gin.Context) {
	var activities []models.Activity

	if err := utils.H.DB.Find(&activities).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(activities) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No activities found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"activities": activities})
}

func (a *ActivityController) CreateActivity(c *gin.Context) {
	var newActivity models.Activity

	if err := c.ShouldBindJSON(&newActivity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := utils.H.DB.Create(&newActivity)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to  create activity"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Activity created",
		"ID":      newActivity.ID,
	})
}

func (a *ActivityController) UpdateActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity
	var updatedFields models.Activity

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&activity, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&activity).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":      activity.ID,
		"message": "Activity updated",
	})
}

func (a *ActivityController) DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&activity, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete activity"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted"})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type EventController struct{}

func (e *EventController) GetEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&event, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Event not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(200, gin.H{"event": event})
}

func (e *EventController) GetAllEvents(c *gin.Context) {
	var events []models.Event

	if err := utils.H.DB.Find(&events).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(events) == 0 {
		c.JSON(404, gin.H{"error": "No events found"})
		return
	}
	c.JSON(200, gin.H{"events": events})
}

func (e *EventController) CreateEvent(c *gin.Context) {
	var newEvent models.Event

	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := utils.H.DB.Create(&newEvent)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Event created successfully",
		"ID":      newEvent.ID,
	})
}

func (e *EventController) UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	var updatedFields models.Event

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&event, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Event not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&event).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"ID":      event.ID,
		"message": "Event updated successfully",
	})
}

func (e *EventController) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	var event models.Event

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&event, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Event deleted successfully"})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type MessageController struct{}

func (m *MessageController) GetMessage(c *gin.Context) {
	id := c.Param("id")
	var message models.Message

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&message, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Message not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	utils.H.Logger.Printf("Fetched Message: %s", message.Content)
	c.JSON(200, gin.H{"message": message})
}

func (m *MessageController) CreateMessage(c *gin.Context) {
	var newMessage models.Message

	if err := c.ShouldBindJSON(&newMessage); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := utils.H.DB.Create(&newMessage)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Message"})
		return
	}

	utils.H.Logger.Printf("New Message Created with ID: %d", newMessage.ID)
	c.JSON(201, gin.H{"ID": newMessage.ID})
}

func (m *MessageController) UpdateMessage(c *gin.Context) {
	id := c.Param("id")
	var message models.Message
	var updatedFields models.Message

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := utils.H.DB.First(&message, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Message not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&message).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update message"})
		utils.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	utils.H.Logger.Printf("Updated message with ID: %d", message.ID)
	c.JSON(200, gin.H{"ID": message.ID})
}

func (m *MessageController) DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	var message models.Message

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&message, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Message not found"})
		return
	}
	utils.H.Logger.Printf("Deleted message with ID: %s", id)
	c.JSON(200, gin.H{"message": "Message deleted successfully"})
}

func (m *MessageController) GetInboxMessages(c *gin.Context) {
}

func (m *MessageController) GetSentMessages(c *gin.Context) {
}

func (m *MessageController) MarkMessageAsRead(c *gin.Context) {
}

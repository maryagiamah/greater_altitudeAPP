package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"net/http"
)

type MessageController struct{}

func (m *MessageController) GetInboxMessages(c *gin.Context) {
	userId := c.GetUint("userId")
	var messages []models.Message

	if err := utils.H.DB.Where("receiver_id = ?", userId).Find(&messages).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Failed to fetch messages"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func (m *MessageController) GetSentMessages(c *gin.Context) {
	userId := c.GetUint("userId")
	var messages []models.Message

	if err := utils.H.DB.Where("sender_id = ?", userId).Find(&messages).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func (m *MessageController) CreateMessage(c *gin.Context) {
	var newMessage models.Message

	if err := c.ShouldBindJSON(&newMessage); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	newMessage.SenderID = c.GetUint("userId")

	result := utils.H.DB.Create(&newMessage)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Message sent"})
}

func (m *MessageController) UpdateMessage(c *gin.Context) {
	id := c.Param("id")
	var message models.Message
	var updatedFields models.Message

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&message, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if message.SenderID != c.GetUint("UserId") {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User can't update this message"})
	}

	result := utils.H.DB.Model(&message).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to  update message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message updated",
		"ID":      message.ID,
	})
}

func (m *MessageController) DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	var message models.Message

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&message, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete message"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}

func (m *MessageController) MarkMessageAsRead(c *gin.Context) {
	id := c.Param("id")
	var message models.Message

	if err := utils.H.DB.First(&message, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	message.IsRead = true
	if err := utils.H.DB.Save(&message).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message marked as read"})
}

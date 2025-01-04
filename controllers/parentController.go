package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/config"
	"greaterAltitudeapp/models"
)

type ParentController struct{}

func (p *ParentController) FetchParent(c *gin.Context) {
	id := c.Param("id")
	var parent models.Parent

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := config.H.DB.Preload("User").First(&parent, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Parent not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	config.H.Logger.Printf("Fetched Parent: %s %s", parent.User.FirstName, parent.User.LastName)
	c.JSON(200, gin.H{"parent": parent})

}

func (p *ParentController) CreateParent(c *gin.Context) {
	var newParent models.Parent

	if err := c.ShouldBindJSON(&newParent); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := config.H.DB.Create(&newParent)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Parent"})
		return
	}

	config.H.Logger.Printf("New Parent Created with ID: %d", newParent.ID)
	c.JSON(201, gin.H{"ID": newParent.ID})

}

func (p *ParentController) UpdateParent(c *gin.Context) {
	id := c.Param("id")
	var parent models.Parent
	var updatedFields models.Parent

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := config.H.DB.First(&parent, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Parent not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	result := config.H.DB.Model(&parent).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update parent"})
		config.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	config.H.Logger.Printf("Updated parent with ID: %d", parent.ID)
	c.JSON(200, gin.H{"ID": parent.ID})
}

func (p *ParentController) DeleteParent(c *gin.Context) {
	id := c.Param("id")
	var parent models.Parent

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := config.H.DB.Delete(&parent, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Parent not found"})
		return
	}
	config.H.Logger.Printf("Deleted Parent with ID: %s", id)
	c.JSON(200, gin.H{"message": "Parent deleted successfully"})
}

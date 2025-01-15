package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/config"
	"greaterAltitudeapp/models"
	"log"
)

type PupilController struct{}

func (p *PupilController) FetchPupil(c *gin.Context) {
	id := c.Param("id")
	var pupil models.Pupil

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := config.H.DB.First(&pupil, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Pupil not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	config.H.Logger.Printf("Fetched Pupil: %s %s", pupil.FirstName, pupil.LastName)
	c.JSON(200, gin.H{"pupil": pupil})
}

func (p *PupilController) CreatePupil(c *gin.Context) {
	var newPupil models.Pupil

	if err := c.ShouldBindJSON(&newPupil); err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := config.H.DB.Create(&newPupil)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Pupil"})
		return
	}

	config.H.Logger.Printf("New Pupil Created with ID: %d", newPupil.ID)
	c.JSON(201, gin.H{"ID": newPupil.ID})
}

func (p *PupilController) UpdatePupil(c *gin.Context) {
	id := c.Param("id")
	var pupil models.Pupil
	var updatedFields models.Pupil

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := config.H.DB.First(&pupil, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Pupil not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := config.H.DB.Model(&pupil).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update pupil"})
		config.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	config.H.Logger.Printf("Updated pupil with ID: %d", pupil.ID)
	c.JSON(200, gin.H{"ID": pupil.ID})
}

func (p *PupilController) DeletePupil(c *gin.Context) {
	id := c.Param("id")
	var pupil models.Pupil

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := config.H.DB.Delete(&pupil, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Pupil not found"})
		return
	}
	config.H.Logger.Printf("Deleted Pupil with ID: %s", id)
	c.JSON(200, gin.H{"message": "Pupil deleted successfully"})
}

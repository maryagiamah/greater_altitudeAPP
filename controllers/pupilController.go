package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"log"
	"net/http"
)

type PupilController struct{}

func (p *PupilController) GetPupil(c *gin.Context) {
	id := c.Param("id")
	var pupil models.Pupil

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&pupil, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Pupil not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"pupil": pupil})
}

func (p *PupilController) GetAllPupils(c *gin.Context) {
	var pupils []models.Pupil

	if err := utils.H.DB.Find(&pupils).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(pupils) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No pupil found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"pupils": pupils})
}

func (p *PupilController) CreatePupil(c *gin.Context) {
	var newPupil models.Pupil

	if err := c.ShouldBindJSON(&newPupil); err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := utils.H.DB.Create(&newPupil)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Pupil"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID":      newPupil.ID,
		"message": "Pupil created",
	})
}

func (p *PupilController) UpdatePupil(c *gin.Context) {
	id := c.Param("id")
	var pupil models.Pupil
	var updatedFields models.Pupil

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&pupil, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Pupil not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&pupil).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to update pupil"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":      pupil.ID,
		"message": "Pupil updated",
	})
}

func (p *PupilController) DeletePupil(c *gin.Context) {
	id := c.Param("id")
	var pupil models.Pupil

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&pupil, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pupil"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Pupil not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pupil deleted successfully"})
}

func (p *PupilController) GetAllClasses(c *gin.Context) {
}

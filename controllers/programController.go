package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/config"
	"greaterAltitudeapp/models"
)

type ProgramController struct{}

func (p *ProgramController) FetchProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := config.H.DB.First(&program, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Program not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	config.H.Logger.Printf("Fetched Program: %s %s", program.Name)
	c.JSON(200, gin.H{"program": program})
}

func (p *ProgramController) CreateProgram(c *gin.Context) {
	var newProgram models.Program

	if err := c.ShouldBindJSON(&newProgram); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := config.H.DB.Create(&newProgram)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Program"})
		return
	}

	config.H.Logger.Printf("New Pupil Created with ID: %d", newProgram.ID)
	c.JSON(201, gin.H{"ID": newProgram.ID})
}

func (p *ProgramController) UpdateProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program
	var updatedFields models.Program

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := config.H.DB.First(&program, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Program not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := config.H.DB.Model(&program).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update program"})
		config.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	config.H.Logger.Printf("Updated program with ID: %d", program.ID)
	c.JSON(200, gin.H{"ID": program.ID})
}

func (p *ProgramController) DeleteProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := config.H.DB.Delete(&program, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Program not found"})
		return
	}
	config.H.Logger.Printf("Deleted Program with ID: %s", id)
	c.JSON(200, gin.H{"message": "Program deleted successfully"})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type ProgramController struct{}

func (p *ProgramController) GetProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&program, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Program not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	utils.H.Logger.Printf("Fetched Program: %s %s", program.Name)
	c.JSON(200, gin.H{"program": program})
}

func (p *ProgramController) GetAllPrograms(c *gin.Context) {
	var programs []models.Program

        if err := utils.H.DB.Find(&programs).Error; err != nil {
                c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
                return
        }

        if len(programs) == 0 {
                c.JSON(404, gin.H{"error": "No program found"})
                return
        }
        c.JSON(200, gin.H{"programs": programs})
}

func (p *ProgramController) CreateProgram(c *gin.Context) {
	var newProgram models.Program

	if err := c.ShouldBindJSON(&newProgram); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := utils.H.DB.Create(&newProgram)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Program"})
		return
	}

	utils.H.Logger.Printf("New Program Created with ID: %d", newProgram.ID)
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

	if err := utils.H.DB.First(&program, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Program not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&program).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update program"})
		utils.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	utils.H.Logger.Printf("Updated program with ID: %d", program.ID)
	c.JSON(200, gin.H{"ID": program.ID})
}

func (p *ProgramController) DeleteProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&program, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Program not found"})
		return
	}
	utils.H.Logger.Printf("Deleted Program with ID: %s", id)
	c.JSON(200, gin.H{"message": "Program deleted successfully"})
}

func (p *ProgramController) GetProgramClasses(c *gin.Context) {
}

func (p *ProgramController) GetProgramActivities(c *gin.Context) {
}

func (p *ProgramController) AddClassToProgram(c *gin.Context) {
}

func (p *ProgramController) AddActivityToProgram (c *gin.Context) {
}

func (p *ProgramController) DeleteActivity (c *gin.Context) {
}

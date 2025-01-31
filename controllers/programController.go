package controllers

import (
	"net/http"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&program, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"program": program,
		"message": "Program created",
	})
}

func (p *ProgramController) GetAllPrograms(c *gin.Context) {
	var programs []models.Program

	if err := utils.H.DB.Find(&programs).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(programs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No program found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"programs": programs})
}

func (p *ProgramController) CreateProgram(c *gin.Context) {
	var newProgram models.Program

	if err := c.ShouldBindJSON(&newProgram); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := utils.H.DB.Create(&newProgram)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Program"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID":      newProgram.ID,
		"message": "Program created",
	})
}

func (p *ProgramController) UpdateProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program
	var updatedFields models.Program

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&program, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&program).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to  update program"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":      program.ID,
		"message": "Program updated",
	})
}

func (p *ProgramController) DeleteProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&program, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Program deleted successfully"})
}

func (p *ProgramController) GetProgramClasses(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.Preload("Classes").First(&program, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"program_classes": program.Classes})
}

func (p *ProgramController) GetProgramActivities(c *gin.Context) {
	id := c.Param("id")
	var program models.Program

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.Preload("Activities").First(&program, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
	}
	c.JSON(http.StatusOK, gin.H{"program_activities": program.Activities})
}

func (p *ProgramController) AddClassToProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program
	var newClass models.Class

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&program, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if err := c.ShouldBindJSON(&newClass); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.Association("Classes").Append(newClass).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add class to program"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class succesfully added to program"})
}

func (p *ProgramController) AddActivityToProgram(c *gin.Context) {
	id := c.Param("id")
	var program models.Program
	var newActivity models.Activity

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&program, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if err := c.ShouldBindJSON(&newActivity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.Model(&program).Association("Activities").Append(newActivity).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add activity to program"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity succesfully added to program"})
}

func (p *ProgramController) DeleteActivity(c *gin.Context) {
}

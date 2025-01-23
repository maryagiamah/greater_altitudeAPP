package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type ClassController struct{}

func (cl *ClassController) GetClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Class not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	utils.H.Logger.Printf("Fetched Class: %s %s", class.Name)
	c.JSON(200, gin.H{"class": class})
}

func (cl *ClassController) CreateClass(c *gin.Context) {
	var newClass models.Class

	if err := c.ShouldBindJSON(&newClass); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := utils.H.DB.Create(&newClass)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Class"})
		return
	}

	utils.H.Logger.Printf("New Class Created with ID: %d", newClass.ID)
	c.JSON(201, gin.H{"ID": newClass.ID})
}

func (cl *ClassController) UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	var updatedFields models.Class

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := utils.H.DB.First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Class not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&class).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update class"})
		return
	}

	utils.H.Logger.Printf("Updated class with ID: %d", class.ID)
	c.JSON(200, gin.H{"ID": class.ID})
}

func (cl *ClassController) DeleteClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&class, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Class  not found"})
		return
	}
	utils.H.Logger.Printf("Deleted Class with ID: %s", id)
	c.JSON(200, gin.H{"message": "Class deleted successfully"})
}

func (cl *ClassController) GetAllClasses(c *gin.Context) {
	var classes []models.Class

	if err := utils.H.DB.Find(&classes).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(classes) == 0 {
		c.JSON(404, gin.H{"error": "No class found"})
		return
	}
	c.JSON(200, gin.H{"classes": classes})
}

func (cl *ClassController) AddPupilToClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	var newPupil models.Pupil

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Class not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if err := c.ShouldBindJSON(&newPupil); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := utils.H.DB.Association("Pupils").Append(newPupil).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to add pupil to class"})
	}

	c.JSON(200, gin.H{"message": "Pupil succesfully added to class"})
}

func (cl *ClassController) AssignTeacherToClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	var newTeacher models.Staff

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Class not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if err := c.ShouldBindJSON(&newTeacher); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := utils.H.DB.Association("Teachers").Append(newTeacher).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to add teacher to class"})
	}

	c.JSON(200, gin.H{"message": "Teacher succesfully added to class"})
}

func (cl *ClassController) GetPupilsInClass(c *gin.Context) {
}

func (cl *ClassController) GetTeachersInClass(c *gin.Context) {
}

func (cl *ClassController) GetClassActivities(c *gin.Context) {
}

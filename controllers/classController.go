package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"net/http"
)

type ClassController struct{}

func (cl *ClassController) GetClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"class": class})
}

func (cl *ClassController) CreateClass(c *gin.Context) {
	var newClass models.Class

	if err := c.ShouldBindJSON(&newClass); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := utils.H.DB.Create(&newClass)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to create class"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Class created",
		"ID":      newClass.ID,
	})
}

func (cl *ClassController) UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	var updatedFields models.Class

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&class).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Fialed to update class"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Class updated",
		"ID":      class.ID,
	})
}

func (cl *ClassController) DeleteClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&class, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete class"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Class  not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}

func (cl *ClassController) GetAllClasses(c *gin.Context) {
	var classes []models.Class

	if err := utils.H.DB.Find(&classes).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(classes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No classes found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"classes": classes})
}

func (cl *ClassController) AddPupilToClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	var newPupil models.Pupil

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if err := c.ShouldBindJSON(&newPupil); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.Model(&class).Association("Pupils").Append(newPupil).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to  add pupil to class"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pupil succesfully added to class"})
}

func (cl *ClassController) AssignTeacherToClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	var newTeacher models.Staff

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if err := c.ShouldBindJSON(&newTeacher); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.Model(&class).Association("Teachers").Append(newTeacher).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add teacher to class"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher succesfully added to class"})
}

func (cl *ClassController) GetPupilsInClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.Preload("Pupils").First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pupils to class"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"class_pupils": class.Pupils})
}

func (cl *ClassController) GetTeachersInClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.Preload("Teachers").First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get teachers in class"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"class_teachers": class.Teachers})
}

func (cl *ClassController) GetClassActivities(c *gin.Context) {
}

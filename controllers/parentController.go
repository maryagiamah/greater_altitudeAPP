package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"net/http"
)

type ParentController struct{}

func (p *ParentController) GetParent(c *gin.Context) {
	id := c.Param("id")
	var parent models.Parent

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.Preload("User").First(&parent, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"parent": parent})

}

func (p *ParentController) GetAllParents(c *gin.Context) {
	var parents []models.Parent

	if err := utils.H.DB.Find(&parents).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(parents) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No parent found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"parents": parents})
}

func (p *ParentController) GetPupilsByParent(c *gin.Context) {
	id := c.Param("id")
	var parent models.Parent

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.Preload("Ward").First(&parent, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"parent_wards": parent.Ward})
}

func (p *ParentController) AddPupilToParent(c *gin.Context) {
	id := c.Param("id")
	var parent models.Parent
	var newPupil models.Pupil

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&parent, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if err := c.ShouldBindJSON(&newPupil); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.Model(&parent).Association("Pupils").Append(newPupil).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add pupil to parent"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pupil succesfully added to parent"})
}

func (p *ParentController) CreateParent(c *gin.Context) {
	var newParent models.Parent
	var user models.User

	if err := c.ShouldBindJSON(&newParent); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&user, newParent.UserID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User row not found"})
		return
	}

	result := utils.H.DB.Create(&newParent)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Parent"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID":      newParent.ID,
		"message": "Parent created",
	})

}

func (p *ParentController) UpdateParent(c *gin.Context) {
	id := c.Param("id")
	var parent models.Parent
	var updatedFields models.Parent

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&parent, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	result := utils.H.DB.Model(&parent).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update parent"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":      parent.ID,
		"message": "Parent updated",
	})
}

func (p *ParentController) DeleteParent(c *gin.Context) {
	id := c.Param("id")
	var parent models.Parent

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&parent, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete parent"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Parent deleted successfully"})
}

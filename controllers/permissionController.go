package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"net/http"
)

type PermissionController struct{}

func (p *PermissionController) GetPermission(c *gin.Context) {
	id := c.Param("id")
	var permission models.Permission

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&permission, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"permission": permission})
}

func (p *PermissionController) GetPermissions(c *gin.Context) {
	var permissions []models.Permission

	if err := utils.H.DB.Find(&permissions).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(permissions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No permissions found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"permissions": permissions})
}

func (p *PermissionController) CreatePermission(c *gin.Context) {
	var newPermission models.Permission

	if err := c.ShouldBindJSON(&newPermission); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := utils.H.DB.Create(&newPermission)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to create permission"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID":      newPermission.ID,
		"message": "Permission created",
	})
}

func (p *PermissionController) UpdatePermission(c *gin.Context) {
	id := c.Param("id")
	var permission models.Permission
	var updatedFields models.Permission

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&permission, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&permission).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update permission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":      permission.ID,
		"message": "Permission updated",
	})
}

func (p *PermissionController) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	var permission models.Permission

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&permission, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete permission"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Permission deleted successfully"})
}

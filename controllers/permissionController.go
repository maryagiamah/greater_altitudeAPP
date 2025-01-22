package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type PermissionController struct{}

func (p *PermissionController) GetPermission(c *gin.Context) {
	id := c.Param("id")
	var permission models.Permission

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&permission, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Permission not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	utils.H.Logger.Printf("Fetched Permission: %s", permission.Name)
	c.JSON(200, gin.H{"permission": permission})
}

func (p *PermissionController) GetPermissions(c *gin.Context) {
	var permissions []models.Permission

	if err := utils.H.DB.Find(&permissions).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(permissions) == 0 {
		c.JSON(404, gin.H{"error": "No permissions found"})
		return
	}
	c.JSON(200, gin.H{"permissions": permissions})
}

func (p *PermissionController) CreatePermission(c *gin.Context) {
	var newPermission models.Permission

	if err := c.ShouldBindJSON(&newPermission); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := utils.H.DB.Create(&newPermission)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create permission"})
		return
	}

	utils.H.Logger.Printf("New Permission Created with ID: %d", newPermission.ID)
	c.JSON(201, gin.H{"ID": newPermission.ID})
}

func (p *PermissionController) UpdatePermission(c *gin.Context) {
	id := c.Param("id")
	var permission models.Permission
	var updatedFields models.Permission

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := utils.H.DB.First(&permission, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Permission not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&permission).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update permission"})
		utils.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	utils.H.Logger.Printf("Updated permission with ID: %d", permission.ID)
	c.JSON(200, gin.H{"ID": permission.ID})
}

func (p *PermissionController) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	var permission models.Permission

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&permission, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Permission not found"})
		return
	}
	utils.H.Logger.Printf("Deleted Permission with ID: %s", id)
	c.JSON(200, gin.H{"message": "Permission deleted successfully"})
}

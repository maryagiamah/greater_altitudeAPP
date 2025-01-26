package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type RoleController struct{}

func (r *RoleController) GetRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&role, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Role not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	utils.H.Logger.Printf("Fetched Role: %s %s", role.Name)
	c.JSON(200, gin.H{"role": role})
}

func (r *RoleController) GetRoles(c *gin.Context) {
	var roles []models.Role

	if err := utils.H.DB.Find(&roles).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(roles) == 0 {
		c.JSON(404, gin.H{"error": "No roles found"})
		return
	}
	c.JSON(200, gin.H{"roles": roles})
}

func (r *RoleController) CreateRole(c *gin.Context) {
	var newRole models.Role

	if err := c.ShouldBindJSON(&newRole); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := utils.H.DB.Create(&newRole)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create role"})
		return
	}

	utils.H.Logger.Printf("New Role Created with ID: %d", newRole.ID)
	c.JSON(201, gin.H{"ID": newRole.ID})
}

func (r *RoleController) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	var updatedFields models.Role

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := utils.H.DB.First(&role, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Role not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&role).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update role"})
		utils.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	utils.H.Logger.Printf("Updated role with ID: %d", role.ID)
	c.JSON(200, gin.H{"ID": role.ID})
}

func (r *RoleController) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&role, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Role not found"})
		return
	}
	utils.H.Logger.Printf("Deleted Role with ID: %s", id)
	c.JSON(200, gin.H{"message": "Role deleted successfully"})
}

func (r *RoleController) UpdateRolePermissions(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	var newPermission models.Permission

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}
	if err := utils.H.DB.Preload("Permissions").First(&role, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Role not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if err := c.ShouldBindJSON(&newPermission); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}
	if err := utils.H.DB.Model(&role).Association("Permissions").Append(newPermission).Error; err != nil {
		utils.H.Logger.Print(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to add permission to role"})
		return
	}

	c.JSON(200, gin.H{"message": "Permission succesfully added to role"})
}

func (r *RoleController) GetPermissionsInRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}
	if err := utils.H.DB.Preload("Permissions").First(&role, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Role not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(200, gin.H{"role_permissions": role.Permissions})
}

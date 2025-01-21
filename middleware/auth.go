package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/config"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("userId", (*claims)["userId"])
		c.Set("role", (*claims)["role"])
		c.Next()
	}
}

func HasPermission(role, allowedPerms ...string) bool {
	var checkRole models.Role

	if err := config.H.DB.Where("name = ?", role).Preload("Permissions").First(&checkRole).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			config.H.Logger.Fatal("Role hasn't been created")
		}
		return false
	}

	if len(permissionNames) == 0 {
		config.H.Logger.Print(err)
		config.H.Logger.Fatal("Role has no permissions")
		return false
	}

	var permissionNames []string
	for i, perm := range checkRole.Permissions {
		permissionNames[i] = perm.Name
	}

	for _, perm := range allowedPerms {
		if !slices.Contains(permissionNames, perm) {
			return false
		}
	}

	return true
}

func RoleMiddleware(allowedPerms ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if HasPermission(role, allowedPerms) {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
	}
}

package middleware

import (
	"errors"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Authorization header is required",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Bearer token is required",
			})
			return
		}
		claims, err := utils.ValidateJWT(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Token Invalid or Expired",
			})
			return
		}

		c.Set("userId", (*claims)["userId"])
		c.Set("role", (*claims)["role"])
		c.Next()
	}
}

func HasPermission(role string, allowedPerms ...string) bool {
	var checkRole models.Role

	if err := utils.H.DB.Where("name = ?", role).Preload("Permissions").First(&checkRole).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.H.Logger.Print("Role hasn't been created")
		} else {
			utils.H.Logger.Printf("Error fetching role: %v", err)
		}
		return false
	}

	if len(checkRole.Permissions) == 0 {
		utils.H.Logger.Print("Role has no permissions")
		return false
	}

	for _, perm := range checkRole.Permissions {
		if slices.Contains(allowedPerms, perm.Name) {
			return true
		}
	}

	return false
}

func RoleMiddleware(allowedPerms ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role == "" {
			utils.H.Logger.Printf("No role found in context")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if !HasPermission(role, allowedPerms...) {
			utils.H.Logger.Printf("Access denied for role: %s", role)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

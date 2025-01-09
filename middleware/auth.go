package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
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

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        role := c.GetString("role")
        for _, allowedRole := range allowedRoles {
            if role == allowedRole {
                c.Next()
                return
            }
        }
        c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
    }
}

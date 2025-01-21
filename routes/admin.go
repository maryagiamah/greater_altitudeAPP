package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterAdminServices(r *gin.Engine) {
	admin := r.Group("/admin", middleware.AuthMiddleware(), middleware.RoleMiddleware("superUser"))
	{
		admin.POST("/login", Login)
		admin.POST("/signup", CreateUser)
		admin.PUT("/users/:id/activate", ActivateUser)
		admin.PUT("/users/:id/deactivate", DeactivateUser)

		admin.POST("/roles", CreateRole)
		admin.GET("/roles", GetRoles)
		admin.PUT("/roles/:id", UpdateRole)
		admin.PUT("/roles/:id/permissions", UpdateRolePermissions)
		admin.GET("/permissions", GetPermissions)
	}
}

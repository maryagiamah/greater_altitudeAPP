package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterAdminServices(rg *gin.RouterGroup) {
	authController := &controllers.AuthController{}
	userController := &controllers.UserController{}
	roleController := &controllers.RoleController{}
	permissionController := &controllers.PermissionController{}

	admin := rg.Group("/admin")
	{
		admin.POST("/login", authController.Login)
		admin.DELETE("/logout", middleware.AuthMiddleware(), authController.Logout)
		admin.POST("/signup", middleware.AuthMiddleware(), authController.SignUp)
		admin.PUT("/users/:id/activate", middleware.AuthMiddleware(), userController.ActivateUser)
		admin.PUT("/users/:id/deactivate", middleware.AuthMiddleware(), userController.DeactivateUser)

		admin.GET("/roles/:id", roleController.GetRole)
		admin.GET("/roles/", roleController.GetRoles)
		admin.POST("/roles", roleController.CreateRole)
		admin.PUT("/roles/:id", roleController.UpdateRole)
		admin.DELETE("/roles/:id", roleController.DeleteRole)
		admin.GET("roles/:id/permissions", roleController.GetPermissionsInRole)
		admin.PUT("/roles/:id/permissions", roleController.UpdateRolePermissions)

		admin.GET("/permissions", permissionController.GetPermissions)
		admin.GET("/permissions/:id", permissionController.GetPermission)
		admin.POST("/permissions", permissionController.CreatePermission)
		admin.PUT("/permissions/:id", permissionController.UpdatePermission)
		admin.DELETE("/permissions/:id", permissionController.DeletePermission)
	}
}

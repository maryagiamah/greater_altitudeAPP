package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
)

func RegisterAdminServices(rg *gin.RouterGroup) {
	authController := &controllers.AuthController{}
	userController := &controllers.UserController{}
	roleController := &controllers.RoleController{}
	permissionController := &controllers.PermissionController{}

	admin := rg.Group("/admin")
	{
		admin.POST("/login", authController.Login)
		admin.POST("/signup", authController.SignUp)
		admin.PUT("/users/:id/activate", userController.ActivateUser)
		admin.PUT("/users/:id/deactivate", userController.DeactivateUser)

		admin.POST("/roles", roleController.CreateRole)
		admin.GET("/roles", roleController.GetRoles)
		admin.PUT("/roles/:id", roleController.UpdateRole)
		admin.DELETE("/roles/:id", roleController.DeleteRole)
		admin.PUT("/roles/:id/permissions", roleController.UpdateRolePermissions)
		admin.GET("/permissions", permissionController.GetPermissions)
	}
}

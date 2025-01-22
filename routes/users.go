package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterUserServices(rg *gin.RouterGroup) {
	userController := &controllers.UserController{}
	parentController := &controllers.ParentController{}
	staffController := &controllers.StaffController{}

	user := rg.Group("/users", middleware.AuthMiddleware())
	{
		user.GET("/", userController.GetAllUsers)
		user.GET("/:id", userController.GetUser)
		user.PUT("/:id", userController.UpdateUser)
		user.GET("/:id/profile", userController.GetUserProfile)
		user.GET("/staffs", userController.GetAllStaffs)
		user.GET("/parents", userController.GetAllParents)

		user.GET("/me", userController.GetAuthenticatedUser)
	}

	parent := rg.Group("/parents", middleware.AuthMiddleware())
	{
		parent.GET("/", parentController.GetAllParents)
		parent.GET("/:id", parentController.GetParent)
		parent.POST("/", parentController.CreateParent)
		parent.PUT("/:id", parentController.UpdateParent)
		parent.DELETE("/:id", parentController.DeleteParent)

		parent.GET("/:id/wards", parentController.GetPupilsByParent)
		parent.POST("/:id/ward", parentController.AddPupilToParent)
	}

	staff := rg.Group("/staff")
	{
		staff.POST("/", staffController.CreateStaff)
		staff.GET("/", staffController.GetAllStaffs)
		staff.GET("/:id", staffController.GetStaff)
		staff.PUT("/:id", staffController.UpdateStaff)
		staff.DELETE("/:id", staffController.DeleteStaff)
	}

}

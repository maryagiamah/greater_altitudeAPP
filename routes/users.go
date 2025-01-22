package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterUserServices(rg *gin.RouterGroup) {
	user := rg.Group("/users", middleware.AuthMiddleware())
	{
		user.GET("/", userController.GetAllUsers)
		user.GET("/:id", userController.FetchUser)
		user.PUT("/:id", userController.UpdateUser)
		user.GET("/:id/profile", getUserProfile)
		user.GET("/staffs", getAllStaff)
		user.GET("/parents", getAllParents)

		user.GET("/me", userController.GetAuthenticatedUser)
	}

	parent := rg.Group("/parents", middleware.AdminMiddleware())
	{
		parent.GET("/", GetAllParents)
		parent.GET("/:id", GetParentDetails)
		parent.POST("/", CreateParent)
		parent.PUT("/:id", UpdateParent)
		parent.DELETE("/:id", DeleteParent)

		parent.GET("/:id/wards", GetPupilsByParent)
		parent.POST("/:id/ward", AddPupilToParent)
	}

	staff := rg.Group("/staff")
	{
		staff.POST("/", createStaff)
		staff.GET("/", listStaff)
		staff.GET("/:id", getStaff)
		staff.PUT("/:id", updateStaff)
		staff.DELETE("/:id", deleteStaff)
	}

}

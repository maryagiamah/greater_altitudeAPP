package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterUserServices(r *gin.Engine) {
	user := r.Group("/users", middleware.AuthMiddleware())
	{
		user.GET("/", userController.GetAllUsers)
		user.GET("/:id", userController.FetchUser)
		user.PUT("/:id", userController.UpdateUser)
		user.GET("/:id/profile", getUserProfile)
		user.GET("/staffs", getAllStaff)
		user.GET("/parents", getAllParents)

		user.GET("/me", userController.GetAuthenticatedUser)
	}

	parent := r.Group("/parents", middleware.AdminMiddleware())
	{
		parents.GET("/", GetAllParents)
		parents.GET("/:id", GetParentDetails)
		parents.POST("/", CreateParent)
		parents.PUT("/:id", UpdateParent)
		parents.DELETE("/:id", DeleteParent)

		parents.GET("/:id/wards", GetPupilsByParent)
		parents.POST("/:id/ward", AddPupilToParent)
	}

	staff := router.Group("/staff")
	{
		staff.POST("/", createStaff)
		staff.GET("/", listStaff)
		staff.GET("/:id", getStaff)
		staff.PUT("/:id", updateStaff)
		staff.DELETE("/:id", deleteStaff)
	}

}

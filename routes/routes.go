package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	pupilController := &controllers.PupilController{}
	parentController := &controllers.ParentController{}
	staffController := &controllers.StaffController{}
	programController := &controllers.ProgramController{}
	enrollmentController := &controllers.EnrollmentController{}
	classController := &controllers.ClassController{}
	userController := &controllers.UserController{}
	authController := &controllers.AuthController{}

	v1 := r.Group("api/v1", middleware.AuthMiddleware())
	{
		v1.GET("/user/:id", middleware.RoleMiddleware("admin"), userController.FetchUser)
		v1.GET("/staff/:id", middleware.RoleMiddleware("admin", "staff"), staffController.FetchStaff)
		v1.GET("/parent/:id", middleware.RoleMiddleware("admin", "staff"), parentController.FetchParent)
		v1.GET("/pupil/:id", middleware.RoleMiddleware("admin", "staff"), pupilController.FetchPupil)
		v1.GET("/program/:id", middleware.RoleMiddleware("admin", "parent"), programController.FetchProgram)
		v1.GET("/class/:id", middleware.RoleMiddleware("admin", "staff"), classController.FetchClass)
		v1.GET("/enrollment/:id", middleware.RoleMiddleware("admin"), enrollmentController.FetchEnrollment)

		v1.POST("/user", middleware.RoleMiddleware("admin"), userController.CreateUser)
		v1.POST("/program", middleware.RoleMiddleware("admin"), programController.CreateProgram)
		v1.POST("/class", middleware.RoleMiddleware("admin"), classController.CreateClass)
		v1.POST("/enrollment", middleware.RoleMiddleware("admin"), enrollmentController.CreateEnrollment)
		v1.POST("/staff", middleware.RoleMiddleware("admin"), staffController.CreateStaff)
		v1.POST("/parent", middleware.RoleMiddleware("admin"), parentController.CreateParent)
		v1.POST("/pupil", middleware.RoleMiddleware("admin"), pupilController.CreatePupil)

		v1.PUT("/user/:id", middleware.RoleMiddleware("admin", "parent", "staff"), userController.UpdateUser)
		v1.PUT("/program/:id", middleware.RoleMiddleware("admin"), programController.UpdateProgram)
		v1.PUT("/class/:id", middleware.RoleMiddleware("admin"), classController.UpdateClass)
		v1.PUT("/enrollment/:id", middleware.RoleMiddleware("admin"), enrollmentController.UpdateEnrollment)
		v1.PUT("/staff/:id", middleware.RoleMiddleware("admin"), staffController.UpdateStaff)
		v1.PUT("/parent/:id", middleware.RoleMiddleware("admin"), parentController.UpdateParent)
		v1.PUT("/pupil/:id", middleware.RoleMiddleware("admin"), pupilController.UpdatePupil)

		v1.DELETE("/user/:id", middleware.RoleMiddleware("admin"), userController.DeleteUser)
		v1.DELETE("/program/:id", middleware.RoleMiddleware("admin"), programController.DeleteProgram)
		v1.DELETE("/enrollment/:id", middleware.RoleMiddleware("admin"), enrollmentController.DeleteEnrollment)
		v1.DELETE("/class/:id", middleware.RoleMiddleware("admin"), classController.DeleteClass)
		v1.DELETE("/staff/:id", middleware.RoleMiddleware("admin"), staffController.DeleteStaff)
		v1.DELETE("/parent/:id", middleware.RoleMiddleware("admin"), parentController.DeleteParent)
		v1.DELETE("/pupil/:id", middleware.RoleMiddleware("admin"), pupilController.DeletePupil)
	}

	auth := r.Group("auth")
	{
		auth.POST("/login", authController.Login)
	}
}

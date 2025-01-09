package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	pupilController := &controllers.PupilController{}
	parentController := &controllers.ParentController{}
	staffController := &controllers.StaffController{}
	programController := &controllers.ProgramController{}
	enrollmentController := &controllers.EnrollmentController{}
	classController := &controllers.ClassController{}

	v1 := r.Group("api/v1", middleware.AuthMiddleware())
	{
		v1.GET("/staff/:id", staffController.FetchStaff)
		v1.GET("/parent/:id", parentController.FetchParent)
		v1.GET("/pupil/:id", pupilController.FetchPupil)
		v1.GET("/program/:id", programController.FetchProgram)
		v1.GET("/class/:id", classController.FetchClass)
		v1.GET("/enrollment/:id", enrollmentController.FetchEnrollment)

		v1.POST("/program", middleware.RoleMiddleware("admin"), programController.CreateProgram)
		v1.POST("/class", middleware.RoleMiddleware("admin"), classController.CreateClass)
		v1.POST("/enrollment", middleware.RoleMiddleware("admin"), enrollmentController.CreateEnrollment)
		v1.POST("/staff", middleware.RoleMiddleware("admin"), staffController.CreateStaff)
		v1.POST("/parent", middleware.RoleMiddleware("admin"), parentController.CreateParent)
		v1.POST("/pupil", middleware.RoleMiddleware("admin"), pupilController.CreatePupil)

		v1.PUT("/program/:id", programController.UpdateProgram)
		v1.PUT("/class/:id", classController.UpdateClass)
		v1.PUT("/enrollment/:id", enrollmentController.UpdateEnrollment)
		v1.PUT("/staff/:id", staffController.UpdateStaff)
		v1.PUT("/parent/:id", parentController.UpdateParent)
		v1.PUT("/pupil/:id", pupilController.UpdatePupil)

		v1.DELETE("/program/:id", programController.DeleteProgram)
		v1.DELETE("/enrollment/:id", enrollmentController.DeleteEnrollment)
		v1.DELETE("/class/:id", classController.DeleteClass)
		v1.DELETE("/staff/:id", staffController.DeleteStaff)
		v1.DELETE("/parent/:id", parentController.DeleteParent)
		v1.DELETE("/pupil/:id", pupilController.DeletePupil)
	}

	/*
	   auth := r.Group("auth")
	   {
	   }
	*/
}

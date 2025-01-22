package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterClassServices(rg *gin.RouterGroup) {

	classController := &controllers.ClassController{}
	pupilController := &controllers.PupilController{}

	class := rg.Group("/classes", middleware.AuthMiddleware())
	{
		class.POST("/", classController.CreateClass)
		class.GET("/", classController.GetAllClasses)
		class.GET("/:id", classController.GetClass)
		class.PUT("/:id", classController.UpdateClass)
		class.DELETE("/:id", classController.DeleteClass)

		class.POST("/:id/pupil", classController.AddPupilToClass)
		class.POST("/:id/teacher", classController.AssignTeacherToClass)
		class.GET("/:id/pupils", classController.GetPupilsInClass)
		class.GET("/:id/teachers", classController.GetTeachersInClass)
		class.GET("/:id/activities", classController.GetClassActivities)
	}

	pupil := rg.Group("/pupils")
	{
		pupil.POST("/", pupilController.CreatePupil)
		pupil.GET("/", pupilController.GetAllPupils)
		pupil.GET("/:id", pupilController.GetPupil)
		pupil.PUT("/:id", pupilController.UpdatePupil)
		pupil.DELETE("/:id", pupilController.DeletePupil)

		pupil.GET("/:id/classes", pupilController.GetAllClasses)
	}

}

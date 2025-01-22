package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterClassServices(rg *gin.RouterGroup) {

	classController := &controllers.ClassController{}

	class := rg.Group("/classes", middleware.AdminMiddleware())
	{
		class.POST("/", CreateClass)
		class.GET("/", GetAllClasses)
		class.GET("/:id", GetClassDetails)
		class.PUT("/:id", UpdateClass)
		class.DELETE("/:id", DeleteClass)

		class.POST("/:id/pupil", AddPupilToClass)
		class.POST("/:id/teacher", AssignTeacherToClass)
		class.GET("/:id/pupils", GetPupilsInClass)
		class.GET("/:id/teachers", GetTeachersInClass)
		class.GET("/:id/activities", getClassActivities)
	}

	pupil := rg.Group("/pupils")
	{
		pupil.POST("/", createPupil)
		pupil.GET("/", listPupils)
		pupil.GET("/:id", getPupil)
		pupil.PUT("/:id", updatePupil)
		pupil.DELETE("/:id", deletePupil)

		pupil.GET("/:id/classes", GetAllClasses)
	}

}

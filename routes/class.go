package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterClassServices(r *gin.Engine) {

	classController := &controllers.ClassController{}

	classes := r.Group("/classes", middleware.AdminMiddleware())
	{
		classes.POST("/", CreateClass)
		classes.GET("/", GetAllClasses)
		classes.GET("/:id", GetClassDetails)
		classes.PUT("/:id", UpdateClass)
		classes.DELETE("/:id", DeleteClass)

		classes.POST("/:id/pupil", AddPupilToClass)
		classes.POST("/:id/teacher", AssignTeacherToClass)
		classes.GET("/:id/pupils", GetPupilsInClass)
		classes.GET("/:id/teachers", GetTeachersInClass)
		classes.GET("/:id/activities", getClassActivities)
	}

	pupil := router.Group("/pupils")
	{
		pupil.POST("/", createPupil)
		pupil.GET("/", listPupils)
		pupil.GET("/:id", getPupil)
		pupil.PUT("/:id", updatePupil)
		pupil.DELETE("/:id", deletePupil)

		pupil.GET("/:id/classes", GetAllClasses)
	}

}

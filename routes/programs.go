package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func ProgramServices(rg *gin.RouterGroup) {
	programController := &controllers.ProgramController{}
	activityController := &controllers.ActivityController{}

	program := rg.Group("/programs", middleware.AuthMiddleware())
	{
		program.POST("/", programController.CreateProgram)
		program.GET("/", programController.GetAllPrograms)
		program.GET("/:id", programController.GetProgram)
		program.PUT("/:id", programController.UpdateProgram)
		program.DELETE("/:id", programController.DeleteProgram)

		program.GET("/:id/classes", programController.GetProgramClasses)
		program.GET("/:id/activities", programController.GetProgramActivities)
		program.POST("/:id/class", programController.AddClassToProgram)
		program.POST("/:id/activity", programController.AddActivityToProgram)
		program.DELETE("/activities/:id", programController.DeleteActivity)
	}

	activity := rg.Group("/activity", middleware.AuthMiddleware())
	{
		activity.POST("/", activityController.CreateActivity)
		activity.GET("/", activityController.GetAllActivities)
		activity.GET("/:id", activityController.GetActivity)
		activity.PUT("/:id", activityController.UpdateActivity)
		activity.DELETE("/:id", activityController.DeleteActivity)
	}

}

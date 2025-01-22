package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func ProgramServices(rg *gin.RouterGroup) {
	programController := &controllers.ProgramController{}

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
		program.POST("/:id/activity",programController. AddActivityToProgram)
		program.DELETE("/activities/:id", programController.DeleteActivity)
	}

}

package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func ProgramServices(r *Engine) {
	programs := r.Group("/programs", middleware.AuthMiddleware())
	{
		programs.POST("/", CreateProgram)
		programs.GET("/", GetAllPrograms)
		programs.GET("/:id", GetProgramDetails)
		programs.PUT("/:id", UpdateProgram)
		programs.DELETE("/:id", DeleteProgram)

		programs.GET("/:id/classes", GetProgramClasses)
		programs.GET("/:id/activities", GetProgramActivities)
		programs.POST("/:id/class", AddClassToProgram)
		programs.POST("/:id/activity", AddActivityToProgram)
		programs.DELETE("/activities/:id", DeleteActivity)
	}

}

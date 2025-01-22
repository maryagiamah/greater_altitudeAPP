package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func ProgramServices(rg *gin.RouterGroup) {
	program := rg.Group("/programs", middleware.AuthMiddleware())
	{
		program.POST("/", CreateProgram)
		program.GET("/", GetAllPrograms)
		program.GET("/:id", GetProgramDetails)
		program.PUT("/:id", UpdateProgram)
		program.DELETE("/:id", DeleteProgram)

		program.GET("/:id/classes", GetProgramClasses)
		program.GET("/:id/activities", GetProgramActivities)
		program.POST("/:id/class", AddClassToProgram)
		program.POST("/:id/activity", AddActivityToProgram)
		program.DELETE("/activities/:id", DeleteActivity)
	}

}

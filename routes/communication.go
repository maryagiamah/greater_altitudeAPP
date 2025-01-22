package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterCommunicationServices(rg *gin.RouterGroup) {

	message := rg.Group("/messages", middleware.AuthMiddleware())
	{
		message.POST("/", CreateMessage)
		message.GET("/inbox", GetInboxMessages)
		message.GET("/sent", GetSentMessages)
		message.PUT("/:id/read", MarkMessageAsRead)
		message.DELETE("/:id", DeleteMessage)
	}

	report := rg.Group("/reports", middleware.AuthMiddleware())
	{
		report.POST("/", CreateReport)
		report.GET("/", GetAllReports)
		report.GET("/:id", GetReportDetails)
		report.PUT("/:id", UpdateReport)
		report.DELETE("/:id", DeleteReport)

		report.GET("/pupil/:pupilId", GetPupilReports)
		report.GET("/teacher/:teacherId", GetTeacherReports)
	}

	event := rg.Group("/events", middleware.AuthMiddleware())
	{
		event.POST("/", CreateEvent)
		event.GET("/", GetAllEvents)
		event.GET("/:id", GetEventDetails)
		event.PUT("/:id", UpdateEvent)
		event.DELETE("/:id", DeleteEvent)
	}

}

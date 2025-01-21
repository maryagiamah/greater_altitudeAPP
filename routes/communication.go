package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterCommunicationServices(r *gin.Engine) {

	messages := r.Group("/messages", middleware.AuthMiddleware())
	{
		messages.POST("/", CreateMessage)
		messages.GET("/inbox", GetInboxMessages)
		messages.GET("/sent", GetSentMessages)
		messages.PUT("/:id/read", MarkMessageAsRead)
		messages.DELETE("/:id", DeleteMessage)
	}

	reports := r.Group("/reports", middleware.AuthMiddleware())
	{
		reports.POST("/", CreateReport)
		reports.GET("/", GetAllReports)
		reports.GET("/:id", GetReportDetails)
		reports.PUT("/:id", UpdateReport)
		reports.DELETE("/:id", DeleteReport)

		reports.GET("/pupil/:pupilId", GetPupilReports)
            reports.GET("/teacher/:teacherId", GetTeacherReports)
	}

	events := r.Group("/events", middleware.AuthMiddleware())
	{
		events.POST("/", CreateEvent)
		events.GET("/", GetAllEvents)
		events.GET("/:id", GetEventDetails)
		events.PUT("/:id", UpdateEvent)
		events.DELETE("/:id", DeleteEvent)
	}

}

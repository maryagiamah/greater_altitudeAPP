package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterCommunicationServices(rg *gin.RouterGroup) {

	messageController := &controllers.MessageController{}
	reportController := &controllers.ReportController{}
	eventController := &controllers.EventController{}

	message := rg.Group("/messages", middleware.AuthMiddleware())
	{
		message.POST("/", messageController.CreateMessage)
		message.GET("/inbox", messageController.GetInboxMessages)
		message.GET("/sent", messageController.GetSentMessages)
		message.PUT("/:id/read", messageController.MarkMessageAsRead)
		message.DELETE("/:id", messageController.DeleteMessage)
	}

	report := rg.Group("/reports", middleware.AuthMiddleware())
	{
		report.POST("/", reportController.CreateReport)
		report.GET("/", reportController.GetAllReports)
		report.GET("/:id", reportController.GetReport)
		report.PUT("/:id", reportController.UpdateReport)
		report.DELETE("/:id", reportController.DeleteReport)

		report.GET("/pupil/:pupilId", reportController.GetPupilReports)
		report.GET("/teacher/:teacherId", reportController.GetTeacherReports)
	}

	event := rg.Group("/events", middleware.AuthMiddleware())
	{
		event.POST("/", eventController.CreateEvent)
		event.GET("/", eventController.GetAllEvents)
		event.GET("/:id", eventController.GetEvent)
		event.PUT("/:id", eventController.UpdateEvent)
		event.DELETE("/:id", eventController.DeleteEvent)
	}

}

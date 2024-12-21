package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	pupilController := &controllers.PupilController{}
	parentController := &controllers.ParentController{}
	staffController := &controllers.StaffController{}

	v1 := r.Group("api/v1")
	{
		v1.GET("/staff/:id", staffController.FetchStaff)
		v1.GET("/parent/:id", parentController.FetchParent)
		v1.GET("/pupil/:id", pupilController.FetchPupil)

		v1.POST("/staff", staffController.CreateStaff)
		v1.POST("/parent", parentController.CreateParent)
		v1.POST("/pupil", pupilController.CreatePupil)

		v1.PUT("/staff/:id", staffController.UpdateStaff)
		v1.PUT("/parent/:id", parentController.UpdateParent)
		v1.PUT("/pupil/:id", pupilController.UpdatePupil)

		v1.DELETE("/staff/:id", staffController.DeleteStaff)
		v1.DELETE("/parent/:id", parentController.DeleteParent)
		v1.DELETE("/pupil/:id", pupilController.DeletePupil)
	}

/*	auth := r.Group("auth")
	{
	} */
}

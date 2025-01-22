package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		RegisterAdminServices(v1)
		RegisterBillingServices(v1)
		RegisterClassServices(v1)
		RegisterCommunicationServices(v1)
		ProgramServices(v1)
		RegisterUserServices(v1)
	}
}

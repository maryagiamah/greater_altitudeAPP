package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterBillingServices(rg *gin.RouterGroup) {
	invoice := rg.Group("/invoices")
	{
		invoice.GET("/", getAllInvoices)
		invoice.GET("/:id", getInvoiceByID)
		invoice.POST("/", createInvoice)
		invoice.POST("/:id/payments", makeInvoicePayment)
		invoice.PUT("/:id", updateInvoice)
		invoice.DELETE("/:id", deleteInvoice)
	}

	payment := rg.Group("/payments")
	{
		payment.GET("/", getAllPayments)
		payment.GET("/:id", getPaymentByID)
		payment.POST("/", createPayment)
		payment.PUT("/:id", updatePayment)
		payment.DELETE("/:id", deletePayment)
	}
}

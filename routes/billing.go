package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterBillingServices(r *gin.Engine) {
	invoiceRoutes := r.Group("/invoices")
	{
		invoiceRoutes.GET("/", getAllInvoices)
		invoiceRoutes.GET("/:id", getInvoiceByID)
		invoiceRoutes.POST("/", createInvoice)
		invoiceRoutes.POST("/:id/payments", makeInvoicePayment)
		invoiceRoutes.PUT("/:id", updateInvoice)
		invoiceRoutes.DELETE("/:id", deleteInvoice)
	}

	paymentRoutes := r.Group("/payments")
	{
		paymentRoutes.GET("/", getAllPayments)
		paymentRoutes.GET("/:id", getPaymentByID)
		paymentRoutes.POST("/", createPayment)
		paymentRoutes.PUT("/:id", updatePayment)
		paymentRoutes.DELETE("/:id", deletePayment)
	}
}

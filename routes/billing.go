package routes

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/controllers"
	"greaterAltitudeapp/middleware"
)

func RegisterBillingServices(rg *gin.RouterGroup) {
	invoiceController := &controllers.InvoiceController{}
	paymentController := &controllers.PaymentController{}

	invoice := rg.Group("/invoices", middleware.AuthMiddleware())
	{
		invoice.GET("/", invoiceController.GetAllInvoices)
		invoice.GET("/:id", invoiceController.GetInvoice)
		invoice.POST("/", invoiceController.CreateInvoice)
		invoice.POST("/:id/payments", invoiceController.MakePayment)
		invoice.PUT("/:id", invoiceController.UpdateInvoice)
		invoice.DELETE("/:id", invoiceController.DeleteInvoice)
	}

	payment := rg.Group("/payments", middleware.AuthMiddleware())
	{
		payment.GET("/", paymentController.GetAllPayments)
		payment.GET("/:id", paymentController.GetPayment)
		payment.POST("/", paymentController.CreatePayment)
		payment.PUT("/:id", paymentController.UpdatePayment)
		payment.DELETE("/:id", paymentController.DeletePayment)
	}
}

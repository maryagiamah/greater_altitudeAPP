package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type InvoiceController struct{}

func (i *InvoiceController) GetInvoice(c *gin.Context) {
	id := c.Param("id")
	var invoice models.Invoice

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&invoice, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Invoice not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	c.JSON(200, gin.H{"invoice": invoice})
}

func (i *InvoiceController) GetAllInvoices(c *gin.Context) {
	var invoices []models.Invoice

	if err := utils.H.DB.Find(&invoices).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(invoices) == 0 {
		c.JSON(404, gin.H{"error": "No invoice found"})
		return
	}
	c.JSON(200, gin.H{"invoices": invoices})
}

func (i *InvoiceController) CreateInvoice(c *gin.Context) {
	var newInvoice models.Invoice

	if err := c.ShouldBindJSON(&newInvoice); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := utils.H.DB.Create(&newInvoice)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Invoice created successfully",
		"ID":      newInvoice.ID,
	})
}

func (i *InvoiceController) UpdateInvoice(c *gin.Context) {
	id := c.Param("id")
	var invoice models.Invoice
	var updatedFields models.Invoice

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&invoice, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Invoice not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&invoice).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Invoice updated successfully",
		"ID":      invoice.ID,
	})
}

func (i *InvoiceController) DeleteInvoice(c *gin.Context) {
	id := c.Param("id")
	var invoice models.Invoice

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&invoice, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Invoice not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Invoice deleted successfully"})
}

func (i *InvoiceController) MakePayment(c *gin.Context) {
	id := c.Param("id")
	var invoice models.Invoice
	var newPayment models.Payment

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&invoice, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Invoice not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	if err := c.ShouldBindJSON(&newPayment); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.Association("Payments").Append(newPayment).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
	}

	c.JSON(200, gin.H{"message": "Payment succesfully added to invoice"})
}

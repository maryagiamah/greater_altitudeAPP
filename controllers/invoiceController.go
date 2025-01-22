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
	utils.H.Logger.Printf("Fetched Invoice: %s", invoice.Description)
	c.JSON(200, gin.H{"invoice": invoice})
}

func (i *InvoiceController) GetAllInvoices(c *gin.Context) {
}

func (i *InvoiceController) CreateInvoice(c *gin.Context) {
	var newInvoice models.Invoice

	if err := c.ShouldBindJSON(&newInvoice); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := utils.H.DB.Create(&newInvoice)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Invoice"})
		return
	}

	utils.H.Logger.Printf("New Invoice Created with ID: %d", newInvoice.ID)
	c.JSON(201, gin.H{"ID": newInvoice.ID})
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
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
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
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update invoice"})
		utils.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	utils.H.Logger.Printf("Updated invoice with ID: %d", invoice.ID)
	c.JSON(200, gin.H{"ID": invoice.ID})
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
	utils.H.Logger.Printf("Deleted Invoice with ID: %s", id)
	c.JSON(200, gin.H{"message": "Invoice deleted successfully"})
}

func (i *InvoiceController) MakePayment(c *gin.Context) {
}

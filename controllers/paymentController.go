package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type PaymentController struct{}

func (p *PaymentController) GetPayment(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&payment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Payment not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	utils.H.Logger.Printf("Fetched Payment: %s", payment.Reference)
	c.JSON(200, gin.H{"payment": payment})
}

func (p *PaymentController) GetAllPayments(c *gin.Context) {
	var payments []models.Payment

	if err := utils.H.DB.Find(&payments).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(payments) == 0 {
		c.JSON(404, gin.H{"error": "No payment found"})
		return
	}
	c.JSON(200, gin.H{"payments": payments})
}

func (p *PaymentController) CreatePayment(c *gin.Context) {
	var newPayment models.Payment

	if err := c.ShouldBindJSON(&newPayment); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := utils.H.DB.Create(&newPayment)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Payment"})
		return
	}

	utils.H.Logger.Printf("New Payment Created with ID: %d", newPayment.ID)
	c.JSON(201, gin.H{"ID": newPayment.ID})
}

func (p *PaymentController) UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment
	var updatedFields models.Payment

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := utils.H.DB.First(&payment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Payment not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&payment).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update payment"})
		utils.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	utils.H.Logger.Printf("Updated payment with ID: %d", payment.ID)
	c.JSON(200, gin.H{"ID": payment.ID})
}

func (p *PaymentController) DeletePayment(c *gin.Context) {
	id := c.Param("id")
	var payment models.Payment

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&payment, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Payment not found"})
		return
	}
	utils.H.Logger.Printf("Deleted Payment with ID: %s", id)
	c.JSON(200, gin.H{"message": "Payment deleted successfully"})
}

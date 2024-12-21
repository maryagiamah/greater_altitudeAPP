package controllers

import (
	"greaterAltitudeapp/config"
	"greaterAltitudeapp/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
type ParentController struct {}

func (p *ParentController) FetchParent(c *gin.Context) {
	id := c.Param("id")
	var parent models.Parent


	if id != "" {
		if err := config.H.DB.First(&parent, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, gin.H{"Error": "Parent not found"})
			} else {
				c.JSON(500, gin.H{"Error": "Internal Server Error"})
			}
			return
		}
		config.H.Logger.Printf("Fetched Parent %d: %s", parent.ID, parent.FirstName)
		c.JSON(200, gin.H{"Parent": parent})
	} else {
		c.JSON(400, gin.H{"Error": "ID cannot be empty"})
	}

}

func (p *ParentController) CreateParent(c *gin.Context) {
}

func (p *ParentController) UpdateParent(c *gin.Context) {
}

func (p *ParentController) DeleteParent(c *gin.Context) {
}

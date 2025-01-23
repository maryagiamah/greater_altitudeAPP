package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
)

type ReportController struct{}

func (r *ReportController) GetReport(c *gin.Context) {
	id := c.Param("id")
	var report models.Report

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&report, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Report not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	utils.H.Logger.Printf("Fetched Report: %s", report.Type)
	c.JSON(200, gin.H{"report": report})
}

func (r *ReportController) GetAllReports(c *gin.Context) {
	var reports []models.Report

	if err := utils.H.DB.Find(&reports).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(reports) == 0 {
		c.JSON(404, gin.H{"error": "No report found"})
		return
	}
	c.JSON(200, gin.H{"reports": reports})
}

func (r *ReportController) CreateReport(c *gin.Context) {
	var newReport models.Report

	if err := c.ShouldBindJSON(&newReport); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	result := utils.H.DB.Create(&newReport)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't create Report"})
		return
	}

	utils.H.Logger.Printf("New Report Created with ID: %d", newReport.ID)
	c.JSON(201, gin.H{"ID": newReport.ID})
}

func (r *ReportController) UpdateReport(c *gin.Context) {
	id := c.Param("id")
	var report models.Report
	var updatedFields models.Report

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not a JSON"})
		return
	}

	if err := utils.H.DB.First(&report, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(404, gin.H{"error": "Report not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&report).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Can't update report"})
		utils.H.Logger.Printf("Update failed: %v", result.Error)
		return
	}

	utils.H.Logger.Printf("Updated report with ID: %d", report.ID)
	c.JSON(200, gin.H{"ID": report.ID})
}

func (r *ReportController) DeleteReport(c *gin.Context) {
	id := c.Param("id")
	var report models.Report

	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&report, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Report not found"})
		return
	}
	utils.H.Logger.Printf("Deleted Report with ID: %s", id)
	c.JSON(200, gin.H{"message": "Report deleted successfully"})
}

func (r *ReportController) GetPupilReports(c *gin.Context) {
}

func (r *ReportController) GetTeacherReports(c *gin.Context) {
}

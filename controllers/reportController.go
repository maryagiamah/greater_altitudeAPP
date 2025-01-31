package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"net/http"
)

type ReportController struct{}

func (r *ReportController) GetReport(c *gin.Context) {
	id := c.Param("id")
	var report models.Report

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.First(&report, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"report": report})
}

func (r *ReportController) GetAllReports(c *gin.Context) {
	var reports []models.Report

	if err := utils.H.DB.Find(&reports).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if len(reports) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No report found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reports": reports})
}

func (r *ReportController) CreateReport(c *gin.Context) {
	var newReport models.Report

	if err := c.ShouldBindJSON(&newReport); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := utils.H.DB.Create(&newReport)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Report"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ID":      newReport.ID,
		"message": "Report Created",
	})
}

func (r *ReportController) UpdateReport(c *gin.Context) {
	id := c.Param("id")
	var report models.Report
	var updatedFields models.Report

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.H.DB.First(&report, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}
	result := utils.H.DB.Model(&report).Updates(updatedFields)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update report"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":      report.ID,
		"message": "Message updated",
	})
}

func (r *ReportController) DeleteReport(c *gin.Context) {
	id := c.Param("id")
	var report models.Report

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	result := utils.H.DB.Delete(&report, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Report deleted successfully"})
}

func (r *ReportController) GetPupilReports(c *gin.Context) {
	pupilId := c.Param("pupilId")
	var report []models.Report

	if pupilId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.Where("pupil_id = ?", pupilId).Find(&report).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"pupil_reports": report})
}

func (r *ReportController) GetTeacherReports(c *gin.Context) {
	teacherId := c.Param("teacherId")
	var report []models.Report

	if teacherId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID cannot be empty"})
		return
	}

	if err := utils.H.DB.Where("teacher_id = ?", teacherId).Find(&report).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"teacher_reports": report})
}

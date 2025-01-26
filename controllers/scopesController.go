package controllers

import (
	"gorm.io/gorm"
	"greaterAltitudeapp/utils"
)

func Pagination(db *gorm.DB, page, page_size int) *gorm.DB {
	return utils.H.DB
}

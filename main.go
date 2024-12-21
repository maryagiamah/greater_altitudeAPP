package main

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/config"
	"greaterAltitudeapp/routes"
	"greaterAltitudeapp/utils"
)

var H *utils.DBHandler

func init() {
	config.InitDB()
}
func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}

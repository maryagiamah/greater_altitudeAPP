package main

import (
	"github.com/gin-gonic/gin"
	"greaterAltitudeapp/utils"
	"greaterAltitudeapp/routes"
)

func init() {
	utils.InitDB()
}
func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	defer utils.CloseDB()
	r.Run(":8080")
}

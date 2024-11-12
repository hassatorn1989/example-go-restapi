package main

import (
	"inv-app/configs"
	"inv-app/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	r := gin.Default()

	configs.ConnectDatabase()

	routes.ApiRoutes(r)
	r.Run(":8080")
}

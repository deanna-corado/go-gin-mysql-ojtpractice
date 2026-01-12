package main

import (
	"go-gin-mysql/config"
	"go-gin-mysql/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}

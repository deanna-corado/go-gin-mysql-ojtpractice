package main

import (
	"go-gin-mysql/config"
	"go-gin-mysql/controllers"
	"go-gin-mysql/repositories"
	"go-gin-mysql/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	r := gin.Default()

	movieRepo := repositories.NewMovieRepository(config.DB)

	movieController := controllers.NewMovieController(movieRepo)

	routes.RegisterRoutes(r, movieController)

	r.Run(":8080")
}

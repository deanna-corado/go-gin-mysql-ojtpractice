package main

import (
	"go-gin-mysql/config"
	"go-gin-mysql/controllers"
	"go-gin-mysql/repositories"
	"go-gin-mysql/routes"
	"go-gin-mysql/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	r := gin.Default()

	movieRepo := repositories.NewMovieRepository(config.DB)
	movieService := services.NewMovieService(movieRepo)
	movieController := controllers.NewMovieController(movieService)

	routes.RegisterRoutes(r, movieController)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}

}

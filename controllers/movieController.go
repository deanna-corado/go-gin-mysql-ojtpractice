// CONTROLLER - BUILD RESPONSES TO REQUESTS AND INTERACT WITH THE DATABASE
package controllers

import (
	"net/http"

	"go-gin-mysql/config"
	"go-gin-mysql/models"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, director FROM movies")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		rows.Scan(&movie.ID, &movie.Title, &movie.Director)
		movies = append(movies, movie)
	}

	c.JSON(http.StatusOK, movies)
}

func AddMovie(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := "INSERT INTO movies (id, title, director) VALUES (?, ?, ?)"
	_, err := config.DB.Exec(query, movie.ID, movie.Title, movie.Director)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// CONTROLLER - BUILD RESPONSES TO REQUESTS AND INTERACT WITH THE DATABASE
package controllers

import (
	"net/http"

	"go-gin-mysql/config"
	"go-gin-mysql/models"

	"github.com/gin-gonic/gin"

	"strconv"
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

func UpdateMovie(c *gin.Context) {
	var movie models.Movie
	id := c.Param("id")

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	query := `UPDATE movies SET title = ?, director = ? WHERE id = ?
`
	result, err := config.DB.Exec(query, movie.Title, movie.Director, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	movie.ID = atoi(id)
	c.JSON(http.StatusOK, movie)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func DeleteMovie(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	query := `DELETE FROM movies WHERE id = ?`
	result, err := config.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}

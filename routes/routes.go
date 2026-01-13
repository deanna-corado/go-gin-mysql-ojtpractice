// ROUTES - FOR CALLING CONTROLLERS BASED ON ENDPOINTS; CALL API
package routes

import (
	// "database/sql"
	// "go-gin-mysql/controllers"
	// "go-gin-mysql/repositories"

	"go-gin-mysql/controllers"

	"github.com/gin-gonic/gin"
)

// FIX BY GROUP
func RegisterRoutes(
	r *gin.Engine,
	movieController *controllers.MovieController,
) {
	//for api versioning
	v1 := r.Group("/api/v1")
	movies := v1.Group("/movies")
	{
		movies.GET("", movieController.GetMovies)
		movies.GET(":id", movieController.GetMovieByID)
		movies.POST("", movieController.AddMovie)
		movies.PUT(":id", movieController.UpdateMovie)
		movies.DELETE(":id", movieController.DeleteMovie)
	}
}

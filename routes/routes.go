// ROUTES - FOR CALLING CONTROLLERS BASED ON ENDPOINTS; CALL API
package routes

import (
	"go-gin-mysql/controllers"
	"go-gin-mysql/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	movieController *controllers.MovieController,
) {
	//for api versioning
	v1 := r.Group("/api/v1")

	//PUBLIC ROUTES
	movies := v1.Group("/movies")
	{
		movies.GET("", movieController.GetMovies)
		movies.GET("/:id", movieController.GetMovieByID)

	}

	//PRIVATE ROUTES (ADMIN ACCESS)
	adminMovies := v1.Group("/admin/movies", middlewares.AuthRequired())
	{
		adminMovies.POST("", movieController.AddMovie)
		adminMovies.PUT("/:id", movieController.UpdateMovie)
		adminMovies.DELETE("/:id", movieController.DeleteMovie)
	}
}

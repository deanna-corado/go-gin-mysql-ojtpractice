// ROUTES - FOR CALLING CONTROLLERS BASED ON ENDPOINTS; CALL API
package routes

import (
	// "database/sql"
	// "go-gin-mysql/controllers"
	// "go-gin-mysql/repositories"

	"go-gin-mysql/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	movieController *controllers.MovieController,
) {
	r.GET("/movies", movieController.GetMovies)

}

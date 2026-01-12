//ROUTES - FOR CALLING CONTROLLERS BASED ON ENDPOINTS; CALL API
package routes

import (
	"go-gin-mysql/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/movies", controllers.GetMovies)
	r.POST("/movies", controllers.AddMovie)
}

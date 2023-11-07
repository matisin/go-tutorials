package cities

import (
	"github.com/gin-gonic/gin"
)

func UseRoutes(r *gin.Engine) {
	cityGroup := r.Group("/cities")
	{
		cityGroup.POST("/", CreateController)
		cityGroup.GET("/", FindAllController)
		cityGroup.GET("/id/:id", FindOneController)
		// cityGroup.DELETE("/id/:id", Delete)
		// cityGroup.PUT("/id/:id", Update)
	}
}

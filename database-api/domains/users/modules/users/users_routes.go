package users

import (
	"github.com/gin-gonic/gin"
)

func UseRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", CreateController)
		userGroup.GET("/", FindAllController)
		userGroup.GET("/id/:id", FindOneController)
		// userGroup.DELETE("/id/:id", Delete)
		// userGroup.PUT("/id/:id", Update)
	}
}

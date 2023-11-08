package users

import (
	"github.com/gin-gonic/gin"
)

func UseRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", CreateController)
		userGroup.GET("/", FindController)
		userGroup.GET("/id/:id", FindOneController)
		userGroup.PUT("/id/:id", UpdateOneController)
		userGroup.DELETE("/id/:id", DeleteOneController)

	}
}

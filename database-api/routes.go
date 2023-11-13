package module

import (

	"github.com/gin-gonic/gin"

	userEntities "database-api/entities/users"
)

struct 



func UseRoutes(r *gin.Engine) {

	userModule := Module.Register(userEntities.User)

	userGroup := r.Group("/users")
	{
		userGroup.POST("/", CreateController)
		userGroup.GET("/", FindController)
		userGroup.GET("/id/:id", FindOneController)
		userGroup.PUT("/id/:id", UpdateOneController)
		userGroup.DELETE("/id/:id", DeleteOneController)

	}
}

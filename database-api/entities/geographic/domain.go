package entitites

import (
	"database-api/domains/geographic/modules/cities"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	routes.BindRoute(City)
	cities.UseRoutes(r)
}

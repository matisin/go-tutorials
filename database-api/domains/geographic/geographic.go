package geographic

import (
	"database-api/domains/geographic/modules/cities"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	cities.UseRoutes(r)
}

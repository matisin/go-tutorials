package users

import (
	"database-api/domains/users/modules/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	users.UseRoutes(r)
}

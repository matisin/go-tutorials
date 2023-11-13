package main

import (
	"database-api/config"
	"database-api/db"
	"database-api/domains/geographic"
	"database-api/domains/users"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	_, err := db.InitDB()
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}

	users.RegisterRoutes(r)
	geographic.RegisterRoutes(r)

	module.RegisterDomain("users", r)

	r.Run(":" + config.AppPort)
}

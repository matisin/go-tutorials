package users

import (
	"database-api/domains/users/entities"

	"github.com/gin-gonic/gin"
)

func CreateController(c *gin.Context) {
	var data entities.User
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Datos de usuario no válidos"})
		return
	}

	newUser, err := CreateService(data)
	if err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	c.JSON(201, newUser)
}

func FindOneController(c *gin.Context) {
	userID := c.Param("id")

	fetchedUser, err := FindOneService(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(200, fetchedUser)
}

func FindAllController(c *gin.Context) {

	fetchedUsers, err := FindAllService()
	if err != nil {
		c.JSON(404, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(200, fetchedUsers)
}

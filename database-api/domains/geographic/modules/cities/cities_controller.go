package cities

import (
	"database-api/domains/geographic/entities"

	"github.com/gin-gonic/gin"
)

func CreateController(c *gin.Context) {
	var data entities.City
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Datos de usuario no válidos"})
		return
	}

	newCity, err := CreateService(data)
	if err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	c.JSON(201, newCity)
}

func FindOneController(c *gin.Context) {
	userID := c.Param("id")

	fetchedCity, err := FindOneService(userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(200, fetchedCity)
}

func FindAllController(c *gin.Context) {

	fetchedCities, err := FindAllService()
	if err != nil {
		c.JSON(404, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(200, fetchedCities)
}

package main

import (
	"database-api/config"
	"database-api/db"
	"database-api/domains/users/modules/users"

	"github.com/gin-gonic/gin"
	// "net/http"
	// "./db"
	// "github.com/gin-gonic/gin"
	// "time"
)

// var db = make(map[string]string)

// func setupRouter() *gin.Engine {
// 	// Disable Console Color
// 	// gin.DisableConsoleColor()
// 	r := gin.Default()

// 	// Ping test
// 	r.GET("/ping", func(c *gin.Context) {
// 		numIteraciones := 100000000000

// 		// Marca de tiempo antes de iniciar los cálculos
// 		inicio := time.Now()

// 		// Realiza muchos cálculos (puedes reemplazar esto con la lógica de tu API)

// 		var resultado int
// 		for i := 0; i < numIteraciones; i++ {
// 			resultado = i * 2
// 			// Puedes hacer algo con el resultado si es necesario
// 			_ = resultado
// 		}

// 		// Marca de tiempo después de completar los cálculos
// 		fin := time.Now()

// 		// Calcula la duración
// 		duracion := fin.Sub(inicio)

// 		respuesta := fmt.Sprintf("Tiempo total: %v ms\n, Resultado: %v", duracion.Milliseconds(), resultado)
// 		c.String(http.StatusOK, respuesta)
// 	})

// 	// Get user value
// 	r.GET("/user/:name", func(c *gin.Context) {
// 		user := c.Params.ByName("name")
// 		value, ok := db[user]
// 		if ok {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
// 		}
// 	})

// 	// Authorized group (uses gin.BasicAuth() middleware)
// 	// Same than:
// 	// authorized := r.Group("/")
// 	// authorized.Use(gin.BasicAuth(gin.Credentials{
// 	//	  "foo":  "bar",
// 	//	  "manu": "123",
// 	//}))
// 	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
// 		"foo":  "bar", // user:foo password:bar
// 		"manu": "123", // user:manu password:123
// 	}))

// 	/* example curl for /admin with basicauth header
// 	   Zm9vOmJhcg== is base64("foo:bar")

// 		curl -X POST \
// 	  	http://localhost:8080/admin \
// 	  	-H 'authorization: Basic Zm9vOmJhcg==' \
// 	  	-H 'content-type: application/json' \
// 	  	-d '{"value":"bar"}'
// 	*/
// 	authorized.POST("admin", func(c *gin.Context) {
// 		user := c.MustGet(gin.AuthUserKey).(string)

// 		// Parse JSON
// 		var json struct {
// 			Value string `json:"value" binding:"required"`
// 		}

// 		if c.Bind(&json) == nil {
// 			db[user] = json.Value
// 			c.JSON(http.StatusOK, gin.H{"status": "ok"})
// 		}
// 	})

// 	return r
// }

func main() {
	r := gin.Default()

	_, err := db.InitDB()
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}

	// Configura rutas para cada dominio
	users.UserRoutes(r)
	// Agrega rutas para otros dominios si es necesario

	r.Run(":" + config.AppPort)
}

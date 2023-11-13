package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"user-service/internal/controller"
	"user-service/internal/core/server"
	"user-service/internal/core/service"
	"user-service/internal/infra/config"
	"user-service/internal/infra/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new instance of the Gin router
	instance := gin.New()
	instance.Use(gin.Recovery())

	// Initialize the database connection
	db, err := repository.NewDB()
	if err != nil {
		log.Fatalf("failed to new database err=%s\n", err.Error())
	}

	// Create the User Repository, Service and Controller
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(instance, userService)

	// Initialize the routes for UserController
	userController.InitRouter()

	// Create the HTTP server
	httpServer := server.NewHttpServer(
		instance,
		config.HttpServerConfig{
			Port: 8000,
		},
	)

	// Start the HTTP server
	httpServer.Start()
	defer httpServer.Stop()

	// Listen for OS signals to perform a graceful shutdown
	log.Println("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}

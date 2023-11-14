package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"users-service/internal/controller"
	"users-service/internal/core/server"
	"users-service/internal/core/service"
	"users-service/internal/infra/config"
	"users-service/internal/infra/repository"

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

	err = db.RunMigrations()
	if err != nil {
		log.Fatalf("failed to run migrations err=%s\n", err.Error())
	} else {
		log.Println("Migrations run succesful")
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
		config.HttpConfig{
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

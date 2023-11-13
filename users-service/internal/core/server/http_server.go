package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"users-service/internal/infra/config"

	"github.com/gin-gonic/gin"
)

const defaultHost = "0.0.0.0"

type HttpServer interface {
	Start()
	Stop()
}

type httpServer struct {
	Port   uint
	server *http.Server
}

func NewHttpServer(router *gin.Engine, config config.HttpConfig) HttpServer {
	return &httpServer{
		Port: config.Port,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", defaultHost, config.Port),
			Handler: router,
		},
	}
}

func (httpServer httpServer) Start() {
	go func() {
		if err := httpServer.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(
				"failed to stater HttpServer listen port %d, err=%s\n",
				httpServer.Port, err.Error(),
			)
		}
	}()
	log.Printf("Start Service with port %d\n", httpServer.Port)
}

func (httpServer httpServer) Stop() {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(3)*time.Second,
	)
	defer cancel()

	if err := httpServer.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown err=%s\n", err.Error())
	}
}

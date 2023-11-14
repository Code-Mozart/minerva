package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func (server *Server) Init() {
	server.router = gin.Default()
	RegisterRoutes(server.router)
}

func (server *Server) Run(address string) {
	if err := server.router.Run(address); err != nil {
		fmt.Printf("Runtime error: %e", err)
	}
}

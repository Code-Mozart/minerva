package server

import (
	"fmt"

	"github.com/Code-Mozart/minerva/minerva-backend/initializers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Environment *initializers.Environment

	engine   *gin.Engine
	router   *Router
	database *gorm.DB
}

func Initialize(env *initializers.Environment) (*Server, error) {
	server := &Server{
		Environment: env,

		engine:   gin.Default(),
		router:   &Router{},
		database: nil,
	}

	server.router.server = server
	RegisterRoutes(server.router)

	database, err := server.InitDatabase()
	if err != nil {
		return nil, err
	}
	server.database = database

	return server, nil
}

func (server *Server) Run(address string) {
	if err := server.engine.Run(address); err != nil {
		fmt.Printf("Runtime error: %e", err)
	}
}

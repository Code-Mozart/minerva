package server

import (
	"github.com/Code-Mozart/minerva/minerva-backend/internal/controllers"
)

func RegisterRoutes(router *Router) {
	router.GET("/status", controllers.Status)
}

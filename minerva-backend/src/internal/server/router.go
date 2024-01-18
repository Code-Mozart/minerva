package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerFunction func(*gin.Context, *gorm.DB) error

type GinMethodFunction func(string, ...gin.HandlerFunc) gin.IRoutes

type Router struct {
	server *Server
}

func (router *Router) GET(route string, handler HandlerFunction) {
	router.register(router.server.engine.GET, route, handler)
}

func (router *Router) POST(route string, handler HandlerFunction) {
	router.register(router.server.engine.POST, route, handler)
}

func (router *Router) PUT(route string, handler HandlerFunction) {
	router.register(router.server.engine.PUT, route, handler)
}

func (router *Router) PATCH(route string, handler HandlerFunction) {
	router.register(router.server.engine.PATCH, route, handler)
}

func (router *Router) DELETE(route string, handler HandlerFunction) {
	router.register(router.server.engine.DELETE, route, handler)
}

func (router *Router) register(method GinMethodFunction, route string, handler HandlerFunction) {
	method(route, withErrorHandling(handler, router.server.database))
}

func withErrorHandling(handler HandlerFunction, database *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := handler(context, database); err != nil {
			HandleError(context, err)
		}
	}
}

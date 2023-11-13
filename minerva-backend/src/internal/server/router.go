package server

import "github.com/gin-gonic/gin"

type HandlerFunction func(*gin.Context) error

type GinMethodFunction func(string, ...gin.HandlerFunc) gin.IRoutes

type Router struct {
	context *gin.Engine
}

func (router *Router) GET(route string, handler HandlerFunction) {
	router.register(router.context.GET, route, handler)
}

func (router *Router) POST(route string, handler HandlerFunction) {
	router.register(router.context.POST, route, handler)
}

func (router *Router) PUT(route string, handler HandlerFunction) {
	router.register(router.context.PUT, route, handler)
}

func (router *Router) PATCH(route string, handler HandlerFunction) {
	router.register(router.context.PATCH, route, handler)
}

func (router *Router) DELETE(route string, handler HandlerFunction) {
	router.register(router.context.DELETE, route, handler)
}

func (router *Router) register(method GinMethodFunction, route string, handler HandlerFunction) {
	method(route, withErrorHandling(handler))
}

func withErrorHandling(handler HandlerFunction) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := handler(context); err != nil {
			HandleError(context, err)
		}
	}
}

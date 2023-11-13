package server

import (
	"github.com/Code-Mozart/minerva/minerva-backend/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(context *gin.Engine) {
	context.GET("/status", controllers.Status)
}

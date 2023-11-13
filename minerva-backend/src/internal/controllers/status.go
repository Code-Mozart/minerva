package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Status(context *gin.Context) {
	context.Status(http.StatusOK)
}

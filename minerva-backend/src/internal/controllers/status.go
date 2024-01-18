package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Status(context *gin.Context, database *gorm.DB) error {
	context.Status(http.StatusOK)
	context.Writer.WriteString("OK")
	return nil
}

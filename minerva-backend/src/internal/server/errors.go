package server

import "github.com/gin-gonic/gin"

func HandleError(context *gin.Context, err error) {
	context.Error(err)
}

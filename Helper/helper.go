package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusBadRequest(ctx *gin.Context, error string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"Error": error,
	})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg ": "DELETE Opening",
	})
}

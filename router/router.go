package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// initialize router
	var router *gin.Engine = gin.Default()

	initializeRoutes(router)

	router.Run(":8080")
}

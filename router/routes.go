package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rochaandrey/goproject.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	var rota *gin.RouterGroup = router.Group("/api/rota")
	{
		rota.POST("/opening", handler.CreateOpeningHandler)

		rota.GET("/opening", handler.ShowOpeningHandler)

		rota.DELETE("/opening", handler.DeleteOpeningHandler)

		rota.PUT("/opening", handler.UpdateOpeningHandler)

		rota.GET("/openings", handler.ListOpeningHandler)
	}
}

package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *LineBotHandler) {
	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/callback", handler.Callback)
}

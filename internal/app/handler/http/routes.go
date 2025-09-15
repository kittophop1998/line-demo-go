package http

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func SetupRoutes(router *gin.Engine, bot *linebot.Client) {
	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	handler := NewLineBotHandler(bot)

	router.POST("/callback", handler.Callback)
}

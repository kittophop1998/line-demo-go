package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func SetupRoutes(router *gin.Engine, bot *linebot.Client) {
	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/callback", func(ctx *gin.Context) {
		events, err := bot.ParseRequest(ctx.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				ctx.Status(http.StatusBadRequest)
			} else {
				ctx.Status(http.StatusInternalServerError)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if message.Text == "ping" {
						if _, err := bot.ReplyMessage(event.ReplyToken,
							linebot.NewTextMessage("pong")).Do(); err != nil {
							log.Println("Failed to reply message:", err)
						}
						continue
					}

					if _, err := bot.ReplyMessage(event.ReplyToken,
						linebot.NewTextMessage("คุณพิมพ์ว่า: "+message.Text)).Do(); err != nil {
						log.Println("Failed to reply message:", err)
					}
				}
			}
		}

		ctx.Status(http.StatusOK)
	})
}

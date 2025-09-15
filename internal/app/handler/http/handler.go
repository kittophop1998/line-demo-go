package http

import (
	"log"
	"net/http"
	"strings"

	"line-bot/internal/app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type LineBotHandler struct {
	Bot      *linebot.Client
	UseCase  *usecase.DebtUseCase
	commands map[string]func(event *linebot.Event)
}

func NewLineBotHandler(bot *linebot.Client, debtUC *usecase.DebtUseCase) *LineBotHandler {
	// ===== Initialize Handler =====
	h := &LineBotHandler{
		Bot:      bot,
		commands: make(map[string]func(event *linebot.Event)),
		UseCase:  debtUC,
	}

	// === Setup commands ====
	h.commands = map[string]func(event *linebot.Event){
		"check debt": func(e *linebot.Event) {
			debt, err := h.UseCase.GetDebts()
			if err != nil {
				h.replyText(e, "Error fetching debt info")
				return
			}
			h.replyText(e, "Your debt: "+debt)
		},
	}

	return h
}

// Healthz สำหรับตรวจสอบ server
func (h *LineBotHandler) Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Callback จัดการ webhook จาก LINE
func (h *LineBotHandler) Callback(c *gin.Context) {
	events, err := h.Bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Status(http.StatusBadRequest)
		} else {
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			h.dispatch(event)
		}
	}

	c.Status(http.StatusOK)
}

// dispatch ตรวจสอบข้อความและเรียก command map
func (h *LineBotHandler) dispatch(event *linebot.Event) {
	switch msg := event.Message.(type) {
	case *linebot.TextMessage:
		command := strings.ToLower(strings.TrimSpace(msg.Text))
		if cmdFunc, ok := h.commands[command]; ok {
			cmdFunc(event)
		} else {
			h.replyText(event, "Unknown command: "+msg.Text)
		}
	default:
		log.Printf("Unhandled message type: %T", msg)
	}
}

// replyText ส่งข้อความตอบกลับ
func (h *LineBotHandler) replyText(event *linebot.Event, text string) {
	if _, err := h.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)).Do(); err != nil {
		log.Printf("Failed to reply to user %s: %v", event.Source.UserID, err)
	}
}

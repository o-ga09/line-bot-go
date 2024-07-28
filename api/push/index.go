package push

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/o-ga09/line-bot-go/pkg/config"
	"github.com/o-ga09/line-bot-go/pkg/logger"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	logger.Logger()
	cfg, err := config.New()
	if err != nil {
		slog.Error(fmt.Sprintf("config error: %v", err))
	}
	access_token := cfg.LineAccesstoken
	userId := cfg.LineUserId
	bot, err := messaging_api.NewMessagingApiAPI(access_token)
	if err != nil {
		slog.Error(fmt.Sprintf("can not connect line messaging api: %v", err))
	}

	retrykey := uuid.New().String()
	replyMessage := "Hello"

	bot.PushMessage(
		&messaging_api.PushMessageRequest{
			To:                   userId,
			NotificationDisabled: true,
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: replyMessage,
				},
			},
		},
		retrykey,
	)
	slog.Info("push message success")
}

package push

import (
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	access_token := os.Getenv("LINE_ACCESS_TOKEN")
	userId := os.Getenv("USERID")
	bot, err := messaging_api.NewMessagingApiAPI(access_token)
	if err != nil {
		log.Println(err)
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
}

package callback

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	channel_secret := os.Getenv("LINE_CHANNEL_SECRET")
	access_token := os.Getenv("LINE_ACCESS_TOKEN")
	bot, err := messaging_api.NewMessagingApiAPI(access_token)

	var reply_message string
	if err != nil {
		log.Fatalf("can not connect line messaging api: %v", err)
	}

	cb, err := webhook.ParseRequest(channel_secret, r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	var name, content, sticker, replymessage string
	for _, event := range cb.Events {
		// eventの種類でスイッチ
		switch e := event.(type) {
		case webhook.MessageEvent:
			// profileを取得
			switch s := e.Source.(type) {
			case webhook.UserSource:
				res, err := bot.GetProfile(s.UserId)
				if err != nil {
					log.Println(err)
				}
				name = res.DisplayName
			}
			// 送信メッセージを取得
			switch m := e.Message.(type) {
			case webhook.TextMessageContent:
				content = m.Text
			case webhook.StickerMessageContent:
				sticker = m.StickerId
			}

			// メッセージを返信する
			if content != "" {
				replymessage = fmt.Sprintf("送ってきたメッセージは、%sです", content)
			} else {
				replymessage = fmt.Sprintf("StickerIdは、%sです", sticker)
			}

			reply_message = fmt.Sprintf("%sさん！ありがとうございます。", name)
			if _, err := bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
				ReplyToken: e.ReplyToken,
				Messages: []messaging_api.MessageInterface{
					messaging_api.TextMessage{
						Text: reply_message,
					},
					messaging_api.TextMessage{
						Text: replymessage,
					},
				},
			}); err != nil {
				log.Print(err)
			}
		}
	}
}

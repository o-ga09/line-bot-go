package callback

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/o-ga09/line-bot-go/pkg/config"
	"github.com/o-ga09/line-bot-go/pkg/logger"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	logger.Logger()
	cfg, err := config.New()
	if err != nil {
		slog.Error(fmt.Sprintf("config error: %v", err))
	}
	channel_secret := cfg.LineChannelscret
	access_token := cfg.LineAccesstoken
	bot, err := messaging_api.NewMessagingApiAPI(access_token)

	var reply_message string
	if err != nil {
		slog.Info(fmt.Sprintf("can not connect line messaging api: %v", err))
	}

	cb, err := webhook.ParseRequest(channel_secret, r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			slog.Info(fmt.Sprintf("invalid signature: %v", err))
			w.WriteHeader(400)
		} else {
			slog.Error(fmt.Sprintf("can not parse request: %v", err))
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
					slog.Error(fmt.Sprintf("can not get profile: %v", err))
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
				slog.Error(fmt.Sprintf("can not reply message: %v", err))
			}
			slog.Info(fmt.Sprintf("message send: %s", reply_message))
		}
	}
}

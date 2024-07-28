# GoでLINE Messaging APIを使用してLINE Botを作成する

## Overview／概要

`line-bot-sdk-go`の`v8`を使用してLINE BotをGoで作成する

## Deploy／デプロイ方法

```
$ npm i -g vercel
$ vercel login
$ vercel --prod
```

## Test／テスト方法

```
$ go test ./...
```

## LINE Messaging APIを使用する

- ユーザー情報取得

```
curl --location 'https://api.line.me/v2/bot/profile/[userId]' \
--header 'Authorization: Bearer [access token]'
```

- プッシュ通知

```
curl --location 'https://api.line.me/v2/bot/message/push' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer [access token]' \
--data '{
    "to": "[userId]",
    "notificationDisabled": false,
    "messages": [
        {
            "type": "text",
            "text": "Hello"
        }
    ]
}'
```

## Licence

MIT @o-ga09

# telegram

Yet another [Telegram Bot API](https://core.telegram.org/bots/api#getupdates) client on golang

Made for personal usage, but feel free to reuse it (it you find something useful)

## Install

`go get -u github.com/iamdimka/go-telegram`

## Usage

```go
  bot := telegram.NewBot(BOT_TOKEN)

  // you can redefine default JSONMarshal/JSONUnmarshal functions
  bot.JSONMarshal = json.Marshal
  bot.JSONUnmarshal = json.Unmarshal

  // or redefine client
  bot.HTTPClient = http.DefaultClient

  // or redefine poll timeout
  bot.PollTimeout = 5 // default 30

  // you can call any api
  result, err := bot.GetMe()

  result2, err := bot.GetUpdates(&telegram.GetUpdatesRequest{Offset: -5})

  // or you can PollUpdates
  channel := PollUpdates(0, "message")
  for update := range channel {
    // do something with update here
  }

  // when channel has been closed, check the poll error
  err = bot.PollError()
```
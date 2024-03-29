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

  // you can call any api
  result, err := bot.GetMe()

  result2, err := bot.GetUpdates(&telegram.GetUpdatesRequest{Offset: -5})

  // or you can PollUpdates
  // it accepts different options:
  // telegram.WithAllowedUpdates(updates ...string)
  // telegram.WithOffset(int)
  // telegram.WithTimeout(int) - set the timeout in GetUpdatesRequest; default = 30
  // telegram.WithLimit(int)
  // telegram.WithClient(*http.Client) - set custom http client
  // telegram.WithContext(context.Context) - set context for each request
  channel := PollUpdates()
  for update := range channel {
    // do something with update here
  }

  // when channel has been closed, check the poll error
  err = bot.PollError()
```

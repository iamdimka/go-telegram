package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	BaseURL = "https://api.telegram.org/"
)

type Bot struct {
	url       string
	pollError error

	HTTPClient    *http.Client
	JSONMarshal   func(interface{}) ([]byte, error)
	JSONUnmarshal func([]byte, interface{}) error
}

func NewBot(token string) *Bot {
	token = strings.TrimPrefix(token, "bot")

	return &Bot{
		url:           BaseURL + "bot" + token + "/",
		HTTPClient:    http.DefaultClient,
		JSONMarshal:   json.Marshal,
		JSONUnmarshal: json.Unmarshal,
	}
}

func (b *Bot) request(method string, request interface{}, result interface{}) error {
	return b.doRequest(b.HTTPClient.Do, method, request, result)
}

func (b *Bot) doRequest(do func(req *http.Request) (*http.Response, error), method string, request interface{}, result interface{}) error {
	httpMethod := http.MethodGet
	var body io.Reader

	if request != nil {
		data, err := b.JSONMarshal(request)
		if err != nil {
			return err
		}

		body = bytes.NewReader(data)
		httpMethod = http.MethodPost
	}

	req, err := http.NewRequest(httpMethod, b.url+method, body)
	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var apiResult ApiResult
	err = b.JSONUnmarshal(data, &apiResult)
	if err != nil {
		return err
	}

	if !apiResult.Ok {
		return &apiResult
	}

	return b.JSONUnmarshal(apiResult.Result, result)
}

func (b *Bot) PollError() error {
	return b.pollError
}

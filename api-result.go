package telegram

import "encoding/json"

type ApiResult struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

func (a *ApiResult) Error() string {
	return a.Description
}

package telegram

import (
	"context"
	"net/http"
)

type pollOptions struct {
	context        context.Context
	client         *http.Client
	offset         int64
	timeout        int
	limit          int
	allowedUpdates []string
}

type pollOpt func(*pollOptions)

func WithAllowedUpdates(updates ...string) pollOpt {
	return func(po *pollOptions) {
		po.allowedUpdates = updates
	}
}

func WithOffset(offset int64) pollOpt {
	return func(po *pollOptions) {
		po.offset = offset
	}
}

func WithTimeout(timeout int) pollOpt {
	return func(po *pollOptions) {
		po.timeout = timeout
	}
}

func WithLimit(limit int) pollOpt {
	return func(po *pollOptions) {
		po.limit = limit
	}
}

func WithContext(ctx context.Context) pollOpt {
	return func(po *pollOptions) {
		po.context = ctx
	}
}

func WithClient(client *http.Client) pollOpt {
	return func(po *pollOptions) {
		po.client = client
	}
}

func (b *Bot) PollUpdates(options ...pollOpt) <-chan *Update {
	opt := &pollOptions{
		timeout: 30,
	}

	for _, fn := range options {
		fn(opt)
	}

	b.pollError = nil
	channel := make(chan *Update, 1)
	request := &GetUpdatesRequest{
		Offset:         opt.offset,
		Limit:          opt.limit,
		Timeout:        opt.timeout,
		AllowedUpdates: opt.allowedUpdates,
	}

	updates := make([]*Update, 0)

	doRequest := func(req *http.Request) (*http.Response, error) {
		client := opt.client
		if client == nil {
			client = b.HTTPClient
		}

		if opt.context != nil {
			req = req.WithContext(opt.context)
		}

		return client.Do(req)
	}

	go func() {
		for {
			err := b.doRequest(doRequest, "getUpdates", request, &updates)

			if err != nil {
				b.pollError = err
				close(channel)
				return
			}

			for _, update := range updates {
				channel <- update

				if update.UpdateId >= request.Offset {
					request.Offset = update.UpdateId + 1
				}
			}

			updates = updates[:0]
		}
	}()

	return channel
}

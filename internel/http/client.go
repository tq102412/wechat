package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Handler func(*Message)

type Client struct {
	c        *http.Client
	handlers []Handler
}

type Optional func(client *Client)

// NewClient
// 获取一个http.Client
func NewClient(options ...Optional) *Client {
	client := &Client{
		c: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       5 * time.Second,
		},
		handlers: nil,
	}

	for _, op := range options {
		op(client)
	}

	return client
}

// AddHandlers
// 添加中间件
func (c *Client) AddHandlers(handlers ...Handler) {
	c.handlers = append(c.handlers, handlers...)
}

// Request
// 发送一个http请求
func (c *Client) Request(ctx context.Context, method string, url string, body io.Reader) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)

	if nil != err {
		return nil, fmt.Errorf("new request error: %w (url: %s, body: %s, method: %s)", err, url, body, method)
	}

	m := NewMessage(SetRequest(req))

	m.handlers = append(m.handlers, c.handlers...)

	m.handlers = append(m.handlers, func(m *Message) {
		resp, err := c.c.Do(req)

		if nil != err {
			m.Error(err)
			return
		}

		m.Resp = &Response{resp}
	})

	m.Next()

	return m.Resp, m.LastError()
}

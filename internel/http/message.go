package http

import (
	"fmt"
	"math"
	"net/http"
)

const abortIndex int8 = math.MaxInt8 / 2

type Message struct {
	Request  *http.Request
	Resp     *Response
	index    int8
	handlers []Handler
	Errors   []*Error
}

// NewMessage
// 获取一个新的Message对象
func NewMessage(optionals ...MessageOp) *Message {
	message := &Message{
		Request:  nil,
		Resp:     nil,
		index:    -1,
		handlers: nil,
	}

	for _, op := range optionals {
		op(message)
	}

	return message
}

// MessageOp 可选参数
type MessageOp func(*Message)

// SetRequest
// 设置http.request 对象
func SetRequest(r *http.Request) MessageOp {
	return func(m *Message) {
		m.Request = r
	}
}

// Abort
// 设置异常，该方法不会终止程序
// 调用该方法可以阻止后续中间件的执行
func (m *Message) Abort() {
	m.index = abortIndex
}

// IsAbort
// 判断是否存在异常
func (m *Message) IsAbort() bool {
	return m.index >= abortIndex
}

// Next
// 执行后续的中间件
func (m *Message) Next() {
	m.index++
	for m.index < int8(len(m.handlers)) {
		m.handlers[m.index](m)
		m.index++
	}
}

// Error
// 用于生成错误
func (m *Message) Error(err error) {
	if nil != err {
		panic(err)
	}

	e := Error{
		err: fmt.Errorf("http request error: %w (request: %v, response: %v)", err, m.Request, m.Resp),
	}

	m.Errors = append(m.Errors, &e)
}

// LastError
// 获取最后一个错误消息
func (m *Message) LastError() *Error {
	errLen := len(m.Errors)

	if errLen == 0 {
		return nil
	}

	return m.Errors[errLen-1]
}

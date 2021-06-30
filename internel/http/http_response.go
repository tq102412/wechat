package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	*http.Response
}

//Content 获取body内容
func (rsp *Response) Content() ([]byte, error) {
	var err error

	defer func() {
		err = rsp.Body.Close()
	}()

	body, err := io.ReadAll(rsp.Body)

	if nil != err {
		return nil, err
	}

	return body, err
}

// ToJson 转换为给定数据类型
func (rsp *Response) ToJson(v interface{}) error {
	return json.NewDecoder(rsp.Body).Decode(v)
}

// Discard 丢弃body内容
func (rsp *Response) Discard() error {
	defer func() {
		_ = rsp.Body.Close()
	}()

	_, err := io.Copy(io.Discard, rsp.Body)

	return err
}

package http

import (
	"github.com/tq102412/wechatsdk"
)

func accessToken(accessToken wechatsdk.AccessTokenInterface) Handler {
	return func(message *Message) {
		token, err := accessToken.GetToken()
		tokenKey := accessToken.GetTokenKey()

		if nil != err {
			message.Error(err)
			message.Abort()
			return
		}

		query := message.Request.URL.Query()
		query.Set(tokenKey, token.AccessToken)
		message.Request.URL.RawQuery = query.Encode()
	}
}

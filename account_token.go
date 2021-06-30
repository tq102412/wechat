package wechatsdk

type AccessTokenInterface interface {
	GetToken() (*Token, error)
	Refresh() (*Token, error)
	GetTokenKey() string
}

type Token struct {
	AccessToken string
	ExpiresIn   int64
}

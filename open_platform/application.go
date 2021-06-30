package open_platform

type OpenPlatform struct {
	request   *BaseRequest
	appid     string
	appsecret string
}

func NewOpenPlatform(appid string, appsecret string, accessToken AccessTokenInterface) *OpenPlatform {
	r := NewBaseRequest()

	r.AccessToken = accessToken

	op := OpenPlatform{
		r,
		appid,
		appsecret,
	}

	return &op
}

func (op *OpenPlatform) GetAccessToken() (*Token, error) {
	return op.request.AccessToken.GetToken()
}

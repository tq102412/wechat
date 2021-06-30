package open_platform

import "log"

func (op *OpenPlatform) CreatePreAuthorizationCode() ([]byte, error) {

	url := "cgi-bin/component/api_create_preauthcode"

	json := map[string]interface{}{
		"component_appid": op.appid,
	}

	rep, err := op.request.HttpPostJson(url, HttpOption{
		json,
		nil,
	})

	return handlerRequestResult(rep, err)

}

func (op *OpenPlatform) HandleAuthorize(authCode string) ([]byte, error) {
	url := "cgi-bin/component/api_query_auth"

	json := map[string]interface{}{
		"component_appid":    op.appid,
		"authorization_code": authCode,
	}

	rep, err := op.request.HttpPostJson(url, HttpOption{
		json,
		nil,
	})

	return handlerRequestResult(rep, err)
}

func (op *OpenPlatform) GetAuthorizerToken(appid string, refreshToken string) ([]byte, error) {

	url := "cgi-bin/component/api_authorizer_token"

	json := map[string]interface{}{
		"component_appid":          op.appid,
		"authorizer_appid":         appid,
		"authorizer_refresh_token": refreshToken,
	}

	rep, err := op.request.HttpPostJson(url, HttpOption{
		json,
		nil,
	})

	return handlerRequestResult(rep, err)

}

func (op *OpenPlatform) GetAuthorizer(appid string) ([]byte, error) {

	url := "cgi-bin/component/api_get_authorizer_info"

	json := map[string]interface{}{
		"component_appid":  op.appid,
		"authorizer_appid": appid,
	}

	rep, err := op.request.HttpPostJson(url, HttpOption{
		json,
		nil,
	})

	return handlerRequestResult(rep, err)
}

func (op *OpenPlatform) GetAuthorizerOption(appid string, optionName string) ([]byte, error) {

	url := "cgi-bin/component/api_get_authorizer_option"

	json := map[string]interface{}{
		"component_appid":  op.appid,
		"authorizer_appid": appid,
		"option_name":      optionName,
	}

	rep, err := op.request.HttpPostJson(url, HttpOption{
		json,
		nil,
	})

	return handlerRequestResult(rep, err)

}

func (op *OpenPlatform) SetAuthorizerOption(appid string, optionName string, optionValue string) ([]byte, error) {

	url := "cgi-bin/component/api_set_authorizer_option"

	json := map[string]interface{}{
		"component_appid":  op.appid,
		"authorizer_appid": appid,
		"option_name":      optionName,
		"option_value":     optionValue,
	}

	rep, err := op.request.HttpPostJson(url, HttpOption{
		json,
		nil,
	})

	return handlerRequestResult(rep, err)

}

func handlerRequestResult(rep Response, err error) ([]byte, error) {
	if nil != err {
		return nil, err
	}

	return rep.GetBody()
}

func GetComponentToken(appid, appsecret string, verifyTicket VerifyTicketInterface) ([]byte, error) {
	url := "cgi-bin/component/api_component_token"

	js := make(map[string]interface{})

	js["component_appid"] = appid
	js["component_appsecret"] = appsecret
	js["component_verify_ticket"] = verifyTicket.Get(appid)

	request := NewBaseRequest()

	log.Debugf("url: %v; options: %v", url, js)

	rep, err := request.Post(url, HttpOption{
		js,
		nil,
	})

	return handlerRequestResult(rep, err)
}

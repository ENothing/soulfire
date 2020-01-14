package login

import (
	"soulfire/pkg/wechat/config"
	"soulfire/utils"
)

type Login struct{}

func (l *Login) Code2Session(code string) map[string]interface{} {

	conf := config.NewConfig()

	url := config.Code2SessionURL + "?appid=" + conf.AppId + "&secret=" + conf.Secret + "&js_code=" + code + "&grant_type=authorization_code"

	response, err := utils.HttpGet(url, []utils.Header{})
	if err != nil {
		panic(err)
	}

	res := utils.JsonDecode(string(response))

	return res
}

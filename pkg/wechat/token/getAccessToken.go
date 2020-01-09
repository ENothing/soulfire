package token

import (
	"soulfire/pkg/wechat/config"
	"soulfire/utils"
)

type Token struct{}

func (t *Token) GetAccessToken() map[string]interface{} {

	conf := config.NewConfig()

	url := config.AccessTokenURL + "&appid=" + conf.AppId + "&secret=" + conf.Secret

	response, err := utils.HttpGet(url)
	if err != nil {
		panic(err)
	}

	res := utils.JsonDecode(string(response))

	return res

}

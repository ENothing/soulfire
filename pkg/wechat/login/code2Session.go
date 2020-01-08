package login

import (
	"fmt"
	"soulfire/pkg/wechat"
	"soulfire/utils"
)

type Login struct {
}

func NewLogin() *Login {

	srv := new(Login)
	return srv
}

func (l Login) Code2Session(code string) interface{} {

	config := wechat.Config{}

	fmt.Println(config.AppId)

	url := wechat.Code2SessionURL + "?appid=" + config.AppId + "&secret=" + config.Secret + "&js_code=" + code + "&grant_type=authorization_code"

	var response []byte

	response, _ = utils.HttpGet(url)
	//if err != nil {
	//	panic(err)
	//}

	res := utils.JsonDecode(string(response))

	return res

	//str := (*string)(unsafe.Pointer(&response))
	//response[0] = 'i'
	//fmt.Println(*str)
}

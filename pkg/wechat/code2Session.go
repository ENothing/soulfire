package wechat

import (
	"fmt"
	config2 "soulfire/pkg/wechat/config"
	"soulfire/utils"
)

func Code2Session(code string) string {

	config := config2.Config{}
	fmt.Println(config.AppId)

	url := config2.Code2SessionURL + "?appid=" + config.AppId + "&secret=" + config.Secret + "&js_code=" + code + "&grant_type=authorization_code"

	var response []byte

	response, _ = utils.HttpGet(url)
	//if err != nil {
	//	panic(err)
	//}

	return fmt.Sprintf("%s", response)

	//str := (*string)(unsafe.Pointer(&response))
	//response[0] = 'i'
	//fmt.Println(*str)
}

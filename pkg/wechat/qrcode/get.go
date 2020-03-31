package qrcode

import (
	"fmt"
	"soulfire/pkg/wechat/config"
	"soulfire/pkg/wechat/token"
	"soulfire/utils"
)





func Get(page ,scene string) map[string]interface{} {


	token := token.Token{}

	tokenMap := token.GetAccessToken()

	url := config.GetWxACodeUnlimitURL + "?access_token="+ (tokenMap["access_token"]).(string)


	data := make(map[string]string)

	data["page"] = page
	data["scene"] = scene


	response, err := utils.HttpPost(url,data ,[]utils.Header{})
	if err != nil {
		panic(err)
	}

	res := utils.JsonDecode(string(response))


	fmt.Println(res)

	return res

}

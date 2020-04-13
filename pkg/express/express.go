package express

import (
	"encoding/json"
	"soulfire/pkg/config"
	"soulfire/pkg/logging"
	"soulfire/utils"
)

/**
com	STRING	必选	快递公司字母简称,可以从接口"快递公司查询" 中查到该信息 如不知道快递公司名,可以使用"auto"代替,此时将自动识别快递单号所属公司（成功率99%，因为一个单号规则可能会映射到多个快递公司。如果识别失败，系统将返回可能的快递公司列表）。不推荐大面积使用auto，建议尽量传入准确的公司编码。
nu	STRING	必选	单号
receiverPhone	STRING	可选	收件人手机号后四位【手机号后四位填一个就行，多填以寄件人为准】、【请填写正确的信息、否则会查询不到】【查询顺丰时，为了保证效率，请尽量提供寄件人或者收件人查询】
senderPhone	STRING	可选	寄件人手机号后四位【手机号后四位填一个就行，多填以寄件人为准】、【请填写正确的信息、否则会查询不到】【查询顺丰时，为了保证效率，请尽量提供寄件人或者收件人查询】
*/

type Express struct {
	ShowApiResCode  int64   `json:"showapi_res_code"`
	ShowApiResError string  `json:"showapi_res_error"`
	ShowApiResBody  ResBody `json:"showapi_res_body"`
}

type ResBody struct {
	Data        interface{} `json:"data"`
	ExpTextName string      `json:"expTextName"`
	Flag        bool        `json:"flag"`
	Status      int64       `json:"status"`
	RetCode     int64       `json:"ret_code"`
}

const ExpInfoUrl = "https://ali-deliver.showapi.com/showapi_expInfo"

func GetExpInfo(com, nu, receiverPhone, senderPhone string) *Express {

	app, _ := config.Cfg.GetSection("aliyun")

	url := ExpInfoUrl + "?com=" + com + "&nu=" + nu + "&receiverPhone=" + receiverPhone + "&senderPhone=" + senderPhone

	header := []utils.Header{
		utils.Header{
			Key:   "Authorization",
			Value: "APPCODE " + app.Key("AppCode").String(),
		},
	}

	response, err := utils.HttpGet(url, header)
	if err != nil {
		return nil
	}

	express := &Express{}

	err = json.Unmarshal(response, express)
	if err != nil {
		logging.Logging(logging.ERR, "快递信息json解析失败："+err.Error())
		return nil
	}

	if express.ShowApiResCode != int64(0) || express.ShowApiResError != "" {
		logging.Logging(logging.ERR, "快递信息接口返回失败："+express.ShowApiResError)
		return nil
	}

	return express

}

package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"soulfire/pkg/config"
	"soulfire/pkg/db"
	"soulfire/pkg/logging"
	"soulfire/utils"
	"time"
)

//const Url = "https://cxkjsms.market.alicloudapi.com/chuangxinsms/dxjk"
const Url = "https://www.fastmock.site/mock/2e0245a5857209a26b82c6a67956af38/soulfire/sms"

type Sms struct {
	mobile string
}
type ResBody struct {
	ReturnStatus  string `json:"ReturnStatus"`
	Message       string `json:"Message"`
	RemainPoint   int64  `json:"RemainPoint"`
	TaskID        int64  `json:"TaskID"`
	SuccessCounts int64  `json:"SuccessCounts"`
}

func New(mobile string) *Sms {
	return &Sms{mobile: mobile}
}

func (s *Sms) SendCode() error {

	app, _ := config.Cfg.GetSection("aliyun")

	code := utils.Code()

	header := []utils.Header{
		utils.Header{
			Key:   "Authorization",
			Value: "APPCODE " + app.Key("AppCode").String(),
		},
	}

	data := make(map[string]interface{})
	data["mobile"] = s.mobile
	data["content"] = "【SoulFire】你的验证码是："+code+"，10分钟内有效！"

	res, err := utils.HttpPost(Url, data, header)
	if err != nil {
		logging.Logging(logging.ERR,"短信接口请求失败原因："+err.Error())
		return err
	}

	resBody := ResBody{}

	err = json.Unmarshal(res, &resBody)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if resBody.ReturnStatus != "Success"{
		logging.Logging(logging.ERR,"短信发送失败原因："+resBody.Message)
		return errors.New(resBody.Message)
	}


	db.RedisDb.Set(s.mobile, code, 10*time.Minute)

	return nil

}

func (s *Sms) GetCode() (string, error) {
	return db.RedisDb.Get(s.mobile).Result()
}

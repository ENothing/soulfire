package sms

import (
	"encoding/json"
	"soulfire/pkg/config"
	"soulfire/pkg/db"
	"soulfire/utils"
	"time"
)

const Url = "http(s)://cxkjsms.market.alicloudapi.com/chuangxinsms/dxjk"

type Sms struct {
	mobile string
}
type Response struct {

}

func New(mobile string) *Sms {
	return &Sms{mobile: mobile}
}

func (s *Sms) SendCode()(  error){

	app, _ := config.Cfg.GetSection("aliyun")

	code := utils.Code()

	header := []utils.Header{
		utils.Header{
			Key:   "Authorization",
			Value: "APPCODE " + app.Key("AppCode").String(),
		},
	}

	_, err := utils.HttpPost(Url,nil, header)
	if err != nil {
		return err
	}




	//err = json.Unmarshal(res,Response)


	db.RedisDb.Set(s.mobile,code,10*time.Minute)


}

func (s *Sms) GetCode() (string, error) {
	return db.RedisDb.Get(s.mobile).Result()
}

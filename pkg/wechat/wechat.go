package wechat

import (
	"soulfire/pkg/config"
	"soulfire/pkg/wechat/login"
)

type Config struct {
	AppId  string
	Secret string
	MchId  string
	Key    string
}

const (
	Code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session"
	AccessTokenURL  = "https://api.weixin.qq.com/cgi-bin/token"
	TemplateSendURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send"
)

func init() {

	app, _ := config.Cfg.GetSection("wechat")

	weConfig := &Config{}
	weConfig.AppId = app.Key("APPID").String()
	weConfig.Secret = app.Key("SECRET").String()
	weConfig.MchId = app.Key("MCHID").String()
	weConfig.Key = app.Key("KEY").String()

}

func Login() *login.Login {
	return login.NewLogin()
}

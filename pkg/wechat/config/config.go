package config

import (
	"soulfire/pkg/config"
)

type Config struct {
	AppId  string
	Secret string
	MchId  string
	Key    string
}

const (
	Code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session"
	AccessTokenURL  = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"
	TemplateSendURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send"
	GetWxACodeUnlimitURL = "https://api.weixin.qq.com/wxa/getwxacodeunlimit"
)

func NewConfig() *Config {
	app, _ := config.Cfg.GetSection("wechat")
	return &Config{
		AppId:  app.Key("APPID").String(),
		Secret: app.Key("SECRET").String(),
		MchId:  app.Key("MCHID").String(),
		Key:    app.Key("MCHID").String(),
	}
}

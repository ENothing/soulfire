package wechat

import (
	"soulfire/pkg/config"
)

var Base Wechat


type Wechat struct {
	Appid string
	Secret string
	Mchid string
	Key string
}

func init()  {

	app, _   := config.Cfg.GetSection("wechat")

	base = Wechat{}
	base.Appid = app.Key("APPID").String()
	base.Secret = app.Key("SECRET").String()
	base.Mchid = app.Key("MCHID").String()
	base.Key = app.Key("KEY").String()

}



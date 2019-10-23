package wechat

import (
	"gin-init/pkg/config"
	"gin-init/utils"
)

var (
	app, _  = config.Cfg.GetSection("wechat")
	appid = app.Key("APPID").String()
	secret = app.Key("SECRET").String()
)

const base_url   = "https://api.weixin.qq.com"

func Code2Session(code string)( data map[string]string){

	url := base_url + "/sns/jscode2session?appid="+appid+"&secret="+secret+"&js_code="+code+"&grant_type=authorization_code"

	data = utils.HttpGet(url)

	return

}
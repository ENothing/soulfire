package wechat

import "soulfire/utils"

type Op struct {
	Wechat
	BaseUrl string
}

func init()  {

	initOp := Op{}
	initOp.BaseUrl = "https://api.weixin.qq.com"

}

func Code2Session(code string)( data map[string]string){

	op := Op{}

	url := op.BaseUrl + "/sns/jscode2session?appid="+op.Appid+"&secret="+op.Secret+"&js_code="+code+"&grant_type=authorization_code"

	data = utils.HttpGet(url)

	return

}
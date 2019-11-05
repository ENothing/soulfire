package wechat

import "fmt"

type Pay struct {
	Wechat
	BaseUrl string
}
func init()  {

	initPay := Op{}
	initPay.BaseUrl = "https://api.mch.weixin.qq.com"

}

func Unifiedorder()  {

	pay := Pay{}

	fmt.Println(pay.Appid)
	fmt.Println(pay.BaseUrl)

}
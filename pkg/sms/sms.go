package sms

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"soulfire/utils"
	"time"
)

const Url = "http(s)://cxkjsms.market.alicloudapi.com/chuangxinsms/dxjk"

type Sms struct {
	mobile string
	cache  *cache.Cache
}

func NewSms(mobile string) *Sms {
	return &Sms{mobile: mobile, cache: cache.New(30*time.Second, 20*time.Minute)}
}

func (s *Sms) SendCode() {

	code := utils.Code()

	fmt.Println(code)
	s.cache.Set(s.mobile, code, cache.DefaultExpiration)

}

func (s *Sms)GetCode() (interface{},bool ){
	fmt.Println(s.mobile)

	return s.cache.Get(s.mobile)

}

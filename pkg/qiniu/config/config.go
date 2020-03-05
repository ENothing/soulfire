package config

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"soulfire/pkg/config"
)

type Config struct {
	AccessKey   string
	SecretKey   string
	MediaUrl string
}

func GetQiniuConfig() *Config {

	app, _ := config.Cfg.GetSection("qiniu")

	accessKey := app.Key("AccessKey").String()
	secretKey := app.Key("SecretKey").String()

	return &Config{
		AccessKey:   accessKey,
		SecretKey:   secretKey,
	}

}

func NewConfig() *storage.Config {

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	return &cfg

}

func NewMac() *qbox.Mac {

	conf := GetQiniuConfig()

	mac := qbox.NewMac(conf.AccessKey, conf.SecretKey)

	return mac

}

func UpToken(bucket string) string {

	mac := NewMac()

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	return putPolicy.UploadToken(mac)

}

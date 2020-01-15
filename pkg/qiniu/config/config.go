package config

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"soulfire/pkg/config"
)

func NewConfig() *storage.Config {

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	return &cfg

}

func UpToken(bucket string) string {

	app, _ := config.Cfg.GetSection("qiniu")

	accessKey := app.Key("AccessKey").String()
	secretKey := app.Key("SecretKey").String()

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := qbox.NewMac(accessKey, secretKey)

	return putPolicy.UploadToken(mac)

}

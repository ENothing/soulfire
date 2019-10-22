package config

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg     *ini.File
	Runmode string
)

func init() {

	var err error
	Cfg, err = ini.Load("conf/conf")
	if err != nil {
		log.Printf("配置加载失败：%s", err)
	}
	RunMode()
}

func RunMode() {
	Runmode = Cfg.Section("").Key("RUN_MODE").MustString("RUN_MODE")
}

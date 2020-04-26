package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
}

func Conf() *Config {
	return &Config{}
}

func (c *Config) Init() {

	viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		panic(err)
	}

	c.watch()

}

func (c *Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}

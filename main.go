package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"soulfire/pkg/config"
	"soulfire/pkg/db"
	"soulfire/router"
)

func main() {

	config.Conf().Init()

	db.DB.Init()
	defer db.DB.Close()

	db.RedisInit()

	fmt.Println(viper.GetString("App.Mode"))

	gin.SetMode(viper.GetString("App.Mode"))

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(g, middlewares...)

	log.Printf("服务开始运行：%s", ":8081")
	log.Printf(http.ListenAndServe(":8081", g).Error())

}

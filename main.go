package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"soulfire/pkg/config"
	"soulfire/pkg/db"
	"soulfire/router"
)

func main() {

	db.DB.Init()
	defer db.DB.Close()

	//db.RedisInit()

	gin.SetMode(config.Runmode)

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(g, middlewares...)

	log.Printf("服务开始运行：%s", ":8081")
	log.Printf(http.ListenAndServe(":8081", g).Error())

}

package main

import (
	"soulfire/pkg/config"
	"soulfire/pkg/db"
	"soulfire/pkg/logging"
	"soulfire/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	db.DB.Init()
	defer db.DB.Close()

	gin.SetMode(config.Runmode)

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(g,middlewares...)


	logging.Logging(logging.INFO,"test")
	log.Printf("服务开始运行：%s",":8080")
	log.Printf(http.ListenAndServe(":8080",g).Error())

}

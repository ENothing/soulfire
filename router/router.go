package router

import (
	"gin-init/controllers/app/v1/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(mw...)

	g.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, "接口路由不存在~")
	})

	app := g.Group("app/v1/")
	{

		u := app.Group("user/")
		{
			u.POST("login",user.Login)
		}


	}


	return g

}

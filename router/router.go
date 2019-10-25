package router

import (
	"gin-init/controllers/app/v1/activity"
	"gin-init/controllers/app/v1/user"
	"gin-init/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(middleware.Options)
	g.Use(mw...)

	g.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, "接口路由不存在~")
	})

	app := g.Group("app/v1")
	{

		u := app.Group("user")
		{
			u.POST("login",user.Login)
		}

		mu := app.Group("user").Use(middleware.Verify())
		{
			mu.GET("")
		}


		a := app.Group("activity")
		{
			a.GET("index",activity.Index)
			a.GET("list",activity.List)
			a.GET("detail/:id",activity.Detail)

		}



	}


	return g

}

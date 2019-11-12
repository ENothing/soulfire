package router

import (
	"soulfire/controllers/app/v1/activity"
	"soulfire/controllers/app/v1/bbs"
	"soulfire/controllers/app/v1/user"
	"soulfire/router/middleware"
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
			a.GET("list",activity.ActivityList)
			a.GET("detail/:id",activity.Detail)
			a.GET("cates",activity.ActivityCates)

		}
		ma := app.Group("activity").Use(middleware.Verify())
		{
			ma.GET("like/:id",activity.Like)
			ma.POST("enter",activity.Enter)
			ma.GET("order/:id",activity.OrderDetail)
			ma.POST("pay",activity.Pay)
			ma.GET("order_list",activity.ActivityOrderList)

		}

		b := app.Group("bbs")
		{
			b.GET("list",bbs.ArticleList)
			b.GET("detail",bbs.Detail)
		}
		mb := app.Group("bbs").Use(middleware.Verify())
		{
			mb.GET("like/:id",bbs.Like)


			mb.POST("post_comment",bbs.PostComment)
		}
	}


	return g

}

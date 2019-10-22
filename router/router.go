package router

import (
	admin_v1_user "gin-init/controllers/admin/v1/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(mw...)

	g.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, "接口路由不存在~")
	})


	admin := g.Group("admin/v1/")
	{
		u := admin.Group("user")
		{

			u.GET("",admin_v1_user.Index)
			u.GET("create/:username/:password",admin_v1_user.Create)
			u.GET("edit/:id",admin_v1_user.Edit)
			u.GET("update/:id",admin_v1_user.Update)
			u.GET("delete/:id",admin_v1_user.Delete)
		}

	}


	app := g.Group("app/v1/")
	{

	}




	return g

}

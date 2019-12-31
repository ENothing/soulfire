package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"soulfire/controllers/app/v1/activity"
	"soulfire/controllers/app/v1/address"
	"soulfire/controllers/app/v1/bbs"
	"soulfire/controllers/app/v1/coupon"
	"soulfire/controllers/app/v1/shop"
	"soulfire/controllers/app/v1/user"
	"soulfire/router/middleware"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(middleware.Options)
	g.Use(mw...)

	g.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, "访问地址不存在~")
	})

	app := g.Group("app/v1")
	{

		u := app.Group("user")
		{
			u.POST("login", user.Login) //登录
		}
		mu := app.Group("user").Use(middleware.Verify())
		{
			mu.GET("")
		}

		a := app.Group("activity")
		{
			a.GET("index", activity.Index)         //活动首页头部
			a.GET("list", activity.ActivityList)   //活动列表
			a.GET("detail/:id", activity.Detail)   //活动详情
			a.GET("cates", activity.ActivityCates) //活动分类

		}
		ma := app.Group("activity").Use(middleware.Verify())
		{
			ma.GET("like/:id", activity.Like)                //用户like
			ma.POST("enter", activity.Enter)                 //活动报名
			ma.GET("order/:id", activity.OrderDetail)        //活动订单详情
			ma.POST("pay", activity.Pay)                     //活动支付
			ma.GET("order_list", activity.ActivityOrderList) //活动订单列表

		}

		b := app.Group("bbs")
		{

			b.GET("user_articles/:user_id", bbs.UserArticleList) //用户个人文章列表

			b.GET("comment_list", bbs.CommentList) //评论列表
		}
		mb := app.Group("bbs").Use(middleware.Verify())
		{
			mb.GET("like/:id", bbs.Like) //文章like

			mb.POST("publish_article", bbs.PublishArticle) //发布文章
			mb.POST("edit_article", bbs.EditArticle)       //编辑文章
			mb.POST("del_article", bbs.DeleteArticle)      //删除文章

			mb.POST("post_comment", bbs.PostComment) //发表评论
			mb.POST("follow", bbs.Follow)            //用户关注

			mb.GET("list", bbs.ArticleList)                //文章列表
			mb.GET("user_detail/:user_id", bbs.UserDetail) //用户个人页面信息
			mb.GET("detail/:id", bbs.Detail)               //文章详情
		}

		s := app.Group("shop")
		{
			s.GET("index", shop.Index)                         //商城首页上部分
			s.GET("list", shop.GoodsList)                      //商品列表
			s.GET("goods_detail/:id", shop.GoodsDetail)        //商品详情
			s.GET("purchasers/:id", shop.PurchasersList)       //商品购买人列表
			s.GET("cates_with_brands", shop.CateWithBrandList) //分类列表
		}

		ms := app.Group("shop").Use(middleware.Verify())
		{
			ms.GET("pre_order_detail/:goods_spu_id", shop.PreOrderDetail) //预订单详情
			ms.POST("buy", shop.Buy)                                      //下单
			ms.GET("order_detail/:order_id", shop.OrderDetail)            //订单详情
		}

		mad := app.Group("address").Use(middleware.Verify())
		{
			mad.GET("list", address.AddressList)                        //地址列表
			mad.POST("add", address.AddAddress)                         //添加地址
			mad.POST("update/:id", address.UpdateAddress)               //更新地址
			mad.GET("update_default/:id", address.UpdateDefaultAddress) //修改默认地址
			mad.GET("detail/:id", address.Detail)                       //地址详情
			mad.GET("del/:id", address.DelAddress)                      //删除地址

		}
		mc := app.Group("coupon").Use(middleware.Verify())
		{
			mc.POST("can_use_coupons", coupon.CanUseCouponsList) //可用优惠券列表
			mc.POST("user_coupons", coupon.UserCouponList)       //用户优惠券列表
		}
	}

	return g

}

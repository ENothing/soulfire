package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"soulfire/controllers/app/v1/activity"
	"soulfire/controllers/app/v1/address"
	"soulfire/controllers/app/v1/bbs"
	"soulfire/controllers/app/v1/coupon"
	"soulfire/controllers/app/v1/shop"
	"soulfire/controllers/app/v1/temp"
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
			mu.POST("feedback", user.PostFeedback)         //提交反馈
			mu.GET("per_collect", user.PersonalCollection) //用户收藏列表
			mu.POST("upload", user.Upload) //用户收藏列表
			mu.GET("info", user.Info) //用户收藏列表
			mu.GET("follow_list", user.FollowsList) //用户收藏列表
			mu.GET("followed_list", user.FollowedList) //用户收藏列表
			mu.GET("article", user.MyArticleList) //用户收藏列表
			//mu.GET("index/:id", user.Index) //用户个人主页
		}

		a := app.Group("activity")
		{
			a.GET("index", activity.Index)         //活动首页头部
			a.GET("detail/:id", activity.Detail)   //活动详情
			a.GET("cates", activity.ActivityCates) //活动分类
			a.GET("dynamic_history", activity.DynamicHistory) //动态获取相似搜索内容
			a.POST("send_sms", activity.SendSms) //发送短信

		}
		ma := app.Group("activity").Use(middleware.Verify())
		{
			ma.GET("list", activity.ActivityList)   //活动列表
			ma.GET("like/:id", activity.Like)                //用户like
			ma.POST("enter", activity.Enter)                 //活动报名
			ma.GET("order/:id", activity.OrderDetail)        //活动订单详情
			ma.GET("finish/:id", activity.FinishOrder)        //订单确认完成

			ma.POST("pay", activity.Pay)                     //活动支付
			ma.GET("order_list", activity.ActivityOrderList) //活动订单列表

			ma.GET("favor/:id", activity.Favor) //收藏活动

			ma.GET("search_history", activity.GetHistory) //获取历史和热门搜索
			ma.GET("del_search_history", activity.DelSearchHistory) //删除历史搜索记录

			ma.POST("initiate_refund", activity.InitiateRefund) //发起退款

		}

		b := app.Group("bbs")
		{

			b.GET("user_articles/:user_id", bbs.UserArticleList) //用户个人文章列表
			b.GET("user_like_articles/:user_id", bbs.UserLikeArticleList) //用户个人文章列表

			b.GET("comment_list", bbs.CommentList) //评论列表

			b.GET("article_cate", bbs.ArticleCateList) //文章分类

		}
		mb := app.Group("bbs").Use(middleware.Verify())
		{
			mb.GET("like/:id", bbs.Like) //文章like

			mb.POST("publish_article", bbs.PublishArticle) //发布文章
			mb.POST("edit_article/:id", bbs.EditArticle)       //编辑文章
			mb.GET("del_article/:id", bbs.DeleteArticle)      //删除文章
			mb.GET("edit_article_detail/:id", bbs.ArticleEditDetail)      //删除文章
			mb.GET("publish/:id", bbs.UpdateArticleToPublish)      //删除文章

			mb.POST("post_comment", bbs.PostComment) //发表评论
			mb.GET("follow/:follow_id", bbs.Follow)            //用户关注
			mb.GET("favor/:id", bbs.Favor)           //收藏活动

			mb.GET("list", bbs.ArticleList)                //文章列表
			mb.GET("user_detail/:user_id", bbs.UserDetail) //用户个人页面信息
			mb.GET("detail/:id", bbs.Detail)               //文章详情
			mb.GET("follows_List/:id", bbs.FollowsList)               //文章详情
			mb.GET("followed_List/:id", bbs.FollowedList)               //文章详情

			mb.POST("upload", bbs.Upload) //文章详情

		}

		s := app.Group("shop")
		{
			s.GET("index", shop.Index)                         //商城首页上部分
			s.GET("goods_detail/:id", shop.GoodsDetail)        //商品详情
			s.GET("purchasers/:id", shop.PurchasersList)       //商品购买人列表
			s.GET("cates_with_brands", shop.CateWithBrandList) //分类列表
			s.GET("dynamic_history", shop.DynamicHistory) //动态获取相似搜索内容
			s.POST("share", shop.Share) //动态获取相似搜索内容

		}

		ms := app.Group("shop").Use(middleware.Verify())
		{
			ms.GET("list", shop.GoodsList)                      //商品列表



			ms.GET("pre_order_detail/:goods_spu_id", shop.PreOrderDetail) //预订单详情
			ms.GET("order_detail/:order_id", shop.OrderDetail)            //订单详情

			ms.POST("buy", shop.Buy) //下单

			ms.POST("order_cancel", shop.CancelOrder)        //取消订单
			ms.POST("initiate_refund", shop.InitiateRefund)  //发起退款
			ms.POST("post_return_info", shop.PostReturnInfo) //填写退款单号
			ms.GET("order_list", shop.OrderList)             //订单列表

			ms.POST("pay", shop.Pay) //支付

			ms.GET("exp_info", shop.GetExpress) //快递查询


			ms.GET("search_history", shop.GetHistory) //获取历史和热门搜索

			ms.GET("del_search_history", shop.DelSearchHistory) //删除历史搜索记录


			ms.POST("upload", shop.Upload) //退款上传图片



		}

		mad := app.Group("address").Use(middleware.Verify())
		{
			mad.GET("list", address.AddressList)                        //地址列表
			mad.POST("add", address.AddAddress)                         //添加地址
			mad.POST("update/:id", address.UpdateAddress)               //更新地址
			mad.GET("update_default/:id", address.UpdateDefaultAddress) //修改默认地址
			mad.GET("detail/:id", address.Detail)                       //地址详情
			mad.GET("detail_to_order", address.DetailToOrder)                       //地址详情
			mad.GET("del/:id", address.DelAddress)                      //删除地址

		}
		mc := app.Group("coupon").Use(middleware.Verify())
		{
			mc.GET("can_use_coupons", coupon.CanUseCouponsList) //可用优惠券列表
			mc.GET("user_coupons", coupon.UserCouponList)       //用户优惠券列表
		}



		t := app.Group("tool")
		{
			t.POST("upload",temp.Upload)
		}


	}

	return g

}

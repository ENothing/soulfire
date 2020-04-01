package shop

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Pay(ctx *gin.Context) {

	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)
	orderId, _ := strconv.ParseInt(ctx.PostForm("order_id"), 10, 64)
	data := make(map[string]interface{})

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	order, err := models.GetOrderById(userId, orderId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	data["out_trade_no"] = order.OrderN
	data["total_amount"] = order.RealPrice
	data["body"] = order.GoodsName + "(" + order.GoodsSpuName + ")"
	data["openid"] = ""

	//todo 调起支付

	//todo 获取支付回调

}

func Notify(ctx *gin.Context) {

	//todo 支付回调

}

package shop

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/express"
	"soulfire/pkg/rsp"
	"strconv"
)

func GetExpress(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	orderId, _ := strconv.ParseInt(ctx.PostForm("order_id"), 10, 64)

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	_, err := models.GetOrderById(userId, orderId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderNotExits, nil, "")
		return
	}

	delivery, err := models.GetDeliveryById(orderId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ExpressGetFailed, nil, "")
		return
	}

	//abbreviation := "yuantong"
	//deliveryN := "YT4282966663164"
	//expressName := "圆通速递"
	abbreviation := delivery.Abbreviation
	deliveryN := delivery.DeliveryN
	expressName := delivery.Name

	expInfo := express.GetExpInfo(abbreviation, deliveryN, "", "")

	if expInfo == nil {
		rsp.JsonResonse(ctx, rsp.ExpressGetFailed, nil, "")
		return
	}
	if expInfo.ShowApiResBody.Flag == false || expInfo.ShowApiResBody.RetCode != int64(0) {
		rsp.JsonResonse(ctx, rsp.ExpressGetFailed, nil, "")
		return
	}

	data := make(map[string]interface{})
	data["path"] = expInfo.ShowApiResBody.Data
	data["delivery_n"] = deliveryN
	data["expressName"] = expressName

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

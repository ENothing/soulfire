package activity

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"soulfire/utils"
	"strconv"
	"time"
)

func InitiateRefund(ctx *gin.Context)  {

	userId,_ := ctx.MustGet("user_id").(int64)
	orderId, _ := strconv.ParseInt(ctx.PostForm("order_id"), 10, 64)
	reason := ctx.PostForm("reason")


	nowTime := utils.TimeFormat(time.Now(), 0)

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	if reason == "" {
		rsp.JsonResonse(ctx, rsp.ActivityOrderRefundReasonEmpty, nil, "")
		return
	}

	order, err := models.GetActivityOrderById(userId, orderId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderNotExits, nil, "")
		return
	}

	//待付款或订单取消
	if order.Status == 0 || order.Status == 1 {
		rsp.JsonResonse(ctx, rsp.ActivityOrderRefundRejected, nil, "")
		return
	}

	if order.RefundId != 0 {
		rsp.JsonResonse(ctx, rsp.ReActivityOrderRefund, nil, "")
		return
	}


	//已完成
	if order.Status == 3 {
		duringDays := utils.BetweenDays(nowTime, order.CompletedAtFormat)
		if duringDays > int64(5) {
			rsp.JsonResonse(ctx, rsp.ActivityOrderRefundTimeOver, nil, "")
			return
		}
	}

	activityRefund := models.ActivityOrderRefund{
		RefundN:utils.Uid("RO"),
		OrderId:orderId,
		Price:order.RealPrice,
		Reason:reason,
		UserId:userId,
	}

	refundId,err := activityRefund.Create()
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ActivityOrderRefundFailed, nil, "")
		return
	}

	err = models.UpdateActivityOrderRefundId(userId,orderId,refundId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ActivityOrderRefundFailed, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}


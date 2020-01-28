package coupon

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func CanUseCouponsList(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	goodsId, _ := strconv.ParseInt(ctx.Query("goods_id"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	data := make(map[string]interface{})

	userCoupons, total, lastPage, err := models.CanUseCouponsPaginate(page, pageSize, userId, goodsId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.CouponsListNotExits, nil, "")
		return
	}

	data["coupons"] = userCoupons
	data["total"] = total
	data["lastPage"] = lastPage

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func UserCouponList(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	status, _ := strconv.ParseInt(ctx.Query("status"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	data := make(map[string]interface{})

	userCoupons, total, lastPage, err := models.UserCouponsPaginate(page, pageSize, userId, status)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.UserCouponsListNotExits, nil, "")
		return
	}

	data["coupons"] = userCoupons
	data["total"] = total
	data["lastPage"] = lastPage

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

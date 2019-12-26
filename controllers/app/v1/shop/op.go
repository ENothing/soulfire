package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Buy(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	//num, _ := strconv.ParseInt(ctx.DefaultPostForm("num","0"), 10, 64)
	//goodsSpuId,_ := strconv.ParseInt(ctx.PostForm("goods_spu_id"), 10, 64)
	goodsId, _ := strconv.ParseInt(ctx.PostForm("goods_id"), 10, 64)
	//addressId,_ := strconv.ParseInt(ctx.PostForm("address_id"), 10, 64)
	couponId, _ := strconv.ParseInt(ctx.PostForm("coupon_id"), 10, 64)

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	//shipAddress,err := models.GetAddressById(addressId,userId)
	//if err == gorm.ErrRecordNotFound {
	//	rsp.JsonResonse(ctx, rsp.AddressNotExits, nil, "")
	//	return
	//}

	userCoupon, err := models.GetUserCouponById(userId, goodsId, couponId)
	if err == gorm.ErrRecordNotFound {
		rsp.JsonResonse(ctx, rsp.CouponNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.CouponNotExits, userCoupon, "")

	//shopOrder := models.ShopOrder{}
	//
	//
	//shopOrder.Create(userId,num, goodsSpuId, addressId, couponId)

}

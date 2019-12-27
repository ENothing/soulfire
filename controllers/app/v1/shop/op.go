package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"soulfire/models"
	"soulfire/pkg/db"
	"soulfire/pkg/rsp"
	"soulfire/utils"
	"strconv"
)

func Buy(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	num, _ := strconv.ParseInt(ctx.DefaultPostForm("num", "1"), 10, 64)
	goodsSpuId, _ := strconv.ParseInt(ctx.PostForm("goods_spu_id"), 10, 64)
	goodsId, _ := strconv.ParseInt(ctx.PostForm("goods_id"), 10, 64)
	addressId, _ := strconv.ParseInt(ctx.PostForm("address_id"), 10, 64)
	couponId, _ := strconv.ParseInt(ctx.PostForm("coupon_id"), 10, 64)
	var realPrice float64
	var discountPrice float64

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	shipAddress, err := models.GetAddressById(addressId, userId)
	if err == gorm.ErrRecordNotFound || err != nil {
		rsp.JsonResonse(ctx, rsp.AddressNotExits, nil, "")
		return
	}

	userCoupon, _ := models.GetUserCouponById(userId, goodsId, couponId)

	goodsSpu, err := models.GetGoodsSpuById(goodsSpuId)
	if err == gorm.ErrRecordNotFound || err != nil {
		rsp.JsonResonse(ctx, rsp.GoodsNotExits, nil, "")
		return
	}
	if goodsSpu.Stock < num {
		rsp.JsonResonse(ctx, rsp.ShopGoodsNotEnough, nil, "")
		return
	}

	totalPrice := goodsSpu.Price * float64(num)

	//优惠类型 1：满减 2：立减 3：打折
	switch userCoupon["coupon_type"] {
	case 1:
		if totalPrice >= (userCoupon["full_price"]).(float64) {
			realPrice = totalPrice - (userCoupon["reduction_price"]).(float64)
			discountPrice = (userCoupon["reduction_price"]).(float64)
		}
		break
	case 2:
		realPrice = totalPrice - (userCoupon["immediately_price"]).(float64)
		discountPrice = (userCoupon["immediately_price"]).(float64)
		break
	case 3:
		realPrice = totalPrice * (userCoupon["discount"]).(float64)
		discountPrice = totalPrice - realPrice
		break
	}

	orderRealPrice := realPrice + goodsSpu.PostPrice

	/* 订单创建事务-START */
	transaction := db.DB.Self.Begin()

	shopOrder := models.ShopOrder{
		UserId:        userId,
		OrderN:        utils.Uid("SF"),
		UserCouponId:  couponId,
		Num:           num,
		UnitPrice:     goodsSpu.Price,
		TotalPrice:    totalPrice,
		RealPrice:     orderRealPrice,
		DiscountPrice: discountPrice,
		PostPrice:     goodsSpu.PostPrice,
		Status:        int64(0),
		Name:          shipAddress.Name,
		Mobile:        shipAddress.Mobile,
		Province:      shipAddress.Province,
		City:          shipAddress.City,
		District:      shipAddress.District,
		DetailAddress: shipAddress.DetailAddress,
	}

	orderId, err := shopOrder.Create(transaction)
	if err != nil {
		transaction.Rollback()
		rsp.JsonResonse(ctx, rsp.ShopOrderCreateFailed, nil, "")
		return
	}
	transaction.Rollback()
	return

	//shopOrderGoods := models.ShopOrderGoods{
	//	OrderId:       orderId,
	//	GoodsId:       goodsSpu.GoodsId,
	//	Num:           num,
	//	UnitPrice:     goodsSpu.Price,
	//	TotalPrice:    totalPrice,
	//	RealPrice:     realPrice,
	//	DiscountPrice: discountPrice,
	//}
	//
	//err = shopOrderGoods.Create()
	//if err != nil {
	//	transaction.Rollback()
	//	rsp.JsonResonse(ctx, rsp.ShopOrderCreateFailed, nil, "")
	//	return
	//}

	///* 减库存-START */
	//
	//err = models.CutGoodsSpuStock(goodsSpuId, num)
	//if err != nil {
	//	transaction.Rollback()
	//	rsp.JsonResonse(ctx, rsp.ShopGoodsNotEnough, nil, "")
	//	return
	//}
	//
	//err = models.CutGoodsStockAndAddSold(goodsId, num)
	//if err != nil {
	//	transaction.Rollback()
	//	rsp.JsonResonse(ctx, rsp.ShopGoodsNotEnough, nil, "")
	//	return
	//}
	//
	///* 减库存-END */

	transaction.Commit()
	/* 订单创建事务-END */

	rsp.JsonResonse(ctx, rsp.OK, orderId, "")
}

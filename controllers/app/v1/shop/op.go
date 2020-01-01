package shop

import (
	"fmt"
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

	userCoupon, userCouponErr := models.GetUserCouponById(userId, goodsId, couponId)

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
	case int64(1):
		if totalPrice >= (userCoupon["full_price"]).(float64) {
			fmt.Println(123)
			realPrice = totalPrice - (userCoupon["reduction_price"]).(float64)
			fmt.Println(realPrice)

			if realPrice < float64(0) {
				realPrice = float64(0)
			}
			discountPrice = (userCoupon["reduction_price"]).(float64)
		}
		break
	case int64(2):
		realPrice = totalPrice - (userCoupon["immediately_price"]).(float64)
		if realPrice < float64(0) {
			realPrice = float64(0)
		}
		discountPrice = (userCoupon["immediately_price"]).(float64)
		break
	case int64(3):
		realPrice = totalPrice * (userCoupon["discount"]).(float64)
		discountPrice = totalPrice - realPrice
		if discountPrice < float64(0) {
			discountPrice = float64(0)
		}
		break
	default:
		realPrice = totalPrice
		discountPrice = float64(0)
	}

	orderRealPrice := realPrice + goodsSpu.PostPrice

	/* 订单创建事务-START */
	transaction := db.DB.Self.Begin()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	shopOrder := models.ShopOrderCreateForm{
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

	shopOrderGoods := models.ShopOrderGoodsCreateForm{
		UserId:        userId,
		OrderId:       orderId,
		GoodsId:       goodsSpu.GoodsId,
		Num:           num,
		UnitPrice:     goodsSpu.Price,
		TotalPrice:    totalPrice,
		RealPrice:     realPrice,
		DiscountPrice: discountPrice,
		SpuId:         goodsSpuId,
	}

	err = shopOrderGoods.Create(transaction)
	if err != nil {
		transaction.Rollback()
		rsp.JsonResonse(ctx, rsp.ShopOrderCreateFailed, nil, "")
		return
	}

	/* 减库存-START */

	err = models.CutGoodsSpuStock(goodsSpuId, num, transaction)
	if err != nil {
		transaction.Rollback()
		rsp.JsonResonse(ctx, rsp.ShopGoodsNotEnough, nil, "")
		return
	}

	err = models.CutGoodsStockAndAddSold(goodsId, num, transaction)
	if err != nil {
		transaction.Rollback()
		rsp.JsonResonse(ctx, rsp.ShopGoodsNotEnough, nil, "")
		return
	}

	/* 减库存-END */

	/* 修改用户优惠券状态-START */
	if userCouponErr != gorm.ErrRecordNotFound && userCouponErr == nil {
		err = models.UpdateUserCouponIsUsed(userId, couponId, transaction)
		if err != nil {
			transaction.Rollback()
			rsp.JsonResonse(ctx, rsp.ShopOrderCreateFailed, nil, "")
			return
		}
	}
	/* 修改用户优惠券状态-END */

	transaction.Commit()
	/* 订单创建事务-END */

	rsp.JsonResonse(ctx, rsp.OK, orderId, "")
}

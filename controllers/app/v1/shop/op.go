package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
	"path"
	"soulfire/models"
	"soulfire/pkg/config"
	"soulfire/pkg/db"
	"soulfire/pkg/qiniu"
	"soulfire/pkg/rsp"
	"soulfire/utils"
	"strconv"
	"time"
)

/**
下单
*/
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
		OrderN:        utils.Uid("SH"),
		UserCouponId:  couponId,
		Num:           num,
		UnitPrice:     goodsSpu.Price,
		TotalPrice:    totalPrice,
		RealPrice:     orderRealPrice,
		DiscountPrice: discountPrice,
		PostPrice:     goodsSpu.PostPrice,
		Status:        models.PendingPay,
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

	//todo 队列:两小时关闭订单

	rsp.JsonResonse(ctx, rsp.OK, orderId, "")
}

/**
取消订单
*/
func CancelOrder(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	orderId, _ := strconv.ParseInt(ctx.PostForm("order_id"), 10, 64)

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	err := models.UpdateOrderStatusToCancel(userId, orderId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderCancelFailed, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

/**
发起退款（待审核）
*/
func InitiateRefund(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	orderId, _ := strconv.ParseInt(ctx.PostForm("order_id"), 10, 64)
	reason := ctx.PostForm("reason")
	rType, _ := strconv.ParseInt(ctx.DefaultPostForm("r_type", "1"), 10, 64)
	reasonPics := ctx.PostForm("imgs")

	nowTime := utils.TimeFormat(time.Now(), 0)

	if reason == "" {
		rsp.JsonResonse(ctx, rsp.ShopOrderRefundReasonEmpty, nil, "")
		return
	}

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	//hasShopOrderRefund, err := models.GetShopOrderRefundByOrderId(userId, orderId)
	//if err == nil && err != gorm.ErrRecordNotFound {
	//	if hasShopOrderRefund.Status == models.PendingReview || hasShopOrderRefund.Status == models.Refunding || hasShopOrderRefund.Status == models.Refunded || hasShopOrderRefund.Status == models.AgreeRefund || hasShopOrderRefund.Status == models.PendingPass {
	//		rsp.JsonResonse(ctx, rsp.ReShopOrderRefund, nil, "")
	//		return
	//	}
	//}

	order, err := models.GetOrderById(userId, orderId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderNotExits, nil, "")
		return
	}

	if order.RefundId != 0 {
		rsp.JsonResonse(ctx, rsp.ReShopOrderRefund, nil, "")
		return
	}

	if order.Status == models.PendingPay || order.Status == models.CancelOrder {
		rsp.JsonResonse(ctx, rsp.ShopOrderRefundRejected, nil, "")
		return
	}

	if order.Status == models.Completed {
		duringDays := utils.BetweenDays(nowTime, order.CompletedAtFormat)
		if duringDays > int64(7) {
			rsp.JsonResonse(ctx, rsp.ShopOrderRefundRejected, nil, "")
			return
		}
	}

	price := order.RealPrice
	if order.Status == models.ToBeReceived || order.Status == models.Completed {
		price = price - order.PostPrice
		if price < float64(0) {
			price = float64(0)
		}
	}

	var imgs string
	if reasonPics != "" {
		imgs = utils.JsonEncode(reasonPics)
	} else {
		imgs = ""
	}

	shopOrderRefund := models.ShopOrderRefund{
		RefundN:    utils.Uid("RO"),
		Price:      price,
		Status:     models.PendingReview,
		RType:      rType,
		ReasonPics: imgs,
		Reason:     reason,
		UserId:     userId,
		OrderId:    orderId,
	}

	refundId, err := shopOrderRefund.Create()
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderRefundFailed, nil, "")
		return
	}

	err = models.UpdateOrderRefundId(userId, orderId, refundId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderRefundFailed, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

/*
退货退款状态填写快递信息
*/
func PostReturnInfo(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	refundId, _ := strconv.ParseInt(ctx.PostForm("refund_id"), 10, 64)
	expressN := ctx.PostForm("express_n")
	expressId, _ := strconv.ParseInt(ctx.PostForm("express_id"), 10, 64)

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	shopOrderRefund, err := models.GetShopOrderRefundById(userId, refundId)
	if err != nil && err == gorm.ErrRecordNotFound {
		rsp.JsonResonse(ctx, rsp.ShopOrderRefundNotExits, nil, "")
		return
	}

	if shopOrderRefund.Status != models.AgreeRefund {
		rsp.JsonResonse(ctx, rsp.ReShopOrderRefund, nil, "")
		return
	}
	shopOrderRefundUpdate := models.ShopOrderRefund{
		ExpressId: expressId,
		ExpressN:  expressN,
		Status:    models.PendingPass,
	}
	err = shopOrderRefundUpdate.UpdateShopOrderRefundExpress(refundId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderRefundPostFailed, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func CancelRefund(ctx *gin.Context)  {

	userId := ctx.MustGet("user_id").(int64)
	orderId, _ := strconv.ParseInt(ctx.PostForm("order_id"), 10, 64)

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	_, err := models.GetOrderById(userId, orderId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderNotExits, nil, "")
		return
	}

	err =  models.UpdateOrderRefundToCancel(userId,orderId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderRefundCancelFailed, nil, "")
		return
	}

	refundOrder := models.ShopOrderRefund{}
	err = refundOrder.Delete(userId,orderId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderRefundCancelFailed, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func Upload(ctx *gin.Context) {

	app, _ := config.Cfg.GetSection("qiniu")
	MediaUrl := app.Key("MediaUrl").String()
	file, _ := ctx.FormFile("file")
	bucket := "soulfire-media"

	ext := path.Ext(file.Filename)
	key := utils.Uid("FE") + ext

	dst := "runtime/tmp/imgs/" + key

	err := ctx.SaveUploadedFile(file, dst)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.UploadErr, nil, "")
		return
	}

	img, err := qiniu.Upload(bucket, dst, "shop/refund/"+key)

	url := MediaUrl + "/" + img

	if err != nil {
		rsp.JsonResonse(ctx, rsp.UploadErr, nil, "")
		return
	}

	_ = os.Remove(dst)

	rsp.JsonResonse(ctx, rsp.OK, url, "")

}

package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func GoodsDetail(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data := make(map[string]interface{})

	shopGoodsDetail, err := models.GetShopGoodsById(id)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.GoodsNotExits, nil, "")
		return
	}

	shopGoodsSpus, err := models.GetGoodsSpusById(id)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.GoodsNotExits, nil, "")
		return
	}

	purchasers, pTotal, _ := models.GetPurchasersById(id)

	data["shop_goods_detail"] = shopGoodsDetail
	data["shop_goods_spus"] = shopGoodsSpus
	data["purchasers"] = purchasers
	data["p_total"] = pTotal

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func PreOrderDetail(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)
	goodsSpuId, _ := strconv.ParseInt(ctx.Param("goods_spu_id"), 10, 64)

	data := make(map[string]interface{}, 0)

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}



	goodsSpu, err := models.GetGoodsSpuById(goodsSpuId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopGoodsSpuNotExits, nil, "")
		return
	}

	coupons,err := models.GetCanUseCoupons(userId, goodsSpu.GoodsId)
	if err != nil || err == gorm.ErrRecordNotFound {
		coupons = nil
	}

	data["goods_spu"] = goodsSpu
	data["coupons"] = coupons

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func OrderDetail(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)
	orderId, _ := strconv.ParseInt(ctx.Param("order_id"), 10, 64)

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	order, err := models.GetOrderDetailById(userId, orderId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, order, "")

}

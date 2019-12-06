package shop

import (
	"github.com/gin-gonic/gin"
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

	userId := ctx.MustGet("user_id").(int64)
	goodsSpuId, _ := strconv.ParseInt(ctx.Param("goods_spu_id"), 10, 64)

	data := make(map[string]interface{}, 0)

	defaultAddress, _ := models.GetDefaultAddress(userId)

	goodsSpu, err := models.GetGoodsSpuById(goodsSpuId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopGoodsSpuNotExits, nil, "")
		return
	}

	couponsCount := models.GetCanUseCouponCountById(userId, goodsSpu.GoodsId)

	data["goods"] = goodsSpu
	data["default_address"] = defaultAddress
	data["coupons_count"] = couponsCount

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

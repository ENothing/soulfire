package shop

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func GoodsDetail(ctx *gin.Context)  {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data := make(map[string]interface{})

	shopGoodsDetail,err := models.GetShopGoodsById(id)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.GoodsNotExits, nil,"")
		return
	}

	shopGoodsSpus ,err := models.GetGoodsSpusById(id)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.GoodsNotExits, nil,"")
		return
	}

	purchasers,_ := models.GetPurchasersById(id)

	data["shop_goods_detail"] = shopGoodsDetail
	data["shop_goods_spus"] = shopGoodsSpus
	data["purchasers"] = purchasers

	rsp.JsonResonse(ctx, rsp.OK, data,"")

}

package shop

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func GoodsList(ctx *gin.Context) {

	cateId, _ := strconv.ParseInt(ctx.DefaultQuery("cate_id", "0"), 10, 64)
	brandId, _ := strconv.ParseInt(ctx.DefaultQuery("brand_id", "0"), 10, 64)
	name := ctx.Query("goods_name")
	sort, _ := strconv.ParseInt(ctx.DefaultQuery("sort", "0"), 10, 64)
	sortType, _ := strconv.ParseInt(ctx.DefaultQuery("sort_type", "0"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	data := make(map[string]interface{})

	goods, total, lastPage, err := models.ShopGoodsPaginate(page, pageSize, sortType, sort, name, cateId, brandId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.GoodsListNotExits, nil, "")
		return
	}

	data["goods"] = goods
	data["total"] = total
	data["lastPage"] = lastPage

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func PurchasersList(ctx *gin.Context) {

	goodsId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	data := make(map[string]interface{})

	purchasers, total, lastPage, err := models.ShopOrderGoodsPaginate(page, pageSize, goodsId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderGoodsNotExits, nil, "")
		return
	}

	data["purchasers"] = purchasers
	data["total"] = total
	data["lastPage"] = lastPage

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func CateWithBrandList(ctx *gin.Context) {

	goodsCates, _ := models.GetCateWithBrand()

	rsp.JsonResonse(ctx, rsp.OK, goodsCates, "")

}

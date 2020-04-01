package shop

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func GoodsList(ctx *gin.Context) {

	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)
	cateId, _ := strconv.ParseInt(ctx.DefaultQuery("cate_id", "0"), 10, 64)
	brandId, _ := strconv.ParseInt(ctx.DefaultQuery("brand_id", "0"), 10, 64)
	kword := ctx.DefaultQuery("kword", "")
	sort, _ := strconv.ParseInt(ctx.DefaultQuery("sort", "0"), 10, 64)
	sortType, _ := strconv.ParseInt(ctx.DefaultQuery("sort_type", "0"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	data := make(map[string]interface{})

	if kword != "" && userId != int64(0){
		shopSearchHistory := models.ShopSearchHistory{
			UserId:userId,
			Kword:kword,
		}

		_ = shopSearchHistory.Create()
	}

	goods, total, lastPage, err := models.ShopGoodsPaginate(page, pageSize, sortType, sort, kword, cateId, brandId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.GoodsListNotExits, nil, "")
		return
	}

	data["goods"] = goods
	data["total"] = total
	data["last_page"] = lastPage

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
	data["last_page"] = lastPage

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func CateWithBrandList(ctx *gin.Context) {

	goodsCates, _ := models.GetCateWithBrand()

	rsp.JsonResonse(ctx, rsp.OK, goodsCates, "")

}

func OrderList(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)
	status := ctx.DefaultQuery("status", "")

	data := make(map[string]interface{})

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	orders, total, lastPage, err := models.ShopOrderPaginate(page, pageSize, userId, status)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ShopOrderGoodsNotExits, nil, "")
		return
	}

	data["orders"] = orders
	data["total"] = total
	data["last_page"] = lastPage

	rsp.JsonResonse(ctx, rsp.OK, data, "")
}

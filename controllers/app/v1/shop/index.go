package shop

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
)

func Index(ctx *gin.Context) {

	data := make(map[string]interface{})

	shopBanners, _ := models.GetBannersByCate(2)
	goodsCates, _ := models.GetGoodsCateLimitNum(8)

	data["shop_banners"] = shopBanners
	data["goods_cates"] = goodsCates

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

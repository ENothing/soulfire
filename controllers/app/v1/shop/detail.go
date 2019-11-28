package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func GoodsDetail(ctx *gin.Context)  {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	fmt.Println(id)

	shopGoodsDetail,err := models.GetShopGoodsById(id)

	fmt.Println(err)
	if err != nil {

		rsp.JsonResonse(ctx, rsp.GoodsNotExits, nil,"")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, shopGoodsDetail,"")

}

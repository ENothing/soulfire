package address

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Detail(ctx *gin.Context) {

	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	shipAddress, err := models.GetAddressById(id, userId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.AddressNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, shipAddress, "")

}

func DetailToOrder(ctx *gin.Context) {

	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)
	id, _ := strconv.ParseInt(ctx.DefaultQuery("id","0"), 10, 64)

	shipAddress, _ := models.GetAddress(id, userId)

	rsp.JsonResonse(ctx, rsp.OK, shipAddress, "")

}

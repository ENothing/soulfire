package address

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Detail(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	shipAddress, err := models.GetAddressById(id, userId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.AddressNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, shipAddress, "")

}

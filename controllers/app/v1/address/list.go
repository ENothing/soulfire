package address

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func AddressList(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	shipAddresses, total, lastPage, err := models.AddressPaginate(page, pageSize, userId)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ActivityListNotExits, nil, "")
		return
	}
	defaultCount := models.GetDefaultAddressCount(userId)
	data := make(map[string]interface{})

	data["total"] = total
	data["last_page"] = lastPage
	data["ship_addresses"] = shipAddresses
	data["default_count"] = defaultCount

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

package address

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func AddressList(ctx *gin.Context)  {

	userId := ctx.MustGet("user_id").(int64)
	page,_ := strconv.ParseInt(ctx.DefaultQuery("page","1"),10,64)
	pageSize,_ := strconv.ParseInt(ctx.DefaultQuery("pageSize","10"),10,64)

	shipAddresses,total,lastPage,err := models.AddressPaginate(page,pageSize,userId)

	if err != nil {
		rsp.JsonResonse(ctx,rsp.ActivityListNotExits,nil,"")
		return
	}

	data := make(map[string]interface{})

	data["total"] = total
	data["last_page"] = lastPage
	data["ship_addresses"] = shipAddresses

	rsp.JsonResonse(ctx,rsp.OK,data,"")

}
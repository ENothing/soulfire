package activity

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Pay(ctx *gin.Context)  {

	//todo 调起支付


	userId := ctx.MustGet("user_id").(int64)

	id,_ := strconv.ParseInt(ctx.PostForm("id"),10,64)


	activityOrder,err := models.GetActivityOrderById(id)

	if err != nil {

		rsp.JsonResonse(ctx,rsp.ActivityOrderNotExits,nil,"")
		return

	}












	
}
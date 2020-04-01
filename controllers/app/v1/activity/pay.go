package activity

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Pay(ctx *gin.Context)  {

	//todo 调起支付

	//
	userId,_ := ctx.MustGet("user_id").(int64)

	id,_ := strconv.ParseInt(ctx.PostForm("id"),10,64)
	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	_,err := models.GetActivityOrderById(id,userId)

	if err != nil {

		rsp.JsonResonse(ctx,rsp.ActivityOrderNotExits,nil,"")
		return

	}

	//todo 返回支付参数


	rsp.JsonResonse(ctx,rsp.OK,nil,"")

}

func Notify(ctx *gin.Context)  {

	//todo 支付回调

}
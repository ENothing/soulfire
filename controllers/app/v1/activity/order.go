package activity

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func OrderDetail(ctx *gin.Context)  {

	userId := ctx.MustGet("user_id").(int64)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	activityOrder,err := models.GetActivityOrderById(id,userId)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.ActivityOrderNotExits, nil,"")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, activityOrder,"")

}

func FinishOrder(ctx *gin.Context)  {

	userId := ctx.MustGet("user_id").(int64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	err := models.UpdateActivityOrderToFinished(userId,id)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ActivityOrderFinishedFailed, nil,"")
		return
	}
	rsp.JsonResonse(ctx, rsp.OK, nil,"")

}
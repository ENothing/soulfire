package activity

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Detail(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	activity, err := models.GetActivityById(id)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil,"")
		return
	}

	err = models.ActivityViewAddOne(id)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil,"")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, activity,"")

}

func OrderDetail(ctx *gin.Context)  {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	activityOrder,err := models.GetActivityOrderById(id)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.ActivityOrderNotExits, nil,"")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, activityOrder,"")

}

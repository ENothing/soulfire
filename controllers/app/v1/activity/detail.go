package activity

import (
	"gin-init/models"
	"gin-init/pkg/rsp"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Detail(ctx *gin.Context)  {

	id,_ := strconv.ParseInt(ctx.Param("id"),10,64)

	activity,err := models.GetActivityById(id)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil)
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, activity)

}
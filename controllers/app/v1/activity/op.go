package activity

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Like(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	likes := models.LikeAndUnlike(userId, id, 1)

	if likes == true {

		err := models.ActivityLikeAddOne(id)

		if err != nil {

			rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil,"")
			return
		}

		rsp.JsonResonse(ctx, rsp.OK, likes,"")
		return

	}

	err := models.ActivityLikeCutOne(id)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil,"")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, likes,"")

}

package activity

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Like(ctx *gin.Context) {

	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	likes := models.LikeAndUnlike(userId, id, 1)

	if likes == true {

		err := models.ActivityLikeAddOne(id)

		if err != nil {

			rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil, "")
			return
		}

		rsp.JsonResonse(ctx, rsp.OK, likes, "")
		return

	}

	err := models.ActivityLikeCutOne(id)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, likes, "")

}

func Favor(ctx *gin.Context) {

	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	favor := models.FavorAndUnFavor(userId, id, 1)

	if favor == true {

		err := models.ActivityFavorAddOne(id)

		if err != nil {

			rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil, "")
			return
		}

		rsp.JsonResonse(ctx, rsp.OK, favor, "")
		return

	}

	err := models.ActivityFavorCutOne(id)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, favor, "")

}

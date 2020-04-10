package activity

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"soulfire/pkg/sms"
	"strconv"
)

func Like(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if userId == int64(0) {
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

	userId,_ := ctx.MustGet("user_id").(int64)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if userId == int64(0) {
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

func SendSms(ctx *gin.Context)  {

	mobile := ctx.PostForm("mobile")


	if mobile == "" {

		rsp.JsonResonse(ctx, rsp.MobileEmpty, nil, "")

	}

	sms.NewSms(mobile).SendCode()




	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func GetCode(ctx *gin.Context)  {

	mobile := ctx.PostForm("mobile")
	fmt.Println(sms.NewSms(mobile).GetCode())


}



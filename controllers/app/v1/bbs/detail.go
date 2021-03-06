package bbs

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/logging"
	"soulfire/pkg/rsp"
	"strconv"
)

func Detail(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId,_ := ctx.MustGet("user_id").(int64)


	article, err := models.GetArticleById(id, userId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
		return
	}

	err = models.ArticleViewAddOne(id)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, article, "")

}

func ArticleEditDetail(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId,_ := ctx.MustGet("user_id").(int64)
	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}


	article, err := models.GetSelfArticleById(id, userId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, article, "")

}

func UserDetail(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	userId,_ := ctx.MustGet("user_id").(int64)

	user := models.UserDetail{}

	userInfo, err := user.GetUserById(id, userId)

	if err != nil {
		logging.Logging(logging.ERR, err)
		rsp.JsonResonse(ctx, rsp.UserNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, userInfo, "")

}

package bbs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/logging"
	"soulfire/pkg/rsp"
	"strconv"
)

func Detail(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId := ctx.MustGet("user_id").(int64)


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
	userId := ctx.MustGet("user_id").(int64)
	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	if userId == 0 {
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
	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)
	fmt.Println(id)
	fmt.Println(userId)

	user := models.UserDetail{}

	userInfo, err := user.GetUserById(id, userId)

	if err != nil {
		logging.Logging(logging.ERR, err)
		rsp.JsonResonse(ctx, rsp.UserNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, userInfo, "")

}

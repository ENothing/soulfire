package bbs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Detail(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId := ctx.MustGet("user_id").(int64)

	fmt.Println(userId)

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

func UserDetail(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	userId := ctx.MustGet("user_id").(int64)

	user := models.User{}

	userInfo, err := user.GetUserById(id, userId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.UserNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, userInfo, "")

}

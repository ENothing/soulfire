package bbs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Detail(ctx *gin.Context) {

	//id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId := ctx.MustGet("user_id").(string)

	fmt.Println(userId)

	//
	//article, err := models.GetArticleById(id,userId)
	//
	//if err != nil {
	//	rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil,"")
	//	return
	//}
	//
	//err = models.ArticleViewAddOne(id)
	//
	//if err != nil {
	//	rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil,"")
	//	return
	//}
	//
	//rsp.JsonResonse(ctx, rsp.OK, article,"")

}

func UserDetail(ctx *gin.Context) {

	userId, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)

	user := models.User{}

	userInfo, err := user.GetUserById(userId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.UserNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, userInfo, "")

}

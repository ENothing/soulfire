package bbs

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func Detail(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	article, err := models.GetArticleById(id)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil,"")
		return
	}

	err = models.ArticleViewAddOne(id)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil,"")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, article,"")

}
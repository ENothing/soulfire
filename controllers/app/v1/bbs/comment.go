package bbs

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)


func PostComment(ctx *gin.Context)  {

	userId := ctx.MustGet("user_id").(int64)
	id, _ := strconv.ParseInt(ctx.PostForm("id"), 10, 64)
	parentId, _ := strconv.ParseInt(ctx.PostForm("parent_id"), 10, 64)
	content := ctx.PostForm("content")

	comment := models.ArticleComment{
		UserId:userId,
		ArticleId:id,
		Content:content,
		ParentId:parentId,
	}

	err := comment.Create()
	if err != nil {

		rsp.JsonResonse(ctx,rsp.ArticleCommentCreateFailed,nil,"")
		return

	}

	rsp.JsonResonse(ctx,rsp.OK,nil,"")

}
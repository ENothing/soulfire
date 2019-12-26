package bbs

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func ArticleList(ctx *gin.Context) {

	userId, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	cateId, _ := strconv.ParseInt(ctx.DefaultQuery("cate_id", "0"), 10, 64)
	title := ctx.Query("title")
	sort, _ := strconv.ParseInt(ctx.DefaultQuery("sort", "0"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	data := make(map[string]interface{})

	articles, total, lastPage, err := models.ArticlePaginate(page, pageSize, sort, cateId, title, userId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
		return
	}

	data["total"] = total
	data["last_page"] = lastPage
	data["articles"] = articles

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func CommentList(ctx *gin.Context) {

	articleId, _ := strconv.ParseInt(ctx.Query("article_id"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	data := make(map[string]interface{})

	articleComments, total, lastPage, err := models.ArticleCommentPaginate(page, pageSize, articleId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleCommentListNotExits, nil, "")
		return
	}

	data["total"] = total
	data["last_page"] = lastPage
	data["articleComments"] = articleComments

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func UserArticleList(ctx *gin.Context) {

	userId, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	data := make(map[string]interface{})

	articles, total, lastPage, err := models.UserArticlePaginate(page, pageSize, userId)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
		return
	}

	data["total"] = total
	data["last_page"] = lastPage
	data["articles"] = articles

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

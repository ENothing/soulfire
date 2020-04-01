package bbs

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func ArticleList(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)
	cateId, _ := strconv.ParseInt(ctx.DefaultQuery("cate_id", "0"), 10, 64)
	title := ctx.Query("title")
	sort, _ := strconv.ParseInt(ctx.DefaultQuery("sort", "0"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "5"), 10, 64)

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
		rsp.JsonResonse(ctx, rsp.OK, "", "")
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
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "5"), 10, 64)

	data := make(map[string]interface{})

	articles, total, lastPage, err := models.UserArticlePaginate(page, pageSize, userId,1)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
		return
	}

	data["total"] = total
	data["last_page"] = lastPage
	data["articles"] = articles

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func UserLikeArticleList(ctx *gin.Context) {

	userId, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	//status := ctx.DefaultQuery("status", "")
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	var dataList interface{}
	var total int64
	var lastPage int64
	var err error
	data := make(map[string]interface{})

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	//if status == "" {
	//	dataList, total, lastPage, err = models.GetActivityFavorPaginate(page, pageSize, userId)
	//} else {
	dataList, total, lastPage, err = models.GetArticleLikePaginate(page, pageSize, userId)
	//}

	data["list"] = dataList
	data["total"] = total
	data["last_page"] = lastPage
	data["err"] = err

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func ArticleCateList(ctx *gin.Context) {

	articleCates, err := models.GetArticleCates()
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleCateGetFailed, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, articleCates, "")

}

func FollowsList(ctx *gin.Context)  {

	userId,_ := ctx.MustGet("user_id").(int64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	var dataList interface{}
	var total int64
	var lastPage int64
	data := make(map[string]interface{})

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}


	dataList, total, lastPage, _ = models.GetUserFollowsPaginate(page, pageSize, userId,id)


	data["list"] = dataList
	data["total"] = total
	data["last_page"] = lastPage

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func FollowedList(ctx *gin.Context)  {

	userId,_ := ctx.MustGet("user_id").(int64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "10"), 10, 64)

	var dataList interface{}
	var total int64
	var lastPage int64
	data := make(map[string]interface{})

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	//if status == "" {
	//	dataList, total, lastPage, err = models.GetActivityFavorPaginate(page, pageSize, userId)
	//} else {
	dataList, total, lastPage, _ = models.GetUserFollowedPaginate(page, pageSize, userId,id)
	//}

	data["list"] = dataList
	data["total"] = total
	data["last_page"] = lastPage

	rsp.JsonResonse(ctx, rsp.OK, data, "")



}

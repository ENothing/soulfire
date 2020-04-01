package user

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func PersonalCollection(ctx *gin.Context) {

	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)
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

func FollowsList(ctx *gin.Context)  {

	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)
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
	dataList, total, lastPage, _ = models.GetFollowsPaginate(page, pageSize, userId)
	//}

	data["list"] = dataList
	data["total"] = total
	data["last_page"] = lastPage

	rsp.JsonResonse(ctx, rsp.OK, data, "")



}

func FollowedList(ctx *gin.Context)  {

	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)
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
	dataList, total, lastPage, _ = models.GetFollowedPaginate(page, pageSize, userId)
	//}

	data["list"] = dataList
	data["total"] = total
	data["last_page"] = lastPage

	rsp.JsonResonse(ctx, rsp.OK, data, "")



}

func MyArticleList(ctx *gin.Context)  {
	userId,_ := strconv.ParseInt(ctx.MustGet("user_id").(string), 10, 64)
	page, _ := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(ctx.DefaultQuery("pageSize", "5"), 10, 64)

	data := make(map[string]interface{})
	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	articles, total, lastPage, err := models.UserArticlePaginate(page, pageSize, userId,0)

	if err != nil {
		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
		return
	}

	data["total"] = total
	data["last_page"] = lastPage
	data["articles"] = articles

	rsp.JsonResonse(ctx, rsp.OK, data, "")
}
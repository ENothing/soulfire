package user

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func PersonalCollection(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
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

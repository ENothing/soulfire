package activity

import (
	"soulfire/models"
	"soulfire/pkg/rsp"
	"github.com/gin-gonic/gin"
	"strconv"
)

func List(ctx *gin.Context)  {

	cateId,_ :=strconv.ParseInt(ctx.DefaultQuery("cate_id","0"),10,64)
	title := ctx.Query("title")
	sort,_ := strconv.ParseInt(ctx.DefaultQuery("sort","0"),10,64)
	page,_ := strconv.ParseInt(ctx.DefaultQuery("page","1"),10,64)
	pageSize,_ := strconv.ParseInt(ctx.DefaultQuery("pageSize","10"),10,64)

	data := make(map[string]interface{})

	activities,total,lastPage,err := models.ActivityPaginate(page,pageSize,sort,cateId,title)

	if err != nil {
		rsp.JsonResonse(ctx,rsp.ActivityListNotExits,nil,"")
		return
	}

	data["total"] = total
	data["last_page"] = lastPage
	data["activities"] = activities

	rsp.JsonResonse(ctx,rsp.OK,data,"")

}

func ActivityCates(ctx *gin.Context)  {

	activityCates,err := models.GetActivityCateLimitNum(-1)

	if err != nil {
		rsp.JsonResonse(ctx,rsp.ActivityCateNotExits,nil,"")
		return
	}

	rsp.JsonResonse(ctx,rsp.OK,activityCates,"")

}
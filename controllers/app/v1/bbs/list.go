package bbs

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"strconv"
)

func ArticleList(ctx *gin.Context)  {

	cateId,_ :=strconv.ParseInt(ctx.DefaultQuery("cate_id","0"),10,64)
	title := ctx.Query("title")
	sort,_ := strconv.ParseInt(ctx.DefaultQuery("sort","0"),10,64)
	page,_ := strconv.ParseInt(ctx.DefaultQuery("page","1"),10,64)
	pageSize,_ := strconv.ParseInt(ctx.DefaultQuery("pageSize","10"),10,64)

	data := make(map[string]interface{})

	activities,total,lastPage,err := models.ActivityPaginate(page,pageSize,sort,cateId,title)

	if err != nil {
		rsp.JsonResonse(ctx,rsp.ArticleNotExits,nil,"")
		return
	}

	data["total"] = total
	data["last_page"] = lastPage
	data["activities"] = activities

	rsp.JsonResonse(ctx,rsp.OK,data,"")










}
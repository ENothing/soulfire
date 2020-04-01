package activity

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
)

func GetHistory(ctx *gin.Context)  {

	userId,_ := ctx.MustGet("user_id").(int64)
	data := make(map[string]interface{})

	shopSearchHistory,_ := models.GetActivityHistoryByUserId(userId)
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	shopSearchHistory = nil
	//}

	shopHotHistory,_ := models.GetActivityHotHistory()
	//if err != nil {
	//	shopHotHistory = nil
	//}


	data["history"] = shopSearchHistory
	data["hot"] = shopHotHistory

	rsp.JsonResonse(ctx,rsp.OK,data,"")
}

func DynamicHistory(ctx *gin.Context)  {

	kword := ctx.Query("kword")

	shopSearchHistory,_ := models.GetActivityDynamicHistory(kword)

	rsp.JsonResonse(ctx,rsp.OK,shopSearchHistory,"")

}

func DelSearchHistory(ctx *gin.Context)  {

	userId,_ := ctx.MustGet("user_id").(int64)

	err := models.DelAllActivityHistoryByUserId(userId)
	if err != nil {

		rsp.JsonResonse(ctx,rsp.ShopHistoryDelFailed,nil,"")
		return

	}

	rsp.JsonResonse(ctx,rsp.OK,nil,"")

}

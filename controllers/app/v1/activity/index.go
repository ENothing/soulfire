package activity

import (
	"gin-init/models"
	"gin-init/pkg/rsp"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	data := make(map[string]interface{})

	activityBanners, err := models.GetBannerByCate(1)
	if err != nil {

		rsp.JsonResonse(c, rsp.BannersNotExits, nil)
		return

	}


	data["activity_banners"] = activityBanners

	//todo 公告


	activityCates,err := models.GetActivityCateLimitNum(5)


	if err != nil {

		rsp.JsonResonse(c, rsp.BannersNotExits, nil)
		return

	}


	data["activity_cates"] = activityCates




	rsp.JsonResonse(c, rsp.OK, data)

}

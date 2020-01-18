package activity

import (
	"soulfire/models"
	"soulfire/pkg/rsp"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	data := make(map[string]interface{})

	activityBanners, err := models.GetBannersByCate(1)
	if err != nil {

		rsp.JsonResonse(ctx, rsp.BannersNotExits, nil,"")
		return

	}

	activityAnnounce, _ := models.GetBannerByCate(4)

	activityCates,err := models.GetActivityCateLimitNum(5)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ActivityCateNotExits, nil,"")
		return
	}

	activityVideo, err := models.GetBannerByCate(3)
	if err != nil {

		rsp.JsonResonse(ctx, rsp.VideoNotExits, nil,"")
		return

	}

	data["activity_banners"] = activityBanners
	data["activity_ann"] = activityAnnounce
	data["activity_cates"] = activityCates
	data["activity_video"] = activityVideo

	rsp.JsonResonse(ctx, rsp.OK, data,"")

}

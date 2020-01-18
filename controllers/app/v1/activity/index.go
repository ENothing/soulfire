package activity

import (
	"github.com/jinzhu/gorm"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	data := make(map[string]interface{})

	activityBanners, err := models.GetBannersByCate(1)
	if err != nil || err == gorm.ErrRecordNotFound {
		data["activity_banners"] = ""
	}else{
		data["activity_banners"] = activityBanners
	}

	activityAnnounce, err := models.GetBannerByCate(4)
	if err != nil ||  err == gorm.ErrRecordNotFound {
		data["activity_ann"] = ""
	}else{
		data["activity_ann"] = activityAnnounce
	}

	activityCates,err := models.GetActivityCateLimitNum(6)
	if err != nil || err == gorm.ErrRecordNotFound {
		data["activity_cates"] = ""
	}else{
		data["activity_cates"] = activityCates
	}

	activityVideo, err := models.GetBannerByCate(3)
	if err != nil || err == gorm.ErrRecordNotFound {
		data["activity_video"] = ""
	}else{
		data["activity_video"] = activityVideo
	}


	rsp.JsonResonse(ctx, rsp.OK, data,"")

}

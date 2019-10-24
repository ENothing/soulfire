package models

import (
	"gin-init/pkg/db"
)

type ActivityCate struct {
	Model
	Name string `json:"name" gorm:"column:name;not null"`
}

func GetActivityCateLimitNum(num int64) ([]*ActivityCate,error) {

	activityCates := make([]*ActivityCate,0)

	res := db.DB.Self.Limit(num).Find(&activityCates)

	return activityCates,res.Error

}




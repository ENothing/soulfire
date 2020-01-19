package models

import (
	"soulfire/pkg/db"
)

type ActivityCate struct {
	Model
	Name    string `json:"name" gorm:"column:name;not null"`
	IconUrl string `json:"icon_url" gorm:"column:icon_url;not null"`
}

func (ActivityCate) TableName() string {
	return "activity_cates"
}

func GetActivityCateLimitNum(num int64) ([]*ActivityCate, error) {

	activityCates := make([]*ActivityCate, 0)

	res := db.DB.Self.Limit(num).Find(&activityCates)

	return activityCates, res.Error

}

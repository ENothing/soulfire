package models

import (
	"github.com/jinzhu/gorm"
	"soulfire/pkg/db"
	"time"
)

type ActivitySearchHistory struct {
	Model
	UserId int64 `json:"user_id" gorm:"column:user_id;not null"`
	Kword string `json:"kword" gorm:"column:kword;not null"`
	CreatedAt  time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

type ActivityHotHistory struct {
	Kword string `json:"kword" gorm:"column:kword;not null"`
}


func (ActivitySearchHistory) TableName() string {
	return "activity_search_histories"
}

func (ActivityHotHistory) TableName() string {
	return "activity_search_histories"
}

func (ash *ActivitySearchHistory)Create()(err error) {

	sh := &ActivitySearchHistory{}
	res := db.DB.Self.Where("user_id = ?",ash.UserId).Where("kword = ?",ash.Kword).First(&sh)
	if  res.Error == gorm.ErrRecordNotFound {

		err = db.DB.Self.Create(&ash).Error

	}else{

		err = db.DB.Self.Model(&ActivitySearchHistory{}).Where("user_id = ?",ash.UserId).Where("kword = ?",ash.Kword).UpdateColumn("updated_at",time.Now()).Error

	}

	return

}

func DelAllActivityHistoryByUserId(userId int64)error  {
	return db.DB.Self.Where("user_id = ?", userId).Delete(&ActivitySearchHistory{}).Error
}


func GetActivityHistoryByUserId(userId int64)(activitySearchHistory []*ActivitySearchHistory,err error)  {

	err = db.DB.Self.Where("user_id = ?",userId).Order("updated_at desc").Find(&activitySearchHistory).Error

	return

}

func GetActivityHotHistory()(activityHotHistory []*ActivityHotHistory,err error){

	err = db.DB.Self.
		Unscoped().
		Select("kword").
		Group("kword").
		Order("count(id) desc").
		Limit(15).
		Find(&ActivityHotHistory{}).Scan(&activityHotHistory).Error

	return

}

func GetActivityDynamicHistory(kword string)(activitySearchHistory []*ActivitySearchHistory,err error)  {

	err = db.DB.Self.Unscoped().Where("kword LIKE ?", "%"+kword+"%").Group("kword").Order("updated_at desc").Limit(20).Find(&activitySearchHistory).Error

	return

}




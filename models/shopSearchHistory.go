package models

import (
	"github.com/jinzhu/gorm"
	"soulfire/pkg/db"
	"time"
)

type ShopSearchHistory struct {
	Model
	UserId int64 `json:"user_id" gorm:"column:user_id;not null"`
	Kword string `json:"kword" gorm:"column:kword;not null"`
	CreatedAt  time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

type shopHotHistory struct {
	Kword string `json:"kword" gorm:"column:kword;not null"`
}


func (ShopSearchHistory) TableName() string {
	return "shop_search_histories"
}

func (ssh *ShopSearchHistory)Create()(err error) {

	sh := &ShopSearchHistory{}
	res := db.DB.Self.Where("user_id = ",ssh.UserId).Where("kword = ?",ssh.Kword).First(&sh)
	if res != nil && res.Error != gorm.ErrRecordNotFound {

		err = db.DB.Self.Create(&ssh).Error

	}else{

		err = db.DB.Self.Where("user_id = ",ssh.UserId).Where("kword = ?",ssh.Kword).UpdateColumn("updated_at",time.Now()).Error

	}

	return

}

func DelAllHistoryByUserId(userId int64)error  {
	return db.DB.Self.Where("user_id = ?", userId).Delete(&ShopSearchHistory{}).Error
}


func GetHistoryByUserId(userId int64)(shopSearchHistory []*ShopSearchHistory,err error)  {

	err = db.DB.Self.Where("user_id = ?",userId).Order("updated_at desc").Find(&shopSearchHistory).Error

	return

}

func GetHotHistory()(shopHotHistory []*shopHotHistory,err error){

	err = db.DB.Self.
		Unscoped().
		Select("kword").
		Group("kword").
		Order("count(id) desc").
		Limit(15).
		Find(&ShopSearchHistory{}).Scan(&shopHotHistory).Error

	return

}

func GetDynamicHistory(kword string)(shopSearchHistory []*ShopSearchHistory,err error)  {

	err = db.DB.Self.Unscoped().Where("kword LIKE ?", "%"+kword+"%").Group("kword").Order("updated_at desc").Limit(20).Find(&shopSearchHistory).Error

	return

}




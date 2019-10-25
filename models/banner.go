package models

import (
	"soulfire/pkg/db"
	"time"
)

type Banner struct {
	Model
	CateId string `json:"cate_id" gorm:"column:cate_id;not null" binding:"required"`
	Thumb string `json:"thumb" gorm:"column:thumb;not null" binding:"required" `
	Url string `json:"url" gorm:"column:url"`
	CreatedAt time.Time `gorm:";column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func (Banner) TableName() string {
	return "banners"
}


func GetBannersByCate(cateId int64) ([]*Banner,error) {

	banners := make([]*Banner,0)

	res := db.DB.Self.Where("cate_id = ?",cateId).Find(&banners)

	return banners,res.Error

}

func GetBannerByCate(cateId int64)(*Banner,error){

	banner := &Banner{}

	res := db.DB.Self.Where("cate_id = ?",cateId).First(&banner)

	return banner,res.Error
}




package models

import (
	"soulfire/pkg/db"
)

type ShopGoodsCate struct {
	Model
	Name   string           `json:"name" gorm:"column:name;not null"`
	Brands []ShopGoodsBrand `json:"brands" gorm:"foreignkey:cate_id;PRELOAD:false"`
}

func (ShopGoodsCate) TableName() string {
	return "shop_goods_cates"
}

func GetGoodsCateLimitNum(num int64) ([]*ShopGoodsCate, error) {

	goodsCates := make([]*ShopGoodsCate, 0)

	res := db.DB.Self.Limit(num).Find(&goodsCates)

	return goodsCates, res.Error

}

func GetCateWithBrand() ([]*ShopGoodsCate, error) {

	goodsCates := make([]*ShopGoodsCate, 0)

	res := db.DB.Self.Preload("Brands").Select("shop_goods_cates.*").Find(&goodsCates)

	return goodsCates, res.Error

}

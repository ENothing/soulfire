package models

import (
	"soulfire/pkg/db"
)

type ShopGoodsCate struct {
	Model
	Name    string `json:"name" gorm:"column:name;not null"`
	IconUrl string `json:"icon_url" gorm:"column:icon_url;not null"`
}

type ShopGoodsCateWithBrand struct {
	ShopGoodsCate
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

func GetCateWithBrand() ([]*ShopGoodsCateWithBrand, error) {

	goodsCates := make([]*ShopGoodsCateWithBrand, 0)

	res := db.DB.Self.Preload("Brands").Select("shop_goods_cates.*").Find(&goodsCates)

	return goodsCates, res.Error

}

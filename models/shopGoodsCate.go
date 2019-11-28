package models

import (
	"soulfire/pkg/db"
)

type ShopGoodsCate struct {
	Model
	Name string `json:"name" gorm:"column:name;not null"`
}

func (ShopGoodsCate) TableName() string {
	return "shop_goods_cates"
}


func GetGoodsCateLimitNum(num int64) ([]*ShopGoodsCate,error) {

	goodsCates := make([]*ShopGoodsCate,0)

	res := db.DB.Self.Limit(num).Find(&goodsCates)

	return goodsCates,res.Error

}




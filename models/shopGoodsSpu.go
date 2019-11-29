package models

import "soulfire/pkg/db"

type ShopGoodsSpu struct {
	Model
	GoodsId string `json:"goods_id" gorm:"column:goods_id;not null"`
	Name    string `json:"name" gorm:"column:name;not null"`
	Price   string `json:"price" gorm:"column:price;not null"`
	Stock   string `json:"stock" gorm:"column:stock;not null"`
}

func (ShopGoodsSpu) TableName() string {
	return "shop_goods_spus"
}

func GetGoodsSpusById(goodsId int64) ([]*ShopGoodsSpu, error) {

	goodsSpus := make([]*ShopGoodsSpu, 0)

	res := db.DB.Self.Where("goods_id = ?",goodsId).Find(&goodsSpus)

	return goodsSpus, res.Error

}

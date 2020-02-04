package models

import (
	"github.com/jinzhu/gorm"
	"soulfire/pkg/db"
)

type ShopGoodsSpu struct {
	Model
	GoodsId   int64   `json:"goods_id" gorm:"column:goods_id;not null"`
	Name      string  `json:"name" gorm:"column:name;not null"`
	Price     float64 `json:"price" gorm:"column:price;not null"`
	Stock     int64   `json:"stock" gorm:"column:stock;not null"`
	//GoodsName string  `json:"goods_name" gorm:"column:goods_name;not null"`
	//Thumb     string  `json:"thumb" gorm:"column:thumb;not null"`
	PostPrice float64 `json:"post_price" gorm:"column:post_price;not null"`
}

func (ShopGoodsSpu) TableName() string {
	return "shop_goods_spus"
}

func CutGoodsSpuStock(goodsSpuId, num int64, transaction *gorm.DB) error {
	shopGoodsSpu := &ShopGoodsSpu{}

	res := transaction.Model(&shopGoodsSpu).
		Where("id = ?", goodsSpuId).
		Where("stock >= ?", num).
		Update("stock", gorm.Expr("stock - ?", num))

	return res.Error
}

func GetGoodsSpuById(goodsSpuId int64) (*ShopGoodsSpu, error) {

	goodsSpu := &ShopGoodsSpu{}

	res := db.DB.Self.
		Where("shop_goods_spus.id = ?", goodsSpuId).
		Joins("LEFT JOIN shop_goods as sg ON sg.id = shop_goods_spus.goods_id").
		Select("shop_goods_spus.*,sg.name as goods_name,sg.thumb as thumb").
		First(&goodsSpu)

	return goodsSpu, res.Error
}

func GetGoodsSpusById(goodsId int64) ([]*ShopGoodsSpu, error) {

	goodsSpus := make([]*ShopGoodsSpu, 0)

	res := db.DB.Self.Where("goods_id = ?", goodsId).Find(&goodsSpus)

	return goodsSpus, res.Error

}

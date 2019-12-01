package models

import (
	"soulfire/pkg/db"
	"time"
)

type ShopOrderGoods struct {
	Model
	OrderId       int64      `json:"order_id" gorm:"column:order_id;not null"`
	GoodsId       int64      `json:"goods_id" gorm:"column:goods_id;not null"`
	Num           int64     `json:"num" gorm:"column:num;not null"`
	UnitPrice     float64     `json:"unit_price" gorm:"column:unit_price;not null"`
	TotalPrice    float64     `json:"total_price" gorm:"column:total_price;not null"`
	RealPrice     float64     `json:"real_price" gorm:"column:real_price;not null"`
	DiscountPrice float64    `json:"discount_price" gorm:"column:discount_price;not null"`
	CreatedAt     time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func GetPurchasersById(goodsId int64)(*[]ShopOrderGoods,error)   {

	purchasers := &[]ShopOrderGoods{}

	res := db.DB.Self.
		Where("goods_id = ?",goodsId).
		Find(&purchasers)

	return purchasers,res.Error

}

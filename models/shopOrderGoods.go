package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math"
	"soulfire/pkg/db"
	"soulfire/utils"
	"time"
)

type ShopOrderGoods struct {
	Model
	UserId          int64      `json:"user_id" gorm:"column:user_id;not null"`
	OrderId         int64      `json:"order_id" gorm:"column:order_id;not null"`
	GoodsId         int64      `json:"goods_id" gorm:"column:goods_id;not null"`
	Num             int64      `json:"num" gorm:"column:num;not null"`
	UnitPrice       float64    `json:"unit_price" gorm:"column:unit_price;not null"`
	TotalPrice      float64    `json:"total_price" gorm:"column:total_price;not null"`
	RealPrice       float64    `json:"real_price" gorm:"column:real_price;not null"`
	DiscountPrice   float64    `json:"discount_price" gorm:"column:discount_price;not null"`
	CreatedAt       time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt       time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	SpuId           int64      `json:"spu_id" gorm:"column:spu_id;not null"`
	NickName        string     `json:"nickname" gorm:"column:nickname;not null"`
	Avatar          string     `json:"avatar" gorm:"column:avatar;not null"`
	CreatedAtFormat string     `gorm:";column:created_at_format" json:"created_at_format"`
	Specification   string     `gorm:";column:specification" json:"specification"`
}

type ShopOrderGoodsCreateForm struct {
	Model
	UserId        int64      `json:"user_id" gorm:"column:user_id;not null"`
	OrderId       int64      `json:"order_id" gorm:"column:order_id;not null"`
	GoodsId       int64      `json:"goods_id" gorm:"column:goods_id;not null"`
	Num           int64      `json:"num" gorm:"column:num;not null"`
	UnitPrice     float64    `json:"unit_price" gorm:"column:unit_price;not null"`
	TotalPrice    float64    `json:"total_price" gorm:"column:total_price;not null"`
	RealPrice     float64    `json:"real_price" gorm:"column:real_price;not null"`
	DiscountPrice float64    `json:"discount_price" gorm:"column:discount_price;not null"`
	CreatedAt     time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	SpuId         int64      `json:"spu_id" gorm:"column:spu_id;not null"`
}

func (ShopOrderGoods) TableName() string {
	return "shop_order_goods"
}

func (ShopOrderGoodsCreateForm) TableName() string {
	return "shop_order_goods"
}

func (sogcf *ShopOrderGoodsCreateForm) Create(transaction *gorm.DB) error {

	return transaction.Create(&sogcf).Error

}

func GetPurchasersById(goodsId int64) ([]*ShopOrderGoods, int64, error) {

	purchasers := make([]*ShopOrderGoods, 0)
	var total int64

	res := db.DB.Self.
		Where("shop_order_goods.goods_id = ?", goodsId).
		Joins("LEFT JOIN users as user ON user.id = shop_order_goods.user_id").
		Joins("LEFT JOIN shop_goods_spus as sgs ON sgs.id = shop_order_goods.goods_spu_id").
		Select("shop_order_goods.*,user.nickname as nickname,user.head_url as avatar, sgs.name as specification").
		Limit(5).
		Find(&purchasers)

	for _, value := range purchasers {

		fmt.Println(value.CreatedAt)
		value.CreatedAtFormat = utils.TimeSpan(value.CreatedAt)

	}

	db.DB.Self.Table("shop_order_goods").Where("goods_id = ?", goodsId).Count(&total)

	return purchasers, total, res.Error

}

func ShopOrderGoodsPaginate(page int64, pageSize int64, goodsId int64) (shopOrderGoods []*ShopOrderGoods, total int64, lastPage int64, err error) {

	shopOrderGoods = make([]*ShopOrderGoods, 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self.
		Where("shop_order_goods.goods_id = ?", goodsId).
		Joins("LEFT JOIN users as user ON user.id = shop_order_goods.user_id").
		Joins("LEFT JOIN shop_goods_spus as sgs ON sgs.id = shop_order_goods.goods_spu_id").
		Select("shop_order_goods.*,user.nickname as nickname,user.head_url as avatar, sgs.name as specification").
		Order("created_at desc").
		Limit(pageSize).
		Offset(offset).
		Find(&shopOrderGoods)

	for _, value := range shopOrderGoods {
		value.CreatedAtFormat = utils.TimeSpan(value.CreatedAt)
	}

	db.DB.Self.Model(&shopOrderGoods).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return shopOrderGoods, total, lastPage, res.Error
}

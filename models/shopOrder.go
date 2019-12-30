package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ShopOrder struct {
	Model
	UserId        int64      `json:"user_id" gorm:"column:user_id;not null"`
	OrderN        string     `json:"order_n" gorm:"column:order_n;not null"`
	UserCouponId  int64      `json:"user_coupon_id" gorm:"column:user_coupon_id;not null"`
	Num           int64      `json:"num" gorm:"column:num;not null"`
	UnitPrice     float64    `json:"unit_price" gorm:"column:unit_price;not null"`
	TotalPrice    float64    `json:"total_price" gorm:"column:total_price;not null"`
	RealPrice     float64    `json:"real_price" gorm:"column:real_price;not null"`
	DiscountPrice float64    `json:"discount_price" gorm:"column:discount_price;not null"`
	PostPrice     float64    `json:"post_price" gorm:"column:post_price;not null"`
	Status        int64      `json:"status" gorm:"column:status;not null"`
	Name          string     `json:"name" gorm:"column:name;not null"`
	Mobile        string     `json:"mobile" gorm:"column:mobile;not null"`
	Province      string     `json:"province" gorm:"column:province;not null"`
	City          string     `json:"city" gorm:"column:city;not null"`
	District      string     `json:"district" gorm:"column:district;not null"`
	DetailAddress string     `json:"detail_address" gorm:"column:detail_address;not null"`
	CreatedAt     time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func (ShopOrder) TableName() string {
	return "shop_orders"
}

func (so *ShopOrder) Create(transaction *gorm.DB) (id int64, err error) {

	err = transaction.Create(&so).Error
	id = so.Id
	return

}

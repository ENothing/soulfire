package models

import "time"

type Coupon struct {
	Model
	Name             string     `json:"name" gorm:"column:name;not null"`
	Type             int64      `json:"type" gorm:"column:type;not null"`
	CouponType       int64      `json:"coupon_type" gorm:"column:coupon_type;not null"`
	Num              int64      `json:"num" gorm:"column:num;not null"`
	FullPrice        float64    `json:"full_price" gorm:"column:full_price;not null"`
	ReductionPrice   float64    `json:"reduction_price" gorm:"column:reduction_price;not null"`
	ImmediatelyPrice float64    `json:"immediately_price" gorm:"column:immediately_price;not null"`
	Discount         float64    `json:"discount" gorm:"column:discount;not null"`
	StartReceiveTime time.Time  `json:"start_receive_time" gorm:"column:start_receive_time;not null"`
	EndReceiveTime   time.Time  `json:"end_receive_time" gorm:"column:end_receive_time;not null"`
	IsTimeing        int64      `json:"is_timeing" gorm:"column:is_timeing;not null"`
	UseDay           int64      `json:"use_day" gorm:"column:use_day;not null"`
	StartUseTime     time.Time  `json:"start_use_time" gorm:"column:start_use_time;not null"`
	EndUseTime       time.Time  `json:"end_use_time" gorm:"column:end_use_time;not null"`
	IsGoods          int64      `json:"is_goods" gorm:"column:is_goods;not null"`
	GoodsIds         string     `json:"goods_ids" gorm:"column:goods_ids;not null"`
	CreatedAt        time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt        time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func (Coupon) TableName() string {
	return "coupons"
}

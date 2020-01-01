package models

import "time"

type ShopOrderRefund struct {
	Model
	OrderGoodsId int64      `json:"order_goods_id" gorm:"column:order_goods_id;not null"`
	RefundNo     string     `json:"refund_no" gorm:"column:refund_no;not null"`
	Price        float64    `json:"price" gorm:"column:price;not null"`
	Status       int64      `json:"status" gorm:"column:status;not null"`
	Type         int64      `json:"type" gorm:"column:type;not null"`
	ReasonPics   string     `json:"reason_pics" gorm:"column:reason_pics;not null"`
	Reason       string     `json:"reason" gorm:"column:reason;not null"`
	ReplyReason  string     `json:"reply_reason" gorm:"column:reply_reason;not null"`
	CreatedAt    time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func (ShopOrderRefund) TableName() string {
	return "shop_order_refunds"
}

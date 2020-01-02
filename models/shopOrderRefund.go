package models

import (
	"soulfire/pkg/db"
	"time"
)

type ShopOrderRefund struct {
	Model
	OrderGoodsId int64      `json:"order_goods_id" gorm:"column:order_goods_id;not null"`
	RefundN      string     `json:"refund_n" gorm:"column:refund_n;not null"`
	Price        float64    `json:"price" gorm:"column:price;not null"`
	Status       int64      `json:"status" gorm:"column:status;not null"`
	RType        int64      `json:"r_type" gorm:"column:r_type;not null"`
	ReasonPics   string     `json:"reason_pics" gorm:"column:reason_pics;not null"`
	Reason       string     `json:"reason" gorm:"column:reason;not null"`
	ReplyReason  string     `json:"reply_reason" gorm:"column:reply_reason;not null"`
	CreatedAt    time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	UserId       int64      `gorm:"column:user_id" sql:"index" json:"user_id"`
}

const (
	PendingReview  int64 = iota //发起退款待审核
	CancelRefund                //取消退款
	Refunding                   //退款中
	Refunded                    //退款完成
	RejectedRefund              //拒绝退款
	AgreeRefund                 //同意退款
)

const (
	OnlyRefund      int64 = 1
	RefundAndReturn int64 = 2
)

func (ShopOrderRefund) TableName() string {
	return "shop_order_refunds"
}

func (sor *ShopOrderRefund) Create() error {

	return db.DB.Self.Create(&sor).Error

}

func GetShopOrderRefundById(userId, orderId, orderGoodsId int64) (*ShopOrderRefund, error) {

	shopOrderRefund := &ShopOrderRefund{}

	res := db.DB.Self.
		Where("user_id = ?", userId).
		Where("order_id = ?", orderId).
		Where("order_goods_id", orderGoodsId).
		First(&shopOrderRefund)

	return shopOrderRefund, res.Error

}

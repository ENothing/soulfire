package models

import (
	"soulfire/pkg/db"
	"time"
)

type ActivityOrderRefund struct {
	Model
	RefundN     string     `json:"refund_n" gorm:"column:refund_n;not null"`
	Price       float64    `json:"price" gorm:"column:price;not null"`
	Status      int64      `json:"status" gorm:"column:status;not null"`
	Reason      string     `json:"reason" gorm:"column:reason;not null"`
	ReplyReason string     `json:"reply_reason" gorm:"column:reply_reason;not null"`
	CreatedAt   time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	UserId      int64      `gorm:"column:user_id" json:"user_id"`
	OrderId     int64      `gorm:"column:order_id" json:"order_id"`
}




func (ActivityOrderRefund) TableName() string {
	return "activity_order_refunds"
}

func (aor *ActivityOrderRefund) Create() (int64, error) {

	res := db.DB.Self.Create(&aor)
	return aor.Id, res.Error

}

func (aor *ActivityOrderRefund)Delete(userId,orderId int64) error  {


	return db.DB.Self.Where("user_id = ?",userId).Where("order_id = ?",orderId).Delete(&aor).Error
}


func (aor *ActivityOrderRefund) UpdateShopOrderRefundExpress(refundId int64) error {

	return db.DB.Self.Model(&ShopOrderRefund{}).Where("id = ?",refundId).Update(&aor).Error

}

func GetActivityOrderRefundById(userId, id int64) (*ActivityOrderRefund, error) {

	activityOrderRefund := &ActivityOrderRefund{}

	res := db.DB.Self.
		Where("user_id = ?", userId).
		Where("id = ?", id).
		First(&activityOrderRefund)

	return activityOrderRefund, res.Error

}


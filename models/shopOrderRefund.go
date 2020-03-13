package models

import (
	"soulfire/pkg/db"
	"time"
)

type ShopOrderRefund struct {
	Model
	//OrderGoodsId int64      `json:"order_goods_id" gorm:"column:order_goods_id;not null"`
	RefundN     string     `json:"refund_n" gorm:"column:refund_n;not null"`
	Price       float64    `json:"price" gorm:"column:price;not null"`
	Status      int64      `json:"status" gorm:"column:status;not null"`
	RType       int64      `json:"r_type" gorm:"column:r_type;not null"`
	ReasonPics  string     `json:"reason_pics" gorm:"column:reason_pics;not null"`
	Reason      string     `json:"reason" gorm:"column:reason;not null"`
	ReplyReason string     `json:"reply_reason" gorm:"column:reply_reason;not null"`
	CreatedAt   time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	UserId      int64      `gorm:"column:user_id" json:"user_id"`
	OrderId     int64      `gorm:"column:order_id" json:"order_id"`
	ExpressId   int64      `gorm:"column:express_id" json:"express_id"`
	ExpressN    string     `gorm:"column:express_n" json:"express_n"`
	RWay        int64      `gorm:"column:r_way"  json:"r_way"`
	RStatus     int64      `gorm:"column:r_status"  json:"r_status"`
}

const (
	PendingReview  int64 = iota //发起退款待同意
	CancelRefund                //取消退款
	Refunding                   //退款中(审核通过)
	Refunded                    //退款完成
	RejectedRefund              //拒绝退款
	AgreeRefund                 //同意退款
	PendingPass                 //待审核
)

const (
	OnlyRefund      int64 = 1 //仅退款
	RefundAndReturn int64 = 2 //退款并且退货
)

func (ShopOrderRefund) TableName() string {
	return "shop_order_refunds"
}

func (sor *ShopOrderRefund) Create() (int64, error) {

	res := db.DB.Self.Create(&sor)
	return sor.Id, res.Error

}

func (sor *ShopOrderRefund)Delete(userId,orderId int64) error  {


	return db.DB.Self.Where("user_id = ?",userId).Where("order_id = ?",orderId).Delete(&sor).Error
}


func (sor *ShopOrderRefund) UpdateShopOrderRefundExpress(refundId int64) error {

	return db.DB.Self.Model(&ShopOrderRefund{}).Where("id = ?",refundId).Update(&sor).Error

}

func GetShopOrderRefundById(userId, id int64) (*ShopOrderRefund, error) {

	shopOrderRefund := &ShopOrderRefund{}

	res := db.DB.Self.
		Where("user_id = ?", userId).
		Where("id = ?", id).
		First(&shopOrderRefund)

	return shopOrderRefund, res.Error

}

func GetShopOrderRefundByOrderId(userId, orderId int64) (*ShopOrderRefund, error) {

	shopOrderRefund := &ShopOrderRefund{}

	res := db.DB.Self.
		Where("user_id = ?", userId).
		Where("order_id = ?", orderId).
		First(&shopOrderRefund)

	return shopOrderRefund, res.Error

}

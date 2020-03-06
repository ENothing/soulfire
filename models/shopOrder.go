package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"soulfire/pkg/db"
	"soulfire/utils"
	"time"
)

type ShopOrder struct {
	Model
	UserId            int64           `json:"user_id" gorm:"column:user_id;not null"`
	OrderN            string          `json:"order_n" gorm:"column:order_n;not null"`
	UserCouponId      int64           `json:"user_coupon_id" gorm:"column:user_coupon_id;not null"`
	Num               int64           `json:"num" gorm:"column:num;not null"`
	UnitPrice         float64         `json:"unit_price" gorm:"column:unit_price;not null"`
	TotalPrice        float64         `json:"total_price" gorm:"column:total_price;not null"`
	RealPrice         float64         `json:"real_price" gorm:"column:real_price;not null"`
	DiscountPrice     float64         `json:"discount_price" gorm:"column:discount_price;not null"`
	PostPrice         float64         `json:"post_price" gorm:"column:post_price;not null"`
	Status            int64           `json:"status" gorm:"column:status;not null"`
	Name              string          `json:"name" gorm:"column:name;not null"`
	Mobile            string          `json:"mobile" gorm:"column:mobile;not null"`
	Province          string          `json:"province" gorm:"column:province;not null"`
	City              string          `json:"city" gorm:"column:city;not null"`
	District          string          `json:"district" gorm:"column:district;not null"`
	DetailAddress     string          `json:"detail_address" gorm:"column:detail_address;not null"`
	CreatedAt         time.Time       `gorm:";column:created_at" json:"created_at"`
	UpdatedAt         time.Time       `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt         *time.Time      `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	CompletedAt       time.Time       `gorm:"column:completed_at"  json:"completed_at"`
	PayAt             time.Time       `gorm:"column:pay_at"  json:"pay_at"`
	CancelAt          time.Time       `gorm:"column:cancel_at"  json:"cancel_at"`
	Thumb             string          `json:"thumb" gorm:"column:thumb;not null"`
	GoodsName         string          `json:"goods_name" gorm:"column:goods_name;not null"`
	GoodsSpuName      string          `json:"goods_spu_name" gorm:"column:goods_spu_name;not null"`
	CreatedAtFormat   string          `json:"created_at_format" gorm:"column:created_at_format"`
	UpdatedAtFormat   string          `json:"updated_at_format" gorm:"column:updated_at_format"`
	CompletedAtFormat string          `json:"completed_at_format" gorm:"column:completed_at_format"`
	RefundId          int64           `json:"refund_id" gorm:"column:refund_id"`
	RefundOrder       ShopOrderRefund `json:"shop_order_refund" gorm:"foreignkey:refund_id;PRELOAD:false"`
}

type ShopOrderCreateForm struct {
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
	RefundId      int64      `json:"refund_id" gorm:"column:refund_id"`
}

const (
	PendingPay    int64 = iota //待付款
	CancelOrder                //取消订单
	ToBeDelivered              //已付款待发货
	ToBeReceived               //已发货待收货
	Completed                  //已收货（完成）
)

func (ShopOrder) TableName() string {
	return "shop_orders"
}
func (ShopOrderCreateForm) TableName() string {
	return "shop_orders"
}

func (so *ShopOrder) AfterFind() (err error) {

	so.CreatedAtFormat = utils.TimeFormat(so.CreatedAt, 0)
	so.UpdatedAtFormat = utils.TimeFormat(so.UpdatedAt, 0)
	so.CompletedAtFormat = utils.TimeFormat(so.CompletedAt, 0)
	return
}

func (sof *ShopOrderCreateForm) Create(transaction *gorm.DB) (id int64, err error) {

	err = transaction.Create(&sof).Error
	id = sof.Id
	return

}

func GetOrderById(userId, orderId int64) (*ShopOrder, error) {

	order := &ShopOrder{}

	res := db.DB.Self.
		Where("shop_orders.id = ?", orderId).
		Where("shop_orders.user_id = ?", userId).
		Joins("LEFT JOIN shop_order_goods as sog ON sog.order_id = shop_orders.id").
		Joins("LEFT JOIN shop_goods as sg ON sg.id = sog.goods_id").
		Joins("LEFT JOIN shop_goods_spus as sgs ON sgs.id = sog.spu_id").
		Select("shop_orders.*,sg.thumb as thumb,sg.name as goods_name,sgs.name as goods_spu_name").
		First(&order)

	return order, res.Error

}

func UpdateOrderStatusToCancel(userId, orderId int64) error {

	nowTime := utils.TimeFormat(time.Now(), 0)

	return db.DB.Self.Model(&ShopOrderCreateForm{}).
		Where("user_id = ?", userId).
		Where("id = ?", orderId).
		Where("status = ?", 0).
		Update("status", CancelOrder).Update("cancel_at", nowTime).Error

}

func UpdateOrderRefundId(userId, orderId, refundId int64) error {

	return db.DB.Self.Model(&ShopOrderCreateForm{}).
		Where("user_id = ?", userId).
		Where("id = ?", orderId).
		Update("refund_id", refundId).Error

}

func GetOrderDetailById(userId, orderId int64) (interface{}, error) {

	data := make(map[string]interface{})
	order := &ShopOrder{}
	refundOrder := &ShopOrderRefund{}

	res := db.DB.Self.
		Where("shop_orders.id = ?", orderId).
		Where("shop_orders.user_id = ?", userId).
		Joins("LEFT JOIN shop_order_goods as sog ON sog.order_id = shop_orders.id").
		Joins("LEFT JOIN shop_goods as sg ON sg.id = sog.goods_id").
		Joins("LEFT JOIN shop_goods_spus as sgs ON sgs.id = sog.spu_id").
		Select("shop_orders.*,sg.thumb as thumb,sg.name as goods_name,sgs.name as goods_spu_name").
		First(&order)

	refundOrderRes := db.DB.Self.Model(&ShopOrderRefund{}).
		Where("order_id = ?", orderId).
		Where("user_id = ?", userId).
		First(&refundOrder)

	data["order"] = order

	if refundOrderRes.Error != nil || refundOrderRes.Error == gorm.ErrRecordNotFound {

		data["refund_order"] = nil

	} else {
		data["refund_order"] = refundOrder
	}

	return data, res.Error

}

func ShopOrderPaginate(page int64, pageSize int64, userId int64, status string) (shopOrder []*ShopOrder, total int64, lastPage int64, err error) {

	shopOrder = make([]*ShopOrder, 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self

	res = res.Where("shop_orders.user_id = ?", userId)

	if status != "" {

		if status == "5" { //退款退货

			res = res.Where("shop_orders.refund_id != ?", 0)

		} else {

			res = res.Where("shop_orders.status = ?", status).Where("shop_orders.refund_id = ?", 0)

		}

	}

	res = res.
		Joins("LEFT JOIN shop_order_goods as sog ON sog.order_id = shop_orders.id").
		Joins("LEFT JOIN shop_goods as sg ON sg.id = sog.goods_id").
		Joins("LEFT JOIN shop_goods_spus as sgs ON sgs.id = sog.spu_id").
		Preload("RefundOrder").
		Select("shop_orders.*,sg.thumb as thumb,sg.name as goods_name,sgs.name as goods_spu_name").
		Order("created_at desc").
		Limit(pageSize).
		Offset(offset).
		Find(&shopOrder)

	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return shopOrder, total, lastPage, res.Error
}

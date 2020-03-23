package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"soulfire/pkg/db"
	"soulfire/utils"
	"time"
)

type ActivityOrder struct {
	Model
	UserId        int64      `json:"user_id" gorm:"column:user_id;not null"`
	ActivityId    int64      `json:"activity_id" gorm:"column:activity_id;not null"`
	RefundId      int64      `json:"refund_id" gorm:"column:refund_id;not null"`
	OrderN        string     `json:"order_n" gorm:"column:order_n;not null"`
	Name          string     `json:"name" gorm:"column:name;not null"`
	Sex           int64      `json:"sex" gorm:"column:sex;not null"`
	Mobile        string     `json:"mobile" gorm:"column:mobile;not null"`
	UnitPrice     float64    `json:"unit_price" gorm:"column:unit_price;not null"`
	PersonNum     int64      `json:"person_num" gorm:"column:person_num;not null"`
	TotalPrice    float64    `json:"total_price" gorm:"column:total_price;not null"`
	RealPrice     float64    `json:"real_price" gorm:"column:real_price;not null"`
	DiscountPrice float64    `gorm:"column:discount_price" json:"discount_price"`
	Status        int64      `gorm:"column:status" json:"status"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	CType         int64      `gorm:"column:c_type" json:"c_type"`
	CNum          string     `gorm:"column:c_num" json:"c_num"`
}

type ActivityCompleted struct {
	ActivityOrder
	CompletedAt   time.Time  `gorm:"column:completed_at"  json:"completed_at"`
}

type ActivityOrderRelated struct {
	ActivityCompleted
	Activity          Activity            `json:"activity" gorm:"foreignkey:ActivityId;PRELOAD:false"`
	RefundOrder       ActivityOrderRefund `json:"activity_refund_order" gorm:"foreignkey:RefundId;PRELOAD:false"`
	CreatedAtFormat   string              `json:"created_at_format" gorm:"column:created_at_format"`
	CompletedAtFormat string              `json:"completed_at_format" gorm:"column:completed_at_format"`

	//Refund      Activity   `json:"activity" gorm:"foreignkey:ActivityId"`
}



func (ActivityOrder) TableName() string {
	return "activity_orders"
}
func (ActivityOrderRelated) TableName() string {
	return "activity_orders"
}
//func (ao *ActivityOrder) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("start_time", time.Now())
//	return nil
//}
func (ao *ActivityOrder) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("completed_at", nil)
	return nil
}
func (ao *ActivityOrderRelated) AfterFind() (err error) {

	ao.CreatedAtFormat = utils.TimeFormat(ao.CreatedAt, 1)
	ao.CompletedAtFormat = utils.TimeFormat(ao.CompletedAt, 0)
	return
}

func (ao *ActivityOrder) Create() (id int64, err error) {

	res := db.DB.Self.Create(&ao)
	return ao.Id, res.Error

}

func GetActivityOrderById(id, userId int64) (*ActivityOrderRelated, error) {

	activityOrder := &ActivityOrderRelated{}

	res := db.DB.Self.Where("id = ?", id).Where("user_id = ?", userId).Preload("Activity").Preload("RefundOrder").First(&activityOrder)

	//db.DB.Self.Model(&activityOrder).Select([]string{"title,thumb"}).Related(&activityOrder.Activity)

	return activityOrder, res.Error

}

func UpdateActivityOrderRefundId(userId, orderId, refundId int64) error {

	return db.DB.Self.Model(&ActivityOrder{}).
		Where("user_id = ?", userId).
		Where("id = ?", orderId).
		Update("refund_id", refundId).Error

}

func UpdateActivityOrderToFinished(userId, orderId int64) error {

	return db.DB.Self.Model(&ActivityCompleted{}).
		Where("user_id = ?", userId).
		Where("id = ?", orderId).
		Update("status", 3).Update("completed_at",time.Now()).Error

}

func ActivityOrderPaginate(page int64, pageSize int64, userId int64, status string) (activityOrder []*ActivityOrderRelated, total int64, lastPage int64, err error) {

	activityOrder = make([]*ActivityOrderRelated, 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self

	res = res.Where("user_id = ?", userId)

	if status != "" {

		if status == "4" { //退款退货

			res = res.Where("refund_id != ?", 0)

		} else {

			res = res.Where("status = ?", status).Where("refund_id = ?", 0)

		}

	}

	res = res.Preload("Activity").Preload("RefundOrder").Order("created_at desc").Limit(pageSize).Offset(offset).Find(&activityOrder)
	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return activityOrder, total, lastPage, res.Error
}

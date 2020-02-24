package models

import (
	"math"
	"soulfire/pkg/db"
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
	DiscountPrice float64    `gorm:";column:discount_price" json:"discount_price"`
	Status        int64      `gorm:";column:status" json:"status"`
	CreatedAt     time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	Activity      Activity   `json:"activity" gorm:"foreignkey:ActivityId"`
}

func (ActivityOrder) TableName() string {
	return "activity_orders"
}

func (ao *ActivityOrder) Create() error {

	return db.DB.Self.Create(&ao).Error

}

func GetActivityOrderById(id,userId int64) (*ActivityOrder, error) {

	activityOrder := &ActivityOrder{}

	db.DB.Self.Where("id = ?", id).Where("user_id = ?",userId).First(&activityOrder)
	res := db.DB.Self.Model(&activityOrder).Select([]string{"title,thumb,kind"}).Related(&activityOrder.Activity)

	return activityOrder, res.Error

}

func ActivityOrderPaginate(page int64, pageSize int64, userId int64, status string) (activity []*Activity, total int64, lastPage int64, err error) {

	activity = make([]*Activity, 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self

	res = res.Where("user_id = ?", userId)

	if status != "" {

		if status == "4" {//退款退货

			res = res.Where("refund_id != ?", 0)

		}else{

			res = res.Where("status = ?", status)

		}

	}

	res = res.Order("created_at desc").Limit(pageSize).Offset(offset).Find(&activity)
	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return activity, total, lastPage, res.Error
}

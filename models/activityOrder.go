package models

import (
	"math"
	"soulfire/pkg/db"
	"time"
)

type ActivityOrder struct {
	Model
	UserId        int64     `json:"user_id" gorm:"column:user_id;not null"`
	ActivityId    int64     `json:"activity_id" gorm:"column:activity_id;not null"`
	RefundId      int64     `json:"refund_id" gorm:"column:refund_id;not null"`
	OrderN        string     `json:"order_n" gorm:"column:order_n;not null"`
	Name          string     `json:"name" gorm:"column:name;not null"`
	Sex           int64     `json:"sex" gorm:"column:sex;not null"`
	Mobile        string     `json:"mobile" gorm:"column:mobile;not null"`
	UnitPrice     string     `json:"unit_price" gorm:"column:unit_price;not null"`
	PersonNum     int64     `json:"person_num" gorm:"column:person_num;not null"`
	TotalPrice    string     `json:"total_price" gorm:"column:total_price;not null"`
	RealPrice     string     `json:"real_price" gorm:"column:real_price;not null"`
	DiscountPrice string  `gorm:";column:discount_price" json:"discount_price"`
	Status        int64  `gorm:";column:status" json:"status"`
	CreatedAt     time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func (ActivityOrder) TableName() string {
	return "activity_orders"
}

func (ao *ActivityOrder) Create() error {

	return db.DB.Self.Create(&ao).Error

}

func ActivityOrderPaginate(page int64, pageSize int64, sort int64, cateId int64, title string) (activity []*Activity, total int64, lastPage int64, err error) {

	activity = make([]*Activity, 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self

	if cateId != 0 {
		res = res.Where("cate_id = ?", cateId)
	}
	if title != "" {
		res = res.Where("title LIKE ?", "%"+title+"%")
	}

	if sort == 0 {
		res = res.Order("created_at desc")
	} else {
		res = res.Order("created_at asc")
	}

	res = res.Limit(pageSize).Offset(offset).Find(&activity)
	db.DB.Self.Model(&activity).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return activity, total, lastPage, res.Error
}

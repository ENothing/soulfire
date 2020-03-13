package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"soulfire/pkg/db"
	"soulfire/utils"
	"time"
)

type Activity struct {
	Model
	Title           string     `json:"title" gorm:"column:title;not null"`
	Intro           string     `json:"intro" gorm:"column:intro;not null"`
	Thumb           string     `json:"thumb" gorm:"column:thumb;not null"`
	CateId          int64      `json:"cate_id" gorm:"column:cate_id;not null"`
	Content         string     `json:"content" gorm:"column:content;not null"`
	Kind            int64      `json:"kind" gorm:"column:kind;not null"`
	CurPrice        float64    `json:"cur_price" gorm:"column:cur_price;not null"`
	OriPrice        float64    `json:"ori_price" gorm:"column:ori_price;not null"`
	StartAt         time.Time  `json:"start_at" gorm:"column:start_at;not null"`
	EndAt           time.Time  `json:"end_at" gorm:"column:end_at;not null"`
	StartEnterAt    time.Time  `json:"start_enter_at" gorm:"column:start_enter_at;not null"`
	EndEnterAt      time.Time  `json:"end_enter_at" gorm:"column:end_enter_at;not null"`
	PersonLimit     int64      `json:"person_limit" gorm:"column:person_limit;not null"`
	View            int64      `json:"view" gorm:"column:view;not null"`
	Likes           int64      `json:"likes" gorm:"column:likes;not null"`
	Favor           int64      `json:"favor" gorm:"column:favor;not null"`
	Sold            int64      `json:"sold" gorm:"column:sold;not null"`
	IsPublish       int64      `json:"is_publish" gorm:"column:is_publish;not null"`
	CreatedAt       time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt       time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	CreatedAtFormat string     `json:"created_at_format" gorm:"column:created_at_format"`
	StartAtFormat   string     `json:"start_at_format" gorm:"column:start_at_format"`
	EndAtFormat     string     `json:"end_at_format" gorm:"column:end_at_format"`
	Mobile          string     `json:"mobile" gorm:"column:mobile"`
	DetailAddress   string     `json:"detail_address" gorm:"column:detail_address"`
	ChargeType      int64      `json:"charge_type" gorm:"column:charge_type"`
}

func (Activity) TableName() string {
	return "activities"
}

func (a *Activity) AfterFind() (err error) {

	a.CreatedAtFormat = utils.TimeFormat(a.CreatedAt, 1)
	a.StartAtFormat = utils.TimeFormat(a.StartAt, 0)
	a.EndAtFormat = utils.TimeFormat(a.EndAt, 0)
	return
}
func ActivityViewAddOne(id int64) error {

	activity := &Activity{}

	res := db.DB.Self.Model(&activity).
		Where("id = ?", id).
		Where("view > 0").
		Update("view", gorm.Expr("view + ?", 1))

	return res.Error

}

func ActivityLikeAddOne(id int64) error {

	activity := &Activity{}

	res := db.DB.Self.Model(&activity).
		Where("id = ?", id).
		Where("likes > 0").
		UpdateColumn("likes", gorm.Expr("likes + ?", 1))

	return res.Error

}

func ActivityLikeCutOne(id int64) error {

	activity := &Activity{}

	res := db.DB.Self.Model(&activity).
		Where("id = ?", id).
		Where("likes > 0").
		UpdateColumn("likes", gorm.Expr("likes - ?", 1))

	return res.Error

}

func ActivityFavorAddOne(id int64) error {

	activity := &Activity{}

	res := db.DB.Self.Model(&activity).
		Where("id = ?", id).
		Where("favor > 0").
		UpdateColumn("favor", gorm.Expr("favor + ?", 1))

	return res.Error

}

func ActivityFavorCutOne(id int64) error {

	activity := &Activity{}

	res := db.DB.Self.Model(&activity).
		Where("id = ?", id).
		Where("favor > 0").
		UpdateColumn("favor", gorm.Expr("favor - ?", 1))

	return res.Error

}

func ActivitySoldOne(id int64) error {

	activity := &Activity{}

	res := db.DB.Self.Model(&activity).
		Where("id = ?", id).
		UpdateColumn("sold", gorm.Expr("sold + ?", 1))

	return res.Error

}

func GetActivityById(id int64) (*Activity, error) {

	activity := &Activity{}

	res := db.DB.Self.Where("id = ?", id).First(&activity)

	return activity, res.Error

}

func ActivityPaginate(page int64, pageSize int64, sort int64, cateId int64, title string) (activity []*Activity, total int64, lastPage int64, err error) {

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
	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return activity, total, lastPage, res.Error
}

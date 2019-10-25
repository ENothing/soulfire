package models

import (
	"gin-init/pkg/db"
	"math"
	"time"
)

type Activity struct {
	Model
	Title        string     `json:"title" gorm:"column:title;not null"`
	Thumb        string     `json:"thumb" gorm:"column:thumb;not null"`
	CateId       string     `json:"cate_id" gorm:"column:cate_id;not null"`
	Content      string     `json:"content" gorm:"column:content;not null"`
	StartAt      string     `json:"start_at" gorm:"column:start_at;not null"`
	EndAt        string     `json:"end_at" gorm:"column:end_at;not null"`
	StartEnterAt string     `json:"start_enter_at" gorm:"column:start_enter_at;not null"`
	EndEnterAt   string     `json:"end_enter_at" gorm:"column:end_enter_at;not null"`
	PersonLimit  string     `json:"person_limit" gorm:"column:person_limit;not null"`
	View         string     `json:"view" gorm:"column:view;not null"`
	Like         string     `json:"like" gorm:"column:like;not null"`
	CreatedAt    time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func (Activity) TableName() string {
	return "activities"
}

func GetActivityById(id int64)(*Activity,error)  {

	activity := &Activity{}

	res := db.DB.Self.Where("id = ?",id).First(&activity)

	return activity,res.Error

}


func Paginate(page int64, pageSize int64, sort int64, cateId int64, title string) (activity []*Activity,total int64,lastPage int64, err error) {

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

	return activity, total,lastPage, res.Error
}

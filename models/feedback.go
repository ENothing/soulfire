package models

import (
	"soulfire/pkg/db"
	"time"
)

type Feedback struct {
	Model
	Title     string    `json:"title" gorm:"column:title;not null"`
	Content   string    `json:"content" gorm:"column:content;not null"`
	UserId    int64     `json:"user_id" gorm:"column:user_id;not null"`
	CreatedAt time.Time `gorm:";column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:";column:updated_at" json:"updated_at"`
	Pics      string    `gorm:";column:pics" json:"pics"`
}

func (Feedback) TableName() string {
	return "feedback"
}

func (f *Feedback) Create() (err error) {
	return db.DB.Self.Create(&f).Error
}

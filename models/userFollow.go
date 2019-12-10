package models

import (
	"github.com/jinzhu/gorm"
	"soulfire/pkg/db"
	"time"
)

type UserFollow struct {
	Model
	UserId    int64     `json:"user_id" gorm:"column:user_id;not null"`
	FollowId  int64     `json:"follow_id" gorm:"column:follow_id;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null"`
}

func (UserFollow) TableName() string {
	return "user_follows"
}

func (uf *UserFollow) Create() error {

	res := db.DB.Self.Create(&uf)

	return res.Error

}

func (uf *UserFollow) Delete() error {

	res := db.DB.Self.Where("user_id = ?", uf.UserId).Where("follow_id = ?", uf.FollowId).Delete(&uf)

	return res.Error
}

func GetUserFollowById(userId, followId int64) bool {

	userFollow := &UserFollow{}

	res := db.DB.Self.Where("user_id = ?", userId).Where("follow_id = ?", followId).First(&userFollow)

	if err := res.Error; err != nil || err == gorm.ErrRecordNotFound {

		return false

	} else {

		return true

	}

}

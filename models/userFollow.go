package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"soulfire/pkg/db"
	"soulfire/utils"
	"time"
)

type UserFollow struct {
	Model
	UserId    int64     `json:"user_id" gorm:"column:user_id;not null"`
	FollowId  int64     `json:"follow_id" gorm:"column:follow_id;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null"`
}

type UserInfoFollow struct {
	UserFollow
	Avatar          string `json:"avatar" gorm:"column:avatar;not null"`
	NickName        string `json:"nickname" gorm:"column:nickname;not null"`
	CreatedAtFormat string `json:"created_at_format" gorm:"column:created_at_format"`
	Attention       int64  `json:"attention" gorm:"column:attention;not null"`
}



func (UserFollow) TableName() string {
	return "user_follows"
}

func (uif *UserInfoFollow) AfterFind() (err error) {

	uif.CreatedAtFormat = utils.TimeFormat(uif.CreatedAt, 0)
	return
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

func GetFollowsPaginate(page int64, pageSize int64, userId int64) (userFollow []*UserInfoFollow, total int64, lastPage int64, err error) {

	offset := (page - 1) * pageSize

	res := db.DB.Self.Where("user_follows.user_id = ?", userId).
		Joins("LEFT JOIN users as u ON u.id = user_follows.follow_id")

	res = res.
		Order("created_at desc").
		Limit(pageSize).Offset(offset).
		Select("user_follows.*,u.head_url as avatar,u.nickname as nickname").
		Find(&userFollow)
	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return userFollow, total, lastPage, res.Error

}
func GetUserFollowsPaginate(page int64, pageSize int64, userId int64,id int64) (userFollow []*UserInfoFollow, total int64, lastPage int64, err error) {

	offset := (page - 1) * pageSize

	res := db.DB.Self.Where("user_follows.user_id = ?", id).
		Joins("LEFT JOIN users as u ON u.id = user_follows.follow_id").
		Joins("LEFT JOIN user_follows as uf ON uf.follow_id = user_follows.follow_id  AND  uf.user_id = ? ",userId)

	res = res.
		Order("created_at desc").
		Limit(pageSize).Offset(offset).
		Select("user_follows.*,u.head_url as avatar,u.nickname as nickname,if(uf.id is null,0,1) as attention").
		Find(&userFollow)
	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return userFollow, total, lastPage, res.Error

}

func GetFollowedPaginate(page int64, pageSize int64, userId int64) (userFollow []*UserInfoFollow, total int64, lastPage int64, err error) {

	offset := (page - 1) * pageSize

	res := db.DB.Self.Where("user_follows.follow_id = ?", userId).
		Joins("LEFT JOIN users as u ON u.id = user_follows.user_id").
		Joins("LEFT JOIN user_follows as uf ON uf.follow_id = user_follows.user_id")

	res = res.
		Order("created_at desc").
		Limit(pageSize).Offset(offset).
		Select("user_follows.*,u.head_url as avatar,u.nickname as nickname,if(uf.id is null,0,1) as attention").
		Find(&userFollow)
	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return userFollow, total, lastPage, res.Error

}

func GetUserFollowedPaginate(page int64, pageSize int64, userId int64,id int64) (userFollow []*UserInfoFollow, total int64, lastPage int64, err error) {

	offset := (page - 1) * pageSize

	res := db.DB.Self.Where("user_follows.follow_id = ?", id).
		Joins("LEFT JOIN users as u ON u.id = user_follows.user_id").
		Joins("LEFT JOIN user_follows as uf ON uf.follow_id = user_follows.user_id")

	res = res.
		Order("created_at desc").
		Limit(pageSize).Offset(offset).
		Select("user_follows.*,u.head_url as avatar,u.nickname as nickname,if(uf.id is null,0,1) as attention").
		Find(&userFollow)
	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return userFollow, total, lastPage, res.Error

}

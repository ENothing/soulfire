package models

import (
	"github.com/jinzhu/gorm"
	"soulfire/pkg/db"
	"time"
)

type User struct {
	Model
	Username   string     `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	ParentId   int64      `json:"parent_id" gorm:"column:parent_id;not null"`
	Openid     string     `json:"openid" gorm:"column:openid;not null"`
	NickName   string     `json:"nickname" gorm:"column:nickname;not null"`
	Gender     int64      `json:"gender" gorm:"column:gender;not null"`
	HeadUrl    string     `json:"head_url" gorm:"column:head_url;not null"`
	Mobile     string     `json:"mobile" gorm:"column:mobile;not null"`
	Email      string     `json:"email" gorm:"column:email;not null"`
	IsBan      int64      `json:"is_ban" gorm:"column:is_ban;not null"`
	Sign       string     `json:"sign" gorm:"column:sign;not null"`
	Follows    int64      `json:"follows" gorm:"column:follows;not null"`
	CreatedAt  time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	IsFollowed int64      `json:"is_followed" gorm:"column:is_followed;not null"`
}

type UserDetail struct {
	User
	Likes int64 `json:"likes" gorm:"column:likes;not null"`
	Attention int64 `json:"attention" gorm:"column:attention;not null"`
	IsMe int64 `json:"is_me" gorm:"column:is_me;not null"`
}

func (User) TableName() string {
	return "users"
}

func GetUserByOpenid(openid string) (*User, error) {

	user := &User{}

	res := db.DB.Self.Where("openid = ?", openid).First(&user)

	return user, res.Error

}

func GetUserInfoById(id int64)(*User, error){

	user := &User{}

	res := db.DB.Self.Where("id = ?", id).Select("id,nickname,head_url,gender,follows, is_followed").First(&user)

	return user, res.Error

}

func (ud *UserDetail) GetUserById(id, userId int64) (*UserDetail, error) {

	res := db.DB.Self.
		Where("users.id = ?", id).
		Joins("LEFT JOIN articles AS a ON a.user_id = users.id").
		Select("users.id as id,users.nickname as nickname,users.gender as gender,users.follows as follows,users.is_followed as is_followed,users.head_url as head_url,users.sign as sign,SUM(a.likes) as likes").
		First(&ud)

	//db.DB.Self.Model(&UserFollow{}).Where("follow_id = ?", id).Count(&ud.Follows)

	isFollowed := GetUserFollowById(userId, ud.Id)

	attention := 0
	if isFollowed {
		attention = 1
	}

	isMe := 0
	if id == userId {
		isMe = 1
	}

	ud.Attention = int64(attention)
	ud.IsMe = int64(isMe)

	return ud, res.Error

}

func (u *User) Create() (id int64, err error) {

	err = db.DB.Self.Create(&u).Error
	id = u.Id

	return
}

func Delete(ids []string) error {

	u := User{}

	return db.DB.Self.Where("id in (?)", ids).Delete(&u).Error
}

func FollowsAddOne(id int64) error {

	user := &User{}

	res := db.DB.Self.Model(&user).
		Where("id = ?", id).
		Where("follows >= 0").
		UpdateColumn("follows", gorm.Expr("follows + ?", 1))

	return res.Error

}

func FollowsCutOne(id int64) error {

	user := &User{}

	res := db.DB.Self.Model(&user).
		Where("id = ?", id).
		Where("follows > 0").
		UpdateColumn("follows", gorm.Expr("follows - ?", 1))

	return res.Error

}
func IsFollowedAddOne(id int64) error {

	user := &User{}

	res := db.DB.Self.Model(&user).
		Where("id = ?", id).
		Where("is_followed >= 0").
		UpdateColumn("is_followed", gorm.Expr("is_followed + ?", 1))

	return res.Error

}

func IsFollowedCutOne(id int64) error {

	user := &User{}

	res := db.DB.Self.Model(&user).
		Where("id = ?", id).
		Where("is_followed > 0").
		UpdateColumn("is_followed", gorm.Expr("is_followed - ?", 1))

	return res.Error

}

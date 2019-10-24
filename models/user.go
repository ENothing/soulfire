package models

import (
	"fmt"
	"gin-init/pkg/db"
	"time"
)

type User struct {
	Model
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	ParentId int64 `json:"parent_id" gorm:"column:parent_id;not null"`
	Openid string `json:"openid" gorm:"column:openid;not null"`
	Nickname string `json:"nickname" gorm:"column:nickname;not null"`
	HeadUrl string `json:"head_url" gorm:"column:head_url;not null"`
	Mobile string `json:"mobile" gorm:"column:mobile;not null"`
	Email string `json:"email" gorm:"column:email;not null"`
	IsBan string `json:"is_ban" gorm:"column:is_ban;not null"`
	CreatedAt time.Time `gorm:";column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func (u *User)GetUserByOpenid(openid string)  {

	res := db.DB.Self.Where("openid = ?",openid).First(&u)


	fmt.Println(res)

}



func (u *User) Create() (id uint64,err error) {

	err = db.DB.Self.Create(&u).Error
	id = u.Id

	return
}

func Delete(ids []string) error {

	u := User{}

	return  db.DB.Self.Where("id in (?)",ids).Delete(&u).Error
}



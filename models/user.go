package models

import "gin-init/pkg/db"

type User struct {
	Model
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=6,max=128"`
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



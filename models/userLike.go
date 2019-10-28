package models

import "soulfire/pkg/db"

type UserLike struct {
	Model
	Own    int64 `json:"own" gorm:"column:own;not null"`
	TypeId int64 `json:"type_id" gorm:"column:type_id;not null"`
	UserId int64 `json:"user_id" gorm:"column:user_id;not null"`
}

func (UserLike) TableName() string {
	return "user_likes"
}

func LikeAndUnlike(userId, typeId, own int64) bool {

	userLike := &UserLike{}

	res := db.DB.Self.
		Where("user_id = ?", userId).
		Where("type_id = ?", typeId).
		Where("own = ?", own).
		First(&userLike)

	if res.Error != nil {

		userLike.Own = own
		userLike.TypeId = typeId
		userLike.UserId = userId

		db.DB.Self.Create(&userLike)

		return true

	}

	db.DB.Self.Delete(&userLike)

	return false

}

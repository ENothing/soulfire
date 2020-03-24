package models

import (
	"fmt"
	"math"
	"soulfire/pkg/db"
)

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

func GetArticleLikePaginate(page int64, pageSize int64, userId int64) (article []*Article, total int64, lastPage int64, err error) {

	var typeIds []int64
	var userLike []*UserLike
	article = make([]*Article, 0)

	db.DB.Self.Where("user_id = ?", userId).Where("own = ?", IsArticle).Find(&userLike).Pluck("type_id", &typeIds)

	fmt.Println(typeIds)

	offset := (page - 1) * pageSize

	res := db.DB.Self.Model(&Article{}).Where("id IN (?)", typeIds)

	res = res.Order("created_at desc")

	res = res.Limit(pageSize).Offset(offset).Find(&article)
	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return article, total, lastPage, res.Error

}

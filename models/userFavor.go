package models

import (
	"math"
	"soulfire/pkg/db"
)

type UserFavor struct {
	Model
	Own    int64 `json:"own" gorm:"column:own;not null"`
	TypeId int64 `json:"type_id" gorm:"column:type_id;not null"`
	UserId int64 `json:"user_id" gorm:"column:user_id;not null"`
}

const (
	IsActivity int64 = 1
	IsArticle  int64 = 2
)

func (UserFavor) TableName() string {
	return "user_favors"
}

func FavorAndUnFavor(userId, typeId, own int64) bool {

	userFavor := &UserFavor{}

	res := db.DB.Self.
		Where("user_id = ?", userId).
		Where("type_id = ?", typeId).
		Where("own = ?", own).
		First(&userFavor)

	if res.Error != nil {

		userFavor.Own = own
		userFavor.TypeId = typeId
		userFavor.UserId = userId

		db.DB.Self.Create(&userFavor)

		return true

	}

	db.DB.Self.Delete(&userFavor)

	return false

}

func GetActivityFavorPaginate(page int64, pageSize int64, userId int64) (activity []*Activity, total int64, lastPage int64, err error) {

	var typeIds []int64
	activity = make([]*Activity, 0)

	db.DB.Self.Where("user_id = ?", userId).Where("own = ?", IsActivity).Find(&UserFavor{}).Pluck("type_id", &typeIds)

	offset := (page - 1) * pageSize

	res := db.DB.Self.Model(&Activity{}).Where("id IN (?)", typeIds)

	res = res.Order("created_at desc")

	res = res.Limit(pageSize).Offset(offset).Find(&activity)
	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return activity, total, lastPage, res.Error

}

func GetArticleFavorPaginate(page int64, pageSize int64, userId int64) (article []*Article, total int64, lastPage int64, err error) {

	var typeIds []int64
	article = make([]*Article, 0)

	db.DB.Self.Where("user_id = ?", userId).Where("own = ?", IsArticle).Find(&UserFavor{}).Pluck("type_id", &typeIds)

	offset := (page - 1) * pageSize

	res := db.DB.Self.Model(&Article{}).Where("id IN (?)", typeIds)

	res = res.Order("created_at desc")

	res = res.Limit(pageSize).Offset(offset).Find(&article)
	res.Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return article, total, lastPage, res.Error

}

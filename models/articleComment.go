package models

import (
	"math"
	"soulfire/pkg/db"
	"time"
)

type ArticleComment struct {
	Model
	UserId    int64      `json:"user_id" gorm:"column:user_id;not null"`
	ArticleId int64      `json:"article_id" gorm:"column:article_id;not null"`
	Content   string     `json:"content" gorm:"column:content;not null"`
	ParentId int64      `json:"parent_id" gorm:"column:parent_id;not null"`
	CreatedAt time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	ArticleComments []ArticleComment `gorm:"foreignkey:ArticleId"`
}

func (ArticleComment) TableName() string {
	return "article_comments"
}

func (ac *ArticleComment) Create() error {

	return db.DB.Self.Create(&ac).Error

}
//func ArticleViewAddOne(id int64) error {
//
//	article := &Article{}
//
//	res := db.DB.Self.Model(&article).
//		Where("id = ?", id).
//		Where("view > 0").
//		UpdateColumn("view", gorm.Expr("view + ?", 1))
//
//	return res.Error
//
//}

//func ArticleLikeAddOne(id int64) error {
//
//	article := &Article{}
//
//	res := db.DB.Self.Model(&article).
//		Where("id = ?", id).
//		Where("likes > 0").
//		UpdateColumn("likes", gorm.Expr("likes + ?", 1))
//
//	return res.Error
//
//}

//func ArticleLikeCutOne(id int64) error {
//
//	article := &Article{}
//
//	res := db.DB.Self.Model(&article).
//		Where("id = ?", id).
//		Where("likes > 0").
//		UpdateColumn("likes", gorm.Expr("likes - ?", 1))
//
//	return res.Error
//
//}

//func GetArticleById(id int64) (*Article, error) {
//
//	article := &Article{}
//
//	res := db.DB.Self.Where("id = ?", id).First(&article)
//
//	return article, res.Error
//
//}

func ArticleCommentPaginate(page int64, pageSize int64, articleId int64) (articleComments []*ArticleComment, total int64, lastPage int64, err error) {

	articleComments = make([]*ArticleComment, 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self.Where("article_id = ?" ,articleId).Related(&articleComments)

	res = res.Limit(pageSize).Offset(offset).Find(&articleComments)
	db.DB.Self.Model(&articleComments).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return articleComments, total, lastPage, res.Error
}

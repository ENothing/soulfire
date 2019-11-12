package models

import (
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

//func ArticlePaginate(page int64, pageSize int64, sort int64, cateId int64, title string) (articles []*Article, total int64, lastPage int64, err error) {
//
//	articles = make([]*Article, 0)
//
//	offset := (page - 1) * pageSize
//
//	res := db.DB.Self
//
//	if cateId != 0 {
//		res = res.Where("cate_id = ?", cateId)
//	}
//	if title != "" {
//		res = res.Where("title LIKE ?", "%"+title+"%")
//	}
//
//	if sort == 0 {
//		res = res.Order("created_at desc")
//	} else {
//		res = res.Order("created_at asc")
//	}
//
//	res = res.Limit(pageSize).Offset(offset).Find(&articles)
//	db.DB.Self.Model(&articles).Count(&total)
//
//	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))
//
//	return articles, total, lastPage, res.Error
//}

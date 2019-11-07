package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"soulfire/pkg/db"
	"time"
)

type Article struct {
	Model
	UserId    string     `json:"user_id" gorm:"column:user_id;not null"`
	Thumb     string     `json:"thumb" gorm:"column:thumb;not null"`
	Title     int64      `json:"title" gorm:"column:title;not null"`
	Content   string     `json:"content" gorm:"column:content;not null"`
	Likes     string     `json:"likes" gorm:"column:likes;not null"`
	View      string     `json:"view" gorm:"column:view;not null"`
	CateId    string     `json:"cate_id" gorm:"column:cate_id;not null"`
	IsPublish string     `json:"is_publish" gorm:"column:is_publish;not null"`
	CreatedAt time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func (Article) TableName() string {
	return "articles"
}

func ArticleViewAddOne(articleId int64) error {

	article := &Article{}

	res := db.DB.Self.Model(&article).
		Where("id = ?", articleId).
		Where("view > 0").
		UpdateColumn("view", gorm.Expr("view + ?", 1))

	return res.Error

}

func ArticleLikeAddOne(activityId int64) error {

	article := &Article{}

	res := db.DB.Self.Model(&article).
		Where("id = ?", activityId).
		Where("likes > 0").
		UpdateColumn("likes", gorm.Expr("likes + ?", 1))

	return res.Error

}

func ArticleLikeCutOne(activityId int64) error {

	article := &Article{}

	res := db.DB.Self.Model(&article).
		Where("id = ?", activityId).
		Where("likes > 0").
		UpdateColumn("likes", gorm.Expr("likes - ?", 1))

	return res.Error

}

func GetArticleById(id int64) (*Article, error) {

	article := &Article{}

	res := db.DB.Self.Where("id = ?", id).First(&article)

	return article, res.Error

}

func ArticlePaginate(page int64, pageSize int64, sort int64, cateId int64, title string) (articles []*Article, total int64, lastPage int64, err error) {

	articles = make([]*Article, 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self

	if cateId != 0 {
		res = res.Where("cate_id = ?", cateId)
	}
	if title != "" {
		res = res.Where("title LIKE ?", "%"+title+"%")
	}

	if sort == 0 {
		res = res.Order("created_at desc")
	} else {
		res = res.Order("created_at asc")
	}

	res = res.Limit(pageSize).Offset(offset).Find(&articles)
	db.DB.Self.Model(&articles).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return articles, total, lastPage, res.Error
}

package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"soulfire/pkg/db"
	"time"
)

type Article struct {
	Model
	UserId     int64      `json:"user_id" gorm:"column:user_id;not null"`
	Thumb      string     `json:"thumb" gorm:"column:thumb;not null"`
	Title      string     `json:"title" gorm:"column:title;not null"`
	Content    string     `json:"content" gorm:"column:content;not null"`
	Likes      int64      `json:"likes" gorm:"column:likes;not null"`
	View       int64      `json:"view" gorm:"column:view;not null"`
	CateId     int64      `json:"cate_id" gorm:"column:cate_id;not null"`
	IsPublish  int64      `json:"is_publish" gorm:"column:is_publish;not null"`
	IsFollowed int64      `json:"is_followed" gorm:"column:is_followed;not null"`
	CreatedAt  time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`

	NickName string `json:"nickname" gorm:"column:nickname;not null"`
	Avatar   string `json:"avatar" gorm:"column:avatar;not null"`
}

func (Article) TableName() string {
	return "articles"
}

func (a *Article) Create() error {

	return db.DB.Self.Create(&a).Error

}

func (a *Article) Update(id int64, userId int64) error {

	return db.DB.Self.Model(&a).Where("id = ?", id).Where("user_id = ?", userId).Updates(&a).Error

}

func (a *Article) Delete(id int64, userId int64) error {

	return db.DB.Self.Where("id = ?", id).Where("user_id = ?", userId).Delete(&a).Error

}

func ArticleViewAddOne(id int64) error {

	article := &Article{}

	res := db.DB.Self.Model(&article).
		Where("id = ?", id).
		Where("view > 0").
		UpdateColumn("view", gorm.Expr("view + ?", 1))

	return res.Error

}

func ArticleLikeAddOne(id int64) error {

	article := &Article{}

	res := db.DB.Self.Model(&article).
		Where("id = ?", id).
		Where("likes > 0").
		UpdateColumn("likes", gorm.Expr("likes + ?", 1))

	return res.Error

}

func ArticleLikeCutOne(id int64) error {

	article := &Article{}

	res := db.DB.Self.Model(&article).
		Where("id = ?", id).
		Where("likes > 0").
		UpdateColumn("likes", gorm.Expr("likes - ?", 1))

	return res.Error

}

func GetArticleById(id, userId int64) (*Article, error) {

	article := &Article{}

	res := db.DB.Self.
		Where("articles.id = ?", id).
		Joins("LEFT JOIN users AS u ON u.id=articles.user_id").
		Select("articles.*,u.nickname as nickname,u.head_url as avatar").
		First(&article)

	isFollowed := GetUserFollowById(userId, article.UserId)

	status := 0
	if isFollowed {
		status = 1
	}

	article.IsFollowed = int64(status)

	return article, res.Error

}

func GetSelfArticleById(id int64, userId int64) (*Article, error) {

	article := &Article{}

	res := db.DB.Self.Where("id = ?", id).Where("user_id = ?", userId).First(&article)

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
		res = res.Order("articles.created_at desc")
	} else {
		res = res.Order("articles.created_at asc")
	}

	res = res.
		Joins("LEFT JOIN users AS u ON u.id=articles.user_id").
		Select("articles.id,articles.user_id,articles.title,articles.thumb,articles.likes").
		Limit(pageSize).
		Offset(offset).
		Find(&articles)
	db.DB.Self.Model(&articles).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return articles, total, lastPage, res.Error
}

func UserArticlePaginate(page int64, pageSize int64, userId int64) (articles []*Article, total int64, lastPage int64, err error) {

	articles = make([]*Article, 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self.Where("user_id = ?", userId)

	res = res.Order("created_at desc")

	res = res.Limit(pageSize).Offset(offset).Find(&articles)
	db.DB.Self.Model(&articles).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return articles, total, lastPage, res.Error
}

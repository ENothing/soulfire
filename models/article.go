package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"soulfire/pkg/db"
	"soulfire/utils"
	"time"
)

type Article struct {
	Model
	UserId     int64      `json:"user_id" gorm:"column:user_id;not null"`
	Thumb      string     `json:"thumb" gorm:"column:thumb;not null"`
	Title      string     `json:"title" gorm:"column:title;not null"`
	Content    string     `json:"content" gorm:"column:content;not null"`
	Likes      int64      `json:"likes" gorm:"column:likes;not null"`
	Favor      int64      `json:"favor" gorm:"column:favor;not null"`
	View       int64      `json:"view" gorm:"column:view;not null"`
	CateId     int64      `json:"cate_id" gorm:"column:cate_id;not null"`
	IsPublish  int64      `json:"is_publish" gorm:"column:is_publish;not null"`
	IsFollowed int64      `json:"is_followed" gorm:"column:is_followed;not null"`
	CreatedAt  time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

type ArticleDetail struct {
	Article
	NickName        string `json:"nickname" gorm:"column:nickname;not null"`
	Avatar          string `json:"avatar" gorm:"column:avatar;not null"`
	Liked           bool   `json:"liked" gorm:"column:liked;not null"`
	Follows         int64  `json:"follows" gorm:"column:follows;not null"`
	CreatedAtFormat string `json:"created_at_format" gorm:"column:created_at_format"`
}

func (Article) TableName() string {
	return "articles"
}

func (a *ArticleDetail) AfterFind() (err error) {

	a.CreatedAtFormat = utils.TimeFormat(a.CreatedAt, 1)
	return
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
		Update("view", gorm.Expr("view + ?", 1))

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

func ArticleFavorAddOne(id int64) error {

	article := &Article{}

	res := db.DB.Self.Model(&article).
		Where("id = ?", id).
		Where("favor > 0").
		UpdateColumn("favor", gorm.Expr("favor + ?", 1))

	return res.Error

}

func ArticleFavorCutOne(id int64) error {

	article := &Article{}

	res := db.DB.Self.Model(&article).
		Where("id = ?", id).
		Where("favor > 0").
		UpdateColumn("favor", gorm.Expr("favor - ?", 1))

	return res.Error

}

func GetArticleById(id, userId int64) (*ArticleDetail, error) {

	article := &ArticleDetail{}

	res := db.DB.Self.
		Where("articles.id = ?", id).
		Joins("LEFT JOIN users AS u ON u.id=articles.user_id").
		Select("articles.*,u.nickname as nickname,u.head_url as avatar,u.follows as follows").
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

func ArticlePaginate(page int64, pageSize int64, sort int64, cateId int64, title string, userId int64) (articles []*ArticleDetail, total int64, lastPage int64, err error) {

	articles = make([]*ArticleDetail, 0)

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
		Joins("LEFT JOIN user_likes AS ul ON ul.type_id=articles.id AND ul.own=2 AND ul.user_id = ?", userId).
		Select("articles.id,articles.user_id,articles.title,articles.thumb,articles.likes,if(ul.id is null,false,true) as liked,u.nickname,u.head_url as avatar,articles.favor").
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

package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"soulfire/pkg/db"
	"time"
)

type ArticleComment struct {
	Model
	UserId          int64            `json:"user_id" gorm:"column:user_id;not null"`
	ArticleId       int64            `json:"article_id" gorm:"column:article_id;not null"`
	Content         string           `json:"content" gorm:"column:content;not null"`
	ParentId        int64            `json:"parent_id" gorm:"column:parent_id;not null"`
	ReplyId        int64            `json:"reply_id" gorm:"column:reply_id;not null"`
	CreatedAt       time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time        `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       *time.Time       `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	SubComments     []ArticleComment `json:"sub_comments" gorm:"foreignkey:parent_id;PRELOAD:false"`
	UserName        string           `json:"username" gorm:"column:username"`
	UserAvatar      string           `json:"user_avatar" gorm:"column:user_avatar"`
	ReplyUserName   string           `json:"reply_username" gorm:"column:reply_username"`
	ReplyUserAvatar string           `json:"reply_user_avatar" gorm:"column:reply_user_avatar"`
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

	res := db.DB.Self.
		Where("article_id = ?", articleId).
		Where("parent_id = ?", 0).
		Preload("SubComments", func(db *gorm.DB)*gorm.DB {
			return db.Joins("LEFT JOIN users as user ON user.id = article_comments.user_id " +
				"LEFT JOIN users as reply_user ON reply_user.id = article_comments.reply_id").
				Select("article_comments.*," +
					"user.nickname as username,user.head_url as user_avatar," +
					"reply_user.nickname as reply_username,reply_user.head_url as reply_user_avatar")
		}).
		Joins("LEFT JOIN users as user ON user.id = article_comments.user_id").
		Select("article_comments.*,user.nickname as username,user.head_url as user_avatar")

	res = res.Limit(pageSize).Offset(offset).Find(&articleComments)

	db.DB.Self.Model(&articleComments).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return articleComments, total, lastPage, res.Error
}

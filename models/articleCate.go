package models

import (
	"soulfire/pkg/db"
)

type ArticleCate struct {
	Model
	Name string `json:"name" gorm:"column:name;not null"`
}

func (ArticleCate) TableName() string {
	return "article_cates"
}


func GetArticleCates() ([]*ArticleCate,error) {

	articleCates := make([]*ArticleCate,0)

	res := db.DB.Self.Find(&articleCates)

	return articleCates,res.Error

}




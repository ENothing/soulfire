package models

import (
	"math"
	"soulfire/pkg/db"
	"time"
)

type ShopGoods struct {
	Model
	CateId int64 `json:"cate_id" gorm:"column:cate_id;not null"`
	BrandId int64 `json:"brand_id" gorm:"column:brand_id;not null"`
	Name string `json:"name" gorm:"column:name;not null"`
	Thumb string `json:"thumb" gorm:"column:thumb;not null"`
	Banners string `json:"banners" gorm:"column:banners;not null"`
	GoodsContent string `json:"goods_content" gorm:"column:goods_content;not null"`
	CurPrice float64 `json:"cur_price" gorm:"column:cur_price;not null"`
	OriPrice float64 `json:"ori_price" gorm:"column:ori_price;not null"`
	Stock int64 `json:"stock" gorm:"column:stock;not null"`
	Sold int64 `json:"sold" gorm:"column:sold;not null"`
	CreatedAt time.Time `gorm:";column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func ShopGoodsPaginate(page int64, pageSize int64, sort int64, cateId int64, title string) (shopGoods []*ShopGoods, total int64, lastPage int64, err error) {

	shopGoods = make([]*ShopGoods, 0)

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

	res = res.Limit(pageSize).Offset(offset).Find(&shopGoods)
	db.DB.Self.Model(&shopGoods).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return shopGoods, total, lastPage, res.Error
}
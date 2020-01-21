package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"math"
	"soulfire/pkg/db"
	"soulfire/utils"
	"time"
)

type ShopGoods struct {
	Model
	CateId          int64       `json:"cate_id" gorm:"column:cate_id;not null"`
	BrandId         int64       `json:"brand_id" gorm:"column:brand_id;not null"`
	Name            string      `json:"name" gorm:"column:name;not null"`
	Thumb           string      `json:"thumb" gorm:"column:thumb;not null"`
	Banners         string      `json:"banners" gorm:"column:banners;not null"`
	GoodsContent    string      `json:"goods_content" gorm:"column:goods_content;not null"`
	CurPrice        float64     `json:"cur_price" gorm:"column:cur_price;not null"`
	OriPrice        float64     `json:"ori_price" gorm:"column:ori_price;not null"`
	Stock           int64       `json:"stock" gorm:"column:stock;not null"`
	Sold            int64       `json:"sold" gorm:"column:sold;not null"`
	CreatedAt       time.Time   `gorm:";column:created_at" json:"created_at"`
	UpdatedAt       time.Time   `gorm:";column:updated_at" json:"updated_at"`
	PublishAt       time.Time   `gorm:";column:publish_at" json:"publish_at"`
	DeletedAt       *time.Time  `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
	DecodeBanners   interface{} `json:"decode_banners" gorm:"column:decode_banners"`
	PublishAtFormat string      `json:"publish_at_format" gorm:"column:publish_at_format"`
}

type Banners struct {
}

func (sg *ShopGoods) AfterFind() (err error) {

	var decodeBanner interface{}
	json.Unmarshal([]byte(sg.Banners), &decodeBanner)
	sg.DecodeBanners = decodeBanner
	sg.PublishAtFormat = utils.TimeFormat(sg.PublishAt, 1)
	return
}

func (ShopGoods) TableName() string {
	return "shop_goods"
}

func CutGoodsStockAndAddSold(goodsId, num int64, transaction *gorm.DB) error {

	shopGoods := &ShopGoods{}

	res := transaction.Model(&shopGoods).
		Where("id = ?", goodsId).
		Where("stock >= ?", num).
		Update("stock", gorm.Expr("stock - ?", num)).
		Update("sold", gorm.Expr("sold + ?", num))

	return res.Error

}

func GetShopGoodsById(id int64) (*ShopGoods, error) {

	shopGoods := &ShopGoods{}

	res := db.DB.Self.Where("id = ? ", id).First(&shopGoods)

	return shopGoods, res.Error

}

func ShopGoodsPaginate(page int64, pageSize int64, sortType int64, sort int64, name string, cateId int64, brandId int64) (shopGoods []*ShopGoods, total int64, lastPage int64, err error) {

	shopGoods = make([]*ShopGoods, 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self

	if cateId != 0 {
		res = res.Where("cate_id = ?", cateId)
	}
	if brandId != 0 {
		res = res.Where("brand_id = ?", brandId)
	}
	if name != "" {
		res = res.Where("name LIKE ?", "%"+name+"%")
	}

	switch sortType {
	case 1: //销量
		res = res.Order("sold desc")
		break
	case 2: //价格
		if sort == 0 {
			res = res.Order("cur_price asc")
		} else {
			res = res.Order("cur_price desc")
		}
		break
	case 3: //新品
		res = res.Order("created_at desc")
		break
	}

	res = res.Limit(pageSize).Offset(offset).Find(&shopGoods)
	db.DB.Self.Model(&shopGoods).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return shopGoods, total, lastPage, res.Error
}

package models

type ShopGoodsBrand struct {
	Model
	Name string `json:"name" gorm:"column:name;not null"`
	CateId string `json:"cate_id" gorm:"column:cate_id;not null"`
	Logo string `json:"logo" gorm:"column:logo;not null"`
}

func (ShopGoodsBrand) TableName() string {
	return "shop_goods_brands"
}



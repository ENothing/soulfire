package models

import (
	"soulfire/pkg/db"
)

type ShopOrderDelivery struct {
	Model
	OrderId      int64  `json:"order_id" gorm:"column:order_id;not null" `
	Abbreviation string `json:"abbreviation" gorm:"column:abbreviation;not null"  `
	Name         string `json:"name" gorm:"column:name"`
	DeliveryN   string `json:"delivery_n" gorm:"column:delivery_n"`
}

func (ShopOrderDelivery) TableName() string {
	return "shop_order_deliveries"
}


func GetDeliveryById(orderId int64) (*ShopOrderDelivery, error) {

	shopOrderDelivery := &ShopOrderDelivery{}

	res := db.DB.Self.Where("order_id = ?", orderId).First(&shopOrderDelivery)

	return shopOrderDelivery, res.Error

}





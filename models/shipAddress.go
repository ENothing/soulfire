package models

import (
	"math"
	"soulfire/pkg/db"
	"time"
)

type ShipAddress struct {
	Model
	UserId        int64     `json:"user_id" gorm:"column:user_id;not null" `
	Name          string     `json:"name" gorm:"column:name;not null"  `
	Mobile        string     `json:"mobile" gorm:"column:mobile"`
	Province      string     `json:"province" gorm:"column:province"`
	City          string     `json:"city" gorm:"column:city"`
	District      string     `json:"district" gorm:"column:district"`
	DetailAddress string     `json:"detail_address" gorm:"column:detail_address"`
	IsDefault     int64     `json:"is_default" gorm:"column:is_default"`
	CreatedAt     time.Time  `gorm:";column:created_at" json:"created_at"`
	UpdatedAt     time.Time  `gorm:";column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

func (ShipAddress) TableName() string {
	return "shipping_addresses"
}

func (sa *ShipAddress)Create() error {

	return db.DB.Self.Create(&sa).Error

}



func GetDefaultAddress(userId int64) (*ShipAddress,error) {

	defaultShipAddress := &ShipAddress{}

	res := db.DB.Self.
		Where("user_id = ?",userId).
		Where("is_default",1).
		First(&defaultShipAddress)


	return defaultShipAddress,res.Error

}

func AddressPaginate(page int64, pageSize int64, userId int64) (shipAddresses []*ShipAddress, total int64, lastPage int64, err error) {

	shipAddresses = make([]*ShipAddress, 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self.Where("user_id = ?", userId).Limit(pageSize).Offset(offset).Find(&shipAddresses)
	db.DB.Self.Model(&shipAddresses).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return shipAddresses, total, lastPage, res.Error

}

package models

import (
	"soulfire/pkg/db"
	"soulfire/utils"
	"time"
)

type UserCoupon struct {
	Model
	UserId      int64     `json:"user_id" gorm:"column:user_id;not null"`
	CouponId    int64     `json:"coupon_id" gorm:"column:coupon_id;not null"`
	ReceiveTime time.Time `json:"receive_time" gorm:"column:receive_time;not null"`
	EndTime     time.Time `json:"end_time" gorm:"column:end_time;not null"`
	IsUsed      int64     `json:"is_used" gorm:"column:is_used;not null"`
	Coupon      Coupon    `json:"coupon" gorm:"column:coupon;not null"`
}

func (UserCoupon) TableName() string {
	return "user_coupons"
}

func GetCanUseCouponCountById(userId int64, goodsId int64) (count int64) {

	nowTime := utils.TimeFormat(time.Now(), 0)

	err := db.DB.Self.Model(&UserCoupon{}).
		Where("is_used = ?", 0).
		Where("user_id = ?", userId).
		Where("end_time >= ?", nowTime).
		Where("is_goods = 1 AND FIND_IN_SET(?,goods_ids)", goodsId).
		Or("is_goods = 0").
		Joins("LEFT JOIN coupons as c on c.id = user_coupons.coupon_id").
		Select("user_coupons.*,coupons.* as coupon").
		Count(&count).Error

	if err != nil {

		return 0

	}

	return count

}

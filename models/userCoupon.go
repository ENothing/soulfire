package models

import (
	"math"
	"soulfire/pkg/db"
	"soulfire/utils"
	"time"
)

type UserCoupon struct {
	Model
	UserId           int64     `json:"user_id" gorm:"column:user_id;not null"`
	CouponId         int64     `json:"coupon_id" gorm:"column:coupon_id;not null"`
	ReceiveTime      time.Time `json:"receive_time" gorm:"column:receive_time;not null"`
	EndTime          time.Time `json:"end_time" gorm:"column:end_time;not null"`
	IsUsed           int64     `json:"is_used" gorm:"column:is_used;not null"`
	Coupon           Coupon    `json:"coupon" gorm:"foreignkey:coupon_id;PRELOAD:false"`
	CouponType       int64     `json:"coupon_type" gorm:"column:coupon_type;not null"`
	FullPrice        float64   `json:"full_price" gorm:"column:full_price;not null"`
	ReductionPrice   float64   `json:"reduction_price" gorm:"column:reduction_price;not null"`
	ImmediatelyPrice float64   `json:"immediately_price" gorm:"column:immediately_price;not null"`
	Discount         float64   `json:"discount" gorm:"column:discount;not null"`
}

func (UserCoupon) TableName() string {
	return "user_coupons"
}

func GetUserCouponById(userId, goodsId, couponId int64) (map[string]interface{}, error) {
	userCoupon := &UserCoupon{}
	nowTime := utils.TimeFormat(time.Now(), 0)

	res := db.DB.Self.
		Joins("LEFT JOIN coupons as c on c.id = user_coupons.coupon_id").
		Where("user_id = ?", userId).
		Where("coupon_id = ?", couponId).
		Where("is_used = ?", 0).
		Where("end_time >= ?", nowTime).
		Where("(is_goods = 1 AND FIND_IN_SET(?,goods_ids)) OR is_goods = 0", goodsId).
		Select("user_coupons.*,c.coupon_type as coupon_type,c.full_price as full_price,c.reduction_price as reduction_price,c.immediately_price as immediately_price,c.discount as discount").
		First(&userCoupon)

	data := make(map[string]interface{}, 0)
	data["coupon_type"] = userCoupon.CouponType
	data["full_price"] = userCoupon.FullPrice
	data["reduction_price"] = userCoupon.ReductionPrice
	data["immediately_price"] = userCoupon.ImmediatelyPrice
	data["discount"] = userCoupon.Discount

	return data, res.Error

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
		Select("user_coupons.*").
		Count(&count).Error

	if err != nil {

		return 0

	}

	return count

}

func CanUseCouponsPaginate(page int64, pageSize int64, userId, goodsId int64) (userCoupon []*UserCoupon, total int64, lastPage int64, err error) {

	userCoupon = make([]*UserCoupon, 0)
	nowTime := utils.TimeFormat(time.Now(), 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self.
		Where("is_used = ?", 0).
		Where("user_id = ?", userId).
		Where("end_time >= ?", nowTime).
		Where("is_goods = 1 AND FIND_IN_SET(?,goods_ids)", goodsId).
		Or("is_goods = 0").
		Joins("LEFT JOIN coupons as c on c.id = user_coupons.coupon_id").
		Preload("Coupon").
		Limit(pageSize).
		Offset(offset).
		Find(&userCoupon)

	db.DB.Self.Model(&userCoupon).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return userCoupon, total, lastPage, res.Error

}

func UserCouponsPaginate(page int64, pageSize int64, userId, status int64) (userCoupon []*UserCoupon, total int64, lastPage int64, err error) {

	userCoupon = make([]*UserCoupon, 0)
	nowTime := utils.TimeFormat(time.Now(), 0)

	offset := (page - 1) * pageSize

	res := db.DB.Self.Where("user_id = ?", userId)

	switch status {
	case 0:
		res = res.Where("is_used = ?", 0).Where("end_time >= ?", nowTime)
		break
	case 1:
		res = res.Where("is_used = ?", 1)
		break
	case 2:
		res = res.Where("is_used = ?", 0).Where("end_time < ?", nowTime)
		break
	default:
		res = res.Where("is_used = ?", 0).Where("end_time >= ?", nowTime)
	}

	res = res.Joins("LEFT JOIN coupons as c on c.id = user_coupons.coupon_id").
		Preload("Coupon").Limit(pageSize).Offset(offset).Find(&userCoupon)

	db.DB.Self.Model(&userCoupon).Count(&total)

	lastPage = int64(math.Ceil(float64(total) / float64(pageSize)))

	return userCoupon, total, lastPage, res.Error

}

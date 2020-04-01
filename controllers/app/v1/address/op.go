package address

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"soulfire/pkg/verify"
	"strconv"
)

type AdForm struct {
	Name          string `json:"name"  valid:"Required" ch:"姓名"`
	Mobile        string `json:"mobile" valid:"Required;Mobile" ch:"手机号"`
	Province      string `json:"province" valid:"Required;" ch:"省"`
	City          string `json:"city" valid:"Required" ch:"市"`
	District      string `json:"district" valid:"Required" ch:"区/县"`
	DetailAddress string `json:"detailAddress" valid:"Required" ch:"详细地址"`
}

func AddAddress(ctx *gin.Context) {

	var isDefault int64

	userId,_ := ctx.MustGet("user_id").(int64)
	name := ctx.PostForm("name")
	mobile := ctx.PostForm("mobile")
	province := ctx.PostForm("province")
	city := ctx.PostForm("city")
	district := ctx.PostForm("district")
	detailAddress := ctx.PostForm("detail_address")
	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	adForm := AdForm{
		name,
		mobile,
		province,
		city,
		district,
		detailAddress,
	}

	message := verify.FormVerify(&adForm)

	if message != nil {

		rsp.JsonResonse(ctx, rsp.AddAddressRequired, nil, message.(string))
		return

	}

	_, err := models.GetDefaultAddress(userId)

	if err == gorm.ErrRecordNotFound {
		isDefault = 1
	} else {
		isDefault = 0
	}

	address := models.ShipAddress{
		UserId:        userId,
		Name:          name,
		Mobile:        mobile,
		Province:      province,
		City:          city,
		District:      district,
		DetailAddress: detailAddress,
		IsDefault:     isDefault,
	}

	err = address.Create()

	if err != nil {

		rsp.JsonResonse(ctx, rsp.AddAddressFailed, nil, "")
		return

	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func UpdateAddress(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	name := ctx.PostForm("name")
	mobile := ctx.PostForm("mobile")
	province := ctx.PostForm("province")
	city := ctx.PostForm("city")
	district := ctx.PostForm("district")
	detailAddress := ctx.PostForm("detail_address")
	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	_, err := models.GetAddressById(id, userId)
	if err != nil {

		rsp.JsonResonse(ctx, rsp.AddressNotExits, nil, "")

	}

	adForm := AdForm{
		name,
		mobile,
		province,
		city,
		district,
		detailAddress,
	}

	message := verify.FormVerify(&adForm)

	if message != nil {

		rsp.JsonResonse(ctx, rsp.AddAddressRequired, nil, message.(string))
		return

	}

	address := models.ShipAddress{
		Name:          name,
		Mobile:        mobile,
		Province:      province,
		City:          city,
		District:      district,
		DetailAddress: detailAddress,
	}

	err = address.Update(id, userId)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.EditAddressFailed, nil, "")
		return

	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func UpdateDefaultAddress(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	err := models.UpdateDefaultAddress(id, userId)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.UpdateDefaultAddressFailed, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func DelAddress(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	shipAddress := models.ShipAddress{}

	err := shipAddress.Delete(id, userId)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.AddressDeleteFailed, nil, "")
		return

	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

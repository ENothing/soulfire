package address

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"soulfire/pkg/verify"
)

type AdForm struct {
	Name          string `json:"name"  valid:"Required" ch:"姓名"`
	Mobile        string `json:"mobile" valid:"Required" ch:"手机号"`
	Province      string `json:"province" valid:"Required;" ch:"省"`
	City          string `json:"city" valid:"Required" ch:"市"`
	District      string `json:"district" valid:"Required" ch:"区/县"`
	DetailAddress string `json:"detailAddress" valid:"Required" ch:"详细地址"`
}

func AddAddress(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	name := ctx.PostForm("name")
	mobile := ctx.PostForm("mobile")
	province := ctx.PostForm("province")
	city := ctx.PostForm("city")
	district := ctx.PostForm("district")
	detailAddress := ctx.PostForm("detail_address")

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


	defaultAddress,err := models.GetDefaultAddress(userId)

	fmt.Println(defaultAddress)
	fmt.Println(err)



	//address := models.ShipAddress{
	//	UserId:userId,
	//	Name:name,
	//	Mobile:mobile,
	//	Province:province,
	//	City:city,
	//	District:district,
	//	DetailAddress:detailAddress,
	//	IsDefault:
	//}
	//
	//
	//err = address.Create()
	//
	//if err != nil {
	//
	//	rsp.JsonResonse(ctx,rsp.AddAddressFailed,nil,"")
	//	return
	//
	//}

	rsp.JsonResonse(ctx,rsp.OK,nil,"")

}

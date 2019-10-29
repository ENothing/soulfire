package activity

import (
	"github.com/gin-gonic/gin"
	"soulfire/pkg/rsp"
	"soulfire/utils"
)

type EnterForm struct {
	Id int64 `json:"id"`
	Fullname string `json:"fullname"  valid:"Required"`
	Gender int64 `json:"gender"`
	Mobile string `json:"mobile" valid:"Required"`
	SmsCode string `json:"sms_code" valid:"Required"`
}

func Enter(ctx *gin.Context) {

	utils.Uid("AO",12)

	//userId := ctx.MustGet("user_id")
	//
	//id,_ := strconv.ParseInt(ctx.PostForm("id"),10,64)
	//Name := ctx.PostForm("name")
	//gender,_ := strconv.ParseInt(ctx.PostForm("gender"),10,64)
	//mobile := ctx.PostForm("mobile")
	//smsCode := ctx.PostForm("sms_code")
	//
	//enter := EnterForm{id,Name,gender,mobile,smsCode}
	//
	//valid := validation.Validation{}
	//
	//valid.Required(enter.Fullname,"fullname").Message("姓名不能为空哦~")
	//valid.Required(enter.Mobile,"mobile").Message("手机号不能为空哦~")
	//valid.Mobile(enter.Mobile,"mobile").Message("手机号格式不正确哦~")
	//valid.Required(enter.SmsCode,"sms_code").Message("验证码不能为空哦~")
	//
	//if valid.HasErrors() {
	//
	//	rsp.JsonResonse(ctx,rsp.EnterActivityRequired ,nil,valid.Errors[0].String())
	//
	//}
	//
	////todo 调起支付
	//
	//activity,err := models.GetActivityById(id)
	//
	//if err != nil {
	//
		rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil,"")
	//
	//}
	//
	//orderN :=
	//
	//models.ActivityOrder{
	//	UserId:userId,
	//	ActivityId:activity.Id,
	//	RefundId:0,
	//	OrderN:
	//	Name
	//	Sex
	//	Mobile
	//	UnitPrice
	//	PersonNum
	//	TotalPrice
	//	RealPrice
	//	DiscountPrice
	//	Status
	//}






































}

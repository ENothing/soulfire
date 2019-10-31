package activity

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

type EnterForm struct {
	Id int64 `json:"id"`
	Fullname string `json:"fullname"  valid:"Required"`
	Gender int64 `json:"gender"`
	Mobile string `json:"mobile" valid:"Required"`
	SmsCode string `json:"sms_code" valid:"Required"`
}

func (enter *EnterForm) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}



func Enter(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	fmt.Println(userId)

	id,_ := strconv.ParseInt(ctx.PostForm("id"),10,64)
	name := ctx.PostForm("name")
	gender,_ := strconv.ParseInt(ctx.PostForm("gender"),10,64)
	mobile := ctx.PostForm("mobile")
	smsCode := ctx.PostForm("sms_code")
	//personNum,_ := strconv.ParseInt(ctx.PostForm("person_num"),10,64)

	enter := EnterForm{id,name,gender,mobile,smsCode}

	valid := validation.Validation{}
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
	//
	//


	b, err := valid.Valid(&enter)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}





	//todo 调起支付

	//activity,err := models.GetActivityById(id)
	//fmt.Println(activity)
	//
	//if err != nil {
	//
	//	rsp.JsonResonse(ctx, rsp.ActivityNotExits, userId,"")
	//	return
	//}

	//orderN := utils.Uid("AO")
	//
	//models.ActivityOrder{
	//	UserId:userId,
	//	ActivityId:int64(activity.Id) ,
	//	RefundId:0,
	//	OrderN:orderN,
	//	Name:name,
	//	Sex:gender,
	//	Mobile:mobile,
	//	UnitPrice:,
	//	PersonNum:1,
	//	TotalPrice:,
	//	RealPrice:,
	//	DiscountPrice:,
	//	Status:,
	//}






































}

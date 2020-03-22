package activity

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"soulfire/pkg/verify"
	"soulfire/utils"
	"strconv"
)

type EnterForm struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"  valid:"Required" ch:"姓名"`
	Gender    int64  `json:"gender"`
	Mobile    string `json:"mobile" valid:"Required;Mobile" ch:"手机号"`
	SmsCode   string `json:"sms_code" valid:"Required" ch:"验证码"`
	CNum      string `json:"c_num" valid:"Required" ch:"证件号码"`
}

/**
活动报名入口
*/
func Enter(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)

	id, _ := strconv.ParseInt(ctx.PostForm("id"), 10, 64)
	name := ctx.PostForm("name")
	gender, _ := strconv.ParseInt(ctx.DefaultPostForm("gender", "1"), 10, 64)
	cType, _ := strconv.ParseInt(ctx.DefaultPostForm("c_type", "1"), 10, 64)
	cNum := ctx.PostForm("c_num")
	mobile := ctx.PostForm("mobile")
	smsCode := ctx.PostForm("sms_code")
	personNum, _ := strconv.ParseInt(ctx.DefaultPostForm("person_num", "1"), 10, 64)
	//couponId, _ := strconv.ParseInt(ctx.PostForm("coupon_id"), 10, 64)

	if userId == 0 {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	enter := EnterForm{
		id,
		name,
		gender,
		mobile,
		smsCode,
		cNum,
	}

	message := verify.FormVerify(&enter)
	if message != nil {
		rsp.JsonResonse(ctx, rsp.EnterActivityRequired, nil, message.(string))
		return
	}

	activity, err := models.GetActivityById(id)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.ActivityNotExits, nil, "")
		return
	}
	if activity.PersonLimit == activity.Sold {
		rsp.JsonResonse(ctx, rsp.ActivityPersonIsFull, nil, "")
		return
	}

	orderN := utils.Uid("AO")

	//todo 优惠券的使用
	discountPrice := 0.00

	totalPrice := activity.CurPrice * float64(personNum)
	realPrice := totalPrice - discountPrice

	activityOrder := models.ActivityOrder{
		UserId:        userId,
		ActivityId:    activity.Id,
		OrderN:        orderN,
		Name:          name,
		Sex:           gender,
		Mobile:        mobile,
		UnitPrice:     activity.CurPrice,
		PersonNum:     personNum,
		TotalPrice:    totalPrice,
		RealPrice:     realPrice,
		DiscountPrice: discountPrice,
		CType:         cType,
		CNum:          cNum,
	}

	id,err = activityOrder.Create()
	if err != nil {
		rsp.JsonResonse(ctx, rsp.CreateActivityOrderFaild, nil, "")
		return
	}

	_ = models.ActivitySoldOne(activity.Id)

	rsp.JsonResonse(ctx, rsp.OK, id, "")

}

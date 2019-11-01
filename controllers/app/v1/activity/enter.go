package activity

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/logging"
	"soulfire/pkg/rsp"
	"soulfire/pkg/verify"
	"soulfire/utils"
	"strconv"
)

type EnterForm struct {
	Id        int64  `json:"id"`
	Fullname  string `json:"fullname"  valid:"Required" ch:"姓名"`
	Gender    int64  `json:"gender"`
	Mobile    string `json:"mobile" valid:"Required;Mobile" ch:"手机号"`
	SmsCode   string `json:"sms_code" valid:"Required" ch:"验证码"`
	PersonNum int64  `json:"person_num" valid:"Required" ch:"人数"`
}

func Enter(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)

	id, _ := strconv.ParseInt(ctx.PostForm("id"), 10, 64)
	name := ctx.PostForm("name")
	gender,_ := strconv.ParseInt(ctx.DefaultPostForm("gender", "1"),10,64)
	mobile := ctx.PostForm("mobile")
	smsCode := ctx.PostForm("sms_code")
	personNum, _ := strconv.ParseInt(ctx.PostForm("person_num"), 10, 64)
	fmt.Println(gender)
	enter := EnterForm{
		id,
		name,
		gender,
		mobile,
		smsCode,
		personNum,
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

	orderN := utils.Uid("AO")

	//todo 优惠券的使用
	discountPrice := 0.00

	totalPrice := activity.CurPrice * float64(personNum)
	realPrice := totalPrice - discountPrice

	activityOrder := models.ActivityOrder{
		UserId:     userId,
		ActivityId: int64(activity.Id),
		OrderN:     orderN,
		Name:       name,
		Sex:        gender,
		Mobile:     mobile,
		UnitPrice:  activity.CurPrice,
		PersonNum:  personNum,
		TotalPrice: totalPrice,
		RealPrice: realPrice,
		DiscountPrice: discountPrice,
	}

	err = activityOrder.Create()

	if err != nil {

		logging.Logging(logging.ERR, err)

		rsp.JsonResonse(ctx, rsp.CreateActivityOrderFaild, nil, "")
		return

	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

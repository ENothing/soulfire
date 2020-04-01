package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/silenceper/wechat"
	"soulfire/models"
	"soulfire/pkg/config"
	"soulfire/pkg/rsp"
	jwt "soulfire/pkg/token"
)

func Login(ctx *gin.Context) {

	code := ctx.PostForm("code")
	iv := ctx.PostForm("iv")
	encryptedData := ctx.PostForm("encryptedData")

	var userId int64
	var nickName string

	if code == "" {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	app, _ := config.Cfg.GetSection("wechat")
	wc := wechat.NewWechat(&wechat.Config{
		AppID:     app.Key("APPID").String(),
		AppSecret: app.Key("SECRET").String(),
	})

	wxa := wc.GetMiniProgram()

	data, err := wxa.Code2Session(code)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.LoginFailed, nil, "")
		return
	}

	user, err := models.GetUserByOpenid(data.OpenID)
	if err != nil && err != gorm.ErrRecordNotFound {
		rsp.JsonResonse(ctx, rsp.DatabaseErr, nil, "")
		return
	}

	if err == gorm.ErrRecordNotFound {

		userInfo, err := wxa.Decrypt(data.SessionKey, encryptedData, iv)
		if err != nil {
			rsp.JsonResonse(ctx, rsp.LoginFailed, nil, "")
			return
		}

		user := models.User{
			Openid:   userInfo.OpenID,
			HeadUrl:  userInfo.AvatarURL,
			NickName: userInfo.NickName,
			Gender:   int64(userInfo.Gender),
		}

		userId, err = user.Create()

		if err != nil {
			rsp.JsonResonse(ctx, rsp.UserCreateFailed, nil, "")
			return
		}

		nickName = userInfo.NickName

	} else {

		userId = user.Id
		nickName = user.NickName
	}

	userToken := jwt.UserToken{
		userId,
		nickName,
		data.OpenID,
		data.SessionKey,
	}

	token, err := jwt.Encode(userToken)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.GenerateTokenErr, nil, "")

	} else {

		rsp.JsonResonse(ctx, rsp.OK, token, "")

	}

}

func Info(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)

	data := make(map[string]interface{})

	user, _ := models.GetUserInfoById(userId)

	activityOrderUnpayCount := models.GetActivityOrderUnpayCount(userId)
	shopOrderUnpayCount := models.GetShopOrderUnpayCount(userId)

	data["user_info"] = user
	data["aorder_unpay"] = activityOrderUnpayCount
	data["sorder_unpay"] = shopOrderUnpayCount

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

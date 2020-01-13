package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"soulfire/models"
	"soulfire/pkg/rsp"
	jwt "soulfire/pkg/token"
	"soulfire/pkg/wechat"
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

	wc := wechat.Wc{}

	data := wc.Login().Code2Session(code)

	if data["errcode"] != nil && data["errcode"] != int64(0) {
		rsp.JsonResonse(ctx, rsp.LoginFailed, nil, (data["errmsg"]).(string))
		return
	}

	openid := (data["openid"]).(string)
	sessionKey := (data["session_key"]).(string)

	user, err := models.GetUserByOpenid(openid)
	if err != nil && err != gorm.ErrRecordNotFound {
		rsp.JsonResonse(ctx, rsp.DatabaseErr, nil, "")
		return
	}
	if err == gorm.ErrRecordNotFound {

		userInfo := wc.Decrypt().UserInfo(sessionKey, encryptedData, iv)
		if userInfo == nil {
			rsp.JsonResonse(ctx, rsp.UserInfoGotFailed, nil, "")
			return
		}

		user := models.User{
			Openid:   openid,
			HeadUrl:  userInfo.HeadUrl,
			NickName: userInfo.NickName,
			Gender:   userInfo.Gender,
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
		openid,
		sessionKey,
	}

	token, err := jwt.Encode(userToken)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.GenerateTokenErr, nil, "")

	} else {

		rsp.JsonResonse(ctx, rsp.OK, token, "")

	}

}

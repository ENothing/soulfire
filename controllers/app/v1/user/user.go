package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/silenceper/wechat"
	"github.com/spf13/viper"
	"soulfire/models"
	"soulfire/pkg/auth"
	"soulfire/pkg/rsp"
	"time"
)

func Login(ctx *gin.Context) {

	code := ctx.PostForm("code")
	iv := ctx.PostForm("iv")
	encryptedData := ctx.PostForm("encryptedData")

	var userId int64

	if code == "" {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	wc := wechat.NewWechat(&wechat.Config{
		AppID:     viper.GetString("Wechat.AppId"),
		AppSecret: viper.GetString("Wechat.Secret"),
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

	} else {

		userId = user.Id
	}

	token,err := generateToken(userId)


	if err != nil {

		rsp.JsonResonse(ctx, rsp.GenerateTokenErr, nil, "")

	} else {

		rsp.JsonResonse(ctx, rsp.OK, token, "")

	}

}

func Info(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)

	data := make(map[string]interface{})

	user, _ := models.GetUserInfoById(userId)

	activityOrderUnpayCount := models.GetActivityOrderUnpayCount(userId)
	shopOrderUnpayCount := models.GetShopOrderUnpayCount(userId)

	data["user_info"] = user
	data["aorder_unpay"] = activityOrderUnpayCount
	data["sorder_unpay"] = shopOrderUnpayCount

	rsp.JsonResonse(ctx, rsp.OK, data, "")

}

func generateToken(userId int64)(string,error) {

	j := auth.NewJWT()
	claims := auth.CustomClaims{
		userId,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 28800), // 过期时间 8小时
			Issuer:    "en",
		},
	}

	token, err:= j.CreateToken(claims)

	return token, err
}

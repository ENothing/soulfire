package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"soulfire/pkg/rsp"
	"soulfire/pkg/wechat"
)

func Login(ctx *gin.Context) {

	code := ctx.PostForm("code")

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
	fmt.Println(data)

	//_,err := models.GetUserByOpenid((data["openid"]).(string))
	//if err != nil {
	//	rsp.JsonResonse(ctx, rsp.DatabaseErr, nil,"")
	//	return
	//}
	//if err == gorm.ErrRecordNotFound {
	//
	//	//user := models.User{
	//	//
	//	//}
	//
	//
	//
	//
	//}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

	//userToken := jwt.UserToken{
	//	1,
	//	"test",
	//	"123",
	//	"321",
	//}
	//
	//token, err := jwt.Encode(userToken)
	//
	//if err != nil {
	//
	//	rsp.JsonResonse(ctx, rsp.GenerateTokenErr, nil,"")
	//
	//}else{
	//
	//	rsp.JsonResonse(ctx, rsp.OK, token,"")
	//
	//}

}

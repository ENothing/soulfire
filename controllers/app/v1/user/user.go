package user

import (
	"github.com/gin-gonic/gin"
	"soulfire/pkg/rsp"
	"soulfire/pkg/wechat"
)

func Login(ctx *gin.Context) {

	code := ctx.PostForm("code")

	if code != "" {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	data := wechat.Login().Code2Session(code)

	//if data["errcode"] == "0" {

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

	//}else{
	//
	rsp.JsonResonse(ctx, rsp.LoginFailed, data, "")
	//
	//}

}

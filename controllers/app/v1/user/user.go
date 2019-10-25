package user

import (
	"gin-init/pkg/rsp"
	jwt "gin-init/pkg/token"
	"gin-init/pkg/wechat"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {



	data := wechat.Code2Session("123")

	if data["errcode"] == "0" {

		userToken := jwt.UserToken{
			1,
			"test",
			"123",
			"321",
		}

		token, err := jwt.Encode(userToken)

		if err != nil {

			rsp.JsonResonse(ctx, rsp.GenerateTokenErr, nil)

		}else{

			rsp.JsonResonse(ctx, rsp.OK, token)

		}

	}else{

		rsp.JsonResonse(ctx, rsp.LoginFailed, data)

	}

}

package middleware

import (
	"soulfire/pkg/rsp"
	jwt "soulfire/pkg/token"
	"github.com/gin-gonic/gin"
	"strings"
)

func Verify() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("Authorization")

		if len(token) == 0 {

			rsp.JsonResonse(ctx,rsp.InvalidToken,nil,"")
			ctx.Abort()
			return
		}

		token = strings.Fields(token)[1]

		userToken, err := jwt.Parse(token)

		if  err != nil{

			rsp.JsonResonse(ctx,rsp.InvalidToken,nil,"")
			ctx.Abort()
			return

		}

		ctx.Set("user_id",userToken.Id)

		ctx.Next()

	}

}
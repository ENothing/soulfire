package middleware

import (
	"github.com/gin-gonic/gin"
	"soulfire/pkg/auth"
	"soulfire/pkg/rsp"
	"strings"
)

func Verify() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("Authorization")
		if s := strings.Split(token, " "); len(s) == 2 {
			token = s[1]
		}
		if token == "" || token == "Bearer" {
			ctx.Set("user_id", "0")
		} else {
			j := auth.NewJWT()
			claims, err := j.ParseToken(token)

			if err != nil {
				if err == auth.TokenExpired {
					rsp.JsonResonse(ctx, rsp.TokenExpried, nil, "")
					ctx.Abort()
					return
				}
				rsp.JsonResonse(ctx, rsp.InvalidToken, nil, "")
				ctx.Abort()
				return
			}
			ctx.Set("user_id", claims.ID)
		}

	}

}

package middleware

import (
	"github.com/gin-gonic/gin"
	jwt "soulfire/pkg/token"
	"strings"
)

func Verify() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("Authorization")

		if len(token) == 0 {

			ctx.Set("user_id", int64(0))

		} else {

			token = strings.Fields(token)[1]

			userToken, err := jwt.Parse(token)

			if err != nil {

				ctx.Set("user_id", int64(0))

			} else {

				ctx.Set("user_id", userToken.Id)
			}

		}

		ctx.Next()

	}

}

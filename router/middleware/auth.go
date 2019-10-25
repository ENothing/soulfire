package middleware

import (
	"fmt"
	"soulfire/pkg/rsp"
	jwt "soulfire/pkg/token"
	"github.com/gin-gonic/gin"
	"strings"
)

func Verify() gin.HandlerFunc {

	return func(context *gin.Context) {

		token := context.Request.Header.Get("Authorization")

		if len(token) == 0 {

			rsp.JsonResonse(context,rsp.InvalidToken,nil)

		}

		token = strings.Fields(token)[1]
		fmt.Println(token)

		_, err := jwt.Parse(token)

		if  err != nil{

			rsp.JsonResonse(context,rsp.InvalidToken,nil)

		}

		context.Next()

	}

}
package user

import (
	"gin-init/pkg/rsp"
	jwt "gin-init/pkg/token"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context)  {


	 token,err := jwt.Encode()

	if err != nil {

		rsp.JsonResonse(c,rsp.GenerateTokenErr,nil)

	}

	rsp.JsonResonse(c,rsp.OK,token)
}


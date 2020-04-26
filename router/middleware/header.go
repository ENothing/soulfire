package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
跨域设置
*/
func Options(c *gin.Context) {

	//if c.Request.Method != "OPTIONS" {
	//	c.Next()
	//}
	//
	//c.Header("Access-Control-Allow-Origin", "*")
	//c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	//c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	////c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	////c.Header("Content-Type", "application/json")
	//c.AbortWithStatus(200)
	method := c.Request.Method

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()


}

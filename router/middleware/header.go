package middleware

import "github.com/gin-gonic/gin"

/**
跨域设置
*/
func Options(c *gin.Context) {

	if c.Request.Method != "OPTIONS" {
		c.Next()
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	//c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	//c.Header("Content-Type", "application/json")
	c.AbortWithStatus(200)

}

package rsp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}


func JsonResonse(c *gin.Context,rsp *Errno,data interface{},message string)  {

	if message != "" {

		rsp.Message = message

	}

	c.JSON(http.StatusOK,Response{
		Code:rsp.Code,
		Message:rsp.Message,
		Data:data,
	})
}

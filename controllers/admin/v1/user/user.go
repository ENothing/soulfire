package user

import (
	"fmt"
	"gin-init/models"
	"gin-init/pkg/rsp"
	"gin-init/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func Index(c *gin.Context) {

	rsp.JsonResonse(c, rsp.OK, nil)

}

func Create(c *gin.Context) {

	username := c.Param("username")
	password := c.Param("password")
	password = utils.Md5(password)

	user := models.User{Username: username, Password: password}

	_, err := user.Create()

	if err != nil {

		fmt.Println(err)
		rsp.JsonResonse(c, rsp.OK, err)

	}
	rsp.JsonResonse(c, rsp.OK, "")

}

func Edit(c *gin.Context) {

	rsp.JsonResonse(c, rsp.OK, c.Param("id"))

}

func Update(c *gin.Context)  {

	rsp.JsonResonse(c, rsp.OK,"")
}


/**
删除（支持批量）
 */
func Delete(c *gin.Context) {

	idsString := c.Param("ids")//"8,9,10"

	ids := strings.Split(idsString,",")

	err := models.Delete(ids)

	if err != nil {

		rsp.JsonResonse(c, rsp.DatabaseErr,"")

	}

	rsp.JsonResonse(c,rsp.OK,"")

}

package utils

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

/**
md5加密
*/
func Md5(password string) (md5str string) {

	data := []byte(password)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)

}

func Uid(prefix string) string {

	uid := uuid.Must(uuid.NewV4(), nil)

	return (prefix + uid.String())[0:24]

}

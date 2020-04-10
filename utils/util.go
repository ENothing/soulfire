package utils

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
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

func Code() string{

	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

}
